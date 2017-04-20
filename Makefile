########################################################################################

# This Makefile generated by GoMakeGen 0.5.0 using next command:
# gomakegen .

########################################################################################

.PHONY = fmt all clean deps

########################################################################################

all: swptop

swptop:
	go build swptop.go

deps:
	git config --global http.https://pkg.re.followRedirects true
	go get -d -v pkg.re/essentialkaos/ek.v8

fmt:
	find . -name "*.go" -exec gofmt -s -w {} \;

clean:
	rm -f swptop

########################################################################################
