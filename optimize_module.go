package experiments

import (
	"llvm.org/llvm/bindings/go/llvm"
)

func optimizeModule(m llvm.Module) bool {
	b := llvm.NewPassManagerBuilder()
	b.SetOptLevel(int(llvm.CodeGenLevelAggressive))

	p := llvm.NewPassManager()
	b.PopulateFunc(p)
	b.Populate(p)

	return p.Run(m)
}
