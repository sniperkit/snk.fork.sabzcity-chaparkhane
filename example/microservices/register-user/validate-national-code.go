/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright 2018 SabzCity. All rights reserved.

package registeruser

// ReqNationalCodeValidation : The request structure of "NationalCodeValidation()".
type ReqNationalCodeValidation struct {
	NationalCode uint `valid:"required" http:"url"`
}

// ResNationalCodeValidation : The response structure of "NationalCodeValidation()".
// http : Headers : Access-Control-Allow-Origin : "https://sabz.city"
// http : Headers : Cache : ""
type ResNationalCodeValidation struct {
	Valid bool `http:"body"`
}

// NationalCodeValidation : Check National code and say it is true or false.
// http-URI : "/{ServiceName}/{ServiceVersion}/{MethodName}/"
// http-Method : "POST"
func NationalCodeValidation(logicRequest *ReqNationalCodeValidation) (*ResNationalCodeValidation, error) {
	var logicResponse ResNationalCodeValidation

	// Check needed logic, read or save to database & ...

	return &logicResponse, nil
}
