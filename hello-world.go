package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Understanding Interfaces in Go",
		Author: "Sammy Shark",
	}
	Print(a)

	b := Book{
		Title:  "All About Go",
		Author: "Jenny Dolphin",
		Pages:  25,
	}
	Print(b)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}

// type Article struct {
// 	Title  string
// 	Author string
// }

// func (a Article) String() string {
// 	return fmt.Sprintf("Article Name %q, Writer: %q", a.Title, a.Author)

// }
// func main() {
// 	a := Article{
// 		Title:  "Programming with C",
// 		Author: "Dennis Riche",
// 	}
// 	b := Article{
// 		Title: "Hello Programmer",
// 	}

// 	fmt.Println("Book Title: ", a.Title, "Author Name: ", a.Author)
// 	fmt.Println("Book Title: ", b.Title, b.Author)

// s := "postgres://user:pass@host.com:5432/path?k=v#f"

// u, err := url.Parse(s)
// if err != nil {
// 	panic(err)
// }

// fmt.Println("u.Scheme", u.Scheme)

// fmt.Println("u.User", u.User)
// fmt.Println("u.User.Username()", u.User.Username())
// p, _ := u.User.Password()
// fmt.Println("u.User.Password()", p)

// fmt.Println("u.Host", u.Host)
// host, port, _ := net.SplitHostPort(u.Host)
// fmt.Println("host", host)
// fmt.Println("port", port)

// fmt.Println("u.Path", u.Path)
// fmt.Println("u.Fragment", u.Fragment)

// fmt.Println("u.RawQuery", u.RawQuery)
// m, _ := url.ParseQuery(u.RawQuery)
// fmt.Println(m)
// fmt.Println(m["k"][0])
// }
