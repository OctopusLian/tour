package main

import (
	"flag"
	"log"
)

var name string

func main() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "Go语言", "帮助信息")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "PHP语言", "帮助信息")

	args := flag.Args()
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = goCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)

}

/*
# go run main.go go -n=neo
flag provided but not defined: -n
Usage of go:
  -name string
    	帮助信息 (default "Go语言")
exit status 2
# go run main.go go -name=neo
2020/07/14 17:15:01 name: neo

*/
