package psd

import (
	"fmt"
)

func ReadLayerAndMaskInfo(f *File, header *Header) {
	var length uint64
	if header.Version == 1 {
		length = uint64(f.ReadUnit32())
	} else {
		length = f.ReadUint64()
	}

	fmt.Println("layer and mask length", length)
	ReadLayerInfo(f, header)
}

func ReadLayerInfo(f *File, header *Header) {
	var length uint64
	if header.Version == 1 {
		length = uint64(f.ReadUnit32())
	} else {
		length = f.ReadUint64()
	}
	layerCount := f.ReadUint16()

	fmt.Println("layer info", length, layerCount)
}
