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
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan types.OBUdata, 128),
	}
}

func main() {
	recv := NewDataReceiver()
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
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
		fmt.Printf("received OBU data [%d] :: <lat %.2f :: long %.2f> \n", data.OBUID, data.Lat, data.Long)
		dr.msgch <- data
	}
}
