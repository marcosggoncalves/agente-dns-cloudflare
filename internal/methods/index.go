package Methods

import (
	"bytes"
	Schema "dnspax/internal/schemas"
	Message "dnspax/pkg/message"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func getDNSRecordID(recordName string) (string, error) {
	url := fmt.Sprintf("%s/zones/%s/dns_records", os.Getenv("APICLOUDFLARE"), os.Getenv("ZONEID"))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Email", os.Getenv("EMAIL"))
	req.Header.Set("X-Auth-Key", os.Getenv("APIKEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var records struct {
		Result []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"result"`
	}

	err = json.NewDecoder(resp.Body).Decode(&records)
	if err != nil {
		return "", err
	}

	for _, record := range records.Result {
		if record.Name == recordName {
			return record.ID, nil
		}
	}

	return "", fmt.Errorf(Message.FALHAR_OBTER_RECORD_ID)
}

func UpdatedDNS(ip string) error {
	recordID, err := getDNSRecordID(os.Getenv("RECORDNAME"))
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/zones/%s/dns_records/%s",
		os.Getenv("APICLOUDFLARE"),
		os.Getenv("ZONEID"), recordID)

	requestBody, err := json.Marshal(Schema.DNSRequest{
		Content:   ip,
		Name:      os.Getenv("RECORDNAME"),
		Proxiable: false,
		Proxied:   false,
		TTL:       1,
		Type:      "A",
		ZoneID:    os.Getenv("ZONEID"),
		ZoneName:  os.Getenv("ZONENAME"),
		ID:        recordID,
	})

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Email", os.Getenv("EMAIL"))
	req.Header.Set("X-Auth-Key", os.Getenv("APIKEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response Schema.DNSResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	if !response.Success {
		return fmt.Errorf(Message.FALHA_ATUALIZACAO_DNS, response.Errors)
	}

	fmt.Printf(Message.IP_ATUALIZADO, ip)
	return nil
}
