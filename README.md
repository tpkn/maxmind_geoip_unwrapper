## MaxMind GeoIP Unwrapper

This tool unwraps ip ranges into a convinient format.

## Usage

```
mgu < "GeoLite2-City-Blocks-IPv4.csv" > "GeoLite2-City-Blocks-IPv4-Wnrapped.csv"
```

## Why?

MaxMind compresses IP address ranges with the same geo-location using a subnet. For example, one of the ip ranges for
Phuket looks like

```
ip,geo_id,lat,lon
1.0.164.24/29,1151254,7.9833,98.3662
```

istead of this

```
ip,geo_id,lat,lon
1.0.164.24,1151254,7.9833,98.3662
1.0.164.25,1151254,7.9833,98.3662
...
1.0.164.29,1151254,7.9833,98.3662
1.0.164.30,1151254,7.9833,98.3662
```

## Pros and cons

Faster lookup and works in all DB, but takes up many times more disk space.

