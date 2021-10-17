package main

import (
	"reflect"
	"strings"
)

const (
	// PropertyNameTag is the field tag used to drive the Pulumi schema name. By using the
	// same tag as Pulumi, we avoid the need to redundantly declare multiple tags. Unfortunately,
	// Pulumi does not permit comma-delimited options in this tag, so we need a separate options one.
	PropertyNameTag = "pulumi"
	// PropertyOptionsTag is the field tag used to control various schema options.
	PropertyOptionsTag = "pschema"
)

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
	var hadTags bool
	var result PropertyOptions

	stag := reflect.StructTag(tag)

	// First see if there is a field name.
	if name, has := stag.Lookup(PropertyNameTag); has {
		hadTags = true
		result.Name = name
	}

	// Next see if there are options and, if so, parse and decode the comma-delimited list.
	if opts, has := stag.Lookup(PropertyOptionsTag); has {
		hadTags = true
		for _, key := range strings.Split(opts, ",") {
			switch key {
			case "optional":
				result.Optional = true
			case "replaces":
				result.Replaces = true
			case "in":
				result.In = true
			case "out":
				result.Out = true
			default:
				if strings.HasPrefix(key, "ref=") {
					result.Ref = key[4:]
				}
			}
		}
	}

	return hadTags, result, nil
}
