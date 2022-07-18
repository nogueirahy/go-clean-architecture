package http

import "fmt"

type HttpAdapter struct{}

func (c *HttpAdapter) Get() float64 {
	fmt.Println("request...")
	return 1200
}
