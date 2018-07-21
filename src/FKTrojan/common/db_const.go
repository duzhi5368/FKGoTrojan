package common

type CommandStatus int

const (
	STATUS_NEW CommandStatus = iota
	STATUS_DOING
	STATUS_DONE
	STATUS_ERR
)
