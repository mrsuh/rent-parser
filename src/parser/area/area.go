package area

import (
	config "github.com/mrsuh/cli-config"
	"rent-parser/src/tomita"
	"log"
	"regexp"
	"encoding/xml"
	"strconv"
	"strings"
)

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

func Parse(text string, response chan float32) {

	conf_instance := config.GetInstance()

	err := conf_instance.Init()

	if err != nil {
		log.Fatal(err)
	}

	conf := conf_instance.Get()

	tom := tomita.NewTomita(conf["tomita.bin"].(string), conf["tomita.conf.area"].(string))

	text = normalize(text)

	response <- getByXML(tom.Parse(text))
}

func normalize(text string) string {

	byte_text := []byte(text)

	re := regexp.MustCompile(`\?\W{0,10}$`)
	if nil != re.Find(byte_text) {
		return ""
	}

	re2 := regexp.MustCompile(`(публиковать.*|правила темы.*|[+\(\)-])`)
	byte_text = re2.ReplaceAll(byte_text, []byte(""))

	byte_text = []byte(strings.Replace(string(byte_text), `\n`, "\n", -1))
	byte_text = []byte(strings.Replace(string(byte_text), "-", " ", -1))

	re3 := regexp.MustCompile(`([\d=\+.\!?])([а-яеёa-z])`)
	byte_text = re3.ReplaceAll(byte_text, []byte("$1 $2"))

	re4 := regexp.MustCompile(`([а-яеёa-z])([\d=\+.\!?])`)
	byte_text = re4.ReplaceAll(byte_text, []byte("$1 $2"))

	return string(byte_text)
}

func getByXML(xml_row string) float32 {

	var document XmlFdoObject

	err := xml.Unmarshal([]byte(xml_row), &document)

	var area float32 = -1

	if err != nil {
		log.Println(err)

		return area
	}

	if len(document.Document.XMLFacts.FactAreaList) == 0 {

		return area
	}

	if len(document.Document.XMLFacts.FactAreaList[0].AreaList) == 0 {

		return area
	}

	area64, err := strconv.ParseFloat(document.Document.XMLFacts.FactAreaList[0].AreaList[0].Val, 32)

	return float32(area64)
}