
## offline wallet & offline sign transaction ##

#### support type
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
    AVAX
```

#### Safety
     1. Addresses are generated offline to avoid accidental deletion or theft through node generation
     2. The private key is encrypted multiple times to reduce the risk of being taken off the pants
     3. Multiple encryption services are deployed in isolation to avoid private theft by wallet managers
     4. Bind the machine code to the signature verification service and the encryption service to prevent brute force cracking of the secret key after the node is relocated
     5. The service can be deployed to isolated nodes without access to the public network, further enhancing the security of the wallet
     6. Service request IP whitelist authorization, prohibit illegal access
     7. The API communication of the signature verification request is transmitted using AES symmetric encryption


#### deploy
```
    1. build

    cd wallet-srv
    
    Note: Configure related encryption keys, address database account passwords, and related parameters conf/config.go
     Import conf/README.sql to the wallet database

    chmod +x build.sh
    ./build.sh


    2. copy the scripts in the bin/ directory to the deployment server
    3. create log directory
    mkdir -p logs/sign logs/encrypt logs/wallet

    4. Start
     EncryptServer ./encrypt-srv -log_dir ./logs [The encryption service is mainly used for the secondary encryption of the private key when the wallet address is generated and the secondary decryption of the private key service when the sign-srv signature is verified, not external]
     SignServer ./sign-srv -log_dir ./logs [It is mainly used for offline signature verification by transaction services, and does not provide public services]
     Generate address ./wallet -log_dir ./logs -num 20000 -coin btc|bch|eth|trx|sol [Generate offline address]
        

    5. Signature verification request API [public gateway: http://xx.x.x.x:4040/v1]

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
        1) /btc/sign POST
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
        2) /eth/sign POST
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
        `
        3) /trx/sign POST
        `
        {
            "from": "THWUyyjjWY5oAc4KBK3QZKZgGJtnS41d2r",
            "raw_data_hex": "0a02cf4e2208fcad8a445616ec3040e0efcf87fd2f5a66080112620a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412310a154152b31daab5a836bef19c5f2c6aad8ff76dc76180121541e4069902588f1111e28c0f90ac696700be7cd57918904e7091accc87fd2f"
        }
        
        `
        4) /sol/sign POST
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
        Note: post request data is transmitted using AES encryption, please refer to the test case for details.

    6. New types are constantly being updated. 
         Telegram: https://t.me/ukbababa (@ukbababa)
```
#### Scope of application
    ```
        1) Exchange
        2) Cryptocurrency Payment
    ```

### Donate && cooperation
```
Technical Support：  Telegram: https://t.me/ukbababa (@ukbababa)

USDT-TRC20/ TRX : TSgPwnsdoFN2zEQw5Fh9iLJnb3Cnxk84gZ
```