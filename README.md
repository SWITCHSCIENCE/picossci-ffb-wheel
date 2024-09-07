# picossci-ffb-wheel

PCゲーム向けの反力付きハンドル型コントローラー用ファームウェア

必須のパーツ:
- [Picossci-Canボード](https://ssci.to/9279)
- [ダイレクトドライブモーター（ゴムタイヤは使いません）](https://ssci.to/9219)
- [24V20A電源ユニット](https://amzn.asia/d/iWJe48B)
- [ハンドル固定プレート（3Dプリント）](https://github.com/SWITCHSCIENCE/Pico-DIYSteeringWithFFB/blob/main/models/steering/steering.stl)
- ハンドルとハンドル固定プレートを留めるワッシャとナット（M4 x 8セット）
- [モーター固定アングル（3Dプリント、耐熱60度程度が必要でABS推奨）](https://github.com/SWITCHSCIENCE/Pico-DIYSteeringWithFFB/blob/main/models/steering/steering-holder.stl)
- 実車用ハンドル(例：https://amzn.asia/d/6pft5gP)
- 固定用の板（450x250x15mmの木材など）
- 板と電源を固定するネジ（M3-20mm x 4本）
- 板とモーター固定アングルを固定するボルト＆ナット（M5-30mm x 4本）
- 板を机に固定するためのクランプ（50mm x 2本）

ハンドル固定プレートとモーターを固定するネジはモーターに付属のものを利用。
ハンドルとハンドル固定プレートを固定するボルトはハンドルに付属のものを利用し、M5ナット６個が別途必要。

オプションパーツ:
- [小型グラフィック液晶ボード](https://ssci.to/2608)

# ファームウェアソースコード

https://github.com/SWITCHSCIENCE/picossci-ffb-wheel

## ビルド

ビルドに必要なツール:
- go v1.23
- tinygo v0.33.0

```
git clone https://github.com/SWITCHSCIENCE/picossci-ffb-wheel
cd picossci-ffb-wheel
tinygo build -target pico -o build/diy-ffb-wheel.uf2 .
```

## 書き込み
ボードのBootSelボタンを押しながらPCに接続
「RPI-RP2」という名称のストレージが見えるようになります。
```
tinygo flash -target pico .
```
