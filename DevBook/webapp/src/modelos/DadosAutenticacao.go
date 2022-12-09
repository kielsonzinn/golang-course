package modelos

//DadosAutenticacao contém o id e o token usuário autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
