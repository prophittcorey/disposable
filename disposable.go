package disposable

import (
	"bytes"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	// A map of disposable domain sources. All sources will be fetched concurrently
	// and merged together.
	Sources = map[string][]byte{
		"https://github.com/disposable/disposable-email-domains/blob/master/domains.txt": []byte{},
	}

	// HTTPClient is used to perform all HTTP requests. You can specify your own
	// to set a custom timeout, proxy, etc.
	HTTPClient = http.Client{
		Timeout: 3 * time.Second,
	}

	// CachePeriod specifies the amount of time an internal cache of disposable email domains are used
	// before refreshing the domains.
	CachePeriod = 45 * time.Minute

	// UserAgent will be used in each request's user agent header field.
	UserAgent = "github.com/prophittcorey/disposable"
)

var (
	domains     = lockabledomains{domains: map[string]struct{}{}}
	lastFetched = time.Now()
)

type lockabledomains struct {
	sync.RWMutex
	domains map[string]struct{}
}

func refreshDomains() error {
	/* aggregate domains concurrently */

	wg := sync.WaitGroup{}

	for url, _ := range Sources {
		wg.Add(1)

		go (func(url string) {
			defer wg.Done()

			req, err := http.NewRequest(http.MethodGet, url, nil)

			if err != nil {
				return
			}

			req.Header.Set("User-Agent", UserAgent)

			res, err := HTTPClient.Do(req)

			if err != nil {
				return
			}

			if bs, err := io.ReadAll(res.Body); err == nil {
				Sources[url] = bs
			}
		})(url)
	}

	wg.Wait()

	/* merge / dedupe all domains */

	ds := map[string]struct{}{}

	for _, bs := range Sources {
		for _, domain := range bytes.Fields(bs) {
			ds[string(domain)] = struct{}{}
		}
	}

	/* update global exit node domains */

	domains.Lock()

	domains.domains = ds
	lastFetched = time.Now()

	domains.Unlock()

	return nil
}

// Check returns true if an email domain is a known disposable domain, false
// otherwise.
func Check(domain string) (bool, error) {
	domains.RLock()

	if len(domains.domains) == 0 || time.Now().After(lastFetched.Add(CachePeriod)) {
		domains.RUnlock()

		if err := refreshDomains(); err != nil {
			return false, err
		}

		domains.RLock()
	}

	defer domains.RUnlock()

	if _, ok := domains.domains[domain]; ok {
		return true, nil
	}

	return false, nil
}

// Domains returns a slice of all known exit node domains.
func Domains() []string {
	domains.RLock()

	if len(domains.domains) == 0 || time.Now().After(lastFetched.Add(CachePeriod)) {
		domains.RUnlock()

		if err := refreshDomains(); err != nil {
			return []string{}
		}

		domains.RLock()
	}

	defer domains.RUnlock()

	ds := []string{}

	for addr := range domains.domains {
		ds = append(ds, addr)
	}

	return ds
}
