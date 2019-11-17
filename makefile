main:
	go build -v --race -o build/plan .

clean:
	rm build -r
