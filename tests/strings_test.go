package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplacer(t *testing.T) {
	r := strings.NewReplacer("#", "o")
	broken := "G#phers"
	fixed := r.Replace(broken)
	assert.Equal(t, "Gophers", fixed)
}

func TestStrings_Field(t *testing.T) {
	assert.Equal(t, []string{"go", "java", "python"}, strings.Fields("go java python"))
	assert.Equal(t, []string{"go", "java", "python"}, strings.Fields("\tgo  \f \u0085 \u00a0 java \n\rpython"))
	assert.Equal(t, []string{}, strings.Fields("\t \n\r    "))
}

func TestStrings_FieldFunc(t *testing.T) {
	assert.Equal(t, []string{"name", "age", "gender"}, strings.FieldsFunc("name,age,gender", func(r rune) bool {
		return r == ','
	}))
}

func TestStrings_Split(t *testing.T) {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))            // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a,b,c", "b"))            // ["a," ",c"]
	fmt.Printf("%q\n", strings.Split("Go社区欢迎你", ""))           // ["G" "o" "社" "区" "欢" "迎" "你"]
	fmt.Printf("%q\n", strings.Split("abc", "de"))             // ["abc"]
	fmt.Printf("%q\n", strings.SplitN("a,b,c,d", ",", 2))      // ["a" "b,c,d"]
	fmt.Printf("%q\n", strings.SplitN("a,b,c,d", ",", 3))      // ["a" "b" "c,d"]
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c,d", ","))     // ["a," "b," "c," "d"]
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d", ",", 2)) // ["a," "b,c,d"]
}

func TestStrings_Map(t *testing.T) {
	ret := strings.Map(func(r rune) rune {
		if r == ',' {
			return '|'
		} else {
			return r
		}
	}, "name,age,gender")
	assert.Equal(t, "name|age|gender", ret)
}
