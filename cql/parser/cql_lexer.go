// Generated from /home/zhichyu/src/github.com/deepfabric/indexer/cql/parser/CQL.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 35, 338,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3,
	30, 3, 30, 3, 30, 5, 30, 258, 10, 30, 3, 30, 5, 30, 261, 10, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 269, 10, 30, 5, 30, 271, 10,
	30, 3, 31, 6, 31, 274, 10, 31, 13, 31, 14, 31, 275, 3, 32, 3, 32, 5, 32,
	280, 10, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 7, 34, 289,
	10, 34, 12, 34, 14, 34, 292, 11, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35,
	5, 35, 299, 10, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3,
	37, 3, 38, 3, 38, 3, 38, 7, 38, 312, 10, 38, 12, 38, 14, 38, 315, 11, 38,
	5, 38, 317, 10, 38, 3, 39, 3, 39, 5, 39, 321, 10, 39, 3, 39, 3, 39, 3,
	40, 3, 40, 7, 40, 327, 10, 40, 12, 40, 14, 40, 330, 11, 40, 3, 41, 6, 41,
	333, 10, 41, 13, 41, 14, 41, 334, 3, 41, 3, 41, 2, 2, 42, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 2, 63,
	2, 65, 2, 67, 32, 69, 2, 71, 2, 73, 2, 75, 33, 77, 2, 79, 34, 81, 35, 3,
	2, 12, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 3, 2, 50, 59, 4, 2,
	36, 36, 94, 94, 10, 2, 36, 36, 49, 49, 94, 94, 100, 100, 104, 104, 112,
	112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72, 99, 104, 3, 2, 51, 59, 5,
	2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 5, 2,
	11, 12, 15, 15, 34, 34, 2, 345, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2,
	2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3,
	2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45,
	3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2,
	53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2,
	2, 67, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2,
	2, 3, 83, 3, 2, 2, 2, 5, 94, 3, 2, 2, 2, 7, 101, 3, 2, 2, 2, 9, 113, 3,
	2, 2, 2, 11, 124, 3, 2, 2, 2, 13, 132, 3, 2, 2, 2, 15, 143, 3, 2, 2, 2,
	17, 149, 3, 2, 2, 2, 19, 155, 3, 2, 2, 2, 21, 163, 3, 2, 2, 2, 23, 169,
	3, 2, 2, 2, 25, 171, 3, 2, 2, 2, 27, 173, 3, 2, 2, 2, 29, 175, 3, 2, 2,
	2, 31, 181, 3, 2, 2, 2, 33, 188, 3, 2, 2, 2, 35, 195, 3, 2, 2, 2, 37, 202,
	3, 2, 2, 2, 39, 210, 3, 2, 2, 2, 41, 218, 3, 2, 2, 2, 43, 223, 3, 2, 2,
	2, 45, 230, 3, 2, 2, 2, 47, 233, 3, 2, 2, 2, 49, 242, 3, 2, 2, 2, 51, 244,
	3, 2, 2, 2, 53, 246, 3, 2, 2, 2, 55, 248, 3, 2, 2, 2, 57, 251, 3, 2, 2,
	2, 59, 270, 3, 2, 2, 2, 61, 273, 3, 2, 2, 2, 63, 277, 3, 2, 2, 2, 65, 283,
	3, 2, 2, 2, 67, 285, 3, 2, 2, 2, 69, 295, 3, 2, 2, 2, 71, 300, 3, 2, 2,
	2, 73, 306, 3, 2, 2, 2, 75, 316, 3, 2, 2, 2, 77, 318, 3, 2, 2, 2, 79, 324,
	3, 2, 2, 2, 81, 332, 3, 2, 2, 2, 83, 84, 7, 75, 2, 2, 84, 85, 7, 70, 2,
	2, 85, 86, 7, 90, 2, 2, 86, 87, 7, 48, 2, 2, 87, 88, 7, 69, 2, 2, 88, 89,
	7, 84, 2, 2, 89, 90, 7, 71, 2, 2, 90, 91, 7, 67, 2, 2, 91, 92, 7, 86, 2,
	2, 92, 93, 7, 71, 2, 2, 93, 4, 3, 2, 2, 2, 94, 95, 7, 85, 2, 2, 95, 96,
	7, 69, 2, 2, 96, 97, 7, 74, 2, 2, 97, 98, 7, 71, 2, 2, 98, 99, 7, 79, 2,
	2, 99, 100, 7, 67, 2, 2, 100, 6, 3, 2, 2, 2, 101, 102, 7, 75, 2, 2, 102,
	103, 7, 70, 2, 2, 103, 104, 7, 90, 2, 2, 104, 105, 7, 48, 2, 2, 105, 106,
	7, 70, 2, 2, 106, 107, 7, 71, 2, 2, 107, 108, 7, 85, 2, 2, 108, 109, 7,
	86, 2, 2, 109, 110, 7, 84, 2, 2, 110, 111, 7, 81, 2, 2, 111, 112, 7, 91,
	2, 2, 112, 8, 3, 2, 2, 2, 113, 114, 7, 75, 2, 2, 114, 115, 7, 70, 2, 2,
	115, 116, 7, 90, 2, 2, 116, 117, 7, 48, 2, 2, 117, 118, 7, 75, 2, 2, 118,
	119, 7, 80, 2, 2, 119, 120, 7, 85, 2, 2, 120, 121, 7, 71, 2, 2, 121, 122,
	7, 84, 2, 2, 122, 123, 7, 86, 2, 2, 123, 10, 3, 2, 2, 2, 124, 125, 7, 75,
	2, 2, 125, 126, 7, 70, 2, 2, 126, 127, 7, 90, 2, 2, 127, 128, 7, 48, 2,
	2, 128, 129, 7, 70, 2, 2, 129, 130, 7, 71, 2, 2, 130, 131, 7, 78, 2, 2,
	131, 12, 3, 2, 2, 2, 132, 133, 7, 75, 2, 2, 133, 134, 7, 70, 2, 2, 134,
	135, 7, 90, 2, 2, 135, 136, 7, 48, 2, 2, 136, 137, 7, 85, 2, 2, 137, 138,
	7, 71, 2, 2, 138, 139, 7, 78, 2, 2, 139, 140, 7, 71, 2, 2, 140, 141, 7,
	69, 2, 2, 141, 142, 7, 86, 2, 2, 142, 14, 3, 2, 2, 2, 143, 144, 7, 83,
	2, 2, 144, 145, 7, 87, 2, 2, 145, 146, 7, 71, 2, 2, 146, 147, 7, 84, 2,
	2, 147, 148, 7, 91, 2, 2, 148, 16, 3, 2, 2, 2, 149, 150, 7, 89, 2, 2, 150,
	151, 7, 74, 2, 2, 151, 152, 7, 71, 2, 2, 152, 153, 7, 84, 2, 2, 153, 154,
	7, 71, 2, 2, 154, 18, 3, 2, 2, 2, 155, 156, 7, 81, 2, 2, 156, 157, 7, 84,
	2, 2, 157, 158, 7, 70, 2, 2, 158, 159, 7, 71, 2, 2, 159, 160, 7, 84, 2,
	2, 160, 161, 7, 68, 2, 2, 161, 162, 7, 91, 2, 2, 162, 20, 3, 2, 2, 2, 163,
	164, 7, 78, 2, 2, 164, 165, 7, 75, 2, 2, 165, 166, 7, 79, 2, 2, 166, 167,
	7, 75, 2, 2, 167, 168, 7, 86, 2, 2, 168, 22, 3, 2, 2, 2, 169, 170, 7, 93,
	2, 2, 170, 24, 3, 2, 2, 2, 171, 172, 7, 46, 2, 2, 172, 26, 3, 2, 2, 2,
	173, 174, 7, 95, 2, 2, 174, 28, 3, 2, 2, 2, 175, 176, 7, 87, 2, 2, 176,
	177, 7, 75, 2, 2, 177, 178, 7, 80, 2, 2, 178, 179, 7, 86, 2, 2, 179, 180,
	7, 58, 2, 2, 180, 30, 3, 2, 2, 2, 181, 182, 7, 87, 2, 2, 182, 183, 7, 75,
	2, 2, 183, 184, 7, 80, 2, 2, 184, 185, 7, 86, 2, 2, 185, 186, 7, 51, 2,
	2, 186, 187, 7, 56, 2, 2, 187, 32, 3, 2, 2, 2, 188, 189, 7, 87, 2, 2, 189,
	190, 7, 75, 2, 2, 190, 191, 7, 80, 2, 2, 191, 192, 7, 86, 2, 2, 192, 193,
	7, 53, 2, 2, 193, 194, 7, 52, 2, 2, 194, 34, 3, 2, 2, 2, 195, 196, 7, 87,
	2, 2, 196, 197, 7, 75, 2, 2, 197, 198, 7, 80, 2, 2, 198, 199, 7, 86, 2,
	2, 199, 200, 7, 56, 2, 2, 200, 201, 7, 54, 2, 2, 201, 36, 3, 2, 2, 2, 202,
	203, 7, 72, 2, 2, 203, 204, 7, 78, 2, 2, 204, 205, 7, 81, 2, 2, 205, 206,
	7, 67, 2, 2, 206, 207, 7, 86, 2, 2, 207, 208, 7, 53, 2, 2, 208, 209, 7,
	52, 2, 2, 209, 38, 3, 2, 2, 2, 210, 211, 7, 72, 2, 2, 211, 212, 7, 78,
	2, 2, 212, 213, 7, 81, 2, 2, 213, 214, 7, 67, 2, 2, 214, 215, 7, 86, 2,
	2, 215, 216, 7, 56, 2, 2, 216, 217, 7, 54, 2, 2, 217, 40, 3, 2, 2, 2, 218,
	219, 7, 71, 2, 2, 219, 220, 7, 80, 2, 2, 220, 221, 7, 87, 2, 2, 221, 222,
	7, 79, 2, 2, 222, 42, 3, 2, 2, 2, 223, 224, 7, 85, 2, 2, 224, 225, 7, 86,
	2, 2, 225, 226, 7, 84, 2, 2, 226, 227, 7, 75, 2, 2, 227, 228, 7, 80, 2,
	2, 228, 229, 7, 73, 2, 2, 229, 44, 3, 2, 2, 2, 230, 231, 7, 75, 2, 2, 231,
	232, 7, 80, 2, 2, 232, 46, 3, 2, 2, 2, 233, 234, 7, 69, 2, 2, 234, 235,
	7, 81, 2, 2, 235, 236, 7, 80, 2, 2, 236, 237, 7, 86, 2, 2, 237, 238, 7,
	67, 2, 2, 238, 239, 7, 75, 2, 2, 239, 240, 7, 80, 2, 2, 240, 241, 7, 85,
	2, 2, 241, 48, 3, 2, 2, 2, 242, 243, 7, 62, 2, 2, 243, 50, 3, 2, 2, 2,
	244, 245, 7, 64, 2, 2, 245, 52, 3, 2, 2, 2, 246, 247, 7, 63, 2, 2, 247,
	54, 3, 2, 2, 2, 248, 249, 7, 62, 2, 2, 249, 250, 7, 63, 2, 2, 250, 56,
	3, 2, 2, 2, 251, 252, 7, 64, 2, 2, 252, 253, 7, 63, 2, 2, 253, 58, 3, 2,
	2, 2, 254, 255, 5, 61, 31, 2, 255, 257, 7, 48, 2, 2, 256, 258, 5, 61, 31,
	2, 257, 256, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 260, 3, 2, 2, 2, 259,
	261, 5, 63, 32, 2, 260, 259, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 271,
	3, 2, 2, 2, 262, 263, 5, 61, 31, 2, 263, 264, 5, 63, 32, 2, 264, 271, 3,
	2, 2, 2, 265, 266, 7, 48, 2, 2, 266, 268, 5, 61, 31, 2, 267, 269, 5, 63,
	32, 2, 268, 267, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 271, 3, 2, 2, 2,
	270, 254, 3, 2, 2, 2, 270, 262, 3, 2, 2, 2, 270, 265, 3, 2, 2, 2, 271,
	60, 3, 2, 2, 2, 272, 274, 5, 65, 33, 2, 273, 272, 3, 2, 2, 2, 274, 275,
	3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 62, 3, 2,
	2, 2, 277, 279, 9, 2, 2, 2, 278, 280, 9, 3, 2, 2, 279, 278, 3, 2, 2, 2,
	279, 280, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 282, 5, 61, 31, 2, 282,
	64, 3, 2, 2, 2, 283, 284, 9, 4, 2, 2, 284, 66, 3, 2, 2, 2, 285, 290, 7,
	36, 2, 2, 286, 289, 5, 69, 35, 2, 287, 289, 10, 5, 2, 2, 288, 286, 3, 2,
	2, 2, 288, 287, 3, 2, 2, 2, 289, 292, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2,
	290, 291, 3, 2, 2, 2, 291, 293, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 293,
	294, 7, 36, 2, 2, 294, 68, 3, 2, 2, 2, 295, 298, 7, 94, 2, 2, 296, 299,
	9, 6, 2, 2, 297, 299, 5, 71, 36, 2, 298, 296, 3, 2, 2, 2, 298, 297, 3,
	2, 2, 2, 299, 70, 3, 2, 2, 2, 300, 301, 7, 119, 2, 2, 301, 302, 5, 73,
	37, 2, 302, 303, 5, 73, 37, 2, 303, 304, 5, 73, 37, 2, 304, 305, 5, 73,
	37, 2, 305, 72, 3, 2, 2, 2, 306, 307, 9, 7, 2, 2, 307, 74, 3, 2, 2, 2,
	308, 317, 7, 50, 2, 2, 309, 313, 9, 8, 2, 2, 310, 312, 9, 4, 2, 2, 311,
	310, 3, 2, 2, 2, 312, 315, 3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 313, 314,
	3, 2, 2, 2, 314, 317, 3, 2, 2, 2, 315, 313, 3, 2, 2, 2, 316, 308, 3, 2,
	2, 2, 316, 309, 3, 2, 2, 2, 317, 76, 3, 2, 2, 2, 318, 320, 9, 2, 2, 2,
	319, 321, 9, 3, 2, 2, 320, 319, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321,
	322, 3, 2, 2, 2, 322, 323, 5, 75, 38, 2, 323, 78, 3, 2, 2, 2, 324, 328,
	9, 9, 2, 2, 325, 327, 9, 10, 2, 2, 326, 325, 3, 2, 2, 2, 327, 330, 3, 2,
	2, 2, 328, 326, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 80, 3, 2, 2, 2,
	330, 328, 3, 2, 2, 2, 331, 333, 9, 11, 2, 2, 332, 331, 3, 2, 2, 2, 333,
	334, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 336,
	3, 2, 2, 2, 336, 337, 8, 41, 2, 2, 337, 82, 3, 2, 2, 2, 17, 2, 257, 260,
	268, 270, 275, 279, 288, 290, 298, 313, 316, 320, 328, 334, 3, 8, 2, 2,
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
	"", "'IDX.CREATE'", "'SCHEMA'", "'IDX.DESTROY'", "'IDX.INSERT'", "'IDX.DEL'",
	"'IDX.SELECT'", "'QUERY'", "'WHERE'", "'ORDERBY'", "'LIMIT'", "'['", "','",
	"']'", "'UINT8'", "'UINT16'", "'UINT32'", "'UINT64'", "'FLOAT32'", "'FLOAT64'",
	"'ENUM'", "'STRING'", "'IN'", "'CONTAINS'", "'<'", "'>'", "'='", "'<='",
	"'>='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "K_UINT8", "K_UINT16",
	"K_UINT32", "K_UINT64", "K_FLOAT32", "K_FLOAT64", "K_ENUM", "K_STRING",
	"K_IN", "K_CONTAINS", "K_LT", "K_BT", "K_EQ", "K_LE", "K_BE", "FLOAT_LIT",
	"STRING", "INT", "IDENTIFIER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "K_UINT8", "K_UINT16", "K_UINT32", "K_UINT64",
	"K_FLOAT32", "K_FLOAT64", "K_ENUM", "K_STRING", "K_IN", "K_CONTAINS", "K_LT",
	"K_BT", "K_EQ", "K_LE", "K_BE", "FLOAT_LIT", "DECIMALS", "EXPONENT", "DECIMAL_DIGIT",
	"STRING", "ESC", "UNICODE", "HEX", "INT", "EXP", "IDENTIFIER", "WS",
}

type CQLLexer struct {
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

func NewCQLLexer(input antlr.CharStream) *CQLLexer {

	l := new(CQLLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "CQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CQLLexer tokens.
const (
	CQLLexerT__0       = 1
	CQLLexerT__1       = 2
	CQLLexerT__2       = 3
	CQLLexerT__3       = 4
	CQLLexerT__4       = 5
	CQLLexerT__5       = 6
	CQLLexerT__6       = 7
	CQLLexerT__7       = 8
	CQLLexerT__8       = 9
	CQLLexerT__9       = 10
	CQLLexerT__10      = 11
	CQLLexerT__11      = 12
	CQLLexerT__12      = 13
	CQLLexerK_UINT8    = 14
	CQLLexerK_UINT16   = 15
	CQLLexerK_UINT32   = 16
	CQLLexerK_UINT64   = 17
	CQLLexerK_FLOAT32  = 18
	CQLLexerK_FLOAT64  = 19
	CQLLexerK_ENUM     = 20
	CQLLexerK_STRING   = 21
	CQLLexerK_IN       = 22
	CQLLexerK_CONTAINS = 23
	CQLLexerK_LT       = 24
	CQLLexerK_BT       = 25
	CQLLexerK_EQ       = 26
	CQLLexerK_LE       = 27
	CQLLexerK_BE       = 28
	CQLLexerFLOAT_LIT  = 29
	CQLLexerSTRING     = 30
	CQLLexerINT        = 31
	CQLLexerIDENTIFIER = 32
	CQLLexerWS         = 33
)
