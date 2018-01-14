package packet

import (
	"testing"

	"github.com/drgomesp/midgard/net"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestParseAccountLogin(t *testing.T) {
	Convey("Given a raw packet", t, func() {
		rawPacket, _ := net.NewRawPacketFromStruct(0x64, struct {
			v      uint32
			us, ps [AccountLoginStrLength]byte
			ct     byte
		}{
			uint32(2018),
			[24]byte{'a', 'b', 'c'},
			[24]byte{'1', '2', '3'},
			byte('t'),
		})

		Convey("When trying to parse an account login from that raw packet", func() {
			packet := &AccountLogin{}
			packet.Parse(rawPacket)

			Convey("Then the account login packet struct should be constructed correctly", func() {
				assert.Equal(t, uint32(2018), packet.Version, "packet.Version actual value is not equal to expected value")
				assert.Equal(t, "abc", packet.Username, "packet.Username actual value is not equal to expected value")
				assert.Equal(t, "123", packet.Password, "packet.Password actual value is not equal to expected value")
				assert.Equal(t, byte('t'), packet.ClientType, "packet.ClientType actual value is not equal to expected value")
			})
		})
	})
}
