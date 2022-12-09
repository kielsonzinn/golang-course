package endereco

import "strings"

//TipoDeEndereco verifica se um endereço tem um tpo valido e o retorna
func TipoDeEndereco(endereco string) string {

	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecosEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalava := strings.Split(enderecosEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalava {
			enderecoTemUmTipoValido = true
			break
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalava)
	}

	return "invalid type"

}
