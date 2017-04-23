package area

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

	tom := tomita.NewTomita(conf["tomita.bin"].(string), conf["tomita.conf.area"].(string))

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

type XmlArea struct {
	XMLName xml.Name `xml:"Area"`
	Val string `xml:"val,attr"`
}

type XmlFactArea struct {
	XMLName xml.Name `xml:"FactArea"`
	AreaList []XmlArea `xml:"Area"`
}

type XmlFacts struct {
	XMLName xml.Name `xml:"facts"`
	FactAreaList []XmlFactArea `xml:"FactArea"`
}

type XmlDocument struct {
	XMLName xml.Name `xml:"document"`
	XMLFacts XmlFacts  `xml:"facts"`
}

type XmlFdoObject struct {
	XMLName xml.Name `xml:"fdo_objects"`
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

	for _,fact := range document.Document.XMLFacts.FactAreaList {
		if len(fact.AreaList) == 0 {
			continue
		}
		contacts = append(contacts, fact.AreaList[0].Val)
	}

	return contacts
}