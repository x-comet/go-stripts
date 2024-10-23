package http_scripts

import (
	"bytes"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"net/http"
	"time"
)

func SendHttpRequest(url string, method string, requestBody []byte) ([]byte, error) {
	// 设置 HTTP 请求方法
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		hlog.Warn("get new http request fail, err: ", err.Error())
		return nil, err
	}

	// 设置 Bearer Token 在 Header 中
	req.Header.Set("Authorization", "Bearer sk-ryfzhvnjbafqxrjqzirnwzgsilzpltdnmificatxembmoqua")
	req.Header.Set("Content-Type", "application/json") // 根据实际需要设置 Content-Type

	client := &http.Client{
		// 超时时间设置为 20 s
		Timeout: 20 * time.Second,
	}
	//hlog.Info("send http request, url: ", req.URL, " req.body: ", string(requestBody))
	resp, err := client.Do(req)
	if err != nil {
		hlog.Warn("do http request fail, err: ", err.Error())
		return nil, errors.New("do request fail")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			hlog.Warn("close body fail, err: ", err.Error())
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		hlog.Warn("read response body fail, err: ", err.Error())
		return nil, err
	}
	//fmt.Println("resp body: ", string(body))

	return body, nil
}

func SendHttpRequestWithoutResponse(url string, method string, requestBody []byte) error {
	// 设置 HTTP 请求方法
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		hlog.Warn("get new http request fail, err: ", err.Error())
		return err
	}

	// 设置 Bearer Token 在 Header 中
	req.Header.Set("Authorization", "Bearer sk-ryfzhvnjbafqxrjqzirnwzgsilzpltdnmificatxembmoqua")
	req.Header.Set("Content-Type", "application/json") // 根据实际需要设置 Content-Type

	client := &http.Client{
		// 超时时间设置为 10 s
		Timeout: 1 * time.Second,
	}
	hlog.Info("send http request, url: ", req.URL, " req.body: ", string(requestBody))
	_, err = client.Do(req)
	if err != nil {
		hlog.Warn("do http request fail, err: ", err.Error())
		return err
	}

	return nil
}
