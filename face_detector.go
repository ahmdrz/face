package face

// #cgo pkg-config: dlib-1
// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -ljpeg
// #include <stdlib.h>
// #include <stdint.h>
// #include "face_detector.h"
import "C"
import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"unsafe"
)

const (
	rectLen = 4
)

var (
	// EmptyImage error of empty image
	EmptyImage = errors.New("empty image")
)

// Detector creates face descriptors for provided images and
// classifies them into categories.
type Detector struct {
	ptr *C.facedetector
}

// NewDetector returns a new detector interface.
func NewDetector() (d *Detector, err error) {
	ptr := C.facedetector_init()

	if ptr.err_str != nil {
		err = fmt.Errorf("error: %s code: %d", C.GoString(ptr.err_str), int(ptr.err_code))
		C.facedetector_free(ptr)
		C.free(unsafe.Pointer(ptr.err_str))
		return
	}

	d = &Detector{ptr}
	return
}

func (d *Detector) detect(imgData []byte) (faces []image.Rectangle, err error) {
	if len(imgData) == 0 {
		err = EmptyImage
		return
	}
	cImgData := (*C.uint8_t)(&imgData[0])
	ret := C.facedetector_detect(d.ptr, cImgData, C.int(len(imgData)))
	defer C.free(unsafe.Pointer(ret))

	if ret.err_str != nil {
		defer C.free(unsafe.Pointer(ret.err_str))
		err = fmt.Errorf("error: %s code: %d", C.GoString(d.ptr.err_str), int(d.ptr.err_code))
		return
	}

	length := int(ret.length)
	faces = make([]image.Rectangle, length)
	if length == 0 {
		return
	}

	defer C.free(unsafe.Pointer(ret.rectangles))

	rDataLen := length * rectLen
	rDataPtr := unsafe.Pointer(ret.rectangles)
	rData := (*[1 << 30]C.long)(rDataPtr)[:rDataLen:rDataLen]

	for i := 0; i < length; i++ {
		x0 := int(rData[i*rectLen])
		y0 := int(rData[i*rectLen+1])
		x1 := int(rData[i*rectLen+2])
		y1 := int(rData[i*rectLen+3])
		faces[i] = image.Rect(x0, y0, x1, y1)
	}
	return
}

// Detect returns all faces found on the provided image.
// Empty list is returned if there are no faces, error is
// returned if there was some error while decoding/processing image.
// Only JPEG format is currently supported. Others will convert
// to JPEG format. Thread-safe.
func (d *Detector) Detect(img image.Image) (faces []image.Rectangle, err error) {
	faces = []image.Rectangle{}
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return
	}
	return d.DetectBytes(buf.Bytes())
}

// DetectBytes returns all faces found on the provided image. Input is byte-array.
// Pass image.Image object to Detect.
func (d *Detector) DetectBytes(img []byte) (faces []image.Rectangle, err error) {
	return d.detect(img)
}

// Close frees resources taken by the Detector. Safe to call multiple
// times.
func (d *Detector) Close() {
	C.facedetector_free(d.ptr)
	d.ptr = nil
}
