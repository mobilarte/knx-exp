# knx-exp

[![Check](https://github.com/mobilarte/knx-exp/actions/workflows/check.yaml/badge.svg?branch=main)](https://github.com/mobilarte/knx-exp/actions/workflows/check.yaml)
[![GoDoc](https://godoc.org/github.com/mobilarte/knx-exp?status.svg)](https://godoc.org/github.com/mobilarte/knx-exp)

This repository is **experimental** and it is only tested with the latest stable version.
The repository was created in 2022, based on the excellent work by Ole Krüger [knx-go](https://github.com/vapourismo/knx-go), which is not maintained anymore.

## Packages

The package structure is the same as in the original repository.

 Package           | Description
-------------------|--------------------------------
 **knx**           | Abstractions to communicate with KNXnet/IP servers
 **knx/knxnet**    | KNXnet/IP protocol services
 **knx/dpt**       | Datapoint types
 **knx/cemi**      | cEMI-encoded frames

Packages are only tested with `Go ^1.26.0` (see Actions).
Because **Type Parameters** are used in `dpt\formats.go`, it will certainly not run
in a version prior to `Go 1.18.0`.
