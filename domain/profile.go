package domain

import "encoding/json"

//<div class="m-btn purple" data-v-bff6f798>离异</div>
//<div class="m-btn purple" data-v-bff6f798>42岁</div>
//<div class="m-btn purple" data-v-bff6f798>魔羯座(12.22-01.19)</div>
//<div class="m-btn purple" data-v-bff6f798>176cm</div>
//<div class="m-btn purple" data-v-bff6f798>68kg</div>
//<div class="m-btn purple" data-v-bff6f798>工作地:银川金凤区</div>
//<div class="m-btn purple" data-v-bff6f798>月收入:3-5千</div>
//<div class="m-btn purple" data-v-bff6f798>农林牧渔</div>
//<div class="m-btn purple" data-v-bff6f798>高中及以下</div>

type Profile struct {
	Name      string
	Gender    string
	Age       int
	Height    int
	Weight    int
	WorkPlace string
	InCome    string
	Marriage  string
	HoKou     string
	XinZuo    string
	House     string
	Car       string
}

func FromJsonObj(o interface{}) (profile Profile, err error) {
	var p Profile
	s, e := json.Marshal(o)
	if e != nil {
		return p, e
	}
	err = json.Unmarshal(s, &p)
	return p, err
}
