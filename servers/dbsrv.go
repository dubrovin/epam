package servers

import (
	"github.com/dubrovin/epam/services"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

// Controller -
type DBController struct {
	db      *services.DataBase
	router  *routing.Router
	errChan chan error
	ttl     time.Duration
}

func NewDBController(db *services.DataBase, router *routing.Router, ttl time.Duration) *DBController {
	return &DBController{
		db:      db,
		router:  router,
		errChan: make(chan error, 100),
		ttl:     ttl,
	}
}

func (c *DBController) Run(addr string, checkerSleep time.Duration) {
	c.registerHandlers()
	go c.db.Checker(checkerSleep)
	go c.ListenAndServe(addr)
	go c.ReadErrChan()
}

// ListenAndServe -
func (c *DBController) ListenAndServe(addr string) {
	log.Print("Listen and server addr = ", addr)
	c.errChan <- fasthttp.ListenAndServe(addr, c.router.HandleRequest)
}

// ReadErrChan -
func (c *DBController) ReadErrChan() {
	for err := range c.errChan {
		log.Print("handlers server error: ", err)
	}
}

func (c *DBController) registerHandlers() {
	c.router.Get("/products", c.GetProducts)
	c.router.Get(`/products/<productid:\d+>/reserve`, c.ReserveProduct)
	c.router.Post(`/reserves/accept`, c.AcceptReserve)
}
