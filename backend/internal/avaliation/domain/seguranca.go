package domain

type Seguranca struct {
	Alergias []Alergia `json:"alergias"`
}

type Alergia struct {
	TipoAlergia        SelectField `json:"tipo_alergia"` // FIXED: Object in JSON
	QuaisAlergias      string      `json:"quais_alergias"`
	Precaucao          SelectField `json:"precaucao"` // FIXED: Object in JSON
	CuidadosPaliativos string      `json:"cuidados_paliativos"`
}
