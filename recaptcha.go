package capsolver

// ReCaptchaQuestion represents a reCAPTCHA v2 classification question code.
type ReCaptchaQuestion string

const (
	ReCaptchaQuestionTaxis            ReCaptchaQuestion = "/m/0pg52"
	ReCaptchaQuestionBus              ReCaptchaQuestion = "/m/01bjv"
	ReCaptchaQuestionSchoolBus        ReCaptchaQuestion = "/m/02yvhj"
	ReCaptchaQuestionMotorcycles      ReCaptchaQuestion = "/m/04_sv"
	ReCaptchaQuestionTractors         ReCaptchaQuestion = "/m/013xlm"
	ReCaptchaQuestionChimneys         ReCaptchaQuestion = "/m/01jk_4"
	ReCaptchaQuestionCrosswalks       ReCaptchaQuestion = "/m/014xcs"
	ReCaptchaQuestionTrafficLights    ReCaptchaQuestion = "/m/015qff"
	ReCaptchaQuestionBicycles         ReCaptchaQuestion = "/m/0199g"
	ReCaptchaQuestionParkingMeters    ReCaptchaQuestion = "/m/015qbp"
	ReCaptchaQuestionCars             ReCaptchaQuestion = "/m/0k4j"
	ReCaptchaQuestionBridges          ReCaptchaQuestion = "/m/015kr"
	ReCaptchaQuestionBoats            ReCaptchaQuestion = "/m/019jd"
	ReCaptchaQuestionPalmTrees        ReCaptchaQuestion = "/m/0cdl1"
	ReCaptchaQuestionMountainsOrHills ReCaptchaQuestion = "/m/09d_r"
	ReCaptchaQuestionFireHydrant      ReCaptchaQuestion = "/m/01pns0"
	ReCaptchaQuestionStairs           ReCaptchaQuestion = "/m/01lynh"
)

// ReCaptchaV2ClassificationTask defines a reCAPTCHA v2 image-classification task.
type ReCaptchaV2ClassificationTask struct {
	Type string `json:"type"`
	// WebsiteURL is the page source URL to improve accuracy.
	WebsiteURL string `json:"websiteURL,omitzero"`
	// WebsiteKey is the site key to improve accuracy.
	WebsiteKey string `json:"websiteKey,omitzero"`
	// Image is the base64-encoded image content without the data URI prefix.
	Image string `json:"image"`
	// Question is the classification question code.
	Question ReCaptchaQuestion `json:"question"`
}

// ReCaptchaType represents the type of reCAPTCHA solution.
type ReCaptchaType string

const (
	// ReCaptchaSingleTile indicates a single-tile reCAPTCHA solution.
	ReCaptchaSingleTile ReCaptchaType = "single"
	// ReCaptchaMultiTile indicates a multi-tile reCAPTCHA solution.
	ReCaptchaMultiTile ReCaptchaType = "multi"
)

// ReCaptchaV2ClassificationSolution holds the response for a reCAPTCHA v2 classification task.
type ReCaptchaV2ClassificationSolution struct {
	// Type is the type of reCAPTCHA solution.
	Type ReCaptchaType `json:"type"`
	// Objects are the indexes of matching tiles for multi-tile questions.
	Objects []int `json:"objects,omitzero"`
	// HasObject indicates presence of the object for single-tile questions.
	HasObject bool `json:"hasObject,omitzero"`
	// Size is the number of tiles to select (e.g., 1 for single, 3–4 for multi).
	Size int `json:"size"`
}

// SolveReCaptchaV2Classification submits a ReCaptchaV2ClassificationTask and returns the solution.
func (s *Session) SolveReCaptchaV2Classification(task ReCaptchaV2ClassificationTask) (*ReCaptchaV2ClassificationSolution, error) {
	task.Type = "ReCaptchaV2Classification"
	res, err := s.Solve(task)
	if err != nil {
		return nil, err
	}
	var solution ReCaptchaV2ClassificationSolution
	if err := res.Unmarshal(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}

// ReCaptchaV2Task defines a reCAPTCHA v2 solve task.
type ReCaptchaV2Task struct {
	Type string `json:"type"`
	// WebsiteURL is the address of the page with the reCAPTCHA widget.
	WebsiteURL string `json:"websiteURL"`
	// WebsiteKey is the site key for the reCAPTCHA widget.
	WebsiteKey string `json:"websiteKey"`
	// Proxy is the proxy string (e.g., "http:ip:port:user:pass").
	Proxy string `json:"proxy,omitzero"`
	// PageAction is the “sa” parameter value from the /anchor payload.
	PageAction string `json:"pageAction,omitzero"`
	// EnterprisePayload carries the enterprise “s” token.
	EnterprisePayload map[string]any `json:"enterprisePayload,omitzero"`
	// IsInvisible marks this as an invisible reCAPTCHA.
	IsInvisible bool `json:"isInvisible,omitzero"`
	// IsSession enables session mode, returning recaptcha-ca-t.
	IsSession bool `json:"isSession,omitzero"`
	// APIDomain is the domain for loading the captcha (usually omitted).
	APIDomain string `json:"apiDomain,omitzero"`
}

// ReCaptchaV2Solution represents the solve result for a reCAPTCHA v2 task.
type ReCaptchaV2Solution struct {
	// UserAgent is the User-Agent string used during solving.
	UserAgent string `json:"userAgent,omitzero"`
	// CreateTime is the token creation timestamp in milliseconds since epoch.
	CreateTime int64 `json:"createTime"`
	// GRecaptchaResponse is the token to submit back to the target site.
	GRecaptchaResponse string `json:"gRecaptchaResponse"`
	// RecaptchaCaT is the session token returned when IsSession was enabled.
	RecaptchaCaT string `json:"recaptcha-ca-t,omitzero"`
	// RecaptchaCaE is an additional cookie parameter returned by some v2 sites.
	RecaptchaCaE string `json:"recaptcha-ca-e,omitzero"`
}

// SolveReCaptchaV2 submits a ReCaptchaV2Task and returns the solution.
func (s *Session) SolveReCaptchaV2(task ReCaptchaV2Task) (*ReCaptchaV2Solution, error) {
	if task.Proxy != "" {
		task.Type = "ReCaptchaV2Task"
	} else {
		task.Type = "ReCaptchaV2TaskProxyLess"
	}
	res, err := s.Solve(task)
	if err != nil {
		return nil, err
	}
	var solution ReCaptchaV2Solution
	if err := res.Unmarshal(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}

// ReCaptchaV3Task defines a reCAPTCHA v3 solve task.
type ReCaptchaV3Task struct {
	Type string `json:"type"`
	// WebsiteURL is the address of the page with the reCAPTCHA.
	WebsiteURL string `json:"websiteURL"`
	// WebsiteKey is the site key for the reCAPTCHA widget.
	WebsiteKey string `json:"websiteKey"`
	// Proxy is the proxy string (e.g., "http://ip:port:user:pass").
	Proxy string `json:"proxy,omitzero"`
	// PageAction is the action parameter value used in grecaptcha.execute.
	PageAction string `json:"pageAction,omitzero"`
	// EnterprisePayload carries the enterprise “s” token.
	EnterprisePayload map[string]any `json:"enterprisePayload,omitzero"`
	// IsSession enables session mode, returning recaptcha-ca-t.
	IsSession bool `json:"isSession,omitzero"`
	// APIDomain is the domain for loading the captcha (usually omitted).
	APIDomain string `json:"apiDomain,omitzero"`
}

// ReCaptchaV3Solution represents the solve result for a reCAPTCHA v3 task.
type ReCaptchaV3Solution struct {
	// UserAgent is the User-Agent string used during solving.
	UserAgent string `json:"userAgent,omitzero"`
	// CreateTime is the token creation timestamp (ms since epoch).
	CreateTime int64 `json:"createTime"`
	// GRecaptchaResponse is the token to submit back to the target site.
	GRecaptchaResponse string `json:"gRecaptchaResponse"`
	// RecaptchaCaT is the session token returned when IsSession was enabled.
	RecaptchaCaT string `json:"recaptcha-ca-t,omitzero"`
	// RecaptchaCaE is an extra cookie parameter returned by some sites.
	RecaptchaCaE string `json:"recaptcha-ca-e,omitzero"`
}

// SolveReCaptchaV3 submits a ReCaptchaV3Task and returns the solution.
func (s *Session) SolveReCaptchaV3(task ReCaptchaV3Task) (*ReCaptchaV3Solution, error) {
	if task.EnterprisePayload != nil {
		if task.Proxy != "" {
			task.Type = "ReCaptchaV3EnterpriseTask"
		} else {
			task.Type = "ReCaptchaV3EnterpriseTaskProxyLess"
		}
	} else {
		if task.Proxy != "" {
			task.Type = "ReCaptchaV3Task"
		} else {
			task.Type = "ReCaptchaV3TaskProxyLess"
		}
	}
	res, err := s.Solve(task)
	if err != nil {
		return nil, err
	}
	var solution ReCaptchaV3Solution
	if err := res.Unmarshal(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}
