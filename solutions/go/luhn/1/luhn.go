package luhn

import(
    "strings"
)

func Valid(id string) bool {
	var formatId string
    soma := 0
    flag := false
    formatId = strings.ReplaceAll(id, " ", "")
	
    if len(formatId) <= 1 {
		return false
	}

	digitos := make([]int, len(formatId))

	for i, char := range formatId {
		if char < '0' || char > '9' {
			return false // Caractere inválido
		}
		digitos[i] = int(char - '0')
	}


	// Percorre da direita para esquerda
    for i := len(digitos) - 1; i >= 0; i-- {
        digito := digitos[i]
		if flag {
			digito *= 2
			if digito > 9 {
				digito -= 9
			}
		}
		
		soma += digito
		flag = !flag
	}
	
    return soma%10 == 0
}
