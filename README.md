#### TODO
- logging
- job status

### Sample
```golang

// Reader

type genericReader struct {
	name    string
	counter int
	limit   int
}

func (this *genericReader) Open() {
	fmt.Println("Opening Source", this.name)
}

func (this *genericReader) Close() {
	fmt.Println("Closing Source", this.name)
}

func (this *genericReader) Read() interface{} {

	for this.counter < this.limit {
		this.counter++
		return strconv.Itoa(this.counter)
	}

	return nil
}


// Processor

type genericProcessor struct{
	append string
}

func (this *genericProcessor) Process(d interface{}) interface{} {
	return d.(string) + this.append
}


// Writer

type genericWriter struct {
	name string
}

func (this *genericWriter) Open() {
	fmt.Println("Opening Writer", this.name)
}

func (this *genericWriter) Write(d interface{}) error {
	fmt.Println(this.name, ":", d)
	return nil
}

func (this *genericWriter) Close() {
	fmt.Println("closing writer", this.name)
}

func main() {
	job := batch.Job{}
	job.Reader(&genericReader{"Reader", 0, 10})
	job.Processor(&genericProcessor{"-"})
	job.Writer(&genericWriter{"Writer"})
	job.Execute()
}
```
