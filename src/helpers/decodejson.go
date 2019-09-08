package helpers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func DecodeJson(c *gin.Context, v interface{}) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &v); err != nil {
		return err
	}

	// reset the response body to initial state
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return nil
}
