package main

import "testing"

func Test_Get(t *testing.T) {

	total := sum(30, 30)

	if total != 60 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 900)
	}

}
