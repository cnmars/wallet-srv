
## 离线钱包和离线验签

#### 已接入主链
```
    BTC
    BCH
    BSV
    DASH
    DOGE
    LTC
    QTUM
    ETH (ETH & ERC20 Token)
    BSC (BSC & BEP20 Token)
    HT (HT & HRC20 Token)
    OKX (OKB & ORC20 Token)
    Matic [Matic & ERC20 Token]
    TRX (TRX & TRC20 Token)
    SOL (SOL & SPL Token)
    ETC (ETC & ERC20 Token)
    FIL
    ADA
    XRP
    LUNA
    NEAR
    DOT
    AVAX （X-Chain）
    FLOW
    FTM
    OP

```

#### 安全可靠
     1. 多重安全防护，降低人为因数引起的资金安全
     2. 避免因脱库、触网、物理备份导致安全事故
     3. 生成地址、离线验签、多重加密、私钥存储全部分离，私钥多重加密可分配多方操作


#### 部署
```
    1. 编译，下载wallet-srv代码，执行 ./build.sh即可生成 wallet/sign-srv/encrypt-srv/wallet-tool等脚本，地址私钥通过加密保存到数据库，因此编译前先准备数据库和修改数据连接 conf/README.sql 和 conf/config.go

    cd wallet-srv

    chmod +x build.sh
    ./build.sh


    2. 复制bin中对应的服务到相应的服务器部署
    3. 创建日志目录，并执行
        mkdir -p logs/sign logs/encrypt logs/wallet
        ./sign-srv

    4. 启动，保存启动顺序
     加密服务： ./encrypt-srv -log_dir ./logs 对生成地址的私钥进行二次加密或解密
     验签服务： ./sign-srv -log_dir ./logs  对交易进行离线验签，提币脚本对验签代码进行广播，验签服务不触外网
     地址生成： ./wallet -log_dir ./logs -num 20000 -coin btc|bch|eth|trx|sol 离线地址生成，保存地址和私钥到数据库
        

    5. 验签API请求方式 [API地址: http://xx.x.x.x:4040/v1]

        0) /hd/sign POST JSON
        `
        {
            "vin": [{
                "address": "mrXUCNgmV439zQjh13N4kSG6DJuqJ9zD4C",
                "txid": "c1076d23d1d92170c01588eace8dfedf5284888402fe79344b796749d5eb19a1",
                "vout": 0,
                "amount": 100000
            }],
            "vout": [{
                "address": "mvaCNagSKWA2cpyXoejJCEAFRdnaPeQ3SD",
                "amount": 10000
            }, {
                "address": "mrXUCNgmV439zQjh13N4kSG6DJuqJ9zD4C",
                "amount": 89742
            }],
            "change": "mrXUCNgmV439zQjh13N4kSG6DJuqJ9zD4C",
            "feekb": 80000,
            "coin":"bch" 
        }

        eg: coin value enum["bch", "bsv", "ltc", "dash", "doge", "qtum"]
        `
        1) /btc/sign POST JSON
        `
        {
            "vin": [{
                "address": "mrXUCNgmV439zQjh13N4kSG6DJuqJ9zD4C",
                "txid": "c1076d23d1d92170c01588eace8dfedf5284888402fe79344b796749d5eb19a1",
                "vout": 0,
                "amount": 100000
            }],
            "vout": [{
                "address": "mvaCNagSKWA2cpyXoejJCEAFRdnaPeQ3SD",
                "amount": 10000
            }, {
                "address": "mrXUCNgmV439zQjh13N4kSG6DJuqJ9zD4C",
                "amount": 89742
            }],
            "change": "mrXUCNgmV439zQjh13N4kSG6DJuqJ9zD4C",
            "feekb": 80000
        }
        `
        2) /eth/sign POST JSON
        `
        {
            "from": "0x7eeE959B97a243233EDd1133c7D706Be97999D49",
            "to": "0x09d247829344c4d1D1A623E01C9387fF978E1cEe",
            "amount": 1000,
            "nonce": 0,
            "gaslimit": 21000,
            "gasprice": 18000000000,
            "contract": "",
            "chainid": 1
        }

        chainid: 1 ETH，56 BSC，61 ETC，128 HECO，137 Matic, 66 OKX，250 FTM 10 Optimism 592 Astar
        `
        3) /trx/sign POST JSON
        `
        {
            "from": "THWUyyjjWY5oAc4KBK3QZKZgGJtnS41d2r",
            "raw_data_hex": "0a02cf4e2208fcad8a445616ec3040e0efcf87fd2f5a66080112620a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412310a154152b31daab5a836bef19c5f2c6aad8ff76dc76180121541e4069902588f1111e28c0f90ac696700be7cd57918904e7091accc87fd2f"
        }
        
        `
        4) /sol/sign POST JSON
        `
        {
            "from": "GwYeRRgRhykYRX9taXWVALXbCJQaUECMG4s6gHS4sxUv",
            "to": "EWmowLYfz49uyNJdT22usS9qQm495qZdKSYsGQjLw6PJ",
            "amount": 10000,
            "last_blockhash": "6MEVqZdfpkunirTp1UHXecYjJmfxuFLwhTQKai7LePKd",
            "token": "",
            "decimal": 9,
        }
        `
        5) /fil/sign POST JSON
        `
        {
            "from": "f1ahx4j4sx4s4xfi7v542xjyhgaur3lcsqtnwvswq",
            "to": "f1a6qvq7evfo5zwspflulju2qfcnzcdo7fsyhmy4q",
            "amount": 0.211212,
            "nonce": 0,
            "gaslimit": 21000,
            "gasfee": 1000,
            "gaspremium": 1000,
        }
        `
        6) /ada/sign POST JSON
        `
        {
            "vin": [
                {
                    "txid": "3462e1c16086ee4f066d74d4e45c498657033b4f1e4688f9d9d6e995d9604c50",
                    "index": 1,
                    "amount": 0.2,
                    "address": "addr1vxr77yvf4ek9q68c9mrflu6l9kuxm4umhyman3yxzx74zss63vkx4"
                }
            ],
            "vout": [
                {
                    "address": "addr1vydusqepys5wv0a030parxcgulj0pahcxx37e87z33yurhs3y02ke",
                    "amount": 0.1
                }
            ],
            "change": "addr1vx8wcckrrv9e27dtcwp33zu796p6s02qy356ykdauhqpdyq6lhxec",
            "fee_param": {
                "txFeePerByte": 121,
                "txFeeFixed": 1212,
                "maxTxSize": 1212,
                "protocolVersion": {
                    "major": 0,
                    "minor": 1
                },
                "minUTxOValue": 1
            }
        }
        `
        7) /xrp/sign POST JSON
        `
        {
            "from": "r4g1C5rYUvoGJZRHQPQKGwS84ZdPsLJcsg",
            "to": "rnvJTfwpLniF4KBf9XzcdXh53WXSTp6ick",
            "amount": "0.211212",
            "fee": 100,
            "assetcode": "native", //xrp
            "assetissuer": "",
            "sequence": 1201, // from sequence index
        }
        `
        8) /luna/sign POST JSON
        `
        {
            "from": "terra1r026gv4gzp8774p3nkn7shweyn42gttvdydfm6",
            "to": "terra1yv0tf06m4plwg5r622pgquwmerae08hjrtqzyz",
            "amount": "0.211212",
            "fee": "0.00001",
            "gaslimit": 100,
            "coin": "luna",
            "sequence": 1201, // from sequence index
        }
        `
        9) /near/sign POST JSON
        `
        {
            "from": "9684664164044ed7955e2bf4062c8e90b26c6f013213e9e72f59c0e633d4b921",
            "to": "75dff9ad7869b27d1ddcd4f8da8f65f8f29ddf6bbaaeb929cbaee1971ab39148",
            "amount": 10000,
            "last_blockhash": "6MEVqZdfpkunirTp1UHXecYjJmfxuFLwhTQKai7LePKd",
            "nonce": 0
        }
        `
        10) /dot/sign POST JSON
        `
        {
            "from": "9684664164044ed7955e2bf4062c8e90b26c6f013213e9e72f59c0e633d4b921",
            "to": "75dff9ad7869b27d1ddcd4f8da8f65f8f29ddf6bbaaeb929cbaee1971ab39148",
            "amount": 10000,
            "fee": 0,
            "last_blockhash": "6MEVqZdfpkunirTp1UHXecYjJmfxuFLwhTQKai7LePKd",
            "nonce": 0,
            "spec_ver": 9190,
            "tran_ver": 14,
        }
        `
        11) /avax/sign POST JSON
        `
        {
            "vin":[
				   {"txid":"xxxxxxx", "vout":0, "amount": 11, "address": "xxxx", "asset_id": "xxxxx"},
				   {"txid":"xxxxxxx", "vout":1, "amount": 111, "address": "xxxx", "asset_id": "xxxxx"},
				   {"txid":"xxxxxxx", "vout":2, "amount": 11, "address": "xxxx", "asset_id": "xxxxx"},
			   ],
		   	"vout":[
				   {"address": "1xxx99a9axxx", "amount": 1212, "asset_id": "xxxxx"},
				   {"address": "1xxx99a9axxx", "amount": 11, "asset_id": "xxxxx"},
				   {"address": "1xxx99a9axxx", "amount": 11, "asset_id": "xxxxx"},
			   ],
			"block_chainid": "xxxxxxxxxxxxx",//
			"chain_id": 1,
			"memo": ""
        }
        `
        12) /flow/sign POST JSON
        `
        {
            "from": "",
			"to": "",
			"player": "",
			"amount": "102.99911",
			"last_blockhash": "" //last block hash
			"token": "1654653399040a61",
			"seq_num": 121
        }
        `

        注：请求API数据需要经过AES加密传输，详情请看test用例

    6. 如需增加其他币种，友情技术支持：  Telegram: https://t.me/ukbababa (@ukbababa)
```
#### 应用场景
    ```
        1) 交易所
        2) 游戏内支付
        3) 需要支持加密支付等业务的场景
    ```

### 求合作或赞赏
```
技术支持： Telegram: https://t.me/ukbababa (@ukbababa)

USDT-TRC20/ TRX : TSgPwnsdoFN2zEQw5Fh9iLJnb3Cnxk84gZ
USDT-ERC20/BSC/MATIC/HECO : 0xA6cAa332420B2d11D7396E95813FE3dc2f5D573B
```