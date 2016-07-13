BINS=basic-auth-proxy

$(BINS): *.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo .

clean:
	rm -rf $(BINS)

all: clean $(BINS)
