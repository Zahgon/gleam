package ui

import (
	"github.com/chrislusf/gleam/pb"
)

var (
	WidthStep       = 20 * m
	HightStep       = 3 * m
	HightStepHeader = 2 * m
	Margin          = 5 * m
	LineLength      = 5 * m
	SmallMargin     = 4
	VerticalGap     = 4 * m
)

type stepGroupPosition struct {
	input  point
	output point
}

func GenSvg(status *pb.FlowExecutionStatus) string { _ = "STUB: not implemented"; return "" }

func doFlowExecutionStatus(canvas *svg.SVG, status *pb.FlowExecutionStatus, width int) (largestWidth, height int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// determine the largest width

// determine input points

func doState(canvas *svg.SVG, input point, state string) (output point) {
	_ = "STUB: not implemented"
	return *new(point)
}

func doStepGroup(canvas *svg.SVG, input point, status *pb.FlowExecutionStatus, stepGroup *pb.FlowExecutionStatus_StepGroup) (output point) {
	_ = "STUB: not implemented"
	return *new(point)
}

func doStep(canvas *svg.SVG, input point, step *pb.FlowExecutionStatus_Step, hasFinished bool) (output point) {
	_ = "STUB: not implemented"
	return *new(point)
}

func getLastStepId(status *pb.FlowExecutionStatus, stepGroup *pb.FlowExecutionStatus_StepGroup) int32 {
	_ = "STUB: not implemented"
	return 0
}

func isStepGroupFinished(status *pb.FlowExecutionStatus, stepGroup *pb.FlowExecutionStatus_StepGroup) bool {
	_ = "STUB: not implemented"
	return false
}

// may be more efficient to map stepId=>size
func collectStepOutputDatasetSize(status *pb.FlowExecutionStatus, stepId int32) (counter int64) {
	_ = "STUB: not implemented"
	return 0
}
