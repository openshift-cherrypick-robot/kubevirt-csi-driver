package util

type StorageClassEnforcement struct {
	AllowList    []string `yaml:"allowList"`
	AllowAll     bool     `yaml:"allowAll"`
	AllowDefault bool     `yaml:"allowDefault"`
}

// Contains tells whether a contains x.
func Contains(arr []string, val string) bool {
	for _, itrVal := range arr {
		if val == itrVal {
			return true
		}
	}
	return false
}
