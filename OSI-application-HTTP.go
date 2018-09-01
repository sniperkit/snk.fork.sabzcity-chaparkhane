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

// Protocol Standard
// http2 : https://httpwg.org/specs/rfc7540.html

// Ready data for logics & do some logic
// - Route by URL
// - Encode||Decode body by mime type header

// Add Server Header to response : "ChaparKhane" || SCP means "SabzCityPlatform".

// If project don't have any logic that support data on e.g. HTTP (restful, ...) we reject request with related erorr.

// It can use for architectures like restful, ...
