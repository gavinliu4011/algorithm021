/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := map[rune]int{}
	for _, ch := range s {
		count[ch]++
	}
	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end
