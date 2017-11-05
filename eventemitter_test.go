package eventemitter

import (
  "testing"
)
var called = false
func add(params ...interface{}) {
  called = true
  ret := 0
  for _, param := range params {
    ret += param.(int)
  }
}

func TestEventEmitter(t *testing.T) {
  em := &EventEmitter{}
  em.on("test", add)
  em.emit("test", 1, 2, 3)
  if called != true {
    t.Errorf("callback should be called, expect true found false")
  }
}