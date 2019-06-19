# Face !
face package implements face detection for Go using dlib, a popular machine learning toolkit. This repo is a fork of [go-face package](https://github.com/Kagami/go-face) written by Kagomi.

### Requirements
According to `go-face` instructions,
To compile go-face you need to have dlib (>= 19.10) and libjpeg development packages installed. It's highly recommended to compile dlib in your machine.

To install in Ubuntu, Debian, Windows and MacOS take a look at [Installation.md](https://github.com/ahmdrz/face/blob/master/Installation.md) file.

### Test
To fetch test data and run tests:

```
go test -v ./...
```