// Diode Network Client
// Copyright 2019-2021 IoT Blockchain Technology Corporation LLC (IBTC)
// Licensed under the Diode License, Version 1.0
package edge

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/diodechain/diode_go_client/util"
)

func TestValidateMerkleTree(t *testing.T) {
	// 9 bit test
	testRoots := &AccountRoots{AccountRoots: [][]uint8{{0x1a, 0xfd, 0xf0, 0x7c, 0x95, 0x2c, 0x83, 0x5b, 0x80, 0x55, 0x5c, 0xc0, 0x24, 0xd6, 0x2d, 0x25, 0xe5, 0xe8, 0x6e, 0x45, 0x4f, 0x52, 0x72, 0xea, 0x4e, 0xee, 0x11, 0x2f, 0x1d, 0xd8, 0xf5, 0x25}, []uint8{0xa4, 0x28, 0xc0, 0x5, 0x77, 0x64, 0x4, 0x31, 0xd1, 0x51, 0x74, 0xf3, 0xcd, 0x8a, 0x40, 0x22, 0xed, 0x47, 0x6c, 0xb3, 0x72, 0x70, 0xfb, 0x22, 0x95, 0x56, 0x77, 0xe7, 0x1, 0x34, 0x88, 0x34}, []uint8{0x5e, 0xb8, 0x98, 0xb, 0x16, 0x6e, 0x93, 0x24, 0x78, 0xd0, 0xdf, 0xb5, 0x92, 0xfb, 0x7, 0xa0, 0x4f, 0xb2, 0x70, 0xf8, 0xcc, 0x82, 0xde, 0x1c, 0x8a, 0x54, 0x26, 0x31, 0x5a, 0x25, 0x28, 0xf8}, []uint8{0x50, 0xdc, 0x3d, 0x15, 0x60, 0xee, 0xd7, 0x24, 0x63, 0x82, 0x4d, 0xd7, 0x73, 0x2c, 0xfa, 0x99, 0xcc, 0x27, 0x5b, 0xa7, 0xb, 0xd4, 0x2a, 0x43, 0xf2, 0xbf, 0xf4, 0x2a, 0xf5, 0xb, 0x5b, 0xa9}, []uint8{0x22, 0xf1, 0x33, 0xd3, 0xc3, 0x6d, 0x6, 0xbe, 0xeb, 0x8c, 0xf7, 0x2b, 0x70, 0x72, 0x3f, 0xd7, 0x58, 0x86, 0xe1, 0xdb, 0xb2, 0x26, 0xba, 0x74, 0x5a, 0x40, 0x41, 0x4e, 0x54, 0x55, 0x45, 0xaf}, []uint8{0xd6, 0x9a, 0xd1, 0x2b, 0x1c, 0x5d, 0x7e, 0x1f, 0x6b, 0x44, 0xe5, 0xb9, 0x26, 0x8f, 0x1, 0xde, 0x5a, 0xe, 0xf, 0x31, 0xc0, 0xf6, 0x41, 0x59, 0xa0, 0x50, 0x5b, 0x9f, 0x42, 0x55, 0x58, 0x4}, []uint8{0x7e, 0xad, 0x2d, 0x77, 0x92, 0xe0, 0xac, 0x98, 0x34, 0x89, 0x3b, 0xb1, 0xf6, 0xdf, 0xe6, 0xbf, 0x8b, 0x46, 0xe0, 0xf4, 0x6d, 0x86, 0x93, 0x9f, 0x72, 0x5, 0xae, 0xcb, 0xca, 0xd5, 0x21, 0x43}, []uint8{0x1b, 0xd3, 0xea, 0xe7, 0x27, 0x8e, 0xbc, 0xb3, 0x83, 0x3c, 0xd0, 0x9d, 0x14, 0xee, 0x11, 0xb, 0x66, 0x91, 0xb, 0x3d, 0xb7, 0x49, 0x9b, 0x95, 0xa8, 0x32, 0x71, 0x28, 0x5a, 0xa4, 0x5d, 0xa5}, []uint8{0x99, 0xf8, 0x9d, 0xb3, 0xc6, 0xd7, 0xfe, 0x63, 0x8b, 0x9d, 0xa0, 0x4e, 0x32, 0x6, 0x7b, 0x4c, 0x5e, 0x87, 0x26, 0xb6, 0xec, 0x2c, 0x35, 0x3a, 0xf3, 0x98, 0x70, 0x89, 0xd2, 0x78, 0x1d, 0x43}, []uint8{0x79, 0xd1, 0x22, 0x3e, 0xf9, 0xc5, 0xd7, 0xfe, 0x88, 0x58, 0xc6, 0xda, 0xf5, 0xa0, 0xec, 0xf6, 0x22, 0x47, 0x4c, 0x86, 0x28, 0x7c, 0x58, 0xf9, 0x4e, 0x26, 0xdf, 0xee, 0x58, 0x3d, 0x37, 0x16}, []uint8{0xff, 0xad, 0x1, 0xed, 0x68, 0xd7, 0xa6, 0x9e, 0x53, 0xc1, 0xe4, 0x5c, 0x7c, 0xc5, 0xbd, 0xaf, 0x5a, 0x2, 0xf1, 0xb4, 0xe2, 0x5c, 0xac, 0xcd, 0x2b, 0x3c, 0x28, 0xc, 0x2a, 0xaf, 0x28, 0x16}, []uint8{0x18, 0xbc, 0xdf, 0xaa, 0xf1, 0x9, 0xa0, 0x8f, 0xdb, 0x83, 0xb9, 0xb0, 0x1d, 0xe, 0x59, 0x18, 0xd8, 0xae, 0xeb, 0x9a, 0x2a, 0x80, 0x5d, 0x20, 0x3, 0xc7, 0x17, 0x64, 0xdd, 0x6, 0xb8, 0x10}, []uint8{0x54, 0xe6, 0xbe, 0xfd, 0x6d, 0xe1, 0x3a, 0xfe, 0x94, 0xe2, 0x26, 0xf0, 0x19, 0x7d, 0xe2, 0x7, 0x5, 0x5d, 0xf4, 0x47, 0x44, 0x1d, 0x39, 0x87, 0x25, 0xf9, 0xb4, 0x2b, 0xc6, 0x1b, 0xa9, 0x87}, []uint8{0xce, 0x5d, 0x9a, 0x22, 0x4e, 0x4f, 0x7b, 0x9e, 0xdc, 0x35, 0xce, 0xf, 0x15, 0xe2, 0x65, 0x41, 0xba, 0x84, 0x22, 0x8b, 0xf0, 0x78, 0xf3, 0xd0, 0xa8, 0x23, 0x93, 0xbf, 0xbb, 0xe9, 0xc9, 0x4c}, []uint8{0xea, 0xd5, 0x64, 0x88, 0x2b, 0xe0, 0x42, 0x83, 0x38, 0x14, 0x87, 0xe5, 0xc0, 0x81, 0x62, 0x6e, 0x48, 0xc4, 0x75, 0xe1, 0xfa, 0xd5, 0xe0, 0x6c, 0x32, 0xef, 0xd2, 0x6c, 0xc4, 0xf5, 0x81, 0x34}, []uint8{0x3, 0x63, 0x23, 0x6b, 0x48, 0x6a, 0x5a, 0x7b, 0x8f, 0x98, 0x80, 0x78, 0xc5, 0xcb, 0xf, 0x83, 0xb, 0xcb, 0xed, 0x92, 0xc8, 0x69, 0xe3, 0xd8, 0xc4, 0x34, 0xcf, 0xc5, 0x75, 0x8, 0x24, 0x1e}}, rawStorageRoot: []uint8(nil), storageRoot: []uint8(nil)}
	key := []uint8{125, 241, 76, 23, 176, 162, 208, 140, 129, 137, 126, 41, 104, 90, 161, 56, 171, 154, 165, 241, 216, 236, 255, 79, 16, 70, 49, 247, 152, 181, 5, 189}
	expected := []byte{}
	rawTestTree := []interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{[]uint8{0x57, 0xe6, 0xa3, 0x26, 0x9, 0x5d, 0x7e, 0x99, 0xb1, 0x50, 0x89, 0xb2, 0x25, 0x4c, 0x7, 0x8c, 0x19, 0xa6, 0x90, 0x41, 0x37, 0x5f, 0xe1, 0xf6, 0x9c, 0x9c, 0xfa, 0x2d, 0x41, 0x55, 0x81, 0xcd}, []interface{}{[]uint8{0xa7, 0xc4, 0x92, 0x39, 0xe7, 0xf6, 0x68, 0xe5, 0x9b, 0x41, 0x9a, 0x5, 0x4a, 0xe8, 0xbe, 0xd2, 0xe0, 0xfc, 0x51, 0x83, 0x63, 0x1f, 0x12, 0x96, 0x5e, 0x1a, 0x5a, 0x1d, 0x8c, 0xf5, 0x2c, 0x5d}, []interface{}{[]interface{}{[]interface{}{[]uint8{0xc}, []uint8{0x6}, []interface{}{[]uint8{0x24, 0x7e, 0xb6, 0x81, 0x93, 0x2c, 0x13, 0x93, 0x1b, 0x1b, 0xc9, 0x24, 0x71, 0xb5, 0x57, 0x56, 0x51, 0x76, 0xbc, 0xbf, 0xf1, 0x9a, 0x63, 0x10, 0x23, 0x93, 0xe0, 0xa3, 0x2c, 0xa0, 0x97, 0x91}, []uint8{0x67, 0x72, 0x65, 0x65, 0x6e, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x16}}, []interface{}{[]uint8{0x69, 0x9e, 0x80, 0x64, 0xf4, 0x5a, 0x12, 0x8a, 0xaf, 0x7c, 0x6d, 0xaf, 0x9, 0x70, 0x64, 0xb7, 0x85, 0xc0, 0x28, 0x78, 0x69, 0x8d, 0x69, 0xf6, 0x38, 0xc, 0x6f, 0x96, 0x77, 0xf1, 0xb3, 0x30}, []uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}}, []interface{}{[]uint8{0x7e, 0x8, 0xcd, 0x79, 0xff, 0x3f, 0xc4, 0xd5, 0xc2, 0x40, 0x1, 0x27, 0x2a, 0x52, 0x7b, 0x92, 0x73, 0xd5, 0xfc, 0xa, 0x38, 0xeb, 0xe3, 0x13, 0x82, 0x6, 0x26, 0xea, 0x97, 0x66, 0xaf, 0xde}, []uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x75, 0x5d, 0xcf, 0x42, 0x7f, 0xc5, 0x35, 0x24, 0x73, 0xf9, 0xee, 0xef, 0x69, 0x27, 0x35, 0x36, 0x1f, 0x15, 0xdb, 0xf5}}}, []uint8{0xfe, 0xe2, 0xa2, 0x67, 0xad, 0x80, 0xaf, 0x3c, 0x5e, 0xc6, 0x55, 0x3d, 0xd8, 0xa8, 0xbf, 0xe3, 0x62, 0xd5, 0xa7, 0xc7, 0x76, 0x2c, 0x38, 0xa4, 0xa5, 0xee, 0xd6, 0xd5, 0x2, 0x7c, 0x77, 0x67}}, []uint8{0x7a, 0x93, 0x21, 0xb, 0x9c, 0x20, 0xcc, 0xbd, 0x9a, 0xb1, 0x41, 0xa7, 0x6e, 0x6c, 0x62, 0xec, 0x76, 0xe8, 0x7d, 0x5c, 0x3e, 0xcd, 0xd7, 0xd5, 0xfc, 0xc6, 0xbd, 0xe, 0x9, 0xb3, 0x43, 0x20}}}}, []uint8{0x72, 0xdf, 0x44, 0x37, 0x54, 0xee, 0xe2, 0x7, 0xd2, 0xdb, 0xce, 0xb9, 0x94, 0x59, 0xaa, 0xcc, 0x10, 0x7c, 0x70, 0xed, 0x3a, 0x9c, 0x32, 0xbd, 0xf4, 0x2d, 0x1a, 0xa4, 0xda, 0x77, 0x63, 0x8e}}, []uint8{0x15, 0x59, 0x62, 0x65, 0xc0, 0x92, 0x9f, 0x8b, 0x18, 0x10, 0x28, 0xea, 0x48, 0x10, 0xbd, 0x1d, 0xbc, 0x34, 0x57, 0x11, 0xd9, 0xd6, 0xfb, 0x68, 0x24, 0xc8, 0x36, 0xad, 0xc0, 0x9d, 0x32, 0x0}}, []uint8{0x89, 0x42, 0x32, 0x45, 0xaa, 0x16, 0xb, 0xc1, 0x2b, 0x35, 0xe5, 0xf9, 0x8d, 0x9d, 0xea, 0xd1, 0xdd, 0xc2, 0x6d, 0xe4, 0x77, 0xd3, 0x8d, 0xb5, 0x27, 0xff, 0xd0, 0x2, 0xdb, 0xa5, 0x57, 0x0}}, []uint8{0x86, 0xc2, 0xfc, 0x9, 0xfe, 0x14, 0xef, 0x51, 0x17, 0x2, 0x39, 0x19, 0x5d, 0xbd, 0xe8, 0x0, 0x30, 0x9e, 0xc5, 0xa6, 0x80, 0x56, 0x3, 0x58, 0x11, 0x29, 0xd9, 0xb6, 0x8a, 0x49, 0x19, 0x2}}
	assertSame(t, rawTestTree, key, expected, testRoots)
}

func TestValidateMerkleTree1(t *testing.T) {
	// 8 bit test 1
	testRoots := &AccountRoots{AccountRoots: [][]uint8{{0x1a, 0xfd, 0xf0, 0x7c, 0x95, 0x2c, 0x83, 0x5b, 0x80, 0x55, 0x5c, 0xc0, 0x24, 0xd6, 0x2d, 0x25, 0xe5, 0xe8, 0x6e, 0x45, 0x4f, 0x52, 0x72, 0xea, 0x4e, 0xee, 0x11, 0x2f, 0x1d, 0xd8, 0xf5, 0x25}, []uint8{0xa4, 0x28, 0xc0, 0x5, 0x77, 0x64, 0x4, 0x31, 0xd1, 0x51, 0x74, 0xf3, 0xcd, 0x8a, 0x40, 0x22, 0xed, 0x47, 0x6c, 0xb3, 0x72, 0x70, 0xfb, 0x22, 0x95, 0x56, 0x77, 0xe7, 0x1, 0x34, 0x88, 0x34}, []uint8{0x5e, 0xb8, 0x98, 0xb, 0x16, 0x6e, 0x93, 0x24, 0x78, 0xd0, 0xdf, 0xb5, 0x92, 0xfb, 0x7, 0xa0, 0x4f, 0xb2, 0x70, 0xf8, 0xcc, 0x82, 0xde, 0x1c, 0x8a, 0x54, 0x26, 0x31, 0x5a, 0x25, 0x28, 0xf8}, []uint8{0x50, 0xdc, 0x3d, 0x15, 0x60, 0xee, 0xd7, 0x24, 0x63, 0x82, 0x4d, 0xd7, 0x73, 0x2c, 0xfa, 0x99, 0xcc, 0x27, 0x5b, 0xa7, 0xb, 0xd4, 0x2a, 0x43, 0xf2, 0xbf, 0xf4, 0x2a, 0xf5, 0xb, 0x5b, 0xa9}, []uint8{0x22, 0xf1, 0x33, 0xd3, 0xc3, 0x6d, 0x6, 0xbe, 0xeb, 0x8c, 0xf7, 0x2b, 0x70, 0x72, 0x3f, 0xd7, 0x58, 0x86, 0xe1, 0xdb, 0xb2, 0x26, 0xba, 0x74, 0x5a, 0x40, 0x41, 0x4e, 0x54, 0x55, 0x45, 0xaf}, []uint8{0xd6, 0x9a, 0xd1, 0x2b, 0x1c, 0x5d, 0x7e, 0x1f, 0x6b, 0x44, 0xe5, 0xb9, 0x26, 0x8f, 0x1, 0xde, 0x5a, 0xe, 0xf, 0x31, 0xc0, 0xf6, 0x41, 0x59, 0xa0, 0x50, 0x5b, 0x9f, 0x42, 0x55, 0x58, 0x4}, []uint8{0x7e, 0xad, 0x2d, 0x77, 0x92, 0xe0, 0xac, 0x98, 0x34, 0x89, 0x3b, 0xb1, 0xf6, 0xdf, 0xe6, 0xbf, 0x8b, 0x46, 0xe0, 0xf4, 0x6d, 0x86, 0x93, 0x9f, 0x72, 0x5, 0xae, 0xcb, 0xca, 0xd5, 0x21, 0x43}, []uint8{0x1b, 0xd3, 0xea, 0xe7, 0x27, 0x8e, 0xbc, 0xb3, 0x83, 0x3c, 0xd0, 0x9d, 0x14, 0xee, 0x11, 0xb, 0x66, 0x91, 0xb, 0x3d, 0xb7, 0x49, 0x9b, 0x95, 0xa8, 0x32, 0x71, 0x28, 0x5a, 0xa4, 0x5d, 0xa5}, []uint8{0x99, 0xf8, 0x9d, 0xb3, 0xc6, 0xd7, 0xfe, 0x63, 0x8b, 0x9d, 0xa0, 0x4e, 0x32, 0x6, 0x7b, 0x4c, 0x5e, 0x87, 0x26, 0xb6, 0xec, 0x2c, 0x35, 0x3a, 0xf3, 0x98, 0x70, 0x89, 0xd2, 0x78, 0x1d, 0x43}, []uint8{0x79, 0xd1, 0x22, 0x3e, 0xf9, 0xc5, 0xd7, 0xfe, 0x88, 0x58, 0xc6, 0xda, 0xf5, 0xa0, 0xec, 0xf6, 0x22, 0x47, 0x4c, 0x86, 0x28, 0x7c, 0x58, 0xf9, 0x4e, 0x26, 0xdf, 0xee, 0x58, 0x3d, 0x37, 0x16}, []uint8{0xff, 0xad, 0x1, 0xed, 0x68, 0xd7, 0xa6, 0x9e, 0x53, 0xc1, 0xe4, 0x5c, 0x7c, 0xc5, 0xbd, 0xaf, 0x5a, 0x2, 0xf1, 0xb4, 0xe2, 0x5c, 0xac, 0xcd, 0x2b, 0x3c, 0x28, 0xc, 0x2a, 0xaf, 0x28, 0x16}, []uint8{0x18, 0xbc, 0xdf, 0xaa, 0xf1, 0x9, 0xa0, 0x8f, 0xdb, 0x83, 0xb9, 0xb0, 0x1d, 0xe, 0x59, 0x18, 0xd8, 0xae, 0xeb, 0x9a, 0x2a, 0x80, 0x5d, 0x20, 0x3, 0xc7, 0x17, 0x64, 0xdd, 0x6, 0xb8, 0x10}, []uint8{0x54, 0xe6, 0xbe, 0xfd, 0x6d, 0xe1, 0x3a, 0xfe, 0x94, 0xe2, 0x26, 0xf0, 0x19, 0x7d, 0xe2, 0x7, 0x5, 0x5d, 0xf4, 0x47, 0x44, 0x1d, 0x39, 0x87, 0x25, 0xf9, 0xb4, 0x2b, 0xc6, 0x1b, 0xa9, 0x87}, []uint8{0xce, 0x5d, 0x9a, 0x22, 0x4e, 0x4f, 0x7b, 0x9e, 0xdc, 0x35, 0xce, 0xf, 0x15, 0xe2, 0x65, 0x41, 0xba, 0x84, 0x22, 0x8b, 0xf0, 0x78, 0xf3, 0xd0, 0xa8, 0x23, 0x93, 0xbf, 0xbb, 0xe9, 0xc9, 0x4c}, []uint8{0xea, 0xd5, 0x64, 0x88, 0x2b, 0xe0, 0x42, 0x83, 0x38, 0x14, 0x87, 0xe5, 0xc0, 0x81, 0x62, 0x6e, 0x48, 0xc4, 0x75, 0xe1, 0xfa, 0xd5, 0xe0, 0x6c, 0x32, 0xef, 0xd2, 0x6c, 0xc4, 0xf5, 0x81, 0x34}, []uint8{0x3, 0x63, 0x23, 0x6b, 0x48, 0x6a, 0x5a, 0x7b, 0x8f, 0x98, 0x80, 0x78, 0xc5, 0xcb, 0xf, 0x83, 0xb, 0xcb, 0xed, 0x92, 0xc8, 0x69, 0xe3, 0xd8, 0xc4, 0x34, 0xcf, 0xc5, 0x75, 0x8, 0x24, 0x1e}}, rawStorageRoot: []uint8(nil), storageRoot: []uint8(nil)}
	key := []uint8{126, 8, 205, 121, 255, 63, 196, 213, 194, 64, 1, 39, 42, 82, 123, 146, 115, 213, 252, 10, 56, 235, 227, 19, 130, 6, 38, 234, 151, 102, 175, 225}
	expected := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6}
	rawTestTree := []interface{}{[]interface{}{[]interface{}{[]uint8{0x80, 0xbd, 0x9e, 0xb0, 0x51, 0x2c, 0x8c, 0x33, 0xf7, 0x6e, 0xb0, 0xac, 0x91, 0xdb, 0x78, 0xac, 0x13, 0xbd, 0xaa, 0x10, 0x3, 0x44, 0x5c, 0x57, 0xef, 0xd0, 0xa1, 0x80, 0x7a, 0x4b, 0x9b, 0x72}, []interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{[]interface{}{[]uint8{0x60, 0xf, 0x8b, 0xee, 0xbd, 0xa5, 0x9d, 0x3f, 0xe9, 0x5d, 0xe4, 0xb5, 0x2, 0x82, 0x98, 0xce, 0xa6, 0x43, 0x9d, 0x5a, 0x6f, 0x1c, 0xc2, 0x5e, 0xc, 0xaf, 0x39, 0x92, 0x17, 0x4a, 0x59, 0xe0}, []interface{}{[]uint8{0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31}, []uint8{0x3}, []interface{}{[]uint8{0x7e, 0x8, 0xcd, 0x79, 0xff, 0x3f, 0xc4, 0xd5, 0xc2, 0x40, 0x1, 0x27, 0x2a, 0x52, 0x7b, 0x92, 0x73, 0xd5, 0xfc, 0xa, 0x38, 0xeb, 0xe3, 0x13, 0x82, 0x6, 0x26, 0xea, 0x97, 0x66, 0xaf, 0xe1}, []uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6}}, []interface{}{[]uint8{0xbf, 0xe4, 0x8d, 0xf5, 0x70, 0xb7, 0x73, 0x8d, 0xe7, 0xd5, 0x3a, 0x2f, 0xab, 0xe6, 0xd6, 0xac, 0x65, 0x6b, 0xd3, 0x6c, 0xee, 0x99, 0xd6, 0xe2, 0xc5, 0x10, 0xcc, 0xa7, 0x40, 0x58, 0x33, 0x4b}, []uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8a, 0x80, 0x73, 0x2a, 0xc4, 0x3d, 0x69, 0xcd, 0xa0, 0x44, 0xe, 0xc7, 0x92, 0x67, 0x8d, 0x47, 0x9f, 0xe1, 0x52, 0xa}}}}, []uint8{0x62, 0xfd, 0x53, 0xdd, 0xa2, 0xce, 0x4a, 0x5, 0xa1, 0x6b, 0x13, 0x23, 0x5a, 0x6b, 0xcd, 0x3f, 0xfe, 0xcd, 0x52, 0x9, 0x9b, 0xd6, 0x36, 0x2f, 0xfe, 0x27, 0x34, 0xcd, 0x26, 0x55, 0x43, 0x90}}, []uint8{0xe4, 0x71, 0x99, 0xd7, 0x1d, 0x9e, 0x3, 0x9d, 0x75, 0xd8, 0x4c, 0xcc, 0xc8, 0x85, 0x83, 0xbf, 0x26, 0x78, 0x17, 0x18, 0x57, 0x50, 0xff, 0x5e, 0x1, 0x4d, 0xf9, 0x40, 0x4a, 0x76, 0x40, 0x90}}, []uint8{0x1e, 0xb5, 0x7d, 0x2e, 0x5b, 0x2f, 0x5, 0xd7, 0x3, 0xa1, 0x46, 0xa3, 0x1f, 0x6d, 0x6a, 0xcd, 0xe3, 0x36, 0x4e, 0x77, 0x91, 0x2a, 0xa, 0x38, 0x63, 0xed, 0xdc, 0xa4, 0xc9, 0x3a, 0xfb, 0x9a}}, []uint8{0x9e, 0xf7, 0x7, 0x28, 0xf1, 0x26, 0x30, 0xb9, 0x4b, 0x69, 0x3, 0xd8, 0x95, 0xad, 0xae, 0x10, 0x76, 0xb2, 0x45, 0x95, 0x87, 0x3c, 0x6b, 0xe0, 0xad, 0x18, 0x4f, 0x7c, 0xfe, 0x5e, 0x60, 0x78}}, []uint8{0x24, 0x97, 0x64, 0xd2, 0x1f, 0x41, 0x34, 0x72, 0x12, 0x3d, 0xeb, 0xd1, 0x7f, 0x4d, 0x6c, 0x4b, 0x12, 0x1b, 0xf8, 0x75, 0x25, 0xf5, 0x8c, 0xb6, 0xb3, 0x87, 0x7c, 0x73, 0x59, 0xb0, 0x50, 0x6f}}}, []uint8{0x8d, 0x63, 0x80, 0x79, 0xe4, 0x9a, 0x2, 0xbb, 0x7, 0xf4, 0xb0, 0x14, 0x55, 0xfd, 0x54, 0x42, 0x82, 0xe7, 0xb0, 0x8c, 0xb2, 0xb1, 0x74, 0xf0, 0x55, 0xd7, 0x4e, 0x7c, 0x32, 0xe3, 0xa4, 0xee}}, []uint8{0x8d, 0x34, 0x4b, 0x66, 0x8a, 0x57, 0x14, 0xae, 0xd, 0x8, 0x51, 0xdd, 0x8d, 0xb0, 0x6, 0x67, 0x56, 0xa1, 0xbc, 0x1d, 0x8, 0x3f, 0xdf, 0x97, 0xf1, 0x73, 0x80, 0x60, 0x5b, 0xc8, 0x4, 0x4c}}
	assertSame(t, rawTestTree, key, expected, testRoots)
}

func TestValidateMerkleTree2(t *testing.T) {
	// 8 bit test 2
	testRoots := &AccountRoots{AccountRoots: [][]uint8{[]uint8{0x1a, 0xfd, 0xf0, 0x7c, 0x95, 0x2c, 0x83, 0x5b, 0x80, 0x55, 0x5c, 0xc0, 0x24, 0xd6, 0x2d, 0x25, 0xe5, 0xe8, 0x6e, 0x45, 0x4f, 0x52, 0x72, 0xea, 0x4e, 0xee, 0x11, 0x2f, 0x1d, 0xd8, 0xf5, 0x25}, []uint8{0xa4, 0x28, 0xc0, 0x5, 0x77, 0x64, 0x4, 0x31, 0xd1, 0x51, 0x74, 0xf3, 0xcd, 0x8a, 0x40, 0x22, 0xed, 0x47, 0x6c, 0xb3, 0x72, 0x70, 0xfb, 0x22, 0x95, 0x56, 0x77, 0xe7, 0x1, 0x34, 0x88, 0x34}, []uint8{0x5e, 0xb8, 0x98, 0xb, 0x16, 0x6e, 0x93, 0x24, 0x78, 0xd0, 0xdf, 0xb5, 0x92, 0xfb, 0x7, 0xa0, 0x4f, 0xb2, 0x70, 0xf8, 0xcc, 0x82, 0xde, 0x1c, 0x8a, 0x54, 0x26, 0x31, 0x5a, 0x25, 0x28, 0xf8}, []uint8{0x50, 0xdc, 0x3d, 0x15, 0x60, 0xee, 0xd7, 0x24, 0x63, 0x82, 0x4d, 0xd7, 0x73, 0x2c, 0xfa, 0x99, 0xcc, 0x27, 0x5b, 0xa7, 0xb, 0xd4, 0x2a, 0x43, 0xf2, 0xbf, 0xf4, 0x2a, 0xf5, 0xb, 0x5b, 0xa9}, []uint8{0x22, 0xf1, 0x33, 0xd3, 0xc3, 0x6d, 0x6, 0xbe, 0xeb, 0x8c, 0xf7, 0x2b, 0x70, 0x72, 0x3f, 0xd7, 0x58, 0x86, 0xe1, 0xdb, 0xb2, 0x26, 0xba, 0x74, 0x5a, 0x40, 0x41, 0x4e, 0x54, 0x55, 0x45, 0xaf}, []uint8{0xd6, 0x9a, 0xd1, 0x2b, 0x1c, 0x5d, 0x7e, 0x1f, 0x6b, 0x44, 0xe5, 0xb9, 0x26, 0x8f, 0x1, 0xde, 0x5a, 0xe, 0xf, 0x31, 0xc0, 0xf6, 0x41, 0x59, 0xa0, 0x50, 0x5b, 0x9f, 0x42, 0x55, 0x58, 0x4}, []uint8{0x7e, 0xad, 0x2d, 0x77, 0x92, 0xe0, 0xac, 0x98, 0x34, 0x89, 0x3b, 0xb1, 0xf6, 0xdf, 0xe6, 0xbf, 0x8b, 0x46, 0xe0, 0xf4, 0x6d, 0x86, 0x93, 0x9f, 0x72, 0x5, 0xae, 0xcb, 0xca, 0xd5, 0x21, 0x43}, []uint8{0x1b, 0xd3, 0xea, 0xe7, 0x27, 0x8e, 0xbc, 0xb3, 0x83, 0x3c, 0xd0, 0x9d, 0x14, 0xee, 0x11, 0xb, 0x66, 0x91, 0xb, 0x3d, 0xb7, 0x49, 0x9b, 0x95, 0xa8, 0x32, 0x71, 0x28, 0x5a, 0xa4, 0x5d, 0xa5}, []uint8{0x99, 0xf8, 0x9d, 0xb3, 0xc6, 0xd7, 0xfe, 0x63, 0x8b, 0x9d, 0xa0, 0x4e, 0x32, 0x6, 0x7b, 0x4c, 0x5e, 0x87, 0x26, 0xb6, 0xec, 0x2c, 0x35, 0x3a, 0xf3, 0x98, 0x70, 0x89, 0xd2, 0x78, 0x1d, 0x43}, []uint8{0x79, 0xd1, 0x22, 0x3e, 0xf9, 0xc5, 0xd7, 0xfe, 0x88, 0x58, 0xc6, 0xda, 0xf5, 0xa0, 0xec, 0xf6, 0x22, 0x47, 0x4c, 0x86, 0x28, 0x7c, 0x58, 0xf9, 0x4e, 0x26, 0xdf, 0xee, 0x58, 0x3d, 0x37, 0x16}, []uint8{0xff, 0xad, 0x1, 0xed, 0x68, 0xd7, 0xa6, 0x9e, 0x53, 0xc1, 0xe4, 0x5c, 0x7c, 0xc5, 0xbd, 0xaf, 0x5a, 0x2, 0xf1, 0xb4, 0xe2, 0x5c, 0xac, 0xcd, 0x2b, 0x3c, 0x28, 0xc, 0x2a, 0xaf, 0x28, 0x16}, []uint8{0x18, 0xbc, 0xdf, 0xaa, 0xf1, 0x9, 0xa0, 0x8f, 0xdb, 0x83, 0xb9, 0xb0, 0x1d, 0xe, 0x59, 0x18, 0xd8, 0xae, 0xeb, 0x9a, 0x2a, 0x80, 0x5d, 0x20, 0x3, 0xc7, 0x17, 0x64, 0xdd, 0x6, 0xb8, 0x10}, []uint8{0x54, 0xe6, 0xbe, 0xfd, 0x6d, 0xe1, 0x3a, 0xfe, 0x94, 0xe2, 0x26, 0xf0, 0x19, 0x7d, 0xe2, 0x7, 0x5, 0x5d, 0xf4, 0x47, 0x44, 0x1d, 0x39, 0x87, 0x25, 0xf9, 0xb4, 0x2b, 0xc6, 0x1b, 0xa9, 0x87}, []uint8{0xce, 0x5d, 0x9a, 0x22, 0x4e, 0x4f, 0x7b, 0x9e, 0xdc, 0x35, 0xce, 0xf, 0x15, 0xe2, 0x65, 0x41, 0xba, 0x84, 0x22, 0x8b, 0xf0, 0x78, 0xf3, 0xd0, 0xa8, 0x23, 0x93, 0xbf, 0xbb, 0xe9, 0xc9, 0x4c}, []uint8{0xea, 0xd5, 0x64, 0x88, 0x2b, 0xe0, 0x42, 0x83, 0x38, 0x14, 0x87, 0xe5, 0xc0, 0x81, 0x62, 0x6e, 0x48, 0xc4, 0x75, 0xe1, 0xfa, 0xd5, 0xe0, 0x6c, 0x32, 0xef, 0xd2, 0x6c, 0xc4, 0xf5, 0x81, 0x34}, []uint8{0x3, 0x63, 0x23, 0x6b, 0x48, 0x6a, 0x5a, 0x7b, 0x8f, 0x98, 0x80, 0x78, 0xc5, 0xcb, 0xf, 0x83, 0xb, 0xcb, 0xed, 0x92, 0xc8, 0x69, 0xe3, 0xd8, 0xc4, 0x34, 0xcf, 0xc5, 0x75, 0x8, 0x24, 0x1e}}, rawStorageRoot: []uint8(nil), storageRoot: []uint8(nil)}
	key := []byte{0x7d, 0xf1, 0x4c, 0x17, 0xb0, 0xa2, 0xd0, 0x8c, 0x81, 0x89, 0x7e, 0x29, 0x68, 0x5a, 0xa1, 0x38, 0xab, 0x9a, 0xa5, 0xf1, 0xd8, 0xec, 0xff, 0x4f, 0x10, 0x46, 0x31, 0xf7, 0x98, 0xb5, 0x5, 0xbd}
	rawTestTree := []interface{}{[]interface{}{[]interface{}{[]uint8{0xd9, 0xbb, 0x20, 0x9b, 0xb8, 0x52, 0x35, 0xfa, 0x38, 0x68, 0xaa, 0x5c, 0x34, 0xf5, 0xd4, 0xbf, 0x1a, 0x6b, 0x6d, 0xaa, 0x10, 0x91, 0x1b, 0xfb, 0x12, 0xd6, 0x45, 0x89, 0x98, 0xa5, 0x33, 0x17}, []interface{}{[]interface{}{[]uint8{0xeb, 0xf4, 0x7c, 0x46, 0x91, 0x8e, 0xb6, 0x36, 0x5, 0xe0, 0xd7, 0x3c, 0x8a, 0x8e, 0x8, 0x56, 0x1d, 0xf7, 0x27, 0x53, 0x91, 0x11, 0x97, 0x83, 0x48, 0x8d, 0x4c, 0x67, 0x7b, 0xda, 0x41, 0xc1}, []interface{}{[]interface{}{[]uint8{0xdc, 0xc9, 0x34, 0x14, 0x43, 0xf3, 0x70, 0x48, 0x47, 0xb0, 0x3a, 0x9d, 0xd8, 0x47, 0xad, 0x4c, 0x7, 0x78, 0xfc, 0xd1, 0x9b, 0x7a, 0x3a, 0x3c, 0xfd, 0x47, 0x65, 0xf5, 0x69, 0x2a, 0x0, 0x3e}, []interface{}{[]uint8{0xd7, 0xcf, 0xca, 0x61, 0x5, 0xfb, 0x3d, 0x60, 0xa7, 0x38, 0xd1, 0x75, 0x7a, 0x2d, 0x45, 0xc6, 0x2c, 0xf4, 0xf3, 0xbe, 0x91, 0x85, 0x49, 0xb7, 0x1, 0x76, 0xb3, 0x77, 0x77, 0x5a, 0xe1, 0xa8}, []interface{}{[]uint8{0x2b}, []uint8{0xe}}}}, []uint8{0x7d, 0x2d, 0x6a, 0xcb, 0x1e, 0xfd, 0x17, 0x4a, 0xbb, 0x94, 0x12, 0xd9, 0x64, 0x57, 0x74, 0xd, 0xe0, 0x23, 0xa3, 0x50, 0x12, 0x24, 0xc5, 0x90, 0xcc, 0xfd, 0x49, 0x5a, 0x38, 0x4b, 0x15, 0x53}}}, []uint8{0xbb, 0xc4, 0xe0, 0x67, 0x72, 0xac, 0xc2, 0x63, 0x7c, 0x66, 0x98, 0x2f, 0x99, 0x3a, 0xc8, 0xfa, 0xf4, 0x7e, 0x64, 0x38, 0x10, 0x3f, 0x53, 0xb, 0xf7, 0xbd, 0xc1, 0x52, 0xce, 0x27, 0xce, 0x80}}}, []uint8{0x19, 0xdd, 0x36, 0xc0, 0x50, 0xb7, 0x5a, 0xa1, 0x19, 0x22, 0x14, 0xa4, 0xe3, 0x55, 0x2a, 0x53, 0x9e, 0x37, 0xd1, 0x2e, 0xc7, 0xfc, 0xf6, 0xb4, 0x88, 0xd0, 0x77, 0x95, 0x75, 0x12, 0xc2, 0x41}}, []uint8{0xda, 0x4c, 0x5c, 0xdc, 0xe6, 0x1c, 0x44, 0x4f, 0x90, 0x18, 0xd, 0xfa, 0xa5, 0x54, 0xba, 0xd4, 0x85, 0xe8, 0x74, 0xb7, 0x35, 0xa5, 0xb6, 0xc0, 0xb8, 0x87, 0x88, 0x87, 0x71, 0x90, 0x8c, 0x64}}
	expected := []byte{}
	assertSame(t, rawTestTree, key, expected, testRoots)
}

func assertSame(t *testing.T, rawTestTree []interface{}, key []uint8, expected []uint8, testRoots *AccountRoots) {
	acvTree, err := NewMerkleTree(rawTestTree)

	if err != nil {
		t.Fatalf("Tree parsing should succeed")
	}

	// Verify the calculated proof value matches the specific known root
	if testRoots.Find(acvTree.RootHash) < 0 {
		t.Fatalf("Generated invalid root hash %s", util.EncodeToString(acvTree.RootHash))
	}
	if testRoots.Find(acvTree.RootHash) != int(acvTree.Modulo) {
		t.Fatalf("Found root hash but modulo is wrong")
	}

	value, err := acvTree.Get(key)
	if len(expected) == 0 {
		if err == nil {
			t.Fatalf("Value should not exist")
		}
		return
	}

	if err != nil {
		t.Fatalf("Value should exist")
	}

	if !bytes.Equal(expected, value) {
		fmt.Printf("Expected: %v got: %v\n", expected, value)
	}
}
