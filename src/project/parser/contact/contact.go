package contact

import (
	"log"
	"strings"
	"regexp"
	"project/tomita"
	"project/config"
	"encoding/xml"
)

func Parse(text string, channel chan []string) {

	conf_instance := config.GetInstance()

	err := conf_instance.Init()

	if err != nil {
		log.Fatal(err)
	}

	conf := conf_instance.Get()

	tom := tomita.NewTomita(conf["tomita.bin"].(string), conf["tomita.conf.contact"].(string))

	channel <- getByXML(tom.Parse(text))
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

type XmlContact struct {
	XMLName xml.Name `xml:"Contact"`
	Val string `xml:"val,attr"`
}

type XmlFactContact struct {
	XMLName xml.Name `xml:"FactContact"`
	ContactList []XmlContact `xml:"Contact"`
}

type XmlFacts struct {
	XMLName xml.Name `xml:"facts"`
	FactContactList []XmlFactContact `xml:"FactContact"`
}

type XmlDocument struct {
	XMLName xml.Name `xml:"document"`
	FactFacts XmlFacts  `xml:"facts"`
}

type XmlFdoObject struct {
	XMLName xml.Name `xml:"fdo_objects"`
	Document XmlDocument `xml:"document"`
}

func getByXML(xml_row string) []string {

	var document XmlFdoObject

	err := xml.Unmarshal([]byte(xml_row), &document)

	contacts := make([]string, 0)

	if err != nil {
		log.Println(err)

		return contacts
	}

	for _,fact := range document.Document.FactFacts.FactContactList {
		if len(fact.ContactList) == 0 {
			continue
		}
		contacts = append(contacts, fact.ContactList[0].Val)
	}

	return contacts
}