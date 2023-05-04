# skrypt pakujący (makepkg) program laplace (projektor)
# parviaij 2023
pkgname='proj'
pkgrel=1
pkgver=1.0.0
arch=('any')
makedepends=('go' 'git')
source=('git+https://github.com/jordus100/laplace-compiled')

_static_files="/usr/share/proj"
_exe_files="/usr/bin"
_src="laplace-compiled/laplace/src"
_repo="laplace-compiled/"

build() {
	cd "$srcdir/$_src"
	go build -o laplace -ldflags "-X main.staticDir=$_static_files" main.go 
}

check() {
	cksums=(SKIP)
}

package() {
	cd "$srcdir/$_src"
	cp "laplace" $pkgdir/$_exe_files
	cd "$srcdir/$_repo"
	cp -r $src/files $pkgdir/$_static_files 
}
