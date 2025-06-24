package agent

import (
	"sort"

	"luxploit.net/pixlic/cmd"
)

type asaMap map[string]string

func (a asaMap) SortedKeys() []string {
	keys := make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// The ASA license is actually highly configurable but its not really worth it,
// so all values are the maximum according to their capabilities seen here:
// http://web.archive.org/web/20110429171852/http://www.cisco.com/en/US/products/ps6120/prod_models_comparison.html
var asaAppliances = asaMap{
	"ASA 5505": "2A800100", // 0b0000000010010?11?10010000001011?
	"ASA 5510": "????????",
	"ASA 5520": "39000000",
	"ASA 5540": "????????",
	"ASA 5500": "38040000",
}

type ASA struct {
	serial    *int64
	model     *string
	devfuture *string
	devpkey   *string
}

func (ASA) Name() string {
	return "ASA"
}

func (a *ASA) Init(flags *cmd.Flags) {
	a.serial = flags.Int64("serial", 0, "Specify your PIX serial number")
	a.model = flags.String("model", "ASA 5505", "Select your ASA model (only non -X supported for now)")
	a.devfuture = flags.String("dev-future", "", "(DEV) Set custom license future")
	a.devpkey = flags.String("dev-pkey", "0x??", "(DEV) Compare against existing p-key")
}

func (ASA) Support() {

}

func (a *ASA) Run() cmd.ActionStatus {
	return cmd.Action_Success
}
