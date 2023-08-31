package types

import "reflect"

type PostRequestHandler[Req any, Res any] func(body Req) (*Res, error)
type XPostRequestHandler func(body any) (*any, error)

func (o *OpenAPI) RegisterPostRequest(path string, req any, res any, opts ...pathOptFunc) {
	pathOpts := pathOpts{}
	for _, poFunc := range opts {
		poFunc(&pathOpts)
	}
	pw := pathWrapper{
		Type:     "POST",
		Request:  req,
		Response: res,
		opts:     pathOpts,
	}

	op := pw.Operations()

	existingPaths, ok := o.Paths[path]
	if !ok {
		// Path does not exist
		o.Paths[path] = Path{
			Post: op,
		}
	} else {
		existingPaths.Post = op
	}
}

type XHandler func(interface{}) (*interface{}, error)

func Post[Req, Res any](o *OpenAPI, path string, handler PostRequestHandler[Req, Res], opts ...pathOptFunc) {
	o.RegisterPostRequest(path, *new(Req), *new(Res), opts...)
}

func (o *OpenAPI) XPost(path string, handler any, opts ...pathOptFunc) {
	fType := reflect.TypeOf(handler)
	argType := fType.In(0)
	v := reflect.New(argType).Elem().Interface()
	o.RegisterPostRequest(path, v, v, opts...)
}
