package core

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type SpotifyID []byte

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func (s SpotifyID) ToBase62() string {
	bi := big.Int{}
	bi.SetBytes([]byte(s))
	rem := big.NewInt(0)
	base := big.NewInt(62)
	zero := big.NewInt(0)
	result := ""
	for bi.Cmp(zero) > 0 {
		_, rem = bi.DivMod(&bi, base, rem)
		result += string(alphabet[int(rem.Uint64())])
	}
	for len(result) < 22 {
		result += "0"
	}
	return reverse(result)
}

func (s SpotifyID) ToHex() string {
	return fmt.Sprintf("%x", s.ToBase62())
}

func NewSpotifyIDFromBase62(id string) (SpotifyID, error) {
	base := big.NewInt(62)
	n := &big.Int{}
	for _, c := range []byte(id) {
		if idx := strings.IndexByte(alphabet, c); idx != -1 {
			d := big.NewInt(int64(idx))
			n = n.Mul(n, base)
			n = n.Add(n, d)
		} else {
			return SpotifyID(nil), errors.New("NewSpotifyIDFromBase62 : char not present in alphabet")
		}
	}
	return SpotifyID(fixSpotifyIDBytesLen(n.Bytes())), nil
}

func NewSpotifyIDFromRaw(data []byte) (SpotifyID, error) {
	if len(data) != 16 {
		return SpotifyID(nil), errors.New("NewSpotifyIDFromRaw : len(data) not equal to 16")
	}
	return SpotifyID(data), nil
}

func fixSpotifyIDBytesLen(nBytes []byte) []byte {
	if len(nBytes) < 16 {
		paddingBytes := make([]byte, 16-len(nBytes))
		nBytes = append(paddingBytes, nBytes...)
	}
	return nBytes
}
