package utils

import "io"

// Close io.Closer ni xavfsiz yopadi, xatoni e'tiborsiz qoldiradi.
func Close(c io.Closer) {
	_ = c.Close()
}
