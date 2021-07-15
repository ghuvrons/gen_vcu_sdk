package gen_vcu_sdk

import (
	"log"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pudjamansyurin/gen_vcu_sdk/report"
	"github.com/pudjamansyurin/gen_vcu_sdk/util"
)

type StatusListenerFunc func(vin int, online bool) error
type ReportListenerFunc func(vin int, result interface{}) error

type Listener struct {
	logging    bool
	statusFunc StatusListenerFunc
	reportFunc ReportListenerFunc
}

func (l *Listener) status(client mqtt.Client, msg mqtt.Message) {
	l.logPaylod(msg)

	if err := l.statusFunc(parseVin(msg.Topic()), isOnline(msg.Payload())); err != nil {
		log.Fatalf("Status callback error, %v\n", err)
	}
}

func (l *Listener) report(client mqtt.Client, msg mqtt.Message) {
	l.logPaylod(msg)

	rpt := report.New(msg.Payload())
	result, err := rpt.DecodeReport()
	if err != nil {
		log.Fatalf("Can't decode report package, %v\n", err)
	}

	if err := l.reportFunc(parseVin(msg.Topic()), result); err != nil {
		log.Fatalf("Report callback error, %v\n", err)
	}
}

func (l *Listener) logPaylod(msg mqtt.Message) {
	if !l.logging {
		log.Printf("[%s] %s\n", msg.Topic(), util.HexString(msg.Payload()))
	}
}

func parseVin(topic string) int {
	s := strings.Split(topic, "/")
	vin, _ := strconv.Atoi(s[1])

	return vin
}

func isOnline(b []byte) bool {
	return b[0] == '1'
}
