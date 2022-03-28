package scanner

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/fatih/color"
	"github.com/k1LoW/gh-grep/gh"
	"github.com/k1LoW/gh-grep/internal"
)

var (
	matchc    = color.New(color.FgRed, color.Bold)
	delimiter = color.New(color.FgCyan).Sprint(":")
)

type RepoOnlyError struct{}

func (e *RepoOnlyError) Error() string {
	return "repository name already displayed"
}

type Opts struct {
	Patterns     []*regexp.Regexp
	Owner        string
	Repo         string
	Include      string
	Exclude      string
	LineNumber   bool
	NameOnly     bool
	RepoOnly     bool
	URL          bool
	Count        bool
	OnlyMatching bool
	Gh           *gh.Gh
}

func Scan(ctx context.Context, fsys fs.FS, w io.Writer, opts *Opts) error {
	return doublestar.GlobWalk(fsys, opts.Include, func(path string, d fs.DirEntry) error {
		if d.IsDir() {
			return nil
		}
		if opts.Exclude != "" {
			match, err := doublestar.PathMatch(opts.Exclude, path)
			if err != nil {
				return err
			}
			if match {
				log.Printf("Exclude %s\n", path)
				return nil
			}
		}
		log.Printf("Search %s\n", path)
		f, err := fsys.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		// TODO: detect encoding
		r := bufio.NewReader(f)
		n := 1
		c := 0
		for {
			str, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			line := strings.TrimRight(str, "\n")

			matches := [][]int{}
			for _, re := range opts.Patterns {
				matches = append(matches, re.FindAllStringIndex(line, -1)...)
			}

			sort.Slice(matches, func(i, j int) bool {
				return matches[i][0] < matches[j][0]
			})

			// TODO: refactor
			i := 0
			f := [][]int{}
			for {
				if i+1 == len(matches) {
					f = append(f, matches[i])
					break
				} else if i+1 > len(matches) {
					break
				}
				if matches[i][1] < matches[i+1][0] {
					f = append(f, matches[i])
				} else if matches[i][1] >= matches[i+1][0] {
					f = append(f, matches[i])
					i += 1
				}
				i += 1
			}
			matches = f

			if len(matches) > 0 {

				// --count
				if opts.Count {
					c += 1
					continue
				}

				// --repo-only
				if opts.RepoOnly {
					if _, err := fmt.Fprintf(w, "%s/%s\n", opts.Owner, opts.Repo); err != nil {
						return err
					}
					return new(RepoOnlyError)
				}

				// --name-only
				if opts.NameOnly {
					if _, err := fmt.Fprintf(w, "%s/%s%s%s\n", opts.Owner, opts.Repo, delimiter, path); err != nil {
						return err
					}
					break
				}

				// --url
				if opts.URL {
					u, err := opts.Gh.ContentURL(ctx, opts.Owner, opts.Repo, path)
					if err != nil {
						return err
					}
					if _, err := fmt.Fprintf(w, "%s#L%d\n", u, n); err != nil {
						return err
					}
					n += 1
					continue
				}

				// --line-number
				if opts.LineNumber {
					if _, err := fmt.Fprintf(w, "%s/%s%s%s%s%d%s%s\n", opts.Owner, opts.Repo, delimiter, path, delimiter, n, delimiter, internal.PrintLine(line, matches, matchc)); err != nil {
						return err
					}
					n += 1
					continue
				}

				// --only-matching
				if opts.OnlyMatching {
					if _, err := fmt.Fprintf(w, "%s/%s%s%s%s%s\n", opts.Owner, opts.Repo, delimiter, path, delimiter, internal.PrintOnlyMatches(line, matches, matchc)); err != nil {
						return err
					}
					n += 1
					continue
				}

				// default
				if _, err := fmt.Fprintf(w, "%s/%s%s%s%s%s\n", opts.Owner, opts.Repo, delimiter, path, delimiter, internal.PrintLine(line, matches, matchc)); err != nil {
					return err
				}
			}
			n += 1
		}

		// --count
		if opts.Count && c > 0 {
			if _, err := fmt.Fprintf(w, "%s/%s%s%s%s%d\n", opts.Owner, opts.Repo, delimiter, path, delimiter, c); err != nil {
				return err
			}
		}

		return nil
	})
}
