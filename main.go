package main

import (
	"civi-backend-challenge/routers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	endpoint := fmt.Sprintf(":%s", os.Getenv("PORT"))
	initRouters := routers.InitRouter()

	server := &http.Server{
		Addr:    endpoint,
		Handler: initRouters,
	}

	_ = server.ListenAndServe()
}
