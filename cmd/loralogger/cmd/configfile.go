package cmd

import (
	"os"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/rbricheno/lora-logger/internal/config"
)

const configTemplate = `[general]
# Log level
#
# debug=5, info=4, warning=3, error=2, fatal=1, panic=0
log_level={{ .General.LogLevel }}


[lora_logger]
# Bind
#
# The interface:port on which the lora-logger will bind for receiving
# data from the packet-forwarder (UDP data).
bind="{{ .LoraLogger.Bind }}"


# Backends
#
# The backends to which the lora-logger will forward the
# packet-forwarder UDP data.
#
# Example:
# [[lora_logger.backend]]
# # Host
# #
# # The host:IP of the backend.
# host="192.16.1.5:1700"
# 
# # Gateway IDs
# #
# # The Gateway IDs to forward data for.
# gateway_ids = [
#   "0101010101010101",
#   "0202020202020202",
# ]
{{ range $index, $element := .LoraLogger.Backends }}
[[lora_logger.backend]]
host="{{ $element.Host }}"

gateway_ids = [
{{ range $index, $element := $element.GatewayIDs -}}
  "{{ $element }}",
{{ end -}}
]
{{ end }}
`

var configCmd = &cobra.Command{
	Use:   "configfile",
	Short: "Print the LoRa Server configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(configTemplate))
		err := t.Execute(os.Stdout, &config.C)
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}
