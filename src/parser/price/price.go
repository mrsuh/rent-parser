package price

import (
	config "github.com/mrsuh/cli-config"
	"rent-parser/src/tomita"
	"log"
	"strings"
	"regexp"
	"encoding/xml"
	"strconv"
)

type Price struct {
	Value     float64
	Position int
	Sequence int
}


type XmlFull struct {
	XMLName xml.Name `xml:"Full"`
	Val string `xml:"val,attr"`
}

type XmlHalf struct {
	XMLName xml.Name `xml:"Half"`
	Val string `xml:"val,attr"`
}

type XmlShort struct {
	XMLName xml.Name `xml:"Short"`
	Val string `xml:"val,attr"`
}

type XmlFactPrice struct {
	XMLName xml.Name `xml:"FactPrice"`
	FullList []XmlFull `xml:"Full"`
	HalfList []XmlHalf `xml:"Half"`
	ShortList []XmlShort `xml:"Short"`
	FirstWord int `xml:"fw,attr"`
	LastWord  int `xml:"lw,attr"`
	Sequence  int `xml:"sn,attr"`
}

type XmlFacts struct {
	XMLName xml.Name `xml:"facts"`
	FactPriceList []XmlFactPrice `xml:"FactPrice"`
}

type XmlDocument struct {
	XMLName xml.Name `xml:"document"`
	XMLFacts XmlFacts  `xml:"facts"`
}

type XmlFdoObject struct {
	XMLName xml.Name `xml:"fdo_objects"`
	Document XmlDocument `xml:"document"`
}

func Parse(text string, response chan int) {

	conf_instance := config.GetInstance()

	err := conf_instance.Init()

	if err != nil {
		log.Fatal(err)
	}

	conf := conf_instance.Get()

	tom := tomita.NewTomita(conf["tomita.bin"].(string), conf["tomita.conf.price"].(string))

	text = normalize(text)

	response <- getByXML(tom.Parse(text))
}

func normalize(text string) string {
	text = strings.ToLower(text)

	byte_text := []byte(text)

	re := regexp.MustCompile(`\?\W{0,10}$`)
	if nil != re.Find(byte_text) {
		return ""
	}

	re2 := regexp.MustCompile(`(?i)(публиковать|варианты .* нашего сайта|правила темы|сайт)(.|\n)*`)
	byte_text = re2.ReplaceAll(byte_text, []byte(""))

	re3 := regexp.MustCompile(`(?i)http(s):(\w|\/|\.)*`)
	byte_text = re3.ReplaceAll(byte_text, []byte(""))

	byte_text = []byte(strings.Replace(string(byte_text), `\n`, "\n", -1))
	byte_text = []byte(strings.Replace(string(byte_text), "-", " ", -1))

	re4 := regexp.MustCompile(`([\d-:\/=\+.\!?\\\\])([а-яеёa-z-:\/=\+.\!?\\\\])`)
	byte_text = re4.ReplaceAll(byte_text, []byte("$1 $2"))

	re5 := regexp.MustCompile(`([а-яеёa-z-:\/=\+.\!?\\\\])([\d-:\/=\+.\!?\\\\])`)
	byte_text = re5.ReplaceAll(byte_text, []byte("$1 $2"))

	re6 := regexp.MustCompile(`(\d+)(\s+){0,5}[^a-zA-Zа-яА-Я\d](\s+){0,5}(\d+)`)
	byte_text = re6.ReplaceAll(byte_text, []byte("$1.$4"))

	return string(byte_text)
}

func getByXML(xml_row string) int {

	var document XmlFdoObject

	err := xml.Unmarshal([]byte(xml_row), &document)

	if err != nil {
		log.Println(err)

		return -1
	}

	price := Price{Value: -1, Position: 99, Sequence: 99}
	for _,fact := range document.Document.XMLFacts.FactPriceList {

		position := fact.FirstWord
		sequence := fact.Sequence

		var value float64
		var parse_err error
		switch true {
		case len(fact.FullList) > 0 && "" != fact.FullList[0].Val:
			value, parse_err = strconv.ParseFloat(fact.FullList[0].Val, 64)
			if nil != parse_err {
				log.Println(parse_err)
			}

			break
		case len(fact.HalfList) > 0 && "" != fact.HalfList[0].Val:
			value_str := fact.HalfList[0].Val

			parts := strings.Split(value_str, ".")

			value_str += strings.Repeat("0", 3 - len(parts[1]))
			value_str = strings.Replace(value_str, ".", "", -1)

			value, parse_err = strconv.ParseFloat(value_str, 64)
			if nil != parse_err {
				log.Println(parse_err)
			}

			break
		case len(fact.ShortList) > 0 && "" != fact.ShortList[0].Val:
			value, parse_err = strconv.ParseFloat(fact.ShortList[0].Val, 64)
			if nil != parse_err {
				log.Println(parse_err)
			}
			value *= 1000

			break
		}

		if sequence == price.Sequence && position < price.Position {
			price.Position = position
			price.Sequence = sequence
			price.Value = value
			continue
		}

		if sequence < price.Sequence {
			price.Position = position
			price.Sequence = sequence
			price.Value = value
		}
	}


	return int(price.Value)
}