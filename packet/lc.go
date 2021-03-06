package packet

import (
    "bytes"
    "encoding/binary"
    "github.com/tiancaiamao/ouster/data"
    "io"
)

type LCLoginOKPacket struct {
    DummyRead
}

func (loginOk LCLoginOKPacket) PacketID() PacketID {
    return PACKET_LC_LOGIN_OK
}
func (loginOk LCLoginOKPacket) PacketSize() uint32 {
    return 5
}

func (loginOk LCLoginOKPacket) String() string {
    return "loginOK"
}

func (loginOk LCLoginOKPacket) Write(buf io.Writer, code uint8) error {
    buf.Write([]byte{0, 0, 231, 254, 255})
    return nil
}

type LCVersionCheckOKPacket struct {
    DummyRead
}

func (v LCVersionCheckOKPacket) PacketID() PacketID {
    return PACKET_LC_VERSION_CHECK_OK
}
func (v LCVersionCheckOKPacket) PacketSize() uint32 {
    return 0
}
func (v LCVersionCheckOKPacket) String() string {
    return "version check ok"
}
func (v LCVersionCheckOKPacket) Write(buf io.Writer, code uint8) error {
    return nil
}

type LCWorldListPacket struct {
    DummyRead
}

func (wl LCWorldListPacket) PacketSize() uint32 {
    return 13
}
func (wl LCWorldListPacket) PacketID() PacketID {
    return PACKET_LC_WORLD_LIST
}
func (wl LCWorldListPacket) String() string {
    return "world list"
}
func (wl LCWorldListPacket) Write(buf io.Writer, code uint8) error {
    buf.Write([]byte{1, 1, 1, 8, 185, 237, 247, 200, 193, 182, 211, 252, 0})
    return nil
}

type LCServerListPacket struct {
    DummyRead

    CurrentWorld uint8
    Size         uint8
    list         []string
}

func (sl *LCServerListPacket) PacketID() PacketID {
    return PACKET_LC_SERVER_LIST
}
func (sl *LCServerListPacket) String() string {
    return "server list"
}
func (sl *LCServerListPacket) PacketSize() uint32 {
    return 22
}
func (sl *LCServerListPacket) Write(buf io.Writer, code uint8) error {
    buf.Write([]byte{1, 2, 0, 6, 183, 226, 178, 226, 199, 248, 0, 1, 8, 185, 237, 247, 200, 193, 182, 211, 252, 0})
    return nil
}

type LCPCListPacket struct {
    DummyRead
    list []data.PCInfo
}

func (pl *LCPCListPacket) PacketID() PacketID {
    return PACKET_LC_PC_LIST
}
func (pl *LCPCListPacket) String() string {
    return "pc list"
}

// TODO
func (pl *LCPCListPacket) PacketSize() uint32 {
    buf := &bytes.Buffer{}
    pl.Write(buf, 0)
    return uint32(buf.Len())
}
func (pl *LCPCListPacket) Write(buf io.Writer, code uint8) error {
    buf.Write([]byte{
        83, 79, 86,
        6, 's', 'l', 'a', 'y', 'e', 'r', 0,
        76, 29, 0, 0, 9, 0, 11, 0, 10, 0, 50, 170, 9, 0, 0, 53, 15, 0, 0, 47, 12, 0, 0, 18, 0, 18, 0, 20, 0, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 144, 1, 170, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100,
        6, 'o', 'u', 's', 't', 'e', 'r', 1, 76, 29, 0, 0, 0, 121, 1, 101, 0, 121, 1, 0, 0, 8, 10, 0, 25, 0, 10, 0, 59, 1, 59, 1, 0, 0, 0, 0, 150, 50, 0, 0, 0, 0, 68, 0, 0, 0, 15, 39, 15, 39, 100,
        7, 'v', 'a', 'm', 'p', 'i', 'r', 'e', 2, 76, 29, 0, 0, 0, 0, 0, 164, 1, 1, 121, 1, 20, 0, 20, 0, 20, 0, 216, 1, 216, 1, 150, 50, 125, 0, 0, 0, 217, 0, 0, 0, 15, 39, 100,
    })
    return nil
}

type LCReconnectPacket struct {
    NotImplementWrite

    Ip   string
    Port uint32
    Key  uint32
}

func (rc *LCReconnectPacket) PacketID() PacketID {
    return PACKET_LC_RECONNECT
}
func (rc *LCReconnectPacket) PacketSize() uint32 {
    return uint32(1 + len(rc.Ip) + 8)
}
func (rc *LCReconnectPacket) String() string {
    return "reconnect"
}
func (rc *LCReconnectPacket) Write(buf io.Writer, code uint8) error {
    //[13 49 57 50 46 49 54 56 46 49 46 49 50 51 14 39 0 0 0 32 6 11]
    binary.Write(buf, binary.LittleEndian, uint8(len(rc.Ip)))
    io.WriteString(buf, rc.Ip)
    binary.Write(buf, binary.LittleEndian, rc.Port)
    binary.Write(buf, binary.LittleEndian, rc.Key)
    return nil
}

func (ret *LCReconnectPacket) Read(reader io.Reader, code uint8) error {
    var sz uint8
    var buf [256]byte
    binary.Read(reader, binary.LittleEndian, &sz)
    _, err := reader.Read(buf[:sz])
    if err != nil {
        return err
    }
    ret.Ip = string(buf[:sz])
    binary.Read(reader, binary.LittleEndian, &ret.Port)
    binary.Read(reader, binary.LittleEndian, &ret.Key)
    return nil
}
