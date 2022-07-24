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

func (f *File) ReadPascalString() string {
	var strLen uint8
	err := binary.Read(f.Buf, binary.BigEndian, &strLen)

	if err != nil {
		log.Fatal(err)
	}

	data := f.ReadString(int32(strLen))

	// padding even
	if (strLen+1)%2 != 0 {
		f.ReadUint8()
	}

	return data
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

func (f *File) ReadUnit32() uint32 {
	var data uint32
	err := binary.Read(f.Buf, binary.BigEndian, &data)

	if err != nil {
		log.Fatal(err)
	}
	return data
}
