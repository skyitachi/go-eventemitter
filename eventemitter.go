package eventemitter

type Callback func (...interface{})

type EventEmitter struct {
  cbMap map[string][]Callback
  EventNames []string
}

func (em *EventEmitter) on(eventName string, cb Callback) {
  if em.cbMap == nil {
    em.cbMap = map[string][]Callback{}
  }
  cbq, ok := em.cbMap[eventName]
  if !ok {
    cbq = []Callback{}
  }
  cbq = append(cbq, cb)
  em.cbMap[eventName] = cbq
}

func (em *EventEmitter) emit(eventName string, params ...interface{}) {
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


