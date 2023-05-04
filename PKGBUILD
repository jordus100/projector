pkgname='proj'
pkgrel=1
pkgver=1.0.0
arch=('any')
makedepends=('go' 'git')
source=('git+https://github.com/jordus100/laplace-compiled')

_scripts_path="/usr/local/sbin"
_static_files_path="/usr/local/share"
_src_path="laplace-compiled/laplace/src"
_root_path="laplace-compiled/"

build() {
	cd "$srcdir/$_src_path"
	go build -o laplace main.go 
}

package() {
	cd "$srcdir/$_src_path"
	cp "laplace" $_scripts_path
	cd "$srcdir/$_root_path"
	cp "proj-start" "proj-klient" $_scripts_path
	chmod a+x "$_scipts_path/proj-klient" "$_scripts_path/proj-start" "$_scripts_path/laplace"
	_pkg_files="$_static_files_path/$pkgname"
	mkdir $_pkg_files
	cp -r "files" $_pkg_files
	chmod -R a+rw $_pkg_files
}
