package appstore

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

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
		return nil, fmt.Errorf("RequestBuilder@buildUri parse: %v", err)
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

//Request method
func (t *Transport) Request(method string, path string, query map[string]interface{}, body map[string]interface{}) (resp *http.Response, err error) {
	if !t.rb.isValidToken() {
		return nil, fmt.Errorf("transport@request invalid token: %v", err)
	}
	//build uri
	uri, err := t.rb.buildUri(path, query)
	if err != nil {
		return nil, fmt.Errorf("transport@request build uri: %v", err)
	}
	//build request
	req, err := http.NewRequest(strings.ToUpper(method), uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("transport@request new request error: %v", err)
	}
	//build headers
	req.Header = t.rb.buildHeaders()
	return t.http.Do(req)
}

//Get method
func (t *Transport) Get(path string, query map[string]interface{}) (resp *http.Response, err error) {
	return t.Request("GET", path, query, nil)
}

//Response wrapper
type Response struct {
	raw *http.Response
	csv *CSV
}

//IsSuccess method
func (r *Response) IsSuccess() bool {
	return r.raw.StatusCode < http.StatusMultipleChoices
}

//GetRawResponse method
func (r *Response) GetRawResponse() *http.Response {
	return r.raw
}

//GetRawBody method
func (r *Response) GetRawBody() (string, error) {
	data, err := r.ReadBody()
	return string(data), err
}

//UnmarshalCSV method
func (r *Response) UnmarshalCSV(v interface{}) error {
	body, err := r.ReadBody()
	if err != nil {
		return err
	}
	return r.csv.Unmarshal(body, v)
}

//UnmarshalError method
func (r *Response) UnmarshalError(v interface{}) error {
	body, err := r.ReadBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

//ReadBody method
func (r *Response) ReadBody() ([]byte, error) {
	defer r.raw.Body.Close()
	if !r.IsSuccess() {
		return ioutil.ReadAll(r.raw.Body)
	} else {
		zr, _ := gzip.NewReader(r.raw.Body)
		defer zr.Close()
		return ioutil.ReadAll(zr)
	}
}

//NewResponse create new response
func NewResponse(raw *http.Response) *Response {
	return &Response{raw: raw, csv: &CSV{}}
}
