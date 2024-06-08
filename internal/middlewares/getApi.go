package middlewares

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"post/internal/database/models"
)

func GetApi() string {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		log.Printf("Увы без ip:%w", err.Error())
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var ipResp models.IPResponse
	err = json.Unmarshal(data, &ipResp)
	if err != nil {
		return ""
	}
	return ipResp.IP
}
