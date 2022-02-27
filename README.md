# strval

[![Go Reference](https://pkg.go.dev/badge/github.com/jcmuller/strval.svg)](https://pkg.go.dev/github.com/jcmuller/strval)
[![CircleCI](https://circleci.com/gh/jcmuller/strval/tree/main.svg?style=svg)](https://circleci.com/gh/jcmuller/strval/tree/main)

Simple strval marshaller. It sorts the output, too. Arrays are
unmarshalled in such a way that they are multiple keys with the same
name, so the output is not valid YAML.

## Example

```golang
//go:embed testdata/*
var testdata embed.FS

func main() {
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
}
```

```
$ go install github.com/jcmuller/strval/cmd/strval@latest
```

```
$ cat testdata/given_simple.yaml
---
foo:
  bar:
    baz: 123
    bar: bar!

bar: 123

baz:
  - foo
  - bar
  - baz

barbaz:
  oi: vey
  bar: barr

bam:
  - foo: oi
  - bar:
      bar: baz
  - baz
```

```
$ go run .
bam.bar.bar: baz
bam.foo: oi
bam: baz
bar: 123
barbaz.bar: barr
barbaz.oi: vey
baz: bar
baz: baz
baz: foo
foo.bar.bar: bar!
foo.bar.baz: 123
```
