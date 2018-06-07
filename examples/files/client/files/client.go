// Code generated by Remoto; DO NOT EDIT.

package files

import (
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/machinebox/remoto/remototypes"
	"github.com/oxtoacart/bpool"
	"github.com/pkg/errors"
)

// ImagesClient accesses remote Images services.
type ImagesClient struct {
	// endpoint is the HTTP endpoint of the remote server.
	endpoint string
	// httpclient is the http.Client to use to make requests.
	httpclient *http.Client
	// bufs is a buffer pool
	bufs *bpool.BufferPool
}

// NewImagesClient makes a new ImagesClient that will
// use the specified http.Client to make requests.
func NewImagesClient(endpoint string, client *http.Client) *ImagesClient {
	return &ImagesClient{
		endpoint:   endpoint,
		httpclient: client,
		bufs:       bpool.NewBufferPool(48),
	}
}

func (c *ImagesClient) Flip(ctx context.Context, request *FlipRequest) (io.ReadCloser, error) {
	b, err := json.Marshal([]interface{}{request})
	if err != nil {
		return nil, errors.Wrap(err, "ImagesClient.Flip: encode request")
	}
	buf := c.bufs.Get()
	defer c.bufs.Put(buf)
	w := multipart.NewWriter(buf)
	w.WriteField("json", string(b))
	if files, ok := ctx.Value(contextKeyFiles).(map[string]file); ok {
		for fieldname, file := range files {
			f, err := w.CreateFormFile(fieldname, file.filename)
			if err != nil {
				return nil, errors.Wrap(err, "ImagesClient.Flip: create form file")
			}
			if _, err := io.Copy(f, file.r); err != nil {
				return nil, errors.Wrap(err, "ImagesClient.Flip: reading file")
			}
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			default:
			}
		}
	}
	if err := w.Close(); err != nil {
		return nil, errors.Wrap(err, "ImagesClient.Flip: write")
	}
	req, err := http.NewRequest(http.MethodPost, c.endpoint+"/remoto/Images.Flip", buf)
	if err != nil {
		return nil, errors.Wrap(err, "ImagesClient.Flip: new request")
	}
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", w.FormDataContentType())
	req = req.WithContext(ctx)
	resp, err := c.httpclient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "ImagesClient.Flip: do")
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, errors.Errorf("ImagesClient.Flip: remote service returned %s", resp.Status)
	}
	return resp.Body, nil
}

// FlipRequest is the request for Images.Flip.
type FlipRequest struct {
	Image remototypes.File `json:"image"`
}

// SetImage sets the file for the Image field.
func (s *FlipRequest) SetImage(ctx context.Context, filename string, r io.Reader) context.Context {
	files, ok := ctx.Value(contextKeyFiles).(map[string]file)
	if !ok {
		files = make(map[string]file)
	}
	fieldname := "files[" + strconv.Itoa(len(files)) + "]"
	files[fieldname] = file{r: r, filename: filename}
	ctx = context.WithValue(ctx, contextKeyFiles, files)
	s.Image = remototypes.File{
		Fieldname: fieldname,
		Filename:  filename,
	}
	return ctx
}

// contextKey is a local context key type.
// see https://medium.com/@matryer/context-keys-in-go-5312346a868d
type contextKey string

func (c contextKey) String() string {
	return "remoto context key: " + string(c)
}

// contextKeyFiles is the context key for the request files.
var contextKeyFiles = contextKey("files")

// file holds info about a file in the context, including
// the io.Reader where the contents will be read from.
type file struct {
	r        io.Reader
	filename string
}

// this is here so we don't get a compiler complaints.
func init() {
	var _ = remototypes.File{}
}
