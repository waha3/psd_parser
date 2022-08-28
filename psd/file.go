package psd

import (
	"bytes"
	"encoding/binary"
	"log"
)

type File struct {
	Buf *bytes.Reader
}

func (f *File) ReadString(n int32) string {
	data := make([]byte, n)
	err := binary.Read(f.Buf, binary.BigEndian, &data)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func (f *File) ReadPascalString() (string, uint8) {
	var strLen uint8
	err := binary.Read(f.Buf, binary.BigEndian, &strLen)

	if err != nil {
		log.Fatal(err)
	}

	// padding even
	if strLen == 0 {
		strLen = 1
	}

	str := f.ReadString(int32(strLen))

	return str, strLen
}

func (f *File) ReadUint8() uint8 {
	var data uint8
	err := binary.Read(f.Buf, binary.BigEndian, &data)

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (f *File) ReadUint16() uint16 {
	var data uint16
	err := binary.Read(f.Buf, binary.BigEndian, &data)

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (f *File) ReadInt16() int16 {
	var data int16
	err := binary.Read(f.Buf, binary.BigEndian, &data)

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (f *File) ReadUnit32() uint32 {
	var data uint32
	err := binary.Read(f.Buf, binary.BigEndian, &data)

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (f *File) ReadUint64() uint64 {
	var data uint64
	err := binary.Read(f.Buf, binary.BigEndian, &data)

	if err != nil {
		log.Fatal(err)
	}
	return data
}
