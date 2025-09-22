package maps

import (
	"sync"
)

// GetMapKeys 获取map的所有keys
//
// @Description: 获取map的所有keys
// @param m map[K]V 要获取keys的map
// @return []K 所有keys
func GetMapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// GetMapKeysMatchCondition 根据过滤条件获取map的keys
//
// @Description: 根据过滤条件获取map的keys
// @param m map[K]V 要获取keys的map
// @param condition func(K, V) bool 过滤条件
// @return []K 所有符合条件的keys
func GetMapKeysMatchCondition[K comparable, V any](m map[K]V, condition func(K, V) bool) []K {
	keys := make([]K, 0, len(m))
	for k, v := range m {
		if condition(k, v) {
			keys = append(keys, k)
		}
	}
	return keys
}

// GetMapValues 获取map的所有values
//
// @Description: 获取map的所有values
// @param m map[K]V 要获取values的map
// @return []V 所有values
func GetMapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// GetMapValuesMatchCondition 根据过滤条件获取map的values
//
// @Description: 根据过滤条件获取map的values
// @param m map[K]V 要获取values的map
// @param condition func(K, V) bool 过滤条件
// @return []V 所有符合条件的values
func GetMapValuesMatchCondition[K comparable, V any](m map[K]V, condition func(K, V) bool) []V {
	values := make([]V, 0, len(m))
	for k, v := range m {
		if condition(k, v) {
			values = append(values, v)
		}
	}
	return values
}

// MergeMaps 合并多个map
//
// @Description: 合并多个map
// @param maps ...map[K]V 要合并的map
// @return map[K]V 合并后的map
func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	mergeMap := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			mergeMap[k] = v
		}
	}
	return mergeMap
}

// MapToSlice map转切片
//
// @Description: map转切片
// @param m map[K]V 要转化的map
// @param extractor func(V) T 转换函数
// @return []T 转换后的切片
func MapToSlice[K comparable, V any, T any](m map[K]V, extractor func(V) T) []T {
	res := make([]T, 0, len(m))
	for _, v := range m {
		res = append(res, extractor(v))
	}
	return res
}

// SliceToMap 切片转map
//
// @Description: 切片转map
// @param s []T 要转化的切片
// @param keyFunc func(T) K 键函数
// @return map[K]T 转换后的map
func SliceToMap[T any, K comparable](s []T, keyFunc func(T) K) map[K]T {
	res := make(map[K]T)
	for _, v := range s {
		key := keyFunc(v)
		res[key] = v
	}
	return res
}

// CopyMap 复制map
//
// @Description: 复制map
// @param oldMap map[K]V 要复制的map
// @return map[K]V 复制后的map
func CopyMap[K comparable, V any](oldMap map[K]V) map[K]V {
	newMap := make(map[K]V, len(oldMap))
	for k, v := range oldMap {
		newMap[k] = v
	}
	return newMap
}

// SyncMapToMap sync.Map转map
//
// @Description: sync.Map转map
// @param syncMap sync.Map 要转化的sync.Map
// @return map[K]V 转换后的map
func SyncMapToMap[K comparable, V any](syncMap sync.Map) map[K]V {
	m := make(map[K]V)
	syncMap.Range(func(key, value any) bool {
		// 尝试将key和value转换为指定的类型
		k, ok1 := key.(K)
		v, ok2 := value.(V)
		if ok1 && ok2 {
			m[k] = v
		}
		return true
	})
	return m
}

// MapToSyncMap map转sync.Map
//
// @Description: map转sync.Map
// @param m map[K]V 要转化的map
// @return *sync.Map 转换后的sync.Map
func MapToSyncMap[K comparable, V any](m map[K]V) *sync.Map {
	syncMap := &sync.Map{}
	if m == nil {
		return syncMap
	}
	for k, v := range m {
		syncMap.Store(k, v)
	}
	return syncMap
}
