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

// IPv6NetworkProtocol :
type IPv6NetworkProtocol struct {
	Version              uint8
	TrafficClass         uint8
	FlowLabel            uint32
	PayloadLength        uint32
	NextHeader           uint8
	HopLimit             uint8
	SourceIPAddress      [16]byte // 128 bit || 16 byte
	DestinationIPAddress [16]byte // 128 bit || 16 byte
}

// https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers
