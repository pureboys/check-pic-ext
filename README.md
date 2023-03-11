# checkpicext - check picture ext tool

## Install

``` go get https://github.com/pureboys/check-pic-ext```


- Code Sample

```
package main

import (
	"fmt"
	checkpicext "github.com/pureboys/check-pic-ext"
	"os"
)

func main() {
	file, err := os.ReadFile("testdata/image1.heic")
	if err != nil {
		panic(err)
	}
	format := checkpicext.GetImageFormat(file)
	fmt.Println(format)
}

```
