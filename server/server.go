package main

import (
	"encoding/json"
	"fmt"

	"github.com/Numenorean/pornhubEvaler"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type EvalData struct {
	Data []byte
}

const port = "6923"

func evalFunc(ctx *fasthttp.RequestCtx) {
	var data EvalData
	err := json.Unmarshal(ctx.PostBody(), &data)
	if err != nil {
		fmt.Fprintf(ctx, "{\"error\":\"%q\"}", err)
		return
	}
	result := pornhubEvaler.Eval(string(data.Data))
	fmt.Fprintf(ctx, "{\"cookie\":\"%s\"}", result)
}

func main() {
	fmt.Println("Server started on port:", port)
	r := router.New()
	r.POST("/eval", evalFunc)
	fasthttp.ListenAndServe(":"+port, r.Handler)
}
