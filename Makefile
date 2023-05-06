check_install:
	which swagger || GO111MODULE=off go install -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	GOROOT=~/../../opt/homebrew/Cellar/go/1.19.5/libexec swagger generate spec -o ./swagger.yaml --scan-models