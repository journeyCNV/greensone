package test

import (
	"fmt"
	"github.com/journeycnv/greensone/gsweb"
)

func TestH1() gsweb.HandlerFunc {
	return func(c *gsweb.Context) error {
		fmt.Println("mid 1")
		c.Next()
		fmt.Println("mid 1 ")
		return nil
	}
}

func TestH2() gsweb.HandlerFunc {
	return func(c *gsweb.Context) error {
		fmt.Println("mid 2")
		c.Next()
		fmt.Println("mid 2 ")
		return nil
	}
}
