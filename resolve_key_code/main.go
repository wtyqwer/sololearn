package main

/* Resolve input keycodes
 *
 * resource:
 * 1. KeyboardEvent: https://css-tricks.com/snippets/javascript/javascript-keycodes/
 */

//import "fmt"
//
//const (
//	KEY_C   = 67
//	KEY_H   = 72
//	KEY_V   = 86
//	KEY_X   = 88
//	KEY_F1  = 112
//	KEY_F11 = 122
//)
//
//var keyboardCommand = map[KeyEvent]string{
//	KeyEvent{
//		ControlKey: true,
//		KeyCode: KEY_C,
//	}: "command_copy",
//	KeyEvent{
//		ControlKey: true,
//		KeyCode: KEY_V,
//	}: "command_paste",
//	KeyEvent{
//		ControlKey: true,
//		KeyCode: KEY_X,
//	}: "command_cut",
//	KeyEvent{
//		ControlKey: true,
//		KeyCode: KEY_F11,
//	}: "command_full_screen",
//	KeyEvent{
//		AltKey:  true,
//		KeyCode: KEY_X,
//	}: "command_close",
//	KeyEvent{
//		AltKey:  true,
//		KeyCode: KEY_H,
//	}: "command_help",
//	KeyEvent{
//		KeyCode: KEY_F1,
//	}: "command_help",
//}
//
//func main() {
//
//	{
//		e := KeyEvent{
//			KeyCode: KEY_H,
//			AltKey:  true,
//		}
//		command := resolveKeyEvent(e)
//		fmt.Printf("%-20s: KeyEvent %+v\n", command, e)
//	}
//	{
//		e := KeyEvent{
//			KeyCode:    KEY_C,
//			ControlKey: true,
//		}
//		command := resolveKeyEvent(e)
//		fmt.Printf("%-20s: KeyEvent %+v\n", command, e)
//	}
//	{
//		e := KeyEvent{
//			KeyCode:    KEY_F11,
//			ControlKey: false,
//			AltKey:     false,
//		}
//		command := resolveKeyEvent(e)
//		fmt.Printf("%-20s: KeyEvent %+v\n", command, e)
//	}
//}
//
//type KeyEvent struct {
//	KeyCode    int
//	ShiftKey   bool
//	AltKey     bool
//	ControlKey bool
//}
//
////func resolveKeyEvent(ev KeyEvent) (action string) {
////	action = keyboardCommand[ev]
////	if action == "" {
////		return "unknown"
////	}
////	return action
////}
//
//// TODO: 🔨 Refactor this function!!
//func resolveKeyEvent(ev KeyEvent) (action string) {
//	if ev.ControlKey && !ev.AltKey {
//		if ev.KeyCode == KEY_C {
//			return "command_copy"
//		} else if ev.KeyCode == KEY_V {
//			return "command_paste"
//		} else if ev.KeyCode == KEY_X {
//			return "command_cut"
//		} else if ev.KeyCode == KEY_F11 {
//			return "command_full_screen"
//		}
//	} else if ev.AltKey && !ev.ControlKey {
//		if ev.KeyCode == KEY_X {
//			return "command_close"
//		} else if ev.KeyCode == KEY_H {
//			return "command_help"
//		}
//	} else {
//		if ev.KeyCode == KEY_F1 {
//			return "command_help"
//		}
//	}
//	return "unknown"
//}
