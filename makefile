main:
	go build -v --race -o build/main .

install:
	mkdir /usr/local/plan
	cp build/main /usr/local/plan
	touch /usr/local/plan/data.json
	echo "{}" > /usr/local/plan/data.json
	chmod 755 /usr/local/plan/main
	ln -s /usr/local/plan/main /usr/bin/plan 
	chmod 755 /usr/bin/plan

uninstall:
	rm -r /usr/local/plan
	rm /usr/local/bin/plan

clean:
	rm build -r

