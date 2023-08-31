package types

type pathWrapper struct {
	Type     string
	Request  any
	Response any
	opts     pathOpts
}

type pathOpts struct {
	Summary     string
	Description string
	Tags        []string
}

type pathOptFunc func(*pathOpts)

func (w *pathWrapper) Operations() *Operation {
	return &Operation{
		Summary: w.opts.Summary,
		Tags:    w.opts.Tags,
		RequestBody: &RequestBody{
			Required: true,
			Content:  getType(w.Request),
		},
		Responses: map[string]Response{
			"200": {
				Description: "",
				Content:     getType(w.Response),
			},
		},
	}
}

func WithSummary(summary string) pathOptFunc {
	return func(po *pathOpts) {
		po.Summary = summary
	}
}

func WithTags(tags []string) pathOptFunc {
	return func(po *pathOpts) {
		po.Tags = tags
	}
}

func WithDesc(desc string) pathOptFunc {
	return func(po *pathOpts) {
		po.Description = desc
	}
}
