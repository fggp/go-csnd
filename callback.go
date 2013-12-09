package csnd6

/*
#include <csound/csound.h>
#include <csound/csound_standard_types.h>

void csoundSetFileOpenCB(CSOUND *csound);
void csoundSetPlayOpenCB(CSOUND *csound);
void csoundSetRtPlayCB(CSOUND *csound);
void csoundSetRecOpenCB(CSOUND *csound);
void csoundSetRtRecordCB(CSOUND *csound);
void csoundSetRtCloseCB(CSOUND *csound);
void csoundSetExternalMidiInOpenCB(CSOUND *csound);
void csoundSetExternalMidiReadCB(CSOUND *csound);
void csoundSetExternalMidiInCloseCB(CSOUND *csound);
void csoundSetExternalMidiOutOpenCB(CSOUND *csound);
void csoundSetExternalMidiWriteCB(CSOUND *csound);
void csoundSetExternalMidiOutCloseCB(CSOUND *csound);
void csoundSetExternalMidiErrorStringCB(CSOUND *csound);
void csoundSetCscoreCB(CSOUND *csound);
void csoundSetInputChannelCB(CSOUND *csound);
void csoundSetOutputChannelCB(CSOUND *csound);
int csoundRegisterSenseEventCB(CSOUND *csound, void *userData, int);
*/
import "C"

import (
	"fmt"
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

type ExternalMidiInOpenHandler func(csound *CSOUND, userData unsafe.Pointer, devName string) int32

var externalMidiInOpen ExternalMidiInOpenHandler

//export goExternalMidiInOpenCB
func goExternalMidiInOpenCB(csound unsafe.Pointer, userData unsafe.Pointer, devName *C.char) int32 {
	if externalMidiInOpen == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	return externalMidiInOpen(&cs, userData, C.GoString(devName))
}

func (csound *CSOUND) SetExternalMidiInOpenCallback(f ExternalMidiInOpenHandler) {
	externalMidiInOpen = f
	C.csoundSetExternalMidiInOpenCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiReadHandler func(csound *CSOUND, userData unsafe.Pointer, buf []uint8) int32

var externalMidiRead ExternalMidiReadHandler

//export goExternalMidiReadCB
func goExternalMidiReadCB(csound unsafe.Pointer, userData unsafe.Pointer, buf *C.uchar, nBytes int32) int32 {
	if externalMidiRead == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	var slice []uint8
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = int(nBytes)
	sliceHeader.Len = int(nBytes)
	sliceHeader.Data = uintptr(unsafe.Pointer(buf))
	return externalMidiRead(&cs, userData, slice)
}

func (csound *CSOUND) SetExternalMidiReadCallback(f ExternalMidiReadHandler) {
	externalMidiRead = f
	C.csoundSetExternalMidiReadCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiInCloseHandler func(csound *CSOUND, userData unsafe.Pointer) int32

var externalMidiInClose ExternalMidiInCloseHandler

//export goExternalMidiInCloseCB
func goExternalMidiInCloseCB(csound unsafe.Pointer, userData unsafe.Pointer) int32 {
	if externalMidiInClose == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	return externalMidiInClose(&cs, userData)
}

func (csound *CSOUND) SetExternalMidiInCloseCallback(f ExternalMidiInCloseHandler) {
	externalMidiInClose = f
	C.csoundSetExternalMidiInCloseCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiOutOpenHandler func(csound *CSOUND, userData unsafe.Pointer, devName string) int32

var externalMidiOutOpen ExternalMidiOutOpenHandler

//export goExternalMidiOutOpenCB
func goExternalMidiOutOpenCB(csound unsafe.Pointer, userData unsafe.Pointer, devName *C.char) int32 {
	if externalMidiOutOpen == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	return externalMidiOutOpen(&cs, userData, C.GoString(devName))
}

func (csound *CSOUND) SetExternalMidiOutOpenCallback(f ExternalMidiOutOpenHandler) {
	externalMidiOutOpen = f
	C.csoundSetExternalMidiOutOpenCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiWriteHandler func(csound *CSOUND, userData unsafe.Pointer, buf []uint8) int32

var externalMidiWrite ExternalMidiWriteHandler

//export goExternalMidiWriteCB
func goExternalMidiWriteCB(csound unsafe.Pointer, userData unsafe.Pointer, buf *C.uchar, nBytes int32) int32 {
	if externalMidiWrite == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	var slice []uint8
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = int(nBytes)
	sliceHeader.Len = int(nBytes)
	sliceHeader.Data = uintptr(unsafe.Pointer(buf))
	return externalMidiWrite(&cs, userData, slice)
}

func (csound *CSOUND) SetExternalMidiWriteCallback(f ExternalMidiWriteHandler) {
	externalMidiWrite = f
	C.csoundSetExternalMidiWriteCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiOutCloseHandler func(csound *CSOUND, userData unsafe.Pointer) int32

var externalMidiOutClose ExternalMidiOutCloseHandler

//export goExternalMidiOutCloseCB
func goExternalMidiOutCloseCB(csound unsafe.Pointer, userData unsafe.Pointer) int32 {
	if externalMidiOutClose == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	return externalMidiOutClose(&cs, userData)
}

func (csound *CSOUND) SetExternalMidiOutCloseCallback(f ExternalMidiOutCloseHandler) {
	externalMidiOutClose = f
	C.csoundSetExternalMidiOutCloseCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiErrorStringHandler func(err int) string

var externalMidiErrorString ExternalMidiErrorStringHandler

//export goExternalMidiErrorStringCB
func goExternalMidiErrorStringCB(err int32) *C.char {
	if externalMidiErrorString == nil {
		return nil
	}
	s := externalMidiErrorString(int(err))
	return C.CString(s)
}

func (csound *CSOUND) SetExternalMidiErrorStringCallback(f ExternalMidiErrorStringHandler) {
	externalMidiErrorString = f
	C.csoundSetExternalMidiErrorStringCB(csound.cs)
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

func (csound *CSOUND) SetCscoreCallback(f CscoreHandler) {
	cscore = f
	C.csoundSetCscoreCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type ChannelHandler func(csound *CSOUND, channelName string,
	channelValue []MYFLT, channelType int)

var inputChannel ChannelHandler

//export goInputChannelCB
func goInputChannelCB(csound unsafe.Pointer, channelName *C.char,
	channelValuePtr unsafe.Pointer, channelType unsafe.Pointer) {
	if inputChannel == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	var length, chnType int
	if channelType == unsafe.Pointer(&C.CS_VAR_TYPE_K) {
		length = 1
		chnType = CSOUND_CONTROL_CHANNEL
	} else if channelType == unsafe.Pointer(&C.CS_VAR_TYPE_A) {
		length = int(C.csoundGetKsmps(cs.cs))
		chnType = CSOUND_AUDIO_CHANNEL
	} else if channelType == unsafe.Pointer(&C.CS_VAR_TYPE_S) {
		length = int(C.csoundGetChannelDatasize(cs.cs, channelName))
		chnType = CSOUND_STRING_CHANNEL
	} else {
		return
	}
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(channelValuePtr)
	inputChannel(&cs, C.GoString(channelName), slice, chnType)
}

func (csound *CSOUND) SetInputChannelCallback(f ChannelHandler) {
	inputChannel = f
	C.csoundSetInputChannelCB(csound.cs)
}

var outputChannel ChannelHandler

//export goOutputChannelCB
func goOutputChannelCB(csound unsafe.Pointer, channelName *C.char,
	channelValuePtr unsafe.Pointer, channelType unsafe.Pointer) {
	if outputChannel == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	var length, chnType int
	if channelType == unsafe.Pointer(&C.CS_VAR_TYPE_K) {
		length = 1
		chnType = CSOUND_CONTROL_CHANNEL
	} else if channelType == unsafe.Pointer(&C.CS_VAR_TYPE_A) {
		length = int(C.csoundGetKsmps(cs.cs))
		chnType = CSOUND_AUDIO_CHANNEL
	} else if channelType == unsafe.Pointer(&C.CS_VAR_TYPE_S) {
		length = int(C.csoundGetChannelDatasize(cs.cs, channelName))
		chnType = CSOUND_STRING_CHANNEL
	} else {
		return
	}
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(channelValuePtr)
	outputChannel(&cs, C.GoString(channelName), slice, chnType)
}

func (csound *CSOUND) SetOutputChannelCallback(f ChannelHandler) {
	outputChannel = f
	C.csoundSetOutputChannelCB(csound.cs)
}

////////////////////////////////////////////////////////////////

type SenseEventHandler func(csound *CSOUND, userData unsafe.Pointer)

const maxNumSenseEvent = 10

var senseEvent []SenseEventHandler = make([]SenseEventHandler, maxNumSenseEvent)
var numSenseEvent int

//export goSenseEventCB
func goSenseEventCB(csound, userData unsafe.Pointer, numFun int32) {
	if senseEvent[numFun] == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	senseEvent[numFun](&cs, userData)
}

func (csound *CSOUND) RegisterSenseEventCallback(f SenseEventHandler,
	userData unsafe.Pointer) (ret int, err error) {
	if numSenseEvent < maxNumSenseEvent {
		ret = int(C.csoundRegisterSenseEventCB(csound.cs, userData, C.int(numSenseEvent)))
		if ret != 0 {
			err = fmt.Errorf("Csound could not register SenseEvent callback %d", numSenseEvent)
		} else {
			senseEvent[numSenseEvent] = f
			numSenseEvent++
		}
		return
	}
	ret = -1
	err = fmt.Errorf("%d SenseEvent callbacks already registered! Max value reached", numSenseEvent)
	return
}

////////////////////////////////////////////////////////////////
