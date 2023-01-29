package gva

import (
	"log"
	"regexp"
)

type Profile struct {
	Name       string  `json:"name"`
	Headers    Headers `json:"headers"`
	URL        string  `json:"url"`
	Port       string  `json:"port"`
	APIVersion string  `json:"api_version"`
}

func (p *Profile) UpdateVersion(nv string, x string)  {
	regexp, err := regexp.Compile(`v[0-9]`)
	
	if err != nil {
		log.Fatal(err)
	}

	match := regexp.ReplaceAllString(p.URL, nv)

	if match != "" {
		p.URL = match
	}

	p.APIVersion = nv
	p.Headers.XAPIVersion = x

}

func (p *Profile) UpdatePort(port string) {
	regexp, err := regexp.Compile(`[0-9]{2,}`)

	if err != nil {
		log.Fatal(err)
	}

	match := regexp.ReplaceAllString(p.URL, port)

	if match != "" {
		p.URL = match
	}

	p.Port = port
}

type Headers struct {
	Accept      string `json:"accept"`
	ContentType string `json:"Content-type"`
	XAPIVersion string `json:"x-api-version"`
}