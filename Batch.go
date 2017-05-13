package main

import (
	"time"
)

/**
 * Reader
 */

type Reader interface {
	open()
	read() Record
	close()
}

/**
 * Processor
 */

type Processor interface {
	processRecord(r Record) Record
}

/**
 * Record
 */

type Record interface {
	header() Header
	payload() map[string]interface{}
}

type Header struct {
	number       int64
	source       string
	creationDate time.Time
}

/**
 * Writer
 */

type Writer interface {
	open()
	writeRecord(record Record)
	close()
}

/**
 * Batch Definition/Flow
 */

type Batch struct {
	reader    Reader
	processor Processor
	writer    Writer
}

func (b *Batch) setReader(r Reader) Batch {
	b.reader = r
	return *b
}

func (b *Batch) setProcessor(p Processor) Batch {
	b.processor = p
	return *b
}

func (b *Batch) setWriter(w Writer) Batch {
	b.writer = w
	return *b
}

func (b *Batch) execute() {

	reader := b.reader
	processor := b.processor
	writer := b.writer

	reader.open()

	record := reader.read()
	for record != nil {
		writer.writeRecord(processor.processRecord(record))
		record = reader.read()
	}

	reader.close()
	writer.close()
}

func main() {
	// something something something
}
