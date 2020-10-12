package bc

type PositionalReader struct {
	data   string
	cursor int
}

func NewPositionalReader(data string) PositionalReader {
	return PositionalReader{data: data, cursor: 0}
}

func (pr *PositionalReader) Read(n int) string {
	r := pr.data[pr.cursor : pr.cursor+n]
	pr.cursor += n
	return r
}
