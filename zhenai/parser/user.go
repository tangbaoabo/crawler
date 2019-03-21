package parser

import (
	"crawler/engine"
	"crawler/domain"
	"regexp"
	"strconv"
)

//<div class="m-btn purple" data-v-bff6f798>离异</div>
//<div class="m-btn purple" data-v-bff6f798>42岁</div>
//<div class="m-btn purple" data-v-bff6f798>魔羯座(12.22-01.19)</div>
//<div class="m-btn purple" data-v-bff6f798>176cm</div>
//<div class="m-btn purple" data-v-bff6f798>68kg</div>
//<div class="m-btn purple" data-v-bff6f798>工作地:银川金凤区</div>
//<div class="m-btn purple" data-v-bff6f798>月收入:3-5千</div>
//<div class="m-btn purple" data-v-bff6f798>农林牧渔</div>
//<div class="m-btn purple" data-v-bff6f798>高中及以下</div>
//<h1 class="nickName" data-v-5b109fc3>随缘</h1>
// <div class="m-btn pink" data-v-bff6f798>籍贯:陕西商洛</div>
//<div class="m-btn pink" data-v-bff6f798>未买车</div>
//<div class="m-btn pink" data-v-bff6f798>租房</div>

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([0-9]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<][丧偶|离异|未婚]?)</div>`)
var xinZuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(.{3}\([0-9]{1,2}\.[0-9]{1,2}-[0-9]{1,2}\.[0-9]{1,2}\))</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([0-9]{2,3})cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([0-9]{2,3})kg</div>`)
var workPlaceRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>工作地:([^<]+)</div>`)
var inComeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)
var nameRe = regexp.MustCompile(`<h1 class="nickName" data-v-5b109fc3>([^<]+)</h1>`)
var hoKouRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>籍贯:([^<]+)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>(.{2}车)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([^<]+房)</div>`)

//id
var IdRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseUser(contents []byte, extraInfo map[string]interface{}) engine.ParseResult {

	profile := domain.Profile{}
	match := extractProfile(contents, ageRe)
	if age, err := strconv.Atoi(match); err == nil {
		profile.Age = age
	} else {
		profile.Age = 0
	}

	profile.Marriage = extractProfile(contents, marriageRe)

	profile.XinZuo = extractProfile(contents, xinZuoRe)

	profile.Gender = extraInfo["sex"].(string)

	height := extractProfile(contents, heightRe)
	if height, err := strconv.Atoi(height); err == nil {
		profile.Height = height
	} else {
		profile.Height = 0
	}

	profile.WorkPlace = extractProfile(contents, workPlaceRe)

	weight := extractProfile(contents, weightRe)
	if weight, err := strconv.Atoi(weight); err == nil {
		profile.Weight = weight
	} else {
		profile.Weight = 0
	}

	profile.InCome = extractProfile(contents, inComeRe)

	profile.Name = extractProfile(contents, nameRe)

	profile.HoKou = extractProfile(contents, hoKouRe)

	profile.Car = extractProfile(contents, carRe)

	profile.House = extractProfile(contents, houseRe)

	url, ok := extraInfo["url"].(string)
	if !ok {
		url = ""
	}
	id := extractProfile([]byte(url), IdRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      id,
				PayLoad: profile,
			},
		},
	}

	return result
}

func extractProfile(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
