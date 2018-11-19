package experiments

import (
	"testing"

	"github.com/k0kubun/pp"
	"llvm.org/llvm/bindings/go/llvm"
)

var caseFunctionType = llvm.PointerType(
	llvm.FunctionType(
		llvm.PointerType(llvm.Int8Type(), 0),
		[]llvm.Type{llvm.PointerType(llvm.Int8Type(), 0)},
		false,
	),
	0,
)

func TestVectoredReturnsWithPointersToResults(t *testing.T) {
	m := llvm.NewModule("foo")

	f := llvm.AddFunction(
		m,
		"constructor",
		llvm.FunctionType(
			llvm.PointerType(llvm.Int8Type(), 0),
			[]llvm.Type{
				llvm.PointerType(llvm.Int8Type(), 0),
				llvm.ArrayType(caseFunctionType, 1),
			},
			false,
		),
	)
	f.FirstParam().SetName("environment")

	b := llvm.NewBuilder()
	b.SetInsertPointAtEnd(llvm.AddBasicBlock(f, ""))

	b.CreateRet(
		b.CreateCall(b.CreateExtractValue(f.Params()[1], 0, ""), []llvm.Value{f.FirstParam()}, ""),
	)

	llvm.VerifyFunction(f, llvm.AbortProcessAction)

	f = llvm.AddFunction(
		m,
		"case",
		llvm.FunctionType(
			llvm.PointerType(llvm.Int8Type(), 0),
			[]llvm.Type{llvm.PointerType(llvm.Int8Type(), 0)},
			false,
		),
	)

	b = llvm.NewBuilder()
	b.SetInsertPointAtEnd(llvm.AddBasicBlock(f, ""))

	b.CreateRet(f.FirstParam())

	llvm.VerifyFunction(f, llvm.AbortProcessAction)

	f = llvm.AddFunction(
		m,
		"switch",
		llvm.FunctionType(
			llvm.PointerType(llvm.Int8Type(), 0),
			[]llvm.Type{llvm.PointerType(llvm.Int8Type(), 0)},
			false,
		),
	)

	b = llvm.NewBuilder()
	b.SetInsertPointAtEnd(llvm.AddBasicBlock(f, ""))

	b.CreateRet(
		b.CreateCall(
			m.NamedFunction("constructor"),
			[]llvm.Value{
				f.FirstParam(),
				b.CreateInsertValue(
					llvm.Undef(llvm.ArrayType(caseFunctionType, 1)),
					m.NamedFunction("case"),
					0,
					"",
				),
			},
			"",
		),
	)

	llvm.VerifyFunction(f, llvm.AbortProcessAction)

	m.Dump()

	pp.Println(optimizeModule(m))

	m.Dump()
}
