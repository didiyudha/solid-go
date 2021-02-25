package main

import (
	"fmt"
	"log"
)

// Document represents document information.
type Document struct {
	Content string
}

// GetContent - get content of a document.
func (d *Document) GetContent() string {
	return d.Content
}

// Printer is an interface that defines print API.
type Printer interface {
	Print(d Document) string
}

// Scanner is an interface that defines scan API.
type Scanner interface {
	Scan(d *Document) error
}

// DocumentMachine - imlementation of Interface Segragation.
type DocumentMachine interface {
	Scanner
	Printer
}

// DocumentMachineImpl - document machine concrete implementation.
type DocumentMachineImpl struct{}

// Print implementation.
func (impl *DocumentMachineImpl) Print(d Document) string {
	return d.GetContent()
}

// Scan implementation.
func (impl *DocumentMachineImpl) Scan(d *Document) error {
	d.Content = "Some content"
	return nil
}

// ScanDocument - scan document.
func ScanDocument(machine Scanner, doc *Document) error {
	return machine.Scan(doc)
}

// PrintDocument - print document.
func PrintDocument(machine Printer, doc Document) string {
	return machine.Print(doc)
}

func main() {
	var document Document
	documentMachine := new(DocumentMachineImpl)

	if err := ScanDocument(documentMachine, &document); err != nil {
		log.Fatal(err)
	}

	fmt.Println(PrintDocument(documentMachine, document))

}
