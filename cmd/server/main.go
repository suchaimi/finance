package main

import "net/http"

func main() {
	const addr = "0.0.0.0:8088"
	server := http.Server{
		Addr: addr,
	}
}
