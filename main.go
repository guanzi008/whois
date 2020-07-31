package main

import (
	"fmt"
	"os"
	"whois/www"
)

func main() {
	if len(os.Args) > 1 {
		result, err := www.GetWhois(os.Args[1])
		if err != nil {
			fmt.Println("Error in whois lookup : %v ", err)
		} else {
			fmt.Println(result)
			fmt.Printf("Nameservers: %v \n", www.ParseNameServers(result))
		}
	} else {
		fmt.Println("请输入域名")
	}
}
