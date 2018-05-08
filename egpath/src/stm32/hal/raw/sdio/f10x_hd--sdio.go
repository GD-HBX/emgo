// +build f10x_hd

// Peripheral: SDIO_Periph  SD host Interface.
// Instances:
//  SDIO  mmap.SDIO_BASE
// Registers:
//  0x00 32  POWER
//  0x04 32  CLKCR
//  0x08 32  ARG
//  0x0C 32  CMD
//  0x10 32  RESPCMD
//  0x14 32  RESP[4] Response registers.
//  0x24 32  DTIMER
//  0x28 32  DLEN
//  0x2C 32  DCTRL
//  0x30 32  DCOUNT
//  0x34 32  STA
//  0x38 32  ICR
//  0x3C 32  MASK
//  0x48 32  FIFOCNT
//  0x80 32  FIFO
// Import:
//  stm32/o/f10x_hd/mmap
package sdio

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PWRCTRL   POWER = 0x03 << 0 //+ PWRCTRL[1:0] bits (Power supply control bits).
	PWRCTRL_0 POWER = 0x01 << 0 //  Bit 0.
	PWRCTRL_1 POWER = 0x02 << 0 //  Bit 1.
)

const (
	PWRCTRLn = 0
)

const (
	CLKDIV   CLKCR = 0xFF << 0  //+ Clock divide factor.
	CLKEN    CLKCR = 0x01 << 8  //+ Clock enable bit.
	PWRSAV   CLKCR = 0x01 << 9  //+ Power saving configuration bit.
	BYPASS   CLKCR = 0x01 << 10 //+ Clock divider bypass enable bit.
	WIDBUS   CLKCR = 0x03 << 11 //+ WIDBUS[1:0] bits (Wide bus mode enable bit).
	WIDBUS_0 CLKCR = 0x01 << 11 //  Bit 0.
	WIDBUS_1 CLKCR = 0x02 << 11 //  Bit 1.
	NEGEDGE  CLKCR = 0x01 << 13 //+ SDIO_CK dephasing selection bit.
	HWFC_EN  CLKCR = 0x01 << 14 //+ HW Flow Control enable.
)

const (
	CLKDIVn  = 0
	CLKENn   = 8
	PWRSAVn  = 9
	BYPASSn  = 10
	WIDBUSn  = 11
	NEGEDGEn = 13
	HWFC_ENn = 14
)

const (
	CMDARG ARG = 0xFFFFFFFF << 0 //+ Command argument.
)

const (
	CMDARGn = 0
)

const (
	CMDINDEX    CMD = 0x3F << 0  //+ Command Index.
	WAITRESP    CMD = 0x03 << 6  //+ WAITRESP[1:0] bits (Wait for response bits).
	WAITRESP_0  CMD = 0x01 << 6  //  Bit 0.
	WAITRESP_1  CMD = 0x02 << 6  //  Bit 1.
	WAITINT     CMD = 0x01 << 8  //+ CPSM Waits for Interrupt Request.
	WAITPEND    CMD = 0x01 << 9  //+ CPSM Waits for ends of data transfer (CmdPend internal signal).
	CPSMEN      CMD = 0x01 << 10 //+ Command path state machine (CPSM) Enable bit.
	SDIOSUSPEND CMD = 0x01 << 11 //+ SD I/O suspend command.
	ENCMDCOMPL  CMD = 0x01 << 12 //+ Enable CMD completion.
	NIEN        CMD = 0x01 << 13 //+ Not Interrupt Enable.
	CEATACMD    CMD = 0x01 << 14 //+ CE-ATA command.
)

const (
	CMDINDEXn    = 0
	WAITRESPn    = 6
	WAITINTn     = 8
	WAITPENDn    = 9
	CPSMENn      = 10
	SDIOSUSPENDn = 11
	ENCMDCOMPLn  = 12
	NIENn        = 13
	CEATACMDn    = 14
)

const (
	CARDSTATUS1 RESP = 0xFFFFFFFF << 0 //+ Card Status.
)

const (
	CARDSTATUS1n = 0
)

const (
	DATATIME DTIMER = 0xFFFFFFFF << 0 //+ Data timeout period..
)

const (
	DATATIMEn = 0
)

const (
	DATALENGTH DLEN = 0x1FFFFFF << 0 //+ Data length value.
)

const (
	DATALENGTHn = 0
)

const (
	DTEN         DCTRL = 0x01 << 0  //+ Data transfer enabled bit.
	DTDIR        DCTRL = 0x01 << 1  //+ Data transfer direction selection.
	DTMODE       DCTRL = 0x01 << 2  //+ Data transfer mode selection.
	DMAEN        DCTRL = 0x01 << 3  //+ DMA enabled bit.
	DBLOCKSIZE   DCTRL = 0x0F << 4  //+ DBLOCKSIZE[3:0] bits (Data block size).
	DBLOCKSIZE_0 DCTRL = 0x01 << 4  //  Bit 0.
	DBLOCKSIZE_1 DCTRL = 0x02 << 4  //  Bit 1.
	DBLOCKSIZE_2 DCTRL = 0x04 << 4  //  Bit 2.
	DBLOCKSIZE_3 DCTRL = 0x08 << 4  //  Bit 3.
	RWSTART      DCTRL = 0x01 << 8  //+ Read wait start.
	RWSTOP       DCTRL = 0x01 << 9  //+ Read wait stop.
	RWMOD        DCTRL = 0x01 << 10 //+ Read wait mode.
	SDIOEN       DCTRL = 0x01 << 11 //+ SD I/O enable functions.
)

const (
	DTENn       = 0
	DTDIRn      = 1
	DTMODEn     = 2
	DMAENn      = 3
	DBLOCKSIZEn = 4
	RWSTARTn    = 8
	RWSTOPn     = 9
	RWMODn      = 10
	SDIOENn     = 11
)

const (
	DATACOUNT DCOUNT = 0x1FFFFFF << 0 //+ Data count value.
)

const (
	DATACOUNTn = 0
)

const (
	CCRCFAIL STA = 0x01 << 0  //+ Command response received (CRC check failed).
	DCRCFAIL STA = 0x01 << 1  //+ Data block sent/received (CRC check failed).
	CTIMEOUT STA = 0x01 << 2  //+ Command response timeout.
	DTIMEOUT STA = 0x01 << 3  //+ Data timeout.
	TXUNDERR STA = 0x01 << 4  //+ Transmit FIFO underrun error.
	RXOVERR  STA = 0x01 << 5  //+ Received FIFO overrun error.
	CMDREND  STA = 0x01 << 6  //+ Command response received (CRC check passed).
	CMDSENT  STA = 0x01 << 7  //+ Command sent (no response required).
	DATAEND  STA = 0x01 << 8  //+ Data end (data counter, SDIDCOUNT, is zero).
	STBITERR STA = 0x01 << 9  //+ Start bit not detected on all data signals in wide bus mode.
	DBCKEND  STA = 0x01 << 10 //+ Data block sent/received (CRC check passed).
	CMDACT   STA = 0x01 << 11 //+ Command transfer in progress.
	TXACT    STA = 0x01 << 12 //+ Data transmit in progress.
	RXACT    STA = 0x01 << 13 //+ Data receive in progress.
	TXFIFOHE STA = 0x01 << 14 //+ Transmit FIFO Half Empty: at least 8 words can be written into the FIFO.
	RXFIFOHF STA = 0x01 << 15 //+ Receive FIFO Half Full: there are at least 8 words in the FIFO.
	TXFIFOF  STA = 0x01 << 16 //+ Transmit FIFO full.
	RXFIFOF  STA = 0x01 << 17 //+ Receive FIFO full.
	TXFIFOE  STA = 0x01 << 18 //+ Transmit FIFO empty.
	RXFIFOE  STA = 0x01 << 19 //+ Receive FIFO empty.
	TXDAVL   STA = 0x01 << 20 //+ Data available in transmit FIFO.
	RXDAVL   STA = 0x01 << 21 //+ Data available in receive FIFO.
	SDIOIT   STA = 0x01 << 22 //+ SDIO interrupt received.
	CEATAEND STA = 0x01 << 23 //+ CE-ATA command completion signal received for CMD61.
)

const (
	CCRCFAILn = 0
	DCRCFAILn = 1
	CTIMEOUTn = 2
	DTIMEOUTn = 3
	TXUNDERRn = 4
	RXOVERRn  = 5
	CMDRENDn  = 6
	CMDSENTn  = 7
	DATAENDn  = 8
	STBITERRn = 9
	DBCKENDn  = 10
	CMDACTn   = 11
	TXACTn    = 12
	RXACTn    = 13
	TXFIFOHEn = 14
	RXFIFOHFn = 15
	TXFIFOFn  = 16
	RXFIFOFn  = 17
	TXFIFOEn  = 18
	RXFIFOEn  = 19
	TXDAVLn   = 20
	RXDAVLn   = 21
	SDIOITn   = 22
	CEATAENDn = 23
)

const (
	CCRCFAILC ICR = 0x01 << 0  //+ CCRCFAIL flag clear bit.
	DCRCFAILC ICR = 0x01 << 1  //+ DCRCFAIL flag clear bit.
	CTIMEOUTC ICR = 0x01 << 2  //+ CTIMEOUT flag clear bit.
	DTIMEOUTC ICR = 0x01 << 3  //+ DTIMEOUT flag clear bit.
	TXUNDERRC ICR = 0x01 << 4  //+ TXUNDERR flag clear bit.
	RXOVERRC  ICR = 0x01 << 5  //+ RXOVERR flag clear bit.
	CMDRENDC  ICR = 0x01 << 6  //+ CMDREND flag clear bit.
	CMDSENTC  ICR = 0x01 << 7  //+ CMDSENT flag clear bit.
	DATAENDC  ICR = 0x01 << 8  //+ DATAEND flag clear bit.
	STBITERRC ICR = 0x01 << 9  //+ STBITERR flag clear bit.
	DBCKENDC  ICR = 0x01 << 10 //+ DBCKEND flag clear bit.
	SDIOITC   ICR = 0x01 << 22 //+ SDIOIT flag clear bit.
	CEATAENDC ICR = 0x01 << 23 //+ CEATAEND flag clear bit.
)

const (
	CCRCFAILCn = 0
	DCRCFAILCn = 1
	CTIMEOUTCn = 2
	DTIMEOUTCn = 3
	TXUNDERRCn = 4
	RXOVERRCn  = 5
	CMDRENDCn  = 6
	CMDSENTCn  = 7
	DATAENDCn  = 8
	STBITERRCn = 9
	DBCKENDCn  = 10
	SDIOITCn   = 22
	CEATAENDCn = 23
)

const (
	CCRCFAILIE MASK = 0x01 << 0  //+ Command CRC Fail Interrupt Enable.
	DCRCFAILIE MASK = 0x01 << 1  //+ Data CRC Fail Interrupt Enable.
	CTIMEOUTIE MASK = 0x01 << 2  //+ Command TimeOut Interrupt Enable.
	DTIMEOUTIE MASK = 0x01 << 3  //+ Data TimeOut Interrupt Enable.
	TXUNDERRIE MASK = 0x01 << 4  //+ Tx FIFO UnderRun Error Interrupt Enable.
	RXOVERRIE  MASK = 0x01 << 5  //+ Rx FIFO OverRun Error Interrupt Enable.
	CMDRENDIE  MASK = 0x01 << 6  //+ Command Response Received Interrupt Enable.
	CMDSENTIE  MASK = 0x01 << 7  //+ Command Sent Interrupt Enable.
	DATAENDIE  MASK = 0x01 << 8  //+ Data End Interrupt Enable.
	STBITERRIE MASK = 0x01 << 9  //+ Start Bit Error Interrupt Enable.
	DBCKENDIE  MASK = 0x01 << 10 //+ Data Block End Interrupt Enable.
	CMDACTIE   MASK = 0x01 << 11 //+ Command Acting Interrupt Enable.
	TXACTIE    MASK = 0x01 << 12 //+ Data Transmit Acting Interrupt Enable.
	RXACTIE    MASK = 0x01 << 13 //+ Data receive acting interrupt enabled.
	TXFIFOHEIE MASK = 0x01 << 14 //+ Tx FIFO Half Empty interrupt Enable.
	RXFIFOHFIE MASK = 0x01 << 15 //+ Rx FIFO Half Full interrupt Enable.
	TXFIFOFIE  MASK = 0x01 << 16 //+ Tx FIFO Full interrupt Enable.
	RXFIFOFIE  MASK = 0x01 << 17 //+ Rx FIFO Full interrupt Enable.
	TXFIFOEIE  MASK = 0x01 << 18 //+ Tx FIFO Empty interrupt Enable.
	RXFIFOEIE  MASK = 0x01 << 19 //+ Rx FIFO Empty interrupt Enable.
	TXDAVLIE   MASK = 0x01 << 20 //+ Data available in Tx FIFO interrupt Enable.
	RXDAVLIE   MASK = 0x01 << 21 //+ Data available in Rx FIFO interrupt Enable.
	SDIOITIE   MASK = 0x01 << 22 //+ SDIO Mode Interrupt Received interrupt Enable.
	CEATAENDIE MASK = 0x01 << 23 //+ CE-ATA command completion signal received Interrupt Enable.
)

const (
	CCRCFAILIEn = 0
	DCRCFAILIEn = 1
	CTIMEOUTIEn = 2
	DTIMEOUTIEn = 3
	TXUNDERRIEn = 4
	RXOVERRIEn  = 5
	CMDRENDIEn  = 6
	CMDSENTIEn  = 7
	DATAENDIEn  = 8
	STBITERRIEn = 9
	DBCKENDIEn  = 10
	CMDACTIEn   = 11
	TXACTIEn    = 12
	RXACTIEn    = 13
	TXFIFOHEIEn = 14
	RXFIFOHFIEn = 15
	TXFIFOFIEn  = 16
	RXFIFOFIEn  = 17
	TXFIFOEIEn  = 18
	RXFIFOEIEn  = 19
	TXDAVLIEn   = 20
	RXDAVLIEn   = 21
	SDIOITIEn   = 22
	CEATAENDIEn = 23
)

const (
	FIFOCOUNT FIFOCNT = 0xFFFFFF << 0 //+ Remaining number of words to be written to or read from the FIFO.
)

const (
	FIFOCOUNTn = 0
)

const (
	FIFODATA FIFO = 0xFFFFFFFF << 0 //+ Receive and transmit FIFO data.
)

const (
	FIFODATAn = 0
)
