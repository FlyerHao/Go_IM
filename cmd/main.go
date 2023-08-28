package main

import (
	"fmt"
	"goIm/router"
	"github.com/gin-gonic/gin"
)



func main() {
    addr := ":8080"
	r := gin.Default()
	r.GET("/send", router.Send{}.SendMessage)
	fmt.Println(addr)
	r.Run(addr)
}
