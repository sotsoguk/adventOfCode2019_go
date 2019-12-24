package intcode

import (
	"fmt"

	"github.com/adventOfCode2019_go/utils/mathUtils"
)

type VM struct {
	Code       []int64
	Input      []int64
	Output     []int64
	InPtr      int64
	RPtr       int64
	Ptr        int64
	Mode       int64 //1: wairing for input	99: exit // 2: output of desired len
	OrigCode   []int64
	OutputWait int64
}

func (vm *VM) Reset() {
	vm.InPtr = 0
	vm.Mode = 0
	vm.RPtr = 0
	vm.Input = nil
	vm.Input = make([]int64, 0)
	vm.Output = nil
	vm.Output = make([]int64, 0)
	vm.Ptr = 0
	vm.Code = nil
	vm.Code = make([]int64, mathUtils.Max(int64(len(vm.OrigCode)), 10000))
	copy(vm.Code, vm.OrigCode)
	// vm.inputLoaded = false
	// vm.blockInput = false
	// vm.enoughInput = false
	vm.OutputWait = 0

}
func (vm *VM) ResetWithoutCode() {
	vm.InPtr = 0
	vm.Mode = 0
	vm.RPtr = 0
	vm.Input = nil
	vm.Input = make([]int64, 0)
	vm.Output = nil
	vm.Output = make([]int64, 0)
	vm.Ptr = 0
}
func (vm *VM) LoadCode(code []int64) {
	vm.OrigCode = make([]int64, len(code))
	copy(vm.OrigCode, code)
}

// func runCode(code *[]int64, input *[]int64) []int64 {
func (vm *VM) LoadInput(input int64) {
	vm.Input = append(vm.Input, input)
	vm.Mode = 0
}

func (vm *VM) LoadInputs(inputs []int64) {
	vm.Input = append(vm.Input, inputs...)
	vm.Mode = 0
}
func (vm *VM) ClearOuput() {
	vm.Output = nil
	vm.Output = make([]int64, 0)
	vm.Mode = 0
}
func (vm *VM) RunCode() {
	// reverseMemory := int64(2000)
	// code := make([]int64, mathUtils.Max(int64(len(origCode)), reverseMemory))
	// copy(code, origCode)
	// result := make([]int64, 0)

	// var ptr, inPtr, rPtr int64
	// fmt.Println(len(vm.input))
	running := true
	printOutput := false
	vm.Output = make([]int64, 0)

	for running {
		op := mathUtils.MakeDigits(vm.Code[vm.Ptr])
		var offset3 int64
		m1, m2, m3 := op[2], op[1], op[0]
		if m3 == 2 {
			offset3 = vm.RPtr
		}
		var opcode, p1, p2 int64
		opcode = 10*op[3] + op[4]
		if opcode == 1 || opcode == 2 || (opcode >= 4 && opcode <= 9) {
			if m1 == 0 {
				p1 = vm.Code[vm.Code[vm.Ptr+1]]
			} else if m1 == 1 {
				p1 = vm.Code[vm.Ptr+1]
			} else if m1 == 2 {
				p1 = vm.Code[vm.Code[vm.Ptr+1]+vm.RPtr]
			}
		}
		if opcode == 1 || opcode == 2 || (opcode > 4 && opcode <= 8) {
			if m2 == 0 {
				p2 = vm.Code[vm.Code[vm.Ptr+2]]
			} else if m2 == 1 {
				p2 = vm.Code[vm.Ptr+2]
			} else if m2 == 2 {
				p2 = vm.Code[vm.Code[vm.Ptr+2]+vm.RPtr]
			}

		}
		switch opcode {
		case 1:

			vm.Code[vm.Code[vm.Ptr+3]+offset3] = p1 + p2
			vm.Ptr += 4
		case 2:

			vm.Code[vm.Code[vm.Ptr+3]+offset3] = p1 * p2
			vm.Ptr += 4
		case 3:
			// fmt.Println("3:before:", vm.InPtr, len(vm.Input))
			if vm.InPtr >= int64(len(vm.Input)) {
				running = false
				vm.Mode = 1
				break
			} else {
				// fmt.Println("3:input:", vm.Input)
				if m1 == 0 {
					vm.Code[vm.Code[vm.Ptr+1]+offset3] = vm.Input[vm.InPtr]
				} else if m1 == 2 {
					vm.Code[vm.Code[vm.Ptr+1]+vm.RPtr] = vm.Input[vm.InPtr]
				}
				vm.Ptr += 2
				vm.Input = vm.Input[1:]
				// vm.InPtr
			}
		case 4:

			if printOutput {
				fmt.Print(p1, ",")
			}
			vm.Output = append(vm.Output, p1)
			vm.Ptr += 2
			if vm.OutputWait > 0 {
				if len(vm.Output) == int(vm.OutputWait) {
					running = false
					vm.Mode = 3
					break
				}
			}
		case 5:
			if p1 != 0 {
				vm.Ptr = p2
			} else {
				vm.Ptr += 3
			}
		case 6:
			if p1 == 0 {
				vm.Ptr = p2
			} else {
				vm.Ptr += 3
			}
		case 7:
			if p1 < p2 {
				vm.Code[vm.Code[vm.Ptr+3]+offset3] = 1
			} else {
				vm.Code[vm.Code[vm.Ptr+3]+offset3] = 0
			}
			vm.Ptr += 4
		case 8:
			if p1 == p2 {
				vm.Code[vm.Code[vm.Ptr+3]+offset3] = 1
			} else {
				vm.Code[vm.Code[vm.Ptr+3]+offset3] = 0
			}
			vm.Ptr += 4
		case 9:
			vm.RPtr += p1
			vm.Ptr += 2
		case 99:
			running = false
			vm.Mode = 99
			break
		default:
			fmt.Println("Unknown Opcode:", opcode)
			running = false
			vm.Mode = 99
			break
		}

	}
	if printOutput {
		fmt.Println()
	}
	// return result
}
