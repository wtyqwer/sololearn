package main

///* Resolve file signatures
//*
//* resource:
//* 1. https://en.wikipedia.org/wiki/List_of_file_signatures
//*/
//
//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	{
//		rawbytes := []byte{
//			0x53, 0x51, 0x4C, 0x69, 0x74, 0x65, 0x20, 0x66,
//			0x6F, 0x72, 0x6D, 0x61, 0x74, 0x20, 0x33, 0x00,
//			0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
//		}
//		fd := resolveFileSignature(rawbytes, "key1")
//		if fd != nil {
//			fmt.Printf("num:%d , %v: %s\n", 1, fd.Extensions, fd.Description)
//		} else {
//			fmt.Printf("unknown file signature\n")
//		}
//	}
//	{
//		rawbytes := []byte{
//			0x53, 0x51, 0x4C, 0x69,
//			0x66, 0x74, 0x79, 0x70, 0x33, 0x67,
//			0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
//		}
//		fd := resolveFileSignature(rawbytes, "key2")
//		if fd != nil {
//			fmt.Printf("num:%d , %v: %s\n", 2, fd.Extensions, fd.Description)
//		} else {
//			fmt.Printf("unknown file signature\n")
//		}
//	}
//	{
//		rawbytes := []byte{
//			0x46, 0x4F, 0x52, 0x4D,
//			0x01, 0x01, 0x01, 0x01,
//			0x41, 0x43, 0x42, 0x4D,
//			0x01, 0x01, 0x01, 0x01,
//		}
//		fd := resolveFileSignature(rawbytes, "key3")
//		if fd != nil {
//			fmt.Printf("num:%d , %v: %s\n", 3, fd.Extensions, fd.Description)
//		} else {
//			fmt.Printf("unknown file signature\n")
//		}
//	}
//	{
//		rawbytes := []byte{
//			0x46, 0x4F, 0x52, 0x4D,
//			0x01, 0x01, 0x01, 0x01,
//			0x01, 0x01, 0x01, 0x01,
//			0x01, 0x01, 0x01, 0x01,
//		}
//		fd := resolveFileSignature(rawbytes, "key4")
//		if fd != nil {
//			fmt.Printf("num:%d , %v: %s\n", 4, fd.Extensions, fd.Description)
//		} else {
//			fmt.Printf("unknown file signature\n")
//		}
//	}
//}
//
//type FileDescriptor struct {
//	Description string
//	Extensions  []string
//}
//
//type RawBytesStruct struct {
//	StartIndex  int
//	EndIndex    int
//	Level       int //0-root/1-parent/2-child
//	RawBytes    []byte
//	Description string
//	Extensions  []string
//	Child       []*RawBytesStruct
//}
//
//var ResultMap = make(map[string]*FileDescriptor)
//
///*func Test(root *RawBytesStruct,resultKey string) func(rawbytes []byte){
//	return func(rawbytes []byte){
//		for _,rawBytesStruct := range rawBytesStructs {
//			if reflect.DeepEqual(rawbytes[rawBytesStruct.StartIndex:rawBytesStruct.EndIndex],rawBytesStruct.RawBytes) {
//				if rawBytesStruct.Level == 0 {
//					continue
//				} else {
//					result := &FileDescriptor{
//						Description: rawBytesStruct.Description,
//						Extensions:  rawBytesStruct.Extensions,
//					}
//					ResultMap[resultKey] = result
//				}
//			}
//		}
//	}
//}*/
//
//func (r *RawBytesStruct) Test(rawbytes []byte, resultKey string) {
//	if r.Level == 0 && len(r.Child) != 0 {
//		for i := 0; i < len(r.Child); i++ {
//			r.Child[i].Test(rawbytes, resultKey)
//		}
//	}
//	if r.Level == 1 && len(r.Child) != 0 {
//		if reflect.DeepEqual(rawbytes[r.StartIndex:r.EndIndex], r.RawBytes) {
//			for i := 0; i < len(r.Child); i++ {
//				r.Child[i].Test(rawbytes, resultKey)
//			}
//		}
//	}
//	if r.Level == 2 {
//		if reflect.DeepEqual(rawbytes[r.StartIndex:r.EndIndex], r.RawBytes) {
//
//			result := &FileDescriptor{
//				Description: r.Description,
//				Extensions:  r.Extensions,
//			}
//			ResultMap[resultKey] = result
//			return
//		}
//	}
//}
//
//// TODO: 🔨 Refactor this function!!
//func resolveFileSignature(rawbytes []byte, resultKey string) *FileDescriptor {
//	/*	if reflect.DeepEqual(rawbytes[:16], []byte{
//			0x53, 0x51, 0x4C, 0x69, 0x74, 0x65, 0x20, 0x66,
//			0x6F, 0x72, 0x6D, 0x61, 0x74, 0x20, 0x33, 0x00,
//		}) {
//			return &FileDescriptor{
//				Description: "SQLite Database",
//				Extensions:  []string{"sqlitedb", "sqlite", "db"},
//			}
//		} else if reflect.DeepEqual(rawbytes[4:10], []byte{
//			0x66, 0x74, 0x79, 0x70, 0x33, 0x67,
//		}) {
//			return &FileDescriptor{
//				Description: "3rd Generation Partnership Project 3GPP and 3GPP2 multimedia files",
//				Extensions:  []string{"3gp", "3g2"},
//			}
//		} else if reflect.DeepEqual(rawbytes[:4], []byte{
//			0x46, 0x4F, 0x52, 0x4D,
//		}) {
//			if reflect.DeepEqual(rawbytes[8:12], []byte{
//				0x49, 0x4C, 0x42, 0x4D,
//			}) {
//				return &FileDescriptor{
//					Description: "IFF Interleaved Bitmap Image",
//					Extensions:  []string{"ilbm", "lbm", "ibm", "iff"},
//				}
//			} else if reflect.DeepEqual(rawbytes[8:12], []byte{
//				0x38, 0x53, 0x56, 0x58,
//			}) {
//				return &FileDescriptor{
//					Description: "IFF 8-Bit Sampled Voice",
//					Extensions:  []string{"8svx", "8sv", "svx", "snd", "iff"},
//				}
//			} else if reflect.DeepEqual(rawbytes[8:12], []byte{
//				0x41, 0x43, 0x42, 0x4D,
//			}) {
//				return &FileDescriptor{
//					Description: "Amiga Contiguous Bitmap",
//					Extensions:  []string{"acbm", "iff"},
//				}
//			} else if reflect.DeepEqual(rawbytes[8:12], []byte{
//				0x41, 0x49, 0x46, 0x46,
//			}) {
//				return &FileDescriptor{
//					Description: "Audio Interchange File Format",
//					Extensions:  []string{"aiff", "aif", "aifc", "snd", "iff"},
//				}
//			}
//		}
//		return nil*/
//
//	var root = &RawBytesStruct{
//		Level: 0,
//	}
//
//	var rawBytesStructs = []*RawBytesStruct{}
//	rawBytesStructs = append(rawBytesStructs, &RawBytesStruct{
//		StartIndex:  0,
//		EndIndex:    16,
//		Level:       2,
//		RawBytes:    []byte{0x53, 0x51, 0x4C, 0x69, 0x74, 0x65, 0x20, 0x66, 0x6F, 0x72, 0x6D, 0x61, 0x74, 0x20, 0x33, 0x00},
//		Description: "SQLite Database",
//		Extensions:  []string{"sqlitedb", "sqlite", "db"},
//	})
//
//	rawBytesStructs = append(rawBytesStructs, &RawBytesStruct{
//		StartIndex:  4,
//		EndIndex:    10,
//		Level:       2,
//		RawBytes:    []byte{0x66, 0x74, 0x79, 0x70, 0x33, 0x67},
//		Description: "3rd Generation Partnership Project 3GPP and 3GPP2 multimedia files",
//		Extensions:  []string{"3gp", "3g2"},
//	})
//	rawBytesStructs = append(rawBytesStructs, &RawBytesStruct{
//		StartIndex: 0,
//		EndIndex:   4,
//		Level:      1,
//		RawBytes:   []byte{0x46, 0x4F, 0x52, 0x4D},
//	})
//	var rawBytesStructs1 = []*RawBytesStruct{}
//	rawBytesStructs1 = append(rawBytesStructs1, &RawBytesStruct{
//		StartIndex:  8,
//		EndIndex:    12,
//		Level:       2,
//		RawBytes:    []byte{0x49, 0x4C, 0x42, 0x4D},
//		Description: "IFF Interleaved Bitmap Image",
//		Extensions:  []string{"ilbm", "lbm", "ibm", "iff"},
//	})
//	rawBytesStructs1 = append(rawBytesStructs1, &RawBytesStruct{
//		StartIndex:  8,
//		EndIndex:    12,
//		Level:       2,
//		RawBytes:    []byte{0x38, 0x53, 0x56, 0x58},
//		Description: "IFF 8-Bit Sampled Voice",
//		Extensions:  []string{"8svx", "8sv", "svx", "snd", "iff"},
//	})
//	rawBytesStructs1 = append(rawBytesStructs1, &RawBytesStruct{
//		StartIndex:  8,
//		EndIndex:    12,
//		Level:       2,
//		RawBytes:    []byte{0x41, 0x43, 0x42, 0x4D},
//		Description: "Amiga Contiguous Bitmap",
//		Extensions:  []string{"acbm", "iff"},
//	})
//	rawBytesStructs1 = append(rawBytesStructs1, &RawBytesStruct{
//		StartIndex:  8,
//		EndIndex:    12,
//		Level:       2,
//		RawBytes:    []byte{0x41, 0x49, 0x46, 0x46},
//		Description: "Audio Interchange File Format",
//		Extensions:  []string{"aiff", "aif", "aifc", "snd", "iff"},
//	})
//	rawBytesStructs[2].Child = rawBytesStructs1
//	root.Child = rawBytesStructs
//	root.Test(rawbytes, resultKey)
//	return ResultMap[resultKey]
//
//}
