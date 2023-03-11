package checkpicext

import (
	"bytes"
	"errors"
	"strings"
)

const (
	UnKnow = "unknow"
	JPEG   = "jpg"
	PNG    = "png"
	GIF    = "gif"
	TIFF   = "tiff"
	WebP   = "webp"
	HEIC   = "heic"
	HEIF   = "heif"
)

var HEICBitMaps = map[string]struct{}{
	"heic": {},
	"heis": {},
	"heix": {},
	"hevc": {},
	"hevx": {},
}

var HEIFBitMaps = map[string]struct{}{
	"mif1": {},
	"msf1": {},
}

type imageType []byte

func GetImageFormat(image []byte) string {

	buffer, err := getBuffer(image)
	if err != nil {
		return UnKnow
	}

	switch {
	case buffer.isJPEG():
		return JPEG
	case buffer.isPNG():
		return PNG
	case buffer.isGIF():
		return GIF
	case buffer.isTIFF():
		return TIFF
	case buffer.isWebP():
		return WebP
	case buffer.isHEIC():
		return HEIC
	case buffer.isHEIF():
		return HEIF
	default:
		return UnKnow
	}
}

func getBuffer(image []byte) (imageType, error) {
	if len(image) < 1 {
		return nil, errors.New("error size")
	}
	max := len(image)
	if max > 50 {
		max = 50
	}
	buffer := make(imageType, max)
	copy(buffer, image[:max])
	return buffer, nil
}

func (i imageType) isJPEG() bool {
	return bytes.Equal([]byte{0xFF}, i[:1])
}

func (i imageType) isPNG() bool {
	return bytes.Equal([]byte{0x89}, i[:1])
}

func (i imageType) isGIF() bool {
	return bytes.Equal([]byte{0x47}, i[:1])
}

func (i imageType) isTIFF() bool {
	return bytes.Equal([]byte{0x49}, i[:1]) || bytes.Equal([]byte{0x4D}, i[:1])
}

func (i imageType) isWebP() bool {
	if bytes.Equal([]byte{0x52}, i[:1]) && len(i) > 12 {
		str := string(i[:12])
		return strings.HasPrefix(str, "RIFF") || strings.HasSuffix(str, "WEBP")
	}
	return false
}

func (i imageType) isHEIC() bool {
	if bytes.Equal([]byte{0x00}, i[:1]) && len(i) > 12 {
		str := string(i[8:12])
		_, hasHEICBitMaps := HEICBitMaps[str]
		return hasHEICBitMaps
	}
	return false
}

func (i imageType) isHEIF() bool {
	if bytes.Equal([]byte{0x00}, i[:1]) && len(i) > 12 {
		str := string(i[8:12])
		_, hasHEIFBitMaps := HEIFBitMaps[str]
		return hasHEIFBitMaps
	}
	return false
}
