package main

import "strconv"

type Flag interface {
	ID() string
	IsSet() bool
	Parse(string) error
}

type BoolFlag struct {
	Name        string
	Description string
	Required    bool
	Default     bool

	isSet bool
	value bool
}

func (bf *BoolFlag) ID() string {
	return bf.Name
}

func (bf *BoolFlag) IsSet() bool {
	return bf.isSet
}

func (bf *BoolFlag) Value() bool {
	if bf.isSet {
		return bf.value
	}

	return bf.Default
}

func (bf *BoolFlag) Set(b bool) {
	bf.value = b
	bf.isSet = true
}

func (bf *BoolFlag) Parse(v string) error {
	b, err := strconv.ParseBool(v)
	if err != nil {
		return err
	}
	bf.Set(b)
	return nil
}

type StringFlag struct {
	Name        string
	Description string
	Required    bool
	Default     string

	isSet bool
	value string
}

func (sf *StringFlag) ID() string {
	return sf.Name
}

func (sf *StringFlag) IsSet() bool {
	return sf.isSet
}

func (sf *StringFlag) Value() string {
	if sf.isSet {
		return sf.value
	}

	return sf.Default
}

func (sf *StringFlag) Set(b string) {
	sf.value = b
	sf.isSet = true
}

func (sf *StringFlag) Parse(v string) error {
	sf.Set(v)
	return nil
}

type IntFlag struct {
	Name        string
	Description string
	Required    bool
	Default     int

	isSet bool
	value int
}

func (sf *IntFlag) ID() string {
	return sf.Name
}

func (sf *IntFlag) IsSet() bool {
	return sf.isSet
}

func (sf *IntFlag) Value() int {
	if sf.isSet {
		return sf.value
	}

	return sf.Default
}

func (sf *IntFlag) Set(b int) {
	sf.value = b
	sf.isSet = true
}

func (sf *IntFlag) Parse(v string) error {
	val, err := strconv.Atoi(v)
	if err != nil {
		return err
	}
	sf.Set(val)
	return nil
}
