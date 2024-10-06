func HashCode(dec string) string {
	size := len(dec)
	hashed := ""
	for _, char := range dec {
		hash := (int(char) + size) % 127
		if hash < 32 || hash > 126 {
			hash += 33
		}
		hashed += string(hash)
	}
	return hashed
}

func HashCode(dec string) string {
	size := len(dec)
	result := ""

	for _, char := range dec {
		ascii := int(char)
		Hash := (ascii + size) % 127

		if Hash < 33 {
			Hash += 33
		}
		result += string(Hash)
	}
	return result
}
