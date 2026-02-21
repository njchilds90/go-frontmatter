package frontmatter

import (
	"errors"
	"regexp"
	"strings"
	"sync"
)

// Version is the library version.
const Version = "v1.0.0"

// Document represents parsed frontmatter + content.
type Document struct {
	Metadata map[string]string
	Content  string
}

var cache sync.Map

var frontmatterRegex = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*\n(.*)$`)

// Parse extracts YAML-style frontmatter from Markdown.
// Returns Document with metadata and content.
func Parse(input string) (*Document, error) {
	if input == "" {
		return &Document{
			Metadata: map[string]string{},
			Content:  "",
		}, nil
	}

	if val, ok := cache.Load(input); ok {
		return val.(*Document), nil
	}

	matches := frontmatterRegex.FindStringSubmatch(input)
	if len(matches) != 3 {
		return &Document{
			Metadata: map[string]string{},
			Content:  input,
		}, nil
	}

	metaBlock := matches[1]
	content := matches[2]

	meta := parseMetadata(metaBlock)

	doc := &Document{
		Metadata: meta,
		Content:  content,
	}

	cache.Store(input, doc)
	return doc, nil
}

// Generate builds markdown with frontmatter from Document.
func Generate(doc *Document) (string, error) {
	if doc == nil {
		return "", errors.New("document is nil")
	}

	if len(doc.Metadata) == 0 {
		return doc.Content, nil
	}

	var builder strings.Builder
	builder.WriteString("---\n")

	for k, v := range doc.Metadata {
		builder.WriteString(k)
		builder.WriteString(": ")
		builder.WriteString(v)
		builder.WriteString("\n")
	}

	builder.WriteString("---\n")
	builder.WriteString(doc.Content)

	return builder.String(), nil
}

// Get returns metadata value.
func (d *Document) Get(key string) string {
	return d.Metadata[key]
}

// Set sets metadata key/value.
func (d *Document) Set(key, value string) {
	if d.Metadata == nil {
		d.Metadata = make(map[string]string)
	}
	d.Metadata[key] = value
}

// Remove deletes metadata key.
func (d *Document) Remove(key string) {
	delete(d.Metadata, key)
}

// Has checks if metadata key exists.
func (d *Document) Has(key string) bool {
	_, ok := d.Metadata[key]
	return ok
}

// ValidateRequired ensures required metadata keys exist.
func (d *Document) ValidateRequired(keys []string) error {
	for _, k := range keys {
		if !d.Has(k) {
			return errors.New("missing required metadata: " + k)
		}
	}
	return nil
}

func parseMetadata(block string) map[string]string {
	meta := make(map[string]string)

	lines := strings.Split(block, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			meta[key] = value
		}
	}

	return meta
}
