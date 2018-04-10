package main

import (
	"net/http"
	"log"
	"encoding/xml"
	"fmt"
)

type AtomFeed struct {
	XMLName  xml.Name `xml:"feed"`
	Xmlns    string   `xml:"xmlns,attr,omitempty"`
	Title    string   `xml:"title,omitempty"`
	Subtitle string   `xml:"subtitle,omitempty"`
	Entries  []AtomEntry
}

type AtomEntry struct {
	XMLName  xml.Name `xml:"entry"`
	Xmlns    string   `xml:"xmlns,attr,omitempty"`
	Title    string   `xml:"title,omitempty"`
	Subtitle string   `xml:"subtitle,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	entry := AtomEntry{}
	atom := AtomFeed{Xmlns: "http://www.w3.org/2005/AtomFeed", Entries: []AtomEntry{entry}}
	w.Header()["Content-Type"] = []string{"application/atom+xml"}
	fmt.Fprint(w, xml.Header)
	xmlEncoder := xml.NewEncoder(w)
	xmlEncoder.Encode(atom)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
