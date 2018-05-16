// Copyright 2018 SabzCity. All rights reserved.

package logic

// ReqNationalCodeValidation : The request structure of "NationalCodeValidation()".
type ReqNationalCodeValidation struct {
	NationalCode uint `valid:"required" restful:"url" jsonrpc:"body" grpc:"body" protobuf:"1" json:"NationalCode" xml:"NationalCode"`
}

// ResNationalCodeValidation : The response structure of "NationalCodeValidation()".
type ResNationalCodeValidation struct {
	Valid bool `restful:"body" jsonrpc:"body" grpc:"body" protobuf:"1" json:"Valid" xml:"Valid"` // UUID of user
}

// NationalCodeValidation : Check National code and say it is true or false.
// restful-URI : "/{ServiceName}/{ServiceVersion}/{MethodName}/"
// restful-Method : "POST"
// restful-Headers-Access-Control-Allow-Origin : "https://sabz.city"
// jsonrpc-URI : "/"
// jsonrpc-Method : "POST"
// jsonrpc-Headers-Access-Control-Allow-Origin : "*"
// ...
func NationalCodeValidation(logicRequest *ReqNationalCodeValidation) (*ResNationalCodeValidation, error) {
	var logicResponse ResNationalCodeValidation

	// Check needed logic, read or save to database & ...

	return &logicResponse, nil
}
