package link

type department interface {
	execute(*Patient)
	SetNext(department)
}