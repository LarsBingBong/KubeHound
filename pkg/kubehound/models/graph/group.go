package graph

type Group struct {
	StoreID      string `json:"storeID" mapstructure:"storeID"`
	App          string `json:"app" mapstructure:"app"`
	Team         string `json:"team" mapstructure:"team"`
	Service      string `json:"service" mapstructure:"service"`
	IsNamespaced bool   `json:"isNamespaced" mapstructure:"isNamespaced"`
	Namespace    string `json:"namespace" mapstructure:"namespace"`
	Name         string `json:"name" mapstructure:"name"`
	Critical     bool   `json:"critical" mapstructure:"critical"`
}
