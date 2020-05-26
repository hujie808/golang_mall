package main

import (
	"golang_mall/services"
	"log"
)

func main(){
	order :=services.NewRedLogService()
	log.Println(order.GetAll(0,5))
}