package main

import (
	"fmt"
	"service"
)

func main() {
	s := service.NewUser()
	_, err := s.Get(7)
	fmt.Printf("%+v\n", err)
}
