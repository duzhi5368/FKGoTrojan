package main

import (
    "testing"
)

func TestGetLogicalDrives(t *testing.T) {
    s := GetLogicalDrives()
    if len(s) == 0 {
        t.Error("get drivers list is nil")
        return
    }
    t.Logf("drivers : %+v\n", s)
}

func TestReadable(t *testing.T) {
    r,err := Readable("C:")
    if err != nil {
        t.Error(err)
        return
    }
    t.Log(r)
}
