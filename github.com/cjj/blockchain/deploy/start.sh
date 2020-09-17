#!/bin/bash

# 本脚本从头构建一个区块链网络
# 请确保cryptogen 和 configtxgen 这两个可执行文件已经被正确安装
# 创建一个通道 jiajiechannel

# 一、环境清理
echo "1.clean environment"
mkdir -p config
mkdir -p crypto-config
rm -fr config/*
rm -fr crypto-config/*
echo "finish clean"

# 二、生成证书和起始区块信息
echo "2.Generate certificate and starting block information"
cryptogen generate --config=./crypto-config.yaml
configtxgen -profile OneOrgOrdererGenesis -outputBlock ./config/genesis.block

# 三、启动区块链网络
echo "3. Blockchain ： start"
docker-compose up -d        # 按照docker-compose.yaml的配置启动区块链网络并在后台运行
echo "Waiting for the startup of the node to complete, wait 10 seconds"
sleep 10                    # 启动整个区块链网络需要一点时间，所以此处等待10s，让区块链网络完全启动

# 四、生成通道(这个动作会创建一个创世交易，也是该通道的创世交易)
# channel === 通道
echo "4. Generate the TX file of the channel (this action will create a creation transaction, which is also the creation transaction of the channel)"
configtxgen -profile TwoOrgChannel -outputCreateChannelTx ./config/jiajiechannel.tx -channelID jiajiechannel

# 五、在区块链上按照刚刚生成的TX文件去创建通道
# 该操作和上面操作不一样的是，这个操作会写入区块链
echo "5. Create a channel on the blockchain according to the TX file just generated"
docker exec cli peer channel create -o orderer.blockchain.com:7050 -c jiajiechannel -f /etc/hyperledger/config/jiajiechannel.tx

# 六、让节点去加入到通道
echo "6. let the node join the channel"
docker exec cli peer channel join -b jiajiechannel.block

# 七、链码安装
# -n 是链码的名字，可以自己随便设置
# -v 就是版本号，就是composer的bna版本
# -p 是目录，目录是基于cli这个docker里面的$GOPATH相对的
# 此处安装的是示例链码，后续课程会自己编写
echo "7. chaincode installation"
docker exec cli peer chaincode install -n blockchain -v 1.0.0 -l golang -p github.com/cjj/blockchain/chaincode/blockchain

#八、实例化链码
#-n 对应前文安装链码的名字 其实就是composer network start bna名字
#-v 为版本号，相当于composer network start bna名字@版本号
#-C 是通道，在fabric的世界，一个通道就是一条不同的链，composer并没有很多提现这点，composer提现channel也就在于多组织时候的数据隔离和沟通使用
#-c 为传参，传入init参数
echo "8. instantiate the chaincode"
docker exec cli peer chaincode instantiate -o orderer.blockchain.com:7050 -C jiajiechannel -n blockchain -l golang -v 1.0.0 -c '{"Args":["init"]}'

# 进行链码交互，验证链码是否正确安装及区块链网络能否正常工作
# docker exec cli peer chaincode invoke -C jiajiechannel -n blockchain -c '{"Args":[""]}'