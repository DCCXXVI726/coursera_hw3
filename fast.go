package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
        jlexer "github.com/mailru/easyjson/jlexer"
        jwriter "github.com/mailru/easyjson/jwriter"

	// "log"
)

type User struct {
	Browsers []string `json:"browsers"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
}


const filePath1 string = "./data/users.txt"

var pattern1 = regexp.MustCompile("Android")
var pattern2 = regexp.MustCompile("MSIE")
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath1)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile("@")
	uniqueBrowsers := 0
	foundUsers := ""

	lines := strings.Split(string(fileContents), "\n")
	seenBrowsers := make(map[string]bool)
	i := -1
	for _, line := range lines {
		i++
		user := User{}
		err := user.UnmarshalJSON([]byte(line))
		if err != nil {
			panic(err)
		}
		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {
			if ok := pattern1.MatchString(browser); ok {
				isAndroid = true
				if ok, _ := seenBrowsers[browser]; !ok {
					seenBrowsers[browser] = true
					uniqueBrowsers++
				}
			}
			if ok := pattern2.MatchString(browser); ok {
				isMSIE = true
				if ok, _ := seenBrowsers[browser]; !ok {
                                        seenBrowsers[browser] = true
                                        uniqueBrowsers++
                                }
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := r.ReplaceAllString(user.Email, " [at] ")
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email)
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", uniqueBrowsers)
}

func easyjson9e1087fdDecodeTemp(in *jlexer.Lexer, out *User) {
        isTopLevel := in.IsStart()
        if in.IsNull() {
                if isTopLevel {
                        in.Consumed()
                }
                in.Skip()
                return
        }
        in.Delim('{')
        for !in.IsDelim('}') {
                key := in.UnsafeFieldName(false)
                in.WantColon()
                if in.IsNull() {
                        in.Skip()
                        in.WantComma()
                        continue
                }
                switch key {
                case "browsers":
                        if in.IsNull() {
                                in.Skip()
                                out.Browsers = nil
                        } else {
                                in.Delim('[')
                                if out.Browsers == nil {
                                        if !in.IsDelim(']') {
                                                out.Browsers = make([]string, 0, 4)
                                        } else {
                                                out.Browsers = []string{}
                                        }
                                } else {
                                        out.Browsers = (out.Browsers)[:0]
                                }
                                for !in.IsDelim(']') {
                                        var v1 string
                                        v1 = string(in.String())
                                        out.Browsers = append(out.Browsers, v1)
                                        in.WantComma()
                                }
                                in.Delim(']')
                        }
                case "email":
                        out.Email = string(in.String())
                case "name":
                        out.Name = string(in.String())
                default:
                        in.SkipRecursive()
                }
                in.WantComma()
        }
       in.Delim('}')
        if isTopLevel {
                in.Consumed()
        }
}
func easyjson9e1087fdEncodeTemp(out *jwriter.Writer, in User) {
        out.RawByte('{')
        first := true
        _ = first
        {
                const prefix string = ",\"browsers\":"
                out.RawString(prefix[1:])
                if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
                        out.RawString("null")
                } else {
                        out.RawByte('[')
                        for v2, v3 := range in.Browsers {
                                if v2 > 0 {
                                        out.RawByte(',')
                                }
                                out.String(string(v3))
                        }
                        out.RawByte(']')
                }
        }
        {
                const prefix string = ",\"email\":"
                out.RawString(prefix)
                out.String(string(in.Email))
        }
        {
                const prefix string = ",\"name\":"
                out.RawString(prefix)
                out.String(string(in.Name))
        }
        out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
        w := jwriter.Writer{}
        easyjson9e1087fdEncodeTemp(&w, v)
        return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
        easyjson9e1087fdEncodeTemp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
        r := jlexer.Lexer{Data: data}
        easyjson9e1087fdDecodeTemp(&r, v)
        return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
        easyjson9e1087fdDecodeTemp(l, v)
}

