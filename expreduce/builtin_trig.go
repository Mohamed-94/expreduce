package expreduce

import "math/big"
import "math"

func mathFnOneParam(fn func(float64) float64) func(*Expression, *EvalState) Ex {
	return (func(this *Expression, es *EvalState) Ex {
		if len(this.Parts) != 2 {
			return this
		}

		flt, ok := this.Parts[1].(*Flt)
		if ok {
			flt64, acc := flt.Val.Float64()
			if acc == big.Exact {
				return NewReal(big.NewFloat(fn(flt64)))
			}
		}
		return this
	})
}

func GetTrigDefinitions() (defs []Definition) {
	defs = append(defs, Definition{
		Name:         "Sin",
		legacyEvalFn: mathFnOneParam(math.Sin),
	})
	defs = append(defs, Definition{
		Name:         "Cos",
		legacyEvalFn: mathFnOneParam(math.Cos),
	})
	defs = append(defs, Definition{
		Name:         "Tan",
		legacyEvalFn: mathFnOneParam(math.Tan),
	})
	defs = append(defs, Definition{
		Name:         "ArcTan",
		legacyEvalFn: mathFnOneParam(math.Atan),
	})
	defs = append(defs, Definition{
		Name:         "Cot",
		legacyEvalFn: mathFnOneParam(func(x float64) float64 {return 1/math.Tan(x)}),
	})
	defs = append(defs, Definition{
		Name:         "TrigExpand",
		OmitDocumentation: true,
	})
	defs = append(defs, Definition{
		Name:         "TrigReduce",
		OmitDocumentation: true,
	})
	return
}
