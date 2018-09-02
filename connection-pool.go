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

// ConnectionData :
var ConnectionData map[uint64]struct { // ConnectionID
	TLSSecretKey         [32]byte // 256bit encription key
	UserID               uint64   // first guest user(0) and can set other ID during connection.
	MaxConcurrentStreams uint16   // or Use PayAsGo Strategy
	DestinationPort      uint16
}

// StreamData : Can use for send or receive data on specific StreamID.
// We must use same StreamID to send and receive data for specific question and answer.
// If a Packet failed, we must make a temperary []byte and add to StreamData after that packet receive correctly.
var StreamData map[uint64]map[uint16]*[]byte // ConnectionID, StreamID, Payload

// PacketData :
type PacketData struct {
	FirstProtocol uint8

	// Network Protocols
	*IPv4NetworkProtocol
	*IPv6NetworkProtocol

	// Session Protocol

	// Transport Protocols
	*SCPTransportProtocol
	*UDPTransportProtocol
	*TCPTransportProtocol
	*QUICTransportProtocol

	// Payload can be any OSI application layer model and store in StreamData structure
}
