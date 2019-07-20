package csnd

/*
#include <csound/csound.h>
#include <csound/csound_standard_types.h>

void csoundSetFileOpenCB(CSOUND *csound);
void csoundSetPlayOpenCB(CSOUND *csound);
void csoundSetRtPlayCB(CSOUND *csound);
void csoundSetRecOpenCB(CSOUND *csound);
void csoundSetRtRecordCB(CSOUND *csound);
void csoundSetRtCloseCB(CSOUND *csound);
void csoundSetAudioDeviceListCB(CSOUND *csound);
void csoundSetExternalMidiInOpenCB(CSOUND *csound);
void csoundSetExternalMidiReadCB(CSOUND *csound);
void csoundSetExternalMidiInCloseCB(CSOUND *csound);
void csoundSetExternalMidiOutOpenCB(CSOUND *csound);
void csoundSetExternalMidiWriteCB(CSOUND *csound);
void csoundSetExternalMidiOutCloseCB(CSOUND *csound);
void csoundSetExternalMidiErrorStringCB(CSOUND *csound);
void csoundSetMidiDeviceListCB(CSOUND *csound);
void csoundSetCscoreCB(CSOUND *csound);
void csoundSetInputChannelCB(CSOUND *csound);
void csoundSetOutputChannelCB(CSOUND *csound);
int csoundRegisterSenseEventCB(CSOUND *csound, void *userData, int);
void csoundSetMakeGraphCB(CSOUND *csound);
void csoundSetDrawGraphCB(CSOUND *csound);
void csoundSetKillGraphCB(CSOUND *csound);
void csoundSetExitGraphCB(CSOUND *csound);
void csoundSetYieldCB(CSOUND *csound);
*/
import "C"

import (
	"fmt"
	"reflect"
	"sync"
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

type FileOpenHandler func(csound CSOUND, pathName string, fileType int, write, temp bool)

var fileOpenH FileOpenHandler

//export goFileOpenCB
func goFileOpenCB(csound unsafe.Pointer, pathName *C.char, fileType, write, temp C.int) {
	if fileOpenH == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	fileOpenH(cs, C.GoString(pathName), int(fileType), write != 0, temp != 0)
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
func (csound CSOUND) SetFileOpenCallback(f FileOpenHandler) {
	fileOpenH = f
	if f == nil {
		C.csoundSetFileOpenCallback(csound.Cs, nil)
		return
	}
	C.csoundSetFileOpenCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type PlayOpenHandler func(csound CSOUND, parm *CsRtAudioParams) int32

var playOpenH PlayOpenHandler

//export goPlayOpenCB
func goPlayOpenCB(csound unsafe.Pointer, parm *C.csRtAudioParams) int32 {
	if playOpenH == nil {
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
	return playOpenH(cs, goParm)
}

// Set a function to be called by Csound for opening real-time
// audio playback.
func (csound CSOUND) SetPlayOpenCallback(f PlayOpenHandler) {
	playOpenH = f
	C.csoundSetPlayOpenCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type RtPlayHandler func(csound CSOUND, outBuf []MYFLT)

var rtPlayH RtPlayHandler

//export goRtPlayCB
func goRtPlayCB(csound unsafe.Pointer, outBuf *C.MYFLT, length int32) {
	if rtPlayH == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(outBuf))
	rtPlayH(cs, slice)
}

// Set a function to be called by Csound for performing real-time
// audio playback.
func (csound CSOUND) SetRtPlayCallback(f RtPlayHandler) {
	rtPlayH = f
	C.csoundSetRtPlayCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type RecOpenHandler func(csound CSOUND, parm *CsRtAudioParams) int32

var recOpenH RecOpenHandler

//export goRecOpenCB
func goRecOpenCB(csound unsafe.Pointer, parm *C.csRtAudioParams) int32 {
	if recOpenH == nil {
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
	return recOpenH(cs, goParm)
}

// Set a function to be called by Csound for opening real-time
// audio recording.
func (csound CSOUND) SetRecOpenCallback(f RecOpenHandler) {
	recOpenH = f
	C.csoundSetRecOpenCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type RtRecordHandler func(csound CSOUND, inBuf []MYFLT) int32

var rtRecordH RtRecordHandler

//export goRtRecordCB
func goRtRecordCB(csound unsafe.Pointer, inBuf *C.MYFLT, length int32) int32 {
	if rtRecordH == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(inBuf))
	return rtRecordH(cs, slice)
}

// Set a function to be called by Csound for performing real-time
// audio recording.
func (csound CSOUND) SetRtRecordCallback(f RtRecordHandler) {
	rtRecordH = f
	C.csoundSetRtRecordCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type RtCloseHandler func(csound CSOUND)

var rtCloseH RtCloseHandler

//export goRtCloseCB
func goRtCloseCB(csound unsafe.Pointer) {
	if rtCloseH == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	rtCloseH(cs)
}

// Set a function to be called by Csound for closing real-time
// audio playback and recording.
func (csound CSOUND) SetRtCloseCallback(f RtCloseHandler) {
	rtCloseH = f
	C.csoundSetRtCloseCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type AudioDeviceListHandler func(csound CSOUND, list *CsoundAudioDevice, isOutput bool) int

var audioDevListH AudioDeviceListHandler

//export goAudioDeviceListCB
func goAudioDeviceListCB(csound unsafe.Pointer, list unsafe.Pointer, isOutput int32) int32 {
	if audioDevListH == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	return int32(audioDevListH(cs, (*CsoundAudioDevice)(list), isOutput != 0))
}

// Set a function that is called to obtain a list of audio devices.
// This should be set by rtaudio modules and should not be set by hosts.
// (See AudioDevList())
func (csound CSOUND) SetAudioDeviceListCallback(f AudioDeviceListHandler) {
	audioDevListH = f
	C.csoundSetAudioDeviceListCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiInOpenHandler func(csound CSOUND, userData unsafe.Pointer, devName string) int32

var externalMidiInOpenH ExternalMidiInOpenHandler

//export goExternalMidiInOpenCB
func goExternalMidiInOpenCB(csound unsafe.Pointer, userData unsafe.Pointer, devName *C.char) int32 {
	if externalMidiInOpenH == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	return externalMidiInOpenH(cs, userData, C.GoString(devName))
}

// Set callback for opening real time MIDI input.
func (csound CSOUND) SetExternalMidiInOpenCallback(f ExternalMidiInOpenHandler) {
	externalMidiInOpenH = f
	C.csoundSetExternalMidiInOpenCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiReadHandler func(csound CSOUND, userData unsafe.Pointer, buf []uint8) int32

var externalMidiReadH ExternalMidiReadHandler

//export goExternalMidiReadCB
func goExternalMidiReadCB(csound unsafe.Pointer, userData unsafe.Pointer, buf *C.uchar, nBytes int32) int32 {
	if externalMidiReadH == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	var slice []uint8
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = int(nBytes)
	sliceHeader.Len = int(nBytes)
	sliceHeader.Data = uintptr(unsafe.Pointer(buf))
	return externalMidiReadH(cs, userData, slice)
}

// Set callback for reading from real time MIDI input.
func (csound CSOUND) SetExternalMidiReadCallback(f ExternalMidiReadHandler) {
	externalMidiReadH = f
	C.csoundSetExternalMidiReadCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiInCloseHandler func(csound CSOUND, userData unsafe.Pointer) int32

var externalMidiInCloseH ExternalMidiInCloseHandler

//export goExternalMidiInCloseCB
func goExternalMidiInCloseCB(csound unsafe.Pointer, userData unsafe.Pointer) int32 {
	if externalMidiInCloseH == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	return externalMidiInCloseH(cs, userData)
}

// Set callback for closing real time MIDI input.
func (csound CSOUND) SetExternalMidiInCloseCallback(f ExternalMidiInCloseHandler) {
	externalMidiInCloseH = f
	C.csoundSetExternalMidiInCloseCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiOutOpenHandler func(csound CSOUND, userData unsafe.Pointer, devName string) int32

var externalMidiOutOpenH ExternalMidiOutOpenHandler

//export goExternalMidiOutOpenCB
func goExternalMidiOutOpenCB(csound unsafe.Pointer, userData unsafe.Pointer, devName *C.char) int32 {
	if externalMidiOutOpenH == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	return externalMidiOutOpenH(cs, userData, C.GoString(devName))
}

// Set callback for opening real time MIDI output.
func (csound CSOUND) SetExternalMidiOutOpenCallback(f ExternalMidiOutOpenHandler) {
	externalMidiOutOpenH = f
	C.csoundSetExternalMidiOutOpenCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiWriteHandler func(csound CSOUND, userData unsafe.Pointer, buf []uint8) int32

var externalMidiWriteH ExternalMidiWriteHandler

//export goExternalMidiWriteCB
func goExternalMidiWriteCB(csound unsafe.Pointer, userData unsafe.Pointer, buf *C.uchar, nBytes int32) int32 {
	if externalMidiWriteH == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	var slice []uint8
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = int(nBytes)
	sliceHeader.Len = int(nBytes)
	sliceHeader.Data = uintptr(unsafe.Pointer(buf))
	return externalMidiWriteH(cs, userData, slice)
}

// Set callback for writing to real time MIDI output.
func (csound CSOUND) SetExternalMidiWriteCallback(f ExternalMidiWriteHandler) {
	externalMidiWriteH = f
	C.csoundSetExternalMidiWriteCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiOutCloseHandler func(csound CSOUND, userData unsafe.Pointer) int32

var externalMidiOutCloseH ExternalMidiOutCloseHandler

//export goExternalMidiOutCloseCB
func goExternalMidiOutCloseCB(csound unsafe.Pointer, userData unsafe.Pointer) int32 {
	if externalMidiOutCloseH == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	return externalMidiOutCloseH(cs, userData)
}

// Set callback for closing real time MIDI output.
func (csound CSOUND) SetExternalMidiOutCloseCallback(f ExternalMidiOutCloseHandler) {
	externalMidiOutCloseH = f
	C.csoundSetExternalMidiOutCloseCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type ExternalMidiErrorStringHandler func(err int) string

var externalMidiErrorStringH ExternalMidiErrorStringHandler

//export goExternalMidiErrorStringCB
func goExternalMidiErrorStringCB(err int32) *C.char {
	if externalMidiErrorStringH == nil {
		return nil
	}
	s := externalMidiErrorStringH(int(err))
	return C.CString(s)
}

// Set callback for converting MIDI error codes to strings.
func (csound CSOUND) SetExternalMidiErrorStringCallback(f ExternalMidiErrorStringHandler) {
	externalMidiErrorStringH = f
	C.csoundSetExternalMidiErrorStringCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type MidiDeviceListHandler func(csound CSOUND, list *CsoundMidiDevice, isOutput bool) int

var midiDevListH MidiDeviceListHandler

//export goMidiDeviceListCB
func goMidiDeviceListCB(csound unsafe.Pointer, list unsafe.Pointer, isOutput int32) int32 {
	if midiDevListH == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	return int32(midiDevListH(cs, (*CsoundMidiDevice)(list), isOutput != 0))
}

// Set a function that is called to obtain a list of MIDI devices.
// This should be set by IO plugins, and should not be set by hosts.
// (See MidiDevList())
func (csound CSOUND) SetMidiDeviceListCallback(f MidiDeviceListHandler) {
	midiDevListH = f
	C.csoundSetMidiDeviceListCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type CscoreHandler func(csound CSOUND)

var cscoreH CscoreHandler

//export goCscoreCB
func goCscoreCB(csound unsafe.Pointer) {
	if cscoreH == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	cscoreH(cs)
}

// Set an external callback for Cscore processing.
// Pass nil to reset to the internal cscore() function
// (which does nothing).
// This callback is retained after a Reset() call.
func (csound CSOUND) SetCscoreCallback(f CscoreHandler) {
	cscoreH = f
	if f == nil {
		C.csoundSetCscoreCallback(csound.Cs, nil)
		return
	}
	C.csoundSetCscoreCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

// TODO:
// setDefaultMessageCallback
// setMesssageCallback

////////////////////////////////////////////////////////////////

type ChannelHandler func(csound CSOUND, channelName string,
	channelValue []MYFLT, channelType int)

var inputChannelH ChannelHandler

//export goInputChannelCB
func goInputChannelCB(csound unsafe.Pointer, channelName *C.char,
	channelValuePtr unsafe.Pointer, channelType unsafe.Pointer) {
	if inputChannelH == nil {
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
	inputChannelH(cs, C.GoString(channelName), slice, chnType)
}

// Set the function which will be called whenever the invalue opcode is used.
func (csound CSOUND) SetInputChannelCallback(f ChannelHandler) {
	inputChannelH = f
	C.csoundSetInputChannelCB(csound.Cs)
}

var outputChannelH ChannelHandler

//export goOutputChannelCB
func goOutputChannelCB(csound unsafe.Pointer, channelName *C.char,
	channelValuePtr unsafe.Pointer, channelType unsafe.Pointer) {
	if outputChannelH == nil {
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
	outputChannelH(cs, C.GoString(channelName), slice, chnType)
}

// Set the function which will be called whenever the outvalue opcode is used.
func (csound CSOUND) SetOutputChannelCallback(f ChannelHandler) {
	outputChannelH = f
	C.csoundSetOutputChannelCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type SenseEventHandler func(csound CSOUND, userData unsafe.Pointer)

const maxNumSenseEvent = 10

var senseEventH []SenseEventHandler = make([]SenseEventHandler, maxNumSenseEvent)
var numSenseEvent int

// Workaround to avoid the 'cgo argument has Go pointer to Go pointer'
// runtime error (since go1.6)
var registry = make(map[int]unsafe.Pointer)
var indexes int
var mutex = sync.Mutex{}

//export goSenseEventCB
func goSenseEventCB(csound, userData unsafe.Pointer, numFun int32) {
	if senseEventH[numFun] == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	mutex.Lock()
	userDataP := registry[*(*int)(userData)]
	mutex.Unlock()
	senseEventH[numFun](cs, userDataP)
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
func (csound CSOUND) RegisterSenseEventCallback(f SenseEventHandler,
	userData unsafe.Pointer) (ret int, err error) {
	if numSenseEvent < maxNumSenseEvent {
		mutex.Lock()
		index := indexes
		registry[index] = userData
		mutex.Unlock()
		ret = int(C.csoundRegisterSenseEventCB(csound.Cs, unsafe.Pointer(&index), C.int(numSenseEvent)))
		if ret != 0 {
			err = fmt.Errorf("Csound could not register SenseEvent callback %d", numSenseEvent)
		} else {
			senseEventH[numSenseEvent] = f
			numSenseEvent++
			indexes++
		}
		return
	}
	ret = -1
	err = fmt.Errorf("%d SenseEvent callbacks already registered! Max value reached", numSenseEvent)
	return
}

////////////////////////////////////////////////////////////////

// TODO:
//   RegisterKeyboardCallback
//   RemoveKeyboardCallback
////////////////////////////////////////////////////////////////

type MakeGraphHandler func(csound CSOUND, windat unsafe.Pointer, name string)
type GraphHandler func(csound CSOUND, windat unsafe.Pointer)
type ExitGraphHandler func(csound CSOUND) int32

var makeGraphH MakeGraphHandler
var drawGraphH, killGraphH GraphHandler
var exitGraphH ExitGraphHandler

//export goMakeGraphCB
func goMakeGraphCB(csound, windat unsafe.Pointer, name *C.char) {
	if makeGraphH == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	makeGraphH(cs, windat, C.GoString(name))
}

//export goDrawGraphCB
func goDrawGraphCB(csound, windat unsafe.Pointer) {
	if drawGraphH == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	drawGraphH(cs, windat)
}

//export goKillGraphCB
func goKillGraphCB(csound, windat unsafe.Pointer) {
	if killGraphH == nil {
		return
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	killGraphH(cs, windat)
}

//export goExitGraphCB
func goExitGraphCB(csound unsafe.Pointer) int32 {
	if exitGraphH == nil {
		return -1
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	return exitGraphH(cs)
}

// Called by external software to set Csound's MakeGraph function.
func (csound CSOUND) SetMakeGraphCallback(f MakeGraphHandler) {
	makeGraphH = f
	C.csoundSetMakeGraphCB(csound.Cs)
}

// Called by external software to set Csound's DrawGraph function.
func (csound CSOUND) SetDrawGraphCallback(f GraphHandler) {
	drawGraphH = f
	C.csoundSetDrawGraphCB(csound.Cs)
}

// Called by external software to set Csound's KillGraph function.
func (csound CSOUND) SetKillGraphCallback(f GraphHandler) {
	killGraphH = f
	C.csoundSetKillGraphCB(csound.Cs)
}

// Called by external software to set Csound's ExitGraph function.
func (csound CSOUND) SetExitGraphCallback(f ExitGraphHandler) {
	exitGraphH = f
	C.csoundSetExitGraphCB(csound.Cs)
}

////////////////////////////////////////////////////////////////

type YieldHandler func(csound CSOUND) bool

var yieldH YieldHandler

//export goYieldCB
func goYieldCB(csound unsafe.Pointer) int32 {
	if yieldH == nil {
		return 0
	}
	cs := CSOUND{(*C.CSOUND)(csound)}
	if yieldH(cs) {
		return 1
	} else {
		return 0
	}
}

// Called by external software to set a function for checking system
// events, yielding cpu time for coopertative multitasking, etc.
// This function is optional. It is often used as a way to 'turn off'
// Csound, allowing it to exit gracefully. In addition, some operations
// like utility analysis routines are not reentrant and you should use
// this function to do any kind of updating during the operation.
// Returns an 'OK to continue' boolean.
func (csound CSOUND) SetYieldCallback(f YieldHandler) {
	yieldH = f
	C.csoundSetYieldCB(csound.Cs)
}
