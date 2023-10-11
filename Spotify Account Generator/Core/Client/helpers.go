package client

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

var (
	httpStatusCodes = map[int]string{100: "Continue", 101: "Switching Protocols", 102: "Processing", 103: "Early Hints", 200: "OK", 201: "Created", 202: "Accepted", 203: "Non-Authoritative Information", 204: "No Content", 205: "Reset Content", 206: "Partial Content", 207: "Multi-Status", 208: "Already Reported", 226: "IM Used", 300: "Multiple Choices", 301: "Moved Permanently", 302: "Found", 303: "See Other", 304: "Not Modified", 305: "Use Proxy", 306: "(Unused)", 307: "Temporary Redirect", 308: "Permanent Redirect", 400: "Bad Request", 401: "Unauthorized", 402: "Payment Required", 403: "Forbidden", 404: "Not Found", 405: "Method Not Allowed", 406: "Not Acceptable", 407: "Proxy Authentication Required", 408: "Request Timeout", 409: "Conflict", 410: "Gone", 411: "Length Required", 412: "Precondition Failed", 413: "Payload Too Large", 414: "URI Too Long", 415: "Unsupported Media Type", 416: "Range Not Satisfiable", 417: "Expectation Failed", 418: "I'm a teapot", 421: "Misdirected Request", 422: "Unprocessable Entity", 423: "Locked", 424: "Failed Dependency", 425: "Too Early", 426: "Upgrade Required", 428: "Precondition Required", 429: "Too Many Requests", 431: "Request Header Fields Too Large", 451: "Unavailable For Legal Reasons", 500: "Internal Server Error", 501: "Not Implemented", 502: "Bad Gateway", 503: "Service Unavailable", 504: "Gateway Timeout", 505: "HTTP Version Not Supported", 506: "Variant Also Negotiates", 507: "Insufficient Storage", 508: "Loop Detected", 510: "Not Extended", 511: "Network Authentication Required"}
	NoHeaders       = map[string]string{}
)

func AddHeaders(common bool, headers map[string]string, req *http.Request) {
	if common {
		for k, v := range map[string]string{
			"authority":                 "www.spotify.com",
			"accept":                    "*/*",
			"accept-language":           "tr-TR",
			"dnt":                       "1",
			"referer":                   "https://www.google.com/",
			"sec-ch-ua":                 `"Google Chrome";v="112", "Not;A=Brand";v="8", "Chromium";v="112"`,
			"sec-ch-ua-mobile":          "?0",
			"sec-ch-ua-platform":        `"Windows"`,
			"sec-fetch-dest":            "document",
			"sec-fetch-mode":            "navigate",
			"sec-fetch-site":            "cross-site",
			"sec-fetch-user":            "?1",
			"upgrade-insecure-requests": "1",
			"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
		} {
			req.Header.Set(k, v)
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func Ok(n int) bool {
	for _, v := range []int{200, 201, 204} {
		if v == n {
			return true
		}
	}

	return false
}

func Request(method, url string, body io.Reader, headers map[string]string, common bool, c tls_client.HttpClient) (Response RequestResponse) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		Response.Error = err
		return
	}

	req.Close = true

	AddHeaders(common, headers, req)

	resp, err := c.Do(req)

	if err != nil {
		Response.Error = err
		return
	}

	Response.Ok = Ok(resp.StatusCode)
	Response.StatusCode = resp.StatusCode
	Response.Request = resp
	Response.StatusCodeDefinition = httpStatusCodes[resp.StatusCode]

	defer resp.Body.Close()

	if Response.Body, err = io.ReadAll(resp.Body); err != nil {
		Response.Error = err
		return
	}

	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		if err := json.NewDecoder(bytes.NewBuffer(Response.Body)).Decode(&Response.Json); err != nil {
			if strings.Contains(err.Error(), "unmarshal array into Go value of type map[string]interface {}") {
				return
			}

			Response.Error = err
			return
		}

	}
	return
}
