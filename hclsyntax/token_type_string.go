// Code generated by "stringer -type TokenType -output token_type_string.go"; DO NOT EDIT.

package hclsyntax

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TokenOBrace-123]
	_ = x[TokenCBrace-125]
	_ = x[TokenOBrack-91]
	_ = x[TokenCBrack-93]
	_ = x[TokenOParen-40]
	_ = x[TokenCParen-41]
	_ = x[TokenOQuote-171]
	_ = x[TokenCQuote-187]
	_ = x[TokenOHeredoc-72]
	_ = x[TokenCHeredoc-104]
	_ = x[TokenStar-42]
	_ = x[TokenSlash-47]
	_ = x[TokenPlus-43]
	_ = x[TokenMinus-45]
	_ = x[TokenPercent-37]
	_ = x[TokenEqual-61]
	_ = x[TokenColonOp-8788]
	_ = x[TokenNotEqual-8800]
	_ = x[TokenLessThan-60]
	_ = x[TokenLessThanEq-8804]
	_ = x[TokenGreaterThan-62]
	_ = x[TokenGreaterThanEq-8805]
	_ = x[TokenAnd-8743]
	_ = x[TokenOr-8744]
	_ = x[TokenBang-33]
	_ = x[TokenDot-46]
	_ = x[TokenComma-44]
	_ = x[TokenEllipsis-8230]
	_ = x[TokenFatArrow-8658]
	_ = x[TokenQuestion-63]
	_ = x[TokenColon-58]
	_ = x[TokenTemplateInterp-8747]
	_ = x[TokenTemplateControl-955]
	_ = x[TokenTemplateSeqEnd-8718]
	_ = x[TokenQuotedLit-81]
	_ = x[TokenStringLit-83]
	_ = x[TokenNumberLit-78]
	_ = x[TokenIdent-73]
	_ = x[TokenComment-67]
	_ = x[TokenNewline-10]
	_ = x[TokenEOF-9220]
	_ = x[TokenBitwiseAnd-38]
	_ = x[TokenBitwiseOr-124]
	_ = x[TokenBitwiseNot-126]
	_ = x[TokenBitwiseXor-94]
	_ = x[TokenStarStar-10138]
	_ = x[TokenApostrophe-39]
	_ = x[TokenBacktick-96]
	_ = x[TokenSemicolon-59]
	_ = x[TokenTabs-9225]
	_ = x[TokenInvalid-65533]
	_ = x[TokenBadUTF8-128169]
	_ = x[TokenQuotedNewline-9252]
	_ = x[TokenNil-0]
}

const _TokenType_name = "TokenNilTokenNewlineTokenBangTokenPercentTokenBitwiseAndTokenApostropheTokenOParenTokenCParenTokenStarTokenPlusTokenCommaTokenMinusTokenDotTokenSlashTokenColonTokenSemicolonTokenLessThanTokenColonTokenGreaterThanTokenQuestionTokenCommentTokenOHeredocTokenIdentTokenNumberLitTokenQuotedLitTokenStringLitTokenOBrackTokenCBrackTokenBitwiseXorTokenBacktickTokenCHeredocTokenOBraceTokenBitwiseOrTokenCBraceTokenBitwiseNotTokenOQuoteTokenCQuoteTokenTemplateControlTokenEllipsisTokenFatArrowTokenTemplateSeqEndTokenAndTokenOrTokenTemplateInterpTokenColonOpTokenNotEqualTokenLessThanEqTokenGreaterThanEqTokenEOFTokenTabsTokenQuotedNewlineTokenStarStarTokenInvalidTokenBadUTF8"

var _TokenType_map = map[TokenType]string{
	0:      _TokenType_name[0:8],
	10:     _TokenType_name[8:20],
	33:     _TokenType_name[20:29],
	37:     _TokenType_name[29:41],
	38:     _TokenType_name[41:56],
	39:     _TokenType_name[56:71],
	40:     _TokenType_name[71:82],
	41:     _TokenType_name[82:93],
	42:     _TokenType_name[93:102],
	43:     _TokenType_name[102:111],
	44:     _TokenType_name[111:121],
	45:     _TokenType_name[121:131],
	46:     _TokenType_name[131:139],
	47:     _TokenType_name[139:149],
	58:     _TokenType_name[149:159],
	59:     _TokenType_name[159:173],
	60:     _TokenType_name[173:186],
	61:     _TokenType_name[186:196],
	62:     _TokenType_name[196:212],
	63:     _TokenType_name[212:225],
	67:     _TokenType_name[225:237],
	72:     _TokenType_name[237:250],
	73:     _TokenType_name[250:260],
	78:     _TokenType_name[260:274],
	81:     _TokenType_name[274:288],
	83:     _TokenType_name[288:302],
	91:     _TokenType_name[302:313],
	93:     _TokenType_name[313:324],
	94:     _TokenType_name[324:339],
	96:     _TokenType_name[339:352],
	104:    _TokenType_name[352:365],
	123:    _TokenType_name[365:376],
	124:    _TokenType_name[376:390],
	125:    _TokenType_name[390:401],
	126:    _TokenType_name[401:416],
	171:    _TokenType_name[416:427],
	187:    _TokenType_name[427:438],
	955:    _TokenType_name[438:458],
	8230:   _TokenType_name[458:471],
	8658:   _TokenType_name[471:484],
	8718:   _TokenType_name[484:503],
	8743:   _TokenType_name[503:511],
	8744:   _TokenType_name[511:518],
	8747:   _TokenType_name[518:537],
	8788:   _TokenType_name[537:549],
	8800:   _TokenType_name[549:562],
	8804:   _TokenType_name[562:577],
	8805:   _TokenType_name[577:595],
	9220:   _TokenType_name[595:603],
	9225:   _TokenType_name[603:612],
	9252:   _TokenType_name[612:630],
	10138:  _TokenType_name[630:643],
	65533:  _TokenType_name[643:655],
	128169: _TokenType_name[655:667],
}

func (i TokenType) String() string {
	if str, ok := _TokenType_map[i]; ok {
		return str
	}
	return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
}
