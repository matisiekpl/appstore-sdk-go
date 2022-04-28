package appstore

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const ResponseContentTypeJson = "application/json; charset=utf-8"
const ResponseContentTypeGzip = "application/a-gzip"

//NewDefaultHttpClient create new http client
func NewDefaultHttpClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:    AppStoreConnectAPIHttpMaxIdleConnection,
		IdleConnTimeout: AppStoreConnectAPIHttpIdleConnectionTimeout,
	}
	return &http.Client{Transport: tr}
}

//RequestBuilder handler
type RequestBuilder struct {
	cfg   *Config
	token *AuthToken
}

//isValidToken method
func (rb *RequestBuilder) isValidToken() bool {
	return rb.token.IsValid()
}

//buildUri method
func (rb *RequestBuilder) buildUri(path string, query map[string]interface{}) (uri *url.URL, err error) {
	u, err := url.Parse(rb.cfg.Uri)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder.buildUri parse: %v", err)
	}
	u.Path = "/" + path
	u.RawQuery = rb.buildQueryParams(query)
	return u, err
}

//buildQueryParams method
func (rb *RequestBuilder) buildQueryParams(query map[string]interface{}) string {
	q := url.Values{}
	if query != nil {
		for k, v := range query {
			q.Set(k, fmt.Sprintf("%v", v))
		}
	}
	return q.Encode()
}

//buildHeaders method
func (rb *RequestBuilder) buildHeaders() http.Header {
	headers := http.Header{}
	headers.Set("Accept", "application/a-gzip")
	headers.Set("Accept-Encoding", "gzip")
	headers.Set("Authorization", "Bearer "+rb.token.Token)
	return headers
}

//BuildRequest method
func (rb *RequestBuilder) BuildRequest(ctx context.Context, method string, path string, query map[string]interface{}, body map[string]interface{}) (req *http.Request, err error) {
	method = strings.ToUpper(method)
	//build uri
	uri, err := rb.buildUri(path, query)
	if err != nil {
		return nil, fmt.Errorf("transport.request build uri: %v", err)
	}
	//build request
	req, err = http.NewRequestWithContext(ctx, method, uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("transport.request new request error: %v", err)
	}
	//build headers
	req.Header = rb.buildHeaders()
	return req, nil
}

//NewHttpTransport create new http transport
func NewHttpTransport(config *Config, token *AuthToken, h *http.Client) *Transport {
	if h == nil {
		h = NewDefaultHttpClient()
	}
	rb := &RequestBuilder{cfg: config, token: token}
	return &Transport{http: h, rb: rb}
}

//Transport wrapper
type Transport struct {
	http *http.Client
	rb   *RequestBuilder
}

//SendRequest method
func (t *Transport) SendRequest(ctx context.Context, method string, path string, query map[string]interface{}, body map[string]interface{}) (resp *http.Response, err error) {
	if !t.rb.isValidToken() {
		return nil, fmt.Errorf("transport.request invalid token: %v", err)
	}
	req, err := t.rb.BuildRequest(ctx, method, path, query, body)
	if err != nil {
		return nil, fmt.Errorf("transport.SendRequest: %v", err)
	}
	return t.http.Do(req)
}

//Get method
func (t *Transport) Get(ctx context.Context, path string, query map[string]interface{}) (resp *http.Response, err error) {
	return t.SendRequest(ctx, http.MethodGet, path, query, nil)
}

//ResponseBody struct
type ResponseBody struct {
	status int
	//ErrorResult Information with error details that an API returns in the response body whenever the API request is not successful.
	// .see https://developer.apple.com/documentation/appstoreconnectapi/errorresponse
	Errors []*Error `json:"errors,omitempty"`
}

//GetError method
func (r *ResponseBody) GetError() string {
	err := ""
	if len(r.Errors) > 0 {
		err = r.Errors[0].Detail
	}
	return err
}

//IsSuccess method
func (r *ResponseBody) IsSuccess() bool {
	return r.status < http.StatusMultipleChoices
}

type ResponseHandlerInterface interface {
	ReadBody(resp *http.Response) ([]byte, error)
	UnmarshalBody(data []byte, v interface{}) error
	RestoreBody(data []byte) (io.ReadCloser, error)
}

type ResponseHandlerJson struct {
}

func (r *ResponseHandlerJson) ReadBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (r *ResponseHandlerJson) UnmarshalBody(data []byte, v interface{}) error {
	return json.Unmarshal(data, &v)
}

func (r *ResponseHandlerJson) RestoreBody(data []byte) (io.ReadCloser, error) {
	return ioutil.NopCloser(bytes.NewBuffer(data)), nil
}

type ResponseHandlerGzip struct {
}

func (r *ResponseHandlerGzip) ReadBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	zr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	return ioutil.ReadAll(zr)
}

func (r *ResponseHandlerGzip) UnmarshalBody(data []byte, v interface{}) error {
	return UnmarshalCSV(data, v)
}

func (r *ResponseHandlerGzip) RestoreBody(data []byte) (io.ReadCloser, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}
	if err = gz.Flush(); err != nil {
		return nil, err
	}

	if err = gz.Close(); err != nil {
		return nil, err
	}
	return ioutil.NopCloser(bytes.NewBuffer(b.Bytes())), nil
}

func NewResponseHandler(contentType string) ResponseHandlerInterface {
	var handler ResponseHandlerInterface
	switch contentType {
	case ResponseContentTypeGzip:
		handler = &ResponseHandlerGzip{}
		break
	case ResponseContentTypeJson:
		handler = &ResponseHandlerJson{}
		break
	default:
		handler = &ResponseHandlerJson{}
	}
	return handler
}
