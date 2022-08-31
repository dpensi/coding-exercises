package list

type linkedList[T any] struct {
	data   *T
	next   List[T]
	last   List[T]
	length int
}

type List[T any] interface {
	Append(List[T])
	Next() List[T]
	Data() *T
	HasNext() bool
	Last() List[T]
	setLast(List[T])
	Length() int
}

func New[T any](data *T) List[T] {
	newList := &linkedList[T]{
		data: data,
		next: nil,
	}
	newList.last = newList
	if data != nil {
		newList.length = 1
	}
	return newList
}

func (l *linkedList[T]) Append(list List[T]) {
	if l.data == nil {
		if linkedList, ok := list.(*linkedList[T]); ok {
			l.data = linkedList.data
			l.next = nil
			l.length = 1
			l.last = l
		}
		return
	}

	if l.next == nil {
		l.next = list
		l.last = l.next
		l.next.setLast(l.next)
		l.length = 2
		return
	}

	l.last.Append(list)
	l.setLast(list.Last())
	l.length += list.Length()
}

func (l *linkedList[T]) Next() List[T] {
	return l.next
}

func (l *linkedList[T]) Data() *T {
	return l.data
}

func (l *linkedList[T]) HasNext() bool {
	return l.next != nil
}

func (l *linkedList[T]) Last() List[T] {
	return l.last
}

func (l *linkedList[T]) setLast(lst List[T]) {
	l.last = lst
}

func (l *linkedList[T]) Length() int {
	return l.length
}
