# skrypt pakujący (makepkg) program laplace (projektor)
# parviaij 2023
pkgname='proj'
pkgrel=1
pkgver=1.0.0
arch=('any')
makedepends=('go' 'git')
source=('git+https://github.com/jordus100/laplace-compiled')
sha256sums=('SKIP')

_static_files="usr/share/proj"
_exe_files="usr/bin"
_src="laplace-compiled/src/laplace"
_repo="laplace-compiled/"

build() {
	cd "$srcdir/$_src"
	go build -o laplace -ldflags "-X main.staticDir=/$_static_files/files" main.go 
}

package() {
	cd "$srcdir/$_src"
	install -D "laplace" "$pkgdir/$_exe_files/laplace"
	cd "$srcdir/$_repo/files"
	install -dD "files" "$pkgdir/$_static_files"
	cp -r . "$pkgdir/$_static_files/files"
}
