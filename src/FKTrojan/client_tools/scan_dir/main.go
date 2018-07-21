package main

import (
	. "FKTrojan/client_tools/common"
	common2 "FKTrojan/common"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func JsonPrintHelp() {

	app := AppUsage{
		Name:    ExeBaseName(),
		Version: "1.0.0",
		Desc:    "this is for scan custom dir, list sub dir and files",
		Parameters: []Parameter{
			{
				LongFmt:  "-dir",
				ShortFmt: "-dir",
				Example:  "c:\\windows\\system32",
				Desc:     "scan dir",
				Required: true,
				Type:     "string",
			},
			{
				LongFmt:  "-depth",
				ShortFmt: "-depth",
				Example:  "2",
				Desc:     "depth of current dir",
				Required: false,
				Type:     "int",
			},
		},
	}
	var b []byte
	if len(os.Args) >= 2 && os.Args[1] == "--HELP" {
		b, _ = json.Marshal(app)
	} else {
		b, _ = json.MarshalIndent(app, " ", " ")
	}
	fmt.Println(string(b))
}
func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			JsonPrintHelp()
			os.Exit(0)
		}
	}
	dir := flag.String("dir", "", "`directory` that you want list sub dir and files")
	depth := flag.Uint("depth", common2.DEPTH_ALL, "`depth` that you want list, default is all")
	flag.Parse()
	if *dir == "" {
		flag.Usage()
		os.Exit(1)
	}
	//fmt.Println(*dir, *depth)
	items, err := Scan(*dir, *depth)
	if err != nil {
		fmt.Printf("scan(\"%s\") error %v\n", *dir, err)
		os.Exit(1)
	}
	b, err := json.MarshalIndent(items, " ", " ")
	if err != nil {
		fmt.Printf("json.MarshalIndent error %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}
