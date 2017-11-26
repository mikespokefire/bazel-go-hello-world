# Bazel and go

Trying out the Bazel build tool with a small monorepo using go.

## Building

```
bazel build //tools:*
```

This will produce 2 binaries: `bazel-bin/tools/hello_world` and `bazel-bin/tools/goodbye_world`.

## Testing

```
bazel test //utils/*
```
