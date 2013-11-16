package csnd6

/*
#include <csound/csound.h>

void csoundSetFileOpenCB(CSOUND *csound);
void csoundSetPlayOpenCB(CSOUND *csound);
void csoundSetRtPlayCB(CSOUND *csound);
void csoundSetRecOpenCB(CSOUND *csound);
void csoundSetRtRecordCB(CSOUND *csound);
void csoundSetRtCloseCB(CSOUND *csound);
void csoundSetCscoreCB(CSOUND *csound);
*/
import "C"

import (
	"reflect"
	"unsafe"
)

type CsRtAudioParams struct {
	devName      string  // device name (NULL/empty: default)
	devNum       int32   // device number (0-1023), 1024: default
	bufSamp_SW   uint32  // buffer fragment size (-b) in sample frames
	bufSamp_HW   int32   // total buffer size (-B) in sample frames
	nChannels    int32   // number of channels
	sampleFormat int32   // sample format (AE_SHORT etc.)
	sampleRate   float32 // sample rate in Hz
}

////////////////////////////////////////////////////////////////

type FileOpenHandler func(csound *CSOUND, pathName string, fileType int, write, temp bool)

var fileOpen FileOpenHandler

//export goFileOpenCB
func goFileOpenCB(csound unsafe.Pointer, pathName *C.char, fileType, write, temp C.int) {
	if fileOpen == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	fileOpen(&cs, C.GoString(pathName), int(fileType), write != 0, temp != 0)
}

func (csound *CSOUND) SetFileOpenCallback(f FileOpenHandler) {
	fileOpen = f
	C.csoundSetFileOpenCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type PlayOpenHandler func(csound *CSOUND, parm *CsRtAudioParams) int32

var playOpen PlayOpenHandler

//export goPlayOpenCB
func goPlayOpenCB(csound unsafe.Pointer, parm *C.csRtAudioParams) int32 {
	if playOpen == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	goParm := &CsRtAudioParams{
		devName:      C.GoString(parm.devName),
		devNum:       int32(parm.devNum),
		bufSamp_SW:   uint32(parm.bufSamp_SW),
		bufSamp_HW:   int32(parm.bufSamp_HW),
		nChannels:    int32(parm.nChannels),
		sampleFormat: int32(parm.sampleFormat),
		sampleRate:   float32(parm.sampleRate),
	}
	return playOpen(&cs, goParm)
}

func (csound *CSOUND) SetPlayOpenCallback(f PlayOpenHandler) {
	playOpen = f
	C.csoundSetPlayOpenCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type RtPlayHandler func(csound *CSOUND, outBuf []MYFLT)

var rtPlay RtPlayHandler

//export goRtPlayCB
func goRtPlayCB(csound unsafe.Pointer, outBuf *C.MYFLT, length int32) {
	if rtPlay == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(outBuf))
	rtPlay(&cs, slice)
}

func (csound *CSOUND) SetRtPlayCallback(f RtPlayHandler) {
	rtPlay = f
	C.csoundSetRtPlayCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type RecOpenHandler func(csound *CSOUND, parm *CsRtAudioParams) int32

var recOpen RecOpenHandler

//export goRecOpenCB
func goRecOpenCB(csound unsafe.Pointer, parm *C.csRtAudioParams) int32 {
	if recOpen == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	goParm := &CsRtAudioParams{
		devName:      C.GoString(parm.devName),
		devNum:       int32(parm.devNum),
		bufSamp_SW:   uint32(parm.bufSamp_SW),
		bufSamp_HW:   int32(parm.bufSamp_HW),
		nChannels:    int32(parm.nChannels),
		sampleFormat: int32(parm.sampleFormat),
		sampleRate:   float32(parm.sampleRate),
	}
	return recOpen(&cs, goParm)
}

func (csound *CSOUND) SetRecOpenCallback(f RecOpenHandler) {
	recOpen = f
	C.csoundSetRecOpenCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type RtRecordHandler func(csound *CSOUND, inBuf []MYFLT) int32

var rtRecord RtRecordHandler

//export goRtRecordCB
func goRtRecordCB(csound unsafe.Pointer, inBuf *C.MYFLT, length int32) int32 {
	if rtRecord == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(inBuf))
	return rtRecord(&cs, slice)
}

func (csound *CSOUND) SetRtRecordCallback(f RtRecordHandler) {
	rtRecord = f
	C.csoundSetRtRecordCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type RtCloseHandler func(csound *CSOUND)

var rtClose RtCloseHandler

//export goRtCloseCB
func goRtCloseCB(csound unsafe.Pointer) {
	if rtClose == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	rtClose(&cs)
}

func (csound *CSOUND) SetRtCloseCallback(f RtCloseHandler) {
	rtClose = f
	C.csoundSetRtCloseCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type CscoreHandler func(csound *CSOUND)

var cscore CscoreHandler

//export goCscoreCB
func goCscoreCB(csound unsafe.Pointer) {
	if cscore == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	cscore(&cs)
}

func (csound *CSOUND) SetCscoreCallbak(f CscoreHandler) {
	cscore = f
	C.csoundSetCscoreCB(csound.cs)
}

////////////////////////////////////////////////////////////////
