package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYAML, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYAML)
	return MapHandler(pathMap, fallback), nil
}

type pathYAML struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAML(bytes []byte) ([]pathYAML, error) {
	var pathURLs []pathYAML
	err := yaml.Unmarshal(bytes, &pathURLs)
	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}

func buildMap(yaml []pathYAML) map[string]string {
	YAMLMap := make(map[string]string)
	for _, data := range yaml {
		YAMLMap[data.Path] = data.URL
	}
	return YAMLMap
}
