# strval

[![Go Reference](https://pkg.go.dev/badge/git.sr.ht/~jcmuller/strval.svg)](https://pkg.go.dev/git.sr.ht/~jcmuller/strval)
[![MIT](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![builds.sr.ht status](https://builds.sr.ht/~jcmuller/strval.svg)](https://builds.sr.ht/~jcmuller/strval)

Simple strval marshaller. It sorts the output, too. Arrays are
unmarshalled in such a way that they are multiple keys with the same
name, so the output is not valid YAML.

## Examples

### As a library

```shell
$ go get git.sr.ht/~jcmuller/strval/cmd/strval
```

```golang
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"git.sr.ht/~jcmuller/strval"
	"gopkg.in/yaml.v3"
)

func main() {
	d, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]any
	if e := yaml.Unmarshal(d, &data); e != nil {
		log.Fatal(e)
	}

	out, e := strval.Marshal(data)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Printf("%s\n", out)
}
```

```shell
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

### As binary

```shell
$ go install git.sr.ht/~jcmuller/strval/cmd/strval@latest
$ strval <testdata/given_simple.yaml
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

### Test data

```yaml
---
# testdata/given_simple.yaml
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
