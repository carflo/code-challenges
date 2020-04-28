default:
	go build -o main

profile:
	./xtime.sh go test -timeout 30s github.com/carflo/code-challgenges/phrases -run TestPhrases -cpuprofile cpuprofile
