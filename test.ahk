LWin & f::

WinGet Style, Style, A
if(Style & 0xC40000) {
  WinSet, Style, -0xC40000, A
  ;WinMaximize, A 

} else {
  WinSet, Style, +0xC40000, A
 ; WinRestore, A
}
return