package heap

const (
	AccPublic       = 0x0001 //      CcFM_____
	AccPrivate      = 0x0002 //      _cFM_____
	AccProtected    = 0x0004 //      _cFM_____
	AccStatic       = 0x0008 //      _cFM_____
	AccFinal        = 0x0010 //      CcFMP____
	AccSuper        = 0x0020 //      C________
	AccSynchronized = 0x0020 //      ___M_____
	AccOpen         = 0x0020 // 9,   _____D___
	AccTransitive   = 0x0020 // 9,   ______R__
	AccVolatile     = 0x0040 //      __F______
	AccBridge       = 0x0040 //      ___M_____
	AccStaticPhase  = 0x0040 // 9,   ______R__
	AccTransient    = 0x0080 //      __F______
	AccVarargs      = 0x0080 // 5.0  ___M_____
	AccNative       = 0x0100 //      ___M_____
	AccInterface    = 0x0200 //      Cc_______
	AccAbstract     = 0x0400 //      Cc_M_____
	AccStrict       = 0x0800 //      ___M_____
	AccSynthetic    = 0x1000 //      CcFMPDRXO
	AccAnnotation   = 0x2000 // 5.0, Cc_______
	AccEnum         = 0x4000 // 5.0, CcF______
	AccModule       = 0x8000 // 9,   C________
	AccMandated     = 0x8000 // ?,   ____PDRXO
)
