package capsolver

// VisionEngineModule represents a supported VisionEngine model.
type VisionEngineModule string

const (
	// Slider1 is the slider puzzle module.
	Slider1 VisionEngineModule = "slider_1"
	// Rotate1 is the first rotate puzzle module.
	Rotate1 VisionEngineModule = "rotate_1"
	// Rotate2 is the second rotate puzzle module.
	Rotate2 VisionEngineModule = "rotate_2"
	// Shein is the Shein-style graphics selection module.
	Shein VisionEngineModule = "shein"
	// ShopReceipt is the receipt question module.
	ShopReceipt VisionEngineModule = "shop_receipt"
	// SpaceDetection is the space detection (counting) module.
	SpaceDetection VisionEngineModule = "space_detection"
	// SliderTemuPlus is the Temu slider puzzle module.
	SliderTemuPlus VisionEngineModule = "slider_temu_plus"
	// SelectTemu is the Temu image selection module.
	SelectTemu VisionEngineModule = "select_temu"
)

// VisionEngineTask defines the parameters for a VisionEngine recognition task.
type VisionEngineTask struct {
	Type string `json:"type"`
	// Module specifies which VisionEngine model to use.
	Module VisionEngineModule `json:"module"`
	// WebsiteURL is the page source URL to improve accuracy.
	WebsiteURL string `json:"websiteURL,omitzero"`
	// Image is the base64-encoded image content without the data URI prefix.
	Image string `json:"image"`
	// ImageBackground is the base64-encoded background image content.
	ImageBackground string `json:"imageBackground"`
	// Question is required for modules that ask a question (e.g. "space_detection").
	Question string `json:"question,omitzero"`
}

// Rect defines a rectangular region in am image.
type Rect struct {
	// X1,Y1 are the coordinates of the top-left corner.
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
	// X2,Y2 are the coordinates of the bottom-right corner.
	X2 int `json:"x2"`
	Y2 int `json:"y2"`
}

// VisionEngineSolution represents the response payload for a VisionEngine task.
type VisionEngineSolution struct {
	// Distance is the slide distance for slider modules.
	Distance float64 `json:"distance,omitzero"`
	// Angle is the rotation angle for rotate modules.
	Angle float64 `json:"angle,omitzero"`
	// Rects are the bounding rectangles for modules returning multiple regions.
	Rects []Rect `json:"rects,omitzero"`
	// Box is the single bounding box for space_detection.
	Box []float64 `json:"box,omitzero"`
	// Text is the recognized text for text-extraction modules.
	Text string `json:"text,omitzero"`
	// Objects are the indexes of matching tiles for selection modules.
	Objects []int `json:"objects,omitzero"`
}

// SolveVisionEngine submits a VisionEngine task and returns the solution.
func (s *Session) SolveVisionEngine(task VisionEngineTask) (*VisionEngineSolution, error) {
	task.Type = "VisionEngine"
	res, err := s.Solve(task)
	if err != nil {
		return nil, err
	}
	var solution VisionEngineSolution
	if err := res.Unmarshal(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}
