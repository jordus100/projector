# skrypt pakujący (makepkg) program laplace (projektor)
# parviaij 2023
pkgname='projector'
pkgrel=1
pkgver=1.0.0
arch=('x86_64')
makedepends=('go' 'git')
source=('git+https://github.com/jordus100/projector')
sha256sums=('SKIP')

_static_files="usr/share/proj"
_exe_files="usr/bin"
_scripts="usr/local/bin"
_service_files="usr/lib/systemd/system"
_src="$pkgname/src"

build() {
	cd "$srcdir/$_src"
	go build -o laplace -ldflags "-X main.staticDir=/$_static_files/files" main.go 
}

package() {
	cd "$srcdir/$_src"
	install -D "laplace" "$pkgdir/$_exe_files/laplace"
	cd "$srcdir/$pkgname"
	install -D "proj-start" "$pkgdir/$_scripts/proj-start"
	install -D "proj-klient" "$pkgdir/$_scripts/proj-klient"
	install -D -m="u=rw g=r o=r" "projector.service" "$pkgdir/$_service_files/projector.service"
	cd "$srcdir/$pkgname/files"
	install -dD "files" "$pkgdir/$_static_files"
	cp -r . "$pkgdir/$_static_files/files"
}
