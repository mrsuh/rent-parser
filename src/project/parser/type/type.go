package parser

import (
	"log"
	"strings"
	"regexp"
	"project/tomita"
	"project/config"
	"encoding/xml"
	"fmt"
)

func Parse(text string, response chan []string) {

	conf_instance := config.GetInstance()

	err := conf_instance.Init()

	if err != nil {
		log.Fatal(err)
	}

	conf := conf_instance.Get()

	tom := tomita.NewTomita(conf["tomita.bin"].(string), conf["tomita.conf.type"].(string))

	response <- getByXML(tom.Parse(text))
}

func normalize(text string) string {
	text = strings.ToLower(text)
	bynary_text := []byte(text)
	re, re_err := regexp.Compile(`\?\W{0,10}$`)

	if re_err != nil {

	}

	if len(re.Find(bynary_text)) > 0 {
		return ""
	}

	re2, re2_err := regexp.Compile(`/(публиковать.*|правила темы.*|[+\(\)-])/ui`)

	if re2_err != nil {

	}

	re2.ReplaceAll(bynary_text, []byte(""))

	//$text = preg_replace('//ui', '', $text);
	//$text = preg_replace('/(\d)\s(\d)/ui', '$1$2', $text);

	return ""
}

type XmlType struct {
	XMLName xml.Name `xml:"Type"`
	Val     string `xml:"val,attr"`
}

type XmlFactRealty struct {
	XMLName xml.Name `xml:"FactRealty"`
	TypeList   []XmlType `xml:"Type"`
}

type XmlFactRent struct {
	XMLName xml.Name `xml:"FactRent"`
	TypeList   []XmlType `xml:"Type"`
}

type XmlFactNeighbor struct {
	XMLName xml.Name `xml:"FactNeighbor"`
	TypeList   []XmlType `xml:"Type"`
}

type XmlFactWrong struct {
	XMLName xml.Name `xml:"FactWrong"`
	TypeList   []XmlType `xml:"Type"`
}

type XmlFactError struct {
	XMLName xml.Name `xml:"FactError"`
	TypeList   []XmlType `xml:"Type"`
}

type XmlFacts struct {
	XMLName          xml.Name `xml:"facts"`
	FactRealtyList   []XmlFactRealty `xml:"FactRealty"`
	FactRentList     []XmlFactRent `xml:"FactRent"`
	FactNeighborList []XmlFactNeighbor `xml:"FactNeighbor"`
	FactWrongList    []XmlFactWrong `xml:"FactWrong"`
	FactErrorList    []XmlFactError `xml:"FactError"`
}

type XmlDocument struct {
	XMLName  xml.Name `xml:"document"`
	XMLFacts XmlFacts  `xml:"facts"`
}

type XmlFdoObject struct {
	XMLName  xml.Name `xml:"fdo_objects"`
	Document XmlDocument `xml:"document"`
}

func getByXML(xml_row string) []string {

	fmt.Println(xml_row)

	var document XmlFdoObject

	err := xml.Unmarshal([]byte(xml_row), &document)

	contacts := make([]string, 0)

	if err != nil {
		log.Println(err)

		return contacts
	}

	for _, fact := range document.Document.XMLFacts.FactRealtyList {
		if len(fact.TypeList) == 0 {
			continue
		}
		contacts = append(contacts, fact.TypeList[0].Val)
	}

	for _, fact := range document.Document.XMLFacts.FactRentList {
		if len(fact.TypeList) == 0 {
			continue
		}
		contacts = append(contacts, fact.TypeList[0].Val)
	}

	for _, fact := range document.Document.XMLFacts.FactNeighborList {
		if len(fact.TypeList) == 0 {
			continue
		}
		contacts = append(contacts, fact.TypeList[0].Val)
	}

	for _, fact := range document.Document.XMLFacts.FactWrongList {
		if len(fact.TypeList) == 0 {
			continue
		}
		contacts = append(contacts, fact.TypeList[0].Val)
	}

	for _, fact := range document.Document.XMLFacts.FactErrorList {
		if len(fact.TypeList) == 0 {
			continue
		}
		contacts = append(contacts, fact.TypeList[0].Val)
	}

	return contacts
}
