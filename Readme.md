The aim of my-bloody-hetzner-sb-notifier is a simple CLI to fetch the current Hetzner Serverbörse deals and filter them according to CLI parameters sorted by score.

For each offer a score is calculated from the amount of HDD space as well as RAM and CPU-Benchnmark for better comparability. 

The CLI interface looks like this:
````
Usage of hetzner-sb-notifier:
  -min-cpu-benchmark int
    	set min benchmark
  -max-cpu-benchmark int
    	set max benchmark (default 20000)
  -max-hdd-count int
    	set max hdd count (default 15)
  -max-hdd-size int
    	set max hdd size (default 6144)
  -max-price float
    	set max price (default 297)
  -max-ram int
    	set max ram (default 256)
  -min-hdd-count int
    	set min hdd count
  -min-hdd-size int
    	set min hdd size
  -min-price float
    	set min price
  -min-ram int
    	set min ram
  -serve-http
    	set serve http
  -serve-http-port int
    	set serve http port (default 8080)
  -output
    	set output: one of table, json (default table)
````       

## Http mode

In HTTP mode (app started with flag -serve-http) hetzner-sb-notifier runs continuously and can be queried via simple HTTP GET request.
CLI parameters are translated to camel case. For example  min-hdd-count becomes minHddCount

## CLI-Example

./hetzner-sb-notifier --max-price 77 --min-ram 128 --min-hdd-count 2 --min-hdd-size 4096
```` 
Got 545 offers. Filtered offers: 3
           ID|     Ram|             HDD|                           CPU|    Price|  Score|  Reduce time|Specials
  SB64-935022|  128 GB|  2x 2 TB (4096)|  Intel Xeon E5-1650V2 (12518)|  64.00 €|  91.84|      47h 48m|ECC, Ent. HDD, iNIC
  SB72-927788|  128 GB|  2x 2 TB (4096)|  Intel Xeon E5-1650V3 (13335)|  72.00 €|  86.17|      21h 08m|ECC, Ent. HDD, iNIC
  SB73-910394|  128 GB|  3x 2 TB (6144)|  Intel Xeon E5-1650V2 (12518)|  73.00 €|  86.13|      03h 04m|ECC, Ent. HDD, iNIC
```` 

## HTTP-Example

./hetzner-sb-notifier -serve-http
```` 
Running http server on address :8080

Got request: GET /?minRam=256&output=table HTTP/1.1
Host: localhost:8080
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8
Accept-Encoding: gzip, deflate, br
Accept-Language: de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7
Cache-Control: max-age=0
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36
```` 

## Build

The Go project uses Go Modules and can be easily build with the wrapper script build.sh:
```` 
chmod +x build.sh
./build.sh
```` 