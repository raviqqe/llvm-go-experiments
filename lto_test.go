package experiments

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestLTO(t *testing.T) {
	m := createModule("foo")

	llvm.LinkModules(m, createModule("bar"))

	m.Dump()

	optimizeWithLTO(m)

	m.Dump()
}

func createModule(funcName string) llvm.Module {
	m := llvm.NewModule("")

	f := llvm.AddFunction(m, "util", llvm.FunctionType(llvm.Int8Type(), nil, false))
	f.SetLinkage(llvm.PrivateLinkage)
	b := llvm.NewBuilder()
	b.SetInsertPointAtEnd(llvm.AddBasicBlock(f, ""))
	b.CreateRet(llvm.ConstInt(llvm.Int8Type(), 42, false))

	ff := llvm.AddFunction(m, funcName, llvm.FunctionType(llvm.Int8Type(), nil, false))
	b = llvm.NewBuilder()
	b.SetInsertPointAtEnd(llvm.AddBasicBlock(ff, ""))
	b.CreateRet(b.CreateCall(f, nil, ""))

	return m
}

func optimizeWithLTO(m llvm.Module) {
	b := llvm.NewPassManagerBuilder()
	b.SetOptLevel(3)
	b.SetSizeLevel(2)

	p := llvm.NewPassManager()
	b.PopulateFunc(p)
	b.Populate(p)
	// b.PopulateLTOPassManager(p, true, true)

	p.Run(m)
}
