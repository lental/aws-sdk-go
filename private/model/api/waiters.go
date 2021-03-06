// +build codegen

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
)

// A Waiter is an individual waiter definition.
type Waiter struct {
	Name          string
	Delay         int
	MaxAttempts   int
	OperationName string `json:"operation"`
	Operation     *Operation
	Acceptors     []WaitAcceptor
}

// A WaitAcceptor is an individual wait acceptor definition.
type WaitAcceptor struct {
	Expected interface{}
	Matcher  string
	State    string
	Argument string
}

// WaitersGoCode generates and returns Go code for each of the waiters of
// this API.
func (a *API) WaitersGoCode() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "import (\n\t%q\n)",
		"github.com/lental/aws-sdk-go/private/waiter")

	for _, w := range a.Waiters {
		buf.WriteString(w.GoCode())
	}
	return buf.String()
}

// used for unmarshaling from the waiter JSON file
type waiterDefinitions struct {
	*API
	Waiters map[string]Waiter
}

// AttachWaiters reads a file of waiter definitions, and adds those to the API.
// Will panic if an error occurs.
func (a *API) AttachWaiters(filename string) {
	p := waiterDefinitions{API: a}

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		panic(err)
	}

	p.setup()
}

func (p *waiterDefinitions) setup() {
	p.API.Waiters = []Waiter{}
	i, keys := 0, make([]string, len(p.Waiters))
	for k := range p.Waiters {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for _, n := range keys {
		e := p.Waiters[n]
		n = p.ExportableName(n)
		e.Name = n
		e.OperationName = p.ExportableName(e.OperationName)
		e.Operation = p.API.Operations[e.OperationName]
		if e.Operation == nil {
			panic("unknown operation " + e.OperationName + " for waiter " + n)
		}
		p.API.Waiters = append(p.API.Waiters, e)
	}
}

// ExpectedString returns the string that was expected by the WaitAcceptor
func (a *WaitAcceptor) ExpectedString() string {
	switch a.Expected.(type) {
	case string:
		return fmt.Sprintf("%q", a.Expected)
	default:
		return fmt.Sprintf("%v", a.Expected)
	}
}

var waiterTmpls = template.Must(template.New("waiterTmpls").Parse(`
{{ define "docstring" -}}
// WaitUntil{{ .Name }} uses the {{ .Operation.API.NiceName }} API operation
// {{ .OperationName }} to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
{{- end }}

{{ define "waiter" }}
{{ template "docstring" . }}
func (c *{{ .Operation.API.StructName }}) WaitUntil{{ .Name }}(input {{ .Operation.InputRef.GoType }}) error {
	waiterCfg  := waiter.Config{
		Operation:   "{{ .OperationName }}",
		Delay:       {{ .Delay }},
		MaxAttempts: {{ .MaxAttempts }},
		Acceptors: []waiter.WaitAcceptor{
			{{ range $_, $a := .Acceptors }}waiter.WaitAcceptor{
				State:    "{{ .State }}",
				Matcher:  "{{ .Matcher }}",
				Argument: "{{ .Argument }}",
				Expected: {{ .ExpectedString }},
			},
			{{ end }}
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
{{- end }}

{{ define "waiter interface" }}
WaitUntil{{ .Name }}({{ .Operation.InputRef.GoTypeWithPkgName }}) error
{{- end }}
`))

// InterfaceSignature returns a string representing the Waiter's interface
// function signature.
func (w *Waiter) InterfaceSignature() string {
	var buf bytes.Buffer
	if err := waiterTmpls.ExecuteTemplate(&buf, "waiter interface", w); err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

// GoCode returns the generated Go code for an individual waiter.
func (w *Waiter) GoCode() string {
	var buf bytes.Buffer
	if err := waiterTmpls.ExecuteTemplate(&buf, "waiter", w); err != nil {
		panic(err)
	}

	return buf.String()
}
