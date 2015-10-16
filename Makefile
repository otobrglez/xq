.PHONY: install-deps

all: xq

install-deps:
	go get -u github.com/moovweb/gokogiri

xq:
	go build -v .

test: xq
	cat ./data/wiki_page.html | ./xq //title | grep Wikipedia
	cat ./data/wiki_page.html | ./xq "//table[@class='multicol'][7]//a/@href" | grep Go
	cat ./data/weather.xml | ./xq "//temp_c" | grep "16"

clean:
	rm -rf ./xq
