// Copyright © 2018 Mephis Pheies <mephistommm@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package iterator

// a interface to traverse struct implement Aggregate
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// a interface represent data store and would return Iterator
type Aggregate interface {
	Iterator() Iterator
}

type Book struct {
	Name string
}

// BookShelf is a data store of book
type BookShelf struct {
	books []*Book
	last  int
}

func NewBookShelf(maxsize int) *BookShelf {
	return &BookShelf{
		books: make([]*Book, maxsize),
		last:  0,
	}
}

// GetBookAt ...
func (bs *BookShelf) GetBookAt(index int) *Book {
	return bs.books[index]
}

// AppendBook ...
func (bs *BookShelf) AppendBook(book *Book) {
	bs.books[bs.last] = book
	bs.last++
}

// GetLength ...
func (bs *BookShelf) GetLength() int {
	return bs.last
}

// Iterator ...
func (bs *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(bs)
}

// BookShelfIterator traverse BookShelf
type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{
		index:     0,
		bookShelf: bookShelf,
	}
}

// HasNext ...
func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.index < bsi.bookShelf.GetLength() {
		return true
	}

	return false
}

// Next ...
func (bsi *BookShelfIterator) Next() interface{} {
	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return book
}
