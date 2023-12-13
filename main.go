package main

import "ifcconv/lib/handler"

func main() {
	r := handler.NewRouter()
	r.Run(":8080")
}
