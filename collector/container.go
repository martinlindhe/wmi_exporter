// returns containers metrics from HCS
// https://github.com/Microsoft/hcsshim

// +build windows

package collector

import (
	"github.com/Microsoft/hcsshim"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

func init() {
	Factories["container"] = NewContainerMetricsCollector
}

// A ContainerMetricsCollector is a Prometheus collector for containers metrics
type ContainerMetricsCollector struct {
	// Presence
	ContainerAvailable *prometheus.Desc
	// memory
	UsageCommitBytes            *prometheus.Desc
	UsageCommitPeakBytes        *prometheus.Desc
	UsagePrivateWorkingSetBytes *prometheus.Desc

	// CPU
	TotalRuntime100ns  *prometheus.Desc
	RuntimeUser100ns   *prometheus.Desc
	RuntimeKernel100ns *prometheus.Desc

	// Network
	BytesReceived          *prometheus.Desc
	BytesSent              *prometheus.Desc
	PacketsReceived        *prometheus.Desc
	PacketsSent            *prometheus.Desc
	DroppedPacketsIncoming *prometheus.Desc
	DroppedPacketsOutgoing *prometheus.Desc
}

func NewContainerMetricsCollector() (Collector, error) {
	const subsystem = "container"
	return &ContainerMetricsCollector{
		ContainerAvailable: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "available"),
			"Available",
			[]string{"container_id"},
			nil,
		),
		UsageCommitBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "memory_usage_commit_bytes"),
			"Memory Usage Commit Bytes",
			[]string{"container_id"},
			nil,
		),
		UsageCommitPeakBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "memory_usage_commit_peak_bytes"),
			"Memory Usage Commit Peak Bytes",
			[]string{"container_id"},
			nil,
		),
		UsagePrivateWorkingSetBytes: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "memory_usage_private_working_set_bytes"),
			"Memory Usage Private Working Set Bytes",
			[]string{"container_id"},
			nil,
		),
		TotalRuntime100ns: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "cpu_usage_seconds_total"),
			"Total Run time in Seconds",
			[]string{"container_id"},
			nil,
		),
		RuntimeUser100ns: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "cpu_usage_seconds_usermode"),
			"Run Time in User mode in Seconds",
			[]string{"container_id"},
			nil,
		),
		RuntimeKernel100ns: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "cpu_usage_seconds_kernelmode"),
			"Run time in Kernel mode in Seconds",
			[]string{"container_id"},
			nil,
		),
		BytesReceived: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "network_receive_bytes_total"),
			"Bytes Received on Interface",
			[]string{"container_id", "interface"},
			nil,
		),
		BytesSent: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "network_transmit_bytes_total"),
			"Bytes Sent on Interface",
			[]string{"container_id", "interface"},
			nil,
		),
		PacketsReceived: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "network_receive_packets_total"),
			"Packets Received on Interface",
			[]string{"container_id", "interface"},
			nil,
		),
		PacketsSent: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "network_transmit_packets_total"),
			"Packets Sent on Interface",
			[]string{"container_id", "interface"},
			nil,
		),
		DroppedPacketsIncoming: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "network_receive_packets_dropped_total"),
			"Dropped Incoming Packets on Interface",
			[]string{"container_id", "interface"},
			nil,
		),
		DroppedPacketsOutgoing: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "network_transmit_packets_dropped_total"),
			"Dropped Outgoing Packets on Interface",
			[]string{"container_id", "interface"},
			nil,
		),
	}, nil
}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *ContainerMetricsCollector) Collect(ch chan<- prometheus.Metric) error {
	if desc, err := c.collect(ch); err != nil {
		log.Error("failed collecting ContainerMetricsCollector metrics:", desc, err)
		return err
	}
	return nil
}

func (c *ContainerMetricsCollector) collect(ch chan<- prometheus.Metric) (*prometheus.Desc, error) {

	// Types Container is passed to get the containers compute systems only
	containers, err := hcsshim.GetContainers(hcsshim.ComputeSystemQuery{Types: []string{"Container"}})
	if err != nil {
		log.Error("Err in Getting containers:", err)
		return nil, err
	}

	if len(containers) == 0 {
		log.Warn("no containers on node")
		return nil, nil
	}

	for _, containerDetails := range containers {
		containerId := containerDetails.ID

		container, err := hcsshim.OpenContainer(containerId)
		if err != nil {
			log.Error("err in opening container: ", containerId)
			continue
		}

		cstats, err := container.Statistics()
		if err != nil {
			log.Error("err in fetching container Statistics: ", containerId)
			continue
		}
		// HCS V1 is for docker runtime. Add the docker:// prefix on container_id
		containerId = "docker://" + containerId

		ch <- prometheus.MustNewConstMetric(
			c.ContainerAvailable,
			prometheus.CounterValue,
			1,
			containerId,
		)
		ch <- prometheus.MustNewConstMetric(
			c.UsageCommitBytes,
			prometheus.GaugeValue,
			float64(cstats.Memory.UsageCommitBytes),
			containerId,
		)
		ch <- prometheus.MustNewConstMetric(
			c.UsageCommitPeakBytes,
			prometheus.GaugeValue,
			float64(cstats.Memory.UsageCommitPeakBytes),
			containerId,
		)
		ch <- prometheus.MustNewConstMetric(
			c.UsagePrivateWorkingSetBytes,
			prometheus.GaugeValue,
			float64(cstats.Memory.UsagePrivateWorkingSetBytes),
			containerId,
		)
		ch <- prometheus.MustNewConstMetric(
			c.TotalRuntime100ns,
			prometheus.CounterValue,
			float64(cstats.Processor.TotalRuntime100ns)*ticksToSecondsScaleFactor,
			containerId,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RuntimeUser100ns,
			prometheus.CounterValue,
			float64(cstats.Processor.RuntimeUser100ns)*ticksToSecondsScaleFactor,
			containerId,
		)
		ch <- prometheus.MustNewConstMetric(
			c.RuntimeKernel100ns,
			prometheus.CounterValue,
			float64(cstats.Processor.RuntimeKernel100ns)*ticksToSecondsScaleFactor,
			containerId,
		)

		if len(cstats.Network) == 0 {
			log.Warn("No Network Stats for container: ", containerId)
			continue
		}

		networkStats := cstats.Network

		for _, interafce := range networkStats {
			ch <- prometheus.MustNewConstMetric(
				c.BytesReceived,
				prometheus.CounterValue,
				float64(interafce.BytesReceived),
				containerId, interafce.EndpointId,
			)
			ch <- prometheus.MustNewConstMetric(
				c.BytesSent,
				prometheus.CounterValue,
				float64(interafce.BytesSent),
				containerId, interafce.EndpointId,
			)
			ch <- prometheus.MustNewConstMetric(
				c.PacketsReceived,
				prometheus.CounterValue,
				float64(interafce.PacketsReceived),
				containerId, interafce.EndpointId,
			)
			ch <- prometheus.MustNewConstMetric(
				c.PacketsSent,
				prometheus.CounterValue,
				float64(interafce.PacketsSent),
				containerId, interafce.EndpointId,
			)
			ch <- prometheus.MustNewConstMetric(
				c.DroppedPacketsIncoming,
				prometheus.CounterValue,
				float64(interafce.DroppedPacketsIncoming),
				containerId, interafce.EndpointId,
			)
			ch <- prometheus.MustNewConstMetric(
				c.DroppedPacketsOutgoing,
				prometheus.CounterValue,
				float64(interafce.DroppedPacketsOutgoing),
				containerId, interafce.EndpointId,
			)
		}
	}

	return nil, nil
}