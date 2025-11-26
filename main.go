package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"log"
	"net"
)

const serverAddr = "127.0.0.1:8080"

func main() {
	// Создание сикретного ключа с помощью rand
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}
	// Создание UDP-адреса
	addr, _ := net.ResolveUDPAddr("udp", serverAddr)
	// Запуск UDP сервера
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Ошибка создание адаптера : %v", err)
	}
	defer conn.Close()

	fmt.Println("VPN слушаю сервер ", serverAddr)
	// Создаем буфер для вход данных
	buf := make([]byte, 2048)
	// Прием пакетов
	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}
		nonceSize := aead.NonceSize()

		if n < nonceSize {
			continue
		}
		nonce := buf[:12]
		ciphertext := buf[12:n]
		plain, err := aead.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			fmt.Println("Ошибка расшифровки")
			continue
		}
		// Вывод результата
		fmt.Printf("От %s: %s\n", clientAddr, string(plain))

	}
}
