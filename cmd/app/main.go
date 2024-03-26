package main

import (
	"bcc-freepass-2023/internal/handler/rest"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	var (
		code int
		err  error
	)

	defer func() {
		if err != nil {
			log.Errorf("[bcc-freepass-2023] in main, err: %v", err)
		}
		os.Exit(code)
	}()

	errChan := make(chan error)
	codeChan := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		wg.Done()
		rest := rest.InitializeServer()
		code, err := rest.RunServer()
		if err != nil {
			errChan <- err
		}
		codeChan <- code
	}()
	wg.Wait()

	err = <-errChan
	code = <-codeChan
}
