package authenticator

import (
	"net/url"
	"testing"
)

type mockFormatter string

func (mf mockFormatter) AsUrlValues(issuer string) url.Values {
	params := url.Values{}
	params.Add("self", string(mf))
	params.Add("issuer", issuer)
	return params
}

func TestLabel_String(t *testing.T) {
	l := Label{
		AccountName: "$0mbody@nowh3re.net",
		Issuer:      "Nowhere Inc.",
	}

	expected := "Nowhere Inc.:$0mbody@nowh3re.net"

	if expected != l.String() {
		t.Errorf("unexpected string\nexpected: %s\n  actual: %s", expected, l.String())
	}
}

func TestKeyUri_String(t *testing.T) {
	cases := []struct {
		label    string
		name     string
		issuer   string
		otpType  string
		params   paramsFormatter
		expected string
	}{
		{
			"Nothing Weird",
			"john",
			"company",
			"mock",
			mockFormatter("vanilla"),
			"otpauth://mock/company:john?issuer=company&self=vanilla",
		},
		{
			"Special Chars",
			"J0hn@example.com",
			"Example Co.",
			"test",
			mockFormatter("$p3c!al C#4rs"),
			"otpauth://test/Example%20Co.:J0hn@example.com?issuer=Example+Co.&self=%24p3c%21al+C%234rs",
		},
	}

	for _, c := range cases {
		ku := KeyUri{
			Type: c.otpType,
			Label: Label{
				AccountName: c.name,
				Issuer:      c.issuer,
			},
			Parameters: c.params,
		}

		kus := ku.String()

		if c.expected != kus {
			t.Errorf("unexpected string\nexpected: %s\n  actual: %s", c.expected, kus)
		}
	}
}
