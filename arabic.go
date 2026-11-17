package elnemer

import (
	"strings"
	"unicode"
)

// StripTashkeel removes Arabic diacritics and vowels.
func StripTashkeel(text string) string {
	var b strings.Builder
	for _, r := range text {
		// Arabic Tashkeel range \u064B - \u0652, \u0670, \u0640 (tatweel)
		if (r >= 0x064B && r <= 0x0652) || r == 0x0670 || r == 0x0640 {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// JummalValue calculates Hisab Al-Jummal (الأبجد الكبير) for Arabic text.
func JummalValue(text string) int {
	values := map[rune]int{
		'ا': 1, 'أ': 1, 'إ': 1, 'آ': 1, 'ٱ': 1,
		'ب': 2, 'ج': 3, 'د': 4, 'ه': 5, 'ة': 5, 'و': 6, 'ؤ': 6,
		'ز': 7, 'ح': 8, 'ط': 9, 'ي': 10, 'ى': 10, 'ئ': 10,
		'ك': 20, 'ل': 30, 'م': 40, 'ن': 50, 'س': 60, 'ع': 70,
		'ف': 80, 'ص': 90, 'ق': 100, 'ر': 200, 'ش': 300, 'ت': 400,
		'ث': 500, 'خ': 600, 'ذ': 700, 'ض': 800, 'ظ': 900, 'غ': 1000,
	}

	clean := StripTashkeel(text)
	total := 0
	for _, r := range clean {
		if v, ok := values[r]; ok {
			total += v
		} else if !unicode.IsSpace(r) {
			// ignore non-arabic or unknown runes
		}
	}
	return total
}
