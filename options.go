package main

import (
	"reflect"
	"strings"
)

// PropertyOptionsTag is the field tag the IDL compiler uses to find property options.
const PropertyOptionsTag = "pulumi"

// PropertyOptions represents a parsed field tag, controlling how properties are treated.
type PropertyOptions struct {
	Name     string // the property name to emit into the package.
	Optional bool   // true if this is an optional property.
	Replaces bool   // true if changing this property triggers a replacement of this resource.
	In       bool   // true if this is part of the resource's input, but not its output, properties.
	Out      bool   // true if the property is part of the resource's output, rather than input, properties.
	Ref      string // required if we're referencing another package's type.
}

// ParsePropertyOptions parses a tag into a structured set of options.
func ParsePropertyOptions(tag string) (bool, PropertyOptions, error) {
	if ptag, has := reflect.StructTag(tag).Lookup(PropertyOptionsTag); has {
		// The first element is the name; all others are optional flags.  All are delimited by commas.
		opts := PropertyOptions{}
		if keys := strings.Split(ptag, ","); len(keys) > 0 {
			opts.Name = keys[0]
			for _, key := range keys[1:] {
				switch key {
				case "optional":
					opts.Optional = true
				case "replaces":
					opts.Replaces = true
				case "in":
					opts.In = true
				case "out":
					opts.Out = true
				default:
					if strings.HasPrefix(key, "ref=") {
						opts.Ref = key[4:]
					}
					// Ignore unrecognized tags:
					// return true, opts, errors.Errorf("unrecognized tag `pulumi:\"%v\"`", key)
				}
			}
		}
		return true, opts, nil
	}
	return false, PropertyOptions{}, nil
}
