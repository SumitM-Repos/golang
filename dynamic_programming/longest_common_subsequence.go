package dynamic_programming

func Longest_Common_SubSequence(seq1, seq2 string, seq1_len, seq2_len int, arr [][]int) int {

	seq1_chars := []rune(seq1)
	seq2_chars := []rune(seq2)

	if seq1_len == 0 || seq2_len == 0 {
		return 0
	} else if seq1_chars[seq1_len-1] == seq2_chars[seq2_len-1] {
		arr[seq1_len][seq2_len] = 1 + Longest_Common_SubSequence(seq1, seq2, seq1_len-1, seq2_len-1, arr)
		return arr[seq1_len][seq2_len]

	} else if arr[seq1_len][seq2_len] != -1 {
		return arr[seq1_len][seq2_len]
	} else {
		arr[seq1_len][seq2_len] = max(Longest_Common_SubSequence(seq1, seq2, seq1_len-1, seq2_len, arr), Longest_Common_SubSequence(seq1, seq2, seq1_len, seq2_len-1, arr))
		return arr[seq1_len][seq2_len]
	}
}
