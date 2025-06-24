package agent

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"luxploit.net/pixlic/cmd"
)

type pixEntry struct {
	future       string
	hasEModel    bool
	supportLevel string
}

type pixMap map[string]pixEntry

func (a pixMap) SortedKeys() []string {
	keys := make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

var pixAppliances = pixMap{
	"PIX 501": {future: "2A800100", supportLevel: Support_Untested},
	"PIX 506": {future: "????????", supportLevel: Support_Unsupported, hasEModel: true},
	"PIX 515": {future: "39000000", supportLevel: Support_Working, hasEModel: true},
	"PIX 520": {future: "????????", supportLevel: Support_Unsupported},
	"PIX 525": {future: "38040000", supportLevel: Support_Working},
	"PIX 535": {future: "39000000", supportLevel: Support_Untested},
}

type PIX struct {
	serial        *int64
	model         *string
	devbruteforce *bool
	devfuture     *string
	devpkey       *string

	working_iterations int
	working_hexSerial  string
	working_future     string
	working_outputPkey string
}

func (PIX) Name() string {
	return "PIX"
}

func (p *PIX) Init(flags *cmd.Flags) {
	p.serial = flags.Int64("serial", 0, "Specify your PIX serial number")
	p.model = flags.String("model", "PIX 515", "Select your PIX model")
	p.devbruteforce = flags.Bool("dev-force", false, "(DEV) Bruteforce license future value")
	p.devfuture = flags.String("dev-future", "", "(DEV) Set custom license future")
	p.devpkey = flags.String("dev-pkey", "0x??", "(DEV) Compare against existing p-key")
}

func (PIX) Support() {
	for _, key := range pixAppliances.SortedKeys() {
		appl := pixAppliances[key]
		name := key
		if appl.hasEModel {
			name += "(e)"
		}

		fmt.Printf("\t%-20s\t%-50s (Future: %s)\n", name, appl.supportLevel, appl.future)
	}
}

func (p *PIX) Run() cmd.ActionStatus {
	if *p.serial == 0 {
		fmt.Println("Invalid/No License Key provided!")
		return cmd.Action_Invalid
	}
	p.working_hexSerial = invertHexBytes(strconv.FormatInt(*p.serial, 16))

	appliance, ok := pixAppliances[*p.model]
	if !ok {
		fmt.Println("Invalid/No PIX model provided!")
		return cmd.Action_Invalid
	}
	p.working_future = appliance.future

	if *p.devfuture != "" {
		p.working_future = *p.devfuture
	}

	p.working_iterations = 1
	fmt.Println("Generating...")
	if *p.devbruteforce {
		p.bruteforceFuture()
	} else {
		p.generateSerial()
	}

	fmt.Printf("> Device:         %s\n", *p.model)
	fmt.Printf("> Future:         %s\n", p.working_future)
	fmt.Printf("> Iterations:     %d\n", p.working_iterations)
	fmt.Printf("> Serial Number:  %d (%s)\n", *p.serial, p.working_hexSerial)
	fmt.Printf("> Activation Key: %s\n", p.working_outputPkey)
	fmt.Println("Generation successful!")

	return cmd.Action_Success
}

func (p *PIX) bruteforceFuture() {
	for idx := range 0x5F5E0FF {
		p.working_future = invertHexBytes(fmt.Sprintf("%08X", idx))
		if p.generateSerial() != nil {
			continue
		}

		if *p.devpkey == p.working_outputPkey {
			return
		}

		p.working_iterations += 1
	}
}

func (p *PIX) generateSerial() error {
	data, err := hex.DecodeString(p.working_future + p.working_hexSerial)
	if err != nil {
		return err
	}
	hash := fmt.Sprintf("%x", md5.Sum(data))

	for idx := 0; idx < 16; idx += 4 {
		part := invertHexBytes(hash[(idx * 2):((idx + 4) * 2)])
		p.working_outputPkey += "0x" + part + " "
	}
	p.working_outputPkey = strings.TrimRight(p.working_outputPkey, " ")
	return nil
}
