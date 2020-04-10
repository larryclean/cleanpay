package v3

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/larry-dev/cleanpay/common/crand"
	"github.com/larry-dev/cleanpay/common/hash"
	"github.com/larry-dev/cleanpay/common/httpclient"
	"net/http"
	"net/url"
	"time"
)

const (
	POST    = "POST"
	GET     = "GET"
	baseUrl = "https://api.mch.weixin.qq.com/v3"

	fmtSign     = "%s\n%s\n%s\n%s\n%s\n"
	fmtAuth     = "WECHATPAY2-SHA256-RSA2048 mchid=\"%s\",nonce_str=\"%s\",signature=\"%s\",timestamp=\"%s\",serial_no=\"%s\""
	createOrder = "/payscore/serviceorder"
	cancelOrder = "/payscore/serviceorder/%s/cancel"
	queryOrder  = "/payscore/serviceorder?service_id=%s&appid=%s"
)

type PayScore struct {
	client     *httpclient.HttpClient
	MchID      string
	SerialNo   string
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func New(mchId, serialNo, ServiceID, keyPath, certPath string) (*PayScore, error) {
	client := httpclient.NewHttpClient()
	client.SetBaseUrl(baseUrl)
	client.SetTimeOut(5 * time.Second)
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("Accept", "application/json")
	client.SetHeader("User-Agent", "clean pay client")
	privateKey, err := hash.GetPrivateKey(keyPath)
	if err != nil {
		return nil, err
	}
	publicKey, err := hash.GetPublicKey(certPath)
	if err != nil {
		return nil, err
	}
	return &PayScore{
		client:     client,
		MchID:      mchId,
		SerialNo:   serialNo,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func (wx *PayScore) CreateOrder(order CreateOrder) (*RspCreateOrder, error) {
	res, err := wx.Do(order, POST, createOrder)
	if err != nil {
		return nil, err
	}
	var result RspCreateOrder
	_ = json.Unmarshal(res, &result)
	return &result, nil
}
func (wx *PayScore) QueryOrder(order QueryOrder) (resp *RspCreateOrder, err error) {
	url := fmt.Sprintf(queryOrder, order.ServiceID, order.AppID)
	if order.OutOrderNo != "" {
		url += "&out_order_no=" + order.OutOrderNo
	} else {
		url += "&query_id=" + order.QueryID
	}
	res, err := wx.Do(order, GET, url)
	if err != nil {
		return nil, err
	}
	var result RspCreateOrder
	_ = json.Unmarshal(res, &result)
	return &result, nil
	//fmt.Println(string(res), err)
}
func (wx *PayScore) CancelOrder(order CancelOrder) {
	uri := fmt.Sprintf(cancelOrder, order.OutOrderNo)
	res, err := wx.Do(order, POST, uri)
	fmt.Println(string(res), err)
}

func (wx *PayScore) Do(req interface{}, method, uri string) ([]byte, error) {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	nonceStr := crand.Letters(16)
	var data []byte
	if method == POST {
		d, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}
		data = d
	}
	header := http.Header{}
	sign, err := wx.sign(method, uri, timestamp, nonceStr, string(data))
	if err != nil {
		return nil, err
	}
	header.Add("Authorization", fmt.Sprintf(fmtAuth, wx.MchID, nonceStr, sign, timestamp, wx.SerialNo))
	if method == POST {
		return wx.client.POST(uri, header, data)
	}
	return wx.client.GET(uri, header, data)
}
func (wx *PayScore) sign(method, uri, timestamp, nonce, body string) (string, error) {
	path := wx.getUri(uri)
	hash := sha256.New()
	fmt.Println(fmt.Sprintf(fmtSign, method, path, timestamp, nonce, body))
	hash.Write([]byte(fmt.Sprintf(fmtSign, method, path, timestamp, nonce, body)))
	cipherdata, err := rsa.SignPKCS1v15(rand.Reader, wx.PrivateKey, crypto.SHA256, hash.Sum(nil))
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(cipherdata)
	return signature, nil
}
func (wx *PayScore) getUri(path string) string {
	ur, _ := url.Parse(baseUrl + path)
	return ur.RequestURI()
}
