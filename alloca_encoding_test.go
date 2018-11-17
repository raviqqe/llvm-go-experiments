package main

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestAllocaEncoding(t *testing.T) {
	m := llvm.NewModule("foo")

	f := llvm.AddFunction(
		m,
		"foo",
		llvm.FunctionType(llvm.ArrayType(llvm.Int8Type(), 8), nil, false),
	)

	b := llvm.NewBuilder()
	b.SetInsertPointAtEnd(llvm.AddBasicBlock(f, ""))

	tt := llvm.StructType(
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
	)

	v := b.CreateAlloca(tt, "")
	b.CreateStore(llvm.ConstNull(tt), v)
	b.Create
	b.CreateRetVoid()

	f.Dump()
}
