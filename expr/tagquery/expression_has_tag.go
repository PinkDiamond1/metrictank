package tagquery

import (
	"strings"

	"github.com/raintank/schema"
)

type expressionHasTag struct {
	expressionCommon
}

func (e *expressionHasTag) GetDefaultDecision() FilterDecision {
	return Fail
}

func (e *expressionHasTag) GetOperator() ExpressionOperator {
	return HAS_TAG
}

func (e *expressionHasTag) GetCostMultiplier() uint32 {
	return 2
}

func (e *expressionHasTag) OperatesOnTag() bool {
	return true
}

func (e *expressionHasTag) ValuePasses(value string) bool {
	return value == e.key
}

func (e *expressionHasTag) GetMetricDefinitionFilter(_ IdTagLookup) MetricDefinitionFilter {
	if e.key == "name" {
		return func(_ schema.MKey, _ string, _ []string) FilterDecision { return Pass }
	}

	resultIfTagIsAbsent := None
	if !metaTagSupport {
		resultIfTagIsAbsent = Fail
	}

	matchPrefix := e.key + "="
	return func(_ schema.MKey, _ string, tags []string) FilterDecision {
		for _, tag := range tags {
			if strings.HasPrefix(tag, matchPrefix) {
				return Pass
			}
		}

		return resultIfTagIsAbsent
	}
}

func (e *expressionHasTag) StringIntoBuilder(builder *strings.Builder) {
	builder.WriteString(e.key)
	builder.WriteString("!=")
}