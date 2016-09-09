package searcher

import "io"

const (
	topBit     = 128
	byteMaxLen = 8+1
)

func Encode(w io.Writer, n int) error {
	bytes := [byteMaxLen]byte{}
	i := byteMaxLen - 1
	for {
		bytes[i] = byte(n % topBit)
		if n < topBit {
			break
		}
		i--
		n /= topBit
	}
	bytes[byteMaxLen-1] += topBit
	_, err := w.Write(bytes[i:])
	return err
}

func Decode(r io.Reader) (n int, err error) {
	b := make([]byte, 1)
	for {
		_, err = r.Read(b)
		if err != nil {
			return
		}
		v := int(b[0])
		if v < topBit {
			n = n*topBit + v
		} else {
			n = n*topBit + v - topBit
			return
		}
	}
}
