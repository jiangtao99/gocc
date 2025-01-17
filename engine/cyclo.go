package engine

import (
	"fmt"
	"regexp"

	"github.com/fzipp/gocyclo"
)

type cyclomaticComplexityExecutor struct {
}

func (c cyclomaticComplexityExecutor) buildDetails(stats gocyclo.Stats) map[string]interface{} {

	details := map[string]interface{}{}

	for _, stat := range stats {
		// example: engine (bigFileExecutor).Compute engine/big_file.go:72:1
		key := fmt.Sprintf("%s %s %s", stat.PkgName, stat.FuncName, stat.Pos)
		details[key] = stat.Complexity
	}
	return details

}

func (c cyclomaticComplexityExecutor) Compute(param Parameter, config Config) Summary {

	var re *regexp.Regexp
	var err error
	if len(config.Cyclo.IgnoreRegx) != 0 {
		re, err = regex(config.Cyclo.IgnoreRegx)
		if err != nil {
			return Summary{Name: CyclomaticComplexity, Err: err}
		}
	}

	stats := gocyclo.Analyze(param.Path, re)

	details := c.buildDetails(stats)
	summary := Summary{
		Name:    CyclomaticComplexity,
		Value:   round(stats.AverageComplexity()),
		Details: details,
	}

	return summary

}
