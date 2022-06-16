package appstore

import (
	"fmt"
	"net/http"
)

//ResourceAbstract base resource
type ResourceAbstract struct {
	transport *Transport
	config    *Config
}

//UnmarshalResponse method
func (ra *ResourceAbstract) unmarshalResponse(resp *http.Response, v interface{}, filterLines bool) error {
	contentType := resp.Header.Get("Content-Type")
	responseHandler := NewResponseHandler(contentType, filterLines)

	bodyBytes, err := responseHandler.ReadBody(resp)
	if err != nil {
		return fmt.Errorf("ResourceAbstract.unmarshalResponse read body: %v", err)
	}
	//reset the response body to the original unread state
	body, err := responseHandler.RestoreBody(bodyBytes)
	if err != nil {
		return fmt.Errorf("ResourceAbstract.unmarshalResponse read body: %v", err)
	}
	resp.Body = body
	return responseHandler.UnmarshalBody(bodyBytes, v)
}

//newResourceAbstract create new resource abstract
func newResourceAbstract(transport *Transport, config *Config) ResourceAbstract {
	return ResourceAbstract{transport: transport, config: config}
}
