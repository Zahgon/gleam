package pb

func (l *Location) URL() string { _ = "STUB: not implemented"; return "" }

// the distance is a relative value, similar to network lantency
func (a *Location) Distance(b *Location) float64 { _ = "STUB: not implemented"; return 0 }

func (a ComputeResource) Minus(b ComputeResource) ComputeResource {
	_ = "STUB: not implemented"
	return *new(ComputeResource)
}

func (a ComputeResource) Plus(b ComputeResource) ComputeResource {
	_ = "STUB: not implemented"
	return *new(ComputeResource)
}

func (a ComputeResource) GreaterThanZero() bool { _ = "STUB: not implemented"; return false }

func (a ComputeResource) IsZero() bool { _ = "STUB: not implemented"; return false }

func (a ComputeResource) Covers(b ComputeResource) bool { _ = "STUB: not implemented"; return false }
