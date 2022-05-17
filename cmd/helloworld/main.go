package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// シャドーイング
	x := 1
	{
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)

	var i int = 134
	// var f float64 = float64(i)
	// var i64 int64 = int64(i)
	var b bool = i != 0
	fmt.Println(b)

	in := 12345
	s := strconv.FormatInt(int64(in), 10)
	fmt.Println(s)

	hoge := 333
	var fuga *int
	fmt.Println(fuga)
	fuga = &hoge
	fmt.Println(fuga)
	fmt.Println(*fuga)

	// array
	var nums [3]int = [3]int{1, 2}
	fmt.Println(nums)
	fmt.Println(nums[1])
	// fmt.Println(nums[3]) invalid argument: array index 3 out of bounds [0:3]
	fmt.Println(len(nums))

	// slice
	var nums1 []int
	nums2 := []int{1, 2, 3, 4, 5}
	nums3 := nums[0:1]
	nums4 := nums[1:3]
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(nums3)
	fmt.Println(nums4)
	nums2[1] = 10000
	nums5 := append(nums2, 100)
	fmt.Println(nums2)
	fmt.Println(nums5)

	// fmt.Println(nums5[10]) panic: runtime error: index out of range [6] with length 6

	// map
	hs := map[int]string{
		200: "OK",
		400: "NG",
	}
	fmt.Println(hs)
	fmt.Println(hs[200])
	// fmt.Println(hs["OK"]) // cannot use "OK" (untyped string constant) as int value in map index
	fmt.Println(hs[300])

	authors := make(map[string][]string)
	authors["GO"] = []string{"山田太郎", "山田花子", "山田二郎"}
	fmt.Println(authors)
	goAuthors, ok := authors["GO"]
	fmt.Println(goAuthors)
	fmt.Println(ok)

	// if
	statusCode := 300
	if statusCode == 200 {
		fmt.Println("no error")
	} else if statusCode < 500 {
		fmt.Println("client error")
	} else {
		fmt.Println("serve error")
	}

	cache := map[string]string{"ok": "キャッシュ"}
	if result, ok := cache["ok"]; ok {
		fmt.Println("has cache value", result)
	}

	// for
	colors := []string{"red", "green", "blue", "white"}
	for i, s := range colors {
		if strings.HasPrefix(s, "gre") {
			fmt.Println("gre", s)
			continue
		}
		if strings.HasSuffix(s, "ue") {
			fmt.Println("bre", s)
			break
		}
		fmt.Println(i, s)
	}

	counter := 0
	for counter < 10 {
		fmt.Println("counter is", counter)
		counter++
	}

	for {
		fmt.Println("counter is", counter)
		if counter == 100 {
			break
		}
		counter++
	}

	// switch
	nu := 10
	switch {
	case nu < 10:
		fmt.Println("under 10")
		fallthrough
	case nu < 100:
		fmt.Println("under 100")
		fallthrough
	case nu < 1000:
		fmt.Println("under 1000")
	}

	fmt.Println(calc(2)(3))

	var b1 Book
	fmt.Println(b1)
	b2 := Book{Title: "hogehoge"}
	fmt.Println(b2)
	b2.Author = "夏目漱石"
	fmt.Println(b2)
	b3 := &Book{Title: "ポインターである"}
	fmt.Println(b3)
	b3.Author = "200"
	fmt.Println(b3)
	fmt.Println(*b3)

	f, err := os.Open("book.json")
	if err != nil {
		log.Fatal("file open error: ", err)
	}
	d := json.NewDecoder(f)
	var b4 BookJson
	d.Decode(&b4)
	fmt.Println(b4)
}

func calc(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

type Book struct {
	Title      string
	Author     string
	Publisher  string
	ReleasedAt time.Time
	ISBN       string
}

type BookJson struct {
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	ISBN       string    `json:"isbn"`
	ReleasedAt time.Time `json:"release_at"`
}
