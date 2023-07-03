这是一个监控 合约发布后 去mint的程序
----

1、监控 ethereum 上的 vmpx 代币的发布
2、发起mint操作


### 编译

```
make build
```

### 运行

* 1、监控 vmpx合约发布

```
./build/trade watch -r <rpc url>
```

```
./build/trade watch -r https://rpc.ankr.com/eth_goerli 
```

* 2、 直接交易

--power    对应的合约内的 power , power 越大, 代币越大, 最大 不超过 195 (不包含)
--c        发送几次 交易(tx)
--contract 交互的合约


```
./build/trade mint -r <rpc url> --contract <contract address> -p <private key> --power <power> -c <count>
```

```
./build/trade mint  -r https://rpc.ankr.com/eth_goerli  --contract 0xF2aDEa4ADb55A8D080e233220d2e30daB90c252b -p abc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1 --power 1 -c 2
```