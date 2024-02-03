# Readme

[-] input your choice: 1
[+] deployer account: 0x223499Cd597f6C3D50946f3A7322a80b33536f83
[+] token: v4.local.w1cOBRuu4jsIhAzY0sRejin-zBNBoub32Cx0p7SHmuQUpz-q4rqzXVh3DjeI6YZOdLi6e5REy19g2edm6jFd8jyhUODAeV1gb7hqcCUNkxPC_6cWOltjH-zEKV6RHsBlvDUPDJbP5nIRJ4ofjcuWlwzt2buGDcISWKvXFDV4CZ6trw.Um91bmRhYm91dA
[+] please transfer more than 0.001 test ether to the deployer account for next step


[+] contract address: 0x03b90F3bC8Fdaf22C291717c001C03F9f8bE4c54
[+] transaction hash: 0xa72d05547a45ccbeb70b836fab94477b3187eea7bfc66cd084439a3004057733


```sh
cd Circom/

circom Roundabout.circom --r1cs --wasm --sym

node Roundabout_js/generate_witness.js Roundabout_js/Roundabout.wasm input.json witness.wtns

snarkjs powersoftau new bn128 12 pot12_0000.ptau -v
 

snarkjs zkey export verificationkey Roundabout_groth16.zkey verification_key.json


snarkjs groth16 prove Roundabout_groth16.zkey witness.wtns proof.json public.json

snarkjs groth16 verify verification_key.json public.json proof.json
[INFO]  snarkJS: OK!

snarkjs zkey export solidityverifier Roundabout_groth16.zkey verifier.sol

snarkjs generatecall

["0x2268a89673fd854f27a45fc264c7fac61a83fcb69cf0f967846dcdcdababa676", "0x21adf2e1155b30df7e0fc256f61afdf65a0c5e27e93dd0598156441a7c4e2b99"],

[["0x0be4f8686b61c452296b30bf1bc4717bc6fff5f10d79d0373299d9956682560f", "0x0fa3fbf3ec14920e7cd363e30ca39fb3c380b57d5fa9cfdd051d743b7c25a91f"],["0x2ff88b545aefc61620390ffdb634e1d4a228ec2104ea90a851eeba37eabe14aa", "0x260cdb579e66f33bee3cfb7cb4b983cf4a1f507b689c6728b7e24debdcb5897b"]],

["0x2ed3e8d8225fc4f27764f3b4886ad67728af69ca0200926814194cd3ace1c176", "0x0b7b1f48473c6d15073f375aeab481903d477f8f0aeb4367b1c2ea58eb9725c3"],

["0x00000000000000000000000000000000000000000000000000000000025d2f72"]

```
