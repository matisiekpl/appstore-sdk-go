package appstore

import "fmt"

//ResourceAbstract base resource
type ResourceAbstract struct {
	transport *Transport
	config    *Config
}

//Get HTTP method wrapper
func (r *ResourceAbstract) get(path string, query map[string]interface{}) (*Response, error) {
	rsp, err := r.transport.Get(path, query)
	if err != nil {
		return nil, fmt.Errorf("ResourceAbstract@get request: %v", err)
	}
	return NewResponse(rsp), nil
}

//Create new resource abstract
func newResourceAbstract(transport *Transport, config *Config) *ResourceAbstract {
	return &ResourceAbstract{transport: transport, config: config}
}
