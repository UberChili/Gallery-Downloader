package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Post struct {
	No  int   `json:"no"`
	Tim int64 `json:"tim"`
	Name string `json:"name"`
	Com string `json:"com"`
	Filename string `json:"filename"`
	Ext string `json:"ext"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

var brd = flag.String("b", "", "Board")
var tnum = flag.String("n", "", "Thread number")

func main() {
	// url := prep_url("g", "99087972")
	flag.Parse()
	url := prep_url(*brd, *tnum)
	fmt.Println("Attempting to get response...")
	// json := get_json(url)
	parse_json(get_json(url))

}

func prep_url(board string, number string) string {
	url := "https://a.4cdn.org/board/thread/thread_num.json"
	return strings.Replace(strings.Replace(url, "board", board, 1), "thread_num", number, 1)
}

func get_json(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get_response: %v\n", err)
		os.Exit(1)
	}
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "get_response: reading %s, %v\n", url, err)
		os.Exit(1)
	}
	// parse_json(b)
	return b
}

func parse_json(b[] byte) {
	var posts Posts
	err := json.Unmarshal(b, &posts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unmarshal: %s\n", b)
	}

	for _, post := range posts.Posts {
		if post.Filename != "" {
			fmt.Println(get_full_filename(post))
		}
	}
}

func get_full_filename(post Post) string {
	return post.Filename + post.Ext
}
