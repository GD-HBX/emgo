// This example tests different ways of coping memory. It also shows how to use
// DMA for memory to memory transfers. In case of STM32F2xx/4xx only DMA2
// supports MTM transfer.
package main

import (
	"delay"
	"fmt"
	"rtos"
	"sync/fence"
	"unsafe"

	"stm32/hal/dma"
	"stm32/hal/irq"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
)

var (
	ch     *dma.Channel
	tce    rtos.EventFlag
)

func init() {
	system.Setup80(0, 0)
	systick.Setup(2e6)

	d := dma.DMA1
	d.EnableClock(true)
	ch = d.Channel(1, 0)
	rtos.IRQ(irq.DMA1_Channel1).Enable()
}

const n = 20 * 1024 / 4

var (
	src = make([]uint32, n)
	dst = make([]uint32, n)
)

func printSpeed(t int64, check bool) {
	t1 := rtos.Nanosec()
	t2 := rtos.Nanosec()
	dt := (t1 - t) - (t2 - t1)
	if check {
		for i := range dst {
			if dst[i] != uint32(i) {
				fmt.Printf(" dst != src\n")
				return
			}
			dst[i] = 0
		}
	}
	fmt.Printf(" %6d kB/s\n", (int64(n*unsafe.Sizeof(dst[0]))*1e6+dt/2)/dt)
}

func copyDMA() {
	tce.Reset(0)
	ch.Setup(dma.MTM | dma.IncP | dma.IncM)
	ch.SetWordSize(unsafe.Sizeof(src[0]), unsafe.Sizeof(dst[0]))
	ch.SetLen(n)
	ch.SetAddrP(unsafe.Pointer(&src[0]))
	ch.SetAddrM(unsafe.Pointer(&dst[0]))
	ch.Clear(dma.EvAll, dma.ErrAll)
	fence.W() // ensure tce.Reset(0) happens before any DMA IRQ
	ch.EnableIRQ(dma.Complete, dma.ErrAll)
	ch.Enable()
	tce.Wait(1, 0)
	if _, err := ch.Status(); err != 0 {
		fmt.Println(err)
	}
}

func main() {
	delay.Millisec(250) // Wait for SWO (press reset if you see nothing).

	for {
		fmt.Printf("Initialize src                        ")
		t := rtos.Nanosec()
		for i := range src {
			src[i] = uint32(i)
		}
		printSpeed(t, false)

		fmt.Printf("for i := range src { dst[i] = src[i] }")
		t = rtos.Nanosec()
		for i := range src {
			dst[i] = src[i]
		}
		printSpeed(t, true)

		fmt.Printf("copy(dst, src)                        ")
		t = rtos.Nanosec()
		copy(dst, src)
		printSpeed(t, true)

		fmt.Printf("DMA                                   ")
		t = rtos.Nanosec()
		copyDMA()
		printSpeed(t, true)
		
		delay.Millisec(2e3)
		fmt.Println()
	}
}

func dmaISR() {
	ch.DisableIRQ(dma.EvAll, dma.ErrAll)
	ch.Disable()
	tce.Signal(1)
}

//emgo:const
//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.DMA1_Channel1: dmaISR,
}
