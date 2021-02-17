package main

import (
	"log"
	"net"
	"os"
	"wuziqi/game"
)

func handleConnection(conn net.Conn) {
	A := game.NewMan(os.Stdin, os.Stdout, game.Black)
	B := game.NewMan(conn, conn, game.White)
	game.Prepare(A, B)
	doer, waiter := A, B
	for {
		if game.Round(doer, waiter) {
			game.Win(doer, waiter)
		}
		doer, waiter = waiter, doer
	}

}

func main() {
	// fmt.Println(table)
	ln, err := net.Listen("tcp", ":7777")
	log.Printf("Server listening on %v\n", ln.Addr())
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		log.Printf("Accecpt connection from %v\n", conn.RemoteAddr())
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
