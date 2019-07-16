package sample

import (
	"testing"
)

/* read an closed channel will return immediately */
func TestReadClosedChannel(t *testing.T) {
	closed := make(chan bool)
	close(closed)
	<-closed
}
