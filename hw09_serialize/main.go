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

type ProtocBinary struct {
	book []byte
}

type Marshaller interface {
	MarshalJSON() ([]byte, error)
	MarshalProto() ([]ProtocBinary, error)
}

type Unmarshaller interface {
	UnmarshalJSON([]byte) error
	UnmarshalProto(marshaled []ProtocBinary) error
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

func (protoBook *ProtocBook) MarshalProto() ([]ProtocBinary, error) {
	finalStructure := make([]ProtocBinary, len(protoBook.Books))
	for i := range protoBook.Books {
		if protoBook.Books[i].ProtoReflect().IsValid() {
			var ctr Message = &protoBook.Books[i]
			marshaled, err := proto.Marshal(ctr)
			if err == nil {
				finalStructure[i].book = marshaled
			} else {
				return nil, err
			}
		}
	}
	return finalStructure, nil
}

func UnmarshalProto(marshaled []ProtocBinary) (ProtocBook, error) {
	var protoBook ProtocBook
	protoBook.Books = make([]protoc.Book, len(marshaled))
	for i := range marshaled {
		err := proto.Unmarshal(marshaled[i].book, &protoBook.Books[i])
		if err != nil {
			return protoBook, err
		}
	}
	return protoBook, nil
}

func ToProtobufStructure(bookSlice *books.JSONBook) ProtocBook {
	var protoBooks ProtocBook
	protoBooks.Books = make([]protoc.Book, len(bookSlice.Books))
	for i, element := range bookSlice.Books {
		protoBooks.Books[i] = protoc.Book{
			ID:     element.ID,
			Title:  element.Title,
			Author: element.Author,
			Year:   element.Year,
			Size:   element.Size,
			Rate:   element.Rate,
		}
	}
	return protoBooks
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

	marshaled, errProto := protoBook.MarshalProto()
	if errProto != nil {
		fmt.Println("error is:", errProto)
	}
	for i := range marshaled {
		fmt.Printf("\n----\nMarshaled protoc message[%v]:%s\n-----\n", i, marshaled[i].book)
	}
	protoBookNew, errUnmarshalProto := UnmarshalProto(marshaled)
	if errUnmarshalProto != nil {
		fmt.Println("error is:", errUnmarshalProto)
	}
	for i := range protoBookNew.Books {
		fmt.Println("Unmarshled Protoc element of protoBookNew[", i, "]:", proto.Clone(&protoBookNew.Books[i]).(*protoc.Book))
	}
}
