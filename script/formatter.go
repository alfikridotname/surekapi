package script

type ScriptFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"nama"`
	Template string `json:"template"`
}

func FormatScript(nd Script) ScriptFormatter {
	ndFormatter := ScriptFormatter{
		ID:       nd.ID,
		Name:     nd.Name,
		Template: nd.Template,
	}

	return ndFormatter
}

func FormatMultipleScript(nd []Script) []ScriptFormatter {
	var scriptFormatters []ScriptFormatter
	for _, n := range nd {
		ndFormatter := FormatScript(n)
		scriptFormatters = append(scriptFormatters, ndFormatter)
	}

	return scriptFormatters
}
