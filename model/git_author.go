package model

import (
	"fmt"
	"time"
)

type GitAuthor struct {
	Name      string
	Email     string
	Timestamp time.Time
}

var DEFAULT_EMAIL string = "john.doe@example.com"
var DEFAULT_NAME string = "John Doe"

func (a *GitAuthor) Content() string {
	timestamp := a.Timestamp.Unix()
	timezone := a.Timestamp.Format("-0700")
	return fmt.Sprintf("%s <%s> %d %s", a.Name, a.Email, timestamp, timezone)
}
