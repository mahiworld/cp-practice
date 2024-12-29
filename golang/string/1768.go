package main

// func mergeAlternately(word1 string, word2 string) string {
//     len1 := len(word1)
//     len2 := len(word2)
//     if len1 == 0{
//         return word2
//     }else if len2 ==0{
//         return word1
//     }
//     length := len1
//     if len2<len1{
//         length = len2
//     }
//     merged := make([]string, len1+len2)
//     index := 0
//     for i := 0; i < length; i++ {
//         merged[index] = string(word1[i])
//         index = index + 1
//         merged[index] = string(word2[i])
//         index = index + 1
//     }
//     if len1>len2{
//         merged = append(merged, word1[index/2:])
//     }else if len1<len2{
//         merged = append(merged, word2[index/2:])
//     }

//     result := strings.Join(merged, "")
//     return result
// }

func mergeAlternately(word1 string, word2 string) string {

	len1 := len(word1)
	len2 := len(word2)

	i := 0
	j := 0

	mergedSlice := make([]byte, 0)

	for i < len1 || j < len2 {
		if i < len1 {
			mergedSlice = append(mergedSlice, word1[i])
			i++
		}

		if j < len2 {
			mergedSlice = append(mergedSlice, word2[j])
			j++
		}
	}

	return string(mergedSlice)
}
