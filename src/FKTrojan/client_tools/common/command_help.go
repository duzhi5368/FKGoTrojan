package common

type Parameter struct {
	LongFmt  string `json:"long_fmt"`
	ShortFmt string `json:"short_fmt"`
	Example  string `json:"example"`
	Desc     string `json:"desc"`
	Required bool   `json:"required"`
	Type     string `json:"type"`
}
type AppUsage struct {
	Name       string      `json:"name"`
	Version    string      `json:"version"`
	Desc       string      `json:"desc"`
	Parameters []Parameter `json"parameters"`
}
