# go-helloWorld
meh

```
package main

import "fmt"
import (
	"strings"
	"sort"
	"strconv"
)

func main (){
	foo := "World"
	bar := "Hello"
	fb := append([]string{}, bar, foo, foo, bar)
	foobar := strings.Join(fb, " ")

	fmt.Println("Sorted: " + strconv.FormatBool(sort.StringsAreSorted(fb)))
	fmt.Println(foobar)

	sort.Strings(fb)
	foobar = strings.Join(fb, " ")
	fmt.Println("Sorted: " + strconv.FormatBool(sort.StringsAreSorted(fb)))
	fmt.Println(foobar)

	fmt.Println("First: " + fb[0])
	fmt.Println("Last: " + fb[len(fb) - 1])

	//# long form
	//var p int64 = 8080
	//port := strconv.FormatInt(p,10)

	//# short form
	p := 8080
	port := strconv.Itoa(p)

	//# mixed form
	//p := 8080
	//port := strconv.FormatInt(int64(p),10)


	fb = append(fb, foobar)
	endpoint := strings.Join(fb, "/")
	url := "https://" + foo + ":" + bar + "@" + "localhost" + ":" + port + "/" + endpoint

	url = strings.Replace(url, bar + " ", bar + "?foobar=" , 1)
	url = strings.Replace(url, " ", "%20", -1)
	fmt.Println(url)
}

```
