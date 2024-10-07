package main

import (
	a "first/firstgolang"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)



func main(){
	fmt.Println( a.AddNew() )

	r :=  echo.New()

	r.GET("/", func (ctx echo.Context) error  {
		data := "Hello index"
		return ctx.String(http.StatusOK,data)
	})

	r.Start(":3000")
}