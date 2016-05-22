package util

func ReverseString(str string) string {
    // Store runes
    n := 0
    runes := make([]rune, len(str))
    for _, r := range str {
        runes[n] = r
        n++
    }
    runes = runes[0:n]

    // Reverse
    for i := 0; i < n / 2; i++ {
        runes[i], runes[n - 1 - i] = runes[n - 1 - i], runes[i]
    }
    return string(runes)
}
