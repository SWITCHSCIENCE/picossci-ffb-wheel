# picossci-ffb-wheel

PCゲーム向けの反力付きハンドル型コントローラー用ファームウェア

必須のパーツ：
- [Picossci-Canボード](https://ssci.to/9279)
- [ダイレクトドライブモーター](https://ssci.to/9219)
- [24V20A電源ユニット](https://amzn.asia/d/iWJe48B)

オプションパーツ：
- [小型グラフィック液晶ボード](https://ssci.to/2608)

# ファームウェアソースコード

https://github.com/SWITCHSCIENCE/picossci-ffb-wheel

## ビルド

- go v1.23
- tinygo v0.33.0

```
git clone https://github.com/SWITCHSCIENCE/picossci-ffb-wheel
cd picossci-ffb-wheel
make build
```

## 書き込み
ボードのBootSelボタンを押しながらPCに接続
「RPI-RP2」という名称のストレージが見えるようになります。
```
make flash
```
