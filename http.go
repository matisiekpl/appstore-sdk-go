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

func NewDefaultHttpClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:    AppStoreConnectAPIHttpMaxIdleConnection,
		IdleConnTimeout: AppStoreConnectAPIHttpIdleConnectionTimeout,
	}
	return &http.Client{Transport: tr}
}

type RequestBuilder struct {
	cfg   *Config
	token *AuthToken
}

func (rb *RequestBuilder) isValidToken() bool {
	return rb.token.IsValid()
}

func (rb *RequestBuilder) buildUri(path string, query map[string]interface{}) (uri *url.URL, err error) {
	u, err := url.Parse(rb.cfg.Uri)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder@buildUri parse: %v", err)
	}
	u.Path = "/" + path
	u.RawQuery = rb.buildQueryParams(query)
	return u, err
}

func (rb *RequestBuilder) buildQueryParams(query map[string]interface{}) string {
	q := url.Values{}
	if query != nil {
		for k, v := range query {
			q.Set(k, fmt.Sprintf("%v", v))
		}
	}
	return q.Encode()
}

func (rb *RequestBuilder) buildHeaders() http.Header {
	headers := http.Header{}
	headers.Set("Accept", "application/a-gzip")
	headers.Set("Accept-Encoding", "gzip")
	headers.Set("Authorization", "Bearer "+rb.token.Token)
	return headers
}

func NewHttpTransport(config *Config, token *AuthToken, h *http.Client) *Transport {
	if h == nil {
		h = NewDefaultHttpClient()
	}
	rb := &RequestBuilder{cfg: config, token: token}
	return &Transport{http: h, rb: rb}
}

type Transport struct {
	http *http.Client
	rb   *RequestBuilder
}

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

func (t *Transport) Get(path string, query map[string]interface{}) (resp *http.Response, err error) {
	return t.Request("GET", path, query, nil)
}

type Response struct {
	raw *http.Response
	csv *CSV
}

func (r *Response) IsSuccess() bool {
	return r.raw.StatusCode < http.StatusMultipleChoices
}

func (r *Response) GetRawResponse() *http.Response {
	return r.raw
}

func (r *Response) GetRawBody() (string, error) {
	data, err := r.ReadBody()
	return string(data), err
}

func (r *Response) UnmarshalCSV(v interface{}) error {
	body, err := r.ReadBody()
	if err != nil {
		return err
	}
	return r.csv.Unmarshal(body, v)
}

func (r *Response) UnmarshalError(v interface{}) error {
	body, err := r.ReadBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

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

func NewResponse(raw *http.Response) *Response {
	return &Response{raw: raw, csv: &CSV{}}
}
