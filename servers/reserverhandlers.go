package servers

import (
	"bytes"
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// ReserveProduct -
func (r *ReserverServer) ReserveProduct(ctx *routing.Context) error {
	ctx.SetContentType("application/json")
	productIDStr := ctx.Param("productid")
	_, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return err
	}
	urlPath := url.URL{
		Scheme: "http",
		Host:   r.dbAddr,
		Path:   fmt.Sprintf("/products/%s/reserve", productIDStr),
	}
	resp, err := http.Get(urlPath.String())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = ctx.WriteData(body)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

// AcceptReserve -
func (r *ReserverServer) AcceptReserve(ctx *routing.Context) error {
	ctx.SetContentType("application/json")

	urlPath := url.URL{
		Scheme: "http",
		Host:   r.dbAddr,
		Path:   "/reserves/accept",
	}
	resp, err := http.Post(urlPath.String(), "application/json", bytes.NewBuffer(ctx.PostBody()))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = ctx.WriteData(body)

	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}
