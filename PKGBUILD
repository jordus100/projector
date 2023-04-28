pkgname='proj'
pkgrel=1
pkgver=1.0.0
arch=('any')
makedepends=('go' 'git')
source=('git+https://github.com/jordus100/laplace-compiled')

_scripts_path="/usr/local/sbin"
_static_files_path="/usr/local/share"

build() {
	cd "$srcdir"
	cd "laplace-compiled"
	go build -o laplace main.go 
}

package() {
	cd "$srcdir"
	cp "proj-klient" "proj-start" "laplace" $_scripts_path
	chmod a+x "$_scipts_path/proj-klient" "$_scripts_path/proj-start" "$_scripts_path/laplace"
	_pkg_files="$_static_files_path/$pkgname"
	mkdir $_pkg_files
	cp -r "files" $_pkg_files
	chmod -R a+rw $_pkg_files
}
