package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/valyala/fasthttp"
)

var sessiond string
var sorun bool = false
var csrf string
var zaman int64
var sessiondweb string

func LoginWeb(username string, password string) string {
	csrf = randomdata.RandStringRunes(15)
	now := time.Now()
	zaman = now.Unix()
	fmt.Println(zaman)
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://instagram.com/accounts/login/ajax/")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edge/101.0.1210.32")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("x-csrftoken", fmt.Sprintf("%s", csrf))
	req.SetBody([]byte(fmt.Sprintf("username=%s&enc_password=#PWD_INSTAGRAM_BROWSER:0:%d:%s", username, zaman, password)))

	err := client.Do(req, resp)
	if err != nil {
		log.Fatal(err)
	}
	fasthttp.ReleaseRequest(req)
	buffer := string(resp.Body())

	if strings.Contains(string(buffer), "checkpoint_required") {
		fmt.Println("WEB Suphe Dusen hesap" + username)
		sorun = true
		file, err := os.OpenFile("websuphe.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s:%s:%s\n", username, password, buffer))
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("yazildi! websorun")
		}
	} else if strings.Contains(string(buffer), "\"oneTapPrompt\":true") {

		deneme := resp.Header
		re := regexp.MustCompile(`sessionid=(.*); Domain=.instagram.com;`)
		match := re.FindStringSubmatch(deneme.String())

		fmt.Println(match[1])
		sessiondweb = match[1]

		fmt.Println("web giris basarili")
		error(err)

		return sessiondweb

	} else {
		fmt.Println("WEB SORUN VAR: " + username)

		fmt.Println(string(buffer))
		sorun = true
		file, err := os.OpenFile("websorun.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s:%s:%s\n", username, password, buffer))
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("yazildi! websorun")
		}
	}
	return sessiondweb
}

func LoginApi(username string, password string) string {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://i.instagram.com/api/v1/accounts/login/")
	req.Header.Add("Accept-Language", "en-US")
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Instagram 76.0.0.15.395 Android (22/5.1.1; 240dpi; 720x1280; OnePlus; A5010; A5010; intel; tr_TR; 138226758)")
	req.SetBody([]byte(fmt.Sprintf("username=%s&password=%s&device_id=android-72ca06792a4875e9&login_attempt_count=0", username, password)))
	println(username)
	println(password)

	err := client.Do(req, resp)
	if err != nil {
		log.Fatal(err)
	}

	fasthttp.ReleaseRequest(req)
	buffer := string(resp.Body())
	println(req)
	println(resp)
	if strings.Contains(string(buffer), "challenge_required") {
		fmt.Println("Suphe Dusen hesap: " + username)
		sorun = true
		file, err := os.OpenFile("suphe.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s:%s:%s\n", username, password, buffer))
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("yazildi! suphe api")
		}
	} else if strings.Contains(string(buffer), "logged_in_user") {

		deneme := resp.Header
		re := regexp.MustCompile(`sessionid=(.*); Domain=.instagram.com;`)
		match := re.FindStringSubmatch(deneme.String())

		fmt.Println(match[1])
		sessiond = match[1]

		error(err)
		fmt.Println("api giris basarili")
		return sessiond

	} else {
		fmt.Println("SORUN VAR: " + username)

		fmt.Println(string(buffer))
		sorun = true

		file, err := os.OpenFile("apisorun.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s:%s:%s\n", username, password, buffer))
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("yazildi! api sorun")
		}
	}

	return sessiond
}
func main() {
	fmt.Println("sayis", SayiGoster())
	/*
		var username string
		var password string
		var sessionid string
		var websession string
		fmt.Print("username:")
		fmt.Scanln(&username)

		fmt.Print("password:")
		fmt.Scanln(&password)

		LoginApi(username, password)

		sessionid = sessiond

		file, err := os.OpenFile("apilogin.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s:%s:%s\n", username, password, sessionid))
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("yazildi! api")
		}

		LoginWeb(username, password)
		websession = sessiondweb

		file2, err := os.OpenFile("weblogin.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file2.Close()
		file2.WriteString(fmt.Sprintf("%s:%s:%s\n", username, password, websession))
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("yazildi! web")
		}
	*/

	fileBytes, err := ioutil.ReadFile("verecegim.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var n int = 0
	var sessionid string
	var websession string
	var password string
	var user string
	var pass string
	sliceData := strings.Split(string(fileBytes), "\n")

	fmt.Print("Kac hesap var?")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		println("ip degis")
		fmt.Scanln()
		data := strings.Split(string(sliceData[i]), ":")
		user = string(data[0])
		pass = string(data[1])
		password = pass
		if strings.Contains(string(pass), "\r") {
			re := regexp.MustCompile(`(.*)\r`)
			passtemiz := re.FindStringSubmatch(pass)

			password = passtemiz[1]
		}

		LoginWeb(user, password)
		if sorun == true {
			sorun = false
			continue
		}
		websession = sessiondweb

		file2, err := os.OpenFile("weblogin.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file2.Close()
		file2.WriteString(fmt.Sprintf("%s:%s:%s\n", user, password, websession))
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("yazildi! web")
		}

		LoginApi(user, password)
		if sorun == true {
			sorun = false
			continue
		}
		sessionid = sessiond

		file, err := os.OpenFile("apilogin.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("%s:%s:%s\n", user, password, sessionid))
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("yazildi! api")
		}

	}
	os.Exit(0)
}
