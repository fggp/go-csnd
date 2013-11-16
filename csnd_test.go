package csnd6

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestInstantiation(t *testing.T) {
	cs := Create(nil)
	if cs.cs == nil {
		t.Errorf("Could not create Csound instance")
	}
	fmt.Println("\n", cs.GetVersion(), cs.GetAPIVersion(), "\n")
	cs.Destroy()
	if cs.cs != nil {
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
	ht := cs.GetHostData()
	if ht != nil {
		t.Errorf("Hostdata should be nil when instance created with nil arg")
	}

	var i interface{} = 1956
	cs.SetHostData(&i)
	ht = cs.GetHostData()
	if ht != &i {
		t.Errorf("Int hostdata read is different of hostdata written")
	} else {
		fmt.Println("\n", *ht, "\n")
	}

	var s interface{} = "Une chaîne de caractères"
	cs.SetHostData(&s)
	ht = cs.GetHostData()
	if ht != &s {
		t.Errorf("String hostdata read is different of hostdata written")
	} else {
		fmt.Println("\n", *ht, "\n")
	}

	cs.SetHostData(nil)
	ht = cs.GetHostData()
	if ht != nil {
		t.Errorf("Hostdata should have been cleared")
	}
	cs.Destroy()

	triangle := make(Triangle, 3)
	triangle[&Point{1, 2, 3}] = "α"
	triangle[&Point{4, 5, 6}] = "β"
	triangle[&Point{7, 8, 9}] = "γ"
	cs = Create((*interface{})(unsafe.Pointer(&triangle)))
	ht = cs.GetHostData()
	pt := (*Triangle)(unsafe.Pointer(ht))
	if pt != &triangle {
		t.Errorf("String hostdata read is different of hostdata written")
	} else {
		fmt.Println("\n", *pt, "\n")
	}

	cs.SetHostData(nil)
	ht = cs.GetHostData()
	if ht != nil {
		t.Errorf("Hostdata should have been cleared")
	}
	cs.Destroy()
}

func TestCsoundParams(t *testing.T) {
	cs := Create(nil)
	var p CsoundParams
	fmt.Println(p)
	cs.GetParams(&p)
	fmt.Println(p)
	p.ring_bell = 1
	cs.SetParams(&p)
	p.ring_bell = 0
	fmt.Println(p)
	cs.GetParams(&p)
	fmt.Println(p)
	cs.SetDebug(true)
	cs.GetParams(&p)
	fmt.Println(p)
	fmt.Println(cs.GetDebug())
	p.ring_bell = 0
	p.debug_mode = 0
	cs.SetParams(&p)
	fmt.Println(cs.GetDebug())
	cs.Destroy()
}

func TestRTAudioIO(t *testing.T) {
	cs := Create(nil)
	var n int
	for {
		if name, mtype, err := cs.GetModule(n); err == CSOUND_SUCCESS {
			fmt.Printf("%2d: %s\t%s\n", n, name, mtype)
			n++
		} else {
			break
		}
	}

	cs.Compile([]string{"dummy", "simple.csd"})
	list := cs.GetAudioDevList(true)
	fmt.Println("\nGetAudioDevList(true)")
	for i := range list {
		fmt.Println(list[i])
	}
	list = cs.GetAudioDevList(false)
	fmt.Println("\nGetAudioDevList(false)")
	for i := range list {
		fmt.Println(list[i])
	}
	fmt.Println()

	cs.Destroy()
}

func TestMidiIO(t *testing.T) {
	cs := Create(nil)
	cs.Compile([]string{"dummy", "simple.csd"})
	list := cs.GetMidiDevList(true)
	fmt.Println()
	for i := range list {
		fmt.Println(list[i])
	}
	list = cs.GetMidiDevList(false)
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
	cs.GetChannelPtr("Zobie", CSOUND_CONTROL_CHANNEL)
	lst, err := cs.ListChannels()
	if err != nil {
		fmt.Println(err)
	} else if lst == nil {
		fmt.Println("Channel list is empty")
	} else {
		fmt.Println(len(lst))
	}
}
