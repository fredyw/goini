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
	"io"
)

// INI is a struct that represents a parsed INI file.
type INI struct {
	ordered      bool
	sectionNames []string
	sections     map[string]*Options
}

// NewINI creates a new INI.
func NewINI(ordered bool) *INI {
	return &INI{
		ordered:      ordered,
		sectionNames: []string{},
		sections:     map[string]*Options{},
	}
}

// HasSection checks if the specified section name exists.
func (ini *INI) HasSection(sectionName string) bool {
	_, found := ini.sections[sectionName]
	return found
}

// HasOption checks if the specified section and option names exist.
func (ini *INI) HasOption(sectionName, optionName string) bool {
	if !ini.HasSection(sectionName) {
		return false
	}
	opts := ini.sections[sectionName]
	return opts.Exist(optionName)
}

// AddSection add a new section. This method returns true if the section name can be
// successfully added. It returns false if the section name already exists.
func (ini *INI) AddSection(sectionName string) bool {
	if !ini.HasSection(sectionName) {
		opts := NewOptions(ini.ordered)
		ini.sections[sectionName] = opts
		if ini.ordered {
			ini.sectionNames = append(ini.sectionNames, sectionName)
		}
		return true
	}
	return false
}

// AddOption adds a new option with specified section and option names. If a section name
// does not exist, it will be automatically created. This method returns true if the option can be
// successfully added. It returns false if the option already exists.
func (ini *INI) AddOption(sectionName, optionName, optionValue string) bool {
	if !ini.HasSection(sectionName) {
		ini.AddSection(sectionName)
	}
	opts := ini.sections[sectionName]
	return opts.Add(optionName, optionValue)
}

// GetOption gets the option value from specified section and option names. If a section
// name does not exist, this method will return false.
func (ini *INI) GetOption(sectionName, optionName string) (string, bool) {
	if opts, ok := ini.sections[sectionName]; ok {
		return opts.Get(optionName)
	}
	return "", false
}

// RemoveSection removes the specified section name. This method returns true if the
// section name can be successfully removed. It returns false if the section name does not
// exist.
func (ini *INI) RemoveSection(sectionName string) bool {
	if !ini.HasSection(sectionName) {
		return false
	}
	delete(ini.sections, sectionName)
	if ini.ordered {
		i := 0
		for idx, name := range ini.sectionNames {
			if name == sectionName {
				i = idx
			}
		}
		ini.sectionNames = append(ini.sectionNames[:i], ini.sectionNames[i+1:]...)
	}
	return true
}

// RemoveOption removes the specified the option name from the specified section name.
// This method returns true if the option name can be successfully removed. It returns
// false if the section name or option name does exist.
func (ini *INI) RemoveOption(sectionName, optionName string) bool {
	if !ini.HasSection(sectionName) {
		return false
	}
	opts := ini.sections[sectionName]
	return opts.Remove(optionName)
}

// Sections returns a list of section names.
func (ini *INI) Sections() []string {
	if !ini.ordered {
		sectionNames := []string{}
		for sectionName := range ini.sections {
			sectionNames = append(sectionNames, sectionName)
		}
		return sectionNames
	}
	return ini.sectionNames
}

// Options returns a list of option names for the specified section name.
func (ini *INI) Options(sectionName string) []string {
	if !ini.HasSection(sectionName) {
		return []string{}
	}
	return ini.sections[sectionName].Options()
}

// Options is a struct that represents INI options.
type Options struct {
	ordered     bool
	optionNames []string
	options     map[string]string
}

// NewOptions creates a new option.
func NewOptions(ordered bool) *Options {
	return &Options{
		ordered:     ordered,
		optionNames: []string{},
		options:     map[string]string{},
	}
}

// Exist checks if the specified option name exists.
func (opts *Options) Exist(optionName string) bool {
	_, found := opts.options[optionName]
	return found
}

// Add adds a new option. This method returns true if the option can be successfully added.
// It returns false if the option already exists.
func (opts *Options) Add(optionName, optionValue string) bool {
	if opts.ordered {
		if !opts.Exist(optionName) {
			opts.optionNames = append(opts.optionNames, optionName)
		}
	}
	opts.options[optionName] = optionValue
	return true
}

// Get gets the option value from the specified option name. If the specified option name
// does not exist, this method will return false.
func (opts *Options) Get(optionName string) (string, bool) {
	if !opts.Exist(optionName) {
		return "", false
	}
	return opts.options[optionName], true
}

// Remove removes the specified option name. This method returns true if the specified
// option name can be successfully removed. It returns false if the option name does not
// exist.
func (opts *Options) Remove(optionName string) bool {
	if !opts.Exist(optionName) {
		return false
	}
	delete(opts.options, optionName)
	if opts.ordered {
		i := 0
		for idx, name := range opts.optionNames {
			if name == optionName {
				i = idx
			}
		}
		opts.optionNames = append(opts.optionNames[:i], opts.optionNames[i+1:]...)
	}
	return true
}

// Options returns a list of option names.
func (opts *Options) Options() []string {
	optionNames := []string{}
	if !opts.ordered {
		for optionName := range opts.options {
			optionNames = append(optionNames, optionName)
		}
		return optionNames
	}
	return opts.optionNames
}

// Read reads an INI from an io.Reader. Passing ordered parameter true will preserve the
// order. Preserving the order will have some performance overhead.
func Read(ordered bool, reader io.Reader) (*INI, error) {
	// TODO
	return nil, nil
}

// ReadFile reads an INI from a file. Passing ordered parameter true will preserve the
// order. Preserving the order will have some performance overhead.
func ReadFile(ordered bool, path string) (*INI, error) {
	// TODO
	return nil, nil
}

// Write writes an INI into an io.Writer.
func Write(ini *INI, writer io.Writer) error {
	// TODO
	return nil
}

// WriteFile writes an INI into a file.
func WriteFile(ini *INI, path string) error {
	// TODO
	return nil
}
