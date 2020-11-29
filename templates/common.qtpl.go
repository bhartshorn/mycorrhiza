// Code generated by qtc from "common.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/common.qtpl:1
package templates

//line templates/common.qtpl:1
import "net/http"

//line templates/common.qtpl:2
import "github.com/bouncepaw/mycorrhiza/user"

//line templates/common.qtpl:3
import "github.com/bouncepaw/mycorrhiza/util"

// This is the <nav> seen on top of many pages.

//line templates/common.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/common.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/common.qtpl:7
type navEntry struct {
	path  string
	title string
}

var navEntries = []navEntry{
	{"page", "Hypha"},
	{"edit", "Edit"},
	{"text", "Raw text"},
	{"history", "History"},
	{"revision", "NOT REACHED"},
	{"rename-ask", "Rename"},
	{"delete-ask", "Delete"},
}

//line templates/common.qtpl:22
func streamnavHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, navType string, revisionHash ...string) {
//line templates/common.qtpl:22
	qw422016.N().S(`
`)
//line templates/common.qtpl:24
	u := user.FromRequest(rq).OrAnon()

//line templates/common.qtpl:25
	qw422016.N().S(`
	<nav class="navlinks">
		<ul>
`)
//line templates/common.qtpl:28
	for _, entry := range navEntries {
//line templates/common.qtpl:29
		if navType == "revision" && entry.path == "revision" {
//line templates/common.qtpl:29
			qw422016.N().S(`			<li><b>`)
//line templates/common.qtpl:30
			qw422016.E().S(revisionHash[0])
//line templates/common.qtpl:30
			qw422016.N().S(`</b></li>
`)
//line templates/common.qtpl:31
		} else if navType == entry.path {
//line templates/common.qtpl:31
			qw422016.N().S(`			<li><b>`)
//line templates/common.qtpl:32
			qw422016.E().S(entry.title)
//line templates/common.qtpl:32
			qw422016.N().S(`</b></li>
`)
//line templates/common.qtpl:33
		} else if entry.path != "revision" && u.Group.CanAccessRoute(entry.path) {
//line templates/common.qtpl:33
			qw422016.N().S(`			<li><a href="/`)
//line templates/common.qtpl:34
			qw422016.E().S(entry.path)
//line templates/common.qtpl:34
			qw422016.N().S(`/`)
//line templates/common.qtpl:34
			qw422016.E().S(hyphaName)
//line templates/common.qtpl:34
			qw422016.N().S(`">`)
//line templates/common.qtpl:34
			qw422016.E().S(entry.title)
//line templates/common.qtpl:34
			qw422016.N().S(`</a></li>
`)
//line templates/common.qtpl:35
		}
//line templates/common.qtpl:36
	}
//line templates/common.qtpl:36
	qw422016.N().S(`		`)
//line templates/common.qtpl:37
	qw422016.N().S(userMenuHTML(u))
//line templates/common.qtpl:37
	qw422016.N().S(`
		</ul>
	</nav>
`)
//line templates/common.qtpl:40
}

//line templates/common.qtpl:40
func writenavHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, navType string, revisionHash ...string) {
//line templates/common.qtpl:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/common.qtpl:40
	streamnavHTML(qw422016, rq, hyphaName, navType, revisionHash...)
//line templates/common.qtpl:40
	qt422016.ReleaseWriter(qw422016)
//line templates/common.qtpl:40
}

//line templates/common.qtpl:40
func navHTML(rq *http.Request, hyphaName, navType string, revisionHash ...string) string {
//line templates/common.qtpl:40
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/common.qtpl:40
	writenavHTML(qb422016, rq, hyphaName, navType, revisionHash...)
//line templates/common.qtpl:40
	qs422016 := string(qb422016.B)
//line templates/common.qtpl:40
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/common.qtpl:40
	return qs422016
//line templates/common.qtpl:40
}

//line templates/common.qtpl:42
func streamuserMenuHTML(qw422016 *qt422016.Writer, u *user.User) {
//line templates/common.qtpl:42
	qw422016.N().S(`
	`)
//line templates/common.qtpl:43
	if user.AuthUsed {
//line templates/common.qtpl:43
		qw422016.N().S(`
			<li class="navlinks__user">
			`)
//line templates/common.qtpl:45
		if u.Group == user.UserAnon {
//line templates/common.qtpl:45
			qw422016.N().S(`
				<a href="/login">Login</a>
			`)
//line templates/common.qtpl:47
		} else {
//line templates/common.qtpl:47
			qw422016.N().S(`
				<a href="/page/`)
//line templates/common.qtpl:48
			qw422016.E().S(util.UserTree)
//line templates/common.qtpl:48
			qw422016.N().S(`/`)
//line templates/common.qtpl:48
			qw422016.E().S(u.Name)
//line templates/common.qtpl:48
			qw422016.N().S(`">`)
//line templates/common.qtpl:48
			qw422016.E().S(u.Name)
//line templates/common.qtpl:48
			qw422016.N().S(`</a>
			`)
//line templates/common.qtpl:49
		}
//line templates/common.qtpl:49
		qw422016.N().S(`
			</li>
	`)
//line templates/common.qtpl:51
	}
//line templates/common.qtpl:51
	qw422016.N().S(`
`)
//line templates/common.qtpl:52
}

//line templates/common.qtpl:52
func writeuserMenuHTML(qq422016 qtio422016.Writer, u *user.User) {
//line templates/common.qtpl:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/common.qtpl:52
	streamuserMenuHTML(qw422016, u)
//line templates/common.qtpl:52
	qt422016.ReleaseWriter(qw422016)
//line templates/common.qtpl:52
}

//line templates/common.qtpl:52
func userMenuHTML(u *user.User) string {
//line templates/common.qtpl:52
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/common.qtpl:52
	writeuserMenuHTML(qb422016, u)
//line templates/common.qtpl:52
	qs422016 := string(qb422016.B)
//line templates/common.qtpl:52
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/common.qtpl:52
	return qs422016
//line templates/common.qtpl:52
}
