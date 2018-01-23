package models

//Usuarios representa a tabela usuários na base de dados. Struct sempre iniciada com letra maiúscula. O campo é representado por letras minúsculas no banco de dados e no json
type Usuarios struct {
	ID    int    `db: "id" json:"id"`
	Nome  string `db:"nome" json"nome"`
	Email string `db:"email" json"email"`
}
