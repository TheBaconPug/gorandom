# gorandom

## A go package for random things

# Usage

## Functions

### RandStr()

``gorandom.RandStr(number of characters)``

### RandInt()

``gorandom.RandInt(number of characters)``

## Example Program

```go
package main

import (
  "github.com/TheBaconPug/gorandom"
)

func main() {
  println(gorandom.RandStr(10))
}
```
