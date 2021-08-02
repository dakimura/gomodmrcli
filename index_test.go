package gomodmrcli_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/dakimura/gomodmrcli"
)

//type RoundTripFunc func(req *http.Request) *http.Response
//
//func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
//	return f(req), nil
//}
//
//func NewTestClient(fn RoundTripFunc) *http.IndexClient {
//	return &http.IndexClient{
//		Transport: RoundTripFunc(fn),
//	}
//}

func TestClient_Index(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		hc                 *http.Client
		since              time.Time
		limit              int
		disableModuleFetch bool
		wantLen            int
		wantErr            bool
	}{
		"ok": {
			hc:                 http.DefaultClient,
			since:              time.Date(2020, 8, 1, 12, 34, 56, 0, time.UTC),
			disableModuleFetch: true,
			wantLen:            2000,
			wantErr:            false,
		},
		"ok/ when limit=10, len(response)=10": {
			hc:      http.DefaultClient,
			limit:   10,
			wantLen: 10,
			wantErr: false,
		},
	}
	for name, tt := range tests {
		name := name
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// --- given ---
			c := gomodmrcli.NewIndexClient(tt.hc)
			// --- when ---
			got, err := c.Index(tt.since, tt.limit, tt.disableModuleFetch)

			// --- when ---
			if (err != nil) != tt.wantErr {
				t.Errorf("Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("len(got) = %v, wantLen = %v", len(got), tt.wantLen)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Index() got = %v, want %v", got, tt.want)
			//}

			//fmt.Println(got)
		})
	}
}
