package main

import (
	"github.com/PorcoGalliard/truck-toll-calculator/types"
	"github.com/gorilla/websocket"
	"log"
	"math"
	"math/rand"
	"time"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

var sendInterval = time.Second

func generateCoordinate() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return f + n
}

func generateLatLong() (float64, float64) {
	return generateCoordinate(), generateCoordinate()
}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func main() {
	obusID := generateOBUIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(obusID); i++ {
			lat, long := generateLatLong()
			data := types.OBUdata{
				OBUID: obusID[i],
				Lat:   lat,
				Long:  long,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval)
	}
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
