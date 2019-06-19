# Face !
face package implements face detection for Go using dlib, a popular machine learning toolkit. This repo is a fork of [go-face package](https://github.com/Kagami/go-face) written by Kagomi.

### Requirements
According to `go-face` instructions,
To compile go-face you need to have dlib (>= 19.10) and libjpeg development packages installed. It's highly recommended to compile dlib in your machine.

#### Ubuntu 16.04, Ubuntu 18.04
You may use dlib PPA which contains latest dlib package compiled with Intel MKL support:

```
sudo add-apt-repository ppa:kagamih/dlib
sudo apt-get update
sudo apt-get install libdlib-dev libjpeg-turbo8-dev
```

#### Ubuntu 18.10+, Debian sid
Latest versions of Ubuntu and Debian provide suitable dlib package so just run:

```
# Ubuntu
sudo apt-get install libdlib-dev libopenblas-dev libjpeg-turbo8-dev
# Debian
sudo apt-get install libdlib-dev libopenblas-dev libjpeg62-turbo-dev
```

#### ONLY FOR UBUNTU 18.10+ AND DEBIAN SID:
It won't install pkgconfig metadata file so create one in /usr/local/lib/pkgconfig/dlib-1.pc with the following content:

```
libdir=/usr/lib/x86_64-linux-gnu
includedir=/usr/include

Name: dlib
Description: Numerical and networking C++ library
Version: 19.10.0
Libs: -L${libdir} -ldlib -lblas -llapack
Cflags: -I${includedir}
Requires:
```

#### MacOS
Make sure you have Homebrew installed.

brew install pkg-config dlib
sed -i '' 's/^Libs: .*/& -lblas -llapack/' /usr/local/lib/pkgconfig/dlib-1.pc

#### Windows
Make sure you have MSYS2 installed.

- Run MSYS2 MSYS shell from Start menu
- Run pacman -Syu and if it asks you to close the shell do that
- Run pacman -Syu again
- Run pacman -S mingw-w64-x86_64-gcc mingw-w64-x86_64-dlib mingw-w64-x86_64-pkg-config
- * If you already have Go and Git installed and available in PATH uncomment set MSYS2_PATH_TYPE=inherit line in msys2_shell.cmd located in MSYS2 installation folder
- * Otherwise run pacman -S mingw-w64-x86_64-go git
- Run MSYS2 MinGW 64-bit shell from Start menu to compile and use go-face

#### Other Systems
Try to install dlib/libjpeg with package manager of your distribution or compile from sources. Note that go-face won't work with old packages of dlib such as libdlib18. Alternatively create issue with the name of your system and someone might help you with the installation process.

