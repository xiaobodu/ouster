package packet

import (
    "bytes"
    "github.com/tiancaiamao/ouster/data"
    "testing"
)

func TestGCUpdateInfoPacket(t *testing.T) {
    // raw := []byte{83, 221, 118, 0, 0, 4, 203, 171, 208, 222, 1, 1, 125, 1, 224, 1, 0, 76, 29, 0, 0, 13, 0, 13, 0, 13, 0, 14, 0, 14, 0, 14, 0, 20, 0, 20, 0, 20, 0, 1, 204, 41, 0, 0, 248, 4, 0, 0, 221, 5, 0, 0, 211, 15, 0, 0, 54, 0, 54, 0, 50, 0, 50, 0, 0, 0, 0, 0, 244, 1, 0, 0, 0, 50, 0, 0, 0, 0, 50, 0, 0, 0, 0, 40, 0, 0, 0, 7, 102, 3, 0, 0, 0, 30, 0, 0, 0, 0, 64, 66, 15, 0, 13, 114, 73, 57, 0, 0, 0, 72, 215, 1, 99, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 222, 118, 0, 0, 35, 0, 0, 0, 184, 11, 0, 0, 0, 0, 4, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 223, 118, 0, 0, 22, 0, 0, 0, 172, 13, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 224, 118, 0, 0, 4, 2, 0, 0, 1, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 20, 0, 0, 0, 4, 0, 225, 118, 0, 0, 15, 0, 0, 0, 124, 21, 0, 0, 0, 0, 4, 0, 0, 0, 0, 1, 0, 0, 0, 2, 3, 226, 118, 0, 0, 14, 0, 0, 0, 148, 17, 0, 0, 0, 0, 4, 0, 0, 0, 0, 1, 0, 0, 0, 4, 3, 227, 118, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 9, 0, 0, 0, 9, 4, 228, 118, 0, 0, 1, 5, 0, 0, 1, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 9, 0, 0, 0, 9, 5, 4, 229, 118, 0, 0, 11, 0, 0, 0, 215, 5, 0, 0, 0, 0, 4, 0, 0, 0, 0, 1, 0, 0, 0, 2, 230, 118, 0, 0, 17, 0, 0, 0, 184, 9, 0, 0, 0, 0, 4, 0, 0, 0, 0, 1, 0, 0, 0, 3, 230, 118, 0, 0, 17, 0, 0, 0, 184, 9, 0, 0, 0, 0, 4, 0, 0, 0, 0, 1, 0, 0, 0, 4, 231, 118, 0, 0, 12, 0, 0, 0, 172, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 1, 0, 0, 0, 7, 0, 0, 0, 64, 0, 17, 232, 200, 7, 1, 17, 2, 49, 4, 1, 12, 15, 4, 3, 94, 0, 96, 0, 56, 0, 4, 5, 0, 6, 0, 7, 0, 8, 0, 3, 4, 194, 229, 177, 246, 172, 1, 17, 0, 222, 0, 6, 196, 171, 184, 163, 184, 224, 174, 1, 36, 0, 228, 0, 8, 199, 193, 183, 185, 181, 229, 184, 175, 158, 0, 23, 0, 235, 0, 0, 17, 0, 0, 0, 0, 152, 42, 0, 0, 0, 0, 0, 0, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 52, 1, 5, 0, 0, 0, 1, 0, 221, 118, 0, 0}
    raw := []byte{79, 170, 47, 0, 0, 6, 'o', 'u', 's', 't', 'e', 'r', 150, 0, 101, 0, 0, 76, 29, 0, 0, 10, 0, 10, 0, 10, 0, 25, 0, 25, 0, 25, 0, 10, 0, 10, 0, 10, 0, 59, 1, 59, 1, 186, 0, 111, 0, 50, 204, 41, 0, 0, 125, 0, 0, 0, 244, 1, 0, 0, 92, 0, 0, 0, 13, 15, 39, 10, 39, 0, 0, 1, 66, 0, 0, 4, 0, 0, 0, 0, 100, 0, 0, 0, 0, 2, 171, 47, 0, 0, 67, 0, 0, 0, 1, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 9, 0, 0, 0, 0, 0, 172, 47, 0, 0, 66, 0, 0, 0, 1, 0, 0, 0, 0, 0, 255, 255, 255, 255, 0, 9, 0, 0, 0, 0, 1, 1, 173, 47, 0, 0, 59, 0, 0, 0, 145, 15, 0, 0, 0, 0, 4, 0, 0, 0, 0, 1, 0, 0, 0, 3, 0, 2, 148, 1, 54, 66, 109, 0, 246, 224, 0, 62, 0, 60, 56, 191, 7, 8, 19, 15, 56, 16, 0, 0, 13, 6, 0, 4, 5, 0, 6, 0, 7, 0, 8, 0, 2, 6, 192, 173, 206, 172, 209, 199, 141, 2, 32, 0, 22, 0, 4, 186, 194, 192, 173, 146, 2, 24, 0, 42, 0, 0, 17, 0, 0, 0, 0, 72, 126, 0, 0, 0, 0, 0, 0, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0}
    buf := bytes.NewBuffer(raw)
    pkt := &GCUpdateInfoPacket{}
    err := pkt.Read(buf, 0)
    if err != nil {
        t.Fail()
    }

    t.Logf("%#v\n", pkt)
    t.Logf("%#v\n", buf.Bytes())
}

func TestCGConnectPacket(t *testing.T) {
    raw := []byte{0, 230, 91, 4, 0, 4, 203, 171, 208, 222, 8, 0, 39, 203, 215, 129}
    // [ 0 0 0 240 1 4 183 232 191 241 0 80 86 192 0 8]
    buf := bytes.NewBuffer(raw)
    pkt := &CGConnectPacket{}
    err := pkt.Read(buf, 0)
    if err != nil {
        t.Fail()
    }
    // t.Logf("%#v\n", pkt)
}

func TestGetSize(t *testing.T) {
    buf := &bytes.Buffer{}
    gearinfo := &data.GearInfo{}
    gearinfo.Write(buf)
    if buf.Len() != gearinfo.getSize() {
        t.Errorf("GearInfo的getSize不对:期待%d 实际%d\n", buf.Len(), gearinfo.getSize())
    }
}
