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

// Return a new CsoundPerformanceThread object.
func NewCsoundPerformanceThread(csound CSOUND) CsoundPerformanceThread {
	cpt := C.NewCsoundPT(csound.cs)
	return CsoundPerformanceThread{cpt}
}

// Free the memory associated with the underlying C++ object.
func (pt CsoundPerformanceThread) Delete() {
	C.DeleteCsoundPT(pt.cpt)
	pt.cpt = nil
}

////////////////////////////////////////////////////////////////

type PTprocessHandler func(cbData unsafe.Pointer)

var ptProcess PTprocessHandler

// Return the process callback as a PTprocessHandler.
func (pt CsoundPerformanceThread) ProcessCallback() PTprocessHandler {
	return ptProcess
}

//export goPTprocessCB
func goPTprocessCB(cbData unsafe.Pointer) {
	if ptProcess == nil {
		return
	}
	ptProcess(cbData)
}

// Set the process callback.
func (pt CsoundPerformanceThread) SetProcessCallback(f PTprocessHandler, cbData unsafe.Pointer) {
	ptProcess = f
	C.CsoundPTsetProcessCB(pt.cpt, cbData)
}

////////////////////////////////////////////////////////////////

// Tell if performance thread is running.
func (pt CsoundPerformanceThread) IsRunning() bool {
	return C.CsoundPTisRunning(pt.cpt) != 0
}

// Return the Csound instance pointer.
func (pt CsoundPerformanceThread) GetCsound() *C.CSOUND {
	return C.CsoundPTgetCsound(pt.cpt)
}

// Return the current status, zero if still playing, positive if
// the end of score was reached or performance was stopped, and
// negative if an error occured.
func (pt CsoundPerformanceThread) GetStatus() int {
	return int(C.CsoundPTgetStatus(pt.cpt))
}

// Continue performance if it was paused.
func (pt CsoundPerformanceThread) Play() {
	C.CsoundPTplay(pt.cpt)
}

// Pause performance (can be continued by calling Play()).
func (pt CsoundPerformanceThread) Pause() {
	C.CsoundPTpause(pt.cpt)
}

// Pause performance unless it is already paused, in which case
// it is continued.
func (pt CsoundPerformanceThread) TogglePause() {
	C.CsoundPTtogglePause(pt.cpt)
}

// Stop performance (cannot be continued).
func (pt CsoundPerformanceThread) Stop() {
	C.CsoundPTstop(pt.cpt)
}

// Send a score event of type 'opcod' (e.g. 'i' for a note event), with
// p-fields in array 'p' (p[0] is p1). If absp2mode is true,
// the start time of the event is measured from the beginning of
// performance, instead of the default of relative to the current time.
func (pt CsoundPerformanceThread) ScoreEvent(absp2mode bool, opcod byte, p []MYFLT) {
	C.CsoundPTscoreEvent(pt.cpt, cbool(absp2mode), C.char(opcod), C.int(len(p)),
		cpMYFLT(&p[0]))
}

// Send a score event as a string, similarly to line events (-L).
func (pt CsoundPerformanceThread) InputMessage(s string) {
	var cmsg *C.char = C.CString(s)
	defer C.free(unsafe.Pointer(cmsg))
	C.CsoundPTinputMessage(pt.cpt, cmsg)
}

// Set the playback time pointer to the specified value (in seconds).
func (pt CsoundPerformanceThread) SetScoreOffsetSeconds(timeVal float64) {
	C.CsoundPTsetScoreOffsetSeconds(pt.cpt, C.double(timeVal))
}

// Wait until the performance is finished or fails, and return a
// positive value if the end of score was reached or Stop() was called,
// and a negative value if an error occured. Also releases any resources
// associated with the performance thread object.
func (pt CsoundPerformanceThread) Join() int {
	return int(C.CsoundPTjoin(pt.cpt))
}

// Wait until all pending messages (pause, send score event, etc.)
// are actually received by the performance thread.
func (pt CsoundPerformanceThread) FlushMessageQueue() {
	C.CsoundPTflushMessageQueue(pt.cpt)
}
