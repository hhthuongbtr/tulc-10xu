build:
	GOOS=linux GOARCH=amd64 go build -o deployment/talarm main.go
	docker build -t talarm deployment
clean:
	rm -f deployment/talarm
run:
	docker run -p 80:8000 -itd -e GIN_MODE=release talarm
