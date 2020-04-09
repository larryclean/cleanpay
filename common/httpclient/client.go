package httpclient

/**
封装方法来自：https://gitee.com/Tencent-BlueKing/bk-cmdb/blob/master/src/common/http/httpclient/client.go
*/
import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	header  map[string]string
	httpCli *http.Client
	baseUrl string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		httpCli: &http.Client{},
		header:  make(map[string]string),
	}
}
func (client *HttpClient) SetBaseUrl(baseUrl string) {
	client.baseUrl = baseUrl
}
func (client *HttpClient) GetClient() *http.Client {
	return client.httpCli
}

func (client *HttpClient) NewTransPort() *http.Transport {
	return &http.Transport{
		ResponseHeaderTimeout: 30 * time.Second,
	}
}

func (client *HttpClient) SetTimeOut(timeOut time.Duration) {
	client.httpCli.Timeout = timeOut
}

func (client *HttpClient) SetHeader(key, value string) {
	client.header[key] = value
}

func (client *HttpClient) GetHeader(key string) string {
	val, _ := client.header[key]
	return val
}

func (client *HttpClient) GET(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "GET", header, data)

}

func (client *HttpClient) POST(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "POST", header, data)
}

func (client *HttpClient) DELETE(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "DELETE", header, data)
}

func (client *HttpClient) PUT(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "PUT", header, data)
}

func (client *HttpClient) GETEx(url string, header http.Header, data []byte) (int, []byte, error) {
	return client.RequestEx(url, "GET", header, data)
}

func (client *HttpClient) POSTEx(url string, header http.Header, data []byte) (int, []byte, error) {
	return client.RequestEx(url, "POST", header, data)
}

func (client *HttpClient) DELETEEx(url string, header http.Header, data []byte) (int, []byte, error) {
	return client.RequestEx(url, "DELETE", header, data)
}

func (client *HttpClient) PUTEx(url string, header http.Header, data []byte) (int, []byte, error) {
	return client.RequestEx(url, "PUT", header, data)
}

func (client *HttpClient) Request(url, method string, header http.Header, data []byte) ([]byte, error) {
	code, body, err := client.RequestEx(url, method, header, data)
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("statuscode:%d", code)
	}
	return body, nil
}

func (client *HttpClient) RequestEx(url, method string, header http.Header, data []byte) (int, []byte, error) {
	if client.baseUrl != "" {
		url = client.baseUrl + url
	}
	var req *http.Request
	var errReq error
	if data != nil {
		req, errReq = http.NewRequest(method, url, bytes.NewReader(data))
	} else {
		req, errReq = http.NewRequest(method, url, nil)
	}

	if errReq != nil {
		return 0, nil, errReq
	}

	req.Close = true

	if header != nil {
		req.Header = header
	}

	for key, value := range client.header {
		req.Header.Set(key, value)
	}

	rsp, err := client.httpCli.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	return rsp.StatusCode, body, err
}

func (client *HttpClient) DoWithTimeout(timeout time.Duration, req *http.Request) (*http.Response, error) {
	ctx, _ := context.WithTimeout(req.Context(), timeout)
	req = req.WithContext(ctx)
	return client.httpCli.Do(req)
}
