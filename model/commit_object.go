package model

import (
	"crypto/sha1"
	"fmt"
)

type CommitObject struct {
	Author    GitAuthor
	Committer GitAuthor
	Git       GitObject
	Tree      string
	Parents   []string
	Message   string
}

func NewCommitObject(author, commiter GitAuthor, tree string, parents []string, message string) *CommitObject {
	c := CommitObject{Author: author, Committer: commiter, Tree: tree, Parents: parents, Message: message, Git: GitObject{}}

	sha := sha1.Sum([]byte(c.FullContent()))
	c.Git.hash = sha

	return &c
}

func (c *CommitObject) Content() string {
	parents := ""
	for _, p := range c.Parents {
		parents += fmt.Sprintf("parent %s\n", p)
	}

	return fmt.Sprintf(
		"tree %s\n%sauthor %s\ncommitter %s\n\n%s\n",
		c.Tree,
		parents,
		c.Author.Content(),
		c.Committer.Content(),
		c.Message)
}

func (c *CommitObject) Length() int {
	return len(c.Content())
}

func (c *CommitObject) FullContent() string {
	return fmt.Sprintf("commit %d\x00%s", c.Length(), c.Content())
}

func (c *CommitObject) Compress() []byte {
	return c.Git.compress(c.FullContent())
}

func (c *CommitObject) Write() error {
	return c.Git.write(c.FullContent())
}
