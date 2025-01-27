package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bouncepaw/mycorrhiza/util"
)

// isCanonicalName checks if the `name` is canonical.
func isCanonicalName(name string) bool {
	return HyphaPattern.MatchString(name)
}

// CanonicalName makes sure the `name` is canonical. A name is canonical if it is lowercase and all spaces are replaced with underscores.
func CanonicalName(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "_"))
}

// naviTitle turns `canonicalName` into html string with each hypha path parts higlighted as links.
// TODO: rework as a template
func naviTitle(canonicalName string) string {
	var (
		html = fmt.Sprintf(`<h1 class="navi-title" id="navi-title">
	<a href="/page/%s">%s</a>`, util.HomePage, util.SiteTitle)
		prevAcc = `/page/`
		parts   = strings.Split(canonicalName, "/")
	)
	for _, part := range parts {
		html += fmt.Sprintf(`
	<span aria-hidden="true">/</span>
	<a href="%s">%s</a>`,
			prevAcc+part,
			strings.Title(part))
		prevAcc += part + "/"
	}
	return html + "</h1>"
}

// HyphaNameFromRq extracts hypha name from http request. You have to also pass the action which is embedded in the url. For url /page/hypha, the action would be "page".
func HyphaNameFromRq(rq *http.Request, action string) string {
	return CanonicalName(strings.TrimPrefix(rq.URL.Path, "/"+action+"/"))
}
