package http

import (
	"os"
    "fmt"
	"time"
	"bytes"
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"post-deployment-smoke-tests/types"
	"post-deployment-smoke-tests/internal/errors"
)

func VariableNotEmpty(variable string) (bool) {
	if variable != "" {
		return true
	}
	return false
}

func SetPlayload(check *types.HttpCheck) *bytes.Buffer {
	if check.Method == "POST" {
		return bytes.NewBuffer([]byte(check.Playload))
	}
	return bytes.NewBuffer([]byte(nil))
}

func HttpwRequest(check *types.HttpCheck, base_url string) (bool) {
	url := fmt.Sprintf("%s%s", base_url, check.Path)

	req, err := http.NewRequest(check.Method, url, SetPlayload(check))
	errors.CheckError(err)

	if len(check.Headers) > 0 {
		for k, v := range check.Headers {
				req.Header.Set(k, v)
		}
	}

	if VariableNotEmpty(check.Auth.Login) &&  VariableNotEmpty(check.Auth.Password) {
		req.SetBasicAuth(check.Auth.Login, check.Auth.Password)
	}

	client := &http.Client{Timeout:time.Second * 5, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},},}
	resp, err := client.Do(req)
	errors.CheckError(err)
	
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	errors.CheckError(err)

	if resp.StatusCode == check.Response_code  {
		if VariableNotEmpty(check.Response_body)  && (check.Response_body != string(body)) {
			fmt.Printf("Name of check: '%s' is failed. Responsed body is not as expected.\n", check.Name)
			return false	
		}
		return true

	} else {
		fmt.Printf("Name of check: '%s' is failed. Responsed status code is %d and expected code is %d.\n", check.Name, resp.StatusCode, check.Response_code)
		// fmt.Println(resp.StatusCode, check.Response_code)
		return false
	}
}

 func GetBaseUrl(check *types.HttpCheck, Base_url string) (string) {
	if VariableNotEmpty(check.Base_url) {
		return check.Base_url
	}
	return Base_url
 }

func DoHttpCheck(config types.HttpConfig) {
	for _, check := range config.Checks {
		if HttpwRequest(check, GetBaseUrl(check, config.Base_url)) {
			fmt.Printf("Name of check: '%s' is passed.\n", check.Name)
		} else {
			os.Exit(1)
		}
	}
}
