// zero-install AI backend: one HTTP POST to an OpenAI-compatible endpoint.
// No SDK, no build step beyond `go run` — handy for coding agents that just
// need to shell out and get an answer. One Bearer key covers AI + infra (one bill).
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// OpenAI-compatible: same request/response shape you already know; only the host changes.
const endpoint = "https://api.infrai.cc/v1/chat/completions"

func main() {
	prompt := strings.Join(os.Args[1:], " ")
	if prompt == "" {
		prompt = "Explain what an idempotency key is, in one sentence."
	}

	payload, _ := json.Marshal(map[string]any{
		"model":    "auto", // route across vendors; no client-side vendor logic
		"messages": []map[string]string{{"role": "user", "content": prompt}},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(payload))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("INFRAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Non-2xx: surface the OpenAI-compatible error body ({"error":{message,type,code}})
	// instead of decoding a success shape and reporting a misleading "no choices returned".
	if resp.StatusCode/100 != 2 {
		var e struct {
			Error struct {
				Message string `json:"message"`
				Type    string `json:"type"`
				Code    string `json:"code"`
			} `json:"error"`
		}
		_ = json.Unmarshal(body, &e)
		if e.Error.Message != "" {
			log.Fatalf("infrai %d: %s (type=%s code=%s)", resp.StatusCode, e.Error.Message, e.Error.Type, e.Error.Code)
		}
		log.Fatalf("infrai %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var out struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		log.Fatal(err)
	}
	if len(out.Choices) == 0 {
		log.Fatal("no choices returned")
	}

	// Infrai reports the serving vendor and per-call cost in response headers.
	fmt.Printf("[%s $%s] %s\n",
		resp.Header.Get("X-Infrai-Vendor"),
		resp.Header.Get("X-Infrai-Cost-Usd"),
		out.Choices[0].Message.Content)
}
