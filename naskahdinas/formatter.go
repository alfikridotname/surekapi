package naskahdinas

type NaskahDinasFormatter struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Template string `json:"template"`
}

func FormatNaskahDinas(nd NaskahDinas) NaskahDinasFormatter {
	ndFormatter := NaskahDinasFormatter{
		ID:       nd.ID,
		Nama:     nd.Nama,
		Template: nd.Template,
	}

	return ndFormatter
}

func FormatMultipleNaskahDinas(nd []NaskahDinas) []NaskahDinasFormatter {
	var ndFormatters []NaskahDinasFormatter
	for _, n := range nd {
		ndFormatter := FormatNaskahDinas(n)
		ndFormatters = append(ndFormatters, ndFormatter)
	}

	return ndFormatters
}
