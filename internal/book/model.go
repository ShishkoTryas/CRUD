package book

import "time"

type Book struct {
	Id    int       `json:id`
	Name  string    `json:name`
	Price int       `json:price`
	Time  time.Time `json:timestamp`
}

type BookInput struct {
	Name  string `json:name`
	Price int    `json:price`
}
