package wxpay

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/larry-dev/cleanpay/common/cmap"
	"hash"
	"strings"
)

const (
	// 签名方式
	SignTypeMd5        = "MD5"
	SignTypeHmacSha256 = "HMAC-SHA256"
)

func Sign(h cmap.H, apiKey string, signType string) cmap.H {
	h["sign_type"] = signType
	var ha hash.Hash
	if signType == SignTypeHmacSha256 {
		ha = hmac.New(sha256.New, []byte(apiKey))
	} else {
		ha = md5.New()
	}
	str := h.Sort()
	fmt.Println(str)
	str = fmt.Sprintf("%s&key=%s", str, apiKey)
	ha.Write([]byte(str))
	h["sign"] = strings.ToUpper(hex.EncodeToString(ha.Sum(nil)))
	return h
}
