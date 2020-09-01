// history/operations.go
// 	Things related to writing history.
package history

import (
	"fmt"

	"github.com/bouncepaw/mycorrhiza/util"
)

// OpType is the type a history operation has. Callers shall set appropriate optypes when creating history operations.
type OpType int

const (
	TypeNone OpType = iota
	TypeEditText
	TypeEditBinary
)

// HistoryOp is an object representing a history operation.
type HistoryOp struct {
	// All errors are appended here.
	Errs    []error
	opType  OpType
	userMsg string
	name    string
	email   string
}

// Operation is a constructor of a history operation.
func Operation(opType OpType) *HistoryOp {
	hop := &HistoryOp{
		Errs:   []error{},
		opType: opType,
	}
	return hop
}

// git operation maker helper
func (hop *HistoryOp) gitop(args ...string) *HistoryOp {
	out, err := gitsh(args...)
	if err != nil {
		fmt.Println("out:", out.String())
		hop.Errs = append(hop.Errs, err)
	}
	return hop
}

// WithFiles stages all passed `paths`. Paths can be rooted or not.
func (hop *HistoryOp) WithFiles(paths ...string) *HistoryOp {
	for i, path := range paths {
		paths[i] = util.ShorterPath(path)
	}
	// 1 git operation is more effective than n operations.
	return hop.gitop(append([]string{"add"}, paths...)...)
}

// Apply applies history operation by doing the commit.
func (hop *HistoryOp) Apply() *HistoryOp {
	hop.gitop(
		"commit",
		"--author='"+hop.name+" <"+hop.email+">'",
		"--message="+hop.userMsg,
	)
	return hop
}

// WithMsg sets what message will be used for the future commit. If user message exceeds one line, it is stripped down.
func (hop *HistoryOp) WithMsg(userMsg string) *HistoryOp {
	for _, ch := range userMsg {
		if ch == '\r' || ch == '\n' {
			break
		}
		hop.userMsg += string(ch)
	}
	return hop
}

// WithSignature sets a signature for the future commit. You need to pass a username only, the rest is upon us (including email and time).
func (hop *HistoryOp) WithSignature(username string) *HistoryOp {
	hop.name = username
	hop.email = username + "@mycorrhiza" // A fake email, why not
	return hop
}