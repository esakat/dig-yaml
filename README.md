# Easy to get value of yaml

Introduction
------------

The dig-yaml package enables to get value of yaml

Installation and usage
----------------------

To install it, run:

    go get github.com/esakat/dig-yaml
    
Example
-------

```Go
package main

import (
    "fmt"
    "log"
    "gopkg.in/yaml.v2"
    "os"
    "github.com/esakat/dig-yaml"
)

func main() {
    f, err := os.Open("testfile.yml")
    defer f.Close()
    if err != nil {
        log.Fatal(err)
    }

    dec := yaml.NewDecoder(f)

    var y map[interface{}]interface{}
    dec.Decode(&y)
    
    // it's like y["hoge"].([]interface{})[0].(map[interface{}]interface{})["key1"]
    key1, err := dig_yaml.DigYaml(y, "hoge", 0, "key1")
    if err != nil {
        log.Fatal(err)
    }

    // Output: test1
    fmt.Println(key1)

}
```

testfile.yml

```yaml
foo: hoge
1: false
hoge:
  - key1: test1
    key2: 1
  - key1: test2
    key2: 2
```
