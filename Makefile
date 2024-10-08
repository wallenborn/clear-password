
all: clearpassword.exe


clean:
	rm -f clearpassword.exe
	rm -f *~
	rm -f .*~

distclean: clean

clearpassword.exe: main.go
	go build

clearpassword: main.go
	env GOOS=linux GOARCH=amd64 go build 

test: 
	cd info; go test
