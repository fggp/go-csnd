#include <csound/csound.hpp>
#include <csound/csPerfThread.hpp>
#include "csound_pt.h"

extern "C" {
  
Cpt NewCsoundPT(CSOUND *csound)
{
  CsoundPerformanceThread *pt = new CsoundPerformanceThread(csound);
  return (void *)pt;
}

void DeleteCsoundPT(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  delete cpt;
}

int CsoundPTisRunning(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  return cpt->isRunning();
}

CSOUND *CsoundPTgetCsound(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  return cpt->GetCsound();
}

int CsoundPTgetStatus(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  return cpt->GetStatus();
}

void CsoundPTplay(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->Play();
}

void CsoundPTpause(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->Pause();
}

void CsoundPTtogglePause(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->TogglePause();
}

void CsoundPTstop(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->Stop();
}

void CsoundPTscoreEvent(Cpt pt, int absp2mode, char opcod, int pcnt, MYFLT *p)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->ScoreEvent(absp2mode, opcod, pcnt, p);
}

void CsoundPTinputMessage(Cpt pt, const char *s)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->InputMessage(s);
}

void CsoundPTsetScoreOffsetSeconds(Cpt pt, double timeVal)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->SetScoreOffsetSeconds(timeVal);
}

void CsoundPTjoin(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->Join();
}

void CsoundPTflushMessageQueue(Cpt pt)
{
  CsoundPerformanceThread *cpt = (CsoundPerformanceThread *)pt;
  cpt->FlushMessageQueue();
}

} // extern "C"
