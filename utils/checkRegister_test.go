package utils

import (
	"fmt"
	"testing"
)

func TestCheckCharacter(t *testing.T) {

	got := CheckCharacter("1tfs901t")

	fmt.Println(got)

}

func TestCheckUserInput(t *testing.T) {
	get := CheckUserInput("wdh", "1tfs901t")
	fmt.Println(get)
}
