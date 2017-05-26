package blackfridaytemplaterenderer

import (
	"bytes"

	"github.com/russross/blackfriday"

	tt "text/template"
)

type RendererTemplates struct {
	ParagraphStart *tt.Template
	ParagraphEnd   *tt.Template
	Emphasis       *tt.Template
	NormalText     *tt.Template
	Entity         *tt.Template
}

// Templated is a type that implements the Renderer interface for Templated output.
//
// Do not create this directly, instead use the TemplatedRenderer function.
type Templated struct {
	rendererTemplates RendererTemplates
}

func TemplatedRenderer(flags int, rendererTemplates RendererTemplates) blackfriday.Renderer {
	return &Templated{rendererTemplates: rendererTemplates}
}

func (options *Templated) GetFlags() int {
	return 0
}

// render code chunks using verbatim, or listings if we have a language
func (options *Templated) BlockCode(out *bytes.Buffer, text []byte, lang string) {

}

func (options *Templated) TitleBlock(out *bytes.Buffer, text []byte) {

}

func (options *Templated) BlockQuote(out *bytes.Buffer, text []byte) {

}

func (options *Templated) BlockHtml(out *bytes.Buffer, text []byte) {

}

func (options *Templated) Header(out *bytes.Buffer, text func() bool, level int, id string) {

}

func (options *Templated) HRule(out *bytes.Buffer) {

}

func (options *Templated) List(out *bytes.Buffer, text func() bool, flags int) {

}

func (options *Templated) ListItem(out *bytes.Buffer, text []byte, flags int) {

}

func (options *Templated) Paragraph(out *bytes.Buffer, text func() bool) {
	marker := out.Len()
	//out.WriteString( rendererTemplates.ParagraphStart.tmpl.Execute()
	options.rendererTemplates.ParagraphStart.Execute(out, nil)
	if !text() {
		out.Truncate(marker)
		return
	}
	options.rendererTemplates.ParagraphEnd.Execute(out, nil)
}

func (options *Templated) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {

}

func (options *Templated) TableRow(out *bytes.Buffer, text []byte) {

}

func (options *Templated) TableHeaderCell(out *bytes.Buffer, text []byte, align int) {

}

func (options *Templated) TableCell(out *bytes.Buffer, text []byte, align int) {

}

// TODO: this
func (options *Templated) Footnotes(out *bytes.Buffer, text func() bool) {

}

func (options *Templated) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {

}

func (options *Templated) AutoLink(out *bytes.Buffer, link []byte, kind int) {

}

func (options *Templated) CodeSpan(out *bytes.Buffer, text []byte) {

}

func (options *Templated) DoubleEmphasis(out *bytes.Buffer, text []byte) {

}

type textContext struct {
	Text string
}

func (options *Templated) Emphasis(out *bytes.Buffer, text []byte) {

	if options.rendererTemplates.Emphasis != nil {
		context := textContext{string(text)}
		options.rendererTemplates.Emphasis.Execute(out, context)
	}
}

func (options *Templated) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {

}

func (options *Templated) LineBreak(out *bytes.Buffer) {

}

func (options *Templated) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {

}

func (options *Templated) RawHtmlTag(out *bytes.Buffer, tag []byte) {
}

func (options *Templated) TripleEmphasis(out *bytes.Buffer, text []byte) {

}

func (options *Templated) StrikeThrough(out *bytes.Buffer, text []byte) {

}

// TODO: this
func (options *Templated) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {

}

type entityContext struct {
	Entity string
}

func (options *Templated) Entity(out *bytes.Buffer, entity []byte) {
	if options.rendererTemplates.Entity != nil {
		context := entityContext{string(entity)}
		options.rendererTemplates.Entity.Execute(out, context)
	}
}

func (options *Templated) NormalText(out *bytes.Buffer, text []byte) {
	if options.rendererTemplates.NormalText != nil {
		context := textContext{string(text)}
		options.rendererTemplates.NormalText.Execute(out, context)
	}
}

// header and footer
func (options *Templated) DocumentHeader(out *bytes.Buffer) {

}

func (options *Templated) DocumentFooter(out *bytes.Buffer) {

}
