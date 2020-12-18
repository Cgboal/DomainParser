package parser

import (
	"index/suffixarray"
	"strings"
)

func CreateTLDIndex(tlds []string) *suffixarray.Index {	
	data := []byte("\x00" + strings.Join(tlds, "\x00") + "\x00")
	sa := suffixarray.New(data)
	return sa
}

