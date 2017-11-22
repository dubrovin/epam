package servers

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
)

// DBServer -
type WatcherServer struct {
	addr string
	dbAddr  string
	router  *routing.Router
	errChan chan error
}

func NewWatcherServer(dbAddr string, router *routing.Router) *WatcherServer {
	return &WatcherServer{
		dbAddr:  dbAddr,
		router:  router,
		errChan: make(chan error, 100),
	}
}

func (w *WatcherServer) Run(addr string) {
	w.registerHandlers()
	go w.ListenAndServe(addr)
	go w.ReadErrChan()
}

// ListenAndServe -
func (w *WatcherServer) ListenAndServe(addr string) {
	log.Print("Listen and server addr = ", addr)
	w.errChan <- fasthttp.ListenAndServe(addr, w.router.HandleRequest)
}

// ReadErrChan -
func (w *WatcherServer) ReadErrChan() {
	for err := range w.errChan {
		log.Print("handlers server error: ", err)
	}
}

func (w *WatcherServer) registerHandlers() {
	w.router.Get("/products", w.GetProducts)
}
