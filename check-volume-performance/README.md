# Check Volume Performance

Check docker volume performance

## How to

```sh
$ docker compose up -d
$ docker exec -it app sh

# inside container
$ cd /var/test
$ time dd if=/dev/zero of=speedtest bs=1024 count=102400

# e.g. result of Intel Mac
# 102400+0 records in
# 102400+0 records out
# 104857600 bytes (100.0MB) copied, 55.297230 seconds, 1.8MB/s
# real	0m 55.30s
# user	0m 0.22s
# sys	0m 14.03s

# e.g. result of Macbook Air M2 (Docker version: v4.41)
102400+0 records in
102400+0 records out
104857600 bytes (100.0MB) copied, 3.367585 seconds, 29.7MB/s
real	0m 3.37s
user	0m 0.02s
sys	0m 1.72s
/var/test # exit

# 102400+0 records in
# 102400+0 records out
# 104857600 bytes (100.0MB) copied, 21.052339 seconds, 4.8MB/s
# real	0m 21.05s
# user	0m 0.18s
# sys	0m 1.99s
```

## What is this?

```
$ time dd if=/dev/zero of=speedtest bs=1024 count=102400
```

- `if=/dev/zero` は入力ファイルとして無限にゼロが続く特殊ファイル /dev/zero を指定
- `of=speedtest` は出力ファイル名を speedtest に指定
- `bs=1024` は1回の読み書き単位（ブロックサイズ）を1024バイト（1KB）に設定
- `count=102400` はこのブロックを102400回繰り返して書き込む

つまり、1024バイト × 102400回 = 104,857,600バイト（約100MB）のゼロデータを speedtest というファイルに書き込む処理をしている