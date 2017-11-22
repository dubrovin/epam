package servers

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
)

// ReserverServer -
type ReserverServer struct {
	addr    string
	dbAddr  string
	router  *routing.Router
	errChan chan error
}

// NewReserverServer -
func NewReserverServer(dbAddr string, router *routing.Router) *ReserverServer {
	return &ReserverServer{
		dbAddr:  dbAddr,
		router:  router,
		errChan: make(chan error, 100),
	}
}

// Run -
func (r *ReserverServer) Run(addr string) {
	r.registerHandlers()
	go r.ListenAndServe(addr)
	go r.ReadErrChan()
}

// ListenAndServe -
func (r *ReserverServer) ListenAndServe(addr string) {
	log.Print("Listen and server addr = ", addr)
	r.errChan <- fasthttp.ListenAndServe(addr, r.router.HandleRequest)
}

// ReadErrChan -
func (r *ReserverServer) ReadErrChan() {
	for err := range r.errChan {
		log.Print("handlers server error: ", err)
	}
}

// Index -
func (r *ReserverServer) Index(ctx *routing.Context) error {
	fmt.Fprint(ctx, "Welcome to reserver!\n")
	return nil
}

func (r *ReserverServer) registerHandlers() {
	r.router.Get("/", r.Index)
	r.router.Get(`/products/<productid:\d+>/reserve`, r.ReserveProduct)
	r.router.Post(`/reserves/accept`, r.AcceptReserve)
}
