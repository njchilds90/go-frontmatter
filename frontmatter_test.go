package frontmatter

import "testing"

func TestParseBasic(t *testing.T) {
	input := `---
title: Hello
author: Nick
---
Content here`

	doc, err := Parse(input)
	if err != nil {
		t.Fatal(err)
	}

	if doc.Get("title") != "Hello" {
		t.Error("title not parsed")
	}

	if doc.Content != "Content here" {
		t.Error("content not parsed")
	}
}

func TestGenerate(t *testing.T) {
	doc := &Document{
		Metadata: map[string]string{
			"title": "Test",
		},
		Content: "Body",
	}

	out, err := Generate(doc)
	if err != nil {
		t.Fatal(err)
	}

	if out == "" {
		t.Error("generate failed")
	}
}
