package batch

import (
	"errors"
	"time"
)

/**
 * Reader
 */

type Reader interface {
	Open()
	Read() Record
	Close()
}

/**
 * Processor
 */

type Processor interface {
	ProcessRecord(r Record) Record
}

/**
 * Record
 */

type Record interface {
	Header() Header
	Payload() map[string]interface{}
}

type Header struct {
	Number   int64
	Source   string
	Creation time.Time
}

/**
 * Writer
 */

type Writer interface {
	Open()
	WriteRecord(record Record)
	Close()
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

func (b *Job) Execute() error {

	reader := b.reader
	processor := b.processor
	writer := b.writer

	if reader == nil {
		return errors.New("There is no Reader")
	}

	if processor == nil {
		return errors.New("There is no Processor")
	}

	if writer == nil {
		return errors.New("There is no Writer")
	}

	/**
	 * Start Our Main Loop
	 */

	reader.Open()

	record := reader.Read()
	for record != nil {
		writer.WriteRecord(processor.ProcessRecord(record))
		record = reader.Read()
	}

	reader.Close()
	writer.Close()

	return nil
}
