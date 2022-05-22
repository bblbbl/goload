package internal

import "github.com/valyala/fasthttp"

type Executor struct {
	config     *Config
	executed   uint
	successful uint
	failed     uint
	timeSpend  int
}

func NewExecutor(c *Config) *Executor {
	ex := new(Executor)
	ex.config = c
	return ex
}

func (e *Executor) Execute() {
	f := func() {
		e.makeRequest()
	}

	for {
		for i := 0; i < e.config.Concurrent; i++ {
			go f()
		}
	}
}

func (e *Executor) makeRequest() {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(e.config.Url)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	_ = fasthttp.Do(req, resp)
	if e.config.RequiredStatus != resp.StatusCode() {
		e.failed++
	} else {
		e.successful++
	}
}
