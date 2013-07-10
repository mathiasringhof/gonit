package main

import (
	"net/http"
)

type SimpleHeaderAuthenticator struct {
	Headername, Password string
}

func (sha *SimpleHeaderAuthenticator) Wrap(handler func(rw http.ResponseWriter, req *http.Request)) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.Header.Get(sha.Headername) != sha.Password {
			http.Error(rw, "Requesting correct authentication in header "+sha.Headername, 401)
			return
		}
		handler(rw, req)
	}
}
