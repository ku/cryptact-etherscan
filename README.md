cryptac
=====
CryptactにEtherのtransaction feeをPAYとして追加するためのCSVを生成します。純粋な送金はSENDFEEとして生成します。

# How to use

## address
CSVで出したいEthereumのアドレスを指定します。

## source
cryptactの`source`です。なんでもいいです。


```
go run main.go adaptor.go  --key [your etherscan.io api key] ether --source MEW --address [0x.......]
```

