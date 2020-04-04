package main

import (
	"fmt"
	"strings"
)

type commandLineParams struct {
	pathPrefix string            // String to prefix to request path
	importMap  map[string]string // Mapping from .proto file name to import path.
}

// parseCommandLineParams breaks the comma-separated list of key=value pairs
// in the parameter (a member of the request protobuf) into a key/value map.
// It then sets command line parameter mappings defined by those entries.
func parseCommandLineParams(parameter string) (clp commandLineParams, err error) {
	clp = commandLineParams{
		importMap: make(map[string]string),
	}
	ps := make(map[string]string)
	for _, p := range strings.Split(parameter, ",") {
		if p == "" {
			continue
		}
		i := strings.Index(p, "=")
		if i < 0 {
			err = fmt.Errorf("invalid parameter %q: expected format of parameter to be k=v", p)
			return
		}
		k := p[0:i]
		v := p[i+1:]
		if v == "" {
			err = fmt.Errorf("invalid parameter %q: expected format of parameter to be k=v", k)
			return
		}
		ps[k] = v
	}

	for k, v := range ps {
		switch {
		case k == "path_prefix":
			clp.pathPrefix = v
		case len(k) > 0 && k[0] == 'M':
			clp.importMap[k[1:]] = v // 1 is the length of 'M'.
		case len(k) > 0 && strings.HasPrefix(k, "go_import_mapping@"):
			clp.importMap[k[18:]] = v // 18 is the length of 'go_import_mapping@'.
		default:
			err = fmt.Errorf("unknown parameter %q", k)
		}
	}

	return
}
