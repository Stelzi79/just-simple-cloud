package types

type CloudFile struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Provider string `json:"provider"`
	Region   string `json:"region"`
}

func (c CloudFile) PrettyPrint() any {
	// panic("unimplemented")
	return nil
}
