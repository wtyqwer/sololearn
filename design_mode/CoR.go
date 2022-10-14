package main

import (
	"fmt"
	"reflect"
)

/* Resolve file signatures
 *
 * resource:
 * 1. https://en.wikipedia.org/wiki/List_of_file_signatures
 */

//resolve:https://refactoringguru.cn/design-patterns/chain-of-responsibility/go/example

//责任链模式   chain of responsibility    亦称： 职责链模式、命令链、CoR、Chain of Command、Chain of Responsibility
func main() {
	{
		rawbytes := []byte{
			0x53, 0x51, 0x4C, 0x69, 0x74, 0x65, 0x20, 0x66,
			0x6F, 0x72, 0x6D, 0x61, 0x74, 0x20, 0x33, 0x00,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		}
		fd := resolveFileSignature(rawbytes)
		if fd != nil {
			fmt.Printf("%v: %s\n", fd.Extensions, fd.Description)
		} else {
			fmt.Printf("unknown file signature\n")
		}
	}
	{
		rawbytes := []byte{
			0x53, 0x51, 0x4C, 0x69,
			0x66, 0x74, 0x79, 0x70, 0x33, 0x67,
			0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		}
		fd := resolveFileSignature(rawbytes)
		if fd != nil {
			fmt.Printf("%v: %s\n", fd.Extensions, fd.Description)
		} else {
			fmt.Printf("unknown file signature\n")
		}
	}
	{
		rawbytes := []byte{
			0x46, 0x4F, 0x52, 0x4D,
			0x01, 0x01, 0x01, 0x01,
			0x41, 0x43, 0x42, 0x4D,
			0x01, 0x01, 0x01, 0x01,
		}
		fd := resolveFileSignature(rawbytes)
		if fd != nil {
			fmt.Printf("%v: %s\n", fd.Extensions, fd.Description)
		} else {
			fmt.Printf("unknown file signature\n")
		}
	}
	{
		rawbytes := []byte{
			0x46, 0x4F, 0x52, 0x4D,
			0x01, 0x01, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01,
			0x01, 0x01, 0x01, 0x01,
		}
		fd := resolveFileSignature(rawbytes)
		if fd != nil {
			fmt.Printf("%v: %s\n", fd.Extensions, fd.Description)
		} else {
			fmt.Printf("unknown file signature\n")
		}
	}
}

type FileDescriptor struct {
	Description string
	Extensions  []string
}

type CaseChain interface {
	HandlerReq(req []byte) *FileDescriptor //处理请求
	setNextCase(next CaseChain) CaseChain  //设置链表的下一节点 并返回
}

type Sqlite struct {
	next CaseChain
}

func (s *Sqlite) HandlerReq(req []byte) *FileDescriptor {
	if reflect.DeepEqual(req[:16], []byte{
		0x53, 0x51, 0x4C, 0x69, 0x74, 0x65, 0x20, 0x66,
		0x6F, 0x72, 0x6D, 0x61, 0x74, 0x20, 0x33, 0x00,
	}) {
		return &FileDescriptor{
			Description: "Sqlite Database",
			Extensions:  []string{"Sqlitedb", "Sqlite", "db"},
		}
	}
	//链上没有下一个条件了
	if s.next == nil {
		return nil
	}

	return s.next.HandlerReq(req)
}

func (s *Sqlite) setNextCase(next CaseChain) CaseChain {
	s.next = next
	return next
}

type ThreeGP struct {
	next CaseChain
}

func (s *ThreeGP) HandlerReq(req []byte) *FileDescriptor {
	if reflect.DeepEqual(req[4:10], []byte{
		0x66, 0x74, 0x79, 0x70, 0x33, 0x67,
	}) {
		return &FileDescriptor{
			Description: "3rd Generation Partnership Project 3GPP and 3GPP2 multimedia files",
			Extensions:  []string{"3gp", "3g2"},
		}
	}
	//链上没有下一个条件了
	if s.next == nil {
		return nil
	}

	return s.next.HandlerReq(req)
}
func (s *ThreeGP) setNextCase(next CaseChain) CaseChain {
	s.next = next
	return next
}

type Ilbm struct {
	next CaseChain
}

func (s *Ilbm) HandlerReq(req []byte) *FileDescriptor {
	if reflect.DeepEqual(req[4:10], []byte{
		0x66, 0x74, 0x79, 0x70, 0x33, 0x67,
	}) {
		return &FileDescriptor{
			Description: "3rd Generation Partnership Project 3GPP and 3GPP2 multimedia files",
			Extensions:  []string{"3gp", "3g2"},
		}
	}
	//链上没有下一个条件了
	if s.next == nil {
		return nil
	}

	return s.next.HandlerReq(req)
}

func (s *Ilbm) setNextCase(next CaseChain) CaseChain {
	s.next = next
	return next
}

type ChildChain struct {
	next CaseChain
}

func (s *ChildChain) HandlerReq(req []byte) *FileDescriptor {
	if reflect.DeepEqual(req[:4], []byte{
		0x46, 0x4F, 0x52, 0x4D,
	}) && s.next != nil {
		return s.next.HandlerReq(req)
	}
	//链上没有下一个条件了
	return nil
}

func (s *ChildChain) setNextCase(next CaseChain) CaseChain {
	s.next = next
	return next
}

type EightBit struct {
	next CaseChain
}

func (s *EightBit) HandlerReq(req []byte) *FileDescriptor {
	if reflect.DeepEqual(req[8:12], []byte{
		0x38, 0x53, 0x56, 0x58,
	}) {
		return &FileDescriptor{
			Description: "IFF 8-Bit Sampled Voice",
			Extensions:  []string{"8svx", "8sv", "svx", "snd", "iff"},
		}
	}
	//链上没有下一个条件了
	if s.next == nil {
		return nil
	}

	return s.next.HandlerReq(req)
}
func (s *EightBit) setNextCase(next CaseChain) CaseChain {
	s.next = next
	return next
}

type Amiga struct {
	next CaseChain
}

func (s *Amiga) HandlerReq(req []byte) *FileDescriptor {
	if reflect.DeepEqual(req[8:12], []byte{
		0x41, 0x43, 0x42, 0x4D,
	}) {
		return &FileDescriptor{
			Description: "Amiga Contiguous Bitmap",
			Extensions:  []string{"acbm", "iff"},
		}
	}
	//链上没有下一个条件了
	if s.next == nil {
		return nil
	}

	return s.next.HandlerReq(req)
}
func (s *Amiga) setNextCase(next CaseChain) CaseChain {
	s.next = next
	return next
}

type Audio struct {
	next CaseChain
}

func (s *Audio) HandlerReq(req []byte) *FileDescriptor {
	if reflect.DeepEqual(req[8:12], []byte{
		0x41, 0x49, 0x46, 0x46,
	}) {
		return &FileDescriptor{
			Description: "Audio Interchange File Format",
			Extensions:  []string{"aiff", "aif", "aifc", "snd", "iff"},
		}
	}
	//链上没有下一个条件了
	if s.next == nil {
		return nil
	}

	return s.next.HandlerReq(req)
}
func (s *Audio) setNextCase(next CaseChain) CaseChain {
	s.next = next
	return next
}

var chainHeader = new(Sqlite) //设置头节点

func init() {
	//初始化责任链
	chainHeader.setNextCase(new(ThreeGP)).
		setNextCase(new(Ilbm)).
		setNextCase(new(ChildChain)).
		setNextCase(new(EightBit)).
		setNextCase(new(Amiga)).
		setNextCase(new(Audio))

}

// TODO: 🔨 Refactor this function!!
func resolveFileSignature(rawbytes []byte) *FileDescriptor {
	return chainHeader.HandlerReq(rawbytes)
}
