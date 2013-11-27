#ifndef CSOUND_PT_H
#define CSOUND_PT_H

#ifdef __cplusplus
extern "C" {
#endif

typedef void* Cpt;

Cpt NewCsoundPT(CSOUND *);
void DeleteCsoundPT(Cpt pt);
int CsoundPTisRunning(Cpt pt);
void CsoundPTsetProcessCB(Cpt pt, void *cbData);
CSOUND *CsoundPTgetCsound(Cpt pt);
int CsoundPTgetStatus(Cpt pt);
void CsoundPTplay(Cpt pt);
void CsoundPTpause(Cpt pt);
void CsoundPTtogglePause(Cpt pt);
void CsoundPTstop(Cpt pt);
void CsoundPTscoreEvent(Cpt pt, int absp2mode, char opcod, int pcnt, MYFLT *p);
void CsoundPTinputMessage(Cpt pt, const char *s);
void CsoundPTsetScoreOffsetSeconds(Cpt pt, double timeVal);
void CsoundPTjoin(Cpt pt);
void CsoundPTflushMessageQueue(Cpt pt);

#ifdef __cplusplus
} // extern "C"
#endif

#endif // CSOUND_PT_H

