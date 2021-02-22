#include "_cgo_export.h"

#include <stdio.h>
#include <stdbool.h>

#include "variant.h"
#include "provider_node.h"

typedef DLR_PROVIDER_NODE_CALLBACK TYPE_CB;
typedef DLR_PROVIDER_NODE_CALLBACKDATA TYPE_CBDATA;


void callCallbackC(TYPE_CB cb, TYPE_CBDATA cbdata, DLR_RESULT result, DLR_VARIANT data){
    cb(cbdata, result, data);
}

extern void nodeCallbackGo(void* userdata, char* address, DLR_VARIANT data, TYPE_CB cb, TYPE_CBDATA cbdata, int action);

void nodeCallbackOnCreate(void* userdata, char* address, DLR_VARIANT data, TYPE_CB cb, TYPE_CBDATA cbdata){
	nodeCallbackGo(userdata, address, data, cb, cbdata, NODE_ACTION_ON_CREATE);
}

void nodeCallbackOnRemove(void* userdata, char* address, TYPE_CB cb, TYPE_CBDATA cbdata){
	nodeCallbackGo(userdata, address, NULL, cb, cbdata, NODE_ACTION_ON_REMOVE);
}

void nodeCallbackOnBrowse(void* userdata, char* address, TYPE_CB cb, TYPE_CBDATA cbdata){
	nodeCallbackGo(userdata, address, NULL, cb, cbdata, NODE_ACTION_ON_BROWSE);
}

void nodeCallbackOnRead(void* userdata, char* address, DLR_VARIANT data, TYPE_CB cb, TYPE_CBDATA cbdata){
	nodeCallbackGo(userdata, address, data, cb, cbdata, NODE_ACTION_ON_READ);
}

void nodeCallbackOnWrite(void* userdata, char* address, DLR_VARIANT data, TYPE_CB cb, TYPE_CBDATA cbdata){
	nodeCallbackGo(userdata, address, data, cb, cbdata, NODE_ACTION_ON_WRITE);
}

void nodeCallbackOnMetadata(void* userdata, char* address, TYPE_CB cb, TYPE_CBDATA cbdata){
	nodeCallbackGo(userdata, address, NULL, cb, cbdata, NODE_ACTION_ON_METADATA);
}