/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
package main

import (
	"sort"
)

type bytes []byte

func (b bytes) Len() int {
	return len(b)
}

func (b bytes) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func groupAnagrams2(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	groups := make(map[string][]string, 0)
	// 1. 循环字符串数组
	for _, str := range strs {
		keyBytes := bytes(str)
		sort.Sort(keyBytes)
		key := string(keyBytes) // 经过自动以排序后的key
		// 2. 将每个字符串按照一定规则排序组成一个有顺序的key存入hashmap中
		if _, ok := groups[key]; ok {
			groups[key] = append(groups[key], str)
		} else {
			groups[key] = []string{str}
		}
	}
	// 3. 将hashmap中每个键的值存入最终结果中
	result := [][]string{}
	for _, val := range groups {
		result = append(result, val)
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	keyMap := make(map[string]int, 0)
	// 3. 将hashmap中每个键的值存入最终结果中
	result := [][]string{}
	// 1. 循环字符串数组
	for _, str := range strs {
		keyBytes := bytes(str)
		sort.Sort(keyBytes)
		key := string(keyBytes) // 经过自动以排序后的key
		// 2. 将每个字符串按照一定规则排序组成一个有顺序的key存入hashmap中
		if index, ok := keyMap[key]; ok {
			result[index] = append(result[index] , str)
		} else {
			keyMap[key] = len(result)
			result = append(result, []string{str})
		}
	}
	return result
}

// @lc code=end
