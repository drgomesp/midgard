package packet

import (
	"github.com/drgomesp/midgard/net"
)

const (
	// AccountLoginStrLength is the maximum length for string fields of the AccountLogin packet
	AccountLoginStrLength = 24
)

// AccountLogin packet (0x64), sent by the client when a login request is made
type AccountLogin struct {
	Version uint32
	// Username, Password [AccountLoginStrLength]byte
	Username, Password string
	ClientType         byte
}

// Parse a raw packet into an account login packet
func (accountLogin *AccountLogin) Parse(p *net.RawPacket) error {
	p.Read(&accountLogin.Version)
	p.ReadString(AccountLoginStrLength, &accountLogin.Username)
	p.ReadString(AccountLoginStrLength, &accountLogin.Password)
	p.Read(&accountLogin.ClientType)

	return nil
}
