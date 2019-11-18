main:
	go build -v --race -o build/main .

install:
	mkdir /usr/local/plan
	cp build/main /usr/local/plan
	touch /usr/local/plan/data.json
	echo "{}" > /usr/local/plan/data.json
	chmod 755 /usr/local/plan/main
	chmod 766 /usr/local/plan/data.json
	ln -s /usr/local/plan/main /usr/bin/plan 
	chmod 755 /usr/bin/plan

uninstall:
	rm -r /usr/local/plan
	rm /usr/bin/plan

clean:
	rm build -r

