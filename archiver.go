package dec

import "io"

type Archiver interface {
	Make()
	Open()

	Read(io.Reader)
	Write(io.Writer)
}
