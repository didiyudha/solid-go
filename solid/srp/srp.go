package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Journal data type.
type Journal struct {
	entryCounter int
	entries      []string
}

// AddEntry add an entry to journal data.
func (j *Journal) AddEntry(text string) int {
	j.entryCounter++
	entry := fmt.Sprintf("%d: %s", j.entryCounter, text)
	j.entries = append(j.entries, entry)
	return j.entryCounter
}

// RemoveEntry removes an entry from journal.
func (j *Journal) RemoveEntry(index int) {
	if index >= len(j.entries) {
		return
	}
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Breaks SRP.

// SaveToFile save journal entry data to file.
func (j *Journal) SaveToFile(fileName string) error {
	return errors.New("not implemented")
}

// LoadFromFile - load journal entry data from file.
func (j *Journal) LoadFromFile(fileName string) error {
	return errors.New("not implemented")
}

// FilePersistence - We need to add one type that responsible to save journal entry data to file
// or to load entry data from file.
type FilePersistence struct{}

// SaveToFile save journal entry data to file.
func (p *FilePersistence) SaveToFile(journal *Journal, fileName string) error {
	data := []byte(strings.Join(journal.entries, "\n"))
	return ioutil.WriteFile(fileName, data, 0644)
}

// LoadFromFile load entry data from file.
func (p *FilePersistence) LoadFromFile(journal *Journal, fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	entry := strings.Split(string(b), "\n")
	journal.entries = entry
	return nil
}

func main() {
	journal := new(Journal)
	journal.AddEntry("ONE")
	journal.AddEntry("TWO")
	journal.AddEntry("THREE")

	fmt.Println("Journal Entries: ", journal.String())

	journal.RemoveEntry(0)

	fmt.Println("Journal Entries: ", journal.String())

	fp := new(FilePersistence)
	if err := fp.SaveToFile(journal, "journal.txt"); err != nil {
		log.Fatalln(err)
	}

	var secondJournal Journal
	if err := fp.LoadFromFile(&secondJournal, "journal.txt"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Second Journal Entries: ", secondJournal.String())
}
