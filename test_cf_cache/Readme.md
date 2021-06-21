## CloudFlare Cache Tester
#### This tool scrapes a URL to gather the local static assets files such as images, CSS and JS. And returns the relevant headers and the status code

### Usage
```
Usage of ./cfcache_darwin.bin:
  -asset string
        The type of the assets (default "all")
        Choices: ['css', 'js', 'images', 'all']
  -url string
        The site URL
```

### Examples
```
./cfcache_darwin.bin --url=https://site.tld
./cfcache_darwin.bin --url=https://site.tld --asset=css
./cfcache_darwin.bin --url=https://site.tld --asset=js
./cfcache_darwin.bin --url=https://site.tld --asset=images
```