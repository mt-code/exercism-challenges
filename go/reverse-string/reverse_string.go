package reverse

func Reverse(original string) string {
	/**
		Convert our string into a rune slice.

		We do that because of the following. If we don't convert
		the string characters to runes but simply reverse the string
		itself, we're in effect reversing the string on a byte-by-byte
		basis. That works fine if we have a string with ASCII characters,
		which each are indeed a single byte big. But characters in a
		UTF-8 encoded string are between 1 and 4 bytes big. So to
		reverse a string in a Unicode-compatible manner we first
		cast the string into a slice of runes.
	 */
	runes := []rune(original)

	for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}