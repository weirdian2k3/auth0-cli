package display

import (
	"strings"

	"github.com/auth0/auth0-cli/internal/ansi"
	"github.com/logrusorgru/aurora"

	"gopkg.in/auth0.v5/management"
)

const (
	logTimeFormat = "2006-01-02T15:04:05+00:00"
	notApplicable = "N/A"
)

var _ View = &logView{}

type logView struct {
	*management.Log
}

func (v *logView) AsTableHeader() []string {
	return []string{"Type", "Description", "Date", "Connection", "Client"}

}

func (v *logView) getConnection() string {
	if v.Details["prompts"] == nil {
		return notApplicable
	}

	prompt, ok := v.Details["prompts"].([]interface{})
	if ok && len(prompt) > 0 {
		dict, ok := prompt[0].(map[string]interface{})
		if ok {
			return dict["connection"].(string)
		} else {
			return notApplicable
		}
	} else {
		return notApplicable
	}

	return notApplicable
}

func (v *logView) AsTableRow() []string {
	typ, desc := typeDescFor(v.Log, false)

	clientName := v.GetClientName()
	if clientName == "" {
		clientName = "N/A"
	}

	return []string{
		typ,
		desc,
		ansi.Faint(timeAgo(v.GetDate())),
		v.getConnection(),
		clientName,
	}
}

func typeDescFor(l *management.Log, noColor bool) (typ, desc string) {
	chunks := strings.Split(l.TypeName(), "(")
	typ = chunks[0]

	if len(chunks) == 2 {
		desc = strings.TrimSuffix(chunks[1], ")")
	}

	if !noColor {
		// colorize the event type field based on whether it's a success or failure
		if strings.HasPrefix(l.GetType(), "s") {
			typ = aurora.Green(typ).String()
		} else if strings.HasPrefix(l.GetType(), "f") {
			typ = aurora.BrightRed(typ).String()
		} else if strings.HasPrefix(l.GetType(), "w") {
			typ = aurora.BrightYellow(typ).String()
		}
	}

	return typ, desc
}

func (r *Renderer) LogList(logs []*management.Log, noColor bool) {
	var res []View
	for _, l := range logs {
		res = append(res, &logView{Log: l})
	}

	r.Results(res)
}
