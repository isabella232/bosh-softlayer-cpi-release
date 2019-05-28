package dispatcher

import (
	"bytes"
	"encoding/json"

	bslaction "bosh-softlayer-cpi/action"
	bslapi "bosh-softlayer-cpi/api"
)

const (
	jsonLogTag = "json"

	jsonCloudErrorType          = "Bosh::Clouds::CloudError"
	jsonCpiErrorType            = "Bosh::Clouds::CpiError"
	jsonNotImplementedErrorType = "Bosh::Clouds::NotImplemented"
)

type Request struct {
	Method    string        `json:"method"`
	Arguments []interface{} `json:"arguments"`

	// context key is ignored
}

type Response struct {
	Result interface{}    `json:"result"`
	Error  *ResponseError `json:"error"`

	Log string `json:"log"`
}

type ResponseError struct {
	Type    string `json:"type"`
	Message string `json:"message"`

	CanRetry bool `json:"ok_to_retry"`
}

func (r ResponseError) Error() string {
	return r.Message
}

type JSON struct {
	actionFactory bslaction.Factory
	caller        Caller
	logger        bslapi.MultiLogger
}

func NewJSON(
	actionFactory bslaction.Factory,
	caller Caller,
	logger        bslapi.MultiLogger,
) JSON {
	return JSON{
		actionFactory: actionFactory,
		caller:        caller,
		logger:        logger,
	}
}

func (c JSON) Dispatch(reqBytes []byte) []byte {
	var req Request

	c.logger.DebugWithDetails(jsonLogTag, "Request bytes", "")

	digitalDecoder := json.NewDecoder(bytes.NewReader(reqBytes))
	digitalDecoder.UseNumber()
	if err := digitalDecoder.Decode(&req); err != nil {
		return c.buildCpiError("Must provide valid JSON payload")
	}

	c.logger.DebugWithDetails(jsonLogTag, "Deserialized request", "")

	if req.Method == "" {
		return c.buildCpiError("Must provide method key")
	}

	if req.Arguments == nil {
		return c.buildCpiError("Must provide arguments key")
	}

	action, err := c.actionFactory.Create(req.Method)
	if err != nil {
		return c.buildNotImplementedError()
	}

	result, err := c.caller.Call(action, req.Arguments)
	if err != nil {
		return c.buildCloudError(err)
	}

	resp := Response{
		Result: result,
		Log:    c.logger.LogBuff.String(),
	}

	c.logger.DebugWithDetails(jsonLogTag, "Deserialized response", "")

	respBytes, err := json.Marshal(resp)
	if err != nil {
		return c.buildCpiError("Failed to serialize result")
	}

	c.logger.DebugWithDetails(jsonLogTag, "Response bytes", "")

	return respBytes
}

func (c JSON) buildCloudError(err error) []byte {
	respErr := Response{
		Log:   c.logger.LogBuff.String(),
		Error: &ResponseError{},
	}

	
	if typedErr, ok := err.(bslapi.CloudError); ok {
		respErr.Error.Type = typedErr.Type()
	} else {
		respErr.Error.Type = jsonCloudErrorType
	}

	respErr.Error.Message = err.Error()

	if typedErr, ok := err.(bslapi.RetryableError); ok {
		respErr.Error.CanRetry = typedErr.CanRetry()
	}

	respErrBytes, err := json.Marshal(respErr)
	if err != nil {
		panic(err)
	}

	c.logger.DebugWithDetails(jsonLogTag, "CloudError response bytes", string(respErrBytes))

	return respErrBytes
}

func (c JSON) buildCpiError(message string) []byte {
	respErr := Response{
		Log: c.logger.LogBuff.String(),
		Error: &ResponseError{
			Type:    jsonCpiErrorType,
			Message: message,
		},
	}

	respErrBytes, err := json.Marshal(respErr)
	if err != nil {
		panic(err)
	}

	c.logger.DebugWithDetails(jsonLogTag, "CpiError response bytes", string(respErrBytes))

	return respErrBytes
}

func (c JSON) buildNotImplementedError() []byte {
	respErr := Response{
		Log: c.logger.LogBuff.String(),
		Error: &ResponseError{
			Type:    jsonNotImplementedErrorType,
			Message: "Must call implemented method",
		},
	}

	respErrBytes, err := json.Marshal(respErr)
	if err != nil {
		panic(err)
	}

	c.logger.DebugWithDetails(jsonLogTag, "NotImplementedError response bytes", string(respErrBytes))

	return respErrBytes
}
