package servers

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
	"io/ioutil"
)

func (w *WatcherServer) GetProducts(ctx *routing.Context) error {
	ctx.SetContentType("application/json")
	urlPath:=url.URL{
		Scheme: "http",
		Host: w.dbAddr,
		Path: "/products",
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
