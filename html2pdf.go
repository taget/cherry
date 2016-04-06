package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("require provide file name!")
		return
	}
	html_file_name := os.Args[1]
	file, err := os.Open(html_file_name) // For read access.

	if err != nil {
		fmt.Println(err)
		return
	}

	doc, err := html.Parse(file)
	if err != nil {
		fmt.Println(err)
		return
	}


	var f func(*html.Node) (*html.Node)
	f = func(n *html.Node) (*html.Node){

		if n.Type == html.ElementNode && n.Data == "li" {
			//fmt.Println("elementnode")
			//fmt.Println(n.Data)
			//fmt.Println(n.Attr)
            //n.RemoveChild(n)
		}
		if n.Type == html.DocumentNode {
			//fmt.Println("Documentnode")
			//fmt.Println(n.Data)
		}
		if n.Type == html.TextNode {
			//fmt.Println("TextNode")
			//fmt.Println(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {

//            fmt.Println(c.Data)
            summary_div := html.Attribute{"", "class", "book-summary"}
            if c.Data == "div" {
                fmt.Println("aaaaaaaaaaaaaaaaaaaaa")
                fmt.Println(c.Type)
                fmt.Println(c.Data)
                fmt.Println(c.Namespace)
                fmt.Println(c.Attr)
                fmt.Println("aaaaaaaaaaaaaaaaaaaaa")
                if c.Attr[0] == summary_div {
                    fmt.Println("jimmy")
                    // this is book-summary div, I need to delete it
                    n.RemoveChild(c)
                }
            } else {
//                fmt.Println("----------------")
            }
            c = f(c)
		}
        return n
	}

    // remove book-summary div from doc
    new_doc := f(doc)

    file.Close()
    // write it back
    fw, err := os.Create(html_file_name + "_new")

    err = html.Render(fw, new_doc)

    defer fw.Close()
}
