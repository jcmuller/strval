# strval

Simple strval marshaller

## Example

```golang
// d is a []byte
var data map[string]interface{}
if e := yaml.Unmarshal(d, &data); e != nil {
  log.Fatal(e)
}

out, e := strval.Marshal(data)
if e != nil {
  log.Fatal(e)
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
