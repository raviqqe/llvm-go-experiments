package experiments

import (
	"testing"

	"github.com/k0kubun/pp"
	"llvm.org/llvm/bindings/go/llvm"
)

func TestAllocaEncoding(t *testing.T) {
	for _, ts := range [][2]llvm.Type{
		{
			llvm.StructType(
				[]llvm.Type{
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
				},
				false,
			),
			llvm.ArrayType(llvm.Int8Type(), 8),
		},
		{
			llvm.ArrayType(llvm.Int8Type(), 8),
			llvm.StructType(
				[]llvm.Type{
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
				},
				false,
			),
		},
		{
			llvm.ArrayType(llvm.Int8Type(), 8),
			llvm.StructType(
				[]llvm.Type{
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int32Type(),
				},
				false,
			),
		},
		{
			llvm.StructType(
				[]llvm.Type{
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int8Type(),
					llvm.Int32Type(),
				},
				false,
			),
			llvm.ArrayType(llvm.Int8Type(), 8),
		},
	} {
		m := llvm.NewModule("foo")

		f := llvm.AddFunction(
			m,
			"foo",
			llvm.FunctionType(ts[1], []llvm.Type{ts[0]}, false),
		)

		b := llvm.NewBuilder()
		b.SetInsertPointAtEnd(llvm.AddBasicBlock(f, ""))

		v := b.CreateAlloca(ts[0], "")
		b.CreateStore(f.FirstParam(), v)
		b.CreateRet(
			b.CreateLoad(
				b.CreateBitCast(
					v,
					llvm.PointerType(ts[1], 0),
					"",
				),
				"",
			),
		)

		f.Dump()

		llvm.VerifyFunction(f, llvm.AbortProcessAction)

		pp.Println(optimizeModule(m))

		f.Dump()
	}
}
