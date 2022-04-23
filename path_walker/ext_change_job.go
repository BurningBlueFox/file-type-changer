package path_walker

type Whitelist map[string]bool

type ExtChangeJob struct {
	Path             string
	NewExtensionType string
	Whitelist        Whitelist
}
