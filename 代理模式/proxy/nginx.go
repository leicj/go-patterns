package proxy

type nginx struct {
	application *application
	maxAllowedRequest int
	rateLimiter map[string]int
}

func NewNginxServer() *nginx {
	return &nginx{
		application: &application{},
		maxAllowedRequest: 2,
		rateLimiter: make(map[string]int),
	}
}

func (n *nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.CheckRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.HandleRequest(url, method)
}

func (n *nginx) CheckRateLimiting(url string) bool {
    if n.rateLimiter[url] == 0 {
        n.rateLimiter[url] = 1
    }
    if n.rateLimiter[url] > n.maxAllowedRequest {
        return false
    }
    n.rateLimiter[url] = n.rateLimiter[url] + 1
    return true
}