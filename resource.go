package appstore_sdk

import "fmt"

type ResourceAbstract struct {
	transport *Transport
	config    *Config
}

func (r *ResourceAbstract) get(path string, query map[string]interface{}) (*Response, error) {
	rsp, err := r.transport.Get(path, query)
	if err != nil {
		return nil, fmt.Errorf("ResourceAbstract@get request: %v", err)
	}
	return &Response{raw: rsp}, nil
}

func newResourceAbstract(transport *Transport, config *Config) *ResourceAbstract {
	return &ResourceAbstract{transport: transport, config: config}
}
