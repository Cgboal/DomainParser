package tests

import (
	"testing"
	"github.com/Cgboal/DomainParser"
	"strings"
)


func TestGetSubdomains(t *testing.T) {
	parser := parser.NewDomainParser()
	testCase := "1.2.3.4.testing.co.uk"
	subdomain := parser.GetSubdomain(testCase)

	if subdomain != "1.2.3.4" {
		t.Errorf("Subdomain Incorrect, got %s, want %s", subdomain, "1.2.3.4" )
	}
}

func TestGetDomain(t *testing.T) {
	parser := parser.NewDomainParser()
	testCase := "1.2.3.4.testing.co.uk"
	domain := parser.GetDomain(testCase)

	if domain != "testing" {
		t.Errorf("Domain Incorrect, got %s, want %s", domain, "testing")
	}
}

func TestGetFQDN(t *testing.T) {
	parser := parser.NewDomainParser()
	testCase := "1.2.3.4.testing.co.uk"
	fqdn := parser.GetFQDN(testCase)

	if fqdn != "testing.co.uk" {
		t.Errorf("FQDN Incorrect, got %s, want %s", fqdn, "testing")
	}
}

func benchmarkParsing(i int, b *testing.B) {
	parser := parser.NewDomainParser()
	testCase := "1.2.3.4.testing.co.uk"
	domain_parts := strings.Split(testCase, ".")
	for n := 0; n < b.N; n++ {
		for x := 0; x < i; x++ {
			parser.FindTldOffset(domain_parts)
		}
	}
}

func BenchmarkParsing10(b *testing.B) { benchmarkParsing(10, b) }
func BenchmarkParsing100(b *testing.B) { benchmarkParsing(100, b) }
func BenchmarkParsing1000(b *testing.B) { benchmarkParsing(1000, b) }
func BenchmarkParsing100000(b *testing.B) { benchmarkParsing(100000, b) }
func BenchmarkParsing1000000(b *testing.B) { benchmarkParsing(1000000, b) }
