package utils

func MapSlice[SourceT any, TargetT any](source []SourceT, mapper func(SourceT) TargetT) []TargetT {
	result := make([]TargetT, len(source))
	for i, e := range source {
		result[i] = mapper(e)
	}
	return result
}
