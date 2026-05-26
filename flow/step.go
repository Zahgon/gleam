package flow

import (
	"github.com/chrislusf/gleam/instruction"
	"github.com/chrislusf/gleam/script"
)

func (fc *Flow) NewStep() (step *Step) { _ = "STUB: not implemented"; return nil }

func (step *Step) NewTask() (task *Task) { _ = "STUB: not implemented"; return nil }

func (step *Step) SetInstruction(prefix string, ins instruction.Instruction) {
	_ = "STUB: not implemented"
	return
}

func (step *Step) RunFunction(task *Task) error { _ = "STUB: not implemented"; return nil }

func (step *Step) GetScriptCommand() *script.Command { _ = "STUB: not implemented"; return nil }
