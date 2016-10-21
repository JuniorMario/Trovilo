package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Conf struct {
	Wordlist  string
	Content   []string
	url       string
	Set_Proxy bool
}

var Config Conf

func check_error(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}

func NewConf(filename string) Conf {
	var res []string
	dat, _ := ioutil.ReadFile(filename)
	res = strings.Split(string(dat), "\n")
	Config = Conf{Wordlist: filename, Content: res}
	return Config
}
func Write(url string) {
	res := []byte(url)
	_ = ioutil.WriteFile("output/res.txt", res, 0644)
}
func Banner() {
	banner, err := ioutil.ReadFile("src/banner.txt")
	check_error(err)
	fmt.Print(string(banner))
}
func Parser(arg string) []string {
	var param []string
	switch string(arg[0]) {
	case ":":
		param = strings.Split(arg, ":")[0:]
	case "/":
		os.Chdir(strings.Split(arg, "/")[1])
		param, _ = filepath.Glob("*.txt")
		os.Chdir("..")
	default:
		param = append(param, arg)
	}
	return param
}
func (Config *Conf) Set_Tor(opt bool) {
	Config.Set_Proxy = opt
}

func (Config *Conf) Request(url string) map[string]string {
	var resp *http.Response
	var err error
	if Config.Set_Proxy {
		httpClient := &http.Client{Transport: T.Tr}
		resp, err = httpClient.Get(url)
	} else {
		resp, err = http.Get(url)
	}
	check_error(err)
	res := map[string]string{
		"status_code": strconv.Itoa(resp.StatusCode),
		"url":         url,
	}
	return res
}
