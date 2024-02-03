# CheckIn

```sh
cd Circom/

circom CheckIn.circom --r1cs --wasm --sym

node CheckIn_js/generate_witness.js CheckIn_js/CheckIn.wasm input.json witness.wtns

snarkjs powersoftau new bn128 12 pot12_0000.ptau -v
 

snarkjs zkey export verificationkey CheckIn_groth16.zkey verification_key.json


snarkjs groth16 prove CheckIn_groth16.zkey witness.wtns proof.json public.json

snarkjs groth16 verify verification_key.json public.json proof.json
[INFO]  snarkJS: OK!

snarkjs zkey export solidityverifier CheckIn_groth16.zkey verifier.sol

snarkjs generatecall
["0x00e7bc1e1a59b576135531739f2e8b628b6018cca84de820ede6bbeee8911e77", "0x02451bc54c644a1d8e7b7aac65a0d7477506b799e9bfac25bbb64088db9286e6"],

[["0x228f1599081b135481595f53b873ba318ec473edbd6397f300cc2cb998752d61", "0x1d2214bc450c1a9fd2f6f1ea3c002e72e0ef44be88cf321969179de2ff21b84e"],["0x0c246c915582c47181a5ae4870c276492f7748aca4f42e9612f7fa62b8e76cdb", "0x0c0259bcc68e22b339657001e302f353ba25dd140b80593b647f597d3c2baa71"]],

["0x27026f666190ad173f3a4647df58429c008a75d8d3c30d962cf81c9b36babd58", "0x0adaf9f3854992002609f5dd0daecfebd285d942ff04aacaef09fb23ee5d852e"],

["0x0000000000000000000000000000000000000000000000000000000000000021"]
```
token:

v4.local.FylI_AHFP4m0BKF_-lcZs1MEsvL-UEGgJO1ZZ_0n3IWTIW2neXA3dR18CMfTVEXvBUW0xVpw6RX1BpGewZFAU8OqcAH4s-pKTuVOBJsV86yQeJxYeZzt0Cb3d1K2SY1u8aC3ci-1Srqp2l-OPNwbQ_E_U6ZMgGb6ghG8PJ6Y1JpGZw.Q2hlY2tpbg

0x6749C7eB82924d27183999Ed5839873e79378659


flag{Fbn9qkfCYSxm}_Checkin