package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	t.Run("Should sum the two numbers", func(t *testing.T) {
		//arrange
		num1 := 1
		num2 := 2

		resultadoEsperado := 3

		//act
		resultado := soma(num1, num2)

		//assert
		assert.Equal(t, resultadoEsperado, resultado)
	})
}

func TestValidaNome(t *testing.T) {
	t.Run("Should return error when name is empty", func(t *testing.T) {
		name := ""

		expectedResult := "nome não preenchido"

		err := validaNome(name)

		assert.Error(t, err)
		assert.Equal(t, expectedResult, err.Error())
	})
}
