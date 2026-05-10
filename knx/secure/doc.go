// Package secure provides support for KNX/IP Secure communication.
//
// This package implements KNX Data Secure as specified in the KNX standard,
// providing encryption and authentication for KNX/IP messages.
//
// The main components are:
//   - DataSecure: Handles encryption/decryption of CEMI frames
//   - SecurityControlField: Manages security parameters for secure frames
//   - Keyring: Loads and manages encryption keys from .knxkeys files
//   - SecurityPrimitives: Low-level cryptographic operations
package secure
