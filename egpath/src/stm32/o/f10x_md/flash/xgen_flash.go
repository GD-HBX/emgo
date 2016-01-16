package flash

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type FLASH_Periph struct {
	ACR      ACR
	KEYR     KEYR
	OPTKEYR  OPTKEYR
	SR       SR
	CR       CR
	AR       AR
	RESERVED RESERVED
	OBR      OBR
	WRPR     WRPR
}

func (p *FLASH_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var FLASH = (*FLASH_Periph)(unsafe.Pointer(uintptr(mmap.FLASH_R_BASE)))

type ACR_Bits uint32

func (b ACR_Bits) Field(mask ACR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ACR_Bits) J(v int) ACR_Bits {
	return ACR_Bits(bits.Make32(v, uint32(mask)))
}

type ACR struct{ mmio.U32 }

func (r *ACR) Bits(mask ACR_Bits) ACR_Bits { return ACR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ACR) StoreBits(mask, b ACR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ACR) SetBits(mask ACR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ACR) ClearBits(mask ACR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ACR) Load() ACR_Bits              { return ACR_Bits(r.U32.Load()) }
func (r *ACR) Store(b ACR_Bits)            { r.U32.Store(uint32(b)) }

type ACR_Mask struct{ mmio.UM32 }

func (rm ACR_Mask) Load() ACR_Bits   { return ACR_Bits(rm.UM32.Load()) }
func (rm ACR_Mask) Store(b ACR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) LATENCY() ACR_Mask {
	return ACR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(LATENCY)}}
}

func (p *FLASH_Periph) HLFCYA() ACR_Mask {
	return ACR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(HLFCYA)}}
}

func (p *FLASH_Periph) PRFTBE() ACR_Mask {
	return ACR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(PRFTBE)}}
}

func (p *FLASH_Periph) PRFTBS() ACR_Mask {
	return ACR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(PRFTBS)}}
}

type KEYR_Bits uint32

func (b KEYR_Bits) Field(mask KEYR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask KEYR_Bits) J(v int) KEYR_Bits {
	return KEYR_Bits(bits.Make32(v, uint32(mask)))
}

type KEYR struct{ mmio.U32 }

func (r *KEYR) Bits(mask KEYR_Bits) KEYR_Bits { return KEYR_Bits(r.U32.Bits(uint32(mask))) }
func (r *KEYR) StoreBits(mask, b KEYR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *KEYR) SetBits(mask KEYR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *KEYR) ClearBits(mask KEYR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *KEYR) Load() KEYR_Bits               { return KEYR_Bits(r.U32.Load()) }
func (r *KEYR) Store(b KEYR_Bits)             { r.U32.Store(uint32(b)) }

type KEYR_Mask struct{ mmio.UM32 }

func (rm KEYR_Mask) Load() KEYR_Bits   { return KEYR_Bits(rm.UM32.Load()) }
func (rm KEYR_Mask) Store(b KEYR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) FKEYR() KEYR_Mask {
	return KEYR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(FKEYR)}}
}

type OPTKEYR_Bits uint32

func (b OPTKEYR_Bits) Field(mask OPTKEYR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OPTKEYR_Bits) J(v int) OPTKEYR_Bits {
	return OPTKEYR_Bits(bits.Make32(v, uint32(mask)))
}

type OPTKEYR struct{ mmio.U32 }

func (r *OPTKEYR) Bits(mask OPTKEYR_Bits) OPTKEYR_Bits { return OPTKEYR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OPTKEYR) StoreBits(mask, b OPTKEYR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OPTKEYR) SetBits(mask OPTKEYR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *OPTKEYR) ClearBits(mask OPTKEYR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *OPTKEYR) Load() OPTKEYR_Bits                  { return OPTKEYR_Bits(r.U32.Load()) }
func (r *OPTKEYR) Store(b OPTKEYR_Bits)                { r.U32.Store(uint32(b)) }

type OPTKEYR_Mask struct{ mmio.UM32 }

func (rm OPTKEYR_Mask) Load() OPTKEYR_Bits   { return OPTKEYR_Bits(rm.UM32.Load()) }
func (rm OPTKEYR_Mask) Store(b OPTKEYR_Bits) { rm.UM32.Store(uint32(b)) }

type SR_Bits uint32

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) BSY() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(BSY)}}
}

func (p *FLASH_Periph) PGERR() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PGERR)}}
}

func (p *FLASH_Periph) WRPRTERR() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(WRPRTERR)}}
}

func (p *FLASH_Periph) EOP() SR_Mask {
	return SR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(EOP)}}
}

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) PG() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(PG)}}
}

func (p *FLASH_Periph) PER() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(PER)}}
}

func (p *FLASH_Periph) MER() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(MER)}}
}

func (p *FLASH_Periph) OPTPG() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(OPTPG)}}
}

func (p *FLASH_Periph) OPTER() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(OPTER)}}
}

func (p *FLASH_Periph) STRT() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(STRT)}}
}

func (p *FLASH_Periph) LOCK() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(LOCK)}}
}

func (p *FLASH_Periph) OPTWRE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(OPTWRE)}}
}

func (p *FLASH_Periph) ERRIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(ERRIE)}}
}

func (p *FLASH_Periph) EOPIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(EOPIE)}}
}

type AR_Bits uint32

func (b AR_Bits) Field(mask AR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AR_Bits) J(v int) AR_Bits {
	return AR_Bits(bits.Make32(v, uint32(mask)))
}

type AR struct{ mmio.U32 }

func (r *AR) Bits(mask AR_Bits) AR_Bits { return AR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AR) StoreBits(mask, b AR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AR) SetBits(mask AR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *AR) ClearBits(mask AR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *AR) Load() AR_Bits             { return AR_Bits(r.U32.Load()) }
func (r *AR) Store(b AR_Bits)           { r.U32.Store(uint32(b)) }

type AR_Mask struct{ mmio.UM32 }

func (rm AR_Mask) Load() AR_Bits   { return AR_Bits(rm.UM32.Load()) }
func (rm AR_Mask) Store(b AR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) FAR() AR_Mask {
	return AR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(FAR)}}
}

type RESERVED_Bits uint32

func (b RESERVED_Bits) Field(mask RESERVED_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED_Bits) J(v int) RESERVED_Bits {
	return RESERVED_Bits(bits.Make32(v, uint32(mask)))
}

type RESERVED struct{ mmio.U32 }

func (r *RESERVED) Bits(mask RESERVED_Bits) RESERVED_Bits {
	return RESERVED_Bits(r.U32.Bits(uint32(mask)))
}
func (r *RESERVED) StoreBits(mask, b RESERVED_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RESERVED) SetBits(mask RESERVED_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *RESERVED) ClearBits(mask RESERVED_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *RESERVED) Load() RESERVED_Bits             { return RESERVED_Bits(r.U32.Load()) }
func (r *RESERVED) Store(b RESERVED_Bits)           { r.U32.Store(uint32(b)) }

type RESERVED_Mask struct{ mmio.UM32 }

func (rm RESERVED_Mask) Load() RESERVED_Bits   { return RESERVED_Bits(rm.UM32.Load()) }
func (rm RESERVED_Mask) Store(b RESERVED_Bits) { rm.UM32.Store(uint32(b)) }

type OBR_Bits uint32

func (b OBR_Bits) Field(mask OBR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OBR_Bits) J(v int) OBR_Bits {
	return OBR_Bits(bits.Make32(v, uint32(mask)))
}

type OBR struct{ mmio.U32 }

func (r *OBR) Bits(mask OBR_Bits) OBR_Bits { return OBR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OBR) StoreBits(mask, b OBR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OBR) SetBits(mask OBR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *OBR) ClearBits(mask OBR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *OBR) Load() OBR_Bits              { return OBR_Bits(r.U32.Load()) }
func (r *OBR) Store(b OBR_Bits)            { r.U32.Store(uint32(b)) }

type OBR_Mask struct{ mmio.UM32 }

func (rm OBR_Mask) Load() OBR_Bits   { return OBR_Bits(rm.UM32.Load()) }
func (rm OBR_Mask) Store(b OBR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) OPTERR() OBR_Mask {
	return OBR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(OPTERR)}}
}

func (p *FLASH_Periph) RDPRT() OBR_Mask {
	return OBR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(RDPRT)}}
}

func (p *FLASH_Periph) USER() OBR_Mask {
	return OBR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(USER)}}
}

type WRPR_Bits uint32

func (b WRPR_Bits) Field(mask WRPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRPR_Bits) J(v int) WRPR_Bits {
	return WRPR_Bits(bits.Make32(v, uint32(mask)))
}

type WRPR struct{ mmio.U32 }

func (r *WRPR) Bits(mask WRPR_Bits) WRPR_Bits { return WRPR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRPR) StoreBits(mask, b WRPR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRPR) SetBits(mask WRPR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *WRPR) ClearBits(mask WRPR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *WRPR) Load() WRPR_Bits               { return WRPR_Bits(r.U32.Load()) }
func (r *WRPR) Store(b WRPR_Bits)             { r.U32.Store(uint32(b)) }

type WRPR_Mask struct{ mmio.UM32 }

func (rm WRPR_Mask) Load() WRPR_Bits   { return WRPR_Bits(rm.UM32.Load()) }
func (rm WRPR_Mask) Store(b WRPR_Bits) { rm.UM32.Store(uint32(b)) }
