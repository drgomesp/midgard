package net

// Packet represents small pieces of information that can be transmitted across networks
type Packet interface {
	// Parse a raw packet into the actual packet structure
	Parse(p *RawPacket) error
}
