package main

import "ByGolang/ByPackage"

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main1() {

	//a := app()
	//b := app()
	//fmt.Println(a("go"))
	//fmt.Println(a("go 1"))
	//fmt.Println(a("go 2"))
	//fmt.Println(b("all"))
	ByPackage.Init()
	ByPackage.Package()
}
