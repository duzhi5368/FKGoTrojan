package main

type CmdResult []map[string]string
type Cmd interface {
	CheckArgs() error
	Before() error
	After() error
	Execute() (CmdResult, error)
	String() string
	//Help()
}
