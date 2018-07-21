/*
Author: FreeKnight
MaxMind GeoIP2关心的用户信息结构
*/
//------------------------------------------------------------
package components
//------------------------------------------------------------
type maxMind struct {
	City struct {
		     GeonameID int `json:"geoname_id"`
		     Names     struct {
				       En string `json:"en"`
				       Ru string `json:"ru"`
				       Zh string `json:"Zh"`
			       } `json:"names"`
	     } `json:"city"`
	Continent struct {
		     Code      string `json:"code"`
		     GeonameID int    `json:"geoname_id"`
		     Names     struct {
				       Ja   string `json:"ja"`
				       PtBR string `json:"pt-BR"`
				       Ru   string `json:"ru"`
				       ZhCN string `json:"zh-CN"`
				       ZhTW string `json:"zh-TW"`
				       De   string `json:"de"`
				       En   string `json:"en"`
				       Es   string `json:"es"`
				       Fr   string `json:"fr"`
			       } `json:"names"`
	     } `json:"continent"`
	Country struct {
		     IsoCode   string `json:"iso_code"`
		     GeonameID int    `json:"geoname_id"`
		     Names     struct {
				       ZhCN string `json:"zh-CN"`
				       ZhTW string `json:"zh-TW"`
				       De   string `json:"de"`
				       En   string `json:"en"`
				       Es   string `json:"es"`
				       Fr   string `json:"fr"`
				       Ja   string `json:"ja"`
				       PtBR string `json:"pt-BR"`
				       Ru   string `json:"ru"`
			       } `json:"names"`
	     } `json:"country"`
	Location struct {
		     AccuracyRadius int     `json:"accuracy_radius"`
		     Latitude       float64 `json:"latitude"`
		     Longitude      float64 `json:"longitude"`
		     MetroCode      int     `json:"metro_code"`
		     TimeZone       string  `json:"time_zone"`
	     } `json:"location"`
	Postal struct {
		     Code string `json:"code"`
	     } `json:"postal"`
	Subdivisions []struct {
		IsoCode   string `json:"iso_code"`
		GeonameID int    `json:"geoname_id"`
		Names     struct {
				  En   string `json:"en"`
				  Es   string `json:"es"`
				  Fr   string `json:"fr"`
				  Ja   string `json:"ja"`
				  PtBR string `json:"pt-BR"`
				  Ru   string `json:"ru"`
				  ZhCN string `json:"zh-CN"`
				  De   string `json:"de"`
			  } `json:"names"`
	} `json:"subdivisions"`
	Traits struct {
		     AutonomousSystemNumber       int    `json:"autonomous_system_number"`
		     AutonomousSystemOrganization string `json:"autonomous_system_organization"`
		     Isp                          string `json:"isp"`
		     Organization                 string `json:"organization"`
		     IPAddress                    string `json:"ip_address"`
	     } `json:"traits"`
}