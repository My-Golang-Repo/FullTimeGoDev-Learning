package main

import (
	"fmt"
	"github.com/PorcoGalliard/truck-toll-calculator/types"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type DataReceiver struct {
	msgch chan types.OBUdata
	conn  *websocket.Conn
	prod  DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {
	var (
		p          DataProducer
		err        error
		kafkaTopic = "obuData"
	)

	p, err = NewKafkaProducer(kafkaTopic)
	if err != nil {
		return nil, err
	}

	p = NewLogMiddleware(p)
	return &DataReceiver{
		msgch: make(chan types.OBUdata, 128),
		prod:  p,
	}, nil
}

func main() {
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
}

func (dr *DataReceiver) produceData(data types.OBUdata) error {
	return dr.prod.ProduceData(data)
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn

	go dr.wsReceiverLoop()
}

func (dr *DataReceiver) wsReceiverLoop() {
	fmt.Println("New OBU Client Connected")
	for {
		var data types.OBUdata
		if err := dr.conn.ReadJSON(&data); err != nil {
			fmt.Println("read error: ", err)
			continue
		}
		if err := dr.produceData(data); err != nil {
			fmt.Println("kafka produce error", err)
		}
	}
}
