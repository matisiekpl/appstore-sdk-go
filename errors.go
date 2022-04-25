package appstore

//ErrorResult Information with error details that an API returns in the response body whenever the API request is not successful.
// .see https://developer.apple.com/documentation/appstoreconnectapi/errorresponse
type ErrorResult struct {
	Errors []*Error `json:"errors"` //An array of one or more errors.
}

//GetError Get first error from stack
func (e *ErrorResult) GetError() *Error {
	return e.Errors[0]
}

//Error The details about one error that is returned when an API request is not successful.
// .see https://developer.apple.com/documentation/appstoreconnectapi/errorresponse/errors
type Error struct {
	Id     string       `json:"id"`     //(Required) A machine-readable code indicating the type of error. The code is a hierarchical value with levels of specificity separated by the '.' character. This value is parseable for programmatic error handling in code.
	Status string       `json:"status"` //(Required) The HTTP status code of the error. This status code usually matches the response's status code; however, if the request produces multiple errors, these two codes may differ.
	Code   string       `json:"code"`   //The unique ID of a specific instance of an error, request, and response. Use this ID when providing feedback to or debugging issues with Apple.
	Title  string       `json:"title"`  //(Required) A summary of the error. Do not use this field for programmatic error handling.
	Detail string       `json:"detail"` //(Required) A detailed explanation of the error. Do not use this field for programmatic error handling.
	Source *ErrorSource `json:"source"` //One of two possible types of values: source.parameter, provided when a query parameter produced the error, or source.JsonPointer, provided when a problem with the entity produced the error.
}

//ErrorSource An object that contains the query parameter that produced the error.
//An object that contains the JSON pointer that indicates the location of the error.
type ErrorSource struct {
	Parameter string `json:"parameter"` //The query parameter that produced the error.
	Pointer   string `json:"pointer"`   //A JSON pointer that indicates the location in the request entity where the error originates
}
