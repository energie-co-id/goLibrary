package influxLib

import (
	"context"
	"fmt"
	"log"

	// "os"
	"strings"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func Query(influxDBURL string, authToken string, org string, start int64, stop int64, topic []string, field []string, interval string, fn string) (*api.QueryTableResult, error) {
	// Create a new client
	client := influxdb2.NewClient(influxDBURL, authToken)
	queryAPI := client.QueryAPI(org)

	// Define your query
	query := fmt.Sprintf(`from(bucket: "evgate")
  |> range(start: %d, stop: %d)
  |> filter(fn: (r) => r["_measurement"] == "mqtt_consumer")
  |> filter(fn: (r) => r["topic"] == "%s")
  |> filter(fn: (r) => r["_field"] == "%s")
  |> aggregateWindow(every: %s, fn: %s, createEmpty: false)
  |> yield(name: "%s")`, start, stop, strings.Join(topic, `" or r["topic"] == "`), strings.Join(field, `" or r["_field"] == "`), interval, fn, fn)

	// if strings.HasSuffix(os.Getenv("STAGE"), "dev") {
	// 	log.Println("query", query)
	// }
	log.Println("query", query)
	// Execute the query
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		fmt.Printf("Query error: %v\n", err)
		return nil, err
	}

	// Close the client
	client.Close()

	return result, err
}

func AirReg(types string, code string) []string {
	onePhase := map[string][]string{
		"active_energy":               {"reg_11", "reg_10"},
		"reactive_energy":             {"reg_13"},
		"terminal_temperature_n":      {"reg_1"},
		"phase_terminal_temperature":  {"reg_0"},
		"MCU_temperature":             {"reg_2"},
		"power":                       {"reg_6", "reg_5"},
		"reactive_power":              {"reg_8"},
		"apparent_power":              {"reg_19"},
		"power_factor":                {"reg_14"},
		"current":                     {"reg_4"},
		"leakage_current_value":       {"reg_22"},
		"leakage_event_current_value": {"reg_23"},
		"voltage":                     {"reg_3"},
		"total_harmonic_of_current":   {"reg_27"},
		"over_current":                {"reg_20"},
		"main_frequency":              {"reg_9"},
		"breaker_open":                {"reg_32"},
		"breaker_status":              {"reg_15"},
		"rssi":                        {"rssi"},
		"alarm":                       {"alarm"},
		"enbyte":                      {"enbyte"},
	}
	threePhase := map[string][]string{
		"active_energy":               {"reg_27", "reg_26"},
		"active_energy_r":             {"reg_21", "reg_20"},
		"active_energy_s":             {"reg_23", "reg_22"},
		"active_energy_t":             {"reg_25", "reg_24"},
		"reactive_energy":             {"reg_60", "reg_59"},
		"terminal_temperature_n":      {"reg_1"},
		"phase_terminal_temperature":  {"reg_4"},
		"phase_terminal_temperature2": {"reg_3"},
		"phase_terminal_temperature3": {"reg_2"},
		"MCU_temperature":             {"reg_0"},
		"power":                       {"reg_31", "reg_30"},
		"power_r":                     {"reg_13", "reg_12"},
		"power_s":                     {"reg_15", "reg_14"},
		"power_t":                     {"reg_19", "reg_18"},
		"reactive_power":              {"reg_44", "reg_43"},
		"apparent_power":              {"reg_52", "reg_51"},
		"power_factor":                {"reg_36"},
		"current":                     {"reg_8"},
		"current2":                    {"reg_9"},
		"current3":                    {"reg_10"},
		"current_n":                   {"reg_11"},
		"leakage_current_value":       {"reg_68"},
		"voltage":                     {"reg_5"},
		"voltage2":                    {"reg_6"},
		"voltage3":                    {"reg_7"},
		"total_harmonic_of_current":   {"reg_65"},
		"total_harmonic_of_current2":  {"reg_66"},
		"total_harmonic_of_current3":  {"reg_67"},
		"over_current":                {"reg_61"},
		"main_frequency":              {"reg_32"},
		"breaker_open":                {"reg_28"},
		"breaker_status":              {"reg_29"},
		"rssi":                        {"rssi"},
		"alarm":                       {"alarm"},
		"enbyte":                      {"enbyte"},
	}
	pro := map[string][]string{
		"standmeter": {"stm"},
		"lastcredit": {"rmn"},
		"power":      {"pwr"},
		"voltage":    {"v"},
		"current":    {"i"},
		"current2":   {"s1"},
		// "current2":   {"s2"},
		"power_factor": {"pf"},
		"rssi":         {"rssi"},
		"alarm":        {"alarm"},
		"enbyte":       {"enbyte"},
	}
	if types == "onePhase" {
		if val, ok := onePhase[code]; ok {
			return val
		}
	}
	if types == "threePhase" {
		if val, ok := threePhase[code]; ok {
			return val
		}
	}
	if types == "pro" {
		if val, ok := pro[code]; ok {
			return val
		}
	}

	return []string{"not found"}
}

func ModbusHeartbeat(types string, code string) []string {
	threePhase := map[string][]string{
		"charger_sn":                  {"sn"},
		"charging_state":              {"charging-state"},
		"charger_status":              {"charger-status"},
		"charger_iso_state":           {"iso-state"},
		"charger_discharged_energy":   {"discharged-energy"},
		"charger_charged_energy":      {"charged-energy"},
		"charging_time":               {"charging-time"},
		"active_energy":               {"energy-total"},
		"active_energy_r":             {"energy-l1"},
		"active_energy_s":             {"energy-l2"},
		"active_energy_t":             {"energy-l3"},
		"reactive_energy":             {"reactive-energy-total"},
		"terminal_temperature_n":      {"temp-neutral"},
		"phase_terminal_temperature":  {"temp-l1"},
		"phase_terminal_temperature2": {"temp-l2"},
		"phase_terminal_temperature3": {"temp-l3"},
		"MCU_temperature":             {"temp-mcu"},
		"power":                       {"power"},
		"power_r":                     {"power-l1"},
		"power_s":                     {"power-l2"},
		"power_t":                     {"power-l3"},
		"reactive_power":              {"reactive-power-total"},
		"apparent_power":              {"apparent-power-total"},
		"power_factor":                {"pf-total"},
		"current":                     {"current-l1"},
		"current2":                    {"current-l2"},
		"current3":                    {"current-l3"},
		"current_n":                   {"current-neutral"},
		"leakage_current_value":       {"leakage-current"},
		"voltage":                     {"voltage-l1"},
		"voltage2":                    {"voltage-l2"},
		"voltage3":                    {"voltage-l3"},
		"total_harmonic_of_current":   {"thd-l1"},
		"total_harmonic_of_current2":  {"thd-l2"},
		"total_harmonic_of_current3":  {"thd-l3"},
		"over_current":                {"over-current"},
		"main_frequency":              {"main-freq"},
		"breaker_open":                {"switch-pos"},
		"breaker_status":              {"breaker-state"},
		"rssi":                        {"rssi"},
		"alarm":                       {"alarm1", "alarm2"},
		"enbyte":                      {"enbyte"},
	}
	if types == "threePhase" {
		if val, ok := threePhase[code]; ok {
			return val
		}
	}

	return []string{"not found"}

}

func EvgateProHeartbeat(code string) string {
	lib := map[string]string{
		"output": "output",
		"online": "online",
		"stm":    "stm",
		"rmn":    "rmn",
		"pwr":    "pwr",
		"v":      "v",
		"i":      "i",
		"s1":     "s1",
		"s2":     "s2",
		"pf":     "pf",
		"sn":     "sn",
		"rssi":   "rssi",
	}
	if val, ok := lib[code]; ok {
		return val
	}
	return "not found"
}

func TopicPrefix(code string) string {
	lib := map[string]string{
		"default":   "gateway",
		"modbus_cp": "modbus/CP-01",
	}
	if val, ok := lib[code]; ok {
		return val
	}
	return "not found"
}
