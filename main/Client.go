package main

import (
	"fmt"

	"github.com/GreyOfShadow/openApi-authentication/common/auth"
)

func main() {

	var config auth.AkskConfig
	config.AccessKey = "{accessKey}"
	config.SecretKey = "{secretKey}"
	config.EndPoint = "{endPoint}"
	config.Uri = "/api/test/example"
	config.Method = "POST"
	var requestBody string = ""
	response := auth.Invoke(config, requestBody)
	fmt.Printf("response body is %#v\n", response)
}
