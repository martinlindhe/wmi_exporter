// returns data points from Win32_PerfRawData_Tcpip_NetworkInterface

// https://technet.microsoft.com/en-us/security/aa394340(v=vs.80) (Win32_PerfRawData_Tcpip_NetworkInterface class)
// https://msdn.microsoft.com/en-us/library/aa394216 (Win32_NetworkAdapter class)
// https://msdn.microsoft.com/en-us/library/aa394353 (Win32_PnPEntity class)

package collector

import (
	"flag"
	"fmt"
	"log"
	"regexp"

	"github.com/StackExchange/wmi"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	Factories["net"] = NewNetworkCollector
}

var (
	nicWhitelist        = flag.String("collector.net.nic-whitelist", ".+", "Regexp of NIC:s to whitelist. NIC name must both match whitelist and not match blacklist to be included.")
	nicBlacklist        = flag.String("collector.net.nic-blacklist", "", "Regexp of NIC:s to blacklist. NIC name must both match whitelist and not match blacklist to be included.")
	nicNameToUnderscore = regexp.MustCompile("[^a-zA-Z0-9]")
)

// A NetworkCollector is a Prometheus collector for WMI Win32_PerfRawData_Tcpip_NetworkInterface metrics
type NetworkCollector struct {
	BytesReceivedTotal       *prometheus.Desc
	BytesSentTotal           *prometheus.Desc
	BytesTotal               *prometheus.Desc
	PacketsOutboundDiscarded *prometheus.Desc
	PacketsOutboundErrors    *prometheus.Desc
	PacketsTotal             *prometheus.Desc
	PacketsReceivedDiscarded *prometheus.Desc
	PacketsReceivedErrors    *prometheus.Desc
	PacketsReceivedTotal     *prometheus.Desc
	PacketsReceivedUnknown   *prometheus.Desc
	PacketsSentTotal         *prometheus.Desc

	nicWhitelistPattern *regexp.Regexp
	nicBlacklistPattern *regexp.Regexp
}

// NewNetworkCollector ...
func NewNetworkCollector() (Collector, error) {
	const subsystem = "net"

	return &NetworkCollector{
		BytesReceivedTotal: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "bytes_received_total"),
			"(Network.BytesReceivedPerSec)",
			[]string{"nic"},
			nil,
		),
		BytesSentTotal: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "bytes_sent_total"),
			"(Network.BytesSentPerSec)",
			[]string{"nic"},
			nil,
		),
		BytesTotal: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "bytes_total"),
			"(Network.BytesTotalPerSec)",
			[]string{"nic"},
			nil,
		),
		PacketsOutboundDiscarded: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "packets_outbound_discarded"),
			"(Network.PacketsOutboundDiscarded)",
			[]string{"nic"},
			nil,
		),
		PacketsOutboundErrors: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "packets_outbound_errors"),
			"(Network.PacketsOutboundErrors)",
			[]string{"nic"},
			nil,
		),
		PacketsReceivedDiscarded: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "packets_received_discarded"),
			"(Network.PacketsReceivedDiscarded)",
			[]string{"nic"},
			nil,
		),
		PacketsReceivedErrors: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "packets_received_errors"),
			"(Network.PacketsReceivedErrors)",
			[]string{"nic"},
			nil,
		),
		PacketsReceivedTotal: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "packets_received_total"),
			"(Network.PacketsReceivedPerSec)",
			[]string{"nic"},
			nil,
		),
		PacketsReceivedUnknown: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "packets_received_unknown"),
			"(Network.PacketsReceivedUnknown)",
			[]string{"nic"},
			nil,
		),
		PacketsTotal: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "packets_total"),
			"(Network.PacketsPerSec)",
			[]string{"nic"},
			nil,
		),
		PacketsSentTotal: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "packets_sent_total"),
			"(Network.PacketsSentPerSec)",
			[]string{"nic"},
			nil,
		),

		nicWhitelistPattern: regexp.MustCompile(fmt.Sprintf("^(?:%s)$", *nicWhitelist)),
		nicBlacklistPattern: regexp.MustCompile(fmt.Sprintf("^(?:%s)$", *nicBlacklist)),
	}, nil
}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *NetworkCollector) Collect(ch chan<- prometheus.Metric) error {
	if desc, err := c.collect(ch); err != nil {
		log.Println("[ERROR] failed collecting net metrics:", desc, err)
		return err
	}
	return nil
}

// mangleNetworkName mangles Network Adapter name (non-alphanumeric to _)
// that is used in Win32_PerfRawData_Tcpip_NetworkInterface.
func mangleNetworkName(name string) string {
	return nicNameToUnderscore.ReplaceAllString(name, "_")
}

type Win32_PerfRawData_Tcpip_NetworkInterface struct {
	BytesReceivedPerSec      uint64
	BytesSentPerSec          uint64
	BytesTotalPerSec         uint64
	Name                     string
	PacketsOutboundDiscarded uint64
	PacketsOutboundErrors    uint64
	PacketsPerSec            uint64
	PacketsReceivedDiscarded uint64
	PacketsReceivedErrors    uint64
	PacketsReceivedPerSec    uint64
	PacketsReceivedUnknown   uint64
	PacketsSentPerSec        uint64
}

func (c *NetworkCollector) collect(ch chan<- prometheus.Metric) (*prometheus.Desc, error) {
	var dst []Win32_PerfRawData_Tcpip_NetworkInterface

	q := wmi.CreateQuery(&dst, "")
	if err := wmi.Query(q, &dst); err != nil {
		return nil, err
	}

	for _, nic := range dst {
		if c.nicBlacklistPattern.MatchString(nic.Name) ||
			!c.nicWhitelistPattern.MatchString(nic.Name) {
			continue
		}

		name := mangleNetworkName(nic.Name)
		if name == "" {
			continue
		}

		// Counters
		ch <- prometheus.MustNewConstMetric(
			c.BytesReceivedTotal,
			prometheus.CounterValue,
			float64(nic.BytesReceivedPerSec),
			name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BytesSentTotal,
			prometheus.CounterValue,
			float64(nic.BytesSentPerSec),
			name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BytesTotal,
			prometheus.CounterValue,
			float64(nic.BytesTotalPerSec),
			name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.PacketsOutboundDiscarded,
			prometheus.CounterValue,
			float64(nic.PacketsOutboundDiscarded),
			name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.PacketsOutboundErrors,
			prometheus.CounterValue,
			float64(nic.PacketsOutboundErrors),
			name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.PacketsTotal,
			prometheus.CounterValue,
			float64(nic.PacketsPerSec),
			name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.PacketsReceivedDiscarded,
			prometheus.CounterValue,
			float64(nic.PacketsReceivedDiscarded),
			name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.PacketsReceivedErrors,
			prometheus.CounterValue,
			float64(nic.PacketsReceivedErrors),
			name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.PacketsReceivedTotal,
			prometheus.CounterValue,
			float64(nic.PacketsReceivedPerSec),
			name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.PacketsReceivedUnknown,
			prometheus.CounterValue,
			float64(nic.PacketsReceivedUnknown),
			name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.PacketsSentTotal,
			prometheus.CounterValue,
			float64(nic.PacketsSentPerSec),
			name,
		)
	}

	return nil, nil
}
