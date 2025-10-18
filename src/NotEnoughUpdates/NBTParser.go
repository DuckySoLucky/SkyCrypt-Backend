package notenoughupdates

/*
NBT Parser for NotEnoughUpdates nbttag format in Kotlin provided by @nea89o: https://github.com/nea89o/Firmament/blob/d2f240ff0ca0d27f417f837e706c781a98c31311/src/main/kotlin/util/LegacyTagParser.kt
Inspired by SkyCryptv2 version: https://github.com/DuckySoLucky/SkyCryptv2/blob/dev/src/lib/server/helper/NotEnoughUpdates/NBTParser.ts
*/

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
)

type TagParsingException struct {
	BaseString string
	Offset     int
	Message    string
}

func (e *TagParsingException) Error() string {
	return fmt.Sprintf("%s at %d in `%s`", e.Message, e.Offset, e.BaseString)
}

type StringRacer struct {
	backing string
	idx     int
	stack   []int
}

func NewStringRacer(backing string) *StringRacer {
	return &StringRacer{
		backing: backing,
		idx:     0,
		stack:   make([]int, 0),
	}
}

func (sr *StringRacer) PushState() {
	sr.stack = append(sr.stack, sr.idx)
}

func (sr *StringRacer) PopState() {
	if len(sr.stack) > 0 {
		sr.idx = sr.stack[len(sr.stack)-1]
		sr.stack = sr.stack[:len(sr.stack)-1]
	}
}

func (sr *StringRacer) DiscardState() {
	if len(sr.stack) > 0 {
		sr.stack = sr.stack[:len(sr.stack)-1]
	}
}

func (sr *StringRacer) Peek(count int) string {
	end := sr.idx + count
	if end > len(sr.backing) {
		end = len(sr.backing)
	}
	return sr.backing[sr.idx:end]
}

func (sr *StringRacer) Finished() bool {
	return len(sr.Peek(1)) == 0
}

func (sr *StringRacer) PeekReq(count int) *string {
	peeked := sr.Peek(count)
	if len(peeked) == count {
		return &peeked
	}
	return nil
}

func (sr *StringRacer) ConsumeCountReq(count int) *string {
	peeked := sr.PeekReq(count)
	if peeked != nil {
		sr.idx += count
	}
	return peeked
}

func (sr *StringRacer) TryConsume(s string) bool {
	if sr.Peek(len(s)) == s {
		sr.idx += len(s)
		return true
	}
	return false
}

func (sr *StringRacer) ConsumeWhile(conditionFn func(string) bool) string {
	consumed := ""
	for conditionFn(consumed + sr.Peek(1)) {
		consumed += sr.Peek(1)
		sr.idx++
	}
	return consumed
}

func (sr *StringRacer) Expect(s string, errorMessage string) {
	if !sr.TryConsume(s) {
		sr.Error(errorMessage)
	}
}

func (sr *StringRacer) Error(errorMessage string) {
	panic(&TagParsingException{
		BaseString: sr.backing,
		Offset:     sr.idx,
		Message:    errorMessage,
	})
}

var NBTPatterns = map[string]*regexp.Regexp{
	"DOUBLE":         regexp.MustCompile(`^([-+]?[0-9]*\.?[0-9]+)[dD]$`),
	"FLOAT":          regexp.MustCompile(`^([-+]?[0-9]*\.?[0-9]+)[fF]$`),
	"BYTE":           regexp.MustCompile(`^([-+]?[0-9]+)[bB]$`),
	"LONG":           regexp.MustCompile(`^([-+]?[0-9]+)[lL]$`),
	"SHORT":          regexp.MustCompile(`^([-+]?[0-9]+)[sS]$`),
	"INTEGER":        regexp.MustCompile(`^([-+]?[0-9]+)$`),
	"DOUBLE_UNTYPED": regexp.MustCompile(`^([-+]?[0-9]*\.?[0-9]+)$`),
	"ROUGH_PATTERN":  regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]*[dDbBfFlLsS]?$`),
}

type NBTTagParser struct {
	racer   *StringRacer
	baseTag interface{}
}

func NewNBTTagParser(s string) *NBTTagParser {
	parser := &NBTTagParser{
		racer: NewStringRacer(s),
	}
	parser.baseTag = parser.ParseTag()
	return parser
}

func (p *NBTTagParser) DigitRange() string {
	return "0123456789-"
}

func ParseNBT(s string) interface{} {
	return NewNBTTagParser(s).baseTag
}

func (p *NBTTagParser) SkipWhitespace() {
	p.racer.ConsumeWhile(func(s string) bool {
		if len(s) == 0 {
			return false
		}
		return unicode.IsSpace(rune(s[len(s)-1]))
	})
}

func (p *NBTTagParser) ParseTag() map[string]interface{} {
	p.SkipWhitespace()
	p.racer.Expect("{", "Expected '{' at start of tag")
	p.SkipWhitespace()

	tag := make(map[string]interface{})

	for !p.racer.TryConsume("}") {
		p.SkipWhitespace()
		key := p.ParseIdentifier()
		p.SkipWhitespace()
		p.racer.Expect(":", "Expected ':' after identifier in tag")
		p.SkipWhitespace()
		value := p.ParseAny()
		tag[key] = value
		p.racer.TryConsume(",")
		p.SkipWhitespace()
	}

	return tag
}

func (p *NBTTagParser) ParseAny() interface{} {
	p.SkipWhitespace()
	nextChar := p.racer.PeekReq(1)
	if nextChar == nil {
		p.racer.Error("Expected new object, found EOF")
	}

	switch *nextChar {
	case "{":
		return p.ParseTag()
	case "[":
		return p.ParseList()
	case "\"":
		return p.ParseStringTag()
	default:
		if strings.Contains(p.DigitRange(), *nextChar) {
			return p.ParseNumericTag()
		}
	}

	p.racer.Error("Unexpected token found. Expected start of new element")
	return nil
}

func (p *NBTTagParser) ParseList() []interface{} {
	p.SkipWhitespace()
	p.racer.Expect("[", "Expected '[' at start of tag")
	p.SkipWhitespace()

	list := make([]interface{}, 0)

	for !p.racer.TryConsume("]") {
		p.SkipWhitespace()
		p.racer.PushState()

		maybeIndex := p.racer.ConsumeWhile(func(s string) bool {
			for _, c := range s {
				if !strings.Contains(p.DigitRange(), string(c)) {
					return false
				}
			}
			return true
		})

		p.SkipWhitespace()
		if !p.racer.TryConsume(":") || len(maybeIndex) == 0 {
			p.racer.PopState()
			list = append(list, p.ParseAny())
		} else {
			p.racer.DiscardState()
			p.SkipWhitespace()
			list = append(list, p.ParseAny())
		}

		p.SkipWhitespace()
		p.racer.TryConsume(",")
	}

	return list
}

func (p *NBTTagParser) ParseQuotedString() string {
	p.SkipWhitespace()
	p.racer.Expect("\"", "Expected '\"' at string start")

	result := ""

	for {
		char := p.racer.ConsumeCountReq(1)
		if char == nil {
			p.racer.Error("Unfinished string")
		}

		if *char == "\"" {
			return result
		}

		if *char == "\\" {
			escaped := p.racer.ConsumeCountReq(1)
			if escaped == nil {
				p.racer.Error("Unfinished backslash escape")
			}

			if *escaped != "\"" && *escaped != "\\" {
				p.racer.idx--
				p.racer.Error(fmt.Sprintf("Invalid backslash escape '%s'", *escaped))
			}

			result += *escaped
		} else {
			result += *char
		}
	}
}

func (p *NBTTagParser) ParseStringTag() string {
	return p.ParseQuotedString()
}

func (p *NBTTagParser) ParseNumericTag() interface{} {
	p.SkipWhitespace()
	text := p.racer.ConsumeWhile(func(s string) bool {
		return NBTPatterns["ROUGH_PATTERN"].MatchString(s)
	})

	if len(text) == 0 {
		p.racer.Error("Expected numeric tag (starting with either -, +, . or a digit)")
	}

	patterns := []string{"DOUBLE", "FLOAT", "BYTE", "LONG", "SHORT", "INTEGER", "DOUBLE_UNTYPED"}

	for _, patternName := range patterns {
		regex := NBTPatterns[patternName]
		match := regex.FindStringSubmatch(text)
		if len(match) > 1 {
			switch patternName {
			case "DOUBLE", "DOUBLE_UNTYPED":
				if val, err := strconv.ParseFloat(match[1], 64); err == nil {
					return val
				}
			case "FLOAT":
				if val, err := strconv.ParseFloat(match[1], 32); err == nil {
					return float32(val)
				}
			case "BYTE":
				if val, err := strconv.ParseInt(match[1], 10, 8); err == nil {
					return int8(val)
				}
			case "LONG":
				if val, err := strconv.ParseInt(match[1], 10, 64); err == nil {
					return val
				}
			case "SHORT":
				if val, err := strconv.ParseInt(match[1], 10, 16); err == nil {
					return int16(val)
				}
			case "INTEGER":
				if val, err := strconv.ParseInt(match[1], 10, 32); err == nil {
					return int32(val)
				}
			}
		}
	}

	panic(fmt.Sprintf("Could not properly parse numeric tag '%s', despite passing verification. BAD DEV", text))
}

func (p *NBTTagParser) ParseIdentifier() string {
	p.SkipWhitespace()
	if p.racer.Peek(1) == "\"" {
		return p.ParseQuotedString()
	}

	return p.racer.ConsumeWhile(func(s string) bool {
		if len(s) == 0 {
			return true
		}
		lastChar := s[len(s)-1]
		return lastChar != ':' && !unicode.IsSpace(rune(lastChar))
	})
}

func ParseNBTToItem(nbtString string) (skycrypttypes.Tag, bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error parsing NBT: %v\n", r)
		}
	}()

	try := func() (skycrypttypes.Tag, bool) {
		parsed := ParseNBT(nbtString)
		if tagMap, ok := parsed.(map[string]interface{}); ok {
			// Marshal the map to JSON, then unmarshal into the struct (So we can cast it as models.Tag)
			data, err := json.Marshal(tagMap)
			if err != nil {
				return skycrypttypes.Tag{}, false
			}
			var item skycrypttypes.Tag
			if err := json.Unmarshal(data, &item); err != nil {
				return skycrypttypes.Tag{}, false
			}
			return item, true
		}
		return skycrypttypes.Tag{}, false
	}

	return try()
}
