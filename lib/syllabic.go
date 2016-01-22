package lib

// Koremutake syllables, from http://shorl.com/koremutake
var SYLLABLES = [...]string {
	"ba", "be", "bi", "bo", "bu", "by", "da", "de", "di", "do", "du", "dy",
	"fa", "fe", "fi", "fo", "fu", "fy", "ga", "ge", "gi", "go", "gu", "gy",
	"ha", "he", "hi", "ho", "hu", "hy", "ja", "je", "ji", "jo", "ju", "jy",
	"ka", "ke", "ki", "ko", "ku", "ky", "la", "le", "li", "lo", "lu", "ly",
	"ma", "me", "mi", "mo", "mu", "my", "na", "ne", "ni", "no", "nu", "ny",
	"pa", "pe", "pi", "po", "pu", "py", "ra", "re", "ri", "ro", "ru", "ry",
	"sa", "se", "si", "so", "su", "sy", "ta", "te", "ti", "to", "tu", "ty",
	"va", "ve", "vi", "vo", "vu", "vy", "bra", "bre", "bri", "bro", "bru",
	"bry", "dra", "dre", "dri", "dro", "dru", "dry", "fra", "fre", "fri", "fro",
	"fru", "fry", "gra", "gre", "gri", "gro", "gru", "gry", "pra", "pre", "pri",
	"pro", "pru", "pry", "sta", "ste", "sti", "sto", "stu", "sty", "tra", "tre",
}

func randSyllable() string {
	return SYLLABLES[randInt(len(SYLLABLES))]
}

func SyllabicPassword(length int) string {
	var result string

	for len(result) < length {
		result += randSyllable()
	}

	return result
}