package main

import (
	. "FKTrojan/client_tools/common"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	app := AppUsage{
		Name:    ExeBaseName(),
		Version: "3.1.4",
		Desc:    "this is a standard help format for client",
		Parameters: []Parameter{
			{
				LongFmt:  "--dir",
				ShortFmt: "-d",
				Example:  "c:\\windows\\system32",
				Desc:     "check dir for update something",
				Required: true,
				Type:     "string",
			},
			{
				LongFmt:  "--verbose",
				ShortFmt: "-v",
				Example:  "8",
				Desc:     "debug log level",
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
