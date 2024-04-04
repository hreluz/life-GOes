package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "http://google.com",
		"AWS":    "https://aws.com",
	}

	fmt.Println(websites["google"])

	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "AWS")
	fmt.Println(websites)
}
