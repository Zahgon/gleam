package pb

func (taskGroupStatus *FlowExecutionStatus_TaskGroup) Track(
	execute func(*FlowExecutionStatus_TaskGroup_Execution) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *FlowExecutionStatus) GetDataset(datasetId int32) *FlowExecutionStatus_Dataset {
	_ = "STUB: not implemented"
	return nil
}

func (m *FlowExecutionStatus) GetDatasetShard(datasetId, datasetShardId int32) *FlowExecutionStatus_DatasetShard {
	_ = "STUB: not implemented"
	return nil
}

func (m *FlowExecutionStatus) GetTask(stepId, taskId int32) *FlowExecutionStatus_Task {
	_ = "STUB: not implemented"
	return nil
}

func (m *FlowExecutionStatus) GetStep(stepId int32) *FlowExecutionStatus_Step {
	_ = "STUB: not implemented"
	return nil
}
