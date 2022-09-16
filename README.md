# KNX-EXP

**Do not rely on this repository! It may suddenly disappear.**

This repository is **experimental**.

It is based on the excellent work by [Ole Krüger](https://github.com/vapourismo/knx-go). Check his `knx-go` package if you need to access a KNXnet/IP device from Go.

## Packages

The package structure is the same as in the original repository.

 Package           | Description
-------------------|--------------------------------
 **knx**           | Abstractions to communicate with KNXnet/IP servers
 **knx/knxnet**    | KNXnet/IP protocol services
 **knx/dpt**       | Datapoint types
 **knx/cemi**      | cEMI-encoded frames

Packages are only tested with `Go ^1.19.1` (see Actions), because **Type Parameters** are used in `dpt\formats.go`.

