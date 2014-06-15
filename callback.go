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

// Set an external callback for receiving notices whenever Csound opens
// a file. The callback is made after the file is successfully opened.
// The following information is passed to the callback:
//     string pathname of the file; either full or relative to current dir
//     int    a file type code from the enumeration CSOUND_FILETYPES
//     bool   true if Csound is writing the file, false if reading
//     bool   true if a temporary file that Csound will delete; false if not
//
// Pass nil to disable the callback.
// This callback is retained after a Reset() call.
func (csound *CSOUND) SetFileOpenCallback(f FileOpenHandler) {
	fileOpen = f
	if f == nil {
		C.csoundSetFileOpenCallback(csound.Cs, nil)
		return
	}
	C.csoundSetFileOpenCB(csound.Cs)
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

// Set a function to be called by Csound for opening real-time
// audio playback.
func (csound *CSOUND) SetPlayOpenCallback(f PlayOpenHandler) {
	playOpen = f
	C.csoundSetPlayOpenCB(csound.Cs)
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

// Set a function to be called by Csound for performing real-time
// audio playback.
func (csound *CSOUND) SetRtPlayCallback(f RtPlayHandler) {
	rtPlay = f
	C.csoundSetRtPlayCB(csound.Cs)
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

// Set a function to be called by Csound for opening real-time
// audio recording.
func (csound *CSOUND) SetRecOpenCallback(f RecOpenHandler) {
	recOpen = f
	C.csoundSetRecOpenCB(csound.Cs)
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

// Set a function to be called by Csound for performing real-time
// audio recording.
func (csound *CSOUND) SetRtRecordCallback(f RtRecordHandler) {
	rtRecord = f
	C.csoundSetRtRecordCB(csound.Cs)
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

// Set a function to be called by Csound for closing real-time
// audio playback and recording.
func (csound *CSOUND) SetRtCloseCallback(f RtCloseHandler) {
	rtClose = f
	C.csoundSetRtCloseCB(csound.Cs)
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

// Set callback for opening real time MIDI input.
func (csound *CSOUND) SetExternalMidiInOpenCallback(f ExternalMidiInOpenHandler) {
	externalMidiInOpen = f
	C.csoundSetExternalMidiInOpenCB(csound.Cs)
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

// Set callback for reading from real time MIDI input.
func (csound *CSOUND) SetExternalMidiReadCallback(f ExternalMidiReadHandler) {
	externalMidiRead = f
	C.csoundSetExternalMidiReadCB(csound.Cs)
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

// Set callback for closing real time MIDI input.
func (csound *CSOUND) SetExternalMidiInCloseCallback(f ExternalMidiInCloseHandler) {
	externalMidiInClose = f
	C.csoundSetExternalMidiInCloseCB(csound.Cs)
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

// Set callback for opening real time MIDI output.
func (csound *CSOUND) SetExternalMidiOutOpenCallback(f ExternalMidiOutOpenHandler) {
	externalMidiOutOpen = f
	C.csoundSetExternalMidiOutOpenCB(csound.Cs)
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

// Set callback for writing to real time MIDI output.
func (csound *CSOUND) SetExternalMidiWriteCallback(f ExternalMidiWriteHandler) {
	externalMidiWrite = f
	C.csoundSetExternalMidiWriteCB(csound.Cs)
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

// Set callback for closing real time MIDI output.
func (csound *CSOUND) SetExternalMidiOutCloseCallback(f ExternalMidiOutCloseHandler) {
	externalMidiOutClose = f
	C.csoundSetExternalMidiOutCloseCB(csound.Cs)
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

// Set callback for converting MIDI error codes to strings.
func (csound *CSOUND) SetExternalMidiErrorStringCallback(f ExternalMidiErrorStringHandler) {
	externalMidiErrorString = f
	C.csoundSetExternalMidiErrorStringCB(csound.Cs)
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

// Set an external callback for Cscore processing.
// Pass nil to reset to the internal cscore() function
// (which does nothing).
// This callback is retained after a Reset() call.
func (csound *CSOUND) SetCscoreCallback(f CscoreHandler) {
	cscore = f
	if f == nil {
		C.csoundSetCscoreCallback(csound.Cs, nil)
		return
	}
	C.csoundSetCscoreCB(csound.Cs)
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
		length = int(C.csoundGetKsmps(cs.Cs))
		chnType = CSOUND_AUDIO_CHANNEL
	} else if channelType == unsafe.Pointer(&C.CS_VAR_TYPE_S) {
		length = int(C.csoundGetChannelDatasize(cs.Cs, channelName))
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

// Set the function which will be called whenever the invalue opcode is used.
func (csound *CSOUND) SetInputChannelCallback(f ChannelHandler) {
	inputChannel = f
	C.csoundSetInputChannelCB(csound.Cs)
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
		length = int(C.csoundGetKsmps(cs.Cs))
		chnType = CSOUND_AUDIO_CHANNEL
	} else if channelType == unsafe.Pointer(&C.CS_VAR_TYPE_S) {
		length = int(C.csoundGetChannelDatasize(cs.Cs, channelName))
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

// Set the function which will be called whenever the outvalue opcode is used.
func (csound *CSOUND) SetOutputChannelCallback(f ChannelHandler) {
	outputChannel = f
	C.csoundSetOutputChannelCB(csound.Cs)
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

// Register a function to be called once in every control period
// by sensevents(). Any number of functions may be registered,
// and will be called in the order of registration.
// The callback function takes two arguments: the Csound instance
// pointer, and the userData pointer as passed to this function.
// This facility can be used to ensure a function is called synchronously
// before every csound control buffer processing. It is important
// to make sure no blocking operations are performed in the callback.
// The callbacks are cleared on Cleanup().
// Returns zero on success.
func (csound *CSOUND) RegisterSenseEventCallback(f SenseEventHandler,
	userData unsafe.Pointer) (ret int, err error) {
	if numSenseEvent < maxNumSenseEvent {
		ret = int(C.csoundRegisterSenseEventCB(csound.Cs, userData, C.int(numSenseEvent)))
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
