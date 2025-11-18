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

所有命令默认通过 `cosmos-chaind` CLI 发送，`[flags]` 中至少需要 `--from <key-name>`、`--chain-id cosmos-chain` 和 `--keyring-backend test` 等签名信息。

| 能力 | CLI 示例 |
| --- | --- |
| 注册用户 | `cosmos-chaind tx core create-user <bech32-address> <bech32-address> "Alice" "first operator" --from alice` |
| 更新用户 | `cosmos-chaind tx core update-user <bech32-address> "" "new alias" "profile" --from alice` |
| 删除用户 | `cosmos-chaind tx core delete-user <bech32-address> --from alice` |
| 铸造代币 | `cosmos-chaind tx core mint <recipient> 1000 stake --from alice` |
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

```bash
curl http://localhost:1317/cosmoschain/core/v1/user/{bech32-address}
curl http://localhost:1317/cosmoschain/core/v1/block/latest
curl "http://localhost:1317/cosmoschain/core/v1/block_record/{height}"
```

要通过 REST 发送交易，可向 `/cosmos/tx/v1beta1/txs` 提交 `cosmoschain.core.v1` 包下的 `MsgMint`、`MsgTransfer`、`MsgRewardMiner` 等消息体。

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers a frontend scaffolding feature (based on Vue) to help you quickly build a web frontend for your blockchain:

Use: `ignite scaffold vue`
This command can be run within your scaffolded blockchain project.


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/cosmos-chain@latest! | sudo bash
```
`username/cosmos-chain` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/ignite/installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.com/invite/ignitecli)
