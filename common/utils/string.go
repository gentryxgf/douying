package utils

import "strings"

func IsInList(list []string, key string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

func JoinStringToPath(localPath, userID, timeStamp, suffix string) string {
	var builder strings.Builder

	l := len(localPath) + len(userID) + len(timeStamp) + len(suffix) + 3
	builder.Grow(l)
	builder.WriteString(localPath)
	builder.WriteString("/" + userID)
	builder.WriteString("_" + timeStamp)
	builder.WriteString("." + suffix)

	return builder.String()
}

func JoinString(a, b string) string {
	var builder strings.Builder

	l := len(a) + len(b)
	builder.Grow(l)
	builder.WriteString(a)
	builder.WriteString(b)

	return builder.String()
}
