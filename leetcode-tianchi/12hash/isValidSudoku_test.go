package hash

//
///*
//33. 有效的数独
//https://tianchi.aliyun.com/oj/problems/ujkbi1ha5crf9wts?spm=5176.15228502.0.0.760279bfXaWCfA
//*/
//import (
//	"testing"
//)
//
//func Test_isValidSudoku(t *testing.T) {
//	type args struct {
//		board [][]byte
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//		{"", args{[][]byte{}}, true},
//		{"", args{}, true},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := isValidSudoku(tt.args.board); got != tt.want {
//				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func isValidSudoku(board [][]byte) bool {
//
//	for i := 0; i < len(board); i++ {
//		rowsMap := make(map[byte]struct{}, 9)
//		for j := 0; j < len(board[i]); j++ {
//			if _, ok := rowsMap[board[i][j]]; ok {
//				return false
//			}
//			rowsMap[board[i][j]] = struct{}{}
//		}
//	}
//
//	for i := 0; i < len(board); i++ {
//		columnsMap := make(map[byte]struct{}, 9)
//		for j := 0; j < len(board[i]); j++ {
//			if _, ok := columnsMap[board[i][j]]; ok && board[i][j] != '.' {
//				return false
//			}
//			columnsMap[board[i][j]] = struct{}{}
//		}
//	}
//	for i, j := 0, 0; i < len(board) && j < len(board[i]); {
//
//	}
//
//	return true
//}
