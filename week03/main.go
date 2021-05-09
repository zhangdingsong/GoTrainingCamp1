package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

var CloseChan chan bool

//start an http server
func startHttpServer() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("SUCCEED IN GETTING RESPONSE.")
	})

	http.HandleFunc("/close", func(writer http.ResponseWriter, request *http.Request) {
		CloseChan <- true
	})

	err := http.ListenAndServe(":8101", nil)
	return err
}

//Close http server
func closeHttpServer() error {
	if <-CloseChan {
		return errors.New("QUIT.")
	}
	return nil
}

//Listen
func main() {

	group, _ := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return startHttpServer()
	})

	group.Go(func() error {
		return closeHttpServer()
	})

	if err := group.Wait(); err != nil {
		panic(err.Error())
	}

}
