/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright 2018 SabzCity

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chaparkhane

// ExtendedError : This is a extended implementation of error.
type ExtendedError struct {
	Text        string
	Information interface{}
	Code        uint
	HTTPStatus  uint
}

// NewError : Returns an error that formats as the given text, Code and httpStatus code.
func NewError(text string, sabzCityCode uint, httpStatus uint) error {
	e := ExtendedError{
		Text:       text,
		Code:       sabzCityCode,
		HTTPStatus: httpStatus}
	return &e
}

// Error : Return text of error.
func (e *ExtendedError) Error() string {
	return e.Text
}

// AddInformationToError : Add information to existing error and return it as new error(pointer)!
func AddInformationToError(err error, information interface{}) error {
	mainError := err.(*ExtendedError)

	return &ExtendedError{
		Text:        mainError.Text,
		Information: information,
		Code:        mainError.Code,
		HTTPStatus:  mainError.HTTPStatus}
}

// IsEqualError : Compare two error.
func IsEqualError(err1, err2 error) bool {
	if err1.(*ExtendedError).Code == err2.(*ExtendedError).Code {
		return true
	}

	return false
}
