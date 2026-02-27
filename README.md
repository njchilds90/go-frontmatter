# go-frontmatter

Minimal, zero-dependency YAML frontmatter parser and generator for Go.

Inspired by Python’s `python-frontmatter`, redesigned for modern AI pipelines.

---

## Features

- Parse YAML-style frontmatter
- Generate markdown with frontmatter
- Metadata Get / Set / Remove
- Required field validation
- Built-in caching
- AI pipeline friendly
- Zero external dependencies
- MIT Licensed

---

## Install

```bash
go get github.com/njchilds90/go-frontmatter
```

---

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/njchilds90/go-frontmatter"
)

func main() {
	input := `---
title: My Post
author: Nick
---
Hello world`

	doc, _ := frontmatter.Parse(input)

	fmt.Println(doc.Get("title"))

	doc.Set("date", "2026-02-21")

	output, _ := frontmatter.Generate(doc)

	fmt.Println(output)
}
```

---

## AI Use Cases

- Structured content pipelines
- RAG document ingestion
- Static site generation
- Markdown CMS systems
- LLM document normalization
- Metadata validation

---

## Release

Tag: v1.0.0  
Title: Initial Stable Release  

---

## License

MIT