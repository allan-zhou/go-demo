package main

import (
	"io/ioutil"
	"fmt"
)

type Page struct {
	Title   string
	Content []byte
}

func (p *Page) Save() (err error) {
	return ioutil.WriteFile(p.Title, p.Content, 0666)
	
}

func (p *Page) Load(title string) (err error) {
	p.Title = title
	content, err := ioutil.ReadFile(title)
	p.Content = content
	return err
}

func main() {
	page := &Page{
		"blog.md",
		[]byte("# header1 \n## header2 \n### header3"),
	}

	page.Save()

	var newpage Page
	newpage.Load("blog.md")
	fmt.Println(string(newpage.Content))
}
