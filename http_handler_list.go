package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apiserver/pkg/authentication/authenticator"
)

var _ http.Handler = &ListNamespacesHandler{}

type ListNamespacesHandler struct {
	log        *slog.Logger
	lister     NamespaceLister
	userHeader string
}

func NewListNamespacesHandler(log *slog.Logger, lister NamespaceLister) http.Handler {
	return &ListNamespacesHandler{
		log:    log,
		lister: lister,
	}
}

func (h *ListNamespacesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Info("received list request")

	ud := r.Context().Value(ContextKeyUserDetails).(*authenticator.Response)

	// retrieve projects as the user
	nn, err := h.lister.ListNamespaces(r.Context(), ud.User.GetName())
	if err != nil {
		serr := &kerrors.StatusError{}
		if errors.As(err, &serr) {
			w.WriteHeader(int(serr.Status().Code))
			h.write(w, []byte(serr.Error()))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		h.write(w, []byte(err.Error()))
		return
	}

	// build response
	// for PoC limited to JSON
	b, err := json.Marshal(nn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.write(w, []byte(err.Error()))
		return
	}

	w.Header().Add(HttpContentType, HttpContentTypeApplication)
	h.write(w, b)
}

func (h *ListNamespacesHandler) write(w http.ResponseWriter, data []byte) bool {
	if _, err := w.Write(data); err != nil {
		h.log.Error("error writing reply", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}
	return true
}
