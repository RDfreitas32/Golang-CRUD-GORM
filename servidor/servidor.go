package servidor

//Animais é a base de atributos de novos animais
type Animais struct {
	ID       uint32 `json:"id"`
	Nome     string `json:"nome"`
	Raca     string `json:"raca"`
	DonoNome string `json:"dononome"`
}
