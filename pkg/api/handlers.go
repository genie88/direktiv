package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/vorteil/direktiv/pkg/ingress"
	"google.golang.org/grpc/status"
)

const (
	// GRPCCommandTimeout : timeout for grpc function calls
	GRPCCommandTimeout = 30 * time.Second
)

type Handlers struct {
	s *Server
}

func (h *Handlers) Namespaces(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetNamespaces(ctx, &ingress.GetNamespacesRequest{})
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	h.s.json.Marshal(w, resp)
}

func (h *Handlers) AddNamespace(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.AddNamespace(ctx, &ingress.AddNamespaceRequest{
		Name: &n,
	})
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	h.s.json.Marshal(w, resp)
}

func (h *Handlers) DeleteNamespace(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.DeleteNamespace(ctx, &ingress.DeleteNamespaceRequest{
		Name: &n,
	})
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	h.s.json.Marshal(w, resp)
}

func (h *Handlers) NamespaceEvent(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	var contentType string
	if typeMap, ok := r.Header["Content-Type"]; ok {
		contentType = typeMap[0]
	}

	switch contentType {
	case "application/cloudevents+json; charset=utf-8":
	case "application/cloudevents+json":
	case "application/json":
		break
	default:
		errResponse(w, http.StatusUnsupportedMediaType, fmt.Errorf("content type '%s' is not supported. supported media types: 'application/json' ", contentType))
		return
	}

	req := ingress.BroadcastEventRequest{
		Namespace:  &n,
		Cloudevent: b,
	}

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.BroadcastEvent(ctx, &req)
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) Secrets(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetSecrets(ctx, &ingress.GetSecretsRequest{
		Namespace: &n,
	})
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	h.s.json.Marshal(w, resp)
}

func (h *Handlers) CreateSecret(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]

	st := new(NameDataTuple)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	err = json.Unmarshal(b, st)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.StoreSecret(ctx, &ingress.StoreSecretRequest{
		Namespace: &n,
		Name:      &st.Name,
		Data:      []byte(st.Data),
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) DeleteSecret(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]

	st := new(NameDataTuple)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	err = json.Unmarshal(b, st)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.DeleteSecret(ctx, &ingress.DeleteSecretRequest{
		Namespace: &n,
		Name:      &st.Name,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) Registries(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetRegistries(ctx, &ingress.GetRegistriesRequest{
		Namespace: &n,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) CreateRegistry(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]

	st := new(NameDataTuple)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	err = json.Unmarshal(b, st)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.StoreRegistry(ctx, &ingress.StoreRegistryRequest{
		Namespace: &n,
		Name:      &st.Name,
		Data:      []byte(st.Data),
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) DeleteRegistry(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]

	st := new(NameDataTuple)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	err = json.Unmarshal(b, st)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.DeleteRegistry(ctx, &ingress.DeleteRegistryRequest{
		Namespace: &n,
		Name:      &st.Name,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) Workflows(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetWorkflows(ctx, &ingress.GetWorkflowsRequest{
		Namespace: &n,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) GetWorkflow(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]
	id := mux.Vars(r)["workflow"]

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetWorkflowById(ctx, &ingress.GetWorkflowByIdRequest{
		Namespace: &n,
		Id:        &id,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) UpdateWorkflow(w http.ResponseWriter, r *http.Request) {

	uid := mux.Vars(r)["workflowUID"]

	var useRevision bool
	rev, err := strconv.Atoi(r.URL.Query().Get("revision"))
	if err == nil {
		useRevision = true
	}

	var logEvent string
	var useLogEvent bool
	if val, ok := r.URL.Query()["logEvent"]; ok {
		logEvent = val[0]
		useLogEvent = true
	}

	revision := int32(rev)
	active := true

	// Read Body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	var contentType string
	if typeMap, ok := r.Header["Content-Type"]; ok {
		contentType = typeMap[0]
	}

	switch contentType {
	case "text/yaml":
	default:
		errResponse(w, http.StatusUnsupportedMediaType, fmt.Errorf("content type '%s' is not supported. supported media types: 'text/yaml'", contentType))
		return
	}

	// Construct direktiv GRPC Request
	request := ingress.UpdateWorkflowRequest{
		Uid:      &uid,
		Active:   &active,
		Workflow: b,
	}

	if useLogEvent {
		request.LogToEvents = &logEvent
	}

	if useRevision {
		request.Revision = &revision
	}

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.UpdateWorkflow(ctx, &request)
	if err != nil {
		// Convert error
		s := status.Convert(err)
		// Catch when user tries to sent array instead of object
		if strings.HasSuffix(s.Message(), "into map[string]interface {}") {
			errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf("workflow is not a object"))
		} else {
			errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		}
		return
	}

	// Write Data
	w.WriteHeader(http.StatusOK)
	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) ToggleWorkflow(w http.ResponseWriter, r *http.Request) {

	uid := mux.Vars(r)["workflowUID"]

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetWorkflowByUid(ctx, &ingress.GetWorkflowByUidRequest{
		Uid: &uid,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	active := !*resp.Active

	resp2, err := h.s.direktiv.UpdateWorkflow(ctx, &ingress.UpdateWorkflowRequest{
		Uid:      &uid,
		Active:   &active,
		Workflow: resp.Workflow,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	// Write Data
	w.WriteHeader(http.StatusOK)
	if err := h.s.json.Marshal(w, resp2); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) CreateWorkflow(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]
	active := true

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	var contentType string
	if typeMap, ok := r.Header["Content-Type"]; ok {
		contentType = typeMap[0]
	}

	if contentType != "text/yaml" {
		errResponse(w, http.StatusUnsupportedMediaType, fmt.Errorf("content type '%s' is not supported. supported media types: 'text/yaml'", contentType))
		return
	}

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.AddWorkflow(ctx, &ingress.AddWorkflowRequest{
		Namespace: &n,
		Active:    &active,
		Workflow:  b,
	})
	if err != nil {
		s := status.Convert(err)
		if strings.HasSuffix(s.Message(), "into map[string]interface {}") {
			errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf("workflow is not a object"))
		} else {
			errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) DeleteWorkflow(w http.ResponseWriter, r *http.Request) {

	uid := mux.Vars(r)["workflowUID"]

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.DeleteWorkflow(ctx, &ingress.DeleteWorkflowRequest{
		Uid: &uid,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) DownloadWorkflow(w http.ResponseWriter, r *http.Request) {

	uid := mux.Vars(r)["workflowUID"]

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetWorkflowByUid(ctx, &ingress.GetWorkflowByUidRequest{
		Uid: &uid,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+*resp.Id+".yaml")
	w.Header().Set("Content-Type", "application/x-yaml")
	if _, err = w.Write(resp.Workflow); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) ExecuteWorkflow(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]
	uid := mux.Vars(r)["workflowUID"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errResponse(w, 0, err)
		return
	}

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.InvokeWorkflow(ctx, &ingress.InvokeWorkflowRequest{
		Namespace:  &n,
		WorkflowId: &uid,
		Input:      b,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) Instances(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]
	l, o := paginationParams(r)

	if l < 1 {
		l = 10
	}

	if o < 0 {
		o = 0
	}

	limit := int32(l)
	offset := int32(o)

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetWorkflowInstances(ctx, &ingress.GetWorkflowInstancesRequest{
		Namespace: &n,
		Offset:    &offset,
		Limit:     &limit,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) GetInstance(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]
	wid := mux.Vars(r)["workflowID"]
	id := mux.Vars(r)["id"]

	iid := fmt.Sprintf("%s/%s/%s", n, wid, id)

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.GetWorkflowInstance(ctx, &ingress.GetWorkflowInstanceRequest{
		Id: &iid,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) CancelInstance(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]
	wid := mux.Vars(r)["workflowID"]
	id := mux.Vars(r)["id"]

	iid := fmt.Sprintf("%s/%s/%s", n, wid, id)

	ctx, cancel := ctxDeadline()
	defer cancel()

	resp, err := h.s.direktiv.CancelWorkflowInstance(ctx, &ingress.CancelWorkflowInstanceRequest{
		Id: &iid,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}

func (h *Handlers) InstanceLogs(w http.ResponseWriter, r *http.Request) {

	n := mux.Vars(r)["namespace"]
	wid := mux.Vars(r)["workflowID"]
	id := mux.Vars(r)["id"]

	iid := fmt.Sprintf("%s/%s/%s", n, wid, id)

	ctx, cancel := ctxDeadline()
	defer cancel()

	l, o := paginationParams(r)
	if l < 1 {
		l = 10
	}

	if o < 0 {
		o = 0
	}

	limit := int32(l)
	offset := int32(o)

	resp, err := h.s.direktiv.GetWorkflowInstanceLogs(ctx, &ingress.GetWorkflowInstanceLogsRequest{
		InstanceId: &iid,
		Limit:      &limit,
		Offset:     &offset,
	})
	if err != nil {
		s := status.Convert(err)
		errResponse(w, convertGRPCStatusCodeToHTTPCode(s.Code()), fmt.Errorf(s.Message()))
		return
	}

	if err := h.s.json.Marshal(w, resp); err != nil {
		errResponse(w, 0, err)
		return
	}

}
