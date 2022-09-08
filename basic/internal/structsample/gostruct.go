package structsample

import (
	"fmt"
	"unsafe"
)

type newUint8 uint8

type goUint8 uint8

type aUint = uint //type alias

type Book struct {
	Name   string
	Author Person
	Publisher
}

type Person struct {
	Name string
	Age  uint8
}

type Publisher struct {
	Name string
	Address
}

type Address struct {
	Streetname string
	Postalcode string
}

func NewBook(bookname string, authorname string) *Book {
	book := Book{
		Name: bookname,
		Author: Person{
			Name: authorname,
			Age:  30,
		},
		Publisher: Publisher{
			Name: "Sunshine Publisher",
			Address: Address{
				Streetname: "Central park",
				Postalcode: "074184",
			},
		},
	}
	return &book
}

func CompareStructType() {
	var t1 newUint8 = 1
	t2 := 6
	t1 = newUint8(t2)
	fmt.Printf("both t1 [%d] and t2 [%d] underlying type are uint8, hence can cast type between t1 and t2", t1, t2)
	aUint := 10
	aUint += 3
	fmt.Println("t3 is alias of uint, implictly cast when operating with uint")
}

func CreateStruct() {
	book := Book{
		Name: "To Whom the Bell Tolls",
		Author: Person{
			Name: "Hemingwei",
			Age:  48,
		},
		Publisher: Publisher{
			Name: "Public Book",
			Address: Address{
				Streetname: "East Lane",
				Postalcode: "014337",
			},
		},
	}
	booksize := unsafe.Sizeof(book)
	bookNameOffset := unsafe.Offsetof(book.Name)
	authorOffset := unsafe.Sizeof(book.Author)
	authorsize := unsafe.Sizeof(book.Author)
	authorNameOffset := unsafe.Offsetof(book.Author.Name)
	publisherSize := unsafe.Sizeof(book.Publisher)
	publisherOffset := unsafe.Offsetof(book.Publisher)
	fmt.Printf("book size [%d], offset of bookname[%d], offset of author[%d], size of author[%d], offset of author name [%d], size of publier[%d], offset of publisher [%d]\n",
		booksize, bookNameOffset, authorOffset, authorsize, authorNameOffset, publisherSize, publisherOffset)
}

func CreateStructWithConstructor(bookname string, authorname string) {
	book := *NewBook(bookname, authorname)
	fmt.Printf("%#v", book)
}
