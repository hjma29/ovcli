package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hjma29/ovcli/oneview"
)

func TestShowSPT(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/sptemplate.json")
	}))
	defer ts.Close()

	tc := oneview.NewFakeClient(ts.URL, new(bytes.Buffer))

	NewShowSPTemplateCmd(tc)
	fmt.Println(tc.Out().string())

}
