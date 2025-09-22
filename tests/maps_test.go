package tests

import (
	"github.com/xierui921326/toolkit/maps"
	"sync"
	"testing"
)

// 测试 GetMapKeys 函数
func TestGetMapKeys(t *testing.T) {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	keys := maps.GetMapKeys(m)
	if len(keys) != len(m) {
		t.Errorf("Expected %d keys, got %d", len(m), len(keys))
	}
	for _, k := range keys {
		if _, exists := m[k]; !exists {
			t.Errorf("Key %s not found in original map", k)
		}
	}
}

// 测试 GetMapKeysMatchCondition 函数
func TestGetMapKeysMatchCondition(t *testing.T) {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	condition := func(k string, v int) bool {
		return v > 1
	}
	keys := maps.GetMapKeysMatchCondition(m, condition)
	for _, k := range keys {
		if m[k] <= 1 {
			t.Errorf("Key %s should not be included as its value is not greater than 1", k)
		}
	}
}

// 测试 GetMapValues 函数
func TestGetMapValues(t *testing.T) {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	values := maps.GetMapValues(m)
	if len(values) != len(m) {
		t.Errorf("Expected %d values, got %d", len(m), len(values))
	}
	valueSet := make(map[int]bool)
	for _, v := range values {
		valueSet[v] = true
	}
	for _, v := range m {
		if !valueSet[v] {
			t.Errorf("Value %d not found in result values", v)
		}
	}
}

// 测试 GetMapValuesMatchCondition 函数
func TestGetMapValuesMatchCondition(t *testing.T) {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	condition := func(k string, v int) bool {
		return v > 1
	}
	values := maps.GetMapValuesMatchCondition(m, condition)
	for _, v := range values {
		if v <= 1 {
			t.Errorf("Value %d should not be included as it is not greater than 1", v)
		}
	}
}

// 测试 MergeMaps 函数
func TestMergeMaps(t *testing.T) {
	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
	}
	m2 := map[string]int{
		"banana": 3,
		"cherry": 4,
	}
	merged := maps.MergeMaps(m1, m2)
	for k := range m1 {
		if _, exists := merged[k]; !exists {
			t.Errorf("key %s not exist in m1", k)
		}
	}
	for k := range m2 {
		if _, exists := merged[k]; !exists {
			t.Errorf("key %s not exist in m2", k)
		}
	}
}

// 测试 MapToSlice 函数
func TestMapToSlice(t *testing.T) {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	extractor := func(v int) int {
		return v * 2
	}
	slice := maps.MapToSlice(m, extractor)
	if len(slice) != len(m) {
		t.Errorf("Expected %d elements in slice, got %d", len(m), len(slice))
	}
	for _, v := range m {
		found := false
		for _, s := range slice {
			if s == v*2 {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Transformed value %d not found in slice", v*2)
		}
	}
}

// 测试 SliceToMap 函数
func TestSliceToMap(t *testing.T) {
	s := []int{1, 2, 3}
	keyFunc := func(v int) int {
		return v * 10
	}
	m := maps.SliceToMap(s, keyFunc)
	if len(m) != len(s) {
		t.Errorf("Expected %d keys in map, got %d", len(s), len(m))
	}
	for _, v := range s {
		key := keyFunc(v)
		if val, exists := m[key]; !exists || val != v {
			t.Errorf("Value for key %d in map does not match original slice", key)
		}
	}
}

// 测试 CopyMap 函数
func TestCopyMap(t *testing.T) {
	oldMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	newMap := maps.CopyMap(oldMap)
	if len(newMap) != len(oldMap) {
		t.Errorf("Expected %d keys in new map, got %d", len(oldMap), len(newMap))
	}
	for k, v := range oldMap {
		if val, exists := newMap[k]; !exists || val != v {
			t.Errorf("Value for key %s in new map does not match original map", k)
		}
	}
}

// 测试 SyncMapToMap 函数
func TestSyncMapToMap(t *testing.T) {
	var syncMap sync.Map
	syncMap.Store("apple", 1)
	syncMap.Store("banana", 2)
	syncMap.Store("cherry", 3)
	m := maps.SyncMapToMap[string, int](syncMap)
	syncMap.Range(func(key, value any) bool {
		k := key.(string)
		v := value.(int)
		if val, exists := m[k]; !exists || val != v {
			t.Errorf("Value for key %s in map does not match sync.Map", k)
		}
		return true
	})
}

// 测试 MapToSyncMap 函数
func TestMapToSyncMap(t *testing.T) {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	syncMap := maps.MapToSyncMap(m)
	for k, v := range m {
		value, exists := syncMap.Load(k)
		if !exists || value.(int) != v {
			t.Errorf("Value for key %s in sync.Map does not match original map", k)
		}
	}
}
