package parser

import (
	config "github.com/mrsuh/cli-config"
	"rent-parser/src/tomita"
	"log"
	"strings"
	"regexp"
	"encoding/xml"
	"fmt"
)

const ROOM = 0
const FLAT_1 = 1
const FLAT_2 = 2
const FLAT_3 = 3
const FLAT_4 = 4
const STUDIO = 5
const WRONG = 6

type Type struct {
	Type     int
	Position int
	Sequence int
}

type XmlType struct {
	XMLName   xml.Name `xml:"Type"`
	Value     string `xml:"val,attr"`
}

type XmlWrong struct {
	XMLName   xml.Name `xml:"Wrong"`
	Value     string `xml:"val,attr"`
}

type XmlError struct {
	XMLName   xml.Name `xml:"Error"`
	Value     string `xml:"val,attr"`
}

type XmlFactRealty struct {
	XMLName  xml.Name `xml:"FactRealty"`
	TypeList []XmlType `xml:"Type"`
	FirstWord int `xml:"fw,attr"`
	LastWord  int `xml:"lw,attr"`
	Sequence  int `xml:"sn,attr"`
}

type XmlFactRent struct {
	XMLName  xml.Name `xml:"FactRent"`
	TypeList []XmlType `xml:"Type"`
	FirstWord int `xml:"fw,attr"`
	LastWord  int `xml:"lw,attr"`
	Sequence  int `xml:"sn,attr"`
}

type XmlFactNeighbor struct {
	XMLName  xml.Name `xml:"FactNeighbor"`
	TypeList []XmlType `xml:"Type"`
	FirstWord int `xml:"fw,attr"`
	LastWord  int `xml:"lw,attr"`
	Sequence  int `xml:"sn,attr"`
}

type XmlFactWrong struct {
	XMLName  xml.Name `xml:"FactWrong"`
	WrongList []XmlWrong `xml:"Wrong"`
	FirstWord int `xml:"fw,attr"`
	LastWord  int `xml:"lw,attr"`
	Sequence  int `xml:"sn,attr"`
}

type XmlFactError struct {
	XMLName  xml.Name `xml:"FactError"`
	ErrorList []XmlError `xml:"Error"`
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

func Parse(text string, response chan int) {

	conf_instance := config.GetInstance()

	err := conf_instance.Init()

	if err != nil {
		log.Fatal(err)
	}

	conf := conf_instance.Get()

	tom := tomita.NewTomita(conf["tomita.bin"].(string), conf["tomita.conf.type"].(string))

	text = normalize(text)

	response <- getByXML(tom.Parse(text))
}

func PreValid(rawText string) bool {

	normalizedText := normalize(rawText)

	rawByteText := []byte(rawText)
	normalizedByteText := []byte(normalizedText)

	flats := []string{
		"кв",
		"ком",
		"одн",
		"дву",
		"тр(ё|e)",
		"студ",
	}

	reFlat := regexp.MustCompile(fmt.Sprintf(`(?i).*(%s).*`, strings.Join(flats, "|")))
	if !reFlat.Match(normalizedByteText) {
		return false
	}

	vkId := regexp.MustCompile(`(?i)^\[id\d+`)
	if vkId.Match(rawByteText) {
		return false
	}

	search := regexp.MustCompile(`(?i)(сним|ищ)(у|ем)[\w]*[^\w](кварт|комн|одн(ок|ушк)|дву(хк|шк)|тр[её](хк|шк))`)
	if search.Match(normalizedByteText) {
		return false
	}

	return true
}

func normalize(raw_text string) string {

	byte_text := []byte(strings.TrimSpace(raw_text))

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

	months := []string{"январ", "феврал", "март", "апрел", "ма(я|е|й)", "июн", "июл", "август", "сентябр", "октябр", "ноябр", "декабр"}

	tmp_re := regexp.MustCompile(fmt.Sprintf(`(?i)\d{0,2}[^0-9\.\!\?;]{0,3}(%s)[а-я]{0,4}`, strings.Join(months, "|")))

	byte_text = tmp_re.ReplaceAll(byte_text, []byte(``))

	flats := [...]string{
		`(?i)(\s|[0-9])кк(\s|\.)`,
		`(?i)(\s|[0-9])ккв(\s|\.)`,
		`(?i)(\s|[0-9])к\.\s{0,2}к(\s|\.)`,
		`(?i)(\s|[0-9])к\.\s{0,2}квартира(\s|\.)`,
		`(?i)(\s|[0-9])к\.\s{0,2}квартиру(\s|\.)`,
		`(?i)(\s|[0-9])к\.\s{0,2}кв(\s|\.)`,
		`(?i)(\s|[0-9])к\.\s{0,2}кварт(\s|\.)`,
		`(?i)(\s|[0-9])комн\.\s{0,2}кв(\s|\.)`,
		`(?i)(\s|[0-9])хкк(\s|\.)`,
		`(?i)(\s|[0-9])х\.\s{0,2}ком.*\b(\s|\.)`,
		`(?i)(\s|[0-9])хк\.\s{0,2}кв.*\b(\s|\.)`,
		`(?i)(\s|[0-9])х\.\s{0,2}к.*\b(\s|\.)кв.*\b(\s|\.)`,
		`(?i)(\s|[0-9])хк\.\s{0,2}к.*\b(\s|\.)кв.*\b(\s|\.)`,
	}

	for _, flat := range flats {
		tmp_re := regexp.MustCompile(flat)
		byte_text = tmp_re.ReplaceAll(byte_text, []byte("$1 комнатная квартира "))
	}

	re4 := regexp.MustCompile(`(?i)\d{1,3}\s{0,10}(кв(\.|\s){0,1}м(\.|\s){0,1}|м²|м(\.|\s))`)
	byte_text = re4.ReplaceAll(byte_text, []byte(" "))

	re5 := regexp.MustCompile(`(?i)([\d-=\+.\!?])([а-яеёa-z])`)
	byte_text = re5.ReplaceAll(byte_text, []byte("$1 $2"))

	re6 := regexp.MustCompile(`(?i)([а-яеёa-z])([\d-=\+.\!?])`)
	byte_text = re6.ReplaceAll(byte_text, []byte("$1 $2"))

	re7 := regexp.MustCompile(`(?i)\sквартир[а-яА-Яeё]*`)
	byte_text = re7.ReplaceAll(byte_text, []byte(" квартира "))

	re8 := regexp.MustCompile(`(?i)\sкомната[а-яА-Яeё]*`)
	byte_text = re8.ReplaceAll(byte_text, []byte(" комната "))

	text := []rune(string(byte_text))

	if len(text) > 500 {
		byte_text = []byte(string(text[:500]))
	}

	return string(byte_text)
}

func getByXML(xml_row string) int {

	if xml_row == "" {
		return WRONG
	}

	var document XmlFdoObject

	err := xml.Unmarshal([]byte(xml_row), &document)

	if err != nil {
		log.Println(err)

		return WRONG
	}

	if len(document.Document.XMLFacts.FactErrorList) > 0 {

		return WRONG
	}

	rent := Type{Type: -1, Position: 99, Sequence: 99}
	for _, fact := range document.Document.XMLFacts.FactRentList {
		if len(fact.TypeList) == 0 {
			continue
		}

		position := fact.FirstWord
		sequence := fact.Sequence
		rtype := getTypeByString(fact.TypeList[0].Value)

		if STUDIO == rtype {
			rent.Position = position
			rent.Sequence = sequence
			rent.Type = rtype

			break
		}

		if sequence == rent.Sequence && position < rent.Position {
			rent.Position = position
			rent.Sequence = sequence
			rent.Type = rtype

			continue
		}

		if sequence < rent.Sequence {
			rent.Position = position
			rent.Sequence = sequence
			rent.Type = rtype

			continue
		}
	}

	neighbor := Type{Type: -1, Position: 99, Sequence: 99}
	for _, fact := range document.Document.XMLFacts.FactNeighborList {
		if len(fact.TypeList) == 0 {
			continue
		}

		position := fact.FirstWord
		sequence := fact.Sequence
		rtype := getTypeByString(fact.TypeList[0].Value)

		if STUDIO == rtype {
			neighbor.Position = position
			neighbor.Sequence = sequence
			neighbor.Type = rtype

			break
		}

		if sequence == neighbor.Sequence && position < neighbor.Position {
			neighbor.Position = position
			neighbor.Sequence = sequence
			neighbor.Type = rtype

			continue
		}

		if sequence < neighbor.Sequence {
			neighbor.Position = position
			neighbor.Sequence = sequence
			neighbor.Type = rtype

			continue
		}
	}

	realty := Type{Type: -1, Position: 99, Sequence: 99}
	for _, fact := range document.Document.XMLFacts.FactRealtyList {
		if len(fact.TypeList) == 0 {
			continue
		}

		position := fact.FirstWord
		sequence := fact.Sequence
		rtype := getTypeByString(fact.TypeList[0].Value)

		if STUDIO == rtype {
			realty.Position = position
			realty.Sequence = sequence
			realty.Type = rtype

			break
		}

		if sequence == realty.Sequence && position < realty.Position {
			realty.Position = position
			realty.Sequence = sequence
			realty.Type = rtype

			continue
		}

		if sequence < realty.Sequence {
			realty.Position = position
			realty.Sequence = sequence
			realty.Type = rtype

			continue
		}
	}

	wrong := Type{Type: -1, Position: 99, Sequence: 99}
	for _, fact := range document.Document.XMLFacts.FactWrongList {
		if len(fact.WrongList) == 0 {
			continue
		}

		position := fact.FirstWord
		sequence := fact.Sequence

		if sequence == wrong.Sequence && position < wrong.Position {
			wrong.Position = position
			wrong.Sequence = sequence
			wrong.Type = WRONG

			continue
		}

		if sequence < wrong.Sequence {
			wrong.Position = position
			wrong.Sequence = sequence
			wrong.Type = WRONG

			continue
		}
	}

	switch true {
	case -1 != wrong.Type &&
		(wrong.Sequence < rent.Sequence || (wrong.Sequence == rent.Sequence && wrong.Position < rent.Position)) &&
		(wrong.Sequence < realty.Sequence || (wrong.Sequence == realty.Sequence && wrong.Position < realty.Position)):
		return WRONG

	case -1 != rent.Type:
		return rent.Type

	case -1 != neighbor.Type:
		return neighbor.Type

	case -1 != wrong.Type:
		return wrong.Type

	case -1 != realty.Type:
		return realty.Type
	default:
		return WRONG
	}

	return WRONG
}

func getTypeByString(raw_text string) int {

	if "" == raw_text {
		return WRONG
	}

	text := strings.ToLower(raw_text)

	if -1 != strings.Index(text, "студи") {
		return STUDIO
	}

	re := regexp.MustCompile(`(^|\W)комнаты($|\W)`)

	if nil != re.Find([]byte(text)) {
		return ROOM
	}

	if -1 != strings.Index(text, "1") {
		return FLAT_1
	}

	if -1 != strings.Index(text, "2") {
		return FLAT_2
	}

	if -1 != strings.Index(text, "3") {
		return FLAT_3
	}

	re2 := regexp.MustCompile(`(([^\d,\.!?]|^)[4-9]\D{0,30}квартир|четыр\Sх|много)|(квартир\D{0,3}1\D.{0,10}комнатн)`)

	if nil != re2.Find([]byte(text)) {
		return FLAT_4
	}

	re3 := regexp.MustCompile(`(^|\W)квартир\W{1,4}($|\W)`)

	if nil != re3.Find([]byte(text)) {
		return FLAT_1
	}

	re4 := regexp.MustCompile(`(^|\W)комнат`)

	if nil != re4.Find([]byte(text)) {
		return ROOM
	}

	return WRONG
}
