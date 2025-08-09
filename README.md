# go-iso639
![](https://img.shields.io/badge/golang-00ADD8?logo=go&amp;logoColor=white)
![](https://img.shields.io/badge/license-MIT-blue)
![](https://img.shields.io/badge/Version-0.0.1-purple)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazybark/go-iso639)](https://goreportcard.com/report/github.com/lazybark/go-iso639)
![GitHub last commit](https://img.shields.io/github/last-commit/lazybark/go-iso639)

A high-performance Go library for working with ISO 639 language codes and metadata. Provides fast lookups for language information including codes, names, scripts, families, regions, and variants.

## Features

- **Ultra-fast lookups** - O(1) hash map-based searches with sub-nanosecond performance
- **Full language data** - ISO 639-1, 639-2, and 639-3 codes with rich metadata
- **Case-insensitive search** - All lookups work regardless of case
- **Multiple search methods** - Find languages by code, name, script, family, region, type, or scope
- **Script detection** - Identify writing systems used by languages
- **Language variants** - Support for dialects and regional variations
- **Thread-safe** - Concurrent-safe initialization and lookups
- **Zero allocations** - Most operations have zero memory overhead
- **Rich metadata** - Native names, language families, geographic regions, and more

## Performance

```
BenchmarkByCode-12              161M ops/sec     7.6 ns/op      0 B/op    0 allocs/op
BenchmarkByName-12              52M ops/sec      22.5 ns/op     8 B/op    1 allocs/op
BenchmarkByNameIgnoreCase-12    122M ops/sec     9.5 ns/op      0 B/op    0 allocs/op
```

## Installation

```bash
go get github.com/lazybark/go-iso639
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/lazybark/go-iso639"
)

func main() {
    // Find language by ISO code.
    lang := iso639.ByCode("en")
    if lang != nil {
        fmt.Printf("Language: %s\n", lang.EnglishName) // Output: English
        fmt.Printf("Family: %s\n", lang.Family)        // Output: Indo-European
        fmt.Printf("Scripts: %v\n", lang.Scripts)      // Output: [Latin]
    }

    // Case-insensitive search.
    lang = iso639.ByCode("EN") 
    lang = iso639.ByCode("en") // Works with any case
    
    // Search by name.
    lang = iso639.ByName("Spanish")
    lang = iso639.ByName("Español") // Works with native names too.
    
    // Find all languages in a family.
    indoEuropean := iso639.ByFamily(iso639.IndoEuropean)
    fmt.Printf("Found %d Indo-European languages\n", len(indoEuropean))
}
```

## Usage Examples

### Finding Languages by Different Criteria

```go
// By ISO codes (639-1, 639-2, 639-3).
english := iso639.ByCode("en")     // ISO 639-1
english = iso639.ByCode("eng")     // ISO 639-2
chinese := iso639.ByCode("zh")     // Chinese (macrolanguage)
mandarin := iso639.ByCode("cmn")   // Mandarin (specific variant)

// By language names (English or native).
spanish := iso639.ByName("Spanish")
spanish = iso639.ByName("Español")
russian := iso639.ByName("Русский")

// Case doesn't matter.
german := iso639.ByCode("DE")
german = iso639.ByName("GERMAN")
```

### Exploring Language Metadata

```go
lang := iso639.ByCode("ar") // Arabic.
if lang != nil {
    fmt.Printf("English Name: %s\n", lang.EnglishName)
    fmt.Printf("Native Names: %v\n", lang.NativeNames)
    fmt.Printf("Language Family: %s\n", lang.Family)
    fmt.Printf("Scripts: %v\n", lang.Scripts)
    fmt.Printf("Regions: %v\n", lang.Regions)
    fmt.Printf("Type: %s\n", lang.Type)        // Living, Ancient, Constructed.
    fmt.Printf("Scope: %s\n", lang.Scope)      // Individual, Macrolanguage.
    
    // Check all ISO codes
    for _, code := range lang.Codes {
        fmt.Printf("ISO Code: %s\n", code)
    }
}
```

### Working with Language Variants

```go
// Chinese is a macrolanguage with many variants.
chinese := iso639.ByCode("zh")
if chinese != nil {
    fmt.Printf("Chinese has %d variants:\n", len(chinese.Variants))
    for _, variant := range chinese.Variants {
        fmt.Printf("- %s (%s)\n", variant.EnglishName, variant.Code)
    }
}

// Find a specific variant.
mandarin := iso639.ByCode("cmn") // Returns the parent Chinese language
cantonese := iso639.ByCode("yue") // Returns the parent Chinese language
```

### Filtering by Language Properties

```go
// Find all living languages.
livingLangs := iso639.ByType(iso639.Living)
fmt.Printf("Living languages: %d\n", len(livingLangs))

// Find all ancient languages.
ancientLangs := iso639.ByType(iso639.Ancient)
fmt.Printf("Ancient languages: %d\n", len(ancientLangs))

// Find all individual languages (not macrolanguages).
individualLangs := iso639.ByScope(iso639.Individual)

// Find all languages using Latin script.
latinScriptLangs := iso639.ByScript("Latin")

// Find all languages spoken in Europe.
europeanLangs := iso639.ByRegion("Europe")

// Find all Indo-European languages.
indoEuropeanLangs := iso639.ByFamily("Indo-European")
```

### Utility Functions

```go
// Validate language codes.
if iso639.IsValidCode("en") {
    fmt.Println("'en' is a valid language code")
}

// Check if a language supports a specific script.
lang := iso639.ByCode("sr") // Serbian
if iso639.HasScript(lang, iso639.CyrillicScript) {
    fmt.Println("Serbian uses Cyrillic script")
}
if iso639.HasScript(lang, iso639.LatinScript) {
    fmt.Println("Serbian also uses Latin script")
}

// Get all languages.
allLangs := iso639.GetAllLanguages()
fmt.Printf("Total languages: %d\n", len(allLangs))
```

## Script Detection

The library includes a script detection package for identifying writing systems:

```go
import "github.com/lazybark/go-iso639/script/detect"

// Detect script from text.
script := detect.DetectScript("Hello World")      // Returns "Latin"
script = detect.DetectScript("Привет мир")       // Returns "Cyrillic"
script = detect.DetectScript("こんにちは")         // Returns "Hiragana"
```

## Development

### Running Tests (requires gotest.tools/gotestsum)

```bash
make test
```

### Running Benchmarks

```bash
make bench
```

### Linting (requires golangci-lint as Docker image)

```bash
make lint
```

## Language Data Coverage

This library includes data for:

- **ISO 639-1** codes (2-letter codes for major languages)
- **ISO 639-2** codes (3-letter codes, both bibliographic and terminologic)
- **ISO 639-3** codes (3-letter codes for all known languages)
- **Language families** (Indo-European, Sino-Tibetan, Afro-Asiatic, etc.)
- **Writing scripts** (Latin, Cyrillic, Arabic, Chinese, etc.)
- **Geographic regions** (Europe, Asia, Africa, Americas, etc.)
- **Language types** (Living, Ancient, Constructed)
- **Language scopes** (Individual languages vs. Macrolanguages)
- **Native names** in original scripts
- **Language variants** and dialects

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.