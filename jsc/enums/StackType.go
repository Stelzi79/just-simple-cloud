package enums

import "github.com/orsinium-labs/enum"

type StackType enum.Member[string]

var (
	Folder  = StackType{"folder"}
	GitRepo = StackType{"gitRepo"}

	Default = StackType{"folder"}

	StackTypes = enum.New(Folder, GitRepo, Default)
)

func (st *StackType) UnmarshalText(text []byte) error {
	parsed := StackTypes.Parse(string(text))
	if parsed == nil {
		parsed = &Default
	}
	*st = *parsed

	return nil
}
