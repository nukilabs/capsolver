package capsolver

// ImageToTextTask defines the parameters for an image-to-text recognition task.
type ImageToTextTask struct {
	Type string `json:"type"`
	// WebsiteURL is the page source URL to improve accuracy.
	WebsiteURL string `json:"websiteURL,omitzero"`
	// Body is the base64-encoded image content without the data URI prefix.
	Body string `json:"body"`
	// Images contains up to 9 base64-encoded images (only for the number module).
	Images []string `json:"images,omitzero"`
	// Module specifies which recognition model to use.
	Module string `json:"module,omitzero"`
	// Score sets the confidence threshold (0.8–1.0) for a match.
	Score float64 `json:"score,omitzero"`
}

// ImageToTextSolution represents the response payload for an ImageToTextTask.
type ImageToTextSolution struct {
	// Text is the recognized text result.
	Text string `json:"text"`
	// Answers contains per-image results when using the number module.
	Answers []string `json:"answers,omitzero"`
}

// SolveImageToText submits an ImageToTextTask and returns the recognized text solution.
func (s *Session) SolveImageToText(task ImageToTextTask) (*ImageToTextSolution, error) {
	task.Type = "ImageToTextTask"
	res, err := s.Solve(task)
	if err != nil {
		return nil, err
	}
	var solution ImageToTextSolution
	if err := res.Unmarshal(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}
