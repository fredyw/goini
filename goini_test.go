// The MIT License (MIT)
//
// Copyright (c) 2016 Fredy Wijaya
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package goini

import (
	"reflect"
	"runtime/debug"
	"testing"
)

// assertEquals asserts if the two objects are the same.
func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		debug.PrintStack()
		t.Errorf("Expected: %s, got: %s\n", expected, actual)
	}
}

// assertNil asserts if the given object is nil.
func assertNil(t *testing.T, actual interface{}) {
	val := reflect.ValueOf(actual)
	if !val.IsNil() {
		debug.PrintStack()
		t.Errorf("Expected: nil, got: %s\n", actual)
	}
}

// assertTrue asserts if the given bool is true
func assertTrue(t *testing.T, b bool) {
	if !b {
		debug.PrintStack()
		t.Errorf("Expected: true, got: false")
	}
}

// assertFalse asserts if the given bool is false
func assertFalse(t *testing.T, b bool) {
	if b {
		debug.PrintStack()
		t.Errorf("Expected: false, got: true")
	}
}

// assertNotNil asserts if the given object is not nil.
func assertNotNil(t *testing.T, actual interface{}) {
	val := reflect.ValueOf(actual)
	if val.IsNil() {
		debug.PrintStack()
		t.Errorf("Expected: not nil, got: %s\n", actual)
	}
}

// assertError asserts if error is set.
func assertError(t *testing.T, err error) {
	if err == nil {
		debug.PrintStack()
		t.Fatal("Expected: error")
	}
}

// assertNoError asserts if error is not set.
func assertNoError(t *testing.T, err error) {
	if err != nil {
		debug.PrintStack()
		t.Fatalf("Expected: no error, got: %s instead", err)
	}
}

// fail fails the test.
func fail(t *testing.T, message ...interface{}) {
	debug.PrintStack()
	t.Fatal(message)
}

func TestINIOrdered(t *testing.T) {
	ini := NewINI(true)

	added := ini.AddOption("section1", "option1", "value1")
	assertTrue(t, added)

	added = ini.AddOption("section1", "option2", "value2")
	assertTrue(t, added)

	added = ini.AddOption("section1", "option2", "value2_modified")
	assertTrue(t, added)

	added = ini.AddOption("section2", "option1", "value1")
	assertTrue(t, added)

	added = ini.AddOption("section2", "option2", "value2")
	assertTrue(t, added)

	added = ini.AddOption("section2", "option2", "value2_modified")
	assertTrue(t, added)

	added = ini.AddSection("section3")
	assertTrue(t, added)

	added = ini.AddSection("section3")
	assertFalse(t, added)

	sections := ini.Sections()
	assertEquals(t, 3, len(sections))
	assertEquals(t, "section1", sections[0])
	assertEquals(t, "section2", sections[1])
	assertEquals(t, "section3", sections[2])

	options := ini.Options("section1")
	assertEquals(t, 2, len(options))
	assertEquals(t, "option1", options[0])
	assertEquals(t, "option2", options[1])

	options = ini.Options("section2")
	assertEquals(t, 2, len(options))
	assertEquals(t, "option1", options[0])
	assertEquals(t, "option2", options[1])

	options = ini.Options("section3")
	assertEquals(t, 0, len(options))

	val, found := ini.GetOption("section1", "option1")
	assertEquals(t, "value1", val)
	assertTrue(t, found)

	val, found = ini.GetOption("section1", "option2")
	assertEquals(t, "value2_modified", val)
	assertTrue(t, found)

	val, found = ini.GetOption("section1", "doesnotexist")
	assertFalse(t, found)

	val, found = ini.GetOption("section2", "option1")
	assertEquals(t, "value1", val)
	assertTrue(t, found)

	val, found = ini.GetOption("section2", "option2")
	assertEquals(t, "value2_modified", val)
	assertTrue(t, found)

	val, found = ini.GetOption("section2", "doesnotexist")
	assertFalse(t, found)

	val, found = ini.GetOption("section3", "doesnotexist")
	assertFalse(t, found)

	added = ini.AddOption("section3", "option1", "value1")
	assertTrue(t, added)

	options = ini.Options("section3")
	assertEquals(t, 1, len(options))
	assertEquals(t, "option1", options[0])

	val, found = ini.GetOption("section3", "option1")
	assertEquals(t, "value1", val)
	assertTrue(t, found)

	removed := ini.RemoveSection("section3")
	assertTrue(t, removed)

	removed = ini.RemoveSection("doesntexist")
	assertFalse(t, removed)

	sections = ini.Sections()
	assertEquals(t, 2, len(sections))
	assertEquals(t, "section1", sections[0])
	assertEquals(t, "section2", sections[1])

	removed = ini.RemoveOption("section1", "option1")
	assertTrue(t, removed)

	options = ini.Options("section1")
	assertEquals(t, 1, len(options))
	assertEquals(t, "option2", options[0])

	val, found = ini.GetOption("section1", "option2")
	assertEquals(t, "value2_modified", val)
	assertTrue(t, found)
}

func TestINIUnordered(t *testing.T) {
	ini := NewINI(false)

	added := ini.AddOption("section1", "option1", "value1")
	assertTrue(t, added)

	added = ini.AddOption("section1", "option2", "value2")
	assertTrue(t, added)

	added = ini.AddOption("section1", "option2", "value2_modified")
	assertTrue(t, added)

	added = ini.AddOption("section2", "option1", "value1")
	assertTrue(t, added)

	added = ini.AddOption("section2", "option2", "value2")
	assertTrue(t, added)

	added = ini.AddOption("section2", "option2", "value2_modified")
	assertTrue(t, added)

	added = ini.AddSection("section3")
	assertTrue(t, added)

	added = ini.AddSection("section3")
	assertFalse(t, added)

	sections := ini.Sections()
	assertEquals(t, 3, len(sections))

	options := ini.Options("section1")
	assertEquals(t, 2, len(options))

	options = ini.Options("section2")
	assertEquals(t, 2, len(options))

	options = ini.Options("section3")
	assertEquals(t, 0, len(options))

	val, found := ini.GetOption("section1", "option1")
	assertEquals(t, "value1", val)
	assertTrue(t, found)

	val, found = ini.GetOption("section1", "option2")
	assertEquals(t, "value2_modified", val)
	assertTrue(t, found)

	val, found = ini.GetOption("section1", "doesnotexist")
	assertFalse(t, found)

	val, found = ini.GetOption("section2", "option1")
	assertEquals(t, "value1", val)
	assertTrue(t, found)

	val, found = ini.GetOption("section2", "option2")
	assertEquals(t, "value2_modified", val)
	assertTrue(t, found)

	val, found = ini.GetOption("section2", "doesnotexist")
	assertFalse(t, found)

	val, found = ini.GetOption("section3", "doesnotexist")
	assertFalse(t, found)

	added = ini.AddOption("section3", "option1", "value1")
	assertTrue(t, added)

	options = ini.Options("section3")
	assertEquals(t, 1, len(options))

	val, found = ini.GetOption("section3", "option1")
	assertEquals(t, "value1", val)
	assertTrue(t, found)

	removed := ini.RemoveSection("section3")
	assertTrue(t, removed)

	removed = ini.RemoveSection("doesntexist")
	assertFalse(t, removed)

	sections = ini.Sections()
	assertEquals(t, 2, len(sections))

	removed = ini.RemoveOption("section1", "option1")
	assertTrue(t, removed)

	options = ini.Options("section1")
	assertEquals(t, 1, len(options))

	val, found = ini.GetOption("section1", "option2")
	assertEquals(t, "value2_modified", val)
	assertTrue(t, found)
}

func TestReadWriteOrdered(t *testing.T) {
	// TODO
}

func TestReadWriteUnordered(t *testing.T) {
	// TODO
}
