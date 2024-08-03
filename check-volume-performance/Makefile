# after `docker compose up1
.PHONY: exec
exec:
	docker exec -it app sh
# cd /var/test
# time dd if=/dev/zero of=speedtest bs=1024 count=102400

# e.g. result of Intel Mac
# 102400+0 records in
# 102400+0 records out
# 104857600 bytes (100.0MB) copied, 55.297230 seconds, 1.8MB/s
# real	0m 55.30s
# user	0m 0.22s
# sys	0m 14.03s

# e.g. result of Apple Sillicon
# 102400+0 records in
# 102400+0 records out
# 104857600 bytes (100.0MB) copied, 21.052339 seconds, 4.8MB/s
# real	0m 21.05s
# user	0m 0.18s
# sys	0m 1.99s

.PHONY: clean
clean:
	rm -rf speedtest
