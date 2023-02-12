/*
Copyright © 2021 Ken'ichiro Oyama <k1lowxb@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/k1LoW/gh-grep/gh"
	"github.com/k1LoW/gh-grep/scanner"
	"github.com/k1LoW/gh-grep/version"
	"github.com/k1LoW/ghfs"
	"github.com/mattn/go-colorable"
	"github.com/spf13/cobra"
)

var (
	opts       scanner.Opts
	patterns   []string
	repos      []string
	ignoreCase bool
	branch     string
	tag        string
	stdout     = colorable.NewColorableStdout()
	stderr     = colorable.NewColorableStderr()
)

var rootCmd = &cobra.Command{
	Use:          "gh-grep [PATTERN]",
	Short:        "Print lines matching a pattern in repositories using GitHub API",
	Long:         `Print lines matching a pattern in repositories using GitHub API`,
	Version:      version.Version,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		if len(patterns) == 0 && len(args) > 0 {
			patterns = []string{args[0]}
		}

		opts.Patterns = []*regexp.Regexp{}
		for _, p := range patterns {
			if ignoreCase {
				p = "(?i)" + p
			}
			re, err := regexp.Compile(p)
			if err != nil {
				return err
			}
			opts.Patterns = append(opts.Patterns, re)
		}

		g, err := gh.New()
		if err != nil {
			return err
		}
		opts.Gh = g
		if len(repos) == 0 {
			repos, err = g.Repositories(ctx, opts.Owner)
			if err != nil {
				return err
			}
		}

		fsOpts := []ghfs.Option{ghfs.Client(g.Client())}
		if branch != "" {
			fsOpts = append(fsOpts, ghfs.Branch(branch))
		}
		if tag != "" {
			fsOpts = append(fsOpts, ghfs.Tag(tag))
		}

		for _, repo := range repos {
			log.Printf("In %s/%s\n", opts.Owner, repo)
			fsys, err := ghfs.New(opts.Owner, repo, fsOpts...)
			if err != nil {
				return err
			}
			opts.Repo = repo
			if err := scanner.Scan(ctx, fsys, stdout, &opts); err != nil {
				if errors.Is(err, &scanner.RepoOnlyError{}) {
					continue
				} else {
					return err
				}
			}
		}
		return nil
	},
}

func Execute() {
	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	log.SetOutput(io.Discard)
	if env := os.Getenv("DEBUG"); env != "" {
		log.SetOutput(os.Stderr)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&opts.Owner, "owner", "", "", "repository owner or org")
	if err := rootCmd.MarkFlagRequired("owner"); err != nil {
		panic(err)
	}
	rootCmd.Flags().StringSliceVarP(&repos, "repo", "", []string{}, "repository name")
	rootCmd.Flags().StringVarP(&branch, "branch", "", "", "branch name")
	rootCmd.Flags().StringVarP(&tag, "tag", "", "", "tag name")
	rootCmd.Flags().StringVarP(&opts.Include, "include", "", "**/*", "search only files that match pattern")
	rootCmd.Flags().StringVarP(&opts.Exclude, "exclude", "", "", "skip files and directories matching pattern")
	rootCmd.Flags().BoolVarP(&opts.LineNumber, "line-number", "n", false, "show line numbers")
	rootCmd.Flags().BoolVarP(&ignoreCase, "ignore-case", "i", false, "case insensitive matching")
	rootCmd.Flags().BoolVarP(&opts.NameOnly, "name-only", "", false, "show only repository:filenames")
	rootCmd.Flags().BoolVarP(&opts.RepoOnly, "repo-only", "", false, "show only repositories")
	rootCmd.Flags().BoolVarP(&opts.URL, "url", "", false, "show URL")
	rootCmd.Flags().BoolVarP(&opts.Count, "count", "c", false, "show the number of matches instead of matching lines")
	rootCmd.Flags().BoolVarP(&opts.OnlyMatching, "only-matching", "o", false, "show only matching parts of a line")
	rootCmd.Flags().StringSliceVarP(&patterns, "", "e", []string{}, "match pattern")
}
