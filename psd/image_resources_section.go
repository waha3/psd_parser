package psd

import (
	"fmt"
	"io"
)

type ImageResourcesSetion struct {
	// Length of image resource section. The length may be zero
	length uint32 `json:"length"`
	// image resources
	Blocks []ImageResourcesBlock `json:"blocks"`
}

type ImageResourcesBlock struct {
	// Signature: '8BIM'
	Signature string `json:"signature"`
	// Unique identifier for the resource. Image resource IDs contains a list of resource IDs used by Photoshop.
	Id uint16 `json:"id"`
	// Name: Pascal string, padded to make the size even (a null name consists of two bytes of 0)
	Name string `json:"name"`
	// Actual size of resource data that follows
	Size uint32 `json:"size"`
	// The resource data, described in the sections on the individual resource types. It is padded to make the size even
	Data string `json:"data"`
}

func (s *ImageResourcesSetion) ReadImageResourcesSetion(f *File) {
	resourcesLen := f.ReadUnit32()

	// fmt.Println("resourcesLen", resourcesLen)

	blockSize := uint32(0)
	for {
		blockSize += ReadImageResourcesBlock(f)

		fmt.Println("blocksize：resourcesLen", blockSize, resourcesLen)

		if blockSize >= resourcesLen {
			fmt.Println("blocksize here", blockSize)
			break
		}
	}
}

func ReadImageResourcesBlock(f *File) uint32 {
	signature := f.ReadString(4)
	id := f.ReadUint16()

	name, nameLen := f.ReadPascalString()
	size := f.ReadUnit32()

	if size%2 != 0 {
		f.ReadUnit32()
		size = size + 1
	}

	fmt.Println("block", signature, id, name, nameLen, size)
	// fmt.Println(len(signature), id, len(name), size)

	switch {
	// case id == 1032:
	// 	ReadGridAndGuides(f)
	// case 1033:
	// case 1036:
	// 	ReadThumbnail(f)
	// case id == 1038:
	// case id == 1073:
	// 	ReadColorSampleHeader(f)
	// case id >= 2000 && id < 2997:
	// 	ReadPathResource(f)
	default:
		f.Buf.Seek(int64(size), io.SeekCurrent)
	}
	return 4 + 2 + uint32(nameLen)
}

// Photoshop stores grid and guides information for an image in an image resource block.
// Each of these resource blocks consists of an initial 16-byte grid and guide header,
// which is always present, followed by 5-byte blocks of specific guide information for guide direction and location,
// which are present if there are guides ( fGuideCount > 0) .
func ReadGridAndGuides(f *File) {
	version := f.ReadUnit32()
	horizontal, vertical := f.ReadUnit32(), f.ReadUnit32()
	fGuideCount := f.ReadUnit32()

	for i := 0; i < int(fGuideCount); i++ {
		location := f.ReadUnit32()
		direction := f.ReadUint8()

		fmt.Println("resourse block", location, direction)
	}

	fmt.Println("grid", version, horizontal, vertical, fGuideCount)
}

// Adobe Photoshop (version 5.0 and later) stores thumbnail information for preview display in an image resource block that consists of an initial 28-byte header,
// followed by a JFIF thumbnail in RGB (red, green, blue) order for both Macintosh and Windows.
func ReadThumbnail(f *File) {
	formatMap := map[uint32]string{
		0: "kRawRGB",
		1: "kJpegRGB",
	}
	format := f.ReadUnit32()
	width := f.ReadUnit32()
	height := f.ReadUnit32()
	widthbytes := f.ReadUnit32()
	totalSize := f.ReadUnit32()
	compressedSize := f.ReadUnit32()
	bitPerPixel := f.ReadUint16()
	planes := f.ReadUint16()

	// TODO
	// JFIF data in RGB format.
	// For resource ID 1033 the data is in BGR format.

	fmt.Println("thumb", formatMap[format], width, height, widthbytes, totalSize, compressedSize, bitPerPixel, planes)
}

// Color samplers resource format
// Adobe Photoshop (version 5.0 and later) stores color samplers information for an image in an image resource block
// that consists of an initial 8-byte color samplers header followed by a variable length block of specific color samplers information
func ReadColorSampleHeader(f *File) {
	version := f.ReadUnit32()
	numOfColor := f.ReadUnit32()

	fmt.Println("colorSampleHeader", version, numOfColor)

	ReadColorSampleBlock(f)
}

func ReadColorSampleBlock(f *File) {
	version := f.ReadUnit32()
	horizontal, vertical := f.ReadUint16(), f.ReadUint16()
	colorSpace := f.ReadUint16()
	depth := f.ReadUint16()

	fmt.Println("colorBlock", version, horizontal, vertical, colorSpace, depth)
}

// Photoshop stores the paths saved with an image in an image resource block
// These resource blocks consist of a series of 26-byte path point records
func ReadPathResource(f *File) {

}
