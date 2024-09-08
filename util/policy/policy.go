package policy

import "github.com/microcosm-cc/bluemonday"

var policy *bluemonday.Policy
var OkElements = []string{"b", "i", "u", "s", "br"}

func New() {
    policy = bluemonday.StrictPolicy()

    policy.AllowElements(OkElements...)
}

func Sanitize(s string) string {
    return policy.Sanitize(s)
}
