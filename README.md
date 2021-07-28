# Run local

```sh
bazel build //...
```

```sh
bazel run //server:server
```
# With Docker

Add this to an .bazelrc file in the root of the project:

```
run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64
```

```sh
bazel build //...
```

```sh
bazel run //server:mango_container
```

```sh
docker run --rm -it -p8000:8000 bazel/server:mango_container
```

Test API

```sh
curl -v http://localhost:8000/releases/viz
```

Generate new pages, execute this command from the root folder:

```sh
zx dlpages.mjs
```
