#include "_cgo_export.h"

#include <stdio.h>
#include <stdbool.h>

#include "variant.h"

extern void responseCallbackGo(DLR_RESULT result, DLR_VARIANT data, void* userdata);

void responseCallbackC(DLR_RESULT result, DLR_VARIANT data, void* userdata){
	responseCallbackGo(result, data, userdata);
}
