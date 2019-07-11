// snmpbeat is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package snmpbeat

import "time"

// snmpbeat stores Otilio snmpbeaturation loaded from .yaml file
type snmpbeat struct {
	Period       time.Duration       `snmpbeat:"period"`
	Hosts        []string            `snmpbeat:"hosts"`
	Port         uint16              `snmpbeat:"port"`
	Community    string              `snmpbeat:"community"`
	User         string              `snmpbeat:"user"`
	AuthPassword string              `snmpbeat:"authpass"`
	PrivPassword string              `snmpbeat:"privpass"`
	Version      string              `snmpbeat:"version"`
	OIDs         []map[string]string `snmpbeat:"oids"`
}

// Defaultsnmpbeat default snmpbeaturation
var Defaultsnmpbeat = snmpbeat{
	Period:    1 * time.Second,
	Port:      161,
	Community: "public",
	Version:   "2c",
}
