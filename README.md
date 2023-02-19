# MJGETTER 💜
mjgetter are a type of tools that are used for information gathering and reconnaissance activities.

## Installation
```
  go get github.com/mjoffsec/mjgetter
```

## Usage
```
  Usage:
  -as string         ASN Lookup (default "unknown")
  -dl                DNS Lookup
  -ep                Extract Page Link
  -fd                Find DNS Host Records
  -fs string         Find Shared DNS Servers e.g: ns1.example.com (default "unknown")
  -gl string         IP Geolocation Lookup (default "unknown")
  -ht                HTTP Headers
  -ir                IP Reverse Lookup
  -rd                Reverse DNS
  -rs                Reverse Analytics Search
  -url string        URL (default "unknown")
  -zt                Zone Transfer Online Test
  
  Example:
  mjgetter -url example.com -dl -as -fd
```

