package Misc

import (
	"encoding/binary"
	"math"
)

func bytes2float32(b []byte) float32 {
	bits := binary.LittleEndian.Uint32(b)
	result := math.Float32frombits(bits)
	return result
}

func float32Tobytes(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, bits)
	return b
}

func bytes2float64(b []byte) float64 {
	bits := binary.LittleEndian.Uint64(b)
	result := math.Float64frombits(bits)
	return result
}

func float64Tobytes(f float64) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, bits)
	return b
}
