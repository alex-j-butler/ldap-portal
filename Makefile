.PHONY: bindata bindata_debug build build_linux_386 build_linux_amd64 build_windows_386 build_windows_amd64 all

bindata:
	GOPATH=`pwd` go get -u github.com/jteeuwen/go-bindata/...
	./bin/go-bindata -pkg views -prefix views -o src/qixalite.com/Ranndom/ldap-portal/views/views.go views/...
	./bin/go-bindata -pkg public -prefix public -o src/qixalite.com/Ranndom/ldap-portal/public/public.go public/...
	
bindata_debug:
	GOPATH=`pwd` go get -u github.com/jteeuwen/go-bindata/...
	./bin/go-bindata -debug -pkg views -prefix views -o src/qixalite.com/Ranndom/ldap-portal/views/views.go views/...
	./bin/go-bindata -debug -pkg public -prefix public -o src/qixalite.com/Ranndom/ldap-portal/public/public.go public/...

build_linux_386: bindata
	GOPATH=`pwd` go get -v qixalite.com/Ranndom/ldap-portal
	GOPATH=`pwd` GOOS=linux GOARCH=386 go build -o ldap-portal-i386 -v src/qixalite.com/Ranndom/ldap-portal/main.go
	
build_linux_amd64: bindata	
	GOPATH=`pwd` go get -v qixalite.com/Ranndom/ldap-portal
	GOPATH=`pwd` GOOS=linux GOARCH=amd64 go build -o ldap-portal-amd64 -v src/qixalite.com/Ranndom/ldap-portal/main.go

build_windows_386: bindata
	GOPATH=`pwd` go get -v qixalite.com/Ranndom/ldap-portal
	GOPATH=`pwd` GOOS=windows GOARCH=386 go build -o ldap-portal-i386.exe -v src/qixalite.com/Ranndom/ldap-portal/main.go
	
build_windows_amd64: bindata	
	GOPATH=`pwd` go get -v qixalite.com/Ranndom/ldap-portal
	GOPATH=`pwd` GOOS=windows GOARCH=amd64 go build -o ldap-portal-amd64.exe -v src/qixalite.com/Ranndom/ldap-portal/main.go

build: bindata
	GOPATH=`pwd` go get -v qixalite.com/Ranndom/ldap-portal
	GOPATH=`pwd` go build -o ldap-portal-default -v src/qixalite.com/Ranndom/ldap-portal/main.go

run: bindata_debug	
	GOPATH=`pwd` go get -v qixalite.com/Ranndom/ldap-portal
	GOPATH=`pwd` go run src/qixalite.com/Ranndom/ldap-portal/main.go

all: build_linux_386 build_linux_amd64 build_windows_386 build_windows_amd64
