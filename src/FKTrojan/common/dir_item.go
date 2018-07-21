package common

type ItemType int
type Item struct {
	FullPath string   `json:"full_path"`
	ItemType ItemType `json:"item_type"`
}

const (
	DIR_ITEM ItemType = iota
	FILE_ITEM
	DEPTH_ALL uint = 0xffffffff
)
