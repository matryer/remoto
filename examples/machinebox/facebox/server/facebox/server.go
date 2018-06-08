// Code generated by Remoto; DO NOT EDIT.

// Package facebox contains the HTTP server for facebox services.
package facebox

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/machinebox/remoto/go/remotohttp"
	"github.com/machinebox/remoto/remototypes"
	"github.com/pkg/errors"
)

// Facebox provides facial detection and recognition in images.
type Facebox interface {
	CheckFaceprint(context.Context, *CheckFaceprintRequest) (*CheckFaceprintResponse, error)

	CheckFile(context.Context, *CheckFileRequest) (*CheckFileResponse, error)

	CheckURL(context.Context, *CheckURLRequest) (*CheckURLResponse, error)

	FaceprintCompare(context.Context, *FaceprintCompareRequest) (*FaceprintCompareResponse, error)

	GetState(context.Context, *GetStateRequest) (*remototypes.FileResponse, error)

	PutState(context.Context, *PutStateRequest) (*PutStateResponse, error)

	RemoveID(context.Context, *RemoveIDRequest) (*RemoveIDResponse, error)

	Rename(context.Context, *RenameRequest) (*RenameResponse, error)

	RenameID(context.Context, *RenameIDRequest) (*RenameIDResponse, error)

	SimilarFile(context.Context, *SimilarFileRequest) (*SimilarFileResponse, error)

	SimilarID(context.Context, *SimilarIDRequest) (*SimilarIDResponse, error)

	SimilarURL(context.Context, *SimilarURLRequest) (*SimilarURLResponse, error)

	TeachFaceprint(context.Context, *TeachFaceprintRequest) (*TeachFaceprintResponse, error)

	TeachFile(context.Context, *TeachFileRequest) (*TeachFileResponse, error)

	TeachURL(context.Context, *TeachURLRequest) (*TeachURLResponse, error)
}

// Run is the simplest way to run the services.
func Run(addr string,
	facebox Facebox,
) error {
	server := New(
		facebox,
	)
	if err := server.Describe(os.Stdout); err != nil {
		return errors.Wrap(err, "describe service")
	}
	if err := http.ListenAndServe(addr, server); err != nil {
		return err
	}
	return nil
}

// New makes a new remotohttp.Server with the specified services
// registered.
func New(
	facebox Facebox,
) *remotohttp.Server {
	server := &remotohttp.Server{
		OnErr: func(w http.ResponseWriter, r *http.Request, err error) {
			fmt.Fprintf(os.Stderr, "%s %s: %s\n", r.Method, r.URL.Path, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
		NotFound: http.NotFoundHandler(),
	}

	RegisterFaceboxServer(server, facebox)
	return server
}

// RegisterFaceboxServer registers a Facebox with a remotohttp.Server.
func RegisterFaceboxServer(server *remotohttp.Server, service Facebox) {
	srv := &httpFaceboxServer{
		service: service,
		server:  server,
	}
	server.Register("/remoto/Facebox.CheckFaceprint", http.HandlerFunc(srv.handleCheckFaceprint))
	server.Register("/remoto/Facebox.CheckFile", http.HandlerFunc(srv.handleCheckFile))
	server.Register("/remoto/Facebox.CheckURL", http.HandlerFunc(srv.handleCheckURL))
	server.Register("/remoto/Facebox.FaceprintCompare", http.HandlerFunc(srv.handleFaceprintCompare))
	server.Register("/remoto/Facebox.GetState", http.HandlerFunc(srv.handleGetState))
	server.Register("/remoto/Facebox.PutState", http.HandlerFunc(srv.handlePutState))
	server.Register("/remoto/Facebox.RemoveID", http.HandlerFunc(srv.handleRemoveID))
	server.Register("/remoto/Facebox.Rename", http.HandlerFunc(srv.handleRename))
	server.Register("/remoto/Facebox.RenameID", http.HandlerFunc(srv.handleRenameID))
	server.Register("/remoto/Facebox.SimilarFile", http.HandlerFunc(srv.handleSimilarFile))
	server.Register("/remoto/Facebox.SimilarID", http.HandlerFunc(srv.handleSimilarID))
	server.Register("/remoto/Facebox.SimilarURL", http.HandlerFunc(srv.handleSimilarURL))
	server.Register("/remoto/Facebox.TeachFaceprint", http.HandlerFunc(srv.handleTeachFaceprint))
	server.Register("/remoto/Facebox.TeachFile", http.HandlerFunc(srv.handleTeachFile))
	server.Register("/remoto/Facebox.TeachURL", http.HandlerFunc(srv.handleTeachURL))

}

type CheckFaceprintRequest struct {
	Faceprints []string `json:"faceprints"`
}

type CheckFaceprintResponse struct {
	Faces []FaceprintFace `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type CheckFileRequest struct {
	File remototypes.File `json:"file"`
}

type CheckFileResponse struct {
	Faces []Face `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type CheckURLRequest struct {
	File remototypes.File `json:"file"`
}

type CheckURLResponse struct {
	Faces []Face `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type Face struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Matched bool `json:"matched"`

	Faceprint string `json:"faceprint"`

	Rect Rect `json:"rect"`
}

type FaceprintCompareRequest struct {
	Target string `json:"target"`

	Faceprints []string `json:"faceprints"`
}

type FaceprintCompareResponse struct {
	Confidences []float64 `json:"confidences"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type FaceprintFace struct {
	Matched bool `json:"matched"`

	Confidence float64 `json:"confidence"`

	ID string `json:"id"`

	Name string `json:"name"`
}

type GetStateRequest struct {
}

type PutStateRequest struct {
	StateFile remototypes.File `json:"state_file"`
}

type PutStateResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type Rect struct {
	Top int `json:"top"`

	Left int `json:"left"`

	Width int `json:"width"`

	Height int `json:"height"`
}

type RemoveIDRequest struct {
	ID string `json:"id"`
}

type RemoveIDResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type RenameIDRequest struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

type RenameIDResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type RenameRequest struct {
	From string `json:"from"`

	To string `json:"to"`
}

type RenameResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type SimilarFace struct {
	Rect Rect `json:"rect"`

	SimilarFaces []Face `json:"similar_faces"`
}

type SimilarFileRequest struct {
	File remototypes.File `json:"file"`
}

type SimilarFileResponse struct {
	Faces []SimilarFace `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type SimilarIDRequest struct {
	ID string `json:"id"`
}

type SimilarIDResponse struct {
	Faces []SimilarFace `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type SimilarURLRequest struct {
	URL string `json:"url"`
}

type SimilarURLResponse struct {
	Faces []SimilarFace `json:"faces"`

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type TeachFaceprintRequest struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Faceprint string `json:"faceprint"`
}

type TeachFaceprintResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type TeachFileRequest struct {
	ID string `json:"id"`

	Name string `json:"name"`

	File remototypes.File `json:"file"`
}

type TeachFileResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

type TeachURLRequest struct {
	ID string `json:"id"`

	Name string `json:"name"`

	URL string `json:"url"`
}

type TeachURLResponse struct {

	// Error is an error message if one occurred.
	Error string `json:"error"`
}

// httpFaceboxServer is an internal type that provides an
// HTTP wrapper around Facebox.
type httpFaceboxServer struct {
	// service is the Facebox being exposed by this
	// server.
	service Facebox
	// server is the remotohttp.Server that this server is
	// registered with.
	server *remotohttp.Server
}

// handleCheckFaceprint is an http.Handler wrapper for Facebox.CheckFaceprint.
func (srv *httpFaceboxServer) handleCheckFaceprint(w http.ResponseWriter, r *http.Request) {
	var reqs []*CheckFaceprintRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]CheckFaceprintResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.CheckFaceprint(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleCheckFile is an http.Handler wrapper for Facebox.CheckFile.
func (srv *httpFaceboxServer) handleCheckFile(w http.ResponseWriter, r *http.Request) {
	var reqs []*CheckFileRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]CheckFileResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.CheckFile(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleCheckURL is an http.Handler wrapper for Facebox.CheckURL.
func (srv *httpFaceboxServer) handleCheckURL(w http.ResponseWriter, r *http.Request) {
	var reqs []*CheckURLRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]CheckURLResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.CheckURL(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleFaceprintCompare is an http.Handler wrapper for Facebox.FaceprintCompare.
func (srv *httpFaceboxServer) handleFaceprintCompare(w http.ResponseWriter, r *http.Request) {
	var reqs []*FaceprintCompareRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]FaceprintCompareResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.FaceprintCompare(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleGetState is an http.Handler wrapper for Facebox.GetState.
func (srv *httpFaceboxServer) handleGetState(w http.ResponseWriter, r *http.Request) {
	var reqs []*GetStateRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	// single file response

	if len(reqs) != 1 {
		if err := remotohttp.EncodeErr(w, r, errors.New("only single requests supported for file response endpoints")); err != nil {
			srv.server.OnErr(w, r, err)
			return
		}
		return
	}

	resp, err := srv.service.GetState(r.Context(), reqs[0])
	if err != nil {
		resp.Error = err.Error()
		if err := remotohttp.Encode(w, r, http.StatusOK, []interface{}{resp}); err != nil {
			srv.server.OnErr(w, r, err)
			return
		}
	}
	if resp.ContentType == "" {
		resp.ContentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", resp.ContentType)
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.QuoteToASCII(resp.Filename))
	if resp.ContentLength > 0 {
		w.Header().Set("Content-Length", strconv.Itoa(resp.ContentLength))
	}
	if _, err := io.Copy(w, resp.Data); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handlePutState is an http.Handler wrapper for Facebox.PutState.
func (srv *httpFaceboxServer) handlePutState(w http.ResponseWriter, r *http.Request) {
	var reqs []*PutStateRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]PutStateResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.PutState(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleRemoveID is an http.Handler wrapper for Facebox.RemoveID.
func (srv *httpFaceboxServer) handleRemoveID(w http.ResponseWriter, r *http.Request) {
	var reqs []*RemoveIDRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]RemoveIDResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.RemoveID(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleRename is an http.Handler wrapper for Facebox.Rename.
func (srv *httpFaceboxServer) handleRename(w http.ResponseWriter, r *http.Request) {
	var reqs []*RenameRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]RenameResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.Rename(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleRenameID is an http.Handler wrapper for Facebox.RenameID.
func (srv *httpFaceboxServer) handleRenameID(w http.ResponseWriter, r *http.Request) {
	var reqs []*RenameIDRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]RenameIDResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.RenameID(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleSimilarFile is an http.Handler wrapper for Facebox.SimilarFile.
func (srv *httpFaceboxServer) handleSimilarFile(w http.ResponseWriter, r *http.Request) {
	var reqs []*SimilarFileRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]SimilarFileResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.SimilarFile(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleSimilarID is an http.Handler wrapper for Facebox.SimilarID.
func (srv *httpFaceboxServer) handleSimilarID(w http.ResponseWriter, r *http.Request) {
	var reqs []*SimilarIDRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]SimilarIDResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.SimilarID(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleSimilarURL is an http.Handler wrapper for Facebox.SimilarURL.
func (srv *httpFaceboxServer) handleSimilarURL(w http.ResponseWriter, r *http.Request) {
	var reqs []*SimilarURLRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]SimilarURLResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.SimilarURL(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleTeachFaceprint is an http.Handler wrapper for Facebox.TeachFaceprint.
func (srv *httpFaceboxServer) handleTeachFaceprint(w http.ResponseWriter, r *http.Request) {
	var reqs []*TeachFaceprintRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]TeachFaceprintResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.TeachFaceprint(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleTeachFile is an http.Handler wrapper for Facebox.TeachFile.
func (srv *httpFaceboxServer) handleTeachFile(w http.ResponseWriter, r *http.Request) {
	var reqs []*TeachFileRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]TeachFileResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.TeachFile(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// handleTeachURL is an http.Handler wrapper for Facebox.TeachURL.
func (srv *httpFaceboxServer) handleTeachURL(w http.ResponseWriter, r *http.Request) {
	var reqs []*TeachURLRequest
	if err := remotohttp.Decode(r, &reqs); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

	resps := make([]TeachURLResponse, len(reqs))
	for i := range reqs {
		resp, err := srv.service.TeachURL(r.Context(), reqs[i])
		if err != nil {
			resps[i].Error = err.Error()
			continue
		}
		resps[i] = *resp
	}
	if err := remotohttp.Encode(w, r, http.StatusOK, resps); err != nil {
		srv.server.OnErr(w, r, err)
		return
	}

}

// this is here so we don't get a compiler complaints.
func init() {
	var _ = remototypes.File{}
	var _ = strconv.Itoa(0)
	var _ = io.EOF
}