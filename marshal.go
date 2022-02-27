package strval

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/go-multierror"
)

func Marshal(in interface{}) ([]byte, error) {
	acc := make([]string, 0, 10)
	if err := processValue("", in, &acc); err != nil {
		return nil, err
	}

	sort.Strings(acc)
	return []byte(strings.Join(acc, "\n")), nil
}

func processValue(prefix string, v interface{}, acc *[]string) error {
	switch vv := v.(type) {
	case string, int:
		*acc = append(*acc, fmt.Sprintf("%s: %v", prefix, vv))
	case map[string]interface{}:
		if err := processMap(prefix, vv, acc); err != nil {
			return err
		}
	case []interface{}:
		if err := processSlice(prefix, vv, acc); err != nil {
			return err
		}
	default:
		return fmt.Errorf("don't know how to handle %T", v)
	}

	return nil
}

func processMap(prefix string, vv map[string]interface{}, acc *[]string) error {
	var errs error
	for k, v := range vv {
		if err := processValue(augmentPrefix(prefix, k), v, acc); err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	if errs != nil {
		return errs
	}

	return nil
}

func processSlice(prefix string, vv []interface{}, acc *[]string) error {
	var errs error
	for _, v := range vv {
		if err := processValue(prefix, v, acc); err != nil {
			errs = multierror.Append(errs, err)
		}
	}

	if errs != nil {
		return errs
	}

	return nil
}

func augmentPrefix(base, add string) string {
	if base == "" {
		return add
	}

	return base + "." + add
}
