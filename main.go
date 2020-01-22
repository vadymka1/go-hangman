package main

import (
	"context"
	"github.com/go-hangman/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

type Server struct {
	http.Server
	shutdownReq chan bool
	reqCount uint32
}

var port = ":9090"

var word string

func StartServer() *Server {

	srv := &Server{
		Server:      http.Server{
			Addr         :port,
			ReadTimeout  : 10*time.Second,
			WriteTimeout : 10*time.Second,
		},
		shutdownReq: make(chan bool),
		//reqCount:    0,
	}

	router := mux.NewRouter()

	router.HandleFunc("/", services.ShowForm).Methods("GET")
	router.HandleFunc("/hangman", services.CheckData).Methods("POST")
	router.HandleFunc("/shutdown", srv.ShutdownHandler)

	srv.Handler = router

	return srv
}

func main()  {

	server := StartServer()

	word = services.GetWord(1)

	done := make(chan bool)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Listen and serve: %v", err)
		}
		log.Println("Server starting")
		done <- true
	}()

	//wait shutdown
	server.WaitShutdown()

	<-done
	log.Printf("DONE!")

}

func (s *Server) ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shutdown server"))

	//Do nothing if shutdown request already issued
	//if s.reqCount == 0 then set to 1, return true otherwise false
	if !atomic.CompareAndSwapUint32(&s.reqCount, 0, 1) {
		log.Printf("Shutdown through API call in progress...")
		return
	}

	go func() {
		s.shutdownReq <- true
	}()
}

func (s *Server) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		log.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-s.shutdownReq:
		log.Printf("Shutdown request (/shutdown %v)", sig)
	}

	log.Printf("Stoping http server ...")

	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		log.Printf("Shutdown request error: %v", err)
	}
}

