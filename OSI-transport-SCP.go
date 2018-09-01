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

// SCPTransportProtocol is our experimental SabzCityProtocol!!
type SCPTransportProtocol struct {
	DestinationPort uint16
	ConnectionID    uint64
	StreamID        uint16
	PacketID        uint16
	Checksum        uint16 // For error detection process.
}

// Why Checksum is here :
// IPv6 Specification: In order to increase performance, and since current link layer technology
// and transport or application layer protocols are assumed to provide sufficient error detection.

// We must indicate some StreamID for protocol internal logical method, like open new connection, close, ...

// Use First StreamID for indicate supported version of protocol.
// Use last StreamID(65535) for send last part of that stream! It means stream has been finished and otherside must continue process.
