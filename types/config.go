package types

// DirectoryConfig Config for one worker
type DirectoryConfig struct {
	Name           string
	Paths          []string
	IgnorePrefixes []string
	Commands       []string
	BreakOnFail    bool
}

// Config Application config
type Config struct {
	Delay int32
	Watch []DirectoryConfig
}
