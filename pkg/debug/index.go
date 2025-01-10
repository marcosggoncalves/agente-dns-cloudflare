package debug

import (
	Message "dnspax/pkg/message"
	"encoding/json"
	"fmt"
	"log"
)

func Json[T any](data T) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf(Message.FALHA_JSON, err)
		return
	}

	fmt.Println(string(jsonData))
}
