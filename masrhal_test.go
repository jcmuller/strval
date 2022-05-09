package strval_test

import (
	"embed"
	"fmt"
	"log"
	"testing"

	strval "github.com/jcmuller/strval"
	_assert "github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed testdata/*
	testdata embed.FS
)

func TestMarshal(t *testing.T) {
	var tests = []struct {
		given    string
		expected string
	}{
		{
			given:    "testdata/given_simple.yaml",
			expected: "testdata/expected_simple.yaml",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.given, func(t *testing.T) {
			assert := _assert.New(t)

			given, err := testdata.ReadFile(tt.given)
			assert.NoError(err)

			values := make(map[string]interface{})
			err = yaml.Unmarshal(given, &values)
			assert.NoError(err)

			expected, err := testdata.ReadFile(tt.expected)
			assert.NoError(err)

			actual, err := strval.Marshal(values)
			if assert.NoError(err) {
				assert.EqualValues(string(expected), string(actual)+"\n")
			}
		})
	}
}

func ExampleMarshal() {
	given, err := testdata.ReadFile("testdata/given_simple.yaml")
	if err != nil {
		log.Fatal(err)
	}

	values := make(map[string]interface{})
	err = yaml.Unmarshal(given, &values)
	if err != nil {
		log.Fatal(err)
	}

	actual, err := strval.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(actual))

	// Output:
	// bam.bar.bar: baz
	// bam.foo: oi
	// bam: baz
	// bar: 123.45
	// barbaz.bar: barr
	// barbaz.oi: vey
	// baz: bar
	// baz: baz
	// baz: false
	// baz: foo
	// baz: true
	// foo.bar.bar: bar!
	// foo.bar.baz: 123
}
