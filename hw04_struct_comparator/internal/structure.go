package internal

type Book struct {
	id     string
	title  string
	author string
	year   uint16
	size   uint16
	rate   float32
}

func NewBook(i string, t string, a string, y uint16, s uint16, r float32) *Book {
	return &Book{
		id:     i,
		title:  t,
		author: a,
		year:   y,
		size:   s,
		rate:   r,
	}
}

func (b *Book) ID() string {
	return b.id
}

func (b *Book) SetID(id string) {
	b.id = id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) Year() uint16 {
	return b.year
}

func (b *Book) SetYear(year uint16) {
	b.year = year
}

func (b *Book) Size() uint16 {
	return b.size
}

func (b *Book) SetSize(size uint16) {
	b.size = size
}

func (b *Book) Rate() float32 {
	return b.rate
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}
