# Changelog

This log file describes briefly the changes made to code as of 2022-07-01 in chronological order. The changes may be minor, important or cosmetic (just the way I like it).

* Minor changes
 
    * Rewrote individual and group address parsing in `cemi/address.go`, added tests. This change may break code, as fewer addresses are now considered valid!
 
    * Added DescriptionRequest and DescriptionResponse in `knxnet/description.go` and `describe.go`. Parses DIBs that are not in sequence (seems to happen!) or unknown DIBs (like DescriptionTypeManufacturerData). Ignores, but logs, completely unknown DIB entries.
 
    * Rewrote `search.go` to use the same DIB parser as `description.go`.
 
    * Added DiagnosticRequest in `diagnostic.go` (unfinished, because I do not have a KNXnet/IP router that responds to diagnostics - reason: they are 14 years old).

    * Rewrote all DPT tests, added many new DPTs, fixed ranges according to [03_07_02-Datapoint-Types-v02.02.01-AS.pdf](https://www.knx.org/wAssets/docs/downloads/Certification/Interworking-Datapoint-types/03_07_02-Datapoint-Types-v02.02.01-AS.pdf).

    * Rewrote `dpt/types_registry.go`, no need for `setup()`.
 
 
* Major changes [To be added shortly]

    * MulticastLoopback option to routing. Investigate how Go handels multicast, some strange behaviour like multicast is forwarded to all interfaces.

    * TCP connections for tunnelling.

    * Adding SendTimer to avoid flooding a KNXnet/IP server.

    * Reworking RoutingBusy by adding random timer and memory to WaitTime.


# Future

* SecureTunnelling.

* More examples.
