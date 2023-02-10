include config.mk
PKGNAME=yo-yo
PKGVERSION=0.0.4
PKGID=com.github.groob.yo-yo

all: build pkg
build:
	/bin/rm -rf build
	GOOS=darwin GOARCH=amd64 go build -o build/yo-yo.amd64
	GOOS=darwin GOARCH=arm64 go build -o build/yo-yo.arm64
	/usr/bin/lipo -create -output pkg/pkgroot/usr/local/bin/yo-yo build/yo-yo.amd64 build/yo-yo.arm64
	@sudo codesign --timestamp --force --deep -s "${DEV_APP_CERT}" pkg/pkgroot/usr/local/bin/yo-yo
pkg: build
	mkdir -p out
	pkgbuild --root pkg/pkgroot --identifier ${PKGID} --version ${PKGVERSION} --scripts pkg/scripts out/${PKGNAME}-${PKGVERSION}.pkg --sign "${DEV_INSTALL_CERT}"
