package can

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type CAN_FilterRegister_Periph struct {
	FR1 FR1
	FR2 FR2
}

func (p *CAN_FilterRegister_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type FR1_Bits uint32

func (b FR1_Bits) Field(mask FR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FR1_Bits) J(v int) FR1_Bits {
	return FR1_Bits(bits.Make32(v, uint32(mask)))
}

type FR1 struct{ mmio.U32 }

func (r *FR1) Bits(mask FR1_Bits) FR1_Bits { return FR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *FR1) StoreBits(mask, b FR1_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FR1) SetBits(mask FR1_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *FR1) ClearBits(mask FR1_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *FR1) Load() FR1_Bits              { return FR1_Bits(r.U32.Load()) }
func (r *FR1) Store(b FR1_Bits)            { r.U32.Store(uint32(b)) }

type FR1_Mask struct{ mmio.UM32 }

func (rm FR1_Mask) Load() FR1_Bits   { return FR1_Bits(rm.UM32.Load()) }
func (rm FR1_Mask) Store(b FR1_Bits) { rm.UM32.Store(uint32(b)) }

type FR2_Bits uint32

func (b FR2_Bits) Field(mask FR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FR2_Bits) J(v int) FR2_Bits {
	return FR2_Bits(bits.Make32(v, uint32(mask)))
}

type FR2 struct{ mmio.U32 }

func (r *FR2) Bits(mask FR2_Bits) FR2_Bits { return FR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *FR2) StoreBits(mask, b FR2_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FR2) SetBits(mask FR2_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *FR2) ClearBits(mask FR2_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *FR2) Load() FR2_Bits              { return FR2_Bits(r.U32.Load()) }
func (r *FR2) Store(b FR2_Bits)            { r.U32.Store(uint32(b)) }

type FR2_Mask struct{ mmio.UM32 }

func (rm FR2_Mask) Load() FR2_Bits   { return FR2_Bits(rm.UM32.Load()) }
func (rm FR2_Mask) Store(b FR2_Bits) { rm.UM32.Store(uint32(b)) }
