package net

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

// RawPacket contains the ID and raw data of a packet
type RawPacket struct {
	*bytes.Buffer

	ID uint16
}

// NewRawPacketFromBuffer creates a new packet from a given buffer
func NewRawPacketFromBuffer(id uint16, buf []byte) *RawPacket {
	return &RawPacket{
		ID:     id,
		Buffer: bytes.NewBuffer(buf),
	}
}

// NewRawPacketFromStruct creates a new packet from a given struct
func NewRawPacketFromStruct(id uint16, s interface{}) (*RawPacket, error) {
	buf := &bytes.Buffer{}

	if err := binary.Write(buf, binary.LittleEndian, s); err != nil {
		return nil, fmt.Errorf("could not create raw packet from struct (%v)", err)
	}

	return NewRawPacketFromBuffer(0x64, buf.Bytes()), nil
}

// Read a piece of data from the buffer
func (p *RawPacket) Read(v interface{}) {
	binary.Read(p.Buffer, binary.LittleEndian, v)
}

// ReadString reads a fixed-size string from the buffer, without the null characters
func (p *RawPacket) ReadString(len int, s *string) {
	b := make([]byte, len)
	p.Buffer.Read(b)
	*s = strings.Trim(string(b), "\x00")
}

// Hex returns the hexadecimal representation of the packet
func (p *RawPacket) Hex() string {
	return fmt.Sprintf("%04x", p.ID)
}

func (p *RawPacket) String() string {
	return fmt.Sprintf("%04x (len=%db)", p.ID, p.Len())
}
