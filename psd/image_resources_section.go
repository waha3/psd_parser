package psd

import (
	"fmt"
	"io"
)

type ImageResourcesSetion struct {
	// Length of image resource section. The length may be zero
	length uint32
	// image resources
	Blocks []ImageResourcesBlock
}

type ImageResourcesBlock struct {
	// Signature: '8BIM'
	Signature string
	// Unique identifier for the resource. Image resource IDs contains a list of resource IDs used by Photoshop.
	Id uint16
	// Name: Pascal string, padded to make the size even (a null name consists of two bytes of 0)
	Name string
	// Actual size of resource data that follows
	Size uint32
	// The resource data, described in the sections on the individual resource types. It is padded to make the size even
	Data string
}

func (s *ImageResourcesSetion) ReadImageResourcesSetion(f *File) {
	resourcesLen := f.ReadUnit32()

	fmt.Println("resourcesLen", resourcesLen)

	blockSize := uint32(0)
	for {
		blockSize += ReadImageResourcesBlock(f)
		fmt.Println("blocksize", blockSize)

		if blockSize >= resourcesLen {
			break
		}
	}
}

func ReadImageResourcesBlock(f *File) uint32 {
	signature := f.ReadString(4)
	id := f.ReadUint16()

	name := f.ReadPascalString()
	size := f.ReadUnit32()

	if size%2 != 0 {
		size = size + 1
	}

	f.Buf.Seek(int64(size), io.SeekCurrent)

	fmt.Println(signature, id, name, size)
	return 4 + 2 + 4 + size
}
