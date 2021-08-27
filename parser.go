package parser

import (
	"index/suffixarray"
	"strings"
	"net/http"
	"io/ioutil"
	"bytes"
)

type Parser struct {
	sa *suffixarray.Index
}

type Domain struct {
	Subdomain string
	Domain string
	TLD string
}

func NewDomainParser() Parser {
	data, err := ioutil.ReadFile("/tmp/.tlds")
	if err != nil {
		data, _ = download()
		ioutil.WriteFile("/tmp/.tlds", data, 0644)
	}

	tlds := strings.Split(string(data), "\n")

	sa := CreateTLDIndex(tlds)
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
		}
		tld_parts := strings.Join(domain_parts[len(domain_parts)-counter:], ".")

		indicies := p.sa.Lookup([]byte(tld_parts), -1)
		if len(indicies) > 0 {
			offset := (len(domain_parts) - (counter +1))
			if offset >= 0 {
				return offset
			}
		}
		counter--
	}

	return 0

}

func (p *Parser) ParseDomain(domain string) Domain {
	domain_parts := strings.Split(domain, ".")
	offset := p.FindTldOffset(domain_parts)
	return Domain{
		Subdomain: strings.Join(domain_parts[:offset], "."),
		Domain: domain_parts[offset],
		TLD: strings.Join(domain_parts[offset + 1:], "."),
	}
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

func download() ([]byte, error) {
	u := "https://publicsuffix.org/list/public_suffix_list.dat"
	resp, err := http.Get(u)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	lines := strings.Split(string(body), "\n")
	var buffer bytes.Buffer

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "// ===BEGIN PRIVATE DOMAINS") {
			break
		}
		if line != "" && !strings.HasPrefix(line, "//") {
			buffer.WriteString(line)
			buffer.WriteString("\n")
		}
	}

	return buffer.Bytes(), nil
}
