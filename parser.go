package parser

import (
	"index/suffixarray"
	"strings"
)

type Parser struct {
	sa *suffixarray.Index
}

func NewDomainParser() Parser {
	sa := CreateTLDIndex()
	return Parser{
		sa: sa,
	}
}

func (p *Parser) FindTldOffset(domain_parts []string) int {
	counter := 2
	for counter > 0 {
		start_point := len(domain_parts) - counter
		if start_point < 0 {
			return 0
			break
		}
		tld_parts := strings.Join(domain_parts[len(domain_parts)-counter:], ".")

		indicies := p.sa.Lookup([]byte(tld_parts), -1)
		if len(indicies) > 0 {
			offset := (len(domain_parts) - (counter +1))
			if offset >= 0 {
				return offset
				break
			}
		}
		counter--
	}

	return 0

}

func (p *Parser) GetDomain(domain string) string{
	domain_parts := strings.Split(domain, ".")
	offset := p.FindTldOffset(domain_parts)
	return domain_parts[offset]

}

func (p *Parser) GetSubdomain(domain string) string {
	domain_parts := strings.Split(domain, ".")
	offset := p.FindTldOffset(domain_parts)
	return strings.Join(domain_parts[:offset], ".")
}

func (p *Parser) GetFQDN(domain string) string {
	domain_parts := strings.Split(domain, ".")
	offset := p.FindTldOffset(domain_parts)
	return strings.Join(domain_parts[offset:], ".")
}

func (p *Parser) GetTld(domain string) string {
	domain_parts := strings.Split(domain, ".")
	offset := p.FindTldOffset(domain_parts)
	return strings.Join(domain_parts[offset + 1:], ".")
}
