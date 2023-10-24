package main

import "time"

type Movie struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Year      int32
	Runtime   int32
	Genres    []string
	Version   int32
}
