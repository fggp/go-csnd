package csnd6

/*
#cgo CFLAGS: -DUSE_DOUBLE=1
#cgo CFLAGS: -I /usr/local/include
#cgo linux CFLAGS: -DLINUX=1
#cgo LDFLAGS: -lcsound64 -lcsnd6

#include <csound/csound.h>

CS_AUDIODEVICE *getAudioDevList(CSOUND *csound, int n, int isOutput)
{
  CS_AUDIODEVICE *devs = (CS_AUDIODEVICE *)malloc(n*sizeof(CS_AUDIODEVICE));
  csoundGetAudioDevList(csound, devs, isOutput);
  return devs;
}

void getAudioDev(CS_AUDIODEVICE *devs, int i, char **pname, char **pid, char **pmodule,
                 int *nchnls, int *flag)
{
  CS_AUDIODEVICE dev = devs[i];
  *pname = dev.device_name;
  *pid = dev.device_id;
  *pmodule = dev.rt_module;
  *nchnls = dev.max_nchnls;
  *flag = dev.isOutput;
}

CS_MIDIDEVICE *getMidiDevList(CSOUND *csound, int n, int isOutput)
{
  CS_MIDIDEVICE *devs = (CS_MIDIDEVICE *)malloc(n*sizeof(CS_MIDIDEVICE));
  csoundGetMIDIDevList(csound, devs, isOutput);
  return devs;
}

void getMidiDev(CS_MIDIDEVICE *devs, int i, char **pname, char** piname, char **pid,
                char **pmodule, int *flag)
{
  CS_MIDIDEVICE dev = devs[i];
  *pname = dev.device_name;
  *piname = dev.interface_name;
  *pid = dev.device_id;
  *pmodule = dev.midi_module;
  *flag = dev.isOutput;
}

void getOpcodeEntry(opcodeListEntry *list, int n,
                    char **opname, char **outypes, char** intypes)
{
  opcodeListEntry entry = list[n];
  *opname = entry.opname;
  *outypes = entry.outypes;
  *intypes = entry.intypes;
}

controlChannelHints_t *getControlChannelInfo(controlChannelInfo_t *list, int i,
                                             char **name, int *type)
{
  *name = list[i].name;
  *type = list[i].type;
  return &list[i].hints;
}

PVSDATEXT *newPvsData()
{
  return (PVSDATEXT *)calloc(sizeof(PVSDATEXT), 1);
}

void freePvsData(PVSDATEXT *p)
{
  if (p)
    free(p);
}

typedef struct namedgen {
  char *name;
  int  gennum;
  struct namedgen *next;
} NAMEDGEN;

int getNumNamedGens(CSOUND *csound)
{
  NAMEDGEN *p;
  int n;

  n = 0;
  p = (NAMEDGEN *)csoundGetNamedGens(csound);
  while (p) {
    n++;
    p = p->next;
  }
  return n;
}

void *getNamedGen(CSOUND *csound, void *currentGen, char **pname, int *num)
{
  NAMEDGEN *p = (NAMEDGEN *)currentGen;
  *pname = p->name;
  *num = p->gennum;
  return p->next;
}

int utilityListLength(char **list)
{
  int n;

  n = 0;
  while (*list++) {
    n++;
  }
  return n;
}

char *utilityName(char **list, int i)
{
  return *(list+i);
}

CsoundRandMTState *newRandMTState()
{
  return (CsoundRandMTState *)malloc(sizeof(CsoundRandMTState));
}

void freeRandMTState(CsoundRandMTState *p)
{
  if (p)
    free(p);
}
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

func cbool(flag bool) C.int {
	if flag {
		return 1
	}
	return 0
}

func cMYFLT(val MYFLT) C.double {
	return C.double(val)
}

func cpMYFLT(pval *MYFLT) *C.double {
	return (*C.double)(pval)
}

func cppMYFLT(ppval **MYFLT) **C.double {
	return (**C.double)(unsafe.Pointer(ppval))
}

func FileOpen(name, mode string) *C.FILE {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var cmode *C.char = C.CString(mode)
	defer C.free(unsafe.Pointer(cmode))
	return C.fopen(cname, cmode)
}

func FileClose(f *C.FILE) {
	C.fclose(f)
}

const (
	CSOUND_SUCCESS        = 0
	CSOUND_ERROR          = -1
	CSOUND_INITIALIZATION = -2
	CSOUND_PERFORMANCE    = -3
	CSOUND_MEMORY         = -4
	CSOUND_SIGNAL         = -5
)

const (
	CSOUNDINIT_NO_SIGNAL_HANDLER = 1
	CSOUNDINIT_NO_ATEXIT         = 2
)

const (
	// This should only be used internally by the original FileOpen()
	// API call or for temp files written with <CsFileB>
	CSFTYPE_UNKNOWN = iota

	CSFTYPE_UNIFIED_CSD // Unified Csound document
	CSFTYPE_ORCHESTRA   // the primary orc file (may be temporary)
	CSFTYPE_SCORE       // the primary sco file (may be temporary)
	// or any additional score opened by Cscore
	CSFTYPE_ORC_INCLUDE   // a file #included by the orchestra
	CSFTYPE_SCO_INCLUDE   // a file #included by the score
	CSFTYPE_SCORE_OUT     // used for score.srt score.xtr cscore.out
	CSFTYPE_SCOT          // Scot score input format
	CSFTYPE_OPTIONS       // for .csoundrc and -@ flag
	CSFTYPE_EXTRACT_PARMS // extraction file specified by -x

	// audio file types that Csound can write (10-19) or read
	CSFTYPE_RAW_AUDIO
	CSFTYPE_IRCAM
	CSFTYPE_AIFF
	CSFTYPE_AIFC
	CSFTYPE_WAVE
	CSFTYPE_AU
	CSFTYPE_SD2
	CSFTYPE_W64
	CSFTYPE_WAVEX
	CSFTYPE_FLAC
	CSFTYPE_CAF
	CSFTYPE_WVE
	CSFTYPE_OGG
	CSFTYPE_MPC2K
	CSFTYPE_RF64
	CSFTYPE_AVR
	CSFTYPE_HTK
	CSFTYPE_MAT4
	CSFTYPE_MAT5
	CSFTYPE_NIST
	CSFTYPE_PAF
	CSFTYPE_PVF
	CSFTYPE_SDS
	CSFTYPE_SVX
	CSFTYPE_VOC
	CSFTYPE_XI
	CSFTYPE_UNKNOWN_AUDIO // used when opening audio file for reading
	// or temp file written with <CsSampleB>

	// miscellaneous music formats
	CSFTYPE_SOUNDFONT
	CSFTYPE_STD_MIDI   // Standard MIDI file
	CSFTYPE_MIDI_SYSEX // Raw MIDI codes eg. SysEx dump

	// analysis formats
	CSFTYPE_HETRO
	CSFTYPE_HETROT
	CSFTYPE_PVC   // original PVOC format
	CSFTYPE_PVCEX // PVOC-EX format
	CSFTYPE_CVANAL
	CSFTYPE_LPC
	CSFTYPE_ATS
	CSFTYPE_LORIS
	CSFTYPE_SDIF
	CSFTYPE_HRTF

	// Types for plugins and the files they read/write
	CSFTYPE_VST_PLUGIN
	CSFTYPE_LADSPA_PLUGIN
	CSFTYPE_SNAPSHOT

	// Special formats for Csound ftables or scanned synthesis
	// matrices with header info
	CSFTYPE_FTABLES_TEXT   // for ftsave and ftload
	CSFTYPE_FTABLES_BINARY // for ftsave and ftload
	CSFTYPE_XSCANU_MATRIX  // for xscanu opcode

	// These are for raw lists of numbers without header info
	CSFTYPE_FLOATS_TEXT    // used by GEN23 GEN28 dumpk readk
	CSFTYPE_FLOATS_BINARY  // used by dumpk readk etc.
	CSFTYPE_INTEGER_TEXT   // used by dumpk readk etc.
	CSFTYPE_INTEGER_BINARY // used by dumpk readk etc.

	// image file formats
	CSFTYPE_IMAGE_PNG

	// For files that don't match any of the above
	CSFTYPE_POSTSCRIPT  // EPS format used by graphs
	CSFTYPE_SCRIPT_TEXT // executable script files (eg. Python)
	CSFTYPE_OTHER_TEXT
	CSFTYPE_OTHER_BINARY
)

type CSOUND struct {
	cs (*C.CSOUND)
}

type MYFLT float64

type CsoundParams struct {
	DebugMode            int32 // debug mode, 0 or 1
	BufferFrames         int32 // number of frames in in/out buffers
	HardwareBufferFrames int32 // ibid. hardware
	Displays             int32 // graph displays, 0 or 1
	AsciiGraphs          int32 // use ASCII graphs, 0 or 1
	PostscriptGraphs     int32 // use postscript graphs, 0 or 1
	MessageLevel         int32 // message printout control
	Tempo                int32 // tempo (sets Beatmode)
	RingBell             int32 // bell, 0 or 1
	UseCscore            int32 // use cscore for processing
	TerminateOnMidi      int32 // terminate performance at the end of midifile, 0 or 1
	HeartBeat            int32 // print heart beat, 0 or 1
	DeferGen01Load       int32 // defer GEN01 load, 0 or 1
	MidiKey              int32 // pfield to map midi key no
	MidiKeyCps           int32 // pfield to map midi key no as cps
	MidiKeyOct           int32 // pfield to map midi key no as oct
	MidiKeyPch           int32 // pfield to map midi key no as pch
	MidiVelocity         int32 // pfield to map midi velocity
	MidiVelocityAmp      int32 // pfield to map midi velocity as amplitude
	NoDefaultPaths       int32 // disable relative paths from files, 0 or 1
	NumberOfThreads      int32 // number of threads for multicore performance
	SyntaxCheckOnly      int32 // do not compile, only check syntax
	CsdLineCounts        int32 // csd line error reporting
	ComputeWeights       int32 // use calculated opcode weights for multicore, 0 or 1
	RealtimeMode         int32 // use realtime priority mode, 0 or 1
	SampleAccurate       int32 // use sample-level score event accuracy
	SampleRateOverride   MYFLT // overriding sample rate
	ControlRateOverride  MYFLT // overriding control rate
	NchnlsOverride       int32 // overriding number of out channels
	NchnlsIoverride      int32 // overriding number of in channels
	E0dbfsOverride       MYFLT // overriding 0dbfs
}

type CsoundAudioDevice struct {
	DeviceName string
	DeviceId   string
	RtModule   string
	MaxNchnls  int
	IsOutput   bool
}

func (dev CsoundAudioDevice) String() string {
	return fmt.Sprintf("(%s, %s, %s, %d, %t)", dev.DeviceName, dev.DeviceId,
		dev.RtModule, dev.MaxNchnls, dev.IsOutput)
}

type CsoundMidiDevice struct {
	DeviceName    string
	InterfaceName string
	DeviceId      string
	MidiModule    string
	IsOutput      bool
}

func (dev CsoundMidiDevice) String() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %t)", dev.DeviceName,
		dev.InterfaceName, dev.DeviceId, dev.MidiModule, dev.IsOutput)
}

type OpcodeListEntry struct {
	Opname  string
	Outypes string
	Intypes string
}

type TREE struct {
	t (*C.TREE)
}

const (
	CSOUND_CONTROL_CHANNEL   = 1
	CSOUND_AUDIO_CHANNEL     = 2
	CSOUND_STRING_CHANNEL    = 3
	CSOUND_PVS_CHANNEL       = 4
	CSOUND_VAR_CHANNEL       = 5
	CSOUND_CHANNEL_TYPE_MASK = 15
	CSOUND_INPUT_CHANNEL     = 16
	CSOUND_OUTPUT_CHANNEL    = 32
)

const (
	CSOUND_CONTROL_CHANNEL_NO_HINTS = 0
	CSOUND_CONTROL_CHANNEL_INT      = 1
	CSOUND_CONTROL_CHANNEL_LIN      = 2
	CSOUND_CONTROL_CHANNEL_EXP      = 3
)

type ControlChannelHints struct {
	Behav      int
	Dflt       MYFLT
	Min        MYFLT
	Max        MYFLT
	X          int
	Y          int
	Width      int
	Height     int
	Attributes string // This member must be set explicitly to NULL if not used
}

type ControlChannelInfo struct {
	Name  string
	Type  int
	Hints ControlChannelHints
}

/*
 * Instantiation
 */
// Initialize Csound library with specific flags. This function is called
// internally by csoundCreate(), so there is generally no need to use it
// explicitly unless you need to avoid default initilization that sets
// signal handlers and atexit() callbacks.
// Return value is zero on success, positive if initialisation was
// done already, and negative on error.
func Initialize(flags int) int {
	return int(C.csoundInitialize(C.int(flags)))
}

func Create(hostData unsafe.Pointer) CSOUND {
	var cs (*C.CSOUND)
	if hostData != nil {
		cs = C.csoundCreate(hostData)
	} else {
		cs = C.csoundCreate(nil)
	}
	return CSOUND{cs}
}

func (csound *CSOUND) Destroy() {
	C.csoundDestroy(csound.cs)
	csound.cs = nil
}

func (csound CSOUND) GetVersion() int {
	return int(C.csoundGetVersion())
}

func (csound CSOUND) GetAPIVersion() int {
	return int(C.csoundGetAPIVersion())
}

/*
 * Performance
 */
func (csound CSOUND) ParseOrc(str string) TREE {
	var cstr *C.char = C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	t := C.csoundParseOrc(csound.cs, cstr)
	return TREE{t}
}

func (csound CSOUND) CompileTree(root TREE) int {
	result := C.csoundCompileTree(csound.cs, root.t)
	return int(result)
}

func (csound CSOUND) DeleteTree(tree TREE) {
	C.csoundDeleteTree(csound.cs, tree.t)
}

func (csound CSOUND) CompileOrc(str string) int {
	var cstr *C.char = C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.csoundCompileOrc(csound.cs, cstr))
}

func (csound CSOUND) EvalCode(str string) MYFLT {
	var cstr *C.char = C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return MYFLT(C.csoundEvalCode(csound.cs, cstr))
}

func (csound CSOUND) InitializeCscore(insco, outsco *C.FILE) int {
	return int(C.csoundInitializeCscore(csound.cs, insco, outsco))
}

func (csound CSOUND) CompileArgs(args []string) int {
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	result := C.csoundCompileArgs(csound.cs, argc, &argv[0])
	for _, arg := range argv {
		C.free(unsafe.Pointer(arg))
	}
	return int(result)
}

func (csound CSOUND) Start() int {
	return int(C.csoundStart(csound.cs))
}

func (csound CSOUND) Compile(args []string) int {
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	result := C.csoundCompile(csound.cs, argc, &argv[0])
	for _, arg := range argv {
		C.free(unsafe.Pointer(arg))
	}
	return int(result)
}

func (csound CSOUND) Perform() int {
	return int(C.csoundPerform(csound.cs))
}

func (csound CSOUND) PerformKsmps() int {
	return int(C.csoundPerformKsmps(csound.cs))
}

func (csound CSOUND) PerformBuffer() int {
	return int(C.csoundPerformBuffer(csound.cs))
}

func (csound CSOUND) Stop() {
	C.csoundStop(csound.cs)
}

func (csound CSOUND) Cleanup() int {
	numSenseEvent = 0
	return int(C.csoundCleanup(csound.cs))
}

func (csound CSOUND) Reset() {
	C.csoundReset(csound.cs)
}

/*
 * Attributes
 */
func (csound CSOUND) GetSr() MYFLT {
	return MYFLT(C.csoundGetSr(csound.cs))
}

func (csound CSOUND) GetKr() MYFLT {
	return MYFLT(C.csoundGetKr(csound.cs))
}

func (csound CSOUND) GetKsmps() int {
	return int(C.csoundGetKsmps(csound.cs))
}

func (csound CSOUND) GetNchnls() int {
	return int(C.csoundGetNchnls(csound.cs))
}

func (csound CSOUND) GetNchnlsInput() int {
	return int(C.csoundGetNchnlsInput(csound.cs))
}

func (csound CSOUND) Get0dBFS() MYFLT {
	return MYFLT(C.csoundGet0dBFS(csound.cs))
}

func (csound CSOUND) GetCurrentTimeSamples() int {
	return int(C.csoundGetCurrentTimeSamples(csound.cs))
}

func (csound CSOUND) GetSizeOfMYFLT() int {
	return int(C.csoundGetSizeOfMYFLT())
}

func (csound CSOUND) GetHostData() unsafe.Pointer {
	return C.csoundGetHostData(csound.cs)
}

func (csound CSOUND) SetHostData(hostData unsafe.Pointer) {
	C.csoundSetHostData(csound.cs, hostData)
}

func (csound CSOUND) SetOption(option string) int {
	var coption *C.char = C.CString(option)
	defer C.free(unsafe.Pointer(coption))
	return int(C.csoundSetOption(csound.cs, coption))
}

func (csound CSOUND) SetParams(p *CsoundParams) {
	pp := &p.DebugMode
	C.csoundSetParams(csound.cs, (*C.CSOUND_PARAMS)(unsafe.Pointer(pp)))
}

func (csound CSOUND) GetParams(p *CsoundParams) {
	pp := &p.DebugMode
	C.csoundGetParams(csound.cs, (*C.CSOUND_PARAMS)(unsafe.Pointer(pp)))
}

func (csound CSOUND) GetDebug() bool {
	return C.csoundGetDebug(csound.cs) != 0
}

func (csound CSOUND) SetDebug(debug bool) {
	C.csoundSetDebug(csound.cs, cbool(debug))
}

/*
 * General Input/Output
 */
func (csound CSOUND) GetOutputName() string {
	return C.GoString(C.csoundGetOutputName(csound.cs))
}

func (csound CSOUND) SetOutput(name, otype, format string) {
	var cname, ctype, cformat *C.char
	cname = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	if len(otype) > 0 {
		ctype = C.CString(otype)
		defer C.free(unsafe.Pointer(ctype))
	} else {
		ctype = nil
	}
	if len(format) > 0 {
		cformat = C.CString(format)
		defer C.free(unsafe.Pointer(cformat))
	} else {
		cformat = nil
	}
	C.csoundSetOutput(csound.cs, cname, ctype, cformat)
}

func (csound CSOUND) SetInput(name string) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetInput(csound.cs, cname)
}

func (csound CSOUND) SetMIDIInput(name string) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetMIDIInput(csound.cs, cname)
}

func (csound CSOUND) SetMIDIFileInput(name string) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetMIDIFileInput(csound.cs, cname)
}

func (csound CSOUND) SetMIDIOutput(name string) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetMIDIOutput(csound.cs, cname)
}

func (csound CSOUND) SetMIDIFileOutput(name string) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetMIDIFileOutput(csound.cs, cname)
}

/*
 * Realtime Audio I/O
 */
func (csound CSOUND) SetRTAudioModule(module string) {
	var cmodule *C.char = C.CString(module)
	defer C.free(unsafe.Pointer(cmodule))
	C.csoundSetRTAudioModule(csound.cs, cmodule)
}

func (csound CSOUND) GetModule(number int) (name, mtype string, error int) {
	var cname, ctype *C.char
	cerror := C.csoundGetModule(csound.cs, C.int(number), &cname, &ctype)
	name = C.GoString(cname)
	mtype = C.GoString(ctype)
	error = int(cerror)
	return
}

func (csound CSOUND) GetInputBufferSize() int {
	return int(C.csoundGetInputBufferSize(csound.cs))
}

func (csound CSOUND) GetOutputBufferSize() int {
	return int(C.csoundGetOutputBufferSize(csound.cs))
}

func (csound CSOUND) GetInputBuffer() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetInputBuffer(csound.cs))
	length := int(C.csoundGetInputBufferSize(csound.cs))
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(buffer))
	return slice
}

func (csound CSOUND) GetOutputBuffer() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetOutputBuffer(csound.cs))
	length := int(C.csoundGetOutputBufferSize(csound.cs))
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(buffer))
	return slice
}

func (csound CSOUND) GetSpin() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetSpin(csound.cs))
	length := csound.GetKsmps() * csound.GetNchnls()
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(buffer))
	return slice
}

func (csound CSOUND) AddSpinSample(frame, channel int, sample MYFLT) {
	C.csoundAddSpinSample(csound.cs, C.int(frame), C.int(channel), cMYFLT(sample))
}

func (csound CSOUND) GetSpout() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetSpout(csound.cs))
	length := csound.GetKsmps() * csound.GetNchnls()
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(buffer))
	return slice
}

func (csound CSOUND) GetSpoutSample(frame, channel int) MYFLT {
	return MYFLT(C.csoundGetSpoutSample(csound.cs, C.int(frame), C.int(channel)))
}

func (csound CSOUND) GetRtRecordUserData() unsafe.Pointer {
	return unsafe.Pointer(C.csoundGetRtRecordUserData(csound.cs))
}

func (csound CSOUND) GetRtPlaydUserData() unsafe.Pointer {
	return unsafe.Pointer(C.csoundGetRtPlayUserData(csound.cs))
}

func (csound CSOUND) SetHostImplementedAudioIO(state, bufSize int) {
	C.csoundSetHostImplementedAudioIO(csound.cs, C.int(state), C.int(bufSize))
}

func (csound CSOUND) GetAudioDevList(isOutput bool) []CsoundAudioDevice {
	cflag := cbool(isOutput)
	n := C.csoundGetAudioDevList(csound.cs, nil, cflag)
	devs := C.getAudioDevList(csound.cs, n, cflag)
	defer C.free(unsafe.Pointer(devs))
	length := int(n)
	var list = make([]CsoundAudioDevice, length)
	var name, id, module *C.char
	var nchnls C.int
	for i := range list {
		C.getAudioDev(devs, C.int(i), &name, &id, &module, &nchnls, &cflag)
		list[i].DeviceName = C.GoString(name)
		list[i].DeviceId = C.GoString(id)
		list[i].RtModule = C.GoString(module)
		list[i].MaxNchnls = int(nchnls)
		list[i].IsOutput = (cflag == 1)
	}
	return list
}

/*
 * Realtime Midi I/O
 */
func (csound CSOUND) SetMIDIModule(module string) {
	var cmodule *C.char = C.CString(module)
	defer C.free(unsafe.Pointer(cmodule))
	C.csoundSetMIDIModule(csound.cs, cmodule)
}

func (csound CSOUND) SetHostImplementedMIDIIO(state bool) {
	C.csoundSetHostImplementedMIDIIO(csound.cs, cbool(state))
}

func (csound CSOUND) GetMidiDevList(isOutput bool) []CsoundMidiDevice {
	cflag := cbool(isOutput)
	n := C.csoundGetMIDIDevList(csound.cs, nil, cflag)
	devs := C.getMidiDevList(csound.cs, n, cflag)
	defer C.free(unsafe.Pointer(devs))
	length := int(n)
	var list = make([]CsoundMidiDevice, length)
	var name, iname, id, module *C.char
	for i := range list {
		C.getMidiDev(devs, C.int(i), &name, &iname, &id, &module, &cflag)
		list[i].DeviceName = C.GoString(name)
		list[i].InterfaceName = C.GoString(iname)
		list[i].DeviceId = C.GoString(id)
		list[i].MidiModule = C.GoString(module)
		list[i].IsOutput = (cflag == 1)
	}
	return list
}

/*
 * Score handling
 */
func (csound CSOUND) ReadScore(str string) int {
	var cstr *C.char = C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.csoundReadScore(csound.cs, cstr))
}

func (csound CSOUND) GetScoreTime() float64 {
	return float64(C.csoundGetScoreTime(csound.cs))
}

func (csound CSOUND) IsScorePending() bool {
	return C.csoundIsScorePending(csound.cs) != 0
}

func (csound CSOUND) SetScorePending(pending bool) {
	C.csoundSetScorePending(csound.cs, cbool(pending))
}

func (csound CSOUND) GetScoreOffsetSeconds() MYFLT {
	return MYFLT(C.csoundGetScoreOffsetSeconds(csound.cs))
}

func (csound CSOUND) SetScoreOffsetSeconds(time MYFLT) {
	C.csoundSetScoreOffsetSeconds(csound.cs, cMYFLT(time))
}

func (csound CSOUND) RewindScore() {
	C.csoundRewindScore(csound.cs)
}

func (csound CSOUND) ScoreSort(inFile, outFile *C.FILE) {
	C.csoundScoreSort(csound.cs, inFile, outFile)
}

func (csound CSOUND) ScoreExtract(inFile, outFile, extractFile *C.FILE) {
	C.csoundScoreExtract(csound.cs, inFile, outFile, extractFile)
}

/*
 * Messages and Text
 */
func (csound CSOUND) GetMessageLevel() int {
	return int(C.csoundGetMessageLevel(csound.cs))
}

func (csound CSOUND) SetMessageLevel(messageLevel int) {
	C.csoundSetMessageLevel(csound.cs, C.int(messageLevel))
}

func (csound CSOUND) CreateMessageBuffer(toStdOut bool) {
	C.csoundCreateMessageBuffer(csound.cs, cbool(toStdOut))
}

func (csound CSOUND) GetFirstMessage() string {
	cmsg := C.csoundGetFirstMessage(csound.cs)
	return C.GoString(cmsg)
}

func (csound CSOUND) GetFirstMessageAttr() int {
	return int(C.csoundGetFirstMessageAttr(csound.cs))
}

func (csound CSOUND) PopFirstMessage() {
	C.csoundPopFirstMessage(csound.cs)
}

func (csound CSOUND) GetMessageCnt() int {
	return int(C.csoundGetMessageCnt(csound.cs))
}

func (csound CSOUND) DestroyMessageBuffer() {
	C.csoundDestroyMessageBuffer(csound.cs)
}

/*
 * Channels, Control and Events
 */
func (csound CSOUND) GetChannelPtr(name string, chnType int) ([]MYFLT, error) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var ptr *MYFLT
	var length int
	switch chnType & CSOUND_CHANNEL_TYPE_MASK {
	case CSOUND_CONTROL_CHANNEL:
		length = 1
	case CSOUND_AUDIO_CHANNEL:
		length = int(C.csoundGetKsmps(csound.cs))
	case CSOUND_STRING_CHANNEL:
		length = int(C.csoundGetChannelDatasize(csound.cs, cname))
	default:
		return nil, fmt.Errorf("%d is not a valid channel type", chnType)
	}
	ret := C.csoundGetChannelPtr(csound.cs, cppMYFLT(&ptr), cname, C.int(chnType))
	switch ret {
	case CSOUND_SUCCESS:
		var slice []MYFLT
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
		sliceHeader.Cap = length
		sliceHeader.Len = length
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		return slice, nil
	case CSOUND_MEMORY:
		return nil, fmt.Errorf("Not enough memory for allocating the channel")
	case CSOUND_ERROR:
		return nil, fmt.Errorf("The specified name or type is invalid")
	default:
		return nil, fmt.Errorf("Unknown error")
	}
}

func (csound CSOUND) ListChannels() ([]ControlChannelInfo, error) {
	var lst *C.controlChannelInfo_t
	n := int(C.csoundListChannels(csound.cs, &lst))
	if n == CSOUND_MEMORY {
		return nil, fmt.Errorf("Not enough memory for allocating channels list")
	} else if n == 0 {
		return nil, nil
	} else {
		var list = make([]ControlChannelInfo, n)
		var name *C.char
		var ctype C.int
		for i := range list {
			hints := C.getControlChannelInfo(lst, C.int(i), &name, &ctype)
			list[i].Name = C.GoString(name)
			list[i].Type = int(ctype)
			list[i].Hints.Behav = int(hints.behav)
			list[i].Hints.Dflt = MYFLT(hints.dflt)
			list[i].Hints.Min = MYFLT(hints.min)
			list[i].Hints.Max = MYFLT(hints.max)
			list[i].Hints.X = int(hints.x)
			list[i].Hints.Y = int(hints.y)
			list[i].Hints.Width = int(hints.width)
			list[i].Hints.Height = int(hints.height)
			list[i].Hints.Attributes = C.GoString(hints.attributes)
		}
		C.csoundDeleteChannelList(csound.cs, lst)
		return list, nil
	}
}

func (csound CSOUND) SetControlChannelHints(name string, hints ControlChannelHints) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var chints C.controlChannelHints_t
	chints.behav = C.controlChannelBehavior(hints.Behav)
	chints.dflt = cMYFLT(hints.Dflt)
	chints.min = cMYFLT(hints.Min)
	chints.max = cMYFLT(hints.Max)
	chints.x = C.int(hints.X)
	chints.y = C.int(hints.Y)
	chints.width = C.int(hints.Width)
	chints.height = C.int(hints.Height)
	if len(hints.Attributes) > 0 {
		chints.attributes = C.CString(name)
		defer C.free(unsafe.Pointer(chints.attributes))
	}
	return int(C.csoundSetControlChannelHints(csound.cs, cname, chints))
}

func (csound CSOUND) GetControlChannelHints(name string) (ControlChannelHints, int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var chints C.controlChannelHints_t
	ret := C.csoundGetControlChannelHints(csound.cs, cname, &chints)
	var hints ControlChannelHints
	if ret == 0 {
		hints.Behav = int(chints.behav)
		hints.Dflt = MYFLT(chints.dflt)
		hints.Min = MYFLT(chints.min)
		hints.Max = MYFLT(chints.max)
		hints.X = int(chints.x)
		hints.Y = int(chints.y)
		hints.Width = int(chints.width)
		hints.Height = int(chints.height)
		hints.Attributes = C.GoString(chints.attributes)
		if chints.attributes != nil {
			defer C.free(unsafe.Pointer(chints.attributes))
		}
	}
	return hints, int(ret)
}

func (csound CSOUND) GetChannelLock(name string) *C.int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.csoundGetChannelLock(csound.cs, cname)
}

func (csound CSOUND) GetControlChannel(name string) (MYFLT, int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var err C.int
	val := MYFLT(C.csoundGetControlChannel(csound.cs, cname, &err))
	return val, int(err)
}

func (csound CSOUND) SetControlChannel(name string, val MYFLT) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetControlChannel(csound.cs, cname, cMYFLT(val))
}

func (csound CSOUND) GetAudioChannel(name string, samples []MYFLT) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundGetAudioChannel(csound.cs, cname, cpMYFLT(&samples[0]))
}

func (csound CSOUND) SetAudioChannel(name string, samples []MYFLT) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetAudioChannel(csound.cs, cname, cpMYFLT(&samples[0]))
}

func (csound CSOUND) GetStringChannel(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	size := C.csoundGetChannelDatasize(csound.cs, cname)
	cstr := (*C.char)(C.malloc(C.size_t(size)))
	defer C.free(unsafe.Pointer(cstr))
	C.csoundGetStringChannel(csound.cs, cname, cstr)
	return C.GoString(cstr)
}

func (csound CSOUND) SetStringChannel(name, str string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.csoundSetStringChannel(csound.cs, cname, cstr)
}

func (csound CSOUND) GetChannelDatasize(name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundGetChannelDatasize(csound.cs, cname))
}

type PVSDATEXT struct {
	Frame   []float32
	CStruct *C.PVSDATEXT
}

func NewPvsData(N, format, overlap, winsize int32) *PVSDATEXT {
	p := PVSDATEXT{Frame: make([]float32, N+2)}
	cp := C.newPvsData()
	cp.N = C.int32(N)
	cp.format = C.int32(format)
	cp.overlap = C.int32(overlap)
	cp.winsize = C.int32(winsize)
	cp.frame = (*C.float)(&p.Frame[0])
	p.CStruct = cp
	return &p
}

func FreeCPvsData(p *PVSDATEXT) {
	C.freePvsData(p.CStruct)
}

func (csound CSOUND) SetPvsChannel(fin *PVSDATEXT, name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundSetPvsChannel(csound.cs, fin.CStruct, cname))
}

func (csound CSOUND) GetPvsChannel(fout *PVSDATEXT, name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundGetPvsChannel(csound.cs, fout.CStruct, cname))
}

func (csound CSOUND) ScoreEvent(eventType byte, pFields []MYFLT) int {
	return int(C.csoundScoreEvent(csound.cs, C.char(eventType),
		cpMYFLT(&pFields[0]), C.long(len(pFields))))
}

func (csound CSOUND) ScoreEventAbsolute(eventType byte, pFields []MYFLT,
	timeOfs float64) int {
	return int(C.csoundScoreEventAbsolute(csound.cs, C.char(eventType),
		cpMYFLT(&pFields[0]), C.long(len(pFields)),
		C.double(timeOfs)))
}

func (csound CSOUND) InputMessage(message string) {
	var cmsg *C.char = C.CString(message)
	defer C.free(unsafe.Pointer(cmsg))
	C.csoundInputMessage(csound.cs, cmsg)
}

func (csound CSOUND) KillInstance(instr MYFLT, instrName string, mode int,
	allow_release bool) int {
	var cname *C.char
	if len(instrName) > 0 {
		cname = C.CString(instrName)
		defer C.free(unsafe.Pointer(cname))
	} else {
		cname = nil
	}
	return int(C.csoundKillInstance(csound.cs, cMYFLT(instr), cname, C.int(mode),
		cbool(allow_release)))
}

func (csound CSOUND) KeyPress(c byte) {
	C.csoundKeyPress(csound.cs, C.char(c))
}

/*
 * Tables
 */
func (csound CSOUND) TableLength(table int) int {
	return int(C.csoundTableLength(csound.cs, C.int(table)))
}

func (csound CSOUND) TableGet(table, index int) MYFLT {
	return MYFLT(C.csoundTableGet(csound.cs, C.int(table), C.int(index)))
}

func (csound CSOUND) TableSet(table, index int, value MYFLT) {
	C.csoundTableSet(csound.cs, C.int(table), C.int(index), cMYFLT(value))
}

func (csound CSOUND) TableCopyOut(table int, dest []MYFLT) {
	cdest := cpMYFLT(&dest[0])
	C.csoundTableCopyOut(csound.cs, C.int(table), cdest)
}

func (csound CSOUND) TableCopyIn(table int, src []MYFLT) {
	csrc := cpMYFLT(&src[0])
	C.csoundTableCopyIn(csound.cs, C.int(table), csrc)
}

func (csound CSOUND) GetTable(tableNum int) ([]MYFLT, error) {
	var tablePtr *MYFLT
	length := int(C.csoundGetTable(csound.cs, cppMYFLT(&tablePtr), C.int(tableNum)))
	if length == -1 {
		return nil, fmt.Errorf("Function table %d does not exist", tableNum)
	}
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(tablePtr))
	return slice, nil
}

/*
 * Function table display
 */
func (csound CSOUND) SetIsGraphable(isGraphable int) int {
	return int(C.csoundSetIsGraphable(csound.cs, C.int(isGraphable)))
}

/*
 * Opcodes
 */
type NamedGen struct {
	Name string
	Num  int
}

func (csound CSOUND) GetNamedGens() []NamedGen {
	n := int(C.getNumNamedGens(csound.cs))
	if n == 0 {
		return nil
	}
	namedGens := make([]NamedGen, n)
	p := C.csoundGetNamedGens(csound.cs)
	var name *C.char
	var num C.int
	for i := range namedGens {
		p = C.getNamedGen(csound.cs, p, &name, &num)
		namedGens[i].Name = C.GoString(name)
		namedGens[i].Num = int(num)
	}
	return namedGens
}

func (csound CSOUND) OpcodeList() []OpcodeListEntry {
	var opcodeList *C.struct_opcodeListEntry
	length := int(C.csoundNewOpcodeList(csound.cs,
		(**_Ctype_opcodeListEntry)(unsafe.Pointer(&opcodeList))))
	var list = make([]OpcodeListEntry, length)
	var opname, outypes, intypes *C.char
	for i := range list {
		C.getOpcodeEntry((*_Ctype_opcodeListEntry)(unsafe.Pointer(opcodeList)),
			C.int(i), &opname, &outypes, &intypes)
		list[i].Opname = C.GoString(opname)
		list[i].Outypes = C.GoString(outypes)
		list[i].Intypes = C.GoString(intypes)
	}
	C.csoundDisposeOpcodeList(csound.cs,
		(*_Ctype_opcodeListEntry)(unsafe.Pointer(opcodeList)))
	return list
}

/*
 * Threading and concurrency
 */
func (csound *CSOUND) GetCurrentThreadId() unsafe.Pointer {
	return unsafe.Pointer(C.csoundGetCurrentThreadId())
}

func (csound *CSOUND) JoinThread(thread unsafe.Pointer) uintptr {
	return uintptr(C.csoundJoinThread(thread))
}

func (csound *CSOUND) CreateThreadLock() unsafe.Pointer {
	return unsafe.Pointer(C.csoundCreateThreadLock())
}

func (csound *CSOUND) WaitThreadLock(lock unsafe.Pointer, ms uint) int {
	return int(C.csoundWaitThreadLock(lock, C.size_t(ms)))
}

func (csound *CSOUND) WaitThreadLockNoTimeout(lock unsafe.Pointer) {
	C.csoundWaitThreadLockNoTimeout(lock)
}

func (csound *CSOUND) NotifyThreadLock(lock unsafe.Pointer) {
	C.csoundNotifyThreadLock(lock)
}

func (csound *CSOUND) DestroyThreadLock(lock unsafe.Pointer) {
	C.csoundDestroyThreadLock(lock)
}

func (csound *CSOUND) CreateMutex(isRecursive bool) unsafe.Pointer {
	return C.csoundCreateMutex(cbool(isRecursive))
}

func (csound *CSOUND) LockMutex(mutex unsafe.Pointer) {
	C.csoundLockMutex(mutex)
}

func (csound *CSOUND) LockMutexNoWait(mutex unsafe.Pointer) int {
	return int(C.csoundLockMutexNoWait(mutex))
}

func (csound *CSOUND) UnlockMutex(mutex unsafe.Pointer) {
	C.csoundUnlockMutex(mutex)
}

func (csound *CSOUND) DestroyMutex(mutex unsafe.Pointer) {
	C.csoundDestroyMutex(mutex)
}

func (csound *CSOUND) CreateBarrier(max uint) unsafe.Pointer {
	return C.csoundCreateBarrier(C.uint(max))
}

func (csound *CSOUND) DestroyBarrier(barrier unsafe.Pointer) int {
	return int(C.csoundDestroyBarrier(barrier))
}

func (csound *CSOUND) WaitBarrier(barrier unsafe.Pointer) int {
	return int(C.csoundWaitBarrier(barrier))
}

func (csound *CSOUND) Sleep(ms uint) {
	C.csoundSleep(C.size_t(ms))
}

/*
 * Miscellaneous functions
 */
func (csound *CSOUND) RunCommand(args []string, noWait bool) int {
	argv := make([]*C.char, len(args)+1)
	for i, arg := range args {
		argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(argv[i]))
	}
	return int(C.csoundRunCommand(&argv[0], cbool(noWait)))
}

func (csound *CSOUND) InitTimerStruct() C.RTCLOCK {
	var rtc C.RTCLOCK
	C.csoundInitTimerStruct(&rtc)
	return rtc
}

func (csound *CSOUND) GetRealTime(rtc *C.RTCLOCK) float64 {
	return float64(C.csoundGetRealTime(rtc))
}

func (csound *CSOUND) GetCPUTime(rtc *C.RTCLOCK) float64 {
	return float64(C.csoundGetCPUTime(rtc))
}

func (csound *CSOUND) GetRandomSeedFromTime() uint32 {
	return uint32(C.csoundGetRandomSeedFromTime())
}

type Cslanguage_t int

const (
	CSLANGUAGE_DEFAULT Cslanguage_t = iota
	CSLANGUAGE_AFRIKAANS
	CSLANGUAGE_ALBANIAN
	CSLANGUAGE_ARABIC
	CSLANGUAGE_ARMENIAN
	CSLANGUAGE_ASSAMESE
	CSLANGUAGE_AZERI
	CSLANGUAGE_BASQUE
	CSLANGUAGE_BELARUSIAN
	CSLANGUAGE_BENGALI
	CSLANGUAGE_BULGARIAN
	CSLANGUAGE_CATALAN
	CSLANGUAGE_CHINESE
	CSLANGUAGE_CROATIAN
	CSLANGUAGE_CZECH
	CSLANGUAGE_DANISH
	CSLANGUAGE_DUTCH
	CSLANGUAGE_ENGLISH_UK
	CSLANGUAGE_ENGLISH_US
	CSLANGUAGE_ESTONIAN
	CSLANGUAGE_FAEROESE
	CSLANGUAGE_FARSI
	CSLANGUAGE_FINNISH
	CSLANGUAGE_FRENCH
	CSLANGUAGE_GEORGIAN
	CSLANGUAGE_GERMAN
	CSLANGUAGE_GREEK
	CSLANGUAGE_GUJARATI
	CSLANGUAGE_HEBREW
	CSLANGUAGE_HINDI
	CSLANGUAGE_HUNGARIAN
	CSLANGUAGE_ICELANDIC
	CSLANGUAGE_INDONESIAN
	CSLANGUAGE_ITALIAN
	CSLANGUAGE_JAPANESE
	CSLANGUAGE_KANNADA
	CSLANGUAGE_KASHMIRI
	CSLANGUAGE_KAZAK
	CSLANGUAGE_KONKANI
	CSLANGUAGE_KOREAN
	CSLANGUAGE_LATVIAN
	CSLANGUAGE_LITHUANIAN
	CSLANGUAGE_MACEDONIAN
	CSLANGUAGE_MALAY
	CSLANGUAGE_MALAYALAM
	CSLANGUAGE_MANIPURI
	CSLANGUAGE_MARATHI
	CSLANGUAGE_NEPALI
	CSLANGUAGE_NORWEGIAN
	CSLANGUAGE_ORIYA
	CSLANGUAGE_POLISH
	CSLANGUAGE_PORTUGUESE
	CSLANGUAGE_PUNJABI
	CSLANGUAGE_ROMANIAN
	CSLANGUAGE_RUSSIAN
	CSLANGUAGE_SANSKRIT
	CSLANGUAGE_SERBIAN
	CSLANGUAGE_SINDHI
	CSLANGUAGE_SLOVAK
	CSLANGUAGE_SLOVENIAN
	CSLANGUAGE_SPANISH
	CSLANGUAGE_SWAHILI
	CSLANGUAGE_SWEDISH
	CSLANGUAGE_TAMIL
	CSLANGUAGE_TATAR
	CSLANGUAGE_TELUGU
	CSLANGUAGE_THAI
	CSLANGUAGE_TURKISH
	CSLANGUAGE_UKRAINIAN
	CSLANGUAGE_URDU
	CSLANGUAGE_UZBEK
	CSLANGUAGE_VIETNAMESE
	CSLANGUAGE_COLUMBIAN
)

func (csound CSOUND) SetLanguage(langCode Cslanguage_t) {
	C.csoundSetLanguage(C.cslanguage_t(langCode))
}

func (csound CSOUND) GetEnv(name string) string {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(C.csoundGetEnv(csound.cs, cname))
}

func (csound CSOUND) SetGlobalEnv(name, value string) int {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var cvalue *C.char
	if len(value) == 0 {
		cvalue = nil
	} else {
		cvalue = C.CString(value)
		defer C.free(unsafe.Pointer(cvalue))
	}
	return int(C.csoundSetGlobalEnv(cname, cvalue))
}

func (csound CSOUND) CreateGlobalVariable(name string, nbytes uint) int {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundCreateGlobalVariable(csound.cs, cname, C.size_t(nbytes)))
}

func (csound CSOUND) QueryGlobalVariable(name string) unsafe.Pointer {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return (unsafe.Pointer)(C.csoundQueryGlobalVariable(csound.cs, cname))
}

func (csound CSOUND) QueryGlobalVariableNoCheck(name string) unsafe.Pointer {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return (unsafe.Pointer)(C.csoundQueryGlobalVariableNoCheck(csound.cs, cname))
}

func (csound CSOUND) DestroyGlobalVariable(name string) int {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundDestroyGlobalVariable(csound.cs, cname))
}

func (csound CSOUND) RunUtility(name string, args []string) int {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	result := C.csoundRunUtility(csound.cs, cname, argc, &argv[0])
	for _, arg := range argv {
		C.free(unsafe.Pointer(arg))
	}
	return int(result)
}

func (csound CSOUND) ListUtilities() ([]string, error) {
	clist := C.csoundListUtilities(csound.cs)
	if clist == nil {
		return nil, fmt.Errorf("ListUtilities error")
	}
	n := int(C.utilityListLength(clist))
	var list = make([]string, n)
	for i := range list {
		list[i] = C.GoString(C.utilityName(clist, C.int(i)))
	}
	C.csoundDeleteUtilityList(csound.cs, clist)
	return list, nil
}

func (csound CSOUND) GetUtilityDescription(utilName string) string {
	var cname *C.char = C.CString(utilName)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(C.csoundGetUtilityDescription(csound.cs, cname))
}

func (csound CSOUND) Rand31(seedVal *int32) int32 {
	return int32(C.csoundRand31((*C.int)(seedVal)))
}

func (csound CSOUND) SeedRandMT(initKey []uint32) *C.CsoundRandMTState {
	p := C.newRandMTState()
	if len(initKey) > 1 {
		C.csoundSeedRandMT(p, (*C.uint32_t)(&initKey[0]), C.uint32_t(len(initKey)))
	} else {
		C.csoundSeedRandMT(p, nil, C.uint32_t(initKey[0]))
	}
	return p
}

func (csound CSOUND) RandMT(p *C.CsoundRandMTState) uint32 {
	return uint32(C.csoundRandMT(p))
}

func (csound CSOUND) FreeRandMTState(p *C.CsoundRandMTState) {
	C.freeRandMTState(p)
}

func (csound CSOUND) CreateCircularBuffer(numelem int) unsafe.Pointer {
	var sample MYFLT
	return unsafe.Pointer(C.csoundCreateCircularBuffer(csound.cs, C.int(numelem),
		C.int(unsafe.Sizeof(sample))))
}

func (csound CSOUND) ReadCircularBuffer(circularBuffer unsafe.Pointer, out []MYFLT,
	items int) int {
	if len(out) < items {
		return 0
	}
	return int(C.csoundReadCircularBuffer(csound.cs, circularBuffer,
		unsafe.Pointer(&out[0]), C.int(items)))
}

func (csound CSOUND) PeekCircularBuffer(circularBuffer unsafe.Pointer, out []MYFLT,
	items int) int {
	if len(out) < items {
		return 0
	}
	return int(C.csoundPeekCircularBuffer(csound.cs, circularBuffer,
		unsafe.Pointer(&out[0]), C.int(items)))
}

func (csound CSOUND) WriteCircularBuffer(circularBuffer unsafe.Pointer, inp []MYFLT,
	items int) int {
	if len(inp) < items {
		return 0
	}
	return int(C.csoundWriteCircularBuffer(csound.cs, circularBuffer,
		unsafe.Pointer(&inp[0]), C.int(items)))
}

func (csound CSOUND) FlushCircularBuffer(circularBuffer unsafe.Pointer) {
	C.csoundFlushCircularBuffer(csound.cs, circularBuffer)
}

func (csound CSOUND) DestroyCircularBuffer(circularBuffer unsafe.Pointer) {
	C.csoundDestroyCircularBuffer(csound.cs, circularBuffer)
}

func (csound CSOUND) OpenLibrary(libraryPath string) (int, unsafe.Pointer) {
	var cpath *C.char = C.CString(libraryPath)
	defer C.free(unsafe.Pointer(cpath))
	var library unsafe.Pointer
	ret := C.csoundOpenLibrary((*unsafe.Pointer)(&library), cpath)
	return int(ret), library
}

func (csound CSOUND) CloseLibrary(library unsafe.Pointer) int {
	return int(C.csoundCloseLibrary(library))
}

func (csound CSOUND) GetLibrarySymbol(library unsafe.Pointer,
	symbolName string) unsafe.Pointer {
	var cname *C.char = C.CString(symbolName)
	defer C.free(unsafe.Pointer(cname))
	return C.csoundGetLibrarySymbol(library, cname)
}
