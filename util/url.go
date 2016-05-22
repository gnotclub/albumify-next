package util

import (
    "bytes"
    "math"
)

var alphabet = []byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func UrlEncode(id int64) string {
    var ret string
    var alphabetLen int64 = int64(len(alphabet))
    for id > 0 {
        ret += string(alphabet[id % alphabetLen])
        id /= alphabetLen
    }
    ret = ReverseString(ret)
    return ret
}

func UrlDecode(url string) int64 {
    digits := make([]byte, 5)
    var ret int64 = 0
    for _, r := range url {
        digits = append(digits, byte(bytes.IndexByte(alphabet, byte(r))))
    }
    for i, r := range digits {
        ret += int64(r) * int64(math.Pow(float64(len(alphabet)), float64(len(digits) - 1 - i)))
    }
    return ret
}
