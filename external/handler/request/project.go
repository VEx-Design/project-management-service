package request

type Project struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Flow            string `json:"flow"`
	ConfigurationID uint   `json:"configurationId"`
}
