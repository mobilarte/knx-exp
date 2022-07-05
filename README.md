# KNX-EXP

This repository is **experimental**. It is based on the excellent work by [Ole Krüger](https://github.com/vapourismo/knx-go).

**Do not rely** on this repository unless you know what you are doing.

## Notable Additions

 Key               | Description
-------------------|--------------------------------------------------------------------
**describe.go**    | Get description from KNXnet/IP server via unicast-UDP. This implementation can also cope with out-of-order DIB and unknown DIBs.
**timer**


## Packages

The package structure is the same as in the original repository.
 Package           | Description
-------------------|--------------------------------------------------------------------
 **knx**           | Abstractions to communicate with KNXnet/IP servers
 **knx/knxnet**    | KNXnet/IP protocol services
 **knx/dpt**       | Datapoint types
 **knx/cemi**      | CEMI-encoded frames

