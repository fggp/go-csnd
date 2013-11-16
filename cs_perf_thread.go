package csnd6

/*
#cgo CXXFLAGS: -DUSE_DOUBLE=1

#include <csound/csound.h>
#include "csound_pt.h"
*/
import "C"

import "unsafe"

type CsoundPerformanceThread struct {
	cpt (C.Cpt)
}

func NewCsoundPerformanceThread(csound CSOUND) CsoundPerformanceThread {
	cpt := C.NewCsoundPT(csound.cs)
	return CsoundPerformanceThread{cpt}
}

func (pt CsoundPerformanceThread) Delete() {
	C.DeleteCsoundPT(pt.cpt)
	pt.cpt = nil
}

func (pt CsoundPerformanceThread) IsRunning() bool {
	return C.CsoundPTisRunning(pt.cpt) != 0
}

func (pt CsoundPerformanceThread) GetCsound() *C.CSOUND {
	return C.CsoundPTgetCsound(pt.cpt)
}

func (pt CsoundPerformanceThread) GetStatus() int {
	return int(C.CsoundPTgetStatus(pt.cpt))
}

func (pt CsoundPerformanceThread) Play() {
	C.CsoundPTplay(pt.cpt)
}

func (pt CsoundPerformanceThread) Pause() {
	C.CsoundPTpause(pt.cpt)
}

func (pt CsoundPerformanceThread) TogglePause() {
	C.CsoundPTtogglePause(pt.cpt)
}

func (pt CsoundPerformanceThread) Stop() {
	C.CsoundPTstop(pt.cpt)
}

func (pt CsoundPerformanceThread) ScoreEvent(absp2mode bool, opcod byte, p []MYFLT) {
	C.CsoundPTscoreEvent(pt.cpt, cbool(absp2mode), C.char(opcod), C.int(len(p)),
		cpMYFLT(&p[0]))
}

func (pt CsoundPerformanceThread) InputMessage(s string) {
	var cmsg *C.char = C.CString(s)
	defer C.free(unsafe.Pointer(cmsg))
	C.CsoundPTinputMessage(pt.cpt, cmsg)
}

func (pt CsoundPerformanceThread) SetScoreOffsetSeconds(timeVal float64) {
	C.CsoundPTsetScoreOffsetSeconds(pt.cpt, C.double(timeVal))
}

func (pt CsoundPerformanceThread) Join() {
	C.CsoundPTjoin(pt.cpt)
}

func (pt CsoundPerformanceThread) FlushMessageQueue() {
	C.CsoundPTflushMessageQueue(pt.cpt)
}
