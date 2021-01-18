package iterator

type Iterator interface {
	first() interface{}
	next() interface{}
}

type Book struct {
	name string
}

type BookGroup struct {
	books []Book
}

func (b *BookGroup) add(book Book) {
	b.books = append(b.books, book)
}

type BookIterator struct {
	g     *BookGroup
	index int
}

func (b *BookGroup) createIterator() *BookIterator {
	return &BookIterator{b, 0}
}

func (b *BookIterator) first() interface{} {
	if len(b.g.books) > 0 {
		b.index = 0
		return b.g.books[b.index]
	}
	return nil
}

func (b *BookIterator) next() interface{} {
	if len(b.g.books) > b.index+1 {
		b.index++
		return b.g.books[b.index]
	}
	return nil
}
