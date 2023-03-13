package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

var (
	httpTimeout = 10 * time.Second
	httpClient  = http.Client{
		Timeout: httpTimeout,
		//Transport: &http.Transport{
		//	TLSClientConfig: &tls.Config{
		//		MinVersion: tls.VersionTLS13,
		//		MaxVersion: tls.VersionTLS13,
		//	},
		//},
	}
)

type HttpResp struct {
	Err        error
	Data       []byte
	Localtion  string
	RequestUrl string
}

func HttpPostWithJson(api string, param interface{}) HttpResp {
	var result HttpResp
	defer func() {
		e := recover()
		if e != nil {
			log.Error(e)
		}
	}()
	var bytesData []byte
	if param != nil {
		bytesData, _ = json.Marshal(param)
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", api, reader)
	if err != nil {
		result.Err = err
		return result
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := httpClient.Do(request)
	if err != nil {
		result.Err = err
		return result
	}
	defer func() {
		resp.Body.Close()
	}()
	if resp == nil {
		result.Err = errors.New("空数据")
		return result
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		result.Err = err
		return result
	}
	result.Data = respBytes
	return result
}

func HttpPostWithJsonAndHeader(api string, header map[string]string, param interface{}) HttpResp {
	var result HttpResp
	defer func() {
		e := recover()
		if e != nil {
			log.Error(e)
		}
	}()
	var bytesData []byte
	if param != nil {
		bytesData, _ = json.Marshal(param)
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", api, reader)
	if err != nil {
		result.Err = err
		return result
	}
	if len(header) > 0 {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := httpClient.Do(request)
	if err != nil {
		result.Err = err
		return result
	}
	defer func() {
		resp.Body.Close()
	}()
	if resp == nil {
		result.Err = errors.New("空数据")
		return result
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		result.Err = err
		return result
	}
	result.Data = respBytes
	return result
}

func HttpGet(api string) HttpResp {
	var result HttpResp
	request, _ := http.NewRequest("GET", api, nil)
	q := request.URL.Query()
	request.URL.RawQuery = q.Encode()
	resp, err := httpClient.Do(request)
	if err != nil {
		result.Err = err
		log.Info("访问:", err)
		return result
	}
	defer func() {
		resp.Body.Close()
	}()
	body, _ := io.ReadAll(resp.Body)
	result.Data = body
	result.RequestUrl = resp.Request.URL.String()
	return result
}
