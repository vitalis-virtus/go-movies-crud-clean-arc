package entity

type Movie struct {
	Id       string `json: "id"`
	Name     string `json:"name"`
	Director string `json:"director"`
	Rating   string `json:"rating"`
}
