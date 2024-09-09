package main

type MassProducerParams struct {
	XlsxPaths           []string `json:"xlsx_paths"`
	DrawingPaths        []string `json:"drawing_paths"`
	VersionNames        []string `json:"version_names"`
	SizeRatio           int      `json:"size_ratio"`
	IsPrintingVersion   bool     `json:"is_printing_version"`
	Overwrite           bool     `json:"overwrite"`
	CommentOverwrite    string   `json:"@overwrite"`
	NewCardsOnly        bool     `json:"new_cards_only"`
	CommentNewCardsOnly string   `json:"@new_cards_only"`
	GeneralPath         string   `json:"general_path"`
	FontPath            string   `json:"font_path"`
}
