package bridge

type computer interface {
	Print()
	SetPrinter(printer)
}