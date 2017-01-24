package main

import "flag"

var host = flag.String("host","localhost:8765", "server address")

func main() {
	RunUi()
}