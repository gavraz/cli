package cli

import "strconv"

type Flag interface {
	// ID is the flag's identifier.
	ID() string
	// Value returns the value of this flag.
	Value() any
	// WithValue attempts to parse the provided string value based on the flag's parsing rules.
	// If successful, it returns a new flag with the parsed value.
	WithValue(string) (Flag, error)
	// Obligatory returns true if this flag's value must be specified.
	Obligatory() bool
}

type BoolFlag struct {
	Name        string
	Description string
	Required    bool
	Default     bool
	isSet       bool
	value       bool
}

func (f BoolFlag) ID() string {
	return f.Name
}

func (f BoolFlag) Obligatory() bool {
	return f.Required
}

func (f BoolFlag) Value() any {
	if f.isSet {
		return f.value
	}

	return f.Default
}

func (f BoolFlag) WithValue(v string) (Flag, error) {
	b, err := strconv.ParseBool(v)
	if err != nil {
		return nil, err
	}
	f.value = b
	f.isSet = true
	return f, nil
}

type StringFlag struct {
	Name        string
	Description string
	Required    bool
	Default     string

	isSet bool
	value string
}

func (f StringFlag) ID() string {
	return f.Name
}

func (f StringFlag) Obligatory() bool {
	return f.Required
}

func (f StringFlag) Value() any {
	if f.isSet {
		return f.value
	}

	return f.Default
}

func (f StringFlag) WithValue(v string) (Flag, error) {
	f.value = v
	f.isSet = true
	return f, nil
}

type IntFlag struct {
	Name        string
	Description string
	Required    bool
	Default     int

	isSet bool
	value int
}

func (f IntFlag) ID() string {
	return f.Name
}

func (f IntFlag) Obligatory() bool {
	return f.Required
}

func (f IntFlag) Value() any {
	if f.isSet {
		return f.value
	}

	return f.Default
}

func (f IntFlag) WithValue(v string) (Flag, error) {
	val, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	f.value = val
	f.isSet = true
	return f, nil
}

type Float32Flag struct {
	Name        string
	Description string
	Required    bool
	Default     float32

	isSet bool
	value float32
}

func (f Float32Flag) ID() string {
	return f.Name
}

func (f Float32Flag) Obligatory() bool {
	return f.Required
}

func (f Float32Flag) Value() any {
	if f.isSet {
		return f.value
	}

	return f.Default
}

func (f Float32Flag) WithValue(v string) (Flag, error) {
	val, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return nil, err
	}
	f.value = float32(val)
	f.isSet = true
	return f, nil
}
