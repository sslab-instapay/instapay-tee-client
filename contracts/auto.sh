~/solidity/build/solc/solc --abi InstaPay.sol -o .
~/solidity/build/solc/solc --bin InstaPay.sol -o .
abigen --abi=InstaPay.abi --bin=InstaPay.bin --pkg=instapay --out=InstaPay.go
