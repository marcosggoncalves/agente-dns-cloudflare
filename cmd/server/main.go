package main

import (
	Handler "dnspax/internal/handlers"
	Ip "dnspax/pkg/ip"
	Message "dnspax/pkg/message"
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(Message.FALHA_ENV)
	}

	for {
		ip, err := Ip.GetPublic()
		if err != nil {
			fmt.Printf(Message.FALHA_CAPTURA_IP, err)
			time.Sleep(1 * time.Minute)
			return
		}

		err = Handler.UpdatedDNS(ip)
		if err != nil {
			fmt.Printf(Message.FALHA_ATUALIZACAO_DNS, err)
		}
		time.Sleep(1 * time.Minute)
	}
}
