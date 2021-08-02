package gomodmrcli_test

import (
	"github.com/dakimura/gomodmrcli"
	"net/http"
	"testing"
)

func TestClient_Proxy(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		hc                 *http.Client
		modulePath         string
		version            string
		disableModuleFetch bool
		wantErr            bool
	}{
		"ok": {
			hc:                 http.DefaultClient,
			modulePath:         "github.com/gin-gonic/gin",
			version:            "v1.7.2",
			disableModuleFetch: true,
			wantErr:            false,
		},
		"ok/a module that doesn't have go.mod": {
			hc:                 http.DefaultClient,
			modulePath:         "github.com/davecgh/go-spew",
			version:            "v1.1.1",
			disableModuleFetch: true,
			wantErr:            false,
		},
	}
	for name, tt := range tests {
		name := name
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// --- given ---
			c := gomodmrcli.NewProxyClient(tt.hc)
			// --- when ---
			got, err := c.Mod(tt.modulePath, tt.version, tt.disableModuleFetch)

			// --- when ---
			if (err != nil) != tt.wantErr {
				t.Fatalf("Mod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Module.Mod.Path != tt.modulePath {
				t.Fatalf("Module.Mod.Path = %v, modulePath = %v", got.Module.Mod.Path, tt.modulePath)
			}
		})
	}
}
