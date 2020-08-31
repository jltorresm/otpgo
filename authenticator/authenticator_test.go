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

func TestKeyUri_Base64QR(t *testing.T) {
	ku := KeyUri{
		Type: "mock",
		Label: Label{
			AccountName: "J0hn@example.com",
			Issuer:      "Example Co.",
		},
		Parameters: mockFormatter("$p3c!al C#4rs"),
	}

	expected := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAAB" +
		"mvDolAAAABlBMVEX///8AAABVwtN+AAACtklEQVR42uyZzY3sIBCEC3HoIyGQiUlstNiax" +
		"JhMCIEjB+R66p4f7+oFYCxN35b9DgNuV1VjfOtbFywhyQfJNazwRJQiO5yutgsBFZBHQ9i" +
		"arzcgAmR3DfAzAaQ82hJ2ZJKeSABcy7xPBhSk51HXAaAvUwLaDxvJgRFZeurXAwApWMIa1" +
		"nivvnZ066A/HXU2YK9eW8LWbvFeR5TSF9sW5wGeFR7tJ4w4MNDRXf9P684FpNo+CnLLZPU" +
		"0fVjwEz4iNgFAEokqYysyfRUKWcLnoC8BIGoT88E1sAKsSNxlD2y3eQD9ydYwgK++et0DO" +
		"wLb8eqdDgAd2jBh48ANCgCqzEdfzwCYW+xh134wRZNdF5DbdQCpHaL/fluz2sdfvzgfQNR" +
		"Mo7uASjGrlO7I9hGQOQAkzYctP/0CPWGRAnfow/mAVCQkdbDAeovKY5EdP2GdB1CFMMfYW" +
		"uaAryRcX7Q52pWAnjqCSvHQ8QE9yS7lcL0ZAKE2TChha9rgiMrDcXvr3AyAmpo2DNXFbmD" +
		"tSdiBH2AqQHfxaLl5rtHXnmwyy+1agD6N9px7NE9KYTE9eSMTAJYnNZmTrLlaw7jumuPAR" +
		"ABST/bDEe8WF2182MMxiJ0OaMiRouMD1HWHnXxXbfYTAbqCoJFxjevrT9nD9lHaKwDCrsO" +
		"D44qXNcP8wh3JfAZANOQocrdkbpFXo9m7oyYBNMQ21zzv1DlHeW7HfDEH0JfwkVZNj2prC" +
		"AOXAgrVL/RZACO+rkV+Z6DTASuN4mpqGnNeofdXnjwfeN0mAVlTjTmvGUjmTIBdJlsbI2q" +
		"m7UmKFOSPPkwB2A1nsVvaWxzRdkE7+6sBOgJxxDWyknx+XhiYCyiv8SHbtR2A5ZcUTwG8L" +
		"pPZPHUPtqjj5EzA+9VzZPWV1SYccv/7RWxy4Fvfmqr+BQAA///VaqA18lDV5wAAAABJRU5" +
		"ErkJggg=="

	qr, err := ku.Base64QR()

	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.FailNow()
	}

	if expected != qr {
		t.Errorf("unexpected qr\nexpected: %s\n  actual: %s", expected, qr)
	}
}
