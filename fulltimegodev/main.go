package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{}
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("Server starting")
	s.loop()
}

func (s *Server) loop() {
mainloop:
	for {
		select {
		case <-s.quitch:
			fmt.Println("Quitting server")
			break mainloop
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default:
		}
	}
	fmt.Println("Server is shutting down gracefully")
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("We receive a message:", msg)
}

func (s *Server) quit() {
	s.quitch <- struct{}{}
}

func main() {
	server := newServer()

	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()

	server.start()
}
