package smileyface

import (
	"io/ioutil"
	"testing"
)

func TestFaceDetector(t *testing.T) {
	bytes, err := ioutil.ReadFile("samples/albert-einstein.jpg")
	if err != nil {
		t.Fatal(err)
		return
	}
	d, err := NewDetector()
	if err != nil {
		t.Fatal(err)
		return
	}
	faces, err := d.Detect(bytes)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(faces)
}
