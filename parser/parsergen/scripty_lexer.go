// Generated from scripty.g4 by ANTLR 4.7.

package parsergen

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 283,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 5, 10, 63, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	7, 12, 69, 10, 12, 12, 12, 14, 12, 72, 11, 12, 3, 13, 3, 13, 3, 13, 6,
	13, 77, 10, 13, 13, 13, 14, 13, 78, 3, 13, 3, 13, 3, 13, 5, 13, 84, 10,
	13, 3, 13, 6, 13, 87, 10, 13, 13, 13, 14, 13, 88, 5, 13, 91, 10, 13, 3,
	13, 3, 13, 6, 13, 95, 10, 13, 13, 13, 14, 13, 96, 3, 13, 5, 13, 100, 10,
	13, 3, 13, 3, 13, 6, 13, 104, 10, 13, 13, 13, 14, 13, 105, 3, 13, 5, 13,
	109, 10, 13, 5, 13, 111, 10, 13, 3, 13, 6, 13, 114, 10, 13, 13, 13, 14,
	13, 115, 3, 13, 3, 13, 7, 13, 120, 10, 13, 12, 13, 14, 13, 123, 11, 13,
	3, 13, 3, 13, 6, 13, 127, 10, 13, 13, 13, 14, 13, 128, 5, 13, 131, 10,
	13, 3, 13, 3, 13, 5, 13, 135, 10, 13, 3, 13, 6, 13, 138, 10, 13, 13, 13,
	14, 13, 139, 5, 13, 142, 10, 13, 3, 13, 5, 13, 145, 10, 13, 3, 13, 6, 13,
	148, 10, 13, 13, 13, 14, 13, 149, 3, 13, 3, 13, 3, 13, 5, 13, 155, 10,
	13, 3, 13, 6, 13, 158, 10, 13, 13, 13, 14, 13, 159, 3, 13, 5, 13, 163,
	10, 13, 3, 13, 5, 13, 166, 10, 13, 5, 13, 168, 10, 13, 3, 14, 5, 14, 171,
	10, 14, 3, 14, 5, 14, 174, 10, 14, 3, 14, 5, 14, 177, 10, 14, 3, 14, 5,
	14, 180, 10, 14, 5, 14, 182, 10, 14, 3, 14, 3, 14, 3, 14, 6, 14, 187, 10,
	14, 13, 14, 14, 14, 188, 3, 14, 5, 14, 192, 10, 14, 3, 14, 5, 14, 195,
	10, 14, 3, 14, 5, 14, 198, 10, 14, 3, 14, 7, 14, 201, 10, 14, 12, 14, 14,
	14, 204, 11, 14, 3, 14, 3, 14, 3, 14, 3, 14, 6, 14, 210, 10, 14, 13, 14,
	14, 14, 211, 3, 14, 5, 14, 215, 10, 14, 3, 14, 5, 14, 218, 10, 14, 3, 14,
	5, 14, 221, 10, 14, 3, 14, 7, 14, 224, 10, 14, 12, 14, 14, 14, 227, 11,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 237,
	10, 14, 12, 14, 14, 14, 240, 11, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 252, 10, 14, 12, 14, 14, 14,
	255, 11, 14, 3, 14, 3, 14, 3, 14, 5, 14, 260, 10, 14, 3, 15, 6, 15, 263,
	10, 15, 13, 15, 14, 15, 264, 3, 15, 3, 15, 3, 16, 3, 16, 7, 16, 271, 10,
	16, 12, 16, 14, 16, 274, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 18, 3, 18, 4, 238, 253, 2, 19, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 3, 2, 24, 4, 2, 62, 62, 64, 64, 5, 2, 44, 45, 47, 47, 49,
	49, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124,
	4, 2, 90, 90, 122, 122, 5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 78, 78, 110,
	110, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 3, 2, 50, 59, 4, 2,
	81, 81, 113, 113, 3, 2, 50, 57, 4, 2, 68, 68, 100, 100, 3, 2, 50, 51, 4,
	2, 76, 76, 108, 108, 6, 2, 68, 68, 87, 87, 100, 100, 119, 119, 4, 2, 84,
	84, 116, 116, 4, 2, 11, 11, 34, 34, 6, 2, 12, 12, 15, 15, 41, 41, 94, 94,
	6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 3, 2, 94, 94, 4, 2, 12, 12, 15, 15,
	2, 340, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 3, 37, 3, 2, 2, 2, 5, 41, 3, 2, 2,
	2, 7, 43, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 47, 3, 2, 2, 2, 13, 49, 3,
	2, 2, 2, 15, 51, 3, 2, 2, 2, 17, 53, 3, 2, 2, 2, 19, 62, 3, 2, 2, 2, 21,
	64, 3, 2, 2, 2, 23, 66, 3, 2, 2, 2, 25, 167, 3, 2, 2, 2, 27, 181, 3, 2,
	2, 2, 29, 262, 3, 2, 2, 2, 31, 268, 3, 2, 2, 2, 33, 277, 3, 2, 2, 2, 35,
	281, 3, 2, 2, 2, 37, 38, 7, 102, 2, 2, 38, 39, 7, 103, 2, 2, 39, 40, 7,
	104, 2, 2, 40, 4, 3, 2, 2, 2, 41, 42, 7, 42, 2, 2, 42, 6, 3, 2, 2, 2, 43,
	44, 7, 43, 2, 2, 44, 8, 3, 2, 2, 2, 45, 46, 7, 125, 2, 2, 46, 10, 3, 2,
	2, 2, 47, 48, 7, 127, 2, 2, 48, 12, 3, 2, 2, 2, 49, 50, 7, 46, 2, 2, 50,
	14, 3, 2, 2, 2, 51, 52, 7, 63, 2, 2, 52, 16, 3, 2, 2, 2, 53, 54, 7, 61,
	2, 2, 54, 18, 3, 2, 2, 2, 55, 56, 7, 40, 2, 2, 56, 63, 7, 40, 2, 2, 57,
	58, 7, 126, 2, 2, 58, 63, 7, 126, 2, 2, 59, 63, 9, 2, 2, 2, 60, 61, 7,
	63, 2, 2, 61, 63, 7, 63, 2, 2, 62, 55, 3, 2, 2, 2, 62, 57, 3, 2, 2, 2,
	62, 59, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 20, 3, 2, 2, 2, 64, 65, 9,
	3, 2, 2, 65, 22, 3, 2, 2, 2, 66, 70, 9, 4, 2, 2, 67, 69, 9, 5, 2, 2, 68,
	67, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2,
	2, 71, 24, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 110, 7, 50, 2, 2, 74, 76,
	9, 6, 2, 2, 75, 77, 9, 7, 2, 2, 76, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2,
	78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 90, 3, 2, 2, 2, 80, 91, 9,
	8, 2, 2, 81, 83, 9, 9, 2, 2, 82, 84, 9, 10, 2, 2, 83, 82, 3, 2, 2, 2, 83,
	84, 3, 2, 2, 2, 84, 86, 3, 2, 2, 2, 85, 87, 9, 11, 2, 2, 86, 85, 3, 2,
	2, 2, 87, 88, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 91,
	3, 2, 2, 2, 90, 80, 3, 2, 2, 2, 90, 81, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2,
	91, 111, 3, 2, 2, 2, 92, 94, 9, 12, 2, 2, 93, 95, 9, 13, 2, 2, 94, 93,
	3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2,
	97, 99, 3, 2, 2, 2, 98, 100, 9, 8, 2, 2, 99, 98, 3, 2, 2, 2, 99, 100, 3,
	2, 2, 2, 100, 111, 3, 2, 2, 2, 101, 103, 9, 14, 2, 2, 102, 104, 9, 15,
	2, 2, 103, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2,
	105, 106, 3, 2, 2, 2, 106, 108, 3, 2, 2, 2, 107, 109, 9, 8, 2, 2, 108,
	107, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 111, 3, 2, 2, 2, 110, 74, 3,
	2, 2, 2, 110, 92, 3, 2, 2, 2, 110, 101, 3, 2, 2, 2, 111, 168, 3, 2, 2,
	2, 112, 114, 9, 11, 2, 2, 113, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115,
	113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 121,
	7, 48, 2, 2, 118, 120, 9, 11, 2, 2, 119, 118, 3, 2, 2, 2, 120, 123, 3,
	2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 131, 3, 2, 2,
	2, 123, 121, 3, 2, 2, 2, 124, 126, 7, 48, 2, 2, 125, 127, 9, 11, 2, 2,
	126, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128,
	129, 3, 2, 2, 2, 129, 131, 3, 2, 2, 2, 130, 113, 3, 2, 2, 2, 130, 124,
	3, 2, 2, 2, 131, 141, 3, 2, 2, 2, 132, 134, 9, 9, 2, 2, 133, 135, 9, 10,
	2, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2, 2, 2,
	136, 138, 9, 11, 2, 2, 137, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139,
	137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 142, 3, 2, 2, 2, 141, 132,
	3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 144, 3, 2, 2, 2, 143, 145, 9, 16,
	2, 2, 144, 143, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 168, 3, 2, 2, 2,
	146, 148, 9, 11, 2, 2, 147, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149,
	147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 165, 3, 2, 2, 2, 151, 166,
	9, 8, 2, 2, 152, 154, 9, 9, 2, 2, 153, 155, 9, 10, 2, 2, 154, 153, 3, 2,
	2, 2, 154, 155, 3, 2, 2, 2, 155, 157, 3, 2, 2, 2, 156, 158, 9, 11, 2, 2,
	157, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159,
	160, 3, 2, 2, 2, 160, 162, 3, 2, 2, 2, 161, 163, 9, 16, 2, 2, 162, 161,
	3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 166, 3, 2, 2, 2, 164, 166, 9, 16,
	2, 2, 165, 151, 3, 2, 2, 2, 165, 152, 3, 2, 2, 2, 165, 164, 3, 2, 2, 2,
	165, 166, 3, 2, 2, 2, 166, 168, 3, 2, 2, 2, 167, 73, 3, 2, 2, 2, 167, 130,
	3, 2, 2, 2, 167, 147, 3, 2, 2, 2, 168, 26, 3, 2, 2, 2, 169, 171, 9, 17,
	2, 2, 170, 169, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2,
	172, 174, 9, 18, 2, 2, 173, 172, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174,
	182, 3, 2, 2, 2, 175, 177, 9, 18, 2, 2, 176, 175, 3, 2, 2, 2, 176, 177,
	3, 2, 2, 2, 177, 179, 3, 2, 2, 2, 178, 180, 9, 17, 2, 2, 179, 178, 3, 2,
	2, 2, 179, 180, 3, 2, 2, 2, 180, 182, 3, 2, 2, 2, 181, 170, 3, 2, 2, 2,
	181, 176, 3, 2, 2, 2, 182, 259, 3, 2, 2, 2, 183, 202, 7, 41, 2, 2, 184,
	197, 7, 94, 2, 2, 185, 187, 9, 19, 2, 2, 186, 185, 3, 2, 2, 2, 187, 188,
	3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 194, 3, 2,
	2, 2, 190, 192, 7, 15, 2, 2, 191, 190, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2,
	192, 193, 3, 2, 2, 2, 193, 195, 7, 12, 2, 2, 194, 191, 3, 2, 2, 2, 194,
	195, 3, 2, 2, 2, 195, 198, 3, 2, 2, 2, 196, 198, 11, 2, 2, 2, 197, 186,
	3, 2, 2, 2, 197, 196, 3, 2, 2, 2, 198, 201, 3, 2, 2, 2, 199, 201, 10, 20,
	2, 2, 200, 184, 3, 2, 2, 2, 200, 199, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2,
	202, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 205, 3, 2, 2, 2, 204,
	202, 3, 2, 2, 2, 205, 260, 7, 41, 2, 2, 206, 225, 7, 36, 2, 2, 207, 220,
	7, 94, 2, 2, 208, 210, 9, 19, 2, 2, 209, 208, 3, 2, 2, 2, 210, 211, 3,
	2, 2, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 217, 3, 2, 2,
	2, 213, 215, 7, 15, 2, 2, 214, 213, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215,
	216, 3, 2, 2, 2, 216, 218, 7, 12, 2, 2, 217, 214, 3, 2, 2, 2, 217, 218,
	3, 2, 2, 2, 218, 221, 3, 2, 2, 2, 219, 221, 11, 2, 2, 2, 220, 209, 3, 2,
	2, 2, 220, 219, 3, 2, 2, 2, 221, 224, 3, 2, 2, 2, 222, 224, 10, 21, 2,
	2, 223, 207, 3, 2, 2, 2, 223, 222, 3, 2, 2, 2, 224, 227, 3, 2, 2, 2, 225,
	223, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 228, 3, 2, 2, 2, 227, 225,
	3, 2, 2, 2, 228, 260, 7, 36, 2, 2, 229, 230, 7, 36, 2, 2, 230, 231, 7,
	36, 2, 2, 231, 232, 7, 36, 2, 2, 232, 238, 3, 2, 2, 2, 233, 234, 7, 94,
	2, 2, 234, 237, 11, 2, 2, 2, 235, 237, 10, 22, 2, 2, 236, 233, 3, 2, 2,
	2, 236, 235, 3, 2, 2, 2, 237, 240, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 238,
	236, 3, 2, 2, 2, 239, 241, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 241, 242,
	7, 36, 2, 2, 242, 243, 7, 36, 2, 2, 243, 260, 7, 36, 2, 2, 244, 245, 7,
	41, 2, 2, 245, 246, 7, 41, 2, 2, 246, 247, 7, 41, 2, 2, 247, 253, 3, 2,
	2, 2, 248, 249, 7, 94, 2, 2, 249, 252, 11, 2, 2, 2, 250, 252, 10, 22, 2,
	2, 251, 248, 3, 2, 2, 2, 251, 250, 3, 2, 2, 2, 252, 255, 3, 2, 2, 2, 253,
	254, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 254, 256, 3, 2, 2, 2, 255, 253,
	3, 2, 2, 2, 256, 257, 7, 41, 2, 2, 257, 258, 7, 41, 2, 2, 258, 260, 7,
	41, 2, 2, 259, 183, 3, 2, 2, 2, 259, 206, 3, 2, 2, 2, 259, 229, 3, 2, 2,
	2, 259, 244, 3, 2, 2, 2, 260, 28, 3, 2, 2, 2, 261, 263, 9, 19, 2, 2, 262,
	261, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 264, 265,
	3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267, 8, 15, 2, 2, 267, 30, 3, 2,
	2, 2, 268, 272, 7, 37, 2, 2, 269, 271, 10, 23, 2, 2, 270, 269, 3, 2, 2,
	2, 271, 274, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273,
	275, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 275, 276, 8, 16, 3, 2, 276, 32,
	3, 2, 2, 2, 277, 278, 11, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 280, 8, 17,
	3, 2, 280, 34, 3, 2, 2, 2, 281, 282, 11, 2, 2, 2, 282, 36, 3, 2, 2, 2,
	52, 2, 62, 70, 78, 83, 88, 90, 96, 99, 105, 108, 110, 115, 121, 128, 130,
	134, 139, 141, 144, 149, 154, 159, 162, 165, 167, 170, 173, 176, 179, 181,
	188, 191, 194, 197, 200, 202, 211, 214, 217, 220, 223, 225, 236, 238, 251,
	253, 259, 264, 272, 4, 2, 3, 2, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'def'", "'('", "')'", "'{'", "'}'", "','", "'='", "';'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "ASSIGNMENT_OP", "STMT_END", "BOOLEAN_OP",
	"ARITHMETIC_OP", "NAME", "NUMBER", "STRING", "WHITESPACE", "COMMENT", "UNKNOWN",
	"ANY",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "ASSIGNMENT_OP", "STMT_END",
	"BOOLEAN_OP", "ARITHMETIC_OP", "NAME", "NUMBER", "STRING", "WHITESPACE",
	"COMMENT", "UNKNOWN", "ANY",
}

type scriptyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewscriptyLexer(input antlr.CharStream) *scriptyLexer {

	l := new(scriptyLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "scripty.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// scriptyLexer tokens.
const (
	scriptyLexerT__0          = 1
	scriptyLexerT__1          = 2
	scriptyLexerT__2          = 3
	scriptyLexerT__3          = 4
	scriptyLexerT__4          = 5
	scriptyLexerT__5          = 6
	scriptyLexerASSIGNMENT_OP = 7
	scriptyLexerSTMT_END      = 8
	scriptyLexerBOOLEAN_OP    = 9
	scriptyLexerARITHMETIC_OP = 10
	scriptyLexerNAME          = 11
	scriptyLexerNUMBER        = 12
	scriptyLexerSTRING        = 13
	scriptyLexerWHITESPACE    = 14
	scriptyLexerCOMMENT       = 15
	scriptyLexerUNKNOWN       = 16
	scriptyLexerANY           = 17
)
