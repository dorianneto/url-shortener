package document

import "testing"

type dataProvider struct {
	key  string
	want string
}

func TestGetByKey(t *testing.T) {
	r := &ReadOutput{Data: map[string]interface{}{
		"Url": "http://www.example.com",
	}}

	cases := &[]dataProvider{
		{"url", "http://www.example.com"},
		{"Url", "http://www.example.com"},
	}

	for _, dp := range *cases {
		got := r.GetByKey(dp.key)

		if got != dp.want {
			t.Errorf("Expected '%s', but got '%s'", dp.want, got)
		}
	}
}
