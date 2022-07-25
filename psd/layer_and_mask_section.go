package psd

import (
	"fmt"
)

func ReadLayerAndMaskInfo(f *File, header *Header) {

	var version uint64
	if header.Version == 1 {
		version = uint64(f.ReadUnit32())
	} else {
		version = f.ReadUint64()
	}

	fmt.Println("version", version)
}

func ReadLayerInfo(f *File, header *Header) {}
