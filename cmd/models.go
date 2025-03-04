package main

import "time"

type Book struct {
	Id int64
	Name string
	Author string
	Published time.Time
	Pages int64
	Lang string
	Publisher string
}
