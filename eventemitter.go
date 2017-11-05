package eventemitter

type Callback func (...interface{})

type EventEmitter struct {
  cbMap map[string][]Callback
}

func (em *EventEmitter) initCbMap() bool {
  if em.cbMap == nil {
    em.cbMap = map[string][]Callback{}
    return false
  }
  return true
}

func (em *EventEmitter) getCbQueue(eventName string) ([]Callback, bool) {
  em.initCbMap()
  ret, ok := em.cbMap[eventName]
  return ret, ok
}

func (em *EventEmitter) On(eventName string, cb Callback) {
  em.initCbMap()
  cbq, ok := em.cbMap[eventName]
  if !ok {
    cbq = []Callback{}
  }
  cbq = append(cbq, cb)
  em.cbMap[eventName] = cbq
}

func (em *EventEmitter) Emit(eventName string, params ...interface{}) {
  if em.cbMap == nil {
    em.cbMap = map[string][]Callback{}
    return
  }
  cbq, ok := em.cbMap[eventName]
  if !ok {
    return
  }
  for _, cb := range cbq {
    cb(params...)
  }
}

func (em *EventEmitter) EventNames() []string {
  init := em.initCbMap()
  if !init {
    return []string{}

  }
  ret := []string{}
  for name := range em.cbMap {
    ret = append(ret, name)
  }
  return ret
}

func (em *EventEmitter) RemoveAllListeners(eventName string) {
  init := em.initCbMap()
  if !init {
    return
  }
  em.cbMap[eventName] = []Callback{}
  delete(em.cbMap, eventName)
}

