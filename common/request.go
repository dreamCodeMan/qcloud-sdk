package common

import (
	"time"

	"fmt"
	"math/rand"
)

// Constants for Qcloud API requests
const (
	Version            = "0.1"
	SignatureVersion   = "1.0"
	SignatureMethod    = "HmacSHA1"
	JSONResponseFormat = "JSON"
	XMLResponseFormat  = "XML"
	CVMRequestMethod   = "GET"
)

type Request struct {
	Version         string
	SecretId        string
	Signature       string
	SignatureMethod string
	Timestamp       int64
	Nonce           string
	Action          string
	Region          string
	RequestClient   string
}

func (request *Request) init(version, action, secretId, region string) {
	request.Timestamp = time.Now().Unix()
	request.SignatureMethod = SignatureMethod
	request.Nonce = fmt.Sprint(uint(rand.Int()))
	request.Action = action
	request.SecretId = secretId
	request.Region = region
	request.Version = version
	request.RequestClient = "QcloudGO_SDK_" + Version
}

type Response struct {
	RequestId string
}

type ErrorResponse struct {
	Response
	Code    string
	Message string
}

// An Error represents a custom error for Qcloud API failure response
type Error struct {
	ErrorResponse
	StatusCode int //Status Code of HTTP Response
}

func (e *Error) Error() string {
	return fmt.Sprintf("Qcloud API Error: RequestId: %s Status Code: %d Code: %s Message: %s", e.RequestId, e.StatusCode, e.Code, e.Message)
}
