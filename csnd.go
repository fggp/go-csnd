package csnd6

/*
#cgo CFLAGS: -DUSE_DOUBLE=1
#cgo CFLAGS: -I /usr/local/include
#cgo linux CFLAGS: -DLINUX=1
#cgo LDFLAGS: -lcsound64

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

void cMessage(CSOUND *csound, char *msg)
{
  csoundMessage(csound, "%s", msg);
}

void cMessageS(CSOUND *csound, int attr, char *msg)
{
  csoundMessageS(csound, attr, "%s", msg);
}

void noMessageCallback(CSOUND* cs, int attr, const char *format, va_list valist)
{
  // Do nothing so that Csound will not print any message,
  // leaving a clean console for our app.
  return;
}

void cNoMessage(CSOUND* cs)
{
  csoundSetMessageCallback(cs, noMessageCallback);
}

void *getOpcodeList(CSOUND *csound, int *pn)
{
  opcodeListEntry *opcodeList;
  *pn = csoundNewOpcodeList(csound, &opcodeList);
  return (void *)opcodeList;
}

void getOpcodeEntry(void *list, int n,
                    char **opname, char **outypes, char** intypes, int* flags )
{
  opcodeListEntry entry = *((opcodeListEntry *)list + n);
  *opname = entry.opname;
  *outypes = entry.outypes;
  *intypes = entry.intypes;
  *flags = entry.flags;
}

void freeOpcodeList(CSOUND *cs, void *list)
{
  csoundDisposeOpcodeList(cs, (opcodeListEntry *)list);
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

#if defined(CSOUND_SPIN_LOCK)
void csSpinLock(int32 *spinlock)
{
  csoundSpinLock(spinlock);
}
void csSpinUnLock(int32 *spinlock)
{
  csoundSpinUnLock(spinlock);
}
#else
void csSpinLock(int32 *spinlock)
{
  return;
}
void csSpinUnLock(int32 *spinlock)
{
  return;
}
#endif

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

// Error Definitions
const (
	CSOUND_SUCCESS        = 0  // Completed successfully.
	CSOUND_ERROR          = -1 // Unspecified failure.
	CSOUND_INITIALIZATION = -2 // Failed during initialization.
	CSOUND_PERFORMANCE    = -3 // Failed during performance.
	CSOUND_MEMORY         = -4 // Failed to allocate requested memory.
	CSOUND_SIGNAL         = -5 // Termination requested by SIGINT or SIGTERM.
)

// Flags for csoundInitialize().
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

// Encapsulates an opaque pointer to a Csound instance
type CSOUND struct {
	Cs (*C.CSOUND)
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
	ComputeWeights       int32 // deprecated, kept for backwards comp.
	RealtimeMode         int32 // use realtime priority mode, 0 or 1
	SampleAccurate       int32 // use sample-level score event accuracy
	SampleRateOverride   MYFLT // overriding sample rate
	ControlRateOverride  MYFLT // overriding control rate
	NchnlsOverride       int32 // overriding number of out channels
	NchnlsIoverride      int32 // overriding number of in channels
	E0dbfsOverride       MYFLT // overriding 0dbfs
	Daemon               int32 // daemon mode
	KsmpsOverride        int32 // ksmps override
	FFT_library          int32 // fft_lib
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

type TREE struct {
	t (*C.TREE)
}

// Constants used by the bus interface (csoundGetChannelPtr() etc.).
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

const (
	// message types (only one can be specified)
	CSOUNDMSG_DEFAULT  = 0x0000 // standard message
	CSOUNDMSG_ERROR    = 0x1000 // error message (initerror, perferror, etc.)
	CSOUNDMSG_ORCH     = 0x2000 // orchestra opcodes (e.g. printks)
	CSOUNDMSG_REALTIME = 0x3000 // for progress display and heartbeat characters
	CSOUNDMSG_WARNING  = 0x4000 // warning messages

	// format attributes (colors etc., use the bitwise OR of any of these:
	CSOUNDMSG_FG_BLACK   = 0x0100
	CSOUNDMSG_FG_RED     = 0x0101
	CSOUNDMSG_FG_GREEN   = 0x0102
	CSOUNDMSG_FG_YELLOW  = 0x0103
	CSOUNDMSG_FG_BLUE    = 0x0104
	CSOUNDMSG_FG_MAGENTA = 0x0105
	CSOUNDMSG_FG_CYAN    = 0x0106
	CSOUNDMSG_FG_WHITE   = 0x0107

	CSOUNDMSG_FG_BOLD      = 0x0008
	CSOUNDMSG_FG_UNDERLINE = 0x0080

	CSOUNDMSG_BG_BLACK   = 0x0200
	CSOUNDMSG_BG_RED     = 0x0210
	CSOUNDMSG_BG_GREEN   = 0x0220
	CSOUNDMSG_BG_ORANGE  = 0x0230
	CSOUNDMSG_BG_BLUE    = 0x0240
	CSOUNDMSG_BG_MAGENTA = 0x0250
	CSOUNDMSG_BG_CYAN    = 0x0260
	CSOUNDMSG_BG_GREY    = 0x0270

	//-------------------------------------------------------------------------
	CSOUNDMSG_TYPE_MASK     = 0x7000
	CSOUNDMSG_FG_COLOR_MASK = 0x0107
	CSOUNDMSG_FG_ATTR_MASK  = 0x0088
	CSOUNDMSG_BG_COLOR_MASK = 0x0270
)

/*
 * Instantiation
 */

// Initialize Csound library with specific flags.
// This function is called internally by csoundCreate(), so there is
// generally no need to use it explicitly unless you need to avoid
// default initialization that sets signal handlers and atexit() callbacks.
// Return value is zero on success, positive if initialisation was
// done already, and negative on error.
func Initialize(flags int) int {
	return int(C.csoundInitialize(C.int(flags)))
}

// Create an instance of Csound.
// Return an object with methods wrapping calls to the Csound API functions.
// The hostData parameter can be nil, or it can be a pointer to any sort of
// data; this pointer can be accessed from the Csound instance
// that is passed to callback routines.
func Create(hostData unsafe.Pointer) CSOUND {
	var cs (*C.CSOUND)
	if hostData != nil {
		cs = C.csoundCreate(hostData)
	} else {
		cs = C.csoundCreate(nil)
	}
	return CSOUND{cs}
}

// Destroy an instance of Csound.
func (csound *CSOUND) Destroy() {
	C.csoundDestroy(csound.Cs)
	csound.Cs = nil
}

// Return the version number
func (csound CSOUND) Version() string {
	n := int(C.csoundGetVersion())
	l1, l3 := n/1000, n%1000
	l2 := l3 / 10
	l3 %= 10
	return fmt.Sprintf("%d.%02d.%d", l1, l2, l3)
}

// Return the API version number
func (csound CSOUND) APIVersion() string {
	n := int(C.csoundGetAPIVersion())
	return fmt.Sprintf("%d.%02d", n/100, n%100)
}

/*
 * Performance
 */

// Parse the given orchestra from an ASCII string into a TREE.
// This can be called during performance to parse new code.
func (csound CSOUND) ParseOrc(orc string) TREE {
	var cstr *C.char = C.CString(orc)
	defer C.free(unsafe.Pointer(cstr))
	t := C.csoundParseOrc(csound.Cs, cstr)
	return TREE{t}
}

// Compile the given TREE node into structs for Csound to use.
// This can be called during performance to compile a new TREE
func (csound CSOUND) CompileTree(root TREE) int {
	result := C.csoundCompileTree(csound.Cs, root.t)
	return int(result)
}

// Asynchronous version of CompileTree().
func (csound CSOUND) CompileTreeAsync(root TREE) int {
	result := C.csoundCompileTreeAsync(csound.Cs, root.t)
	return int(result)
}

// Free the resources associated with the TREE tree.
// This function should be called whenever the TREE was
// created with ParseOrc and memory can be deallocated.
func (csound CSOUND) DeleteTree(tree TREE) {
	C.csoundDeleteTree(csound.Cs, tree.t)
}

// Parse, and compile the given orchestra from an ASCII string,
// also evaluating any global space code (i-time only)
// this can be called during performance to compile a new orchestra.
//      orc := "instr 1 \n a1 rand 0dbfs/4 \n out a1 \n"
//      csound.CompileOrc(orc)
func (csound CSOUND) CompileOrc(orc string) int {
	var cstr *C.char = C.CString(orc)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.csoundCompileOrc(csound.Cs, cstr))
}

// Async version of CompileOrc().
// The code is parsed and compiled, then placed on a queue for
// asynchronous merge into the running engine, and evaluation.
// The function returns following parsing and compilation.
func (csound CSOUND) CompileOrcAsync(orc string) int {
	var cstr *C.char = C.CString(orc)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.csoundCompileOrcAsync(csound.Cs, cstr))
}

//   Parse and compile an orchestra given on a string,
//   evaluating any global space code (i-time only).
//   On SUCCESS it returns a value passed to the
//   'return' opcode in global space.
//       code := "i1 = 2 + 2 \n return i1 \n"
//       retval := csound.EvalCode(code)
func (csound CSOUND) EvalCode(code string) MYFLT {
	var cstr *C.char = C.CString(code)
	defer C.free(unsafe.Pointer(cstr))
	return MYFLT(C.csoundEvalCode(csound.Cs, cstr))
}

// Prepare an instance of Csound for Cscore
// processing outside of running an orchestra (i.e. "standalone Cscore").
// It is an alternative to PreCompile(), Compile(), and
// Perform*() and should not be used with these functions.
//
// You must call this function before using the interface in "cscore.h"
// when you do not wish to compile an orchestra.
// Pass it the already open *C.FILE pointers to the input and
// output score files.
//
// It returns CSOUND_SUCCESS on success and CSOUND_INITIALIZATION or other
// error code if it fails.
func (csound CSOUND) InitializeCscore(insco, outsco *C.FILE) int {
	return int(C.csoundInitializeCscore(csound.Cs, insco, outsco))
}

//  Read arguments, parse and compile an orchestra, read, process and
//  load a score.
func (csound CSOUND) CompileArgs(args []string) int {
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	result := C.csoundCompileArgs(csound.Cs, argc, &argv[0])
	for _, arg := range argv {
		C.free(unsafe.Pointer(arg))
	}
	return int(result)
}

// Prepares Csound for performance.
// Normally called after compiling a csd file or an orc file, in which
// case score preprocessing is performed and performance terminates
// when the score terminates.
// However, if called before compiling a csd file or an orc file,
// score preprocessing is not performed and "i" statements are dispatched
// as real-time events, the <CsOptions> tag is ignored, and performance
// continues indefinitely or until ended using the API.
// NB: this is called internally by Compile(), therefore
// it is only required if performance is started without
// a call to that function.
func (csound CSOUND) Start() int {
	return int(C.csoundStart(csound.Cs))
}

// Compile Csound input files (such as an orchestra and score, or a CSD)
// as directed by the supplied command-line arguments,
// but does not perform them. Return a non-zero error code on failure.
// This function cannot be called during performance, and before a
// repeated call, Reset() needs to be called.
// In this (host-driven) mode, the sequence of calls should be as follows:
//       csound.Compile(args)
//       for csound.PerformBuffer() == 0 {
//       }
//       csound.Cleanup()
//       csound.Reset()
// Calls Start() internally.
func (csound CSOUND) Compile(args []string) int {
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	result := C.csoundCompile(csound.Cs, argc, &argv[0])
	for _, arg := range argv {
		C.free(unsafe.Pointer(arg))
	}
	return int(result)
}

// Compile a Csound input file (.csd file)
// which includes command-line arguments,
// but does not perform the file. Return a non-zero error code on failure.
// In this (host-driven) mode, the sequence of calls should be as follows:
//       csound.CompileCsd(fileName)
//       for csound.PerformBuffer() == 0 {
//       }
//       csound.Cleanup()
//       csound.Reset()
// NB: this function can be called during performance to
// replace or add new instruments and events.
// On a first call and if called before csound.Start(), this function
// behaves similarly to csound.Compile()
func (csound CSOUND) CompileCsd(fileName string) int {
	var cfileName *C.char = C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	return int(C.csoundCompileCsd(csound.Cs, cfileName))
}

// Compile a Csound input file contained in a string of text,
// which includes command-line arguments, orchestra, score, etc.,
// but does not perform the file. Returns a non-zero error code on failure.
// In this (host-driven) mode, the sequence of calls should be as follows:
//       csound.CompileCsdText(csound, csdText)
//       for csound.PerformBuffer() == 0 {
//       }
//       csound.Cleanup()
//       csound.Reset()
// NB: A temporary file is created, the csd_text is written to the temporary
// file, and csound.CompileCsd is called with the name of the temporary file,
// which is deleted after compilation. Behavior may vary by platform.
func (csound CSOUND) CompileCsdText(csdText string) int {
	var cCsdText *C.char = C.CString(csdText)
	defer C.free(unsafe.Pointer(cCsdText))
	return int(C.csoundCompileCsdText(csound.Cs, cCsdText))
}

// Senses input events and performs audio output until the end of score
// is reached (positive return value), an error occurs (negative return
// value), or performance is stopped by calling Stop() from another
// thread (zero return value).
//
// Note that Compile() or CompileOrc(), ReadScore(),
// Start() must be called first.
//
// In the case of zero return value, Perform() can be called again
// to continue the stopped performance. Otherwise, Reset() should be
// called to clean up after the finished or failed performance.
func (csound CSOUND) Perform() int {
	return int(C.csoundPerform(csound.Cs))
}

// Senses input events, and performs one control sample worth (ksmps) of
// audio output.
//
// Note that Compile() or CompileOrc(), ReadScore(),
// Start() must be called first.
//
// Return false during performance, and true when performance is finished.
// If called until it returns true, will perform an entire score.
// Enables external software to control the execution of Csound,
// and to synchronize performance with audio input and output.
func (csound CSOUND) PerformKsmps() int {
	return int(C.csoundPerformKsmps(csound.Cs))
}

// Performs Csound, sensing real-time and score events
// and processing one buffer's worth (-b frames) of interleaved audio.
// Return a pointer to the new output audio in 'outputAudio'
//
// Note that Compile must be called first, then call
// OutputBuffer() and InputBuffer() to get the pointer
// to csound's I/O buffers.
//
// Return false during performance, and true when performance is finished.
func (csound CSOUND) PerformBuffer() int {
	return int(C.csoundPerformBuffer(csound.Cs))
}

// Stops a Perform() running in another thread. Note that it is
// not guaranteed that Perform() has already stopped when this
// function returns.
func (csound CSOUND) Stop() {
	C.csoundStop(csound.Cs)
}

// Prints information about the end of a performance, and closes audio
// and MIDI devices.
//
// Note: after calling Cleanup(), the operation of the perform
// functions is undefined.
func (csound CSOUND) Cleanup() int {
	numSenseEvent = 0
	return int(C.csoundCleanup(csound.Cs))
}

// Resets all internal memory and state in preparation for a new performance.
// Enables external software to run successive Csound performances
// without reloading Csound. Implies Cleanup(), unless already called.
func (csound CSOUND) Reset() {
	C.csoundReset(csound.Cs)
}

// Starts the UDP server on a supplied port number.
// Returns CSOUND_SUCCESS if server has been started successfully,
// otherwise, CSOUND_ERROR.
func (csound CSOUND) UDPServerStart(port uint) int {
	return int(C.csoundUDPServerStart(csound.Cs, C.uint(port)))
}

// Returns the port number on which the server is running, or
// CSOUND_ERROR if the server is not running.
func (csound CSOUND) UDPServerStatus() int {
	return int(C.csoundUDPServerStatus(csound.Cs))
}

// Closes the UDP server, returning CSOUND_SUCCESS if the
// running server was successfully closed, CSOUND_ERROR otherwise.
func (csound CSOUND) UDPServerClose() int {
	return int(C.csoundUDPServerClose(csound.Cs))
}

/*
 * Attributes
 */

// Return the number of audio sample frames per second.
func (csound CSOUND) Sr() MYFLT {
	return MYFLT(C.csoundGetSr(csound.Cs))
}

// Return the number of control samples per second.
func (csound CSOUND) Kr() MYFLT {
	return MYFLT(C.csoundGetKr(csound.Cs))
}

// Return the number of audio sample frames per control sample.
func (csound CSOUND) Ksmps() int {
	return int(C.csoundGetKsmps(csound.Cs))
}

// Return the number of audio output channels. Set through the nchnls
// header variable in the csd file.
func (csound CSOUND) Nchnls() int {
	return int(C.csoundGetNchnls(csound.Cs))
}

// Return the number of audio input channels. Set through the
// nchnls_i header variable in the csd file. If this variable is
// not set, the value is taken from nchnls.
func (csound CSOUND) NchnlsInput() int {
	return int(C.csoundGetNchnlsInput(csound.Cs))
}

// Return the 0dBFS level of the spin/spout buffers.
func (csound CSOUND) Get0dBFS() MYFLT {
	return MYFLT(C.csoundGet0dBFS(csound.Cs))
}

// Return the A4 frequency reference.
func (csound CSOUND) A4() MYFLT {
	return MYFLT(C.csoundGetA4(csound.Cs))
}

// Return the current performance time in samples.
func (csound CSOUND) CurrentTimeSamples() int {
	return int(C.csoundGetCurrentTimeSamples(csound.Cs))
}

// Return the size of MYFLT in bytes.
func (csound CSOUND) SizeOfMYFLT() int {
	return int(C.csoundGetSizeOfMYFLT())
}

// Return host data.
func (csound CSOUND) HostData() unsafe.Pointer {
	return C.csoundGetHostData(csound.Cs)
}

// Set host data.
func (csound CSOUND) SetHostData(hostData unsafe.Pointer) {
	C.csoundSetHostData(csound.Cs, hostData)
}

// Set a single csound option (flag). Return CSOUND_SUCCESS on success.
// NB: blank spaces are not allowed
func (csound CSOUND) SetOption(option string) int {
	var coption *C.char = C.CString(option)
	defer C.free(unsafe.Pointer(coption))
	return int(C.csoundSetOption(csound.Cs, coption))
}

//  Configure Csound with a given set of parameters defined in
//  the CsoundParams structure. These parameters are the part of the
//  OPARMS struct that are configurable through command line flags.
//  The CsoundParams structure can be obtained using Params().
//  These options should only be changed before performance has started.
func (csound CSOUND) SetParams(p *CsoundParams) {
	pp := &p.DebugMode
	C.csoundSetParams(csound.Cs, (*C.CSOUND_PARAMS)(unsafe.Pointer(pp)))
}

//  Get the current set of parameters from a CSOUND instance in
//  a CsoundParams structure. See SetParams().
func (csound CSOUND) Params(p *CsoundParams) {
	pp := &p.DebugMode
	C.csoundGetParams(csound.Cs, (*C.CSOUND_PARAMS)(unsafe.Pointer(pp)))
}

// Return whether Csound is set to print debug messages sent through the
// DebugMsg() internal API function.
func (csound CSOUND) Debug() bool {
	return C.csoundGetDebug(csound.Cs) != 0
}

// Set whether Csound prints debug messages from the DebugMsg() internal
// API function.
func (csound CSOUND) SetDebug(debug bool) {
	C.csoundSetDebug(csound.Cs, cbool(debug))
}

/*
 * General Input/Output
 */

// Return the audio output name (-o).
func (csound CSOUND) OutputName() string {
	return C.GoString(C.csoundGetOutputName(csound.Cs))
}

//  Set output destination, type and format
//  type can be one of  "wav","aiff", "au","raw", "paf", "svx", "nist", "voc",
//  "ircam","w64","mat4", "mat5", "pvf","xi", "htk","sds","avr","wavex","sd2",
//  "flac", "caf","wve","ogg","mpc2k","rf64", or nil (use default or
//  realtime IO).
//  format can be one of "alaw", "schar", "uchar", "float", "double", "long",
//  "short", "ulaw", "24bit", "vorbis", or nil (use default or realtime IO).
//   For RT audio, use DeviceId from CsoundAudioDevice for a given audio device.
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
	C.csoundSetOutput(csound.Cs, cname, ctype, cformat)
}

// Get output type and format.
func (csound CSOUND) OutputFormat() (otype, format string) {
	type_ := make([]byte, 6)
	format_ := make([]byte, 8)
	ctype := (*C.char)(unsafe.Pointer(&type_[0]))
	cformat := (*C.char)(unsafe.Pointer(&format_[0]))
	C.csoundGetOutputFormat(csound.Cs, ctype, cformat)
	otype = C.GoString(ctype)
	format = C.GoString(cformat)
	return
}

// Set input source.
func (csound CSOUND) SetInput(name string) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetInput(csound.Cs, cname)
}

// Set MIDI input device name/number.
func (csound CSOUND) SetMIDIInput(name string) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetMIDIInput(csound.Cs, cname)
}

// Set MIDI file input name.
func (csound CSOUND) SetMIDIFileInput(name string) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetMIDIFileInput(csound.Cs, cname)
}

// Set MIDI output device name/number.
func (csound CSOUND) SetMIDIOutput(name string) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetMIDIOutput(csound.Cs, cname)
}

// Set MIDI file output name.
func (csound CSOUND) SetMIDIFileOutput(name string) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetMIDIFileOutput(csound.Cs, cname)
}

/*
 * Realtime Audio I/O
 */

// Set the current RT audio module.
func (csound CSOUND) SetRTAudioModule(module string) {
	var cmodule *C.char = C.CString(module)
	defer C.free(unsafe.Pointer(cmodule))
	C.csoundSetRTAudioModule(csound.Cs, cmodule)
}

// Retrieve a module name and type ("audio" or "midi") given a
// number. Modules are added to list as csound loads them. Return
// CSOUND_SUCCESS on success and CSOUND_ERROR if module number
// was not found
//
//   var name, mtype string
//   err := CSOUND_SUCCESS
//   n := 0;
//   for err != CSOUND_ERROR {
//       name, mtype, err = csound.Module(n++)
//       fmt.Printf("Module %d:  %s (%s)\n", n, name, mtype)
//   }
func (csound CSOUND) Module(number int) (name, mtype string, error int) {
	var cname, ctype *C.char
	cerror := C.csoundGetModule(csound.Cs, C.int(number), &cname, &ctype)
	name = C.GoString(cname)
	mtype = C.GoString(ctype)
	error = int(cerror)
	return
}

// Return the number of samples in Csound input buffer.
func (csound CSOUND) InputBufferSize() int {
	return int(C.csoundGetInputBufferSize(csound.Cs))
}

// Return the number of samples in Csound output buffer.
func (csound CSOUND) OutputBufferSize() int {
	return int(C.csoundGetOutputBufferSize(csound.Cs))
}

// Return the Csound audio input buffer as a []MYFLT.
// Enable external software to write audio into Csound before calling
// PerformBuffer.
func (csound CSOUND) InputBuffer() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetInputBuffer(csound.Cs))
	length := int(C.csoundGetInputBufferSize(csound.Cs))
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(buffer))
	return slice
}

// Return the Csound audio output buffer as a []MYFLT.
// Enable external software to read audio from Csound after calling
// PerformBuffer.
func (csound CSOUND) OutputBuffer() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetOutputBuffer(csound.Cs))
	length := int(C.csoundGetOutputBufferSize(csound.Cs))
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(buffer))
	return slice
}

// Return the Csound audio input working buffer (spin) as a []MYFLT.
// Enable external software to write audio into Csound before calling
// PerformKsmps.
func (csound CSOUND) Spin() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetSpin(csound.Cs))
	length := csound.Ksmps() * csound.Nchnls()
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(buffer))
	return slice
}

// Clear the input buffer (spin).
func (csound CSOUND) ClearSpin() {
	C.csoundClearSpin(csound.Cs)
}

// Add the indicated sample into the audio input working buffer (spin);
// this only ever makes sense before calling PerformKsmps().
// The frame and channel must be in bounds relative to ksmps and nchnls.
// NB: the spin buffer needs to be cleared at every k-cycle by calling
// ClearSpin().
func (csound CSOUND) AddSpinSample(frame, channel int, sample MYFLT) {
	C.csoundAddSpinSample(csound.Cs, C.int(frame), C.int(channel), cMYFLT(sample))
}

// Set the audio input working buffer (spin) to the indicated sample
// this only ever makes sense before calling PerformKsmps().
// The frame and channel must be in bounds relative to ksmps and nchnls.
func (csound CSOUND) SetSpinSample(frame, channel int, sample MYFLT) {
	C.csoundSetSpinSample(csound.Cs, C.int(frame), C.int(channel), cMYFLT(sample))
}

// Return the Csound audio output working buffer (spout) as a []MYFLT.
// Enable external software to read audio from Csound after calling
// PerformKsmps.
func (csound CSOUND) Spout() []MYFLT {
	buffer := (*MYFLT)(C.csoundGetSpout(csound.Cs))
	length := csound.Ksmps() * csound.Nchnls()
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(buffer))
	return slice
}

// Return the indicated sample from the Csound audio output
// working buffer (spout); only ever makes sense after calling
// PerformKsmps(). The frame and channel must be in bounds
// relative to ksmps and nchnls.
func (csound CSOUND) SpoutSample(frame, channel int) MYFLT {
	return MYFLT(C.csoundGetSpoutSample(csound.Cs, C.int(frame), C.int(channel)))
}

// Return a pointer to user data for real time audio input.
func (csound CSOUND) RtRecordUserData() unsafe.Pointer {
	return unsafe.Pointer(C.csoundGetRtRecordUserData(csound.Cs))
}

// Return a pointer to user data for real time audio output.
func (csound CSOUND) RtPlaydUserData() unsafe.Pointer {
	return unsafe.Pointer(C.csoundGetRtPlayUserData(csound.Cs))
}

// Calling this function with a non-zero 'state' value between
// Create() and the start of performance will disable all default
// handling of sound I/O by the Csound library, allowing the host
// application to use the spin/spout/input/output buffers directly.
// For applications using spin/spout, bufSize should be set to 0.
// If 'bufSize' is greater than zero, the buffer size (-b) will be
// set to the integer multiple of ksmps that is nearest to the value
// specified.
func (csound CSOUND) SetHostImplementedAudioIO(state, bufSize int) {
	C.csoundSetHostImplementedAudioIO(csound.Cs, C.int(state), C.int(bufSize))
}

// This function can be called to obtain a list of available
// input or output audio devices (isOutput=true for out
// devices, false for in devices).
//
//   list := csound.AudioDevList(true)
//   for i := range list {
//       fmt.Printf("%d: %s (%s), %d chan\n",
//             i, list[i].DeviceId, list[i].DeviceName, list[i].MaxNchnls)
//   }
func (csound CSOUND) AudioDevList(isOutput bool) []CsoundAudioDevice {
	cflag := cbool(isOutput)
	n := C.csoundGetAudioDevList(csound.Cs, nil, cflag)
	devs := C.getAudioDevList(csound.Cs, n, cflag)
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

// Set the current MIDI IO module.
func (csound CSOUND) SetMIDIModule(module string) {
	var cmodule *C.char = C.CString(module)
	defer C.free(unsafe.Pointer(cmodule))
	C.csoundSetMIDIModule(csound.Cs, cmodule)
}

// Call this function with state true if the host is implementing
// MIDI via the callbacks below.
func (csound CSOUND) SetHostImplementedMIDIIO(state bool) {
	C.csoundSetHostImplementedMIDIIO(csound.Cs, cbool(state))
}

// This function can be called to obtain a list of available
// input or output midi devices. (see also AudioDevList())
func (csound CSOUND) MidiDevList(isOutput bool) []CsoundMidiDevice {
	cflag := cbool(isOutput)
	n := C.csoundGetMIDIDevList(csound.Cs, nil, cflag)
	devs := C.getMidiDevList(csound.Cs, n, cflag)
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

// Read, preprocess, and load a score from an ASCII string.
// It can be called repeatedly, with the new score events
// being added to the currently scheduled ones.
func (csound CSOUND) ReadScore(str string) int {
	var cstr *C.char = C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	return int(C.csoundReadScore(csound.Cs, cstr))
}

// Asynchronous version of ReadScore().
func (csound CSOUND) ReadScoreAsync(str string) {
	var cstr *C.char = C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.csoundReadScoreAsync(csound.Cs, cstr)
}

// Return the current score time in seconds
// since the beginning of performance.
func (csound CSOUND) ScoreTime() float64 {
	return float64(C.csoundGetScoreTime(csound.Cs))
}

// Tell whether Csound score events are performed or not, independently
// of real-time MIDI events (see SetScorePending()).
func (csound CSOUND) IsScorePending() bool {
	return C.csoundIsScorePending(csound.Cs) != 0
}

// Set whether Csound score events are performed or not (real-time
// events will continue to be performed). Can be used by external software,
// such as a VST host, to turn off performance of score events (while
// continuing to perform real-time events), for example to
// mute a Csound score while working on other tracks of a piece, or
// to play the Csound instruments live.
func (csound CSOUND) SetScorePending(pending bool) {
	C.csoundSetScorePending(csound.Cs, cbool(pending))
}

// Return the score time beginning at which score events will
// actually immediately be performed (see SetScoreOffsetSeconds()).
func (csound CSOUND) ScoreOffsetSeconds() MYFLT {
	return MYFLT(C.csoundGetScoreOffsetSeconds(csound.Cs))
}

// Csound score events prior to the specified time are not performed, and
// performance begins immediately at the specified time (real-time events
// will continue to be performed as they are received).
// Can be used by external software, such as a VST host,
// to begin score performance midway through a Csound score,
// for example to repeat a loop in a sequencer, or to synchronize
// other events with the Csound score.
func (csound CSOUND) SetScoreOffsetSeconds(time MYFLT) {
	C.csoundSetScoreOffsetSeconds(csound.Cs, cMYFLT(time))
}

// Rewind a compiled Csound score to the time specified with
// SetScoreOffsetSeconds().
func (csound CSOUND) RewindScore() {
	C.csoundRewindScore(csound.Cs)
}

// Sort score file 'inFile' and write the result to 'outFile'.
// The Csound instance should be initialised with PreCompile()
// before calling this function, and Reset() should be called
// after sorting the score to clean up. On success, zero is returned.
func (csound CSOUND) ScoreSort(inFile, outFile *C.FILE) int {
	return int(C.csoundScoreSort(csound.Cs, inFile, outFile))
}

// Extract from 'inFile', controlled by 'extractFile', and write
// the result to 'outFile'. The Csound instance should be initialised
// with PreCompile() before calling this function, and Reset()
// should be called after score extraction to clean up.
// The return value is zero on success.
func (csound CSOUND) ScoreExtract(inFile, outFile, extractFile *C.FILE) int {
	return int(C.csoundScoreExtract(csound.Cs, inFile, outFile, extractFile))
}

/*
 * Messages and Text
 */

// Displays an informal message.
func (csound CSOUND) Message(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	var cval *C.char = C.CString(s)
	defer C.free(unsafe.Pointer(cval))
	C.cMessage(csound.Cs, cval)
}

// Print message with special attributes (see const above for the list of
// available attributes). With attr=0, csoundMessageS() is identical to
// csoundMessage().
func (csound CSOUND) MessageS(attr int, format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	var cval *C.char = C.CString(s)
	defer C.free(unsafe.Pointer(cval))
	C.cMessageS(csound.Cs, C.int(attr), cval)
}

// Return the Csound message level (from 0 to 231).
func (csound CSOUND) MessageLevel() int {
	return int(C.csoundGetMessageLevel(csound.Cs))
}

// Set the Csound message level (from 0 to 231).
func (csound CSOUND) SetMessageLevel(messageLevel int) {
	C.csoundSetMessageLevel(csound.Cs, C.int(messageLevel))
}

// This is not an API function. It passes to csoundSetMessageCallback a
// callback function that does nothing, so that Csound will not print
// any message.
func (csound CSOUND) NoMessage() {
	C.cNoMessage(csound.Cs)
}

// Create a buffer for storing messages printed by Csound.
//
// Should be called after creating a Csound instance and the buffer
// can be freed by calling DestroyMessageBuffer() before
// deleting the Csound instance. You will generally want to call
// Cleanup() to make sure the last messages are flushed to
// the message buffer before destroying Csound.
//
// If 'toStdOut' is true, the messages are also printed to
// stdout and stderr (depending on the type of the message),
// in addition to being stored in the buffer.
//
// Using the message buffer ties up the internal message callback, so
// SetMessageCallback should not be called after creating the
// message buffer.
func (csound CSOUND) CreateMessageBuffer(toStdOut bool) {
	C.csoundCreateMessageBuffer(csound.Cs, cbool(toStdOut))
}

// Return the first message from the buffer.
func (csound CSOUND) FirstMessage() string {
	cmsg := C.csoundGetFirstMessage(csound.Cs)
	return C.GoString(cmsg)
}

// Return the attribute parameter (see msg_attr.h) of the first message
// in the buffer.
func (csound CSOUND) FirstMessageAttr() int {
	return int(C.csoundGetFirstMessageAttr(csound.Cs))
}

// Remove the first message from the buffer.
func (csound CSOUND) PopFirstMessage() {
	C.csoundPopFirstMessage(csound.Cs)
}

// Return the number of pending messages in the buffer.
func (csound CSOUND) MessageCnt() int {
	return int(C.csoundGetMessageCnt(csound.Cs))
}

// Release all memory used by the message buffer.
func (csound CSOUND) DestroyMessageBuffer() {
	C.csoundDestroyMessageBuffer(csound.Cs)
}

/*
 * Channels, Control and Events
 */

// Return a pointer to the specified channel of the bus as a []MYFLT,
// creating the channel first if it does not exist yet.
// 'type' must be the bitwise OR of exactly one of the following values,
//   CSOUND_CONTROL_CHANNEL
//     control data (one MYFLT value)
//   CSOUND_AUDIO_CHANNEL
//     audio data (GetKsmps() MYFLT values)
//   CSOUND_STRING_CHANNEL
//     string data (MYFLT values with enough space to store
//     ChannelDatasize() characters, including the
//     NULL character at the end of the string)
// and at least one of these:
//   CSOUND_INPUT_CHANNEL
//   CSOUND_OUTPUT_CHANNEL
//
// If the channel already exists, it must match the data type
// (control, audio, or string), however, the input/output bits are
// OR'd with the new value. Note that audio and string channels
// can only be created after calling Compile(), because the
// storage size is not known until then.
//
// The returned error is nil on success, or an error message,
//   "Not enough memory for allocating the channel" (CSOUND_MEMORY)
//   "The specified name or type is invalid" (CSOUND_ERROR)
// or, if a channel with the same name but incompatible type
// already exists, the type of the existing channel. In the case
// of any non-nil error value, the channel pointer is set to nil.
//
// Note: to find out the type of a channel without actually
// creating or changing it, set 'chnType' to zero, so that the error
// value will be either the type of the channel, or CSOUND_ERROR
// if it does not exist.
//
// Operations on the channel pointer are not thread-safe by default. The host is
// required to take care of threadsafety by
//   1) with control channels use __sync_fetch_and_add() or
//      __sync_fetch_and_or() gcc atomic builtins to get or set a channel,
//      if available.
//   2) For string and audio channels (and controls if option 1 is not
//      available), retrieve the channel lock with ChannelLock()
//      and use SpinLock() and SpinUnLock() to protect access
//      to the channel.
// See Top/threadsafe.c in the Csound library sources for
// examples. Optionally, use the channel get/set functions
// which are threadsafe by default.
func (csound CSOUND) ChannelPtr(name string, chnType int) ([]MYFLT, error) {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var ptr *MYFLT
	var length int
	switch chnType & CSOUND_CHANNEL_TYPE_MASK {
	case CSOUND_CONTROL_CHANNEL:
		length = 1
	case CSOUND_AUDIO_CHANNEL:
		length = int(C.csoundGetKsmps(csound.Cs))
	case CSOUND_STRING_CHANNEL:
		length = int(C.csoundGetChannelDatasize(csound.Cs, cname))
	default:
		return nil, fmt.Errorf("%d is not a valid channel type", chnType)
	}
	ret := C.csoundGetChannelPtr(csound.Cs, cppMYFLT(&ptr), cname, C.int(chnType))
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
		return nil, fmt.Errorf("The specified channel name or type is invalid")
	case CSOUND_CONTROL_CHANNEL:
		return nil, fmt.Errorf("A control channel named %s already exists", name)
	case CSOUND_AUDIO_CHANNEL:
		return nil, fmt.Errorf("An audio channel named %s already exists", name)
	case CSOUND_STRING_CHANNEL:
		return nil, fmt.Errorf("A string channel named %s already exists", name)
	default:
		return nil, fmt.Errorf("Unknown error")
	}
}

// Return a list of allocated channels. A ControlChannelInfo
// structure contains the channel characteristics.
// The error value is nil, or a CSOUND_MEMORY message, if there is not enough
// memory for allocating the list. In the case of no channels or an error, the
// list is set to nil.
//
// Notes: The list will become inconsistant
// after calling Reset().
func (csound CSOUND) ListChannels() ([]ControlChannelInfo, error) {
	var lst *C.controlChannelInfo_t
	n := int(C.csoundListChannels(csound.Cs, &lst))
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
		C.csoundDeleteChannelList(csound.Cs, lst)
		return list, nil
	}
}

// Set parameters hints for a control channel. These hints have no internal
// function but can be used by front ends to construct GUIs or to constrain
// values. See the ControlChannelHints structure for details.
// Returns zero on success, or a non-zero error code on failure:
//   CSOUND_ERROR:  the channel does not exist, is not a control channel,
//                  or the specified parameters are invalid
//   CSOUND_MEMORY: could not allocate memory
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
	return int(C.csoundSetControlChannelHints(csound.Cs, cname, chints))
}

// Return special parameters (assuming there are any) of a control channel,
// previously set with SetControlChannelHints() or the chnparams
// opcode.
// If the channel exists, is a control channel, the channel hints
// are stored in the ControlChannelHints structure.
//
// The return value is zero if the channel exists and is a control
// channel, otherwise, an error code is returned.
func (csound CSOUND) ControlChannelHints(name string) (ControlChannelHints, int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var chints C.controlChannelHints_t
	ret := C.csoundGetControlChannelHints(csound.Cs, cname, &chints)
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

// Recover a pointer to a lock for the specified channel called 'name'.
// The returned lock can be locked/unlocked  with the SpinLock()
// and SpinUnLock() functions.
// Return the address of the lock or nil if the channel does not exist
func (csound CSOUND) ChannelLock(name string) *C.int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.csoundGetChannelLock(csound.Cs, cname)
}

// Retrieve the value of control channel identified by name.
// The error (or success) code
// finding or accessing the channel is returned as well.
func (csound CSOUND) ControlChannel(name string) (MYFLT, int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var err C.int
	val := MYFLT(C.csoundGetControlChannel(csound.Cs, cname, &err))
	return val, int(err)
}

// Set the value of control channel identified by name.
func (csound CSOUND) SetControlChannel(name string, val MYFLT) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.csoundSetControlChannel(csound.Cs, cname, cMYFLT(val))
}

// Copy the audio channel identified by name into array
// samples which should contain enough memory for ksmps MYFLTs.
func (csound CSOUND) AudioChannel(name string, samples []MYFLT) {
	if len(samples) >= csound.Ksmps() {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))
		C.csoundGetAudioChannel(csound.Cs, cname, cpMYFLT(&samples[0]))
	}
}

// Set the audio channel identified by name with data from array
// samples which should contain at least ksmps MYFLTs.
func (csound CSOUND) SetAudioChannel(name string, samples []MYFLT) {
	if len(samples) >= csound.Ksmps() {
		cname := C.CString(name)
		defer C.free(unsafe.Pointer(cname))
		C.csoundSetAudioChannel(csound.Cs, cname, cpMYFLT(&samples[0]))
	}
}

// Return a copy of the string channel identified by name.
func (csound CSOUND) StringChannel(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	size := C.csoundGetChannelDatasize(csound.Cs, cname)
	cstr := (*C.char)(C.malloc(C.size_t(size)))
	defer C.free(unsafe.Pointer(cstr))
	C.csoundGetStringChannel(csound.Cs, cname, cstr)
	return C.GoString(cstr)
}

// Set the string channel identified by name with str.
func (csound CSOUND) SetStringChannel(name, str string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.csoundSetStringChannel(csound.Cs, cname, cstr)
}

// Return the size of data stored in a channel; for string channels
// this might change if the channel space gets reallocated.
func (csound CSOUND) ChannelDatasize(name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundGetChannelDatasize(csound.Cs, cname))
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

// Send a PVSDATEX fin to the pvsin opcode (f-rate) for channel 'name'.
// Return zero on success, CSOUND_ERROR if the index is invalid or
// fsig framesizes are incompatible,
// CSOUND_MEMORY if there is not enough memory to extend the bus.
func (csound CSOUND) SetPvsChannel(fin *PVSDATEXT, name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundSetPvsChannel(csound.Cs, fin.CStruct, cname))
}

// Receive a PVSDAT fout from the pvsout opcode (f-rate) at channel 'name'.
// Return zero on success, CSOUND_ERROR if the index is invalid or
// if fsig framesizes are incompatible,
// CSOUND_MEMORY if there is not enough memory to extend the bus.
func (csound CSOUND) PvsChannel(fout *PVSDATEXT, name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundGetPvsChannel(csound.Cs, fout.CStruct, cname))
}

// Send a new score event. 'eventType' is the score event type ('a', 'i', 'q',
// 'f', or 'e').
// 'pFields' is slice of MYFLT
// with all the pfields for this event, starting with the p1 value
// specified in pFields[0].
func (csound CSOUND) ScoreEvent(eventType byte, pFields []MYFLT) int {
	return int(C.csoundScoreEvent(csound.Cs, C.char(eventType),
		cpMYFLT(&pFields[0]), C.long(len(pFields))))
}

// Asynchronous version of ScoreEvent().
func (csound CSOUND) ScoreEventAsync(eventType byte, pFields []MYFLT) {
	C.csoundScoreEventAsync(csound.Cs, C.char(eventType),
		cpMYFLT(&pFields[0]), C.long(len(pFields)))
}

// Like ScoreEvent(), this function inserts a score event, but
// at absolute time with respect to the start of performance, or from an
// offset set with timeOfs.
func (csound CSOUND) ScoreEventAbsolute(eventType byte, pFields []MYFLT,
	timeOfs float64) int {
	return int(C.csoundScoreEventAbsolute(csound.Cs, C.char(eventType),
		cpMYFLT(&pFields[0]), C.long(len(pFields)),
		C.double(timeOfs)))
}

// Asynchronous version of ScoreEventAbsolute().
func (csound CSOUND) ScoreEventAbsoluteAsync(eventType byte, pFields []MYFLT,
	timeOfs float64) {
	C.csoundScoreEventAbsoluteAsync(csound.Cs, C.char(eventType),
		cpMYFLT(&pFields[0]), C.long(len(pFields)),
		C.double(timeOfs))
}

// Input a string (as if from a console), used for line events.
func (csound CSOUND) InputMessage(message string) {
	var cmsg *C.char = C.CString(message)
	defer C.free(unsafe.Pointer(cmsg))
	C.csoundInputMessage(csound.Cs, cmsg)
}

// Asynchronous version of InputMessage().
func (csound CSOUND) InputMessageAsync(message string) {
	var cmsg *C.char = C.CString(message)
	defer C.free(unsafe.Pointer(cmsg))
	C.csoundInputMessageAsync(csound.Cs, cmsg)
}

// Kill off one or more running instances of an instrument identified
// by instr (number) or instrName (name). If instrName is nil, the
// instrument number is used.
// Mode is a sum of the following values:
//   0, 1, 2: kill all instances (0), oldest only (1), or newest (2)
//   4: only turnoff notes with exactly matching (fractional) instr number
//   8: only turnoff notes with indefinite duration (p3 < 0 or MIDI)
// allowRelease: if true, the killed instances are allowed to release.
func (csound CSOUND) KillInstance(instr MYFLT, instrName string, mode int,
	allowRelease bool) int {
	var cname *C.char
	if len(instrName) > 0 {
		cname = C.CString(instrName)
		defer C.free(unsafe.Pointer(cname))
	} else {
		cname = nil
	}
	return int(C.csoundKillInstance(csound.Cs, cMYFLT(instr), cname, C.int(mode),
		cbool(allowRelease)))
}

// Set the ASCII code of the most recent key pressed.
// This value is used by the 'sensekey' opcode if a callback
// for returning keyboard events is not set (see
// RegisterKeyboardCallback()).
func (csound CSOUND) KeyPress(c byte) {
	C.csoundKeyPress(csound.Cs, C.char(c))
}

/*
 * Tables
 */

// Return the length of a function table (not including the guard point),
// or -1 if the table does not exist.
func (csound CSOUND) TableLength(table int) int {
	return int(C.csoundTableLength(csound.Cs, C.int(table)))
}

// Return the value of a slot in a function table.
// The table number and index are assumed to be valid.
func (csound CSOUND) TableGet(table, index int) MYFLT {
	return MYFLT(C.csoundTableGet(csound.Cs, C.int(table), C.int(index)))
}

// Set the value of a slot in a function table.
// The table number and index are assumed to be valid.
func (csound CSOUND) TableSet(table, index int, value MYFLT) {
	C.csoundTableSet(csound.Cs, C.int(table), C.int(index), cMYFLT(value))
}

// Copy the contents of a function table into a supplied array dest.
// The table number is assumed to be valid, and the destination needs to
// have sufficient space to receive all the function table contents.
func (csound CSOUND) TableCopyOut(table int, dest []MYFLT) {
	cdest := cpMYFLT(&dest[0])
	C.csoundTableCopyOut(csound.Cs, C.int(table), cdest)
}

// Asynchronous version of TableCopyOut().
func (csound CSOUND) TableCopyOutAsync(table int, dest []MYFLT) {
	cdest := cpMYFLT(&dest[0])
	C.csoundTableCopyOutAsync(csound.Cs, C.int(table), cdest)
}

// Copy the contents of an array src into a given function table.
// The table number is assumed to be valid, and the table needs to
// have sufficient space to receive all the array contents.
func (csound CSOUND) TableCopyIn(table int, src []MYFLT) {
	csrc := cpMYFLT(&src[0])
	C.csoundTableCopyIn(csound.Cs, C.int(table), csrc)
}

// Asynchronous version of TableCopyIn().
func (csound CSOUND) TableCopyInAsync(table int, src []MYFLT) {
	csrc := cpMYFLT(&src[0])
	C.csoundTableCopyInAsync(csound.Cs, C.int(table), csrc)
}

// Return a pointer to function table 'tableNum' as a []MYFLT.
// If the table does not exist, the pointer is set to nil and
// an error is returned.
func (csound CSOUND) Table(tableNum int) ([]MYFLT, error) {
	var tablePtr *MYFLT
	length := int(C.csoundGetTable(csound.Cs, cppMYFLT(&tablePtr), C.int(tableNum)))
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

// Return a pointer to the arguments used to generate
// function table 'tableNum' as a []MYFLT.
// If the table does not exist, the pointer is set to nil and
// an error is returned.
// NB: the argument list starts with the GEN number and is followed by
// its parameters. eg. f 1 0 1024 10 1 0.5  yields the list [10.0, 1.0, 0.5]
//
func (csound CSOUND) TableArgs(tableNum int) ([]MYFLT, error) {
	var argsPtr *MYFLT
	length := int(C.csoundGetTableArgs(csound.Cs, cppMYFLT(&argsPtr), C.int(tableNum)))
	if length == -1 {
		return nil, fmt.Errorf("Function table %d does not exist", tableNum)
	}
	var slice []MYFLT
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(unsafe.Pointer(argsPtr))
	return slice, nil
}

// Check if a given GEN number num is a named GEN.
// If so, it returns the string length.
// Otherwise it returns 0.
func (csound CSOUND) IsNamedGEN(num int) int {
	return int(C.csoundIsNamedGEN(csound.Cs, C.int(num)))
}

// Get the GEN name from a number num, if this is a named GEN.
// The final parameter is the max len of the string.
func (csound CSOUND) NamedGEN(num, namelen int) string {
	name := make([]byte, namelen)
	cname := (*C.char)(unsafe.Pointer(&name[0]))
	C.csoundGetNamedGEN(csound.Cs, C.int(num), cname, C.int(namelen))
	return string(name[:namelen])
}

/*
 * Function table display
 */

// Tell Csound whether external graphic table display is supported.
// Return the previously set value (initially zero).
func (csound CSOUND) SetIsGraphable(isGraphable int) int {
	return int(C.csoundSetIsGraphable(csound.Cs, C.int(isGraphable)))
}

/*
 * Opcodes
 */

type NamedGen struct {
	Name string
	Num  int
}

// Find the list of named gens.
func (csound CSOUND) NamedGens() []NamedGen {
	n := int(C.getNumNamedGens(csound.Cs))
	if n == 0 {
		return nil
	}
	namedGens := make([]NamedGen, n)
	p := C.csoundGetNamedGens(csound.Cs)
	var name *C.char
	var num C.int
	for i := range namedGens {
		p = C.getNamedGen(csound.Cs, p, &name, &num)
		namedGens[i].Name = C.GoString(name)
		namedGens[i].Num = int(num)
	}
	return namedGens
}

type OpcodeListEntry struct {
	Opname  string
	Outypes string
	Intypes string
	Flags   int
}

// Get an alphabetically sorted list of all opcodes.
// Should be called after externals are loaded by Compile().
// Return the number of opcodes, or a negative error code on failure.
func (csound CSOUND) OpcodeList() []OpcodeListEntry {
	var opcodeList unsafe.Pointer
	var length int
	opcodeList = C.getOpcodeList(csound.Cs, (*C.int)(unsafe.Pointer(&length)))
	if length < 0 {
		return nil
	}
	var list = make([]OpcodeListEntry, length)
	var opname, outypes, intypes *C.char
	var flags C.int
	for i := range list {
		C.getOpcodeEntry(opcodeList, C.int(i), &opname, &outypes, &intypes, &flags)
		list[i].Opname = C.GoString(opname)
		list[i].Outypes = C.GoString(outypes)
		list[i].Intypes = C.GoString(intypes)
		list[i].Flags = int(flags)
	}
	C.freeOpcodeList(csound.Cs, unsafe.Pointer(opcodeList))
	return list
}

// TODO
//  AppendOpcode

/*
 * Threading and concurrency
 */

// TODO
//    createThread

// Return the ID of the currently executing thread,
// or nil for failure.
//
// NOTE: The return value can be used as a pointer
// to a thread object, but it should not be compared
// as a pointer. The pointed to values should be compared,
// and the user must free the pointer after use.
func (csound CSOUND) CurrentThreadId() unsafe.Pointer {
	return unsafe.Pointer(C.csoundGetCurrentThreadId())
}

// Wait until the indicated thread's routine has finished.
// Return the value returned by the thread routine.
func (csound CSOUND) JoinThread(thread unsafe.Pointer) uintptr {
	return uintptr(C.csoundJoinThread(thread))
}

// Create and return a monitor object, or nil if not successful.
// The object is initially in signaled (notified) state.
func (csound CSOUND) CreateThreadLock() unsafe.Pointer {
	return unsafe.Pointer(C.csoundCreateThreadLock())
}

// Wait on the indicated monitor object for the indicated period.
// The function returns either when the monitor object is notified,
// or when the period has elapsed, whichever is sooner; in the first case,
// zero is returned.
//
// If 'milliseconds' is zero and the object is not notified, the function
// will return immediately with a non-zero status.
func (csound CSOUND) WaitThreadLock(lock unsafe.Pointer, ms uint) int {
	return int(C.csoundWaitThreadLock(lock, C.size_t(ms)))
}

// Wait on the indicated monitor object until it is notified.
// This function is similar to WaitThreadLock() with an infinite
// wait time, but may be more efficient.
func (csound CSOUND) WaitThreadLockNoTimeout(lock unsafe.Pointer) {
	C.csoundWaitThreadLockNoTimeout(lock)
}

// Notify the indicated monitor object.
func (csound CSOUND) NotifyThreadLock(lock unsafe.Pointer) {
	C.csoundNotifyThreadLock(lock)
}

// Destroy the indicated monitor object.
func (csound CSOUND) DestroyThreadLock(lock unsafe.Pointer) {
	C.csoundDestroyThreadLock(lock)
}

// Create and return a mutex object, or nil if not successful.
//
// Mutexes can be faster than the more general purpose monitor objects
// returned by CreateThreadLock() on some platforms, and can also
// be recursive, but the result of unlocking a mutex that is owned by
// another thread or is not locked is undefined.
//
// If 'isRecursive' is true, the mutex can be re-locked multiple
// times by the same thread, requiring an equal number of unlock calls;
// otherwise, attempting to re-lock the mutex results in undefined
// behavior.
//
// Note: the handles returned by CreateThreadLock() and
// CreateMutex() are not compatible.
func (csound CSOUND) CreateMutex(isRecursive bool) unsafe.Pointer {
	return C.csoundCreateMutex(cbool(isRecursive))
}

// Acquire the indicated mutex object; if it is already in use by
// another thread, the function waits until the mutex is released by
// the other thread.
func (csound CSOUND) LockMutex(mutex unsafe.Pointer) {
	C.csoundLockMutex(mutex)
}

// Acquire the indicated mutex object and return zero, unless it is
// already in use by another thread, in which case a non-zero value is
// returned immediately, rather than waiting until the mutex becomes
// available.
//
// Note: this function may be unimplemented on Windows.
func (csound CSOUND) LockMutexNoWait(mutex unsafe.Pointer) int {
	return int(C.csoundLockMutexNoWait(mutex))
}

// Release the indicated mutex object, which should be owned by
// the current thread, otherwise the operation of this function is
// undefined. A recursive mutex needs to be unlocked as many times
// as it was locked previously.
func (csound CSOUND) UnlockMutex(mutex unsafe.Pointer) {
	C.csoundUnlockMutex(mutex)
}

// Destroy the indicated mutex object. Destroying a mutex that
// is currently owned by a thread results in undefined behavior.
func (csound CSOUND) DestroyMutex(mutex unsafe.Pointer) {
	C.csoundDestroyMutex(mutex)
}

// Create a Thread Barrier. Max value parameter should be equal to
// number of child threads using the barrier plus one for the
// master thread.
func (csound CSOUND) CreateBarrier(max uint) unsafe.Pointer {
	return C.csoundCreateBarrier(C.uint(max))
}

// Destroy a Thread Barrier.
func (csound CSOUND) DestroyBarrier(barrier unsafe.Pointer) int {
	return int(C.csoundDestroyBarrier(barrier))
}

// Wait on the thread barrier.
func (csound CSOUND) WaitBarrier(barrier unsafe.Pointer) int {
	return int(C.csoundWaitBarrier(barrier))
}

//func (csound CSOUND) CreateCondVar
//func (csound CSOUND) CondWait
//func (csound CSOUND) CondSignal

// Wait for at least the specified number of milliseconds,
// yielding the CPU to other threads.
func (csound CSOUND) Sleep(ms uint) {
	C.csoundSleep(C.size_t(ms))
}

// Lock the specified spinlock.
// If the spinlock is not locked, lock it and return;
// if is is locked, wait until it is unlocked, then lock it and return.
// Uses atomic compare and swap operations that are safe across processors
// and safe for out of order operations,
// and which are more efficient than operating system locks.
// Use spinlocks to protect access to shared data, especially in functions
// that do little more than read or write such data, for example:
//
//   var lock int32
//   func write(cs CSOUND, frames, signal) {
//       cs.spinLock(&lock)
//       for frame := range frames {
//           global_buffer[frame] = global_buffer[frame] + signal[frame]
//       }
//       cs.spinUnlock(&lock)
func (csound CSOUND) SpinLock(spinlock *int32) {
	C.csSpinLock((*C.int32)(spinlock))
}

// Unlock the specified spinlock; (see SpinLock()).
func (csound CSOUND) SpinUnLock(spinlock *int32) {
	C.csSpinUnLock((*C.int32)(spinlock))
}

/*
 * Miscellaneous functions
 */

// Run an external command with the arguments specified in 'args'.
// args[0] is the name of the program to execute (if not a full path
// file name, it is searched in the directories defined by the PATH
// environment variable).
//
// If 'noWait' is false, the function waits until the external program
// finishes, otherwise it returns immediately. In the first case, a
// non-negative return value is the exit status of the command (0 to
// 255), otherwise it is the PID of the newly created process.
// On error, a negative value is returned.
func (csound CSOUND) RunCommand(args []string, noWait bool) int {
	argv := make([]*C.char, len(args)+1)
	for i, arg := range args {
		argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(argv[i]))
	}
	return int(C.csoundRunCommand(&argv[0], cbool(noWait)))
}

// Initialise a timer structure.
func (csound CSOUND) InitTimerStruct() C.RTCLOCK {
	var rtc C.RTCLOCK
	C.csoundInitTimerStruct(&rtc)
	return rtc
}

// Return the elapsed real time (in seconds) since the specified timer
// structure was initialised.
func (csound CSOUND) RealTime(rtc *C.RTCLOCK) float64 {
	return float64(C.csoundGetRealTime(rtc))
}

// Return the elapsed CPU time (in seconds) since the specified timer
// structure was initialised.
func (csound CSOUND) CPUTime(rtc *C.RTCLOCK) float64 {
	return float64(C.csoundGetCPUTime(rtc))
}

// Return a 32-bit unsigned integer to be used as seed from current time.
func (csound CSOUND) RandomSeedFromTime() uint32 {
	return uint32(C.csoundGetRandomSeedFromTime())
}

type Cslanguage_t int

// List of languages
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

// Set language to 'langCode' (langCode can be for example
// CSLANGUAGE_ENGLISH_UK or CSLANGUAGE_FRENCH or many others,
// see n_getstr.h for the list of languages). This affects all
// Csound instances running in the address space of the current
// process. The special language code CSLANGUAGE_DEFAULT can be
// used to disable translation of messages and free all memory
// allocated by a previous call to SetLanguage().
// SetLanguage() loads all files for the selected language
// from the directory specified by the CSSTRNGS environment
// variable.
func (csound CSOUND) SetLanguage(langCode Cslanguage_t) {
	C.csoundSetLanguage(C.cslanguage_t(langCode))
}

// Get the value of environment variable 'name', searching
// in this order: local environment of 'csound', variables
// set with SetGlobalEnv(), and system environment variables.
// Should be called after PreCompile() or Compile().
// Return value is nil if the variable is not set.
func (csound CSOUND) Env(name string) string {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(C.csoundGetEnv(csound.Cs, cname))
}

// Set the global value of environment variable 'name' to 'value',
// or delete variable if 'value' is nil.
// It is not safe to call this function while any Csound instances
// are active.
// Return zero on success.
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

// Allocate nbytes bytes of memory that can be accessed later by calling
// QueryGlobalVariable() with the specified name; the space is
// cleared to zero.
//
// Return CSOUND_SUCCESS on success, CSOUND_ERROR in case of invalid
// parameters (zero nbytes, invalid or already used name), or
// CSOUND_MEMORY if there is not enough memory.
func (csound CSOUND) CreateGlobalVariable(name string, nbytes uint) int {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundCreateGlobalVariable(csound.Cs, cname, C.size_t(nbytes)))
}

// Get pointer to space allocated with the name "name".
// Returns nil if the specified name is not defined.
func (csound CSOUND) QueryGlobalVariable(name string) unsafe.Pointer {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return (unsafe.Pointer)(C.csoundQueryGlobalVariable(csound.Cs, cname))
}

// This function is the same as QueryGlobalVariable(), except the
// variable is assumed to exist and no error checking is done.
// Faster, but may crash or return an invalid pointer if 'name' is
// not defined.
func (csound CSOUND) QueryGlobalVariableNoCheck(name string) unsafe.Pointer {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return (unsafe.Pointer)(C.csoundQueryGlobalVariableNoCheck(csound.Cs, cname))
}

// Free memory allocated for "name" and remove "name" from the database.
// Return value is CSOUND_SUCCESS on success, or CSOUND_ERROR if the name is
// not defined.
func (csound CSOUND) DestroyGlobalVariable(name string) int {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.csoundDestroyGlobalVariable(csound.Cs, cname))
}

// Run utility with the specified name and command line arguments.
// Should be called after loading utility plugins with PreCompile();
// use Reset() to clean up after calling this function.
// Returns zero if the utility was run successfully.
func (csound CSOUND) RunUtility(name string, args []string) int {
	var cname *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	argc := C.int(len(args))
	argv := make([]*C.char, argc)
	for i, arg := range args {
		argv[i] = C.CString(arg)
	}
	result := C.csoundRunUtility(csound.Cs, cname, argc, &argv[0])
	for _, arg := range argv {
		C.free(unsafe.Pointer(arg))
	}
	return int(result)
}

// Return a list of registered utility names.
// The return value may be nil in case of an error.
func (csound CSOUND) ListUtilities() ([]string, error) {
	clist := C.csoundListUtilities(csound.Cs)
	if clist == nil {
		return nil, fmt.Errorf("ListUtilities error")
	}
	n := int(C.utilityListLength(clist))
	var list = make([]string, n)
	for i := range list {
		list[i] = C.GoString(C.utilityName(clist, C.int(i)))
	}
	C.csoundDeleteUtilityList(csound.Cs, clist)
	return list, nil
}

// Get utility description.
// Return nil if the utility was not found, or it has no description,
// or an error occured.
func (csound CSOUND) UtilityDescription(utilName string) string {
	var cname *C.char = C.CString(utilName)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(C.csoundGetUtilityDescription(csound.Cs, cname))
}

// Simple linear congruential random number generator:
//   (*seedVal) = (*seedVal) * 742938285 % 2147483647
// The initial value of *seedVal must be in the range 1 to 2147483646.
// Return the next number from the pseudo-random sequence,
// in the range 1 to 2147483646.
func (csound CSOUND) Rand31(seedVal *int32) int32 {
	return int32(C.csoundRand31((*C.int)(seedVal)))
}

// Initialise Mersenne Twister (MT19937) random number generator,
// using len(initKey) unsigned 32 bit values from 'initKey' as seed.
// One has to free the memory used to store the PNRG state, when the PNRG
// is not needed anymore (see FreeRandMTState())
func (csound CSOUND) SeedRandMT(initKey []uint32) *C.CsoundRandMTState {
	p := C.newRandMTState()
	if len(initKey) > 1 {
		C.csoundSeedRandMT(p, (*C.uint32_t)(&initKey[0]), C.uint32_t(len(initKey)))
	} else {
		C.csoundSeedRandMT(p, nil, C.uint32_t(initKey[0]))
	}
	return p
}

// Return next random number from MT19937 generator.
// The PRNG must be initialised first by calling SeedRandMT().
func (csound CSOUND) RandMT(p *C.CsoundRandMTState) uint32 {
	return uint32(C.csoundRandMT(p))
}

// Free the memory pointed to by the C.CsoundRandMTState pointer.
func (csound CSOUND) FreeRandMTState(p *C.CsoundRandMTState) {
	C.freeRandMTState(p)
}

// Create circular buffer with numelem number of MYFLT elements.
// It should be used like:
//   rb := csound.CreateCircularBuffer(1024)
func (csound CSOUND) CreateCircularBuffer(numelem int) unsafe.Pointer {
	var sample MYFLT
	return unsafe.Pointer(C.csoundCreateCircularBuffer(csound.Cs, C.int(numelem),
		C.int(unsafe.Sizeof(sample))))
}

// Read from circular buffer
//   circular_buffer - pointer to an existing circular buffer
//   out - preallocated buffer with at least items number of elements, where
//         buffer contents will be read into
//   items - number of samples to be read
// Return the actual number of items read (0 <= n <= items)
func (csound CSOUND) ReadCircularBuffer(circularBuffer unsafe.Pointer, out []MYFLT,
	items int) int {
	if len(out) < items {
		return 0
	}
	return int(C.csoundReadCircularBuffer(csound.Cs, circularBuffer,
		unsafe.Pointer(&out[0]), C.int(items)))
}

// Read from circular buffer without removing them from the buffer.
//   circular_buffer - pointer to an existing circular buffer
//   out - preallocated buffer with at least items number of elements, where
//         buffer contents will be read into
// items - number of samples to be read
// Return the actual number of items read (0 <= n <= items)
func (csound CSOUND) PeekCircularBuffer(circularBuffer unsafe.Pointer, out []MYFLT,
	items int) int {
	if len(out) < items {
		return 0
	}
	return int(C.csoundPeekCircularBuffer(csound.Cs, circularBuffer,
		unsafe.Pointer(&out[0]), C.int(items)))
}

// Write to circular buffer
//   circular_buffer - pointer to an existing circular buffer
//   inp - buffer with at least items number of elements to be written into
//         circular buffer
//   items - number of samples to be written
// Returns the actual number of samples written (0 <= n <= items)
func (csound CSOUND) WriteCircularBuffer(circularBuffer unsafe.Pointer, inp []MYFLT,
	items int) int {
	if len(inp) < items {
		return 0
	}
	return int(C.csoundWriteCircularBuffer(csound.Cs, circularBuffer,
		unsafe.Pointer(&inp[0]), C.int(items)))
}

// Empty circular buffer of any remaining data.
// This function shuould only be used if there is no reader actively
// getting data from the buffer.
//   circular_buffer - pointer to an existing circular buffer
func (csound CSOUND) FlushCircularBuffer(circularBuffer unsafe.Pointer) {
	C.csoundFlushCircularBuffer(csound.Cs, circularBuffer)
}

// Free circular buffer.
func (csound CSOUND) DestroyCircularBuffer(circularBuffer unsafe.Pointer) {
	C.csoundDestroyCircularBuffer(csound.Cs, circularBuffer)
}

// Platform-independent function to load a shared library.
func (csound CSOUND) OpenLibrary(libraryPath string) (int, unsafe.Pointer) {
	var cpath *C.char = C.CString(libraryPath)
	defer C.free(unsafe.Pointer(cpath))
	var library unsafe.Pointer
	ret := C.csoundOpenLibrary((*unsafe.Pointer)(&library), cpath)
	return int(ret), library
}

// Platform-independent function to unload a shared library.
func (csound CSOUND) CloseLibrary(library unsafe.Pointer) int {
	return int(C.csoundCloseLibrary(library))
}

// Platform-independent function to get a symbol address in a shared library.
func (csound CSOUND) LibrarySymbol(library unsafe.Pointer,
	symbolName string) unsafe.Pointer {
	var cname *C.char = C.CString(symbolName)
	defer C.free(unsafe.Pointer(cname))
	return C.csoundGetLibrarySymbol(library, cname)
}
