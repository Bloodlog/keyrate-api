package main

import (
	"key-rate-api/src/Routes"
)

func main() {
	r := Routes.SetUpRouter()
	r.Run()
}
