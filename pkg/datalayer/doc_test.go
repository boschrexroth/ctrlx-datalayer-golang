/**
 * MIT License
 *
 * Copyright (c) 2021-2022 Bosch Rexroth AG
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
package datalayer_test

import (
	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	fbs "github.com/boschrexroth/ctrlx-datalayer-golang/pkg/fbs/comm/datalayer"
)

func ExampleMetaDataBuilder() {
	m := datalayer.NewMetaDataBuilder(datalayer.AllowedOperationRead|datalayer.AllowedOperationWrite, "simple description", "")
	m.DisplayName("myexample").NodeClass(fbs.NodeClassVariable)

	m.AddReference(datalayer.ReferenceTypeWrite, "type/myexample/write")
	m.AddReference(datalayer.ReferenceTypeRead, "type/myexample/read")

	m.AddExtension("key", "value")

	v := m.Build()
	defer datalayer.DeleteVariant(v)
}
