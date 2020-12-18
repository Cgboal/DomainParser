##### Doamin Parser

A super fast DomainParser written in Go. 

This library allows you to very quickly seperate domain names into subdomains, domains, tlds, and fqdns.

Here are some benchmarks which demonstrate the speed: 
```
→ go test -bench=.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            [4af0b93]
goos: linux
goarch: amd64
pkg: github.com/Cgboal/DomainParser/tests
BenchmarkParsing10-12         	 299260	     3995 ns/op
BenchmarkParsing100-12        	  29043	    39320 ns/op
BenchmarkParsing1000-12       	   2955	   406441 ns/op
BenchmarkParsing100000-12     	     28	 40247358 ns/op
BenchmarkParsing1000000-12    	      3	399338124 ns/op
PASS
ok  	github.com/Cgboal/DomainParser/tests	10.446s
```

As can be seen above, this parser is capable of processing 10 million DNS names in 0.399 seconds.

Example usage can be seen in the tests directory. 
