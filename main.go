package main

import (
	"flag"
	"fmt"
	"utils"
)

func brute(wordlist []string, util utils.Conf) {
	for _, endswith := range wordlist {
		res := util.Request("http://sitefuleiro.webnode.com.br/" + endswith)
		if res["status_code"] == "200" {
			utils.Write(res["url"])
		}
	}
}

func parse(Wordlist string, Tor bool) {
	var conf utils.Conf
	param := utils.Parser(Wordlist)
	if Tor {
		tor := utils.NewTor()
		tor.Check_IP()
		conf.Set_Tor(true)
	}
	for _, wordlist := range param {
		conf := utils.NewConf(wordlist)
		go brute(conf.Content, conf)
	}
}

func main() {
	utils.Banner()
	var Wordlist = flag.String("Wordlist", "0", "Wordlist Path.")
	var Tor = flag.Bool("tor", false, "Seta o Tor")
	flag.Parse()
	parse(*Wordlist, *Tor)
	var input string
	fmt.Scanln(&input)
}
