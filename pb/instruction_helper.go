package pb

func (m *DatasetShard) Name() string { _ = "STUB: not implemented"; return "" }

func (m *DatasetShardLocation) Address() string { _ = "STUB: not implemented"; return "" }

func (m *InstructionSet) InstructionNames() (stepNames []string) {
	_ = "STUB: not implemented"
	return nil
}

func (i *Instruction) SetInputLocations(locations []DataLocation) {
	_ = "STUB: not implemented"
	return
}

func (i *Instruction) SetOutputLocations(locations []DataLocation) {
	_ = "STUB: not implemented"
	return
}

func (i *Instruction) GetName() string { _ = "STUB: not implemented"; return "" }
