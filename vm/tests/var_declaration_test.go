package tests

import (
	"github.com/thanhhuy5902/Flux_lang/shared"
	"github.com/thanhhuy5902/Flux_lang/vm"
	"testing"
)

func TestVM_VarDeclaration(t *testing.T) {
	t.Run("create invalid number variables", func(t *testing.T) {
		myVM := vm.NewFluxVirtualMachine()
		result := myVM.Execute(&shared.ExecutionParams{
			EntryPoint: "var_declaration_test.flux",
		})

		if len(result.ErrorCollector.GetErrors()) != 0 {
			t.Errorf(result.ErrorCollector.GetErrors()[0].MessageFmt, result.ErrorCollector.GetErrors()[0].Args...)
		}

		if result.ElapsedTime > 1000 {
			t.Errorf("Execution time too long: %v", result.ElapsedTime)
		}
	})
	t.Run("create valid number variables", func(t *testing.T) {
		myVM := vm.NewFluxVirtualMachine()
		result := myVM.Execute(&shared.ExecutionParams{
			SourceCode: `
				num c {(5+3)/(4+4)}
				num d {c*2/c}
			`,
		})
		if len(result.ErrorCollector.GetErrors()) != 0 {
			t.Errorf(result.ErrorCollector.GetErrors()[0].MessageFmt, result.ErrorCollector.GetErrors()[0].Args...)
		}
		if result.RuntimeException != nil {
			t.Errorf(result.RuntimeException.MessageFmt, result.RuntimeException.Args...)
		}
		if result.ElapsedTime > 1000 {
			t.Errorf("Execution time too long: %v", result.ElapsedTime)
		}

		if d, err := myVM.GetVarTable().GetNumValue("d"); err != nil || d != 2 {
			t.Errorf("Expected 2, got %v", d)
		}
	})
	t.Run("create valid num var and assign to the other num var", func(t *testing.T) {
		myVM := vm.NewFluxVirtualMachine()
		result := myVM.Execute(&shared.ExecutionParams{
			SourceCode: `
				num a {2}
				num b {a}
			`,
		})
		if len(result.ErrorCollector.GetErrors()) != 0 {
			t.Errorf(result.ErrorCollector.GetErrors()[0].MessageFmt, result.ErrorCollector.GetErrors()[0].Args...)
		}
		if result.RuntimeException != nil {
			t.Errorf(result.RuntimeException.MessageFmt, result.RuntimeException.Args...)
		}
		if result.ElapsedTime > 1000 {
			t.Errorf("Execution time too long: %v", result.ElapsedTime)
		}
		if a, err := myVM.GetVarTable().GetNumValue("a"); err != nil || a != 2 {
			t.Errorf("Expected 2, got %v", a)
		}
		if b, err := myVM.GetVarTable().GetNumValue("b"); err != nil || b != 2 {
			t.Errorf("Expected 2, got %v", b)
		}
	})
	t.Run("create valid text variables", func(t *testing.T) {
		myVM := vm.NewFluxVirtualMachine()
		result := myVM.Execute(&shared.ExecutionParams{
			SourceCode: `
				text c {'hello'}
				text d {'hello'+'hehe'+'world'}
				text e {c+d}
			
				
			`,
		})
		if len(result.ErrorCollector.GetErrors()) != 0 {
			t.Errorf(result.ErrorCollector.GetErrors()[0].MessageFmt, result.ErrorCollector.GetErrors()[0].Args...)
		}
		if result.RuntimeException != nil {
			t.Errorf(result.RuntimeException.MessageFmt, result.RuntimeException.Args...)
		}
		if result.ElapsedTime > 1000 {
			t.Errorf("Execution time too long: %v", result.ElapsedTime)
		}

		if c, err := myVM.GetVarTable().GetTextValue("c"); err != nil || c != "hello" {
			t.Errorf("Expected hello, got %v", c)
		}
		if d, err := myVM.GetVarTable().GetTextValue("d"); err != nil || d != "helloworld" {
			t.Errorf("Expected helloworld, got %v", d)
		}
		if e, err := myVM.GetVarTable().GetTextValue("e"); err != nil || e != "hellohelloworld1" {
			t.Errorf("Expected hellohelloworld, got %v", e)
		}

	})
}
