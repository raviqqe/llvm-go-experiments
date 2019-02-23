package experiments

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestAtomicOps(t *testing.T) {
	m := llvm.NewModule("foo")

	f := llvm.AddFunction(m, "foo", llvm.FunctionType(llvm.Void(), nil, false))

	b := llvm.NewBuilder()
	b.SetInsertPointAtEnd(llvm.AddBasicBlock(f, ""))

	f.Dump()

	llvm.VerifyFunction(f, llvm.AbortProcessAction)
}
