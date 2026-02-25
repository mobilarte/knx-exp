# knx-exp

[![Check](https://github.com/mobilarte/knx-exp/actions/workflows/check.yaml/badge.svg?branch=main)](https://github.com/mobilarte/knx-exp/actions/workflows/check.yaml)
[![GoDoc](https://godoc.org/github.com/mobilarte/knx-exp?status.svg)](https://godoc.org/github.com/mobilarte/knx-exp)
[![Go Report Card](https://goreportcard.com/badge/github.com/mobilarte/knx-exp)](https://goreportcard.com/report/github.com/mobilarte/knx-exp)

This repository is **experimental** and it is only tested against the 
latest stable version of Go.

The repository was created in 2022, based on the excellent work by 
Ole Krüger [knx-go](https://github.com/vapourismo/knx-go), which is not maintained anymore.
It was not created as a fork, because it was meant to be temporary.

## Packages

The package structure is the same as in the original repository.

 Package           | Description
-------------------|--------------------------------
 **knx**           | Abstractions to communicate with KNXnet/IP servers
 **knx/knxnet**    | KNXnet/IP protocol services
 **knx/dpt**       | Datapoint types
 **knx/cemi**      | cEMI-encoded frames

Because **Type Parameters** are used in `knx/dpt/formats.go`, it will not run
with a version prior to Go 1.18.0.
