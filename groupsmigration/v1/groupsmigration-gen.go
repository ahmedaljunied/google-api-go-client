// Package groupsmigration provides access to the Groups Migration API.
//
// See https://developers.google.com/google-apps/groups-migration/
//
// Usage example:
//
//   import "google.golang.org/api/groupsmigration/v1"
//   ...
//   groupsmigrationService, err := groupsmigration.New(oauthHttpClient)
package groupsmigration // import "google.golang.org/api/groupsmigration/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "groupsmigration:v1"
const apiName = "groupsmigration"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/groups/v1/groups/"

// OAuth2 scopes used by this API.
const (
	// Manage messages in groups on your domain
	AppsGroupsMigrationScope = "https://www.googleapis.com/auth/apps.groups.migration"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Archive = NewArchiveService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Archive *ArchiveService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewArchiveService(s *Service) *ArchiveService {
	rs := &ArchiveService{s: s}
	return rs
}

type ArchiveService struct {
	s *Service
}

// Groups: JSON response template for groups migration API.
type Groups struct {
	// Kind: The kind of insert resource this is.
	Kind string `json:"kind,omitempty"`

	// ResponseCode: The status of the insert request.
	ResponseCode string `json:"responseCode,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Groups) MarshalJSON() ([]byte, error) {
	type noMethod Groups
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "groupsmigration.archive.insert":

type ArchiveInsertCall struct {
	s                *Service
	groupId          string
	urlParams_       gensupport.URLParams
	media_           io.Reader
	resumableBuffer_ *gensupport.ResumableBuffer
	mediaType_       string
	mediaSize_       int64 // mediaSize, if known.  Used only for calls to progressUpdater_.
	progressUpdater_ googleapi.ProgressUpdater
	ctx_             context.Context
}

// Insert: Inserts a new mail into the archive of the Google group.
func (r *ArchiveService) Insert(groupId string) *ArchiveInsertCall {
	c := &ArchiveInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.groupId = groupId
	return c
}

// Media specifies the media to upload in a single chunk. At most one of
// Media and ResumableMedia may be set.
func (c *ArchiveInsertCall) Media(r io.Reader, options ...googleapi.MediaOption) *ArchiveInsertCall {
	opts := googleapi.ProcessMediaOptions(options)
	chunkSize := opts.ChunkSize
	r, c.mediaType_ = gensupport.DetermineContentType(r, opts.ContentType)
	c.media_, c.resumableBuffer_ = gensupport.PrepareUpload(r, chunkSize)
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be
// canceled with ctx. ResumableMedia is deprecated in favour of Media.
// At most one of Media and ResumableMedia may be set. mediaType
// identifies the MIME media type of the upload, such as "image/png". If
// mediaType is "", it will be auto-detected. The provided ctx will
// supersede any context previously provided to the Context method.
func (c *ArchiveInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *ArchiveInsertCall {
	c.ctx_ = ctx
	rdr := gensupport.ReaderAtToReader(r, size)
	rdr, c.mediaType_ = gensupport.DetermineContentType(rdr, mediaType)
	c.resumableBuffer_ = gensupport.NewResumableBuffer(rdr, googleapi.DefaultUploadChunkSize)
	c.media_ = nil
	c.mediaSize_ = size
	return c
}

// ProgressUpdater provides a callback function that will be called
// after every chunk. It should be a low-latency function in order to
// not slow down the upload operation. This should only be called when
// using ResumableMedia (as opposed to Media).
func (c *ArchiveInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *ArchiveInsertCall {
	c.progressUpdater_ = pu
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ArchiveInsertCall) Fields(s ...googleapi.Field) *ArchiveInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
// This context will supersede any context previously provided to the
// ResumableMedia method.
func (c *ArchiveInsertCall) Context(ctx context.Context) *ArchiveInsertCall {
	c.ctx_ = ctx
	return c
}

func (c *ArchiveInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{groupId}/archive")
	if c.media_ != nil || c.resumableBuffer_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		protocol := "multipart"
		if c.resumableBuffer_ != nil {
			protocol = "resumable"
		}
		c.urlParams_.Set("uploadType", protocol)
	}
	urls += "?" + c.urlParams_.Encode()
	body = new(bytes.Buffer)
	ctype := "application/json"
	if c.media_ != nil {
		var combined io.ReadCloser
		combined, ctype = gensupport.CombineBodyMedia(body, ctype, c.media_, c.mediaType_)
		defer combined.Close()
		body = combined
	}
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"groupId": c.groupId,
	})
	if c.resumableBuffer_ != nil {
		req.Header.Set("X-Upload-Content-Type", c.mediaType_)
	}
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "groupsmigration.archive.insert" call.
// Exactly one of *Groups or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Groups.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ArchiveInsertCall) Do(opts ...googleapi.CallOption) (*Groups, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	if c.resumableBuffer_ != nil {
		loc := res.Header.Get("Location")
		rx := &gensupport.ResumableUpload{
			Client:    c.s.client,
			UserAgent: c.s.userAgent(),
			URI:       loc,
			Media:     c.resumableBuffer_,
			MediaType: c.mediaType_,
			Callback: func(curr int64) {
				if c.progressUpdater_ != nil {
					c.progressUpdater_(curr, c.mediaSize_)
				}
			},
		}
		ctx := c.ctx_
		if ctx == nil {
			ctx = context.TODO()
		}
		res, err = rx.Upload(ctx)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
	}
	ret := &Groups{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Inserts a new mail into the archive of the Google group.",
	//   "httpMethod": "POST",
	//   "id": "groupsmigration.archive.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "message/rfc822"
	//     ],
	//     "maxSize": "16MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/groups/v1/groups/{groupId}/archive"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/groups/v1/groups/{groupId}/archive"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "groupId": {
	//       "description": "The group ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{groupId}/archive",
	//   "response": {
	//     "$ref": "Groups"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.groups.migration"
	//   ],
	//   "supportsMediaUpload": true
	// }

}
