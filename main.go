package main

import (
	"context"

	log "github.com/hashicorp/go-hclog"

	"github.com/hashicorp/nomad/plugins"
	gitexec "github.com/pnocera/nomad-git-plugin/git"
)

func main() {
	// Serve the plugin
	plugins.ServeCtx(factory)
}

// factory returns a new instance of a nomad driver plugin
func factory(ctx context.Context, log log.Logger) interface{} {
	return gitexec.NewGitExecDriver(ctx, log)
}
