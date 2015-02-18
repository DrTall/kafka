package proto

import (
	"bytes"
	"testing"
)

var (
	keyint = 12
	bint8  = []byte{0x0c}
	bint64 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c}

	keystr = "Ash nazg durbatulûk, ash nazg gimbatul, ash nazg thrakatulûk, agh burzum-ishi krimpatul" // google LOTR one ring
	bstr   = []byte{0x00, 0x59, 0x41, 0x73, 0x68, 0x20, 0x6e, 0x61, 0x7a, 0x67, 0x20, 0x64, 0x75, 0x72, 0x62, 0x61, 0x74, 0x75, 0x6c, 0xc3, 0xbb, 0x6b, 0x2c, 0x20, 0x61, 0x73, 0x68, 0x20, 0x6e, 0x61, 0x7a, 0x67, 0x20, 0x67, 0x69, 0x6d, 0x62, 0x61, 0x74, 0x75, 0x6c, 0x2c, 0x20, 0x61, 0x73, 0x68, 0x20, 0x6e, 0x61, 0x7a, 0x67, 0x20, 0x74, 0x68, 0x72, 0x61, 0x6b, 0x61, 0x74, 0x75, 0x6c, 0xc3, 0xbb, 0x6b, 0x2c, 0x20, 0x61, 0x67, 0x68, 0x20, 0x62, 0x75, 0x72, 0x7a, 0x75, 0x6d, 0x2d, 0x69, 0x73, 0x68, 0x69, 0x20, 0x6b, 0x72, 0x69, 0x6d, 0x70, 0x61, 0x74, 0x75, 0x6c}
	bbyte  = []byte{0x00, 0x00, 0x00, 0x59, 0x41, 0x73, 0x68, 0x20, 0x6e, 0x61, 0x7a, 0x67, 0x20, 0x64, 0x75, 0x72, 0x62, 0x61, 0x74, 0x75, 0x6c, 0xc3, 0xbb, 0x6b, 0x2c, 0x20, 0x61, 0x73, 0x68, 0x20, 0x6e, 0x61, 0x7a, 0x67, 0x20, 0x67, 0x69, 0x6d, 0x62, 0x61, 0x74, 0x75, 0x6c, 0x2c, 0x20, 0x61, 0x73, 0x68, 0x20, 0x6e, 0x61, 0x7a, 0x67, 0x20, 0x74, 0x68, 0x72, 0x61, 0x6b, 0x61, 0x74, 0x75, 0x6c, 0xc3, 0xbb, 0x6b, 0x2c, 0x20, 0x61, 0x67, 0x68, 0x20, 0x62, 0x75, 0x72, 0x7a, 0x75, 0x6d, 0x2d, 0x69, 0x73, 0x68, 0x69, 0x20, 0x6b, 0x72, 0x69, 0x6d, 0x70, 0x61, 0x74, 0x75, 0x6c}
)

var b = bytes.NewBuffer(nil)

func getTestEncoder() *encoder {
	b.Reset()
	return NewEncoder(b)
}

func failByteNotEqual(t *testing.T, a, b []byte) {
	if bytes.Equal(a, b) {
		return
	}
	t.Fatalf("% x != % x", a, b)
}

func TestEncoder(t *testing.T) {
	e := getTestEncoder()
	e.Encode(int8(keyint))
	failByteNotEqual(t, b.Bytes(), bint8)

	e = getTestEncoder()
	e.Encode(int64(keyint))
	failByteNotEqual(t, b.Bytes(), bint64)

	e = getTestEncoder()
	e.Encode(string(keystr))
	failByteNotEqual(t, b.Bytes(), bstr)

	e = getTestEncoder()
	e.Encode([]byte(keystr))
	failByteNotEqual(t, b.Bytes(), bbyte)
}

func getTestDecoder(b []byte) *decoder {
	buf := bytes.NewBuffer(b)
	return NewDecoder(buf)
}

func TestDecoder(t *testing.T) {
	d := getTestDecoder(bint8)
	if d.DecodeInt8() != int8(keyint) {
		t.Fatalf("int8 decoding failed")
	}

	d = getTestDecoder(bint64)
	if d.DecodeInt64() != int64(keyint) {
		t.Fatalf("int64 decoding failed")
	}

	d = getTestDecoder(bstr)
	if d.DecodeString() != keystr {
		t.Fatalf("string decoding failed")
	}

	d = getTestDecoder(bbyte)
	failByteNotEqual(t, d.DecodeBytes(), []byte(keystr))
}
