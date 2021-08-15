MY_VAR=$(shell go run main.go --version)

version:	
	@echo $(MY_VAR)
clean:
	rm -fv superserver~
build:
	make clean
	env GOOS="linux" GOARCH="amd64" go build -o "superserver" -ldflags="-s -w" main.go	
docker:
	docker build -t ginui-server:$(MY_VAR) .
docker-test:
	docker run -it --rm -p 1986:1986 --name factory-test ginui-server:$(MY_VAR)
	docker logs -f factory-test