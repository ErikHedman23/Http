package main

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	pU := new(ParsedURL)
	password, passwordSet := parsedUrl.User.Password()
	if passwordSet == true {
		pU.password = password
	} else {
		pU.password = ""
	}

	pU.username = parsedUrl.User.Username()
	pU.protocol = parsedUrl.Scheme
	pU.hostname = parsedUrl.Hostname()
	pU.port = parsedUrl.Port()
	pU.pathname = parsedUrl.Path
	pU.search = parsedUrl.RawQuery
	pU.hash = parsedUrl.Fragment

	return *pU
}
