package capsolver

// AwsWafClassificationTask defines the parameters for an AWS WAF image-classification task.
type AwsWafClassificationTask struct {
	Type string `json:"type"` // always "AwsWafClassification"
	// WebsiteURL is the page source URL to improve accuracy.
	WebsiteURL string `json:"websiteURL,omitzero"`
	// Images contains 1 base64-encoded image (or up to 9 for grid puzzles).
	Images []string `json:"images"`
	// Question is the classification code (e.g. "aws:toycarcity:carcity" or "aws:grid:bed").
	Question string `json:"question"`
}

// AwsWafClassificationSolution holds the result returned synchronously by createTask.
type AwsWafClassificationSolution struct {
	// Box is the [x, y] coordinate for point-based puzzles (e.g. toycarcity).
	Box []float64 `json:"box,omitzero"`
	// Objects are the indexes of matching tiles for grid puzzles.
	Objects []int `json:"objects,omitzero"`
	// Distance is the slide distance for slider puzzles (if supported).
	Distance float64 `json:"distance,omitzero"`
}

// SolveAwsWafClassification submits an AwsWafClassificationTask and returns the solution.
func (s *Session) SolveAwsWafClassification(task AwsWafClassificationTask) (*AwsWafClassificationSolution, error) {
	task.Type = "AwsWafClassification"
	res, err := s.createTask(task)
	if err != nil {
		return nil, err
	}
	var solution AwsWafClassificationSolution
	if err := res.Unmarshal(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}

// AntiAwsWafTask defines an AWS WAF captcha solve task.
type AntiAwsWafTask struct {
	Type string `json:"type"`
	// Proxy is the proxy string (e.g., "http://ip:port:user:pass").
	Proxy string `json:"proxy,omitzero"`
	// WebsiteURL is the URL of the page returning the captcha challenge.
	WebsiteURL string `json:"websiteURL"`
	// AwsKey is the `key` value extracted from the captcha page.
	AwsKey string `json:"awsKey,omitzero"`
	// AwsIv is the `iv` value extracted from the captcha page.
	AwsIv string `json:"awsIv,omitzero"`
	// AwsContext is the `context` value extracted from the captcha page.
	AwsContext string `json:"awsContext,omitzero"`
	// AwsChallengeJS is the `challenge.js` link returned by the captcha page.
	AwsChallengeJS string `json:"awsChallengeJS,omitzero"`
	// AwsProblemURL is the `problem` endpoint URL containing visualSolutionsRequired, etc.
	AwsProblemURL string `json:"awsProblemUrl,omitzero"`
}

// AntiAwsWafSolution represents the solve result for an AWS WAF captcha task.
type AntiAwsWafSolution struct {
	// Cookie is the DataDome cookie string to use for subsequent requests.
	Cookie string `json:"cookie"`
}

// SolveAntiAwsWaf submits an AntiAwsWafTask and returns the solution.
func (s *Session) SolveAntiAwsWaf(task AntiAwsWafTask) (*AntiAwsWafSolution, error) {
	if task.Proxy != "" {
		task.Type = "AntiAwsWafTask"
	} else {
		task.Type = "AntiAwsWafTaskProxyLess"
	}
	res, err := s.createTask(task)
	if err != nil {
		return nil, err
	}
	var solution AntiAwsWafSolution
	if err := res.Unmarshal(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}
