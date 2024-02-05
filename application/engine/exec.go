package engine

import (
	"github.com/alioth-center/infrastructure/exit"
	"github.com/alioth-center/infrastructure/network/rpc"
	"os"
	"sync"
)

func Exec(collection, servingAddress string) {
	sync.OnceFunc(func() {
		// use given collection name as output file
		outputFile := collection

		// if output file is not given, use first cli argument
		if outputFile == "" && len(os.Args) > 1 {
			outputFile = os.Args[1]
		}

		// if first cli argument is not given, use environment variable
		if outputFile == "" {
			outputFile = os.Getenv("LOG_FILE")
		}

		// if environment variable is not given, use default value
		if outputFile == "" {
			outputFile = "./restoration_collection.log"
		}

		// initialize restoration service
		log := NewLogger(outputFile)
		service := NewService(log)
		exitChan := make(chan struct{}, 3)

		// register exit handler
		exit.Register(func(sig string) string {
			exitChan <- struct{}{}
			return "restoration collection service is shutting down"
		}, "restoration collection service")

		// start serving
		engine := rpc.NewEngine()
		engine.AddService(service)
		engine.ServeAsync(servingAddress, exitChan)

		exit.BlockedUntilTerminate()
	})()
}
