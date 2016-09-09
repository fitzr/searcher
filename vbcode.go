package searcher

import "io"

const (
	topBit     = 128
	byteMaxLen = 10
)

func Encode(w io.Writer, n uint64) error {
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

func Decode(r io.Reader) (n uint64, err error) {
	b := make([]byte, 1)
	for {
		_, err = r.Read(b)
		if err != nil {
			return
		}
		v := uint64(b[0])
		if v < topBit {
			n = n*topBit + v
		} else {
			n = n*topBit + v - topBit
			return
		}
	}
}
