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
