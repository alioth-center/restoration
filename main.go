package main

import (
	"github.com/alioth-center/restoration/application/engine"
	"os"
)

func main() {
	engine.Exec(os.Getenv("LOG_FILE"), "0.0.0.0:28081")
}
