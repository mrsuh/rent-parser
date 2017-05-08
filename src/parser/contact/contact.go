package contact

import (
	config "github.com/mrsuh/cli-config"
	"rent-parser/src/tomita"
	"log"
	"strings"
	"regexp"
	"encoding/xml"
)

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

func Parse(text string, channel chan []string) {

	conf_instance := config.GetInstance()

	err := conf_instance.Init()

	if err != nil {
		log.Fatal(err)
	}

	conf := conf_instance.Get()

	tom := tomita.NewTomita(conf["tomita.bin"].(string), conf["tomita.conf.contact"].(string))

	text = normalize(text)

	channel <- getByXML(tom.Parse(text))
}

func normalize(text string) string {
	text = strings.ToLower(text)

	byte_text := []byte(text)

	re := regexp.MustCompile(`\?\W{0,10}$`)
	if nil != re.Find(byte_text) {
		return ""
	}

	re2 := regexp.MustCompile(`(публиковать.*|правила темы.*|[+\(\)-])`)
	byte_text = re2.ReplaceAll(byte_text, []byte(""))

	byte_text = []byte(strings.Replace(string(byte_text), `\n`, "\n", -1))
	byte_text = []byte(strings.Replace(string(byte_text), "-", " ", -1))

	re3 := regexp.MustCompile(`[+\(\)-]`)
	byte_text = re3.ReplaceAll(byte_text, []byte(""))

	re4 := regexp.MustCompile(`(\d)\s(\d)`)
	byte_text = re4.ReplaceAll(byte_text, []byte("$1$2"))

	re5 := regexp.MustCompile(`([=\+.\!?])(\d)`)
	byte_text = re5.ReplaceAll(byte_text, []byte("$1 $2"))

	re6 := regexp.MustCompile(`(\d)([=\+.\!?])`)
	byte_text = re6.ReplaceAll(byte_text, []byte("$1 $2"))

	return string(byte_text)
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

		re := regexp.MustCompile(`(\d|)(\d{10}$)`)
		contact := re.ReplaceAll([]byte(fact.ContactList[0].Val), []byte("$2"))

		contacts = append(contacts, string(contact))
	}

	return contacts
}