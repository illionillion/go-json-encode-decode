package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Post struct {
	Title         string   `json:"post_title"`
	Tags          []string `json:"tags"`
	User          User     `json:"user"`
	Memo          *string  `json:"memo"`
	MemoOmitEmpty *string  `json:"memo_omit_empty,omitempty"`
	privateMemo   *string
}

func main() {

	// 配列の保存
	s :=[]string{"a","b","c"}
	sj, _ := json.Marshal(s)
	fmt.Printf("sj=%+v\n\n", sj)
	fmt.Printf("string(sj)=%+v\n\n", string(sj))

	if err := os.WriteFile("slice.json", sj , 0755); err != nil {
		log.Fatal(err)
	}

	// オブジェクトの保存
	m := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": []int{4,5,6},
	}

	mj, _ := json.Marshal(m)
	fmt.Printf("mj=%+v\n\n", mj)
	fmt.Printf("string(mj)=%+v\n\n", string(mj))

	if err := os.WriteFile("obj.json", mj , 0755); err != nil {
		log.Fatal(err)
	}

	// こうやって型をつけることも可能
	p := Post{
		Title: "Hello",
		Tags: []string{"aaa", "bbb"},
		User: User{
			Name: "jhon",
			Age: 21,
		},
	}

	fmt.Printf("p=%+v\n\n", p)

	pj, _ := json.Marshal(p)
	fmt.Printf("pj=%+v\n\n", pj)
	fmt.Printf("string(pj)=%+v\n\n", string(pj))

	if err := os.WriteFile("post.json", pj , 0755); err != nil {
		log.Fatal(err)
	}

	// 読み込み
	bytes, err := os.ReadFile("post.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bytes=%+v\n\n", bytes)
	fmt.Printf("string(bytes)=%+v\n\n", string(bytes))

	var post Post
	if err := json.Unmarshal(bytes, &post); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("post=%+v\n\n", post)

}