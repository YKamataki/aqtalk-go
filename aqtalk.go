package aqtalkgo

import (
	"C"
)

type PARAM struct {
	// 基本素片
	Base VoiceBase
	// 話速
	// 50 <= Speed <= 300 default:100
	Speed int
	// 音量
	// 0 <= Vol <= 300 default: 300
	Vol int
	// 高さ
	// 20 <= Pitch <= 200 default:(Baseに依存)
	Pitch int
	// アクセント
	// 0 <= Accent <= 200 default:(Baseに依存)
	Accent int
	// 音程1
	// 0 <= Lmd <= 200 default:100
	Lmd int
	// 音程2(サンプリング周波数)
	// 50 <= Fsc <= 200 default: 100
	Fsc int
}

// 基本素片
type VoiceBase int

const (
	VoiceBaseF1E VoiceBase = iota
	VoiceBaseF2E VoiceBase = iota
	VoiceBaseM1E VoiceBase = iota
)

// プリセットの声質
var (
	// 女性1
	PresetVoiceF1 PARAM = PARAM{
		VoiceBaseF1E,
		100,
		100,
		100,
		100,
		100,
		100,
	}
	// 女性2
	PresetVoiceF2 PARAM = PARAM{
		VoiceBaseF2E,
		100,
		100,
		77,
		150,
		100,
		100,
	}
	// 女性3
	PresetVoiceF3 PARAM = PARAM{
		VoiceBaseF1E,
		80,
		100,
		100,
		100,
		61,
		148,
	}
	// 男性1
	PresetVoiceM1 PARAM = PARAM{
		VoiceBaseF1E,
		80,
		100,
		100,
		100,
		61,
		148,
	}
	// 男性2
	PresetVoiceM2 PARAM = PARAM{
		VoiceBaseM1E,
		100,
		100,
		30,
		100,
		100,
		100,
	}
	// ロボット1
	PresetVoiceR1 PARAM = PARAM{
		VoiceBaseM1E,
		105,
		100,
		45,
		130,
		120,
		100,
	}
	// ロボット2
	PresetVoiceR2 PARAM = PARAM{
		VoiceBaseF2E,
		70,
		100,
		50,
		50,
		50,
		180,
	}
)
