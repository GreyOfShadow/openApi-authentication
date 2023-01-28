package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/GreyOfShadow/openApi-authentication/common/model"
)

func Call(requestMsg model.HttpRequestMsg) string {
	client := &http.Client{}
	req, err := http.NewRequest(requestMsg.Method, requestMsg.Url, strings.NewReader(requestMsg.Body))
	if err != nil {
		// handle error
		fmt.Printf("错误Call：%v", err)
		return ""
	}
	for key := range requestMsg.Headers {
		req.Header.Set(key, requestMsg.Headers[key])
	}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Printf("错误ReadAll：%v", err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("错误Calldefer：%v", err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Printf("错误ReadAll：%v", err)
		return ""
	}
	fmt.Println(resp)
	return string(body)
}
