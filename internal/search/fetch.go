package search

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	readability "codeberg.org/readeck/go-readability/v2"
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

const (
	fetchTimeout  = 10 * time.Second
	fetchMaxBody  = 2 << 20
	fetchMaxChars = 8000
)

var errBlockedAddress = errors.New("adresse reseau non autorisee")

func newFetchClient() *http.Client {
	transport := &http.Transport{
		Proxy: nil,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			host, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}
			ips, err := net.DefaultResolver.LookupIPAddr(ctx, host)
			if err != nil {
				return nil, err
			}
			d := &net.Dialer{Timeout: 5 * time.Second}
			for _, ip := range ips {
				if !isPublicIP(ip.IP) {
					continue
				}
				return d.DialContext(ctx, network, net.JoinHostPort(ip.String(), port))
			}
			return nil, errBlockedAddress
		},
	}
	return &http.Client{
		Timeout:   fetchTimeout,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return errors.New("trop de redirections")
			}
			if req.URL.Scheme != "http" && req.URL.Scheme != "https" {
				return errBlockedAddress
			}
			return nil
		},
	}
}

func isPublicIP(ip net.IP) bool {
	if ip == nil || ip.IsLoopback() || ip.IsPrivate() || ip.IsUnspecified() ||
		ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.IsMulticast() {
		return false
	}
	if v4 := ip.To4(); v4 != nil {
		if v4[0] == 100 && v4[1] >= 64 && v4[1] <= 127 {
			return false
		}
		return true
	}
	if len(ip) == net.IPv6len && (ip[0]&0xfe) == 0xfc {
		return false
	}
	return true
}

func validateFetchURL(raw string) (*url.URL, error) {
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil {
		return nil, errors.New("URL invalide")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, fmt.Errorf("schema non supporte: %s", u.Scheme)
	}
	if u.Hostname() == "" {
		return nil, errors.New("hote manquant")
	}
	return u, nil
}

func (s *Searcher) Fetch(ctx context.Context, raw string) (string, string, error) {
	u, err := validateFetchURL(raw)
	if err != nil {
		return "", "", err
	}
	if title, md, err := s.fetchDirect(ctx, u); err == nil {
		return title, md, nil
	} else {
		// Repli : lecteur Jina (necessite la cle Jina, deja utilisee pour
		// la recherche). Utile quand le fetch direct echoue : page
		// JavaScript, 403, extraction impossible...
		if key := strings.TrimSpace(s.key("jina")); key != "" {
			if md, err := s.fetchViaJina(ctx, u.String(), key); err == nil {
				return "", md, nil
			}
		}
		return "", "", err
	}
}

// fetchDirect recupere une page en HTTP direct et en extrait le contenu
// lisible en Markdown (readability + html-to-markdown).
func (s *Searcher) fetchDirect(ctx context.Context, u *url.URL) (string, string, error) {
	client := s.fetchClient
	if client == nil {
		client = newFetchClient()
	}
	ctx, cancel := context.WithTimeout(ctx, fetchTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; cetas-lite)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,text/plain")
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "" && !strings.Contains(ct, "html") && !strings.Contains(ct, "text/plain") {
		return "", "", fmt.Errorf("type de contenu non supporte: %s", ct)
	}
	article, err := readability.FromReader(io.LimitReader(resp.Body, fetchMaxBody), u)
	if err != nil || article.Node == nil {
		return "", "", errors.New("extraction impossible")
	}
	md, err := htmltomarkdown.ConvertNode(article.Node)
	if err != nil || strings.TrimSpace(string(md)) == "" {
		var sb strings.Builder
		if terr := article.RenderText(&sb); terr == nil {
			md = []byte(sb.String())
		}
	}
	content := truncateRunes(string(md), fetchMaxChars)
	if content == "" {
		return "", "", errors.New("contenu vide")
	}
	return sanitize(article.Title()), content, nil
}

// fetchViaJina recupere une page via le lecteur Jina (https://r.jina.ai/),
// qui renvoie directement du Markdown lisible sans extraction locale.
// Repli quand le fetch direct echoue.
func (s *Searcher) fetchViaJina(ctx context.Context, target, apiKey string) (string, error) {
	base := s.endpoints["jina_reader"]
	if base == "" {
		base = "https://r.jina.ai/"
	}
	ctx, cancel := context.WithTimeout(ctx, fetchTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, base+target, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "text/plain")
	client := s.client
	if client == nil {
		client = &http.Client{Timeout: fetchTimeout}
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("jina reader: HTTP %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, fetchMaxBody))
	if err != nil {
		return "", err
	}
	md := truncateRunes(strings.TrimSpace(string(body)), fetchMaxChars)
	if md == "" {
		return "", errors.New("contenu vide")
	}
	return md, nil
}
