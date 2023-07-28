// pode ser usado com adresses_test
package adresses

import "testing"

type testScenario struct {
	insertedAdress string
	expectedReturn string
}

// para testar todos os diretórios, usa-se o comando "go test ./..."
// para rodar os testes de maneira passo a passo, parecido com um debug, "go test -v"
// para saber quantos porcento de cobertura de testes a função ou funções de teste estão oferecendo: "go test --cover"
func TestIsValidType(t *testing.T) {
	// para rodar as funções de teste em paralelo
	t.Parallel()

	// declarando um slice de testScenario
	testScenarios := []testScenario{
		{"street ABC", "street"},
		{"highway of immigrants", "highway"},
		{"Time Square", "Invalid type"},
		{"road whatever", "road"},
		{"STREET OF DUMBS", "STREET"},
		{"", "Invalid type"},
	}

	for _, scenario := range testScenarios {
		adressTypeReceived := IsValidType(scenario.insertedAdress)
		if adressTypeReceived != scenario.expectedReturn {
			t.Errorf(
				"The type of adress received: %s is not the adress type expected: %s", adressTypeReceived, scenario.expectedReturn,
			)
		}
	}
}

func WhateverTest(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Wrong test!")
	}
}

// para gerar relatórios de cobertura de testes é usado "go test --coverprofile" nomedoarquivo.txt
// para ler esse arquivo adequadamento usa-se: "go tool cover --func=nomedoarquivo.txt" ou go tool cover --html=nomedoarquivo.txt
