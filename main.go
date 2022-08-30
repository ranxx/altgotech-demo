package main

import (
	"fmt"
	"net/http"

	"github.com/ranxx/altgotech-demo/router"
)

func main() {

	http.ListenAndServe(":9123", router.Router())

	fmt.Println("hello")
}
