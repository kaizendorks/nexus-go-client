package models

type StackTraceElement struct {
	ClassName    string `json:"className"`
	FileName     string `json:"fileName"`
	LineNumber   int32  `json:"lineNumber"`
	MethodName   string `json:"methodName"`
	NativeMethod bool   `json:"nativeMethod"`
}

type Throwable struct {
	Cause            *Throwable           `json:"cause"`
	LocalizedMessage string               `json:"localizedMessage"`
	Message          string               `json:"message"`
	StackTrace       []*StackTraceElement `json:"stackTrace"`
	Suppressed       []*Throwable         `json:"suppressed"`
}

type SystemStatus struct {
	Details   map[string]interface{} `json:"details"`
	Duration  int64                  `json:"duration"`
	Error     *Throwable             `json:"error"`
	Healthy   bool                   `json:"healthy"`
	Message   string                 `json:"message"`
	Time      int64                  `json:"time"`
	Timestamp string                 `json:"timestamp"`
}

type StatusCheckResponse map[string]SystemStatus
