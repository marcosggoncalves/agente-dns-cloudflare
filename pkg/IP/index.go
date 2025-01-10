package Ip

import (
	"io/ioutil"
	"net/http"
	"os"
)

func GetPublic() (string, error) {
	response, err := http.Get(os.Getenv("APIFY"))
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	ip, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}
