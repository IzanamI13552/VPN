package main

import (
	"golang.zx2c4.com/wintun"
	"log"
)

func main() {
	adapter, err := wintun.CreateAdapter("MyWintun", "MyWintun", nil)
	if err != nil {
		log.Fatalf("Ошибка создание адаптера : %v", err)
	}
	log.Printf("Адаптер создан : %v", adapter)
}
