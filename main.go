package main

import (
	"codegen/service"
	_ "codegen/service"
)

func main() {
	service.NewGenCode().Gen()
}