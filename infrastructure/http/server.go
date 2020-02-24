package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// ListenAndServe starts a new web server of a provided addr
func ListenAndServe(addr string) {
	srv := &http.Server{
		Handler: router(),
		Addr: addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15* time.Second,
		IdleTimeout: 10 * time.Second,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatalf("failed to start http server: %v", err)
		}
	}()

	<-interrupt
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Fatal(srv.Shutdown(ctx))
	log.Println("Server gracefully shutdown")


}