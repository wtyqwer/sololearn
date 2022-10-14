package main

//
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
//type IFileSignature interface {
//	Handler(rawbytes []byte) *FileDescriptor
//	setNext(IFileSignature) IFileSignature
//}
//
//type sqliteFile struct {
//	next IFileSignature
//}
//
//func (s *sqliteFile) Handler(rawbytes []byte) *FileDescriptor {
//	if reflect.DeepEqual(rawbytes[:16], []byte{
//		0x53, 0x51, 0x4C, 0x69, 0x74, 0x65, 0x20, 0x66,
//		0x6F, 0x72, 0x6D, 0x61, 0x74, 0x20, 0x33, 0x00}) {
//		return &FileDescriptor{
//			Description: "SQLite Database",
//			Extensions:  []string{"sqlitedb", "sqlite", "db"},
//		}
//	}
//	if s.next == nil {
//		return nil
//	}
//	return s.next.Handler(rawbytes)
//}
//func (s *sqliteFile) setNext(fs IFileSignature) IFileSignature {
//	s.next = fs
//	return fs
//}
//
//type threegpFile struct {
//	next IFileSignature
//}
//
//func (s *threegpFile) Handler(rawbytes []byte) *FileDescriptor {
//	if reflect.DeepEqual(rawbytes[4:10], []byte{
//		0x66, 0x74, 0x79, 0x70, 0x33, 0x67}) {
//		return &FileDescriptor{
//			Description: "3rd Generation Partnership Project 3GPP and 3GPP2 multimedia files",
//			Extensions:  []string{"3gp", "3g2"},
//		}
//	}
//	if s.next == nil {
//		return nil
//	}
//	return s.next.Handler(rawbytes)
//}
//func (s *threegpFile) setNext(fs IFileSignature) IFileSignature {
//	s.next = fs
//	return fs
//}
//
//type imageIFFFile struct {
//	next IFileSignature
//}
//
//func (s *imageIFFFile) Handler(rawbytes []byte) *FileDescriptor {
//	if reflect.DeepEqual(rawbytes[:4], []byte{
//		0x46, 0x4F, 0x52, 0x4D}) && reflect.DeepEqual(rawbytes[8:12], []byte{
//		0x49, 0x4C, 0x42, 0x4D}) {
//		return &FileDescriptor{
//			Description: "IFF Interleaved Bitmap Image",
//			Extensions:  []string{"ilbm", "lbm", "ibm", "iff"},
//		}
//	}
//	if s.next == nil {
//		return nil
//	}
//	return s.next.Handler(rawbytes)
//}
//func (s *imageIFFFile) setNext(fs IFileSignature) IFileSignature {
//	s.next = fs
//	return fs
//}
//
//type voiceIFFFile struct {
//	next IFileSignature
//}
//
//func (s *voiceIFFFile) Handler(rawbytes []byte) *FileDescriptor {
//	if reflect.DeepEqual(rawbytes[:4], []byte{
//		0x46, 0x4F, 0x52, 0x4D}) && reflect.DeepEqual(rawbytes[8:12], []byte{
//		0x38, 0x53, 0x56, 0x58}) {
//		return &FileDescriptor{
//			Description: "IFF 8-Bit Sampled Voice",
//			Extensions:  []string{"8svx", "8sv", "svx", "snd", "iff"},
//		}
//	}
//	if s.next == nil {
//		return nil
//	}
//	return s.next.Handler(rawbytes)
//}
//func (s *voiceIFFFile) setNext(fs IFileSignature) IFileSignature {
//	s.next = fs
//	return fs
//}
//
//type bitmapIFFFile struct {
//	next IFileSignature
//}
//
//func (s *bitmapIFFFile) Handler(rawbytes []byte) *FileDescriptor {
//	if reflect.DeepEqual(rawbytes[:4], []byte{
//		0x46, 0x4F, 0x52, 0x4D}) && reflect.DeepEqual(rawbytes[8:12], []byte{
//		0x41, 0x43, 0x42, 0x4D}) {
//		return &FileDescriptor{
//			Description: "Amiga Contiguous Bitmap",
//			Extensions:  []string{"acbm", "iff"},
//		}
//	}
//	if s.next == nil {
//		return nil
//	}
//	return s.next.Handler(rawbytes)
//}
//func (s *bitmapIFFFile) setNext(fs IFileSignature) IFileSignature {
//	s.next = fs
//	return fs
//}
//
//type audioIFFFile struct {
//	next IFileSignature
//}
//
//func (s *audioIFFFile) Handler(rawbytes []byte) *FileDescriptor {
//	if reflect.DeepEqual(rawbytes[:4], []byte{
//		0x46, 0x4F, 0x52, 0x4D}) && reflect.DeepEqual(rawbytes[8:12], []byte{
//		0x41, 0x49, 0x46, 0x46}) {
//		return &FileDescriptor{
//			Description: "Audio Interchange File Format",
//			Extensions:  []string{"aiff", "aif", "aifc", "snd", "iff"},
//		}
//	}
//	if s.next == nil {
//		return nil
//	}
//	return s.next.Handler(rawbytes)
//}
//func (s *audioIFFFile) setNext(fs IFileSignature) IFileSignature {
//	s.next = fs
//	return fs
//}
//
//var (
//	header = new(sqliteFile)
//)
//
//func init() {
//	header.setNext(new(threegpFile)).
//		setNext(new(imageIFFFile)).
//		setNext(new(voiceIFFFile)).
//		setNext(new(bitmapIFFFile)).
//		setNext(new(audioIFFFile))
//}
//
//// TODO:  Refactor this function!!
//func resolveFileSignature(rawbytes []byte) *FileDescriptor {
//	return header.Handler(rawbytes)
//}
