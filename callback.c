#include "_cgo_export.h"

extern void goFileOpenCB(void *, char *, int, int, int);

void csoundFileOpenCB(CSOUND *csound, const char *pathName, int fileType,
                      int write, int temp)
{
  goFileOpenCB((void *)csound, (char *)pathName, fileType, write, temp);
}

void csoundSetFileOpenCB(CSOUND *csound)
{
  csoundSetFileOpenCallback(csound, csoundFileOpenCB);
}

/*////////////////////////////////////////////////////////////*/

extern int goPlayOpenCB(void *, csRtAudioParams *);

int csoundPlayOpenCB(CSOUND *csound, const csRtAudioParams *parm)
{
  return goPlayOpenCB((void *)csound, (csRtAudioParams *)parm);
}

void csoundSetPlayOpenCB(CSOUND *csound)
{
  csoundSetPlayopenCallback(csound, csoundPlayOpenCB);
}

/*////////////////////////////////////////////////////////////*/

extern void goRtPlayCB(void *, MYFLT *, int);

void csoundRtPlayCB(CSOUND *csound, const MYFLT *outbuf, int nbytes)
{
  goRtPlayCB((void *)csound, (MYFLT *)outbuf, nbytes/sizeof(MYFLT));
}

void csoundSetRtPlayCB(CSOUND *csound)
{
  csoundSetRtplayCallback(csound, csoundRtPlayCB);
}

/*////////////////////////////////////////////////////////////*/

extern int goRecOpenCB(void *, csRtAudioParams *);

int csoundRecOpenCB(CSOUND *csound, const csRtAudioParams *parm)
{
  return goRecOpenCB((void *)csound, (csRtAudioParams *)parm);
}

void csoundSetRecOpenCB(CSOUND *csound)
{
  csoundSetRecopenCallback(csound, csoundPlayOpenCB);
}

/*////////////////////////////////////////////////////////////*/

extern int goRtRecordCB(void *, MYFLT *, int);

int csoundRtRecordCB(CSOUND *csound, MYFLT *inbuf, int nbytes)
{
  return goRtRecordCB((void *)csound, inbuf, nbytes/sizeof(MYFLT));
}

void csoundSetRtRecordCB(CSOUND *csound)
{
  csoundSetRtrecordCallback(csound, csoundRtRecordCB);
}

/*////////////////////////////////////////////////////////////*/

extern void goRtCloseCB(void *);

void csoundRtCloseCB(CSOUND *csound)
{
  goRtCloseCB((void *)csound);
}

void csoundSetRtCloseCB(CSOUND *csound)
{
  csoundSetRtcloseCallback(csound, csoundRtCloseCB);
}

/*////////////////////////////////////////////////////////////*/

extern void goCscore(void *);

void csoundCscoreCB(CSOUND *csound)
{
  goCscoreCB((void *)csound);
}

void csoundSetCscoreCB(CSOUND *csound)
{
  csoundSetCscoreCallback(csound, csoundCscoreCB);
}

/*////////////////////////////////////////////////////////////*/


