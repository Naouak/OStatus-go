package atom

import (
	"reflect"
	"encoding/xml"
)

type Entry struct {
	XMLName  xml.Name `xml:"entry"`
	Xmlns    string   `xml:"xmlns,attr,omitempty"`
	Title    string   `xml:"title,omitempty"`
	Subtitle string   `xml:"subtitle,omitempty"`
}

type atomTag struct {
	name       string
	fieldValue interface{}
}

func addField(tag atomTag, entry Entry) Entry {
	switch tag.name {
	case "title":
		entry.Title = tag.fieldValue.(string)
	case "subtitle":
		entry.Subtitle = tag.fieldValue.(string)
	}
	return entry
}

func ParseEntry(v interface{}) Entry {
	entry := Entry{}
	typeOf := reflect.TypeOf(v)

	for i := 0; i < typeOf.NumField(); i++ {
		field := reflect.TypeOf(v).Field(i)
		entry = addField(atomTag{
			name:       field.Tag.Get("atom"),
			fieldValue: reflect.ValueOf(v).Field(i),
		}, entry)
	}

	return Entry{}
}
