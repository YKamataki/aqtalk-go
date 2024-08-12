package aqtalkgo

import (
	// #cgo LDFLAGS: -llibAquesTalk10
	// #include "AquesTalk.h"
	// typedef struct _AQTK_PARAM_ VOICESPEC;
	"C"
)

import (
	"fmt"
	"unsafe"
)

// TODO: Convert between C struct and Go struct
/*
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
*/

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
	PresetVoiceF1 = C.gVoice_F1
	// 女性2
	PresetVoiceF2 = C.gVoice_F2
	// 女性3
	PresetVoiceF3 = C.gVoice_F3
	// 男性1
	PresetVoiceM1 = C.gVoice_M1
	// 男性2
	PresetVoiceM2 = C.gVoice_M2
	// ロボット1
	PresetVoiceR1 = C.gVoice_R1
	// ロボット2
	PresetVoiceR2 = C.gVoice_R2
)

// 開発者ライセンスキーを設定する
// 不正なキーでもエラーとならないことがあるが、制限は解除されない
func SetDevKey(key string) error {
	devKey := C.CString(key)
	isOK := C.AquesTalk_SetDevKey(devKey)
	if isOK == 0 {
		return nil
	} else {
		return fmt.Errorf("invalid development key")
	}
}

// 利用者ライセンスキーを設定する
// 不正なキーでもまれにエラーとならないことがあるが、ライセンス無しのままとなる
func SetUsrKey(key string) error {
	usrKey := C.CString(key)
	isOK := C.AquesTalk_SetUsrKey(usrKey)
	if isOK == 0 {
		return nil
	} else {
		return fmt.Errorf("invalid user key")
	}
}

// 音声を合成する
// 音声記号列はUTF-8でエンコードされている必要がある
// 戻り値: WAVフォーマットの音声データ
func Synthe(pparam *C.VOICESPEC, koe string) ([]byte, error) {
	size := C.int(0)
	wav := C.AquesTalk_Synthe_Utf8(pparam, C.CString(koe), &size)
	if size == 0 {
		return nil, fmt.Errorf("failed to synthe")
	}
	ptr := unsafe.Pointer(wav)
	wavgo := (*[]byte)(ptr)

	// free
	defer C.AquesTalk_FreeWave(wav)

	return *wavgo, nil
}
