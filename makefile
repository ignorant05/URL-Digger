build: 
	go build  -o udig cmd/app/main.go 

run: 
	go build  -o udig cmd/app/main.go 
	./udig

clean: 
	rm udig 

debug: 
	dlv debug cmd/app/main.go



