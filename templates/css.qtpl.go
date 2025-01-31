// Code generated by qtc from "css.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/css.qtpl:1
package templates

//line templates/css.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/css.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/css.qtpl:1
func StreamDefaultCSS(qw422016 *qt422016.Writer) {
//line templates/css.qtpl:1
	qw422016.N().S(`
@media screen and (min-width: 700px) {
	main {margin: 0 auto; width: 700px;}
}
@media screen and (max-width: 700px) {
	main {margin: 0; width: 100%;}
}
*, *::before, *::after {box-sizing: border-box;}
html {height:100%; padding:0; background-color:#ddd;
background-image: url("data:image/svg+xml,%3Csvg width='42' height='44' viewBox='0 0 42 44' xmlns='http://www.w3.org/2000/svg'%3E%3Cg id='Page-1' fill='none' fill-rule='evenodd'%3E%3Cg id='brick-wall' fill='%23bbbbbb' fill-opacity='0.4'%3E%3Cpath d='M0 0h42v44H0V0zm1 1h40v20H1V1zM0 23h20v20H0V23zm22 0h20v20H22V23z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");} /* heropatterns.com */
body {height:100%; margin:0; font-size:16px; font-family:sans-serif;}
main {padding:1rem; background-color: white; box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.2); }
main > form {margin-bottom:1rem;}
textarea {font-size:15px;}
.edit {height:100%;}
.edit-form {height:90%;}
.edit-form textarea {width:100%;height:90%;}

main h1:not(.navi-title) {font-size:1.7rem;}
blockquote {border-left: 4px black solid; margin-left: 0; padding-left: 1rem;}
.wikilink_new {color:#a55858;}
.wikilink_new:visited {color:#a55858;}
.wikilink_external::after {content:"🌐"; margin-left: .5rem; font-size: small; text-decoration: none; align: bottom;}
article code {background-color:#eee; padding: .1rem .3rem; border-radius: .25rem; font-size: 90%; }
article pre.codeblock {background-color:#eee; padding:.5rem; white-space: pre-wrap; border-radius: .25rem;}
.codeblock code {padding:0; font-size:15px;}
.transclusion code, .transclusion .codeblock {background-color:#ddd;}
.transclusion {background-color:#eee; border-radius: .25rem; }
.transclusion__content > *:not(.binary-container) {margin: 0.5rem; }
.transclusion__link {display: block; text-align: right; font-style: italic; margin-top: 0.5rem; color: black; text-decoration: none;}
.transclusion__link::before {content: "⇐ ";}

.binary-container_with-img img,
.binary-container_with-video video,
.binary-container_with-audio audio {width: 100%}
.navi-title a {text-decoration:none;}
.img-gallery { text-align: center; margin-top: .25rem; }
.img-gallery_many-images { background-color: #eee; border-radius: .25rem; padding: .5rem; }
.img-gallery img { max-width: 100%; max-height: 50vh; }
figure { margin: 0; }
figcaption { padding-bottom: .5rem; }

nav ul {display:flex; padding-left:0; flex-wrap:wrap; margin-top:0;}
nav ul li {list-style-type:none;margin-right:1rem;}

#new-name {width:100%;}
.navlinks__user {font-style:italic;}

.rc-entry { display: grid; list-style-type: none; padding: .25rem; background-color: #eee; grid-template-columns: 1fr 1fr; }
.rc-entry__time { font-style: italic; }
.rc-entry__hash { font-style: italic; text-align: right; }
.rc-entry__links { grid-column: 1 / span 2; }
.rc-entry__author { font-style: italic; }
`)
//line templates/css.qtpl:54
}

//line templates/css.qtpl:54
func WriteDefaultCSS(qq422016 qtio422016.Writer) {
//line templates/css.qtpl:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/css.qtpl:54
	StreamDefaultCSS(qw422016)
//line templates/css.qtpl:54
	qt422016.ReleaseWriter(qw422016)
//line templates/css.qtpl:54
}

//line templates/css.qtpl:54
func DefaultCSS() string {
//line templates/css.qtpl:54
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/css.qtpl:54
	WriteDefaultCSS(qb422016)
//line templates/css.qtpl:54
	qs422016 := string(qb422016.B)
//line templates/css.qtpl:54
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/css.qtpl:54
	return qs422016
//line templates/css.qtpl:54
}
