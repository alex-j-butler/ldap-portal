build_linux_386:
	GOPATH=`pwd` go get -v qixalite.com/Ranndom/ldap-portal
	GOPATH=`pwd` GOOS=linux GOARCH=386 go build -o ldap-portal-i386 -v src/qixalite.com/Ranndom/ldap-portal/main.go
	
build_linux_amd64:
	GOPATH=`pwd` go get -v qixalite.com/Ranndom/ldap-portal
	GOPATH=`pwd` GOOS=linux GOARCH=amd64 go build -o ldap-portal-amd64 -v src/qixalite.com/Ranndom/ldap-portal/main.go

build_windows_386:
	GOPATH=`pwd` go get -v qixalite.com/Ranndom/ldap-portal
	GOPATH=`pwd` GOOS=windows GOARCH=386 go build -o ldap-portal-i386.exe -v src/qixalite.com/Ranndom/ldap-portal/main.go
	
build_windows_amd64:
	GOPATH=`pwd` go get -v qixalite.com/Ranndom/ldap-portal
	GOPATH=`pwd` GOOS=windows GOARCH=amd64 go build -o ldap-portal-amd64.exe -v src/qixalite.com/Ranndom/ldap-portal/main.go

build:
	GOPATH=`pwd` go get -v qixalite.com/Ranndom/ldap-portal
	GOPATH=`pwd` go build -o ldap-portal-default -v src/qixalite.com/Ranndom/ldap-portal/main.go