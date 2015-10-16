.PHONY: install-deps

all: xq

install-deps:
	# go get -u github.com/moovweb/gokogiri

xq: install-deps
	go build -v .

test: xq
	cat ./data/wiki_page.html | ./xq //title

clean:
	rm -rf ./xq
