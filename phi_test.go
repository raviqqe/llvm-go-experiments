package experiments

import (
	"fmt"
	"testing"

	"github.com/llvm-mirror/llvm/bindings/go/llvm"
)

func TestPhi(t *testing.T) {
	m := llvm.NewModule("foo")

	f := llvm.AddFunction(
		m,
		"foo",
		llvm.FunctionType(llvm.Int8Type(), nil, false),
	)

	b := llvm.NewBuilder()
	b.SetInsertPointAtEnd(llvm.AddBasicBlock(f, ""))

	p := llvm.AddBasicBlock(f, "")
	b.SetInsertPointAtEnd(p)
	v := b.CreatePHI(llvm.Int8Type(), "")
	b.CreateRet(v)

	for i := 0; i < 3; i++ {
		bb := llvm.AddBasicBlock(f, fmt.Sprintf("case-%v", i))
		b.SetInsertPointAtEnd(bb)
		b.CreateBr(p)

		v.AddIncoming(
			[]llvm.Value{llvm.ConstInt(llvm.Int8Type(), uint64(i), false)},
			[]llvm.BasicBlock{bb},
		)
	}

	f.Dump()
}
