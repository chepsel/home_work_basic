package main

import (
	"fmt"

	"github.com/chepsel/home_work_basic/hw09_serialize/books"
	"github.com/chepsel/home_work_basic/hw09_serialize/protoc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ProtocBook struct {
	Books []protoc.Book
}

type Marshaller interface {
	MarshalJSON() ([]byte, error)
	MarshalProto() ([]byte, error)
}

type Unmarshaller interface {
	UnmarshalJSON([]byte) error
	UnmarshalProto(marshaled []byte) error
}

type Message interface {
	String() string
	ProtoReflect() protoreflect.Message
	GetID() string
	GetRate() float32
	GetYear() uint32
	GetSize() uint32
	GetTitle() string
	GetAuthor() string
}

func MarshalProto(protoBook *protoc.BooksSlice) ([]byte, error) {
	var marshaled []byte
	var err error
	clonedBook := proto.Clone(protoBook).(*protoc.BooksSlice)
	if protoBook.ProtoReflect().IsValid() {
		marshaled, err = proto.Marshal(clonedBook)
		if err != nil {
			return nil, err
		}
	}
	return marshaled, nil
}

func UnmarshalProto(marshaled []byte) (*protoc.BooksSlice, error) {
	protoBook := &protoc.BooksSlice{}
	err := proto.Unmarshal(marshaled, protoBook)
	if err != nil {
		result := proto.Clone(protoBook).(*protoc.BooksSlice)
		return result, err
	}
	return protoBook, nil
}

func ToProtobufStructure(bookSlice *books.JSONBook) *protoc.BooksSlice {
	protoBooks := &protoc.BooksSlice{}
	protoBooks.Books = make([]*protoc.Book, len(bookSlice.Books))
	for i, element := range bookSlice.Books {
		protoBooks.Books[i] = &protoc.Book{
			ID:     element.ID,
			Title:  element.Title,
			Author: element.Author,
			Year:   element.Year,
			Size:   element.Size,
			Rate:   element.Rate,
		}
	}
	result := proto.Clone(protoBooks).(*protoc.BooksSlice)
	return result
}

func main() {
	bookSlice := &books.JSONBook{}
	bookSlice.Books = make([]books.Book, 5)
	bookSlice.Books[0] = books.Book{
		ID:     "978",
		Title:  "Мюриель ежик",
		Author: "Мюриель Барбери",
		Year:   2009,
		Size:   400,
		Rate:   2.4,
	}
	bookSlice.Books[1] = books.Book{
		ID:     "978",
		Title:  "ежик Барбери",
		Author: "Мюриель Барбери",
		Year:   2009,
		Rate:   2.4,
	}
	bookSlice.Books[4] = books.Book{
		ID:     "3232-re",
		Title:  "tttt",
		Author: "aaaa",
		Rate:   3.6,
	}
	result, err := bookSlice.MarshalJSON()
	if err != nil {
		fmt.Println("error is:", err)
	}
	fmt.Println("JSON from structure:", string(result))
	bookNew := &books.JSONBook{}

	if err = bookNew.UnmarshalJSON(result); err != nil {
		fmt.Println("error is:", err)
	}

	for i, element := range bookNew.Books {
		fmt.Println("Unmarshled JSON element of bookNew[", i, "]:", element)
	}

	protoBook := ToProtobufStructure(bookNew)

	marshaled, errProto := MarshalProto(protoBook)
	if errProto != nil {
		fmt.Println("error is:", errProto)
	}

	// proto.Message(protoBook.ProtoReflect())

	fmt.Printf("\n----\nMarshaled protoc message:%s\n-----\n", marshaled)
	// for i := range marshaled {
	// 	fmt.Printf("\n----\nMarshaled protoc message[%v]:%s\n-----\n", i, marshaled[i].book)
	// }
	protoBookNew, errUnmarshalProto := UnmarshalProto(marshaled)
	if errUnmarshalProto != nil {
		fmt.Println("error is:", errUnmarshalProto)
	}
	for i := range protoBookNew.Books {
		fmt.Println("Unmarshled Protoc element of protoBookNew[", i, "]:", protoBookNew.Books[i])
	}
}
