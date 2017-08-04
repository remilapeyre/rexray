// +build !client
// +build !controller
// +build libstorage_integration_driver_linux

package module

import (
	gofigCore "github.com/akutz/gofig"
	gofig "github.com/akutz/gofig/types"
)

func init() {
	cfg := gofigCore.NewRegistration("Module")
	cfg.SetYAML(`
rexray:
    modules:
        default-docker:
            type:     docker
            desc:     The default docker module.
            host:     unix:///run/docker/plugins/rexray.sock
            spec:     /etc/docker/plugins/rexray.spec
            disabled: false
`)
	cfg.Key(gofig.String, "", "10s", "", "rexray.module.startTimeout")
	gofigCore.Register(cfg)
}
