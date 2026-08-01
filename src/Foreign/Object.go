package Object

import "gopurs/output/gopurs_runtime"

func _CopyST(m map[string]interface{}) func(interface{}) interface{} {
	return func(_ interface{}) interface{} {
		newMap := make(map[string]interface{})
		for k, v := range m {
			newMap[k] = v
		}
		return &newMap
	}
}

var Empty = map[string]interface{}{}

func RunST(f func(interface{}) interface{}) interface{} {
	val := f(nil).(gopurs_runtime.Value)
	mPtr := val.PtrVal().(*map[string]interface{})
	return *mPtr
}

func _FmapObject(m0 map[string]interface{}, f func(interface{}) interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range m0 {
		m[k] = f(v)
	}
	return m
}

func _MapWithKey(m0 map[string]interface{}, f func(string) func(interface{}) interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range m0 {
		m[k] = f(k)(v)
	}
	return m
}

func _FoldM(bind func(interface{}) func(interface{}) interface{}) func(func(interface{}) func(string) func(interface{}) interface{}) func(interface{}) func(map[string]interface{}) interface{} {
	return func(f func(interface{}) func(string) func(interface{}) interface{}) func(interface{}) func(map[string]interface{}) interface{} {
		return func(mz interface{}) func(map[string]interface{}) interface{} {
			return func(m map[string]interface{}) interface{} {
				acc := mz
				g := func(k string) func(interface{}) interface{} {
					return func(z interface{}) interface{} {
						return f(z)(k)(m[k])
					}
				}
				for k := range m {
					acc = bind(acc)(g(k))
				}
				return acc
			}
		}
	}
}

func _FoldSCObject(m map[string]interface{}, z interface{}, f func(interface{}) func(string) func(interface{}) interface{}, fromMaybe func(interface{}) func(interface{}) interface{}) interface{} {
	acc := z
	for k, v := range m {
		maybeR := f(acc)(k)(v)
		r := fromMaybe(nil)(maybeR)
		if r == nil {
			return acc
		} else {
			acc = r
		}
	}
	return acc
}

func All(f func(string) func(interface{}) bool) func(map[string]interface{}) bool {
	return func(m map[string]interface{}) bool {
		for k, v := range m {
			if !f(k)(v) {
				return false
			}
		}
		return true
	}
}

func Size(m map[string]interface{}) int64 {
	return int64(len(m))
}

func _Lookup(no interface{}, yes func(interface{}) interface{}, k string, m map[string]interface{}) interface{} {
	if val, ok := m[k]; ok {
		return yes(val)
	}
	return no
}

func _LookupST(no interface{}, yes func(interface{}) interface{}, k string, m *map[string]interface{}) func(interface{}) interface{} {
	return func(_ interface{}) interface{} {
		if val, ok := (*m)[k]; ok {
			return yes(val)
		}
		return no
	}
}

func ToArrayWithKey(f func(string) func(interface{}) interface{}) func(map[string]interface{}) []interface{} {
	return func(m map[string]interface{}) []interface{} {
		r := make([]interface{}, 0, len(m))
		for k, v := range m {
			r = append(r, f(k)(v))
		}
		return r
	}
}

func Keys(m map[string]interface{}) []interface{} {
	r := make([]interface{}, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}
