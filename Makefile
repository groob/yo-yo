PKGNAME=yo-yo
PKGVERSION=0.0.2
PKGID=com.github.groob.yo-yo

all: build pkg
build:
	go build -o pkg/pkgroot/usr/local/bin/yo-yo
pkg: build
	mkdir -p out
	pkgbuild --root pkg/pkgroot --identifier ${PKGID} --version ${PKGVERSION} --scripts pkg/scripts out/${PKGNAME}-${PKGVERSION}.pkg

	
