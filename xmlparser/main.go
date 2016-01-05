package main

import (
    "io/ioutil"
    "encoding/xml"
    "fmt"
    )

type div struct {
    XMLName xml.Name `xml:"div"`
    Content string `xml:",chardata"`
}

func main() {
    d := div{}
    xmlContent, _  := ioutil.ReadFile("example.xml")
    if err := xml.Unmarshal(xmlContent, &d); err != nil {
        panic(err)
    }
    fmt.Println("XMLName : ", d.XMLName)
    fmt.Println("Content : ", d.Content)
    parseSO()
}
type Rss2 struct {
    ItemList []Item `xml:"channel>item"`
}
type Item struct {
    Title       string      `xml:"title"`
    Link        string      `xml:"link"`
    Description string      `xml:"description"`
    PubDate     string      `xml:"pubDate"`
    GUID        string      `xml:"guid"`
}

func parseSO() {
    r := Rss2{}
    xmlContent, _ := ioutil.ReadFile("example2.xml")
    if err := xml.Unmarshal(xmlContent, &r); err != nil {
        panic(err)
    }
    fmt.Println("Sample ", r)

}
