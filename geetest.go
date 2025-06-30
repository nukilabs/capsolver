package capsolver

// GeeTestTaskProxyless defines a Geetest (v3/v4) solve task.
type GeeTestTask struct {
	Type string `json:"type"`
	// WebsiteURL is the page URL where Geetest is used.
	WebsiteURL string `json:"websiteURL"`
	// GT is the Geetest v3 site key.
	GT string `json:"gt,omitzero"`
	// Challenge is the Geetest v3 challenge token.
	Challenge string `json:"challenge,omitzero"`
	// CaptchaID is the Geetest v4 captcha identifier.
	CaptchaID string `json:"captchaId,omitzero"`
	// GeetestAPIServerSubdomain overrides the API subdomain (e.g. "api.geetest.com").
	GeetestAPIServerSubdomain string `json:"geetestApiServerSubdomain,omitzero"`
}

// GeeTestSolution holds the solve result for both Geetest v3 and v4.
type GeeTestSolution struct {
	// Challenge echoes back the challenge token (v3).
	Challenge string `json:"challenge,omitzero"`
	// Validate is the validation token to submit (v3).
	Validate string `json:"validate,omitzero"`
	// CaptchaID echoes back the captcha_id from the challenge (v4).
	CaptchaID string `json:"captcha_id,omitzero"`
	// CaptchaOutput is the captcha_output token (v4).
	CaptchaOutput string `json:"captcha_output,omitzero"`
	// GenTime is the generation timestamp (v4).
	GenTime string `json:"gen_time,omitzero"`
	// LotNumber is the lot_number value (v4).
	LotNumber string `json:"lot_number,omitzero"`
	// PassToken is the pass_token value (v4).
	PassToken string `json:"pass_token,omitzero"`
	// RiskType indicates the risk_type (e.g., "slide") (v4).
	RiskType string `json:"risk_type,omitzero"`
}

// SolveGeeTest submits a GeeTestTask and returns the solution.
func (s *Session) SolveGeeTest(task GeeTestTask) (*GeeTestSolution, error) {
	task.Type = "GeeTestTaskProxyLess"
	res, err := s.Solve(task)
	if err != nil {
		return nil, err
	}
	var solution GeeTestSolution
	if err := res.Unmarshal(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}
