package urlshort

import (
	"net/http"
	"log"
  "encoding/json"
	"fmt"

	yamlV2 "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request){
		log.Println("Request Path:", r.URL.Path)

		path, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w,r)
		}
	}

}

func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
	  return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
  }

func parseYAML(yaml []byte) (dst[]map[string]string, err error) {
	err = yamlV2.Unmarshal(yaml, &dst)
	return dst, nil
}

func JSONHandler(jsonDoc []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJson, err := parseJSON(jsonDoc)
	if err != nil {
	  return nil, err
	}
	pathMap := buildMap(parsedJson)
	return MapHandler(pathMap, fallback), nil
  }

func parseJSON(jsonDoc []byte) (dst[]map[string]string, err error) {
	err = json.Unmarshal(jsonDoc, &dst)
	return dst, nil
}

func SQLHandler(sqldata []map[string]string, fallback http.Handler) (http.HandlerFunc, error) {

	pathMap := buildMap(sqldata)
	fmt.Println("pathMap: ", pathMap)
	fmt.Println("sqldata: ", sqldata)
	return MapHandler(pathMap, fallback), nil
  }

func buildMap(parsed []map[string]string) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, entry := range parsed {
		key := entry["path"]
		pathsToUrls[key] = entry["url"]
	}
	return pathsToUrls
}
