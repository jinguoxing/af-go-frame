package contract

import "net/http"

const HttpEngineKey = "af:http_engine"

type HttpEngine interface {

    HttpEngine() http.Handler
}