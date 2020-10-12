package appstore_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	if query != nil {
		q := u.Query()
		for k, v := range query {
			q.Set(k, fmt.Sprintf("%v", v))
		}
		u.RawQuery = q.Encode()
	}
	return u, err
}

func (rb *RequestBuilder) buildBody(data map[string]interface{}) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder@buildBody json convert: %v", err)
	}
	return bytes.NewBuffer(b), nil
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
	//build body
	bodyReader, err := t.rb.buildBody(body)
	if err != nil {
		return nil, fmt.Errorf("transport@request build request body: %v", err)
	}
	//build request
	req, err := http.NewRequest(strings.ToUpper(method), uri.String(), bodyReader)
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

func (r *Response) ReadBody() ([]byte, error) {
	defer r.raw.Body.Close()
	return ioutil.ReadAll(r.raw.Body)
}

func NewResponse(raw *http.Response) *Response {
	return &Response{raw: raw, csv: &CSV{}}
}
