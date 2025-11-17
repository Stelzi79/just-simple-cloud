package types

type CloudType int

const (
	local CloudType = iota
)

type CloudFile struct {
	Name     string    `json:"name"`
	Type     CloudType `json:"type"`
	Provider string    `json:"provider"`
	Region   string    `json:"region"`
}

func (c CloudFile) PrettyPrint() any {
	// panic("unimplemented")
	return nil
}
