package batch

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
 * Job Definition/Flow
 */

type Job struct {
	reader    Reader
	processor Processor
	writer    Writer
}

func (b *Job) Reader(r Reader) Job {
	b.reader = r
	return *b
}

func (b *Job) Processor(p Processor) Job {
	b.processor = p
	return *b
}

func (b *Job) Writer(w Writer) Job {
	b.writer = w
	return *b
}

func (b *Job) Execute() {

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
