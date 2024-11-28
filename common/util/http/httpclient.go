package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTPClient 是一个封装了常见 HTTP 请求功能的工具类
type HTTPClient struct {
	client *http.Client
}

var Client *HTTPClient

func init() {
	Client = NewHTTPClient()
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:          500,
				MaxIdleConnsPerHost:   300,
				IdleConnTimeout:       30 * time.Second,
				TLSHandshakeTimeout:   30 * time.Second,
				ExpectContinueTimeout: 5 * time.Second,
			},
		},
	}
}

// Get 发送一个 GET 请求
func (c *HTTPClient) Get(url string) ([]byte, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// PostJSON 发送一个 POST 请求，请求体是 JSON 格式
func (c *HTTPClient) PostJSON(url string, body interface{}) []byte {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil
	}

	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return response
}

// PostFormData sends a POST request with form data and sets a Cookie
func (c *HTTPClient) PostFormDataAndCookie(urlStr string, data map[string]interface{}, cookie string) ([]byte, *http.Response, error) {
	// Construct form data
	formData := url.Values{}
	for key, value := range data {
		formData.Set(key, fmt.Sprint(value))
	}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, resp, nil
}

// PostFormData sends a POST request with form data and sets a Cookie
func (c *HTTPClient) PostFormDataStrAndCookie(urlStr string, data map[string]string, cookie string) ([]byte, *http.Response, error) {
	// Construct form data
	formData := url.Values{}
	for key, value := range data {
		formData.Set(key, value)
	}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return body, resp, nil
}

// PostFormData sends a POST request with form data
func (c *HTTPClient) PostFormData(urlStr string, data map[string]string) ([]byte, *http.Response) {
	// Construct form data
	formData := url.Values{}
	for key, value := range data {
		formData.Set(key, value)
	}

	// Create a new request with form data
	resp, err := c.client.PostForm(urlStr, formData)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}

	return body, resp
}
func HttpPost(urlP string, res interface{}) []byte {
	bytesData, _ := json.Marshal(res)

	resp, err := http.Post(urlP, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func CurlCommand(req *http.Request) string {
	var cmdBuf bytes.Buffer

	cmdBuf.WriteString("curl -X " + req.Method + " '" + "http://" + req.Host + req.URL.String() + "' \\\n")

	req.Header.Del("Host") // curl 自动添加 Host
	//var timestamp, nonce, signature string
	for key, values := range req.Header {
		for _, value := range values {
			if key == "Content-Length" {
				continue
			}
			//if key == "X-Timestamp" {
			//	timestamp = value
			//}
			//if key == "X-Nonce" {
			//	nonce = value
			//}
			//if key == "X-Sign" {
			//	signature = value
			//}
			cmdBuf.WriteString("    --header '" + key + ": " + value + "' \\\n")
		}
	}

	//if timestamp == "" || nonce == "" || signature == "" {
	//	logger.LogrusObj.Errorf("防刷策略进行")
	//	panic("请求头错误")
	//}

	//println(timestamp, nonce, signature)
	if req.Body != nil {
		bodyBytes, _ := ioutil.ReadAll(req.Body)
		bodyStr := string(bodyBytes)
		//md5 := gmd5.Md5(timestamp + nonce + bodyStr)
		//if md5 != signature {
		//	panic("sign错误")
		//}
		if len(bodyBytes) > 1024 {
			bodyStr = string(bodyBytes[:1024]) + "..."
		}

		cmdBuf.WriteString("    --data-raw '" + bodyStr + "'")
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	return cmdBuf.String()
}
func HttpOriginalFormat(req *http.Request) string {
	var sb strings.Builder

	// 请求行
	sb.WriteString(fmt.Sprintf("%s %s %s\r\n", req.Method, req.URL.Path, req.Proto))

	// 请求头部
	req.Header.Del("Host") // HTTP 请求中 Host 由客户端自动添加
	for key, values := range req.Header {
		for _, value := range values {
			if key == "Content-Length" {
				continue
			}
			sb.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
		}
	}

	// 空行分隔符
	sb.WriteString("\r\n")

	// 请求体
	if req.Body != nil {
		bodyBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return "无法读取请求体"
		}
		//defer req.Body.Close()
		sb.WriteString(string(bodyBytes))
	}

	return sb.String()
}
