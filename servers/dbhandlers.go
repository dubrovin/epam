package servers

import (
	"encoding/json"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"strconv"
)

// AcceptResp -
type AcceptResp struct {
	Accepted bool `json:"accepted"`
}

// HashResp -
type HashResp struct {
	Hash string `json:"hash"`
}

// GetProducts -
func (c *DBServer) GetProducts(ctx *routing.Context) error {
	ctx.SetContentType("application/json")

	products := c.db.GetProducts()

	jsonData, err := json.Marshal(products)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return err
	}

	err = ctx.WriteData(jsonData)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

// ReserveProduct -
func (c *DBServer) ReserveProduct(ctx *routing.Context) error {
	ctx.SetContentType("application/json")
	productIDStr := ctx.Param("productid")
	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return err
	}

	hash, err := c.db.ReserveProduct(productID, c.ttl)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}

	jsonData, err := json.Marshal(HashResp{Hash: hash})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return err
	}

	err = ctx.WriteData(jsonData)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

// AcceptReserve -
func (c *DBServer) AcceptReserve(ctx *routing.Context) error {
	ctx.SetContentType("application/json")

	requestData := &HashResp{}
	err := json.Unmarshal(ctx.PostBody(), requestData)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return err
	}

	ok, err := c.db.AcceptReserve(requestData.Hash)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}

	jsonData, err := json.Marshal(AcceptResp{Accepted: ok})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return err
	}

	err = ctx.WriteData(jsonData)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}
