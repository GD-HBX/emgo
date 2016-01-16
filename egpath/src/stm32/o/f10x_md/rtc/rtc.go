// Peripheral: RTC_Periph  Real-Time Clock.
// Instances:
//  RTC  mmap.RTC_BASE
// Registers:
//  0x00 16  CRH
//  0x04 16  CRL
//  0x08 16  PRLH
//  0x0C 16  PRLL
//  0x10 16  DIVH
//  0x14 16  DIVL
//  0x18 16  CNTH
//  0x1C 16  CNTL
//  0x20 16  ALRH
//  0x24 16  ALRL
// Import:
//  stm32/o/f10x_md/mmap
package rtc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	SECIE CRH_Bits = 0x01 << 0 //+ Second Interrupt Enable.
	ALRIE CRH_Bits = 0x01 << 1 //+ Alarm Interrupt Enable.
	OWIE  CRH_Bits = 0x01 << 2 //+ OverfloW Interrupt Enable.
)

const (
	SECF  CRL_Bits = 0x01 << 0 //+ Second Flag.
	ALRF  CRL_Bits = 0x01 << 1 //+ Alarm Flag.
	OWF   CRL_Bits = 0x01 << 2 //+ OverfloW Flag.
	RSF   CRL_Bits = 0x01 << 3 //+ Registers Synchronized Flag.
	CNF   CRL_Bits = 0x01 << 4 //+ Configuration Flag.
	RTOFF CRL_Bits = 0x01 << 5 //+ RTC operation OFF.
)
