package capsolver

// ErrorCode represents the CapSolver API error code string.
type ErrorCode string

const (
	// ServiceUnavailable indicates the service is temporarily unavailable.
	// It is possible that the server pressure is high. Please try again later.
	// If it continues to appear, contact customer service.
	ServiceUnavailable ErrorCode = "ERROR_SERVICE_UNAVALIABLE"

	// RateLimit indicates the request frequency/rate limit has been exceeded.
	// This happens when you exceed your service package's request limit.
	RateLimit ErrorCode = "ERROR_RATE_LIMIT"

	// InvalidTaskData means the submitted task data was invalid.
	// Please check the details of errorDescription for what was wrong.
	InvalidTaskData ErrorCode = "ERROR_INVALID_TASK_DATA"

	// BadRequest indicates a general request error.
	// If this persists, contact customer service.
	BadRequest ErrorCode = "ERROR_BAD_REQUEST"

	// TaskIDInvalid means the provided task ID does not exist or is invalid.
	// This may happen if a wrong or expired ID is used.
	TaskIDInvalid ErrorCode = "ERROR_TASKID_INVALID"

	// TaskTimeout means the task has timed out after waiting too long (120s).
	// The captcha could not be solved in time.
	TaskTimeout ErrorCode = "ERROR_TASK_TIMEOUT"

	// SettlementFailed indicates a failure in mission point settlement.
	// Please check your balance or contact customer service.
	SettlementFailed ErrorCode = "ERROR_SETTLEMENT_FAILED"

	// KeyDeniedAccess means the provided account key is incorrect.
	// Check your clientKey and ensure it matches the one from your personal center.
	KeyDeniedAccess ErrorCode = "ERROR_KEY_DENIED_ACCESS"

	// ZeroBalance indicates the account balance is insufficient.
	// You need to top up your account to continue usage.
	ZeroBalance ErrorCode = "ERROR_ZERO_BALANCE"

	// TaskNotSupported means the captcha type is incorrect or not supported.
	// Please verify the task type you're submitting.
	TaskNotSupported ErrorCode = "ERROR_TASK_NOT_SUPPORTED"

	// CaptchaUnsolvable means the captcha could not be recognized.
	// No credits were deducted. Please try again.
	CaptchaUnsolvable ErrorCode = "ERROR_CAPTCHA_UNSOLVABLE"

	// UnknownQuestion means the submitted question ID was invalid or unknown.
	// Usually indicates an unprocessable task.
	UnknownQuestion ErrorCode = "ERROR_UNKNOWN_QUESTION"

	// ProxyBanned means the proxy IP was banned by the target service.
	// Please use a different proxy.
	ProxyBanned ErrorCode = "ERROR_PROXY_BANNED"

	// InvalidImage means the image size does not meet the requirements.
	// Ensure your input image is valid and properly formatted.
	InvalidImage ErrorCode = "ERROR_INVALID_IMAGE"

	// ParseImageFail means the server failed to parse your image.
	// Check if your image's BASE64 encoding is correct.
	ParseImageFail ErrorCode = "ERROR_PARSE_IMAGE_FAIL"

	// IPBanned indicates your IP has been temporarily blocked.
	// If many errors occur rapidly (e.g., 1000 in 1 minute), a 30-minute block is triggered.
	IPBanned ErrorCode = "ERROR_IP_BANNED"

	// KeyTempBlocked means your key has been temporarily blocked due to too many errors.
	// It will be automatically unblocked in 5 minutes. Please try again later.
	KeyTempBlocked ErrorCode = "ERROR_KEY_TEMP_BLOCKED"
)

// Error represents a CapSolver API error response.
type Error struct {
	ID          int       `json:"errorId"`
	Code        ErrorCode `json:"errorCode"`
	Description string    `json:"errorDescription"`
}

func (e Error) Error() string {
	return "[capsolver] " + e.Description
}
