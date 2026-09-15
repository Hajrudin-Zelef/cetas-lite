package search

import (
	"context"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var (
	ddgLinkRe = regexp.MustCompile(`(?s)<a[^>]*class="result__a"[^>]*href="([^"]+)"[^>]*>(.*?)</a>`)
	ddgSnipRe = regexp.MustCompile(`(?s)<a[^>]*class="result__snippet"[^>]*>(.*?)</a>`)
	ddgTagRe  = regexp.MustCompile(`<[^>]+>`)
)

// searchDuckDuckGo interroge DuckDuckGo sans cle (endpoint HTML).
// C'est le repli gratuit : "sans cle, seul DuckDuckGo est utilise".
func (s *Searcher) searchDuckDuckGo(ctx context.Context, query string, maxResults int) ([]Hit, error) {
	endpoint := "https://html.duckduckgo.com/html/?" + url.Values{"q": {query}}.Encode()
	req, err := s.jsonRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/html")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0 Safari/537.36")
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == 202 {
		return nil, fmt.Errorf("duckduckgo: limite anti-bot atteinte (HTTP %d)", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusOK {
		_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4096))
		return nil, fmt.Errorf("duckduckgo: HTTP %d", resp.StatusCode)
	}
	raw, err := io.ReadAll(io.LimitReader(resp.Body, 2<<20))
	if err != nil {
		return nil, err
	}
	return parseDDGHTML(string(raw), maxResults), nil
}

// parseDDGHTML extrait les resultats du HTML de DuckDuckGo.
// Les liens passent par un redirecteur (/l/?uddg=<url encodee>) : on
// recupere l'URL reelle via le parametre uddg.
func parseDDGHTML(page string, max int) []Hit {
	linkMs := ddgLinkRe.FindAllStringSubmatch(page, -1)
	snipMs := ddgSnipRe.FindAllStringSubmatch(page, -1)
	out := make([]Hit, 0, len(linkMs))
	for i, m := range linkMs {
		if len(out) >= max {
			break
		}
		href := html.UnescapeString(m[1])
		if u, err := url.Parse(href); err == nil && u.Host == "duckduckgo.com" {
			if real := u.Query().Get("uddg"); real != "" {
				href = real
			}
		}
		title := ddgText(m[2])
		desc := ""
		if i < len(snipMs) {
			desc = ddgText(snipMs[i][1])
		}
		if href == "" || title == "" || strings.HasPrefix(href, "/") {
			continue
		}
		out = append(out, Hit{
			Title:       sanitize(title),
			URL:         sanitize(href),
			Description: sanitize(desc),
			Source:      "duckduckgo",
		})
	}
	return out
}

func ddgText(raw string) string {
	return strings.Join(strings.Fields(ddgTagRe.ReplaceAllString(html.UnescapeString(raw), " ")), " ")
}
