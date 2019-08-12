package protocol

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

func WriteMsg(w io.Writer, data []byte) (err error) {
	if err = binary.Write(w, binary.LittleEndian, uint32(len(data))); err != nil {
		return
	}

	if len(data) > 0 {
		if _, err = w.Write(data); err != nil {
			return
		}
	}

	return
}

func ReadMsg(r io.Reader, buff []byte) (msgbody []byte, err error) {
	var size uint32
	if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
		return
	}

	if size > uint32(cap(buff)) {
		err = errors.New(fmt.Sprintf("message size exceed, package size: %d, buffer cap: %d", size, cap(buff)))
		return
	}

	buf := buff[0:size]
	if _, e := io.ReadFull(r, buf); e != nil {
		err = e
		return
	}

	msgbody = buf
	return
}
