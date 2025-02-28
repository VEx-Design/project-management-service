package request

type Project struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Flow            string `json:"flow"`
	TypesConfig     string `json:"typesConfig"`
	ConfigurationID uint   `json:"configurationId"`
}
