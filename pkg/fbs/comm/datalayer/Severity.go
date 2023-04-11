// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import "strconv"

type Severity int8

const (
	SeverityEmergency     Severity = 0
	SeverityAlert         Severity = 1
	SeverityCritical      Severity = 2
	SeverityError         Severity = 3
	SeverityWarning       Severity = 4
	SeverityNotice        Severity = 5
	SeverityInformational Severity = 6
	SeverityDebug         Severity = 7
)

var EnumNamesSeverity = map[Severity]string{
	SeverityEmergency:     "Emergency",
	SeverityAlert:         "Alert",
	SeverityCritical:      "Critical",
	SeverityError:         "Error",
	SeverityWarning:       "Warning",
	SeverityNotice:        "Notice",
	SeverityInformational: "Informational",
	SeverityDebug:         "Debug",
}

var EnumValuesSeverity = map[string]Severity{
	"Emergency":     SeverityEmergency,
	"Alert":         SeverityAlert,
	"Critical":      SeverityCritical,
	"Error":         SeverityError,
	"Warning":       SeverityWarning,
	"Notice":        SeverityNotice,
	"Informational": SeverityInformational,
	"Debug":         SeverityDebug,
}

func (v Severity) String() string {
	if s, ok := EnumNamesSeverity[v]; ok {
		return s
	}
	return "Severity(" + strconv.FormatInt(int64(v), 10) + ")"
}
