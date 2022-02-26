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
