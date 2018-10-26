The aim of my-bloody-hetzner-sb-notifier is a simple CLI to fetch the current Hetzner Serverbörse deals and filter them according to CLI parameters sorted by score.

The score is calculated from the amount of HDD space as well as RAM and CPU-Benchnmark for better comparability. 

The CLI interface looks like this:
````
Usage of hetzner-sb-notifier:
  -alert-on-score int
    	set alert on score
  -max-benchmark int
    	set max benchmark (default 20000)
  -max-hdd-count int
    	set max hdd count (default 15)
  -max-hdd-size int
    	set max hdd size (default 6144)
  -max-price float
    	set max price (default 297)
  -max-ram int
    	set max ram (default 256)
  -min-benchmark int
    	set min benchmark
  -min-hdd-count int
    	set min hdd count
  -min-hdd-size int
    	set min hdd size
  -min-price float
    	set min price
  -min-ram int
    	set min ram
````       

## Example

./hetzner-sb-notifier --max-price 77 --min-ram 128 --min-hdd-count 2 --min-hdd-size 4096
```` 
Got 545 offers. Filtered offers: 3
           ID|     Ram|             HDD|                           CPU|    Price|  Score|  Reduce time|Specials
  SB64-935022|  128 GB|  2x 2 TB (4096)|  Intel Xeon E5-1650V2 (12518)|  64.00 €|  91.84|      47h 48m|ECC, Ent. HDD, iNIC
  SB72-927788|  128 GB|  2x 2 TB (4096)|  Intel Xeon E5-1650V3 (13335)|  72.00 €|  86.17|      21h 08m|ECC, Ent. HDD, iNIC
  SB73-910394|  128 GB|  3x 2 TB (6144)|  Intel Xeon E5-1650V2 (12518)|  73.00 €|  86.13|      03h 04m|ECC, Ent. HDD, iNIC
```` 

## Build

The Go project uses Go Modules and can be easily build with the wrapper script build.sh:
```` 
chmod +x build.sh
./build.sh
```` 