package actionlint

import "strings"

func isLocalUsesSpec(spec string) bool {
	return strings.HasPrefix(spec, "./") || strings.HasPrefix(spec, "$/")
}

func normalizeLocalUsesSpec(spec string) string {
	if strings.HasPrefix(spec, "$/") {
		return "./" + strings.TrimPrefix(spec, "$/")
	}
	return spec
}

func isLocalActionUsesSpec(spec string) bool {
	if strings.HasPrefix(spec, "$/") {
		return !strings.Contains(strings.TrimPrefix(spec, "$/"), "@")
	}
	return strings.HasPrefix(spec, "./")
}
