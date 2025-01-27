// history/operations.go
// 	Things related to writing history.
package history

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/bouncepaw/mycorrhiza/user"
	"github.com/bouncepaw/mycorrhiza/util"
)

// gitMutex is used for blocking git operations to avoid clashes.
var gitMutex = sync.Mutex{}

// OpType is the type a history operation has. Callers shall set appropriate optypes when creating history operations.
type OpType int

const (
	TypeNone OpType = iota
	TypeEditText
	TypeEditBinary
	TypeDeleteHypha
	TypeRenameHypha
)

// HistoryOp is an object representing a history operation.
type HistoryOp struct {
	// All errors are appended here.
	Errs    []error
	Type    OpType
	userMsg string
	name    string
	email   string
}

// Operation is a constructor of a history operation.
func Operation(opType OpType) *HistoryOp {
	gitMutex.Lock()
	hop := &HistoryOp{
		Errs:  []error{},
		name:  "anon",
		email: "anon@mycorrhiza",
		Type:  opType,
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

// WithError appends the `err` to the list of errors.
func (hop *HistoryOp) WithError(err error) *HistoryOp {
	hop.Errs = append(hop.Errs, err)
	return hop
}

// WithFilesRemoved git-rm-s all passed `paths`. Paths can be rooted or not. Paths that are empty strings are ignored.
func (hop *HistoryOp) WithFilesRemoved(paths ...string) *HistoryOp {
	args := []string{"rm", "--quiet", "--"}
	for _, path := range paths {
		if path != "" {
			args = append(args, path)
		}
	}
	return hop.gitop(args...)
}

// WithFilesRenamed git-mv-s all passed keys of `pairs` to values of `pairs`. Paths can be rooted ot not. Empty keys are ignored.
func (hop *HistoryOp) WithFilesRenamed(pairs map[string]string) *HistoryOp {
	for from, to := range pairs {
		if from != "" {
			os.MkdirAll(filepath.Dir(to), 0777)
			if err := Rename(from, to); err != nil {
				hop.Errs = append(hop.Errs, err)
			}
		}
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
	gitMutex.Unlock()
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

// WithUser sets a user for the commit.
func (hop *HistoryOp) WithUser(u *user.User) *HistoryOp {
	if u.Group != user.UserAnon {
		hop.name = u.Name
		hop.email = u.Name + "@mycorrhiza"
	}
	return hop
}
