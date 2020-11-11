package discord

// Command represents a simple discord
type Command struct {
	Name        string
	Aliases     []string
	Description string
	Usage       string
	Example     string
	IgnoreCase  bool
	Handler     ExecutionHandler
}
