/*
 * Copyright (c) 2025 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package markdown

import (
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
)

// Markdown represents a markdown object.
type Markdown struct {
	text     string
	md       goldmark.Markdown
	parser   parser.Parser
	renderer renderer.Renderer
}

// New creates a new markdown object from a UTF-8 bytes text.
func New(text string) *Markdown {
	return &Markdown{text: text}
}

// Options sets the options for the markdown object.
func (p *Markdown) Options(opts ...goldmark.Option) *Markdown {
	p.md = goldmark.New(opts...)
	return p
}

// Parser sets the parser for the markdown object.
func (p *Markdown) Parser(parser parser.Parser) *Markdown {
	p.parser = parser
	return p
}

// Renderer sets the renderer for the markdown object.
func (p *Markdown) Renderer(renderer renderer.Renderer) *Markdown {
	p.renderer = renderer
	return p
}

// Convert interprets the markdown object and write rendered
// contents to a writer w.
func (p *Markdown) Convert(writer io.Writer, opts ...parser.ParseOption) error {
	md := p.md
	if md == nil {
		md = goldmark.New()
	}
	if p.parser != nil {
		md.SetParser(p.parser)
	}
	if p.renderer != nil {
		md.SetRenderer(p.renderer)
	}
	return md.Convert([]byte(p.text), writer, opts...)
}
