package random_test

import (
	"fmt"
	"testing"

	"github.com/kaschula/random"
)

func TestARandomStringIsGenerated(t *testing.T) {
	length := 20
	randomOne, errOne := random.GenerateString(length)
	randomTwo, errTwo := random.GenerateString(length)

	if randomOne == randomTwo || errOne != nil || errTwo != nil {
		t.Fatal("Values should not match")
	}
}

func TestARandomStringIsGeneratedByGenerator(t *testing.T) {
	generator := random.NewRandomStringGenerator(20)

	randomOne, errOne := generator.Generate()
	randomTwo, errTwo := generator.Generate()

	if randomOne == randomTwo || errOne != nil || errTwo != nil {
		t.Fatal("Values should not match")
	}
}

func TestARandomStringCanBeGeneratedAtASpecifiedLength(t *testing.T) {
	tests := []int{2, 4, 8, 10, 12, 14, 20, 30, 40, 80}

	for _, desiredLength := range tests {
		random, err := random.GenerateString(desiredLength)

		if err != nil {
			t.Fatal("Error")
		}

		fmt.Println(len(random))

		if len(random) != desiredLength {
			t.Fatal("Length should be")
		}
	}
}

func TestARandomStringCanBeGeneratedWithGenerator(t *testing.T) {
	tests := []int{2, 4, 8, 10, 12, 14, 20, 30, 40, 80}

	for _, desiredLength := range tests {
		generator := random.NewRandomStringGenerator(desiredLength)
		random, err := generator.Generate()

		if err != nil {
			t.Fatal("Error")
		}

		fmt.Println(len(random))

		if len(random) != desiredLength {
			t.Fatal("Length should be")
		}
	}
}
