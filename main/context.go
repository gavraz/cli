package main

import "context"

type Flags struct {
	ctx context.Context
}

func (f *Flags) String(name string) string {
	return f.ctx.Value(name).(Flag).Value().(string)
}

func (f *Flags) Bool(name string) bool {
	return f.ctx.Value(name).(Flag).Value().(bool)
}

func (f *Flags) Int(name string) int {
	return f.ctx.Value(name).(Flag).Value().(int)
}

func (f *Flags) Float32(name string) float32 {
	return f.ctx.Value(name).(Flag).Value().(float32)
}
