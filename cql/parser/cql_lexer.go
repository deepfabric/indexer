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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 31, 271,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 7, 27, 222,
	10, 27, 12, 27, 14, 27, 225, 11, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28,
	5, 28, 232, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 31, 7, 31, 245, 10, 31, 12, 31, 14, 31, 248, 11, 31,
	5, 31, 250, 10, 31, 3, 32, 3, 32, 5, 32, 254, 10, 32, 3, 32, 3, 32, 3,
	33, 3, 33, 7, 33, 260, 10, 33, 12, 33, 14, 33, 263, 11, 33, 3, 34, 6, 34,
	266, 10, 34, 13, 34, 14, 34, 267, 3, 34, 3, 34, 2, 2, 35, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 2, 57, 2, 59, 2, 61, 29, 63, 2,
	65, 30, 67, 31, 3, 2, 12, 4, 2, 36, 36, 94, 94, 10, 2, 36, 36, 49, 49,
	94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2, 50, 59,
	67, 72, 99, 104, 3, 2, 51, 59, 3, 2, 50, 59, 4, 2, 71, 71, 103, 103, 4,
	2, 45, 45, 47, 47, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92,
	97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 274, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43,
	3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2,
	51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2,
	2, 67, 3, 2, 2, 2, 3, 69, 3, 2, 2, 2, 5, 80, 3, 2, 2, 2, 7, 87, 3, 2, 2,
	2, 9, 99, 3, 2, 2, 2, 11, 110, 3, 2, 2, 2, 13, 118, 3, 2, 2, 2, 15, 129,
	3, 2, 2, 2, 17, 135, 3, 2, 2, 2, 19, 143, 3, 2, 2, 2, 21, 149, 3, 2, 2,
	2, 23, 151, 3, 2, 2, 2, 25, 153, 3, 2, 2, 2, 27, 155, 3, 2, 2, 2, 29, 161,
	3, 2, 2, 2, 31, 168, 3, 2, 2, 2, 33, 175, 3, 2, 2, 2, 35, 182, 3, 2, 2,
	2, 37, 187, 3, 2, 2, 2, 39, 194, 3, 2, 2, 2, 41, 197, 3, 2, 2, 2, 43, 206,
	3, 2, 2, 2, 45, 208, 3, 2, 2, 2, 47, 210, 3, 2, 2, 2, 49, 212, 3, 2, 2,
	2, 51, 215, 3, 2, 2, 2, 53, 218, 3, 2, 2, 2, 55, 228, 3, 2, 2, 2, 57, 233,
	3, 2, 2, 2, 59, 239, 3, 2, 2, 2, 61, 249, 3, 2, 2, 2, 63, 251, 3, 2, 2,
	2, 65, 257, 3, 2, 2, 2, 67, 265, 3, 2, 2, 2, 69, 70, 7, 75, 2, 2, 70, 71,
	7, 70, 2, 2, 71, 72, 7, 90, 2, 2, 72, 73, 7, 48, 2, 2, 73, 74, 7, 69, 2,
	2, 74, 75, 7, 84, 2, 2, 75, 76, 7, 71, 2, 2, 76, 77, 7, 67, 2, 2, 77, 78,
	7, 86, 2, 2, 78, 79, 7, 71, 2, 2, 79, 4, 3, 2, 2, 2, 80, 81, 7, 85, 2,
	2, 81, 82, 7, 69, 2, 2, 82, 83, 7, 74, 2, 2, 83, 84, 7, 71, 2, 2, 84, 85,
	7, 79, 2, 2, 85, 86, 7, 67, 2, 2, 86, 6, 3, 2, 2, 2, 87, 88, 7, 75, 2,
	2, 88, 89, 7, 70, 2, 2, 89, 90, 7, 90, 2, 2, 90, 91, 7, 48, 2, 2, 91, 92,
	7, 70, 2, 2, 92, 93, 7, 71, 2, 2, 93, 94, 7, 85, 2, 2, 94, 95, 7, 86, 2,
	2, 95, 96, 7, 84, 2, 2, 96, 97, 7, 81, 2, 2, 97, 98, 7, 91, 2, 2, 98, 8,
	3, 2, 2, 2, 99, 100, 7, 75, 2, 2, 100, 101, 7, 70, 2, 2, 101, 102, 7, 90,
	2, 2, 102, 103, 7, 48, 2, 2, 103, 104, 7, 75, 2, 2, 104, 105, 7, 80, 2,
	2, 105, 106, 7, 85, 2, 2, 106, 107, 7, 71, 2, 2, 107, 108, 7, 84, 2, 2,
	108, 109, 7, 86, 2, 2, 109, 10, 3, 2, 2, 2, 110, 111, 7, 75, 2, 2, 111,
	112, 7, 70, 2, 2, 112, 113, 7, 90, 2, 2, 113, 114, 7, 48, 2, 2, 114, 115,
	7, 70, 2, 2, 115, 116, 7, 71, 2, 2, 116, 117, 7, 78, 2, 2, 117, 12, 3,
	2, 2, 2, 118, 119, 7, 75, 2, 2, 119, 120, 7, 70, 2, 2, 120, 121, 7, 90,
	2, 2, 121, 122, 7, 48, 2, 2, 122, 123, 7, 85, 2, 2, 123, 124, 7, 71, 2,
	2, 124, 125, 7, 78, 2, 2, 125, 126, 7, 71, 2, 2, 126, 127, 7, 69, 2, 2,
	127, 128, 7, 86, 2, 2, 128, 14, 3, 2, 2, 2, 129, 130, 7, 89, 2, 2, 130,
	131, 7, 74, 2, 2, 131, 132, 7, 71, 2, 2, 132, 133, 7, 84, 2, 2, 133, 134,
	7, 71, 2, 2, 134, 16, 3, 2, 2, 2, 135, 136, 7, 81, 2, 2, 136, 137, 7, 84,
	2, 2, 137, 138, 7, 70, 2, 2, 138, 139, 7, 71, 2, 2, 139, 140, 7, 84, 2,
	2, 140, 141, 7, 68, 2, 2, 141, 142, 7, 91, 2, 2, 142, 18, 3, 2, 2, 2, 143,
	144, 7, 78, 2, 2, 144, 145, 7, 75, 2, 2, 145, 146, 7, 79, 2, 2, 146, 147,
	7, 75, 2, 2, 147, 148, 7, 86, 2, 2, 148, 20, 3, 2, 2, 2, 149, 150, 7, 93,
	2, 2, 150, 22, 3, 2, 2, 2, 151, 152, 7, 46, 2, 2, 152, 24, 3, 2, 2, 2,
	153, 154, 7, 95, 2, 2, 154, 26, 3, 2, 2, 2, 155, 156, 7, 87, 2, 2, 156,
	157, 7, 75, 2, 2, 157, 158, 7, 80, 2, 2, 158, 159, 7, 86, 2, 2, 159, 160,
	7, 58, 2, 2, 160, 28, 3, 2, 2, 2, 161, 162, 7, 87, 2, 2, 162, 163, 7, 75,
	2, 2, 163, 164, 7, 80, 2, 2, 164, 165, 7, 86, 2, 2, 165, 166, 7, 51, 2,
	2, 166, 167, 7, 56, 2, 2, 167, 30, 3, 2, 2, 2, 168, 169, 7, 87, 2, 2, 169,
	170, 7, 75, 2, 2, 170, 171, 7, 80, 2, 2, 171, 172, 7, 86, 2, 2, 172, 173,
	7, 53, 2, 2, 173, 174, 7, 52, 2, 2, 174, 32, 3, 2, 2, 2, 175, 176, 7, 87,
	2, 2, 176, 177, 7, 75, 2, 2, 177, 178, 7, 80, 2, 2, 178, 179, 7, 86, 2,
	2, 179, 180, 7, 56, 2, 2, 180, 181, 7, 54, 2, 2, 181, 34, 3, 2, 2, 2, 182,
	183, 7, 71, 2, 2, 183, 184, 7, 80, 2, 2, 184, 185, 7, 87, 2, 2, 185, 186,
	7, 79, 2, 2, 186, 36, 3, 2, 2, 2, 187, 188, 7, 85, 2, 2, 188, 189, 7, 86,
	2, 2, 189, 190, 7, 84, 2, 2, 190, 191, 7, 75, 2, 2, 191, 192, 7, 80, 2,
	2, 192, 193, 7, 73, 2, 2, 193, 38, 3, 2, 2, 2, 194, 195, 7, 75, 2, 2, 195,
	196, 7, 80, 2, 2, 196, 40, 3, 2, 2, 2, 197, 198, 7, 69, 2, 2, 198, 199,
	7, 81, 2, 2, 199, 200, 7, 80, 2, 2, 200, 201, 7, 86, 2, 2, 201, 202, 7,
	67, 2, 2, 202, 203, 7, 75, 2, 2, 203, 204, 7, 80, 2, 2, 204, 205, 7, 85,
	2, 2, 205, 42, 3, 2, 2, 2, 206, 207, 7, 62, 2, 2, 207, 44, 3, 2, 2, 2,
	208, 209, 7, 64, 2, 2, 209, 46, 3, 2, 2, 2, 210, 211, 7, 63, 2, 2, 211,
	48, 3, 2, 2, 2, 212, 213, 7, 62, 2, 2, 213, 214, 7, 63, 2, 2, 214, 50,
	3, 2, 2, 2, 215, 216, 7, 64, 2, 2, 216, 217, 7, 63, 2, 2, 217, 52, 3, 2,
	2, 2, 218, 223, 7, 36, 2, 2, 219, 222, 5, 55, 28, 2, 220, 222, 10, 2, 2,
	2, 221, 219, 3, 2, 2, 2, 221, 220, 3, 2, 2, 2, 222, 225, 3, 2, 2, 2, 223,
	221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 226, 3, 2, 2, 2, 225, 223,
	3, 2, 2, 2, 226, 227, 7, 36, 2, 2, 227, 54, 3, 2, 2, 2, 228, 231, 7, 94,
	2, 2, 229, 232, 9, 3, 2, 2, 230, 232, 5, 57, 29, 2, 231, 229, 3, 2, 2,
	2, 231, 230, 3, 2, 2, 2, 232, 56, 3, 2, 2, 2, 233, 234, 7, 119, 2, 2, 234,
	235, 5, 59, 30, 2, 235, 236, 5, 59, 30, 2, 236, 237, 5, 59, 30, 2, 237,
	238, 5, 59, 30, 2, 238, 58, 3, 2, 2, 2, 239, 240, 9, 4, 2, 2, 240, 60,
	3, 2, 2, 2, 241, 250, 7, 50, 2, 2, 242, 246, 9, 5, 2, 2, 243, 245, 9, 6,
	2, 2, 244, 243, 3, 2, 2, 2, 245, 248, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2,
	246, 247, 3, 2, 2, 2, 247, 250, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 249,
	241, 3, 2, 2, 2, 249, 242, 3, 2, 2, 2, 250, 62, 3, 2, 2, 2, 251, 253, 9,
	7, 2, 2, 252, 254, 9, 8, 2, 2, 253, 252, 3, 2, 2, 2, 253, 254, 3, 2, 2,
	2, 254, 255, 3, 2, 2, 2, 255, 256, 5, 61, 31, 2, 256, 64, 3, 2, 2, 2, 257,
	261, 9, 9, 2, 2, 258, 260, 9, 10, 2, 2, 259, 258, 3, 2, 2, 2, 260, 263,
	3, 2, 2, 2, 261, 259, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 66, 3, 2,
	2, 2, 263, 261, 3, 2, 2, 2, 264, 266, 9, 11, 2, 2, 265, 264, 3, 2, 2, 2,
	266, 267, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268,
	269, 3, 2, 2, 2, 269, 270, 8, 34, 2, 2, 270, 68, 3, 2, 2, 2, 11, 2, 221,
	223, 231, 246, 249, 253, 261, 267, 3, 8, 2, 2,
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
	"'IDX.SELECT'", "'WHERE'", "'ORDERBY'", "'LIMIT'", "'['", "','", "']'",
	"'UINT8'", "'UINT16'", "'UINT32'", "'UINT64'", "'ENUM'", "'STRING'", "'IN'",
	"'CONTAINS'", "'<'", "'>'", "'='", "'<='", "'>='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "K_UINT8", "K_UINT16",
	"K_UINT32", "K_UINT64", "K_ENUM", "K_STRING", "K_IN", "K_CONTAINS", "K_LT",
	"K_BT", "K_EQ", "K_LE", "K_BE", "STRING", "INT", "IDENTIFIER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "K_UINT8", "K_UINT16", "K_UINT32", "K_UINT64",
	"K_ENUM", "K_STRING", "K_IN", "K_CONTAINS", "K_LT", "K_BT", "K_EQ", "K_LE",
	"K_BE", "STRING", "ESC", "UNICODE", "HEX", "INT", "EXP", "IDENTIFIER",
	"WS",
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
	CQLLexerK_UINT8    = 13
	CQLLexerK_UINT16   = 14
	CQLLexerK_UINT32   = 15
	CQLLexerK_UINT64   = 16
	CQLLexerK_ENUM     = 17
	CQLLexerK_STRING   = 18
	CQLLexerK_IN       = 19
	CQLLexerK_CONTAINS = 20
	CQLLexerK_LT       = 21
	CQLLexerK_BT       = 22
	CQLLexerK_EQ       = 23
	CQLLexerK_LE       = 24
	CQLLexerK_BE       = 25
	CQLLexerSTRING     = 26
	CQLLexerINT        = 27
	CQLLexerIDENTIFIER = 28
	CQLLexerWS         = 29
)
