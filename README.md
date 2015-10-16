# xq

[xq] is a lightweight command-line XPATH processor (for HTML)

- [Oto Brglez](https://github.com/otobrglez)

## Usage

Extracting simple content from HTML with XPATH:

    curl -s https://news.ycombinator.com/ | ./xq "//title"
    #=> Hacker News

Extracting content from XML with XPATH:
    
    curl -s http://w1.weather.gov/xml/current_obs/KBOS.xml | ./xq //current_observation/temp_f
    # => 59.0

A bit more sophisticated XPATH from standard input: 

    cat ./data/wiki_page.html | ./xq "//table[@class='multicol'][7]//a/@href"
    #=>
    #...
    #/wiki/Go_(programming_language)
    #...

Extracting from file
 
    ./xq "//table[@class='multicol'][7]//a/@href" ./data/wiki_page.html

## License & Author

[Oto Brglez](https://github.com) - MIT License.

[xq]:https://github.com/otobrglez/xq