# golang-ui
Sandbox for Golang UI

## Enabling CGO on Windows

1. Install [MSYS2](https://www.msys2.org/) 64 bit (msys2-x86_64)
2. Update (See [here](https://github.com/msys2/msys2/wiki/MSYS2-installation))
  - `pacman -Syu`
  - `pacman -Syuu`
  - `pacman -Su`
  - many times
3. Install `gcc` (See [here](https://stackoverflow.com/questions/51724007/msys2-install-gcc-or-toolchain))
  - `pacman -S mingw-w64-x86_64-gcc`
  - `pacman -S mingw-w64-i686-gcc` for 32 bit version
4. Add to PATH
  - C:\msys64\mingw64\bin
  - C:\msys64\usr\bin
5. ...
6. PROFIT!

```go
package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import "unsafe"

func main() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
```
### Caveat

Cross-compiling
(eg. building 32-bit .exe on 64-bit Windows with CGO)
is not working.
