package capsolver

// DataDomeSliderTask defines the parameters for a DataDome slider/interstitial task.
type DataDomeSliderTask struct {
	Type string `json:"type"`
	// CaptchaURL is the URL to the DataDome captcha; must include t=fe.
	CaptchaURL string `json:"captchaUrl"`
	// UserAgent must match the UA used when requesting the target site.
	UserAgent string `json:"userAgent"`
	// Proxy is the proxy string (e.g., "http:ip:port:user:pass").
	Proxy string `json:"proxy"`
}

// DataDomeSolution represents the solution returned by getTaskResult for DataDome.
type DataDomeSolution struct {
	// UserAgent echoes back the User-Agent used during solving.
	UserAgent string `json:"userAgent"`
	// Cookie is the DataDome cookie string to use for subsequent requests.
	Cookie string `json:"cookie"`
}

// SolveDataDomeSlider submits a DataDomeSliderTask and returns the solution.
func (s *Session) SolveDataDomeSlider(task DataDomeSliderTask) (*DataDomeSolution, error) {
	task.Type = "DatadomeSliderTask"
	res, err := s.createTask(task)
	if err != nil {
		return nil, err
	}
	var solution DataDomeSolution
	if err := res.Unmarshal(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}
