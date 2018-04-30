package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/dreamCodeMan/qcloud-sdk/util"
)

type Client struct {
	SecretId     string
	SecretKey    string
	region       string
	endpoint     string
	Debug        bool
	httpClient   *http.Client
	version      string
	businessInfo string
	userAgent    string
}

// NewClient creates a new instance of CVM client
func (client *Client) Init(endpoint, version, secretId, secretKey, region string) {
	client.SecretId = secretId
	client.SecretKey = secretKey
	client.Debug = false
	client.httpClient = &http.Client{}
	client.endpoint = endpoint
	client.version = version
	client.region = region
}

// Invoke sends the raw HTTP request for CVM services
func (client *Client) Invoke(action string, args interface{}, response interface{}) error {

	request := Request{}
	request.init(client.version, action, client.SecretId, client.region)

	query := util.ConvertToQueryValues(request)
	util.SetQueryValues(args, &query)
	//util.EncodeStruct(args, &query)
	//delete null map
	delete(query, "")

	// Sign request
	signature := util.CreateSignatureForRequest(CVMRequestMethod, client.endpoint, client.SecretKey, &query)

	// Generate the request URL
	requestURL := client.endpoint + "?" + query.Encode() + "&Signature=" + url.QueryEscape(signature)

	httpReq, err := http.NewRequest(CVMRequestMethod, requestURL, nil)

	if err != nil {
		return GetClientError(err)
	}

	// TODO move to util and add build val flag
	httpReq.Header.Set("X-SDK-Client", `QcloudGO/`+Version+client.businessInfo)

	httpReq.Header.Set("User-Agent", httpReq.UserAgent()+" "+client.userAgent)

	t0 := time.Now()
	httpResp, err := client.httpClient.Do(httpReq)
	t1 := time.Now()
	if err != nil {
		return GetClientError(err)
	}
	statusCode := httpResp.StatusCode

	if client.Debug {
		log.Printf("Invoke %s %s %d (%v)", CVMRequestMethod, requestURL, statusCode, t1.Sub(t0))
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		return GetClientError(err)
	}

	if client.Debug {
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, body, "", "    ")
		log.Println(string(prettyJSON.Bytes()))
	}

	if statusCode >= 400 && statusCode <= 599 {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(body, &errorResponse)
		CVMError := &Error{
			ErrorResponse: errorResponse,
			StatusCode:    statusCode,
		}
		return CVMError
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return GetClientError(err)
	}

	return nil
}

func GetClientErrorFromString(str string) error {
	return &Error{
		ErrorResponse: ErrorResponse{
			Code:    "QcloudGoClientFailure",
			Message: str,
		},
		StatusCode: -1,
	}
}

func GetClientError(err error) error {
	return GetClientErrorFromString(err.Error())
}
