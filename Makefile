build:
	GOOS=linux GOARCH=amd64 go build -o deployment/tulc-10xu main.go
	docker build -t tulc-10xu deployment
clean:
	rm -f deployment/tulc-10xu
run:
	docker run -p 80:8000 -itd -e GIN_MODE=release tulc-10xu
