package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func serve(ctx context.Context, addr string) error {
	mux := http.NewServeMux()
	srv := &http.Server{Addr: addr, Handler: mux}
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello")
	})
	go func() {
		<-ctx.Done()
		shutdownctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		srv.Shutdown(shutdownctx)
	}()

	return srv.ListenAndServe()
}

func processSignal(ctx context.Context, cs chan os.Signal) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case s := <-cs:
			switch s {
			case syscall.SIGTERM, syscall.SIGINT:
				return errors.New("exit")
			default:
			}
		}
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return serve(ctx, "server1")
	})
	g.Go(func() error {
		return serve(ctx,"server2")
	})
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
	g.Go(func() error {
		return processSignal(ctx, s)
	})
	if err := g.Wait(); err != nil {
		fmt.Println("Server Error:%v\n", err)
	}
}