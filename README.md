# xq

[xq] is a lightweight command-line XPath processor for HTML and XML.

[![Build Status](https://travis-ci.org/otobrglez/xq.svg?branch=master)](https://travis-ci.org/otobrglez/xq)

## Usage

Extracting simple content from HTML with XPath:

    curl -s https://news.ycombinator.com/ | ./xq "//title"
    #=> Hacker News

Extracting content from XML with XPath:
    
    curl -s http://w1.weather.gov/xml/current_obs/KBOS.xml | ./xq //current_observation/temp_f
    # => 59.0

A bit more sophisticated XPath from standard input: 

    cat ./data/wiki_page.html | ./xq "//table[@class='multicol'][7]//a/@href"
    #=>
    #...
    #/wiki/Go_(programming_language)
    #...

Extracting from local file,... Trivial. :)
 
    ./xq "//table[@class='multicol'][7]//a/@href" ./data/wiki_page.html

## License & Author

[Oto Brglez](https://github.com) - MIT License.

[xq]:https://github.com/otobrglez/xq
