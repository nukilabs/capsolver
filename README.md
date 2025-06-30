# CapSolver Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/nukilabs/capsolver.svg)](https://pkg.go.dev/github.com/nukilabs/capsolver)

Go SDK for easy integration with the [CapSolver API](https://docs.capsolver.com).

## Installation

```bash
go get github.com/nukilabs/capsolver
```

## Usage

```go
package main

import "github.com/nukilabs/capsolver"

func main() {
  // create a session with your API key
  client := capsolver.New("YOUR_CLIENT_KEY")

  // build a reCAPTCHA v2 task (proxy-less by default)
  task := capsolver.ReCaptchaV2Task{
    WebsiteURL: "https://example.com",
    WebsiteKey: "SITE_KEY",
  }

  // send the task and retrieve the solution
  sol, err := client.SolveReCaptchaV2(task)
  if err != nil {
    panic(err) // handle error appropriately
  }

  // print the token to submit to the target site
  println(sol.GRecaptchaResponse)
}
```

## Supported Captcha Types

- Image-to-text (OCR)  
- GeeTest v3 / v4  
- reCAPTCHA v2 & v3  
- reCAPTCHA classification  
- Cloudflare Turnstile  
- AWS WAF classification & challenge  
- DataDome slider/interstitial  
- MtCaptcha token  
- VisionEngine puzzles  

---

© 2025 nukilabs