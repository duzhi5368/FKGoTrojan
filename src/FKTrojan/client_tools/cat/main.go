package main

import (
	. "FKTrojan/client_tools/common"
	"FKTrojan/stream_utils"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func JsonPrintHelp() {

	app := AppUsage{
		Name:    ExeBaseName(),
		Version: "1.0.0",
		Desc:    "cat file tools",
		Parameters: []Parameter{
			{
				LongFmt:  "-f",
				ShortFmt: "-f",
				Example:  "c:/windows/system32/drivers/etc",
				Desc:     "filename",
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

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			JsonPrintHelp()
			os.Exit(0)
		}
	}
	fileName := flag.String("f", "", "")
	flag.Parse()
	err := stream_utils.FileToStream(*fileName, os.Stdout, false)
	if err != nil {
		fmt.Printf("error accurred : %v\n", err)
	}
}
