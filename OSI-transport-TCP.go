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

// It is the layer that we must detect transport protocol e.g. TCP, UDP, SPX, ...
// We must call next layer by detected rules.

// If project don't have any logic that support data on e.g. UDP (DNS, ...) we reject request with related erorr.

// TCPTransportProtocol : https://en.wikipedia.org/wiki/Transmission_Control_Protocol#TCP_segment_structure
type TCPTransportProtocol struct {
	SourcePort           uint16
	DestinationPort      uint16
	SequenceNumber       uint32
	AcknowledgmentNumber uint32
	DataOffset           uint8
	Reserved             uint8
	Flags                uint16 // I think it must be multi bool value not integer!
	WindowSize           uint16
	Checksum             uint16
	UrgentPointer        uint16
	// Options
}
