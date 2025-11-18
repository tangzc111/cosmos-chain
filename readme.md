# cosmos-chain
**cosmos-chain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## 功能总览

`core` 模块提供以下能力，满足代币和账户的闭环需求：

- **用户管理**：注册、更新、删除链上用户档案，强制地址唯一。
- **代币生产**：通过 `MsgMint` 铸造新代币后立即发放给指定注册用户。
- **账户转账**：`MsgTransfer` 核验双方身份后调用 bank keeper 完成余额转移。
- **矿工体系**：注册矿工、更新算力描述，并可使用 `MsgRewardMiner` 自动铸造奖励，Keeper 会累积奖励总额。
- **区块追踪**：`BeginBlocker` 会记录每个区块的高度、时间、提议人和哈希，配合查询可实现区块浏览。
- **API 暴露**：所有消息与查询都通过 gRPC 及 gRPC-Gateway 自动提供 REST 访问。

## 常用交易命令

### 创建用户

- 创建两个账号
  ```
    cosmos-chaind keys add laotang --keyring-backend test
    cosmos-chaind keys add laotangtest --keyring-backend test
  ```

- 得到相应的地址、助记词

  ```
  address: cosmos1gkrhqj3mnaju7efyxwp2wqjdjk0r85ya2zu75q
     name: laotang
   pubkey: '{
              "@type":"/cosmos.crypto.secp256k1.PubKey",
              "key": "AofCE89JhO8L51a8xGXHAs64c2SLuabrLRJDs5JOxBwF"
            }'
     type: local

  - 助记词
  area method rapid north suit much tackle such spin uncle black connect treat express memory language useless huge major oyster clay vast laptop candy

  address: cosmos1a7hjvlv6g9kr5fqvecmhc8nt7fje96g5n7vk0d
     name: laotangtest
   pubkey: '{"
              @type":"/cosmos.crypto.secp256k1.PubKey","key":"AqGoREgvhaF2oNuVFsP56JmPDUdinKdoFcqJ0l03w7To"
            }'
     type: local

  - 助记词
  debris resource inherit safe crater average domain hollow chat because later tuition keep throw leg blade knife monkey palm alert month come sense final
  ```

- 通过系统创建的账号 alice 转一笔钱给新建账号（未通过充值转账，链上状态没有这个账号）
  ```
  cosmos-chaind tx bank send \
    $(cosmos-chaind keys show alice -a --home ~/.cosmos-chain --keyring-backend test) \
    $(cosmos-chaind keys show laotang -a --home ~/.cosmos-chain --keyring-backend test) \
    1000000stake \
    --from alice \
    --gas auto --gas-adjustment 1.2 --fees 2000stake
  ```

- 通过 createUser 将账户注册到链上
  ```
  cosmos-chaind tx core create-user \
    cosmos1gkrhqj3mnaju7efyxwp2wqjdjk0r85ya2zu75q \
    cosmos1gkrhqj3mnaju7efyxwp2wqjdjk0r85ya2zu75q \
    "laotang" "test create user" --from laotang
  ```

### 查询用户

- 查看所有用户

  ```
  cosmos-chaind q core list-user
  ```

- 查询单个用户

  ```
  cosmos-chaind q core get-user cosmos1gkrhqj3mnaju7efyxwp2wqjdjk0r85ya2zu75q
  ```

- 查询用户余额

  ```
  cosmos-chaind query bank balances $(cosmos-chaind keys show laotang -a)
  ```

### 代币操作

- 代币的转账

  ```
  cosmos-chaind tx core transfer cosmos1gkrhqj3mnaju7efyxwp2wqjdjk0r85ya2zu75q 50 stake --from laotangtest
  ```

- 代币的生产(增额)

  ```
  cosmos-chaind tx core mint cosmos1gkrhqj3mnaju7efyxwp2wqjdjk0r85ya2zu75q 1000 stake \
    --from laotangtest \
    --gas auto --gas-adjustment 1.2 --fees 2000stake
  ```

所有命令默认通过 `cosmos-chaind` CLI 发送，`[flags]` 中至少需要 `--from <key-name>`、`--chain-id cosmoschain` 和 `--keyring-backend test` 等签名信息。

| 能力 | CLI 示例 |
| --- | --- |
| 注册用户 | `cosmos-chaind tx core create-user <bech32-address> <bech32-address> "Alice" "first operator" --from alice` |
| 更新用户 | `cosmos-chaind tx core update-user <bech32-address> "" "new alias" "profile" --from alice` |
| 删除用户 | `cosmos-chaind tx core delete-user <bech32-address> --from alice` |
| 铸造代币 | `cosmos-chaind tx core mint <recipient> <amount> <denom> --from alice` |
| 账户转账 | `cosmos-chaind tx core transfer <to> 50 stake --from bob` |
| 注册矿工 | `cosmos-chaind tx core create-miner <addr> <addr> 10 "validator" --from <addr>` |
| 更新矿工 | `cosmos-chaind tx core update-miner <addr> "" 20 "fast node" --from <addr>` |
| 删除矿工 | `cosmos-chaind tx core delete-miner <addr> --from <addr>` |
| 发放矿工奖励 | `cosmos-chaind tx core reward-miner <miner-addr> 5 stake --from alice` |

> 提示：`create-user`/`create-miner` 的第一个参数 `index` 必须与地址一致，CLI 会按照 proto 字段顺序依次读取参数。

## 查询与 REST API

默认端口：

- gRPC：`localhost:9090`
- gRPC-Gateway/REST：`http://localhost:1317`

常见查询命令：

```
# 查看所有用户
cosmos-chaind q core list-user

# 查询单个用户
cosmos-chaind q core get-user <bech32-address>

# 查看矿工列表
cosmos-chaind q core list-miner

# 当前区块信息
cosmos-chaind q core latest-block

# 按高度查看区块
cosmos-chaind q core get-block-record <height>
```

对应的 REST 端点可直接用 `curl` 访问，例如：

```
# 查询指定用户信息
curl http://localhost:1317/cosmoschain/core/v1/user/{bech32-address}
# 查询最新区块信息
curl http://localhost:1317/cosmoschain/core/v1/block/latest
# 查询在区块高度 {height} 时记录的信息
curl "http://localhost:1317/cosmoschain/core/v1/block_record/{height}"
```
