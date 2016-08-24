package csnd6

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestInstantiation(t *testing.T) {
	cs := Create(nil)
	if cs.Cs == nil {
		t.Errorf("Could not create Csound instance")
	}
	fmt.Println("\n", cs.Version(), cs.APIVersion(), "\n")
	cs.Destroy()
	if cs.Cs != nil {
		t.Errorf("Csound was destroyed and opaque pointer is not cleared!")
	}
}

type Point struct{ x, y, z int }

func (point Point) String() string {
	return fmt.Sprintf("(%d, %d, %d)", point.x, point.y, point.z)
}

type Triangle map[*Point]string

func TestHostData(t *testing.T) {
	cs := Create(nil)
	ht := cs.HostData()
	if ht != nil {
		t.Errorf("Hostdata should be nil when instance created with nil arg")
	}

	i := 1956
	cs.SetHostData(unsafe.Pointer(&i))
	ht = cs.HostData()
	pi := (*int)(ht)
	if pi != &i {
		t.Errorf("Int hostdata read is different of hostdata written")
	} else {
		fmt.Println("\n", *pi, "\n")
	}

	s := "Une chaîne de caractères"
	cs.SetHostData(unsafe.Pointer(&s))
	ht = cs.HostData()
	ps := (*string)(ht)
	if ps != &s {
		t.Errorf("String hostdata read is different of hostdata written")
	} else {
		fmt.Println("\n", *ps, "\n")
	}

	cs.SetHostData(nil)
	ht = cs.HostData()
	if ht != nil {
		t.Errorf("Hostdata should have been cleared")
	}
	cs.Destroy()

	triangle := make(Triangle, 3)
	triangle[&Point{1, 2, 3}] = "α"
	triangle[&Point{4, 5, 6}] = "β"
	triangle[&Point{7, 8, 9}] = "γ"
	s1 := "togodo tsoin tsoin"
	cs = Create(unsafe.Pointer(&s1))
	ht = cs.HostData()
	pt := (*string)(unsafe.Pointer(ht))
	if pt != &s1 {
		t.Errorf("String hostdata read is different of hostdata written")
	} else {
		fmt.Println("\n", *pt, "\n")
	}

	cs.SetHostData(nil)
	ht = cs.HostData()
	if ht != nil {
		t.Errorf("Hostdata should have been cleared")
	}
	cs.Destroy()
}

func TestCsoundParams(t *testing.T) {
	cs := Create(nil)
	var p CsoundParams
	fmt.Println(p)
	cs.Params(&p)
	fmt.Println(p)
	p.RingBell = 1
	cs.SetParams(&p)
	p.RingBell = 0
	fmt.Println(p)
	cs.Params(&p)
	fmt.Println(p)
	cs.SetDebug(true)
	cs.Params(&p)
	fmt.Println(p)
	fmt.Println(cs.Debug())
	p.RingBell = 0
	p.DebugMode = 0
	cs.SetParams(&p)
	fmt.Println(cs.Debug())
	cs.Destroy()
}

func TestRTAudioIO(t *testing.T) {
	cs := Create(nil)
	var name, mtype string
	err := CSOUND_SUCCESS
	n := 0
	for err != CSOUND_ERROR {
		name, mtype, err = cs.Module(n)
		n++
		fmt.Printf("Module %d:  %s (%s)\n", n, name, mtype)
	}

	cs.Compile([]string{"dummy", "simple.csd"})
	list := cs.AudioDevList(true)
	fmt.Println("\nAudioDevList(true)")
	for i := range list {
		fmt.Printf("%d: %s (%s), %d chan\n",
			i, list[i].DeviceId, list[i].DeviceName, list[i].MaxNchnls)
	}
	list = cs.AudioDevList(false)
	fmt.Println("\nAudioDevList(false)")
	for i := range list {
		fmt.Printf("%d: %s (%s), %d chan\n",
			i, list[i].DeviceId, list[i].DeviceName, list[i].MaxNchnls)
	}
	fmt.Println()
	cs.Destroy()
}

func TestMidiIO(t *testing.T) {
	cs := Create(nil)
	cs.Compile([]string{"dummy", "simple.csd"})
	list := cs.MidiDevList(true)
	fmt.Println()
	for i := range list {
		fmt.Println(list[i])
	}
	list = cs.MidiDevList(false)
	fmt.Println()
	for i := range list {
		fmt.Println(list[i])
	}
	fmt.Println()
	cs.Destroy()
}

func TestChannels(t *testing.T) {
	cs := Create(nil)
	cs.Compile([]string{"dummy", "simple.csd"})
	cs.Start()
	cs.ChannelPtr("Zobie", CSOUND_CONTROL_CHANNEL)
	lst, err := cs.ListChannels()
	if err != nil {
		fmt.Println(err)
	} else if lst == nil {
		fmt.Println("Channel list is empty")
	} else {
		fmt.Println(len(lst))
	}
}

func TestOpcodeList(t *testing.T) {
	cs := Create(nil)
	if list := cs.OpcodeList(); list != nil {
		fmt.Println(list)
		fmt.Println(len(list))
	}
	cs.Destroy()
}

func TestNamedGens(t *testing.T) {
	cs := Create(nil)
	namedGens := cs.NamedGens()
	fmt.Println(namedGens)
	cs.Destroy()
}

func TestRunCommand(t *testing.T) {
	cs := Create(nil)
	cs.RunCommand([]string{"ls", "-a"}, false)
	cs.Destroy()
}

func TestUtilities(t *testing.T) {
	cs := Create(nil)
	if list, err := cs.ListUtilities(); err == nil {
		for _, name := range list {
			fmt.Printf("%s: %s\n", name, cs.UtilityDescription(name))
		}
	}
	cs.Destroy()
}

func TestRand31(t *testing.T) {
	cs := Create(nil)
	seed := int32(1956)
	for i := 0; i < 1000; i++ {
		n := cs.Rand31(&seed)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
	cs.Destroy()
}

func TestRandMT(t *testing.T) {
	cs := Create(nil)
	p := cs.SeedRandMT([]uint32{1956})
	for i := 0; i < 1000; i++ {
		n := cs.RandMT(p)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
	cs.Destroy()
}

func TestMessages(t *testing.T) {
	cs := Create(nil)
	version := cs.Version()
	apiVersion := cs.APIVersion()
	sr := cs.Sr()
	ksmps := cs.Ksmps()
	cs.Message("\n%s, %s, %f, %d\n\n", version, apiVersion, sr, ksmps)
	cs.MessageS(CSOUNDMSG_FG_RED|CSOUNDMSG_FG_UNDERLINE, "\n%s, %s, %f, %d\n\n",
		version, apiVersion, sr, ksmps)
	cs.Destroy()
}
