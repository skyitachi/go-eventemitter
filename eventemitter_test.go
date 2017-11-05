package eventemitter

import (
  "testing"
)
var called = false

func simpleFind(key string, data []string) bool {
  for _, v := range data {
    if v == key {
      return true
    }
  }
  return false
}

func add(params ...interface{}) {
  called = true
  ret := 0
  for _, param := range params {
    ret += param.(int)
  }
}

func TestEventEmitter(t *testing.T) {
  em := &EventEmitter{}
  em.On("test", add)
  em.Emit("test", 1, 2, 3)
  if called != true {
    t.Errorf("callback should be called, expect true found false")
  }
}

func TestEventNames(t *testing.T) {
  em := &EventEmitter{}
  em.On("test", add)
  em.On("test", add)
  em.On("test2", add)
  eventNames := em.EventNames()
  if len(eventNames) != 2 {
    t.Errorf("eventNames expect length 2 found %d", len(eventNames))
  }
  if !simpleFind("test", eventNames) {
    t.Errorf("event test should be in eventNames, but found false")
  }
  if !simpleFind("test2", eventNames) {
    t.Errorf("event test2 should be in eventNames, but found false")
  }
}