package batch

import "errors"

/**
 * Reader
 */

type Reader interface {
	Open()
	Read() (interface{}, error)
	Close()
}

/**
 * Processor
 */

type Processor interface {
	Process(r interface{}) (interface{}, error)
}

/**
 * Writer
 */

type Writer interface {
	Open()
	Write(record interface{}) error
	Close()
}

/**
 * Job Definition/Flow
 * Reader -> Processor -> Writer
 */

type Job struct {
	reader    Reader
	processor Processor
	writer    Writer
}

func (b *Job) Reader(r Reader) {
	b.reader = r
}

func (b *Job) Processor(p Processor) {
	b.processor = p
}

func (b *Job) Writer(w Writer) {
	b.writer = w
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
	 * todo: handle the errors :(
	 */

	reader.Open()

	record, _ := reader.Read()
	for record != nil {
		processed, _ := processor.Process(record)
		writer.Write(processed)
		record, _ = reader.Read()
	}

	reader.Close()
	writer.Close()

	return nil
}
