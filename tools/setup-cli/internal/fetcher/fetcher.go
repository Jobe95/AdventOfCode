package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Jobe95/AdventOfCode/internal/config"
)

func FetchInput(year, day int) (string, error) {
	session, err := config.GetSession()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func FetchExample(year, day int) (string, error) {
	session, err := config.GetSession()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return extractExample(string(body)), nil
}

func extractExample(html string) string {
	// Look for <pre><code>...</code></pre> blocks
	start := strings.Index(html, "<pre><code>")
	if start == -1 {
		return ""
	}
	start += len("<pre><code>")

	end := strings.Index(html[start:], "</code></pre>")
	if end == -1 {
		return ""
	}

	example := html[start : start+end]

	// Basic HTML entity decoding
	example = strings.ReplaceAll(example, "&lt;", "<")
	example = strings.ReplaceAll(example, "&gt;", ">")
	example = strings.ReplaceAll(example, "&amp;", "&")
	example = strings.ReplaceAll(example, "&#39;", "'")
	example = strings.ReplaceAll(example, "&quot;", "\"")

	return strings.TrimSpace(example)
}
