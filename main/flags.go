package main

import "strconv"

type Flag interface {
	// ID is the flag's identifier.
	ID() string
	// Value returns the value of this flag.
	Value() any
	// Parse tries to parse the provided string value according to the parse rules of this flag.
	// On success, it returns a new flag with the new interpreted value.
	Parse(string) (Flag, error)
}

type BoolFlag struct {
	Name        string
	Description string
	Required    bool
	Default     bool
	isSet       bool
	value       bool
}

func (bf BoolFlag) ID() string {
	return bf.Name
}

func (bf BoolFlag) Value() any {
	if bf.isSet {
		return bf.value
	}

	return bf.Default
}

func (bf BoolFlag) Parse(v string) (Flag, error) {
	b, err := strconv.ParseBool(v)
	if err != nil {
		return nil, err
	}
	bf.value = b
	bf.isSet = true
	return bf, nil
}

type StringFlag struct {
	Name        string
	Description string
	Required    bool
	Default     string

	isSet bool
	value string
}

func (sf StringFlag) ID() string {
	return sf.Name
}

func (sf StringFlag) Value() any {
	if sf.isSet {
		return sf.value
	}

	return sf.Default
}

func (sf StringFlag) Parse(v string) (Flag, error) {
	sf.value = v
	sf.isSet = true
	return sf, nil
}

type IntFlag struct {
	Name        string
	Description string
	Required    bool
	Default     int

	isSet bool
	value int
}

func (sf IntFlag) ID() string {
	return sf.Name
}

func (sf IntFlag) Value() any {
	if sf.isSet {
		return sf.value
	}

	return sf.Default
}

func (sf IntFlag) Parse(v string) (Flag, error) {
	val, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	sf.value = val
	sf.isSet = true
	return sf, nil
}
