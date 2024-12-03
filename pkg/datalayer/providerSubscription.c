/**
 * MIT License
 *
 * Copyright (c) 2021-2024 Bosch Rexroth AG
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
#include "_cgo_export.h"

#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

#include <variant.h>

#ifdef __cplusplus
extern "C" {
#endif

// Wrapper of a list of DLR_NOTIFY_ITEM
typedef struct VEC_NOTIFY_ITEM
{
    size_t count;              //! count of requests
    DLR_NOTIFY_ITEM* items;    //! array of DLR_NOTIFY_ITEM
} VEC_NOTIFY_ITEM;

// CreateProviderSubscriptionsNotifyItems create a wrapper of a list of DLR_NOTIFY_ITEM
unsigned long long CreateProviderSubscriptionsNotifyItems(size_t size)
{
    VEC_NOTIFY_ITEM *p = malloc(sizeof(VEC_NOTIFY_ITEM));
    p->items = malloc(sizeof(DLR_NOTIFY_ITEM) * size);
    p->count = size;
    return (unsigned long long)p;
}

// DeleteProviderSubscriptionsNotifyItems remove wrapper
void DeleteProviderSubscriptionsNotifyItems(unsigned long long p)
{
    VEC_NOTIFY_ITEM *ptr = (VEC_NOTIFY_ITEM *)p;
    free(ptr->items);
    free(ptr);
}

// SetProviderSubscriptionsNotifyItem set the DLR_NOTIFY_ITEM
void SetProviderSubscriptionsNotifyItem(unsigned long long p, uint32_t index, DLR_VARIANT data, DLR_VARIANT info)
{
    VEC_NOTIFY_ITEM *ptr = (VEC_NOTIFY_ITEM *)p;
    if (index >= ptr->count)
    {
        return;
    }
    ptr->items[index].data = data;
    ptr->items[index].info = info;
}

// GetProviderSubscriptionsNotifyItems gets the list of DLR_NOTIFY_ITEM
DLR_NOTIFY_ITEM* GetProviderSubscriptionsNotifyItems(unsigned long long p)
{
    VEC_NOTIFY_ITEM *ptr = (VEC_NOTIFY_ITEM *)p;
    return ptr->items;
}

#ifdef __cplusplus
}
#endif
