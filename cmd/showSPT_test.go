package cmd

import (
	"bytes"
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hjma29/ovcli/oneview"
)

// func TestMain(m *testing.M) {

// 	os.Exit(m.Run())

// }

func TestShowSPT(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc(oneview.SPTemplateURL, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/sptemplate.raw")
	})
	mux.HandleFunc(oneview.EGURL, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/eg.raw")
	})
	mux.HandleFunc(oneview.ServerHWTypeURL, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/serverhwt.raw")
	})

	ts := httptest.NewServer(mux)

	defer ts.Close()

	tc := oneview.NewFakeClient(ts.URL, new(bytes.Buffer))

	cmd := NewShowSPTemplateCmd(tc)
	//cmd.Flags().Set("debug", "true")
	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}

	out, err := ioutil.ReadAll(tc.Out)
	if err != nil {
		t.Error("couldn't read fake client output buffer")
	}
	//fmt.Println(string(out))

	gf, err := ioutil.ReadFile("testdata/showspt.golden")
	if err != nil {
		t.Error("could not open file \"showspt.golden\" file")
	}
	//fmt.Println(string(gf))

	if !bytes.Equal(out, gf) {
		t.Error("show sptemplate displays different than golden file")

	}

	// if _, err := io.Copy(os.Stdout, tc.Out); err != nil {
	// 	log.Fatal(err)
	// }

}

func init() {

	flag.BoolVar(&Debugmode, "d", false, "turn on debug info display")
	flag.Parse()

	initConfig()

}
