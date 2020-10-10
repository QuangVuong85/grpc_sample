package utils

import (
	"io/ioutil"
	"net/http"
)

func GetLocalIP() []byte {
	url := "https://api.ipify.org?format=text"
	// https://www.ipify.org
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return ip
}
