package naming_test

import (
	"testing"

	"v.io/v23/ipc/version"
	"v.io/v23/naming"
)

func TestFormat(t *testing.T) {
	testcases := []struct {
		network, address string
		opts             []naming.EndpointOpt
		output           string
	}{
		{"tcp", "127.0.0.1:21", []naming.EndpointOpt{}, "@2@tcp@127.0.0.1:21@@@@@"},
		{"tcp", "127.0.0.1:21", []naming.EndpointOpt{naming.NullRoutingID}, "@2@tcp@127.0.0.1:21@00000000000000000000000000000000@@@@"},
		{"tcp", "127.0.0.1:21", []naming.EndpointOpt{version.IPCVersionRange{3, 5}}, "@2@tcp@127.0.0.1:21@@3@5@@"},
		{"tcp", "127.0.0.1:21", []naming.EndpointOpt{naming.ServesMountTableOpt(true)}, "@3@tcp@127.0.0.1:21@@@@m@@"},
		{"tcp", "127.0.0.1:21", []naming.EndpointOpt{naming.ServesMountTableOpt(false)}, "@2@tcp@127.0.0.1:21@@@@@"},
		{"tcp", "127.0.0.1:22", []naming.EndpointOpt{naming.BlessingOpt("batman@dccomics.com")}, "@4@tcp@127.0.0.1:22@@@@s@batman@dccomics.com@@"},
		{"tcp", "127.0.0.1:22", []naming.EndpointOpt{naming.BlessingOpt("batman@dccomics.com"), naming.BlessingOpt("bugs@bunny.com"), naming.ServesMountTableOpt(true)}, "@4@tcp@127.0.0.1:22@@@@m@batman@dccomics.com,bugs@bunny.com@@"},
		{"tcp", "127.0.0.1:22", []naming.EndpointOpt{naming.BlessingOpt("batman@dccomics.com"), naming.BlessingOpt("bugs@bunny.com"), version.IPCVersionRange{5, 7}}, "@4@tcp@127.0.0.1:22@@5@7@s@batman@dccomics.com,bugs@bunny.com@@"},
		{"tcp", "127.0.0.1:22", []naming.EndpointOpt{naming.BlessingOpt("@s@@")}, "@4@tcp@127.0.0.1:22@@@@s@@s@@@@"},
	}
	for i, test := range testcases {
		str := naming.FormatEndpoint(test.network, test.address, test.opts...)
		if str != test.output {
			t.Errorf("%d: unexpected endpoint string: got %q != %q", i, str, test.output)
		}
	}
}
