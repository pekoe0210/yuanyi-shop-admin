package main

import (
	"fmt"
	"net/http"
	"github.com/pekoe0210/yuanyi-shop-admin/controller"
)

func main()  {
	http.HandleFunc("/index",controller.Index)
	http.ListenAndServe(":8080",nil)
	fmt.Println("yuanyi-web")
}