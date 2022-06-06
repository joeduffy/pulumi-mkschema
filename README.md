# MkSchema

A tool to generate Pulumi Package schemas from Go type definitions.

This tool translates annotated Go files into Pulumi component schema metadata, which is in JSON Schema. This
allows a component author to begin by writing the same Go structures, which is often more natural. It also has
the benefit of using these same structures in the implementation of the component itself.

To use the tool, first install it:

```bash
go install github.com/joeduffy/pulumi-mkschema@latest
```

It accepts two arguments: the Pulumi Package name and the Go package path to generate types from:

```bash
pulumi-mkschema [PULUMI-PKG-NAME] [GO-SOURCE-PKG]
```

## How it works

MkSchema will parse and semantically analyze the Go package's metadata. It looks for publicly exported
types that are annotated as either resources or complex types.

Resource types are any structs that embed the `pulumi.ResourceState` resource:

```go
type MyComponent struct {
    pulumi.ResourceState
    ...
}
```

Complex types are any structs that have ``pulumi:"..."`` annotated fields within them:

```go
type MyStateStruct struct {
    Name string `pulumi:"name"`
    ...
}
```

## Pulumi tag options

The ``pulumi:"..."`` tags can be used to control schema generation behavior. Similar to familiar Go
tags like ``json:"..."``, the first element is the Pulumi name, followed by optional comma-delimited options.

These options include:

* `optional`: mark that the property is optional (default is required)
* `replaces`: indicate that a property, if changed, implies replacement behavior
* `in`: indicate that a property is input-only
* `out`: indicate that a property is output-only
* `ref`: reference an externally defined type, rather than intra-package (which is the default)
