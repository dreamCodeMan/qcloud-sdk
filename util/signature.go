package util

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

//CreateSignature creates signature for string following Qcloud rules
func CreateSignatureSHA256(stringToSignature, accessKeySecret string) string {
	// Crypto by HMAC-SHA256
	hmacSha := hmac.New(sha256.New, []byte(accessKeySecret))
	hmacSha.Write([]byte(stringToSignature))
	sign := hmacSha.Sum(nil)

	// Encode to Base64
	base64Sign := base64.StdEncoding.EncodeToString(sign)

	return base64Sign
}

//CreateSignature creates signature for string following Qcloud rules
func CreateSignatureSHA1(stringToSignature, accessKeySecret string) string {
	// Crypto by HMAC-SHA1
	hmacSha := hmac.New(sha1.New, []byte(accessKeySecret))
	hmacSha.Write([]byte(stringToSignature))
	sign := hmacSha.Sum(nil)

	// Encode to Base64
	base64Sign := base64.StdEncoding.EncodeToString(sign)

	return base64Sign
}

func percentReplace(str string) string {
	//str = url.QueryEscape(str)
	//str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)

	return str
}

func percentReplaceHttp(str string) string {
	str = strings.Replace(str, "http://", "", -1)
	str = strings.Replace(str, "https://", "", -1)

	return str
}

func CreateSignatureForRequest(method, apiHost, secretKey string, values *url.Values) string {
	keys := make([]string, 0, len(*values))
	for k := range *values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	kvs := make([]string, 0, len(keys))
	for _, k := range keys {
		kvs = append(kvs, fmt.Sprintf("%s=%s", k, values.Get(k)))
	}
	queryStr := strings.Join(kvs, "&")
	reqStr := percentReplace(fmt.Sprintf("%s%s?%s", method, percentReplaceHttp(apiHost), queryStr))

	return CreateSignatureSHA1(reqStr, secretKey)
}
