package telebot_test

import (
	"github.com/mohsensamiei/gopher/v3/pkg/telebot"
	"reflect"
	"testing"
)

func TestArguments_ParseAndString(t *testing.T) {
	tests := []struct {
		input    string
		expected telebot.Arguments
	}{
		{
			input: "connect_123_abc-456-def",
			expected: telebot.Arguments{
				"connect": {"123", "abc-456-def"},
			},
		},
		{
			input: "key1_value1_value2__key2_val1_val2_val3__key3_val1",
			expected: telebot.Arguments{
				"key1": {"value1", "value2"},
				"key2": {"val1", "val2", "val3"},
				"key3": {"val1"},
			},
		},
		{
			input:    "",
			expected: telebot.NewArguments(),
		},
		{
			input: "singlekey_singlevalue",
			expected: telebot.Arguments{
				"singlekey": {"singlevalue"},
			},
		},
	}

	for _, tt := range tests {
		args := telebot.NewArguments()
		err := args.Parse(tt.input)
		if err != nil {
			t.Fatalf("Parse failed: %v", err)
		}
		if !reflect.DeepEqual(args, tt.expected) {
			t.Errorf("Parse mismatch.\nGot: %+v\nExpected: %+v", args, tt.expected)
		}

		str := args.String()
		args2 := telebot.NewArguments()
		err = args2.Parse(str)
		if err != nil {
			t.Fatalf("Parse after String failed: %v", err)
		}
		if !reflect.DeepEqual(args2, tt.expected) {
			t.Errorf("Parse after String mismatch.\nGot: %+v\nExpected: %+v", args2, tt.expected)
		}
	}
}

func TestArguments_ExistsAndGet(t *testing.T) {
	args := telebot.NewArguments()
	args.Set("key1", "v1", "v2")
	args.Set("key2", "val1")

	if !args.Exists("key1") {
		t.Error("Exists should return true for key1")
	}
	if args.Exists("key3") {
		t.Error("Exists should return false for non-existent key3")
	}

	vals, ok := args.Get("key1")
	if !ok || !reflect.DeepEqual(vals, []string{"v1", "v2"}) {
		t.Errorf("Get failed for key1. Got: %v", vals)
	}

	vals, ok = args.Get("key3")
	if ok || vals != nil {
		t.Errorf("Get should return false for non-existent key3. Got: %v", vals)
	}
}

func TestArguments_SetAndAdd(t *testing.T) {
	args := telebot.NewArguments()

	args.Set("key1", "v1")
	if vals, _ := args.Get("key1"); !reflect.DeepEqual(vals, []string{"v1"}) {
		t.Errorf("Set failed, got %v", vals)
	}

	args.Add("key1", "v2", "v3")
	if vals, _ := args.Get("key1"); !reflect.DeepEqual(vals, []string{"v1", "v2", "v3"}) {
		t.Errorf("Add failed, got %v", vals)
	}

	args.Add("key2", "val1")
	if vals, _ := args.Get("key2"); !reflect.DeepEqual(vals, []string{"val1"}) {
		t.Errorf("Add failed for new key, got %v", vals)
	}
}

func TestArguments_ParseErrors(t *testing.T) {
	args := telebot.NewArguments()

	// missing underscore
	if err := args.Parse("keyonly"); err == nil {
		t.Error("Parse should fail on invalid format")
	}

	// empty value is ok
	if err := args.Parse("key_"); err != nil {
		t.Errorf("Parse should succeed for empty value, got %v", err)
	}
	if vals, ok := args.Get("key"); !ok || !reflect.DeepEqual(vals, []string{""}) {
		t.Errorf("Empty value parsing failed, got %v", vals)
	}
}

func TestArguments_StringDeterministic(t *testing.T) {
	args := telebot.NewArguments()
	args.Set("a", "1", "2")
	args.Set("b", "x")

	str := args.String()
	args2 := telebot.NewArguments()
	if err := args2.Parse(str); err != nil {
		t.Errorf("Parse after String failed: %v", err)
	}

	if !reflect.DeepEqual(args, args2) {
		t.Errorf("Parse/String roundtrip failed.\nGot: %+v\nExpected: %+v", args2, args)
	}
}
