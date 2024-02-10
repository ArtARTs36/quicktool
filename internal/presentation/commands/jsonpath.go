package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PaesslerAG/jsonpath"

	"github.com/artarts36/quicktool/internal/infrastructure/filesystem"
	"github.com/artarts36/quicktool/internal/presentation/interaction"
)

type JSONPath struct {
	fs filesystem.FileSystem
}

func NewJSONPath(
	fs filesystem.FileSystem,
) *JSONPath {
	return &JSONPath{
		fs: fs,
	}
}

func (c *JSONPath) Definition() *interaction.Definition {
	return &interaction.Definition{
		Name:        "jsonpath",
		Description: "Get value from json by path",
		Args: []*interaction.DefinitionArg{
			{
				Name:        "path",
				Description: "json path",
				Required:    true,
			},
			{
				Name:        "source",
				Description: "json path (string of content, file path)",
				Required:    false,
			},
		},
	}
}

func (c *JSONPath) Execute(_ *interaction.Context, env *interaction.Env) error {
	path := env.Input.Argument("path")
	source := env.Input.Argument("source")

	if path != "" && !strings.Contains(path, "$.") {
		path = fmt.Sprintf("$.%s", path)
	}

	sourceJSONVal, err := c.getJSONInterface(source, env)
	if err != nil {
		return fmt.Errorf("unable scan json: %s", err.Error())
	}

	val, err := jsonpath.Get(path, sourceJSONVal)
	if err != nil {
		return fmt.Errorf("unable to get value by path: %s", err.Error())
	}

	env.PrintText(fmt.Sprintf("value: %v", val))

	return nil
}

func (c *JSONPath) getJSONInterface(source string, env *interaction.Env) (interface{}, error) {
	if source == "" {
		return c.captureJSONFromInput(env)
	}

	if fileExists, err := c.fs.Exists(source); err != nil {
		return nil, fmt.Errorf("unable to check file existence")
	} else if fileExists {
		return c.unmarshalFile(source)
	}

	sourceURL, err := url.Parse(source)
	if err != nil {
		return fmt.Errorf("invalid url: %s", err), nil
	}

	resp, err := http.DefaultClient.Do(&http.Request{
		URL: sourceURL,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to execute request: %s", err.Error())
	}

	if err = resp.Body.Close(); err != nil {
		return nil, fmt.Errorf("unable to close request: %s", err)
	}

	return c.unmarshalHTTPResponse(resp)
}

func (c *JSONPath) captureJSONFromInput(env *interaction.Env) (interface{}, error) {
	sourceJSONVal := interface{}(nil)
	reader := bufio.NewReader(os.Stdin)

	env.PrintText("Enter JSON: ")
	source, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("unable to read input source: %s", err.Error())
	}

	err = json.Unmarshal(source, &sourceJSONVal)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON: %s", err.Error())
	}

	return sourceJSONVal, nil
}

func (c *JSONPath) unmarshalFile(path string) (interface{}, error) {
	sourceJSONVal := interface{}(nil)

	bs, err := c.fs.GetContent(path)
	if err != nil {
		return nil, fmt.Errorf("unable to load file: %s", err.Error())
	}

	err = json.Unmarshal(bs, &sourceJSONVal)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON: %s", err.Error())
	}

	return sourceJSONVal, nil
}

func (c *JSONPath) unmarshalHTTPResponse(response *http.Response) (interface{}, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %s", err.Error())
	}

	sourceJSONVal := interface{}(nil)
	err = json.Unmarshal(body, &sourceJSONVal)
	if err != nil {
		return nil, fmt.Errorf("given invalid json: %s", err.Error())
	}

	return sourceJSONVal, nil
}
