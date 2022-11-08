package steps

import (
	"bytes"
	"github.com/cucumber/godog"
	"github.com/dimayaschhu/vocabulary/pkg/utils"
	"io"
	"net/http"
	"strings"
)

type HTTPStepHandler struct {
	client             *http.Client
	responseStatusCode int
	response           string
	objectMatcher      *utils.ObjectMatcher
}

func NewHTTPStepHandler(objectMatcher *utils.ObjectMatcher) *HTTPStepHandler {
	client := &http.Client{}

	return &HTTPStepHandler{client: client, objectMatcher: objectMatcher}
}

func (h *HTTPStepHandler) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)" with JSON body:`, h.sendJSONRequest)
	ctx.Step(`^method should return status code (\d+) and JSON response:`, h.shouldReturnJsonResponse)
}
func (h *HTTPStepHandler) shouldReturnJsonResponse(statusCode int, expectedJson string) error {

	return nil
}

func (h *HTTPStepHandler) sendJSONRequest(method string, path string, reqData string) error {
	req, err := http.NewRequest(method, "http://localhost:8080"+path, strings.NewReader(reqData))
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json")

	response, err := h.client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		return err
	}

	h.responseStatusCode = response.StatusCode
	h.response = buf.String()

	return nil
}
