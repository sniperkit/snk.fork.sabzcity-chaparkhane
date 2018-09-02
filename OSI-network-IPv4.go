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

// JUST DO IN PERSIAOS!!!!!!!!
// In PersiaOS we don't care about e.g. tcp overhead to manage packets!! Application must do it!!

// IPv4NetworkProtocol :
type IPv4NetworkProtocol struct {
	Version              uint8  // https://en.wikipedia.org/wiki/IPv4#Version
	IHL                  uint8  // https://en.wikipedia.org/wiki/IPv4#IHL
	TOS                  uint8  // type-of-service
	TotalLength          uint16 // https://en.wikipedia.org/wiki/IPv4#Total_Length
	Identification       uint16 // https://en.wikipedia.org/wiki/IPv4#Identification
	Flags                uint8  // https://en.wikipedia.org/wiki/IPv4#Flags
	FragmentOffset       uint16 // https://en.wikipedia.org/wiki/IPv4#Fragment_Offset
	TimeToLive           uint8  // https://en.wikipedia.org/wiki/IPv4#TTL
	Protocol             uint8  // https://en.wikipedia.org/wiki/IPv4#Protocol
	HeaderChecksum       uint16 // https://en.wikipedia.org/wiki/IPv4#Header_Checksum
	SourceIPAddress      uint32 // https://en.wikipedia.org/wiki/IPv4#Source_address
	DestinationIPAddress uint32 // https://en.wikipedia.org/wiki/IPv4#Destination_address
	Options              []byte // options, extension headers
}

// https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers
