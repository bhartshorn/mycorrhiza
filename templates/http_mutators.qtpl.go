// Code generated by qtc from "http_mutators.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/http_mutators.qtpl:1
package templates

//line templates/http_mutators.qtpl:1
import "net/http"

//line templates/http_mutators.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/http_mutators.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/http_mutators.qtpl:3
func StreamEditHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, textAreaFill, warning string) {
//line templates/http_mutators.qtpl:3
	qw422016.N().S(`
<main class="edit">
`)
//line templates/http_mutators.qtpl:5
	qw422016.N().S(navHTML(rq, hyphaName, "edit"))
//line templates/http_mutators.qtpl:5
	qw422016.N().S(`
	<h1>Edit `)
//line templates/http_mutators.qtpl:6
	qw422016.E().S(hyphaName)
//line templates/http_mutators.qtpl:6
	qw422016.N().S(`</h1>
	`)
//line templates/http_mutators.qtpl:7
	qw422016.N().S(warning)
//line templates/http_mutators.qtpl:7
	qw422016.N().S(`
	<form method="post" class="edit-form"
			action="/upload-text/`)
//line templates/http_mutators.qtpl:9
	qw422016.E().S(hyphaName)
//line templates/http_mutators.qtpl:9
	qw422016.N().S(`">
		<textarea name="text">`)
//line templates/http_mutators.qtpl:10
	qw422016.E().S(textAreaFill)
//line templates/http_mutators.qtpl:10
	qw422016.N().S(`</textarea>
		<br/>
		<input type="submit"/>
		<a href="/page/`)
//line templates/http_mutators.qtpl:13
	qw422016.E().S(hyphaName)
//line templates/http_mutators.qtpl:13
	qw422016.N().S(`">Cancel</a>
	</form>
</main>
`)
//line templates/http_mutators.qtpl:16
}

//line templates/http_mutators.qtpl:16
func WriteEditHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, textAreaFill, warning string) {
//line templates/http_mutators.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/http_mutators.qtpl:16
	StreamEditHTML(qw422016, rq, hyphaName, textAreaFill, warning)
//line templates/http_mutators.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line templates/http_mutators.qtpl:16
}

//line templates/http_mutators.qtpl:16
func EditHTML(rq *http.Request, hyphaName, textAreaFill, warning string) string {
//line templates/http_mutators.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/http_mutators.qtpl:16
	WriteEditHTML(qb422016, rq, hyphaName, textAreaFill, warning)
//line templates/http_mutators.qtpl:16
	qs422016 := string(qb422016.B)
//line templates/http_mutators.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/http_mutators.qtpl:16
	return qs422016
//line templates/http_mutators.qtpl:16
}
