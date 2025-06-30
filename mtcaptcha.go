package capsolver

// MtCaptchaTask defines an MTCaptcha token solve task.
type MtCaptchaTask struct {
	Type string `json:"type"`
	// WebsiteURL is the URL of the page protected by MtCaptcha.
	WebsiteURL string `json:"websiteURL"`
	// WebsiteKey is the public domain key for MtCaptcha (e.g. "MTPublic-xxx").
	WebsiteKey string `json:"websiteKey"`
	// Proxy specifies your proxy (e.g., "http://ip:port:user:pass").
	Proxy string `json:"proxy,omitzero"`
}

// MtCaptchaSolution represents the result returned by getTaskResult for MTCaptcha.
type MtCaptchaSolution struct {
	// Token is the captcha token to submit to the target site.
	Token string `json:"token"`
}

// SolveMtCaptcha submits an MtCaptchaTask and returns the solution.
func (s *Session) SolveMtCaptcha(task MtCaptchaTask) (*MtCaptchaSolution, error) {
	if task.Proxy != "" {
		task.Type = "MtCaptchaTask"
	} else {
		task.Type = "MtCaptchaTaskProxyLess"
	}
	res, err := s.Solve(task)
	if err != nil {
		return nil, err
	}
	var solution MtCaptchaSolution
	if err := res.Unmarshal(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}
