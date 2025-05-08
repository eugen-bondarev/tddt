package util

import "strings"

func SanitizeEndpoint(endpoint string) string {
	return strings.ReplaceAll(strings.ReplaceAll(endpoint, "http://", ""), "https://", "")
}

func GetBucketName(path string) (string, string) {
	parts := strings.Split(path, "/")
	return parts[0], strings.Join(parts[1:], "/")
}
