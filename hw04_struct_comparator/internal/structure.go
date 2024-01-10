package internal

type Book struct {
	id     int
	title  string
	author string
	year   uint16
	size   uint16
	rate   float32
}

func NewBook() Book {
	return Book{
		id:     0,
		title:  "Легенды и мифы южного бутово",
		author: "Шекспир Игнат Кимерсенович",
		year:   1907,
		size:   8,
		rate:   5,
	}
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) SetID(pages int) {
	b.id = pages
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

func IntSeq() func() int { // замыкание
	i := 0
	return func() int {
		i++
		return i
	}
}
