package client

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ZMuSiShui/ips/conf"
	"github.com/matryer/try"
	log "github.com/sirupsen/logrus"
)

func loadFromInternet(url string) (ranges conf.IPRangesDoc, err error) {
	var request *http.Request
	request, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Debugf("Error: %v", err)
		os.Exit(2)
	}

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 3,
			DisableKeepAlives:   false,
		},
		Timeout: time.Duration(6) * time.Second,
	}

	var resp *http.Response

	err = try.Do(func(attempt int) (bool, error) {
		var rErr error
		resp, err = client.Do(request)
		return attempt < 3, rErr
	})

	if err != nil {
		log.Debugf("Error: %v", err)
		os.Exit(2)
	}

	defer resp.Body.Close()

	var syncRespBodyBytes []byte
	syncRespBodyBytes, err = getResponseBody(resp)

	if err != nil {
		log.Debugf("Error: %v", err)
		os.Exit(2)
	}

	err = json.Unmarshal(syncRespBodyBytes, &ranges)
	return
}

// 获取响应结构体
func getResponseBody(resp *http.Response) (body []byte, err error) {
	var output io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		output, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Debugf("Error: %v", err)
			os.Exit(2)
		}
		if err != nil {
			log.Debugf("Error: %v", err)
			os.Exit(2)
		}
	default:
		output = resp.Body
		if err != nil {
			log.Debugf("Error: %v", err)
			os.Exit(2)
		}
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(output)
	if err != nil {
		log.Debugf("Error: %v", err)
		os.Exit(2)
	}
	body = buf.Bytes()
	return
}
