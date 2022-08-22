package main

import "fmt"

var HttpCodeMap = map[int]string{
	100: "Continue",
	101: "Switching Protocol",
	102: "Processing",
	103: "Early Hints",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information\"",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	207: "Multi-Status",
	208: "Already Reported",
	226: "IM Used",
	300: "Multiple Choice",
	301: "Moved Permanently",
	302: "Found",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	306: "Switch Proxy",
	307: "Temporary Redirect",
	308: "Permanent Redirect",
}

func main() {
	{
		code := 200
		fmt.Printf("code [%d], message: %s\n", code, resolveHttpStatusCode(code))
	}
	{
		code := 226
		fmt.Printf("code [%d], message: %s\n", code, resolveHttpStatusCode(code))
	}
}

// TODO: 🔨 Refactor this function!!
func resolveHttpStatusCode(code int) string {
	if msg, ok := HttpCodeMap[code]; ok {

		return msg
	}
	return ""
}

/*
// TODO: 🔨 Refactor this function!!
func resolveHttpStatusCode(code int) string {
    if code == 100 {
        return "Continue"
    } else if code == 101 {
        return "Switching Protocol"
    } else if code == 102 {
        return "Processing"
    } else if code == 103 {
        return "Early Hints"
    } else if code == 200 {
        return "OK"
    } else if code == 201 {
        return "Created"
    } else if code == 202 {
        return "Accepted"
    } else if code == 203 {
        return "Non-Authoritative Information"
    } else if code == 204 {
        return "No Content"
    } else if code == 205 {
        return "Reset Content"
    } else if code == 206 {
        return "Partial Content"
    } else if code == 207 {
        return "Multi-Status"
    } else if code == 208 {
        return "Already Reported"
    } else if code == 226 {
        return "IM Used"
    } else if code == 300 {
        return "Multiple Choice"
    } else if code == 301 {
        return "Moved Permanently"
    } else if code == 302 {
        return ""
    } else if code == 303 {
        return "See Other"
    } else if code == 304 {
        return "Not Modified"
    } else if code == 305 {
        return "Use Proxy"
    } else if code == 306 {
        return "Switch Proxy"
    } else if code == 307 {
        return "Temporary Redirect"
    } else if code == 308 {
        return "Permanent Redirect"
    }
    return ""
}


*/
