package engine

var executeHub = map[string]Executor{}

func init() {
	executeHub[CyclomaticComplexity] = &cyclomaticComplexityExecutor{}
	executeHub[UintTest] = &unitTestExecutor{}
	executeHub[BigFile] = &bigFileExecutor{}
	executeHub[LongFunc] = &longFuncExecutor{}
}
