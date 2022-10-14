package main

///* Resolve file signatures
// *
// * resource:
// * 1. https://en.wikipedia.org/wiki/List_of_file_signatures
// */
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
//		fd := resolveFileSignature(rawbytes)
//		if fd != nil {
//			fmt.Printf("%v: %s\n", fd.Extensions, fd.Description)
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
//		fd := resolveFileSignature(rawbytes)
//		if fd != nil {
//			fmt.Printf("%v: %s\n", fd.Extensions, fd.Description)
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
//		fd := resolveFileSignature(rawbytes)
//		if fd != nil {
//			fmt.Printf("%v: %s\n", fd.Extensions, fd.Description)
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
//		fd := resolveFileSignature(rawbytes)
//		if fd != nil {
//			fmt.Printf("%v: %s\n", fd.Extensions, fd.Description)
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
//// TODO: 🔨 Refactor this function!!
//func resolveFileSignature(rawbytes []byte) *FileDescriptor {
//	if reflect.DeepEqual(rawbytes[:16], []byte{
//		0x53, 0x51, 0x4C, 0x69, 0x74, 0x65, 0x20, 0x66,
//		0x6F, 0x72, 0x6D, 0x61, 0x74, 0x20, 0x33, 0x00,
//	}) {
//		return &FileDescriptor{
//			Description: "SQLite Database",
//			Extensions:  []string{"sqlitedb", "sqlite", "db"},
//		}
//	} else if reflect.DeepEqual(rawbytes[4:10], []byte{
//		0x66, 0x74, 0x79, 0x70, 0x33, 0x67,
//	}) {
//		return &FileDescriptor{
//			Description: "3rd Generation Partnership Project 3GPP and 3GPP2 multimedia files",
//			Extensions:  []string{"3gp", "3g2"},
//		}
//	} else if reflect.DeepEqual(rawbytes[:4], []byte{
//		0x46, 0x4F, 0x52, 0x4D,
//	}) {
//		if reflect.DeepEqual(rawbytes[8:12], []byte{
//			0x49, 0x4C, 0x42, 0x4D,
//		}) {
//			return &FileDescriptor{
//				Description: "IFF Interleaved Bitmap Image",
//				Extensions:  []string{"ilbm", "lbm", "ibm", "iff"},
//			}
//		} else if reflect.DeepEqual(rawbytes[8:12], []byte{
//			0x38, 0x53, 0x56, 0x58,
//		}) {
//			return &FileDescriptor{
//				Description: "IFF 8-Bit Sampled Voice",
//				Extensions:  []string{"8svx", "8sv", "svx", "snd", "iff"},
//			}
//		} else if reflect.DeepEqual(rawbytes[8:12], []byte{
//			0x41, 0x43, 0x42, 0x4D,
//		}) {
//			return &FileDescriptor{
//				Description: "Amiga Contiguous Bitmap",
//				Extensions:  []string{"acbm", "iff"},
//			}
//		} else if reflect.DeepEqual(rawbytes[8:12], []byte{
//			0x41, 0x49, 0x46, 0x46,
//		}) {
//			return &FileDescriptor{
//				Description: "Audio Interchange File Format",
//				Extensions:  []string{"aiff", "aif", "aifc", "snd", "iff"},
//			}
//		}
//	}
//	return nil
//}
