
# The Ethereal Quest

## Story
Welcome to Ethereal. Here, your quest is to forge the Seed Blade, a powerful sword assembled from 10 mystical parts. With each part, your blade becomes stronger, symbolizing your growing cryptographic mastery.

Your challenge: to mint Gems with this blade. 
Mint 10 gems, and you demonstrate your skill and balance. 
Mint 20, and you push the boundaries of your abilities. But beware, each gem brings both power and peril. Will your Seed Blade be a tool of wisdom or lead to your downfall? The path is yours to choose.

## Environment Setup

### Starting an Ethereum Node
Launch an Ethereum node:
```sh
anvil
```

### Deploying the Contract
Navigate to the contract directory and deploy the contract:
```sh
# Enter the contract directory
cd mint-contracts

# Deploy the Mint contract
# 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
# 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

forge create Mint --private-key=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

[⠑] Compiling...
No files changed, compilation skipped
Deployer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Deployed to: 0xaad2c7ca5815a54048651a6A7fA9C6aB23c57360
Transaction hash: 0x8684b904d01c9eb2e2438feec39f3a6e1a8c2e45788be1afbd3fe2a94167cd56

# onchain
# e5e8bb4e61f5afe8bf82afec48adf1a5ddd799f042b1c71d3fe5cf26f7f00c31
# 0x78B8Af83b0DEF7e7f89Cd964f9E3921F9685844b
#
# 0xaad2c7ca5815a54048651a6A7fA9C6aB23c57360


# Return to the previous directory
cd -
```

## Starting the Challenge

### Initial Setup
Change to the client directory:
```sh
cd mint-client
```

## Challenge Steps

### 1. The Initiation of the Ethereal Quest
Begin your journey by registering:
```sh
go run . -contract 0xaad2c7ca5815a54048651a6A7fA9C6aB23c57360 -action register

go run . -contract 0xaad2c7ca5815a54048651a6A7fA9C6aB23c57360 -action register


(*big.Int)(0x1400007e020)(9660373881149821873842640389441473316526837076079981247269879061521518137036)
(*bn254.G1Affine)(0x1400002e200)(E([16579489176552295705242449577783279595253357420721807897481407885455124092810,9287744212686234159127478616021006142688015174700233433508952403714876268325]))
```

go run . -contract 0xaad2c7ca5815a54048651a6A7fA9C6aB23c57360 -action query

([]*big.Int) {
}
(struct { Registered bool; Nonce *big.Int; Commitment mint.PairingG1Point; PublicKey mint.PairingG1Point; Escaped0 bool; Escaped1 bool }) {
 Registered: (bool) true,
 Nonce: (*big.Int)(0x140001d5fc0)(1),
 Commitment: (mint.PairingG1Point) {
  X: (*big.Int)(0x140001d5fe0)(7285655873873301917135752923664251176745661555797735151349054290203129140248),
  Y: (*big.Int)(0x140002c8000)(11471138561143431018588512664712300671586266765465791120619532196692255941463)
 },
 PublicKey: (mint.PairingG1Point) {
  X: (*big.Int)(0x140002c8020)(16579489176552295705242449577783279595253357420721807897481407885455124092810),
  Y: (*big.Int)(0x140002c8040)(9287744212686234159127478616021006142688015174700233433508952403714876268325)
 },
 Escaped0: (bool) false,
 Escaped1: (bool) false
}


### 2. The Forging of Ethereum Gems
Mint Ethereum Gems, each one bringing you closer to your goal:
```sh
go run . -contract 0xaad2c7ca5815a54048651a6A7fA9C6aB23c57360  -action mint

&{0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 2 0x102ff0540 <nil> 1789797558 <nil> <nil> 3000000 context.Background false} {9797219847158177989996093187014441826373233429936066039861110838466042263021 18924977759823587178657915824281038560035262976116212999801159651786147984736} {false [8128389501791719627 11874428499506069321 15242590414069608886 2095521652416477102]}

transactor &{0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 2 0x105294400 <nil> 1789797558 <nil> <nil> 3000000 context.Background false}
proofPoint &{9797219847158177989996093187014441826373233429936066039861110838466042263021 18924977759823587178657915824281038560035262976116212999801159651786147984736}
value {false [8128389501791719627 11874428499506069321 15242590414069608886 2095521652416477102]}

transactor &{0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 2 0x1046c03e0 <nil> 1789797558 <nil> <nil> 3000000 context.Background false}
proofPoint &{9797219847158177989996093187014441826373233429936066039861110838466042263021 18924977759823587178657915824281038560035262976116212999801159651786147984736}
value 13153802600923833277929595706871746245972085086634174882780020810967887597771

go run . -contract 0xaad2c7ca5815a54048651a6A7fA9C6aB23c57360 -action mint
Minting the 1st gem
```

### 3. The Path of Redemption
Should you falter, replay and learn from your missteps:
```sh
go run . -contract 0xaad2c7ca5815a54048651a6A7fA9C6aB23c57360 -action replay
```

### 4. Unveil Your Treasures

Verifying the Escape0 Flag: If you've successfully minted 10 gems, the "escape0" flag should now be set, symbolizing a key achievement in your quest. This flag reflects your adeptness in navigating the challenges and evading the watchful eyes of the Watchdog.

Verifying the Escape1 Flag: For the bold adventurers who pushed their limits to mint 20 gems, the "escape1" flag would be their coveted prize. Achieving this flag marks you as one of the few who have reached the pinnacle of daring and strategy in the quest.
To check your achievements and claim your flags, invoke the following incantation:
```sh
go run . -contract 0xaad2c7ca5815a54048651a6A7fA9C6aB23c57360 -action query

([]*big.Int) (len=9 cap=9) {
 (*big.Int)(0x140000c5660)(615298855253563631201382826604470010682646043121472040112462928602133707211),
 (*big.Int)(0x140000c5680)(19706443080845389820124982736666949309525571927495182176705271257965082659720),
 (*big.Int)(0x140000c56a0)(119341808897078565849392865455188375890158855761293753046064705325596279132),
 (*big.Int)(0x140000c56c0)(18612234252842456253032712253335019755642431843014093373014176666107955599601),
 (*big.Int)(0x140000c56e0)(10715961137797832983211077895164562636920005291203533607140715267367867155907),
 (*big.Int)(0x140000c5700)(2636569199964016391987777855412917789214554920703303279136786376610182999464),
 (*big.Int)(0x140000c5720)(15205524770659243482478843072836194383964461467991065114077437934710870671680),
 (*big.Int)(0x140000c5740)(8414204427821629910313153438452068918373229869480497662417109798825803015875),
 (*big.Int)(0x140000c5760)(13136554349491357515460431597077269999238348029104518306380327834431685264364)
}
(struct { Registered bool; Nonce *big.Int; Commitment mint.PairingG1Point; PublicKey mint.PairingG1Point; Escaped0 bool; Escaped1 bool }) {
 Registered: (bool) true,
 Nonce: (*big.Int)(0x140002e0200)(10),
 Commitment: (mint.PairingG1Point) {
  X: (*big.Int)(0x140002e0220)(7285655873873301917135752923664251176745661555797735151349054290203129140248),
  Y: (*big.Int)(0x140002e0240)(11471138561143431018588512664712300671586266765465791120619532196692255941463)
 },
 PublicKey: (mint.PairingG1Point) {
  X: (*big.Int)(0x140002e0260)(16579489176552295705242449577783279595253357420721807897481407885455124092810),
  Y: (*big.Int)(0x140002e0280)(9287744212686234159127478616021006142688015174700233433508952403714876268325)
 },
 Escaped0: (bool) false,
 Escaped1: (bool) false
}
```

## Command Options
Use these options for specific actions:
```bash
go run .
# Options:
# -action [register, mint, replay, query]
# -contract [contract_address]
# -rpc [RPC_endpoint]
# -account [account_address]
# -privateKey [private_key]
```
