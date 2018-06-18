package sdmmc

import (
	"sync/fence"
	"unsafe"

	"sdcard"
)

// Driver implements sdcard.Host interface.
type Driver struct {
	driverCommon
	isr  func(*Driver)
	addr uintptr
	n    int
	err  Error
	dtc  DataCtrl
}

// MakeDriver returns initialized SPI driver that uses provided SPI peripheral.
func MakeDriver(p *Periph) Driver {
	return Driver{driverCommon: driverCommon{p: p}}
}

// NewDriver provides convenient way to create heap allocated Driver struct.
func NewDriver(p *Periph) *Driver {
	d := new(Driver)
	*d = MakeDriver(p)
	return d
}

func (d *Driver) Periph() *Periph {
	return d.p
}

func (d *Driver) Err(clear bool) error {
	if d.err == 0 {
		return nil
	}
	var err error
	if d.err == ErrCmdTimeout {
		err = sdcard.ErrCmdTimeout
	} else {
		err = d.err
	}
	if clear {
		d.n = 0
		d.err = 0
	}
	return err
}

// SetBusClock sets SD bus clock frequency (freqhz <= 0 disables clock). If
// pwrsave is true the clock output is automatically disabled when bus is idle.
func (d *Driver) SetClock(freqhz int, pwrsave bool) {
	d.setClock(freqhz, pwrsave)
}

// SetBusWidth sets the SD bus width.
func (d *Driver) SetBusWidth(width sdcard.BusWidth) {
	d.setBusWidth(width)
}

func (d *Driver) ISR() {
	d.isr(d)
}

func (d *Driver) cmdISR() {
	p := d.p
	_, err := p.Status()
	if err != 0 || d.n == 0 {
		p.SetIRQMask(0, 0)
		d.done.Signal(1)
	}
	irqs := DataEnd
	if d.dtc&Recv != 0 {
		d.isr = (*Driver).recvISR
		irqs |= RxHalfFull
	} else {
		irqs |= TxHalfEmpty
		d.isr = (*Driver).sendISR
		p.SetDataCtrl(d.dtc)
	}
	p.SetIRQMask(irqs, ErrAll)
}

func (d *Driver) recvISR() {
	p := d.p
	addr := d.addr
	n := d.n
	for n >= 8 {
		ev, err := p.Status()
		if err != 0 {
			goto done
		}
		if ev&RxHalfFull == 0 {
			goto wait
		}
		addr = burstCopyPTM(p.raw.FIFO.Addr(), addr)
		n -= 8
	}
	ev, err := p.Status()
	if err != 0 {
		goto done
	}
	if ev&DataEnd == 0 {
		goto wait
	}
	for n > 0 {
		*(*uint32)(unsafe.Pointer(addr)) = p.Load()
		*(*uint32)(unsafe.Pointer(addr + 4)) = p.Load()
		addr += 8
		n--
	}
done:
	p.SetIRQMask(0, 0)
	d.done.Signal(1)
wait:
	d.addr = addr
	d.n = n
}

func (d *Driver) sendISR() {
	p := d.p
	addr := d.addr
	n := d.n
	for n >= 8 {
		ev, err := p.Status()
		if err != 0 {
			goto done
		}
		if ev&TxHalfEmpty == 0 {
			goto wait
		}
		addr = burstCopyMTP(addr, p.raw.FIFO.Addr())
		n -= 8
	}
	ev, err := p.Status()
	if err != 0 {
		goto done
	}
	if ev&DataEnd == 0 {
		goto wait
	}
	for n > 0 {
		p.Store(*(*uint32)(unsafe.Pointer(addr)))
		p.Store(*(*uint32)(unsafe.Pointer(addr + 4)))
		addr += 8
		n--
	}
done:
	p.SetIRQMask(0, 0)
	d.done.Signal(1)
wait:
	d.addr = addr
	d.n = n
}

// SendCmd sends the cmd to the card and receives its response, if any. Short
// response is returned in r[0]. Long is returned in r[0:3] (r[0] contains the
// least significant bits, r[3] contains the most significant bits). If preceded
// by SetupData, SendCmd performs the data transfer.
func (d *Driver) SendCmd(cmd sdcard.Command, arg uint32) (r sdcard.Response) {
	if d.err != 0 {
		return
	}
	var waitFor Event
	if cmd&sdcard.HasResp != 0 {
		waitFor = CmdRespOK
	} else {
		waitFor = CmdSent
	}
	d.isr = (*Driver).cmdISR
	d.done.Reset(0)
	p := d.p
	p.Clear(EvAll, ErrAll)
	p.SetArg(arg)
	p.SetCmd(CmdEna | Command(cmd)&255)
	fence.W()                     // Orders writes to normal and IO memory.
	p.SetIRQMask(waitFor, ErrAll) // After SetCmd because of spurious IRQs.
	d.done.Wait(1, 0)
	if _, err := p.Status(); err != 0 {
		if rt := cmd & sdcard.RespType; rt == sdcard.R3 || rt == sdcard.R4 {
			err &^= ErrCmdCRC
		}
		if err != 0 {
			d.err = err
			d.dtc = 0
			return
		}
	}
	if cmd&sdcard.HasResp != 0 {
		if cmd&sdcard.LongResp != 0 {
			r[3] = p.Resp(0) // Most significant bits.
			r[2] = p.Resp(1)
			r[1] = p.Resp(2)
			r[0] = p.Resp(3) // Least significant bits.
		} else {
			r[0] = p.Resp(0)
		}
	}
	if d.dtc == 0 {
		return
	}
	if d.dtc&Stream == 0 {
		// Wait for data CRC.
		for {
			ev, err := p.Status()
			if err != 0 {
				d.err = err
				break
			}
			if ev&DataBlkEnd != 0 {
				break
			}
		}
	}
	d.dtc = 0
}

// SetupData setups the data transfer for subsequent command.
func (d *Driver) SetupData(mode sdcard.DataMode, buf sdcard.Data) {
	if d.err != 0 || len(buf) == 0 {
		return
	}
	d.dtc = DTEna | DataCtrl(mode)
	d.addr = uintptr(unsafe.Pointer(&buf[0]))
	d.n = -len(buf)
	p := d.p
	p.SetDataLen(len(buf) * 8)
	if d.dtc&Recv != 0 {
		p.SetDataCtrl(d.dtc)
	}
}
