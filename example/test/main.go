package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	conf := map[string]string{
		"VerifyPayURI":   "https://account.dragonest.com/oauth/pay/order_verify?app_order_id=%v&client_id=%v&client_secret=%v&sign=%v",
		"PayCallBackURI": "/ly",
		"AppId":          "bd0f5e83300e3c00",
		"SecrectKey":     "afcf80822bcbd4491ac4e86ae5c96f70",
	}
	params := map[string]string{"amount": "6.00", "app_order_id": "15972129602229153", "app_uid": "CN8YZ5", "product_id": "600", "uid": "49915203", "sign": "b4843a1def81d90d451c53d5903febbf"}
	rsp, err := http.Get(fmt.Sprintf(conf["VerifyPayURI"], params["app_order_id"], conf["AppId"], conf["SecrectKey"],
		params["sign"]))
	if err != nil {
		panic(err)
	}
	dat, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	spew.Dump(string(dat))
}

func main2() {
	url := "http://172.16.212.9:14000/ly"
	body := `{"amount":"6.00","app_order_id":"155971487120","app_uid":"CU","product_id":"55","uid":"33556429","sign":"9bd9720b94b2381e4e634e42dee5fd5e"}`
	rsp, err := http.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	dat, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	spew.Dump(string(dat))
}
