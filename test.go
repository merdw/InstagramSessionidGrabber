package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

var i int
var isim string
var mail string
var nick string
var numara string
var biosu string
var urlp string
var sessionid string
var tisim string
var tmail string
var tnick string
var tnumara string
var tbiosu string
var turlp string

type Welcome struct {
	User   User   `json:"user"`
	Status string `json:"status"`
}

type User struct {
	Username          string `json:"username"`
	FullName          string `json:"full_name"`
	Biography         string `json:"biography"`
	ExternalURL       string `json:"external_url"`
	PhoneNumber       string `json:"phone_number"`
	CountryCode       int64  `json:"country_code"`
	NationalNumber    int64  `json:"national_number"`
	Gender            int64  `json:"gender"`
	Email             string `json:"email"`
	NeedsPhoneConfirm bool   `json:"needs_phone_confirm"`
	TrustedUsername   string `json:"trusted_username"`
	TrustDays         int64  `json:"trust_days"`
}

func error(err interface{}) {
	if err != nil {
		panic(err)
	}

}

func SayiGoster() int {

	func() {
		i = 10

	}()
	return i
}

var guid string

func GetInfo(sessionid string) {

	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://i.instagram.com/api/v1/accounts/current_user/?edit=true")
	req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionid))
	req.Header.SetMethod("GET")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Instagram 76.0.0.15.395 Android (22/5.1.1; 240dpi; 1280x720; OnePlus; A5010; A5010; intel; tr_TR; 138226758)")
	req.Header.Add("device_id", "android-JDS095823049")

	err := client.Do(req, resp)
	if err != nil {
		log.Fatal(err)
	}
	fasthttp.ReleaseRequest(req)
	buffer := string(resp.Body())
	error(err)
	kisiJson := buffer

	var Welcome Welcome
	json.Unmarshal([]byte(kisiJson), &Welcome)
	isim = Welcome.User.FullName
	mail = Welcome.User.Email
	nick = Welcome.User.Username
	numara = Welcome.User.PhoneNumber
	biosu = Welcome.User.Biography
	urlp = Welcome.User.ExternalURL

}

func checkswap(sessionid string, target string) {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://i.instagram.com/api/v1/users/check_username/")
	req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionid))
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Instagram 85.0.0.21.100 Android (28/9; 380dpi; 1080x2147; OnePlus; HWEVA; OnePlus6T; qcom; en_US; 146536611")
	req.Header.Add("device_id", "android-JDS095823049")
	req.SetBody([]byte(fmt.Sprintf("username=%s", target)))
	err := client.Do(req, resp)
	if err != nil {
		log.Fatal(err)
	}
	fasthttp.ReleaseRequest(req)
	buffer := resp.Body()
	if string(buffer) == "This username isn't available." {
		fmt.Println("This user is NOT swappable")
	} else if string(buffer) == "This username isn't available. Please try again." {
		fmt.Println("This user is swappable.	")
	} else {
		fmt.Println(string(buffer) + "\n")
	}
	fasthttp.ReleaseResponse(resp)
}

func TestInfo(sessionid string) {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://i.instagram.com/api/v1/accounts/edit_profile/")
	req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionid))
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("user-agent", "Instagram 85.0.0.21.100 Android (28/9; 380dpi; 1080x2147; OnePlus; HWEVA; OnePlus6T; qcom; en_US; 146536611")
	req.Header.Add("device_id", "android-JDS095823049")

	req.SetBody([]byte(fmt.Sprintf("first_name=%s&email=%s&username=%s&phone_number=%s&biography=%s&external_url=%s", isim, mail, nick, numara, biosu, urlp)))

	err := client.Do(req, resp)

	if err != nil {
		log.Fatal(err)
	}
	fasthttp.ReleaseRequest(req)
	buffer := resp.Body()
	if string(buffer) == "This username isn't available." {
		fmt.Println("This user is NOT swappable\n")
	} else if string(buffer) == "This username isn't available. Please try again." {
		fmt.Println("This user is swappable.\n")
	} else {
		fmt.Println("apitest:")
		fmt.Println(string(buffer) + "\n")
	}
	fasthttp.ReleaseResponse(resp)
}

var rl int = 0
var attempts int = 0
var claimed bool = false
var severe_err bool = false
var target string
var ilk int = 0

func claimRequest(sessionid string, target string) {
	ilk++

	for !claimed && !severe_err {
		client := &fasthttp.Client{}
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		req.SetRequestURI("https://i.instagram.com/api/v1/accounts/edit_profile/")
		req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionid))
		req.Header.SetMethod("POST")
		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("user-agent", "Instagram 85.0.0.21.100 Android (28/9; 380dpi; 1080x2147; OnePlus; HWEVA; OnePlus6T; qcom; en_US; 146536611")
		req.Header.Add("device_id", "android-JDS095823049")
		req.SetBody([]byte(fmt.Sprintf("first_name=%s&email=%s&username=%s&phone_number=%s&biography=%s&external_url=%s", isim, mail, target, numara, biosu, urlp)))
		err := client.Do(req, resp)
		if err != nil {
			log.Fatal(err)
		}
		fasthttp.ReleaseRequest(req)
		buffer := resp.Body()
		if strings.Contains(string(buffer), "few minutes") {
			rl++

		} else if strings.Contains(string(buffer), "\"status\":\"ok\"") {
			fmt.Println(string(buffer))
			fmt.Println("claimed with api")
			claimed = true
		} else if strings.Contains(string(buffer), "logged") {
			severe_err = true
			fmt.Println("err::session went invalid")
			os.Exit(1)
		} else {
			attempts++
		}
		fasthttp.ReleaseResponse(resp)
	}

}

func GetInfoTarget(sessionid string) {

	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://i.instagram.com/api/v1/accounts/current_user/?edit=true")
	req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionid))
	req.Header.SetMethod("GET")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Instagram 76.0.0.15.395 Android (22/5.1.1; 240dpi; 1280x720; OnePlus; A5010; A5010; intel; tr_TR; 138226758)")
	req.Header.Add("device_id", "android-JDS095823049")

	err := client.Do(req, resp)
	if err != nil {
		log.Fatal(err)
	}
	fasthttp.ReleaseRequest(req)
	buffer := string(resp.Body())
	error(err)
	kisiJson := buffer

	var Welcome Welcome
	json.Unmarshal([]byte(kisiJson), &Welcome)
	tisim = Welcome.User.FullName
	tmail = Welcome.User.Email
	tnick = Welcome.User.Username
	tnumara = Welcome.User.PhoneNumber
	tbiosu = Welcome.User.Biography
	turlp = Welcome.User.ExternalURL

}
