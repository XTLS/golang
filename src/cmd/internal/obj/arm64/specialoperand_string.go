// Code generated by "stringer -type SpecialOperand -trimprefix SPOP_"; DO NOT EDIT.

package arm64

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SPOP_PLDL1KEEP-0]
	_ = x[SPOP_BEGIN-0]
	_ = x[SPOP_PLDL1STRM-1]
	_ = x[SPOP_PLDL2KEEP-2]
	_ = x[SPOP_PLDL2STRM-3]
	_ = x[SPOP_PLDL3KEEP-4]
	_ = x[SPOP_PLDL3STRM-5]
	_ = x[SPOP_PLIL1KEEP-6]
	_ = x[SPOP_PLIL1STRM-7]
	_ = x[SPOP_PLIL2KEEP-8]
	_ = x[SPOP_PLIL2STRM-9]
	_ = x[SPOP_PLIL3KEEP-10]
	_ = x[SPOP_PLIL3STRM-11]
	_ = x[SPOP_PSTL1KEEP-12]
	_ = x[SPOP_PSTL1STRM-13]
	_ = x[SPOP_PSTL2KEEP-14]
	_ = x[SPOP_PSTL2STRM-15]
	_ = x[SPOP_PSTL3KEEP-16]
	_ = x[SPOP_PSTL3STRM-17]
	_ = x[SPOP_DAIFSet-18]
	_ = x[SPOP_DAIFClr-19]
	_ = x[SPOP_EQ-20]
	_ = x[SPOP_NE-21]
	_ = x[SPOP_HS-22]
	_ = x[SPOP_LO-23]
	_ = x[SPOP_MI-24]
	_ = x[SPOP_PL-25]
	_ = x[SPOP_VS-26]
	_ = x[SPOP_VC-27]
	_ = x[SPOP_HI-28]
	_ = x[SPOP_LS-29]
	_ = x[SPOP_GE-30]
	_ = x[SPOP_LT-31]
	_ = x[SPOP_GT-32]
	_ = x[SPOP_LE-33]
	_ = x[SPOP_AL-34]
	_ = x[SPOP_NV-35]
	_ = x[SPOP_END-36]
}

const _SpecialOperand_name = "PLDL1KEEPPLDL1STRMPLDL2KEEPPLDL2STRMPLDL3KEEPPLDL3STRMPLIL1KEEPPLIL1STRMPLIL2KEEPPLIL2STRMPLIL3KEEPPLIL3STRMPSTL1KEEPPSTL1STRMPSTL2KEEPPSTL2STRMPSTL3KEEPPSTL3STRMDAIFSetDAIFClrEQNEHSLOMIPLVSVCHILSGELTGTLEALNVEND"

var _SpecialOperand_index = [...]uint8{0, 9, 18, 27, 36, 45, 54, 63, 72, 81, 90, 99, 108, 117, 126, 135, 144, 153, 162, 169, 176, 178, 180, 182, 184, 186, 188, 190, 192, 194, 196, 198, 200, 202, 204, 206, 208, 211}

func (i SpecialOperand) String() string {
	if i < 0 || i >= SpecialOperand(len(_SpecialOperand_index)-1) {
		return "SpecialOperand(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SpecialOperand_name[_SpecialOperand_index[i]:_SpecialOperand_index[i+1]]
}
