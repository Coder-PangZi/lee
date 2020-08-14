package main

import (
	"testing"
)

type metaCond byte

type cond []metaCond

type conditions []cond

func NewCond(mc ...metaCond) cond {
	c := make(cond, 0, 6)
	for _, m := range mc {
		c = append(c, m)
	}
	return c
}

func NewConditions(c ...cond) conditions {
	cs := make(conditions, 0, 6)
	for _, m := range c {
		cs = append(cs, m)
	}
	return cs
}

func TestName(t *testing.T) {
}
