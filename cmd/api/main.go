package main

import "log"

func main() {
	app := &application{
		config: loadConfig(),
	}

	log.Fatal(app.serve(app.mount()))
}
