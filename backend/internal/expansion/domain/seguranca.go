package domain

type Seguranca struct {
	Alergias []Alergia `json:"alergias"`
}

type Alergia struct {
	TipoAlergia        string `json:"tipo_alergia"` // FIXED: Object in JSON
	QuaisAlergias      string `json:"quais_alergias"`
	Precaucao          string `json:"precaucao"` // FIXED: Object in JSON
	CuidadosPaliativos string `json:"cuidados_paliativos"`
}
