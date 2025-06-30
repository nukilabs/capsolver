package capsolver

// AntiTurnstileTask defines a Cloudflare Turnstile solve task.
type AntiTurnstileTask struct {
	Type string `json:"type"`
	// WebsiteURL is the address of the target page.
	WebsiteURL string `json:"websiteURL"`
	// WebsiteKey is the Turnstile website key.
	WebsiteKey string `json:"websiteKey"`
	// Metadata holds extra Turnstile data (e.g., "action", "cdata").
	Metadata map[string]string `json:"metadata,omitzero"`
}

// AntiTurnstileSolution represents the solve result for a Turnstile task.
type AntiTurnstileSolution struct {
	// Token is the captcha token to submit to the target site.
	Token string `json:"token"`
	// Type indicates the type of cloudflare task solved.
	Type string `json:"type"`
	// UserAgent is the User-Agent string used during solving.
	UserAgent string `json:"userAgent,omitzero"`
}

// SolveAntiTurnstile submits an AntiTurnstileTaskProxyLess and returns the solution.
func (s *Session) SolveAntiTurnstile(task AntiTurnstileTask) (*AntiTurnstileSolution, error) {
	task.Type = "AntiTurnstileTaskProxyLess"
	res, err := s.createTask(task)
	if err != nil {
		return nil, err
	}
	var solution AntiTurnstileSolution
	if err := res.Unmarshal(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}
