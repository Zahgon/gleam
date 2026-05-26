// shell.go defines how a Shell script should be executed on agents.
package script

type Operation struct {
	Type string
	Code string
}

type ShellScript struct {
	initCode   string
	env        []string
	operations []*Operation
}

func NewShellScript() *ShellScript { _ = "STUB: not implemented"; return nil }

func (c *ShellScript) Init(code string) { _ = "STUB: not implemented"; return }

func (c *ShellScript) Name() string { _ = "STUB: not implemented"; return "" }

func (c *ShellScript) GetCommand() *Command { _ = "STUB: not implemented"; return nil }

func (c *ShellScript) Pipe(code string) *ShellScript { _ = "STUB: not implemented"; return nil }
