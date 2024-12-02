package iso639

type LanguageCode string

// Language represents an ISO 639 language.
type Language struct {
	Codes       []string  // All ISO codes (ISO 639-1, ISO 639-2, ISO 639-3)
	EnglishName string    // English name of the primary language
	NativeNames []string  // Native names of the language
	Variants    []Variant // Specific dialects or sub-languages
}

type Variant struct {
	Code        string   // Unique ISO 639-3 code for the variant
	EnglishName string   // English name of the variant
	NativeNames []string // Native names of the variant
}
