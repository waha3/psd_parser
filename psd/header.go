package psd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type Header struct {
	// always equal to '8BPS
	Signature string `json:"signature"`
	// always equal to 1 or 2 (**PSB** version is 2)
	Version uint16 `json:"version"`
	// The number of channels in the image, including any alpha channels. Supported range is 1 to 56.
	Channels uint16 `json:"channels"`
	// The height of the image in pixels. Supported range is 1 to 30,000
	Height uint32 `json:"height"`
	// The width of the image in pixels. Supported range is 1 to 30,000
	Width uint32 `json:"width"`
	// the number of bits per channel. Supported values are 1, 8, 16 and 32
	Depth uint16 `json:"depth"`
	// The color mode of the file. Supported values are: Bitmap = 0; Grayscale = 1; Indexed = 2; RGB = 3; CMYK = 4; Multichannel = 7; Duotone = 8; Lab = 9
	ColorMode string `json:"colorMode"`
	// The length of the following color data
	// Only indexed color and duotone (see the mode field in the File header section) have color mode data.
	// For all other modes, this section is just the 4-byte length field, which is set to zero
	ColorModeData uint32 `json:"colorModeData"`
}

func (h *Header) ReadHeader(f *File) *Header {
	colorModeMap := map[uint16]string{
		0: "Bitmap",
		1: "Grayscale",
		2: "Indexed",
		3: "RGB",
		4: "CMYK",
		7: "Multichannel",
		8: "Duotone",
		9: "Lab",
	}
	signature := f.ReadString(4)

	if signature != "8BPS" {
		log.Fatal("it should be .psd file")
	}

	fmt.Println(signature)
	// psd = 1, psb = 2
	version := f.ReadUint16()
	// Reserved: must be zero 6bytes
	f.Buf.Seek(6, io.SeekCurrent)
	channels := f.ReadUint16()
	height := f.ReadUnit32()
	width := f.ReadUnit32()
	depth := f.ReadUint16()
	mode := f.ReadUint16()
	colorModeData := f.ReadUnit32()

	if colorModeData != 0 {
		log.Fatal("目前不支持 indexed color duotone color")
	}
	// @todo 解析两种模式

	headerInfo := &Header{
		signature,
		version,
		channels,
		height,
		width,
		depth,
		colorModeMap[mode],
		colorModeData,
	}

	data, err := json.Marshal(headerInfo)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("headinfo:%s\n", data)
	return headerInfo
}
