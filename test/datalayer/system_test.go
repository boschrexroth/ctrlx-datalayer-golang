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
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSystem(t *testing.T) {
	assert.NotPanics(t, func() { datalayer.DeleteSystem(nil) })
}

func TestSystemNew(t *testing.T) {
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)
	assert.NotNil(t, s)
}

func TestSystemStart(t *testing.T) {
	b := []bool{true, false}
	for _, e := range b {
		s := datalayer.NewSystem("")
		defer datalayer.DeleteSystem(s)
		assert.NotNil(t, s)
		assert.NotPanics(t, func() { s.Start(e) })
	}
}

func TestSystemStop(t *testing.T) {
	b := []bool{true, false}
	for _, e := range b {
		s := datalayer.NewSystem("")
		defer datalayer.DeleteSystem(s)
		assert.NotNil(t, s)
		assert.NotPanics(t, func() { s.Stop(e) })
	}
}

func TestSystemFactory(t *testing.T) {
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)
	assert.NotNil(t, s)
	f := s.Factory()
	assert.NotNil(t, f)
}

func TestSystemConverter(t *testing.T) {
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)
	assert.NotNil(t, s)
	c := s.JSONConverter()
	assert.NotNil(t, c)
}

func TestSystemSetBfbsPath(t *testing.T) {
	s := datalayer.NewSystem("")
	defer datalayer.DeleteSystem(s)
	assert.NotNil(t, s)
	assert.NotPanics(t, func() { s.SetBfbsPath("test_path") })
}
