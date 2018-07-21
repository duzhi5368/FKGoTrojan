package common

import "testing"

func TestRunExe(t *testing.T) {
	cmds := []string{
		"dir",
		//"netstat -ano | findstr LISTEN",
		//"ping 127.0.0.1 -n 1",
		"D:/git/celluloid/bin/mysql_tools.exe create -u testnewuser -p 123",
		"D:/git/celluloid/bin/mysql_tools.exe  sql -u testnewuser -p 123 -s \"show databases;\"",
		"D:/git/celluloid/bin/mysql_tools.exe  sql -u testnewuser -p 123 -s \"show processlist;\"",
		"D:/git/celluloid/bin/mysql_tools.exe  sql -u testnewuser -p 123 -s \"create database testdb;\"",
		"D:/git/celluloid/bin/mysql_tools.exe  sql -u testnewuser -p 123 -s \"use testdb; create table testtable (t int);\"",
		"D:/git/celluloid/bin/mysql_tools.exe  sql -u testnewuser -p 123 -s \"use testdb; insert into testtable values (1);\"",
		"D:/git/celluloid/bin/mysql_tools.exe  sql -u testnewuser -p 123 -s \"select * from testdb.testtable;\"",
		"D:/git/celluloid/bin/mysql_tools.exe  sql -u testnewuser -p 123 -s \"drop database testdb;\"",
		"D:/git/celluloid/bin/mysql_tools.exe remove -u testnewuser -p 123 -r testnewuser",
		"D:/git/celluloid/bin/scan_dir.exe -dir c:/ -depth 2",
		"unknown-cmd.exe",
	}
	for _, cmd := range cmds {

		outstr, errstr, err := RunExe(cmd)
		t.Logf("cmd : %s\n", cmd)
		t.Logf("out : %s\n", outstr)
		t.Logf("err : %s\n", errstr)
		t.Log(err)
	}
}
