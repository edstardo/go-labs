package main

import (
	"fmt"
)

type DB struct {
	name string
	user string
	pass string
	host string
	port string
}

func NewDB(name, user, pass string, options ...DBOption) *DB {
	db := &DB{
		name: name,
		user: user,
		pass: pass,
		host: "localhost",
		port: "5432",
	}

	for _, opt := range options {
		opt(db)
	}

	return db
}

type DBOption func(*DB)

func WithHost(host string) DBOption {
	return func(db *DB) {
		db.host = host
	}
}

func WithPort(port string) DBOption {
	return func(db *DB) {
		db.port = port
	}
}

func main() {
	fmt.Println("Options Design Pattern")

	fmt.Println(NewDB("cxm", "cxm", "cxm"))
	fmt.Println(NewDB("cxm", "cxm", "cxm", WithHost("127.0.0.1")))
	fmt.Println(NewDB("cxm", "cxm", "cxm", WithPort("8080")))
	fmt.Println(NewDB("cxm", "cxm", "cxm", WithPort("8080"), WithHost("127.0.0.1")))
}
