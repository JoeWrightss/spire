<<<<<<< HEAD
## Contributing to Spire

This document is a guide to how to get started contributing to Spire.

### Writing guides or documentation

Often the biggest issue facing open-source projects is a lack of good
documentation, and Spire is no exception here. If you have ideas for
specific pieces of documentation which are absent, feel free to open a
specific issue for that.

We also gladly accept patches for documentation. Anything from fixing
a typo to writing a full tutorial is a great way to help the
project.

### Reporting bugs, issues, or unexpected behavior

If you encounter anything that is broken, confusing, or could be
better, you should
[open an issue](https://github.com/non/spire/issues). You don't have
to know *why* the error is occuring, or even that an error happens at
all.

If you are trying to do something with Spire, and are having a hard
time, it could be any of the following issues:

 * an actual bug or error
 * an omission or problem with the API
 * a confusing edge case
 * a documentation problem

Feel free to open a bug before you're sure which of these is
happening.  You can also ask questions on the
[mailing list](http://groups.google.com/group/spire-math/) to get
other people's opinions.

### Creating or improving tests

Spire uses [ScalaTest](www.scalatest.org) and
[ScalaCheck](http://scalacheck.org/) to test our code. The tests
fulfill a number of important functions:

 * ensure our algorithms return correct results
 * check the visibility of our type class instances
 * confirm that the API works as we expect
 * test edge cases which might otherwise be missed

If you find a bug you are also encouraged to submit a test case (the
code you tried that failed). Adding these failing cases to Spire's
tests provides a good way to ensure the bug is really fixed, and is
also a good opportunity to start contributing.

ALso, when you notice places that lack tests (or where the tests are
sparse, incomplete, or just ugly) feel free to submit a pull request
with improvements!

### Submitting patches or code

Spire is on Github to make it easy to fork the code and change it.
There are very few requirements but here are some suggestions for what
makes a good pull request.

If you're writing a small amount of code to fix a bug, feel free to
just open a pull request immediately. You can even attach some code
snippets to the issue if that's easier.

For adding new code to Spire, it's often smart to create a topic
branch that can be used to collaborate on the design. Features that
require a lot of code, or which represent a big change to Spire, tend
not to get merged to master as quickly. For this kind of work, you
should submit a pull request from your branch, but we will probably
leave the PR open for awhile while commenting on it.

You can always email the list, or visit the `#spire-math` IRC channel
to get a second opinion on your idea or design.

### Ask questions and make suggestions

Spire strives to be an excellent part of the Scala ecosystem. We
welcome your questions about how Spire works now, and your ideas for
how to make it even better!
=======
# Contributor guidelines and Governance

Please see
[CONTRIBUTING](https://github.com/spiffe/spiffe/blob/master/CONTRIBUTING.md)
and
[GOVERNANCE](https://github.com/spiffe/spiffe/blob/master/GOVERNANCE.md)
from the SPIFFE project.

# Prerequisites

For basic development you will need:

* **Go 1.11** or higher (https://golang.org/dl/)

For development that requires changes to the gRPC interfaces you will need:

* The protobuf compiler (https://github.com/google/protobuf)
* The protobuf documentation generator (https://github.com/pseudomuto/protoc-gen-doc)
* protoc-gen-go and protoc-gen-spireplugin (`make utils`)


#  Building

Since go modules are used, this repository can live in any folder on your local disk (it is not required to be in GOPATH).

A Makefile is provided for common actions.

* `make all` - installs 3rd-party dependencies, build all binaries, and run all tests
* `make` - builds all binaries
* `make cmd/spire-agent` - builds one binary
* `make test` - runs all tests

**Other Makefile targets**

* `vendor` - Make vendored copy of dependencies using go mod
* `race-test` - run `go test -race`
* `clean` - cleans `vendor` directory
* `distclean` - removes caches in addition to `make clean`
* `utils` - installs gRPC related development utilities
* `help` - shows makefile targets and description

## Development in Docker

You can either build Spire on your host or in a Ubuntu docker container. In both cases you will use
the same Makefile commands.

To run in a docker container set the environment variable `SPIRE_DEV_HOST` to `docker` like so:

```
$ export SPIRE_DEV_HOST=docker
```

To set up the build container and run bash within it:

```
$ make container
$ make cmd
```

Because the docker container shares `$GOPATH/pkg/mod` you will not have to re-install the go dependencies every time you run the container.

## CI

The script `build.sh` manages the CI build process, implementing several unique steps and sanity
checks. It is also used to bootstrap the Go environment in the Docker container.

* `setup` - download and install necessary build tools into the directory `.build-<os>-<arch>`
* `protobuf` - regenerate the gRPC pb.go and README.md files
* `protobuf_verify` - check that the checked-in generated code is up-to-date
* `distclean` - calls `make distclean` and removes the directory `.build-<os>-<arch>`
* `artifact` - generate a `.tgz` containing all of the SPIFFE binaries
* `test` - when called from within a Travis-CI build, runs coverage tests in addition to the
  regular tests
* `utils` - calls `make utils` and installs additional packages for the CI build
* `eval $(build.sh env)` - configure GOPATH, GOROOT and PATH to use the private build tool directory


# Conventions

In addition to the conventions covered in the SPIFFE project's
[CONTRIBUTING](https://github.com/spiffe/spiffe/blob/master/CONTRIBUTING.md), the following
conventions apply to the SPIRE repository:

## Directory layout

`/cmd/{spire-server,spire-agent}/`

The CLI implementations of the agent and server commands

`/pkg/{agent,server}/`

The main logic of the agent and server processes and their support packages

`/pkg/common/`

Common functionality for agent, server, and plugins

`/plugin/{agent,server}/<name>/`

The implementation of each plugin and their support packages

`/proto/{agent,server,api,common}/<name>/`

gRPC .proto files, their generated .pb.go, and README_pb.md.

The protobuf package names should be `spire.{server,agent,api,common}.<name>` and the go package name
should be specified with `option go_package = "<name>";`

## Interfaces

Packages should be exported through interfaces. Interaction with packages must be done through these
interfaces

Interfaces should be defined in their own file, named (in lowercase) after the name of the
interface. eg. `foodata.go` implements `type FooData interface{}`

## Mocks

Unit tests should avoid mock tests as much as possible. When necessary we should inject mocked
object generated through mockgen

# Git hooks

We have checked in a pre-commit hook which enforces `go fmt` styling. Please install it
before sending a pull request. From the project root:

```
ln -s ../../.githooks/pre-commit .git/hooks/pre-commit
```
>>>>>>> upstream/master
