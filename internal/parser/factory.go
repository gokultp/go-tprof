package parser

// GetParser is a factory for parsers
func GetParser(txt string, failFunc string) Parser {
	if p := NewDescriptionParser(txt); p.IsAbleToParse() {
		return p
	}
	if p := NewPackageStatusParser(txt); p.IsAbleToParse() {
		return p
	}
	if p := NewTestStatusParser(txt); p.IsAbleToParse() {
		return p
	}
	if p := NewSpippedPackageParser(txt); p.IsAbleToParse() {
		return p
	}
	if ailFunc != "" {
		return NewErrorParser(txt)
	}
	return NewDefaultParser(txt)
}
