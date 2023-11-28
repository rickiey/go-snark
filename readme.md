# 卡机 WGPU

## 编译

> make all

### 注意

* 编译时启用 supra,需要 ubuntu 20.04.4 cuda 12 以上版本

> export FFI_USE_CUDA_SUPRASEAL=1
如果遇到 /usr/bin/ld: cannot find -lcudart_static 错误
参考 [官方文档](https://lotus.filecoin.io/tutorials/lotus-miner/supra-seal/)
> find /usr -name "libcudart_static*"
> export LIBRARY_PATH=$LIBRARY_PATH:/usr/local/cuda-12.2/targets/x86_64-linux/lib

* 多卡机使用多个进程时需要绑定显卡，并且使用不同的 TMPDIR 目录（默认/tmp, 更改不同的 TMPDIR 环境变量 ），否则会出现显卡锁冲突。
    > export TMPDIR=/tmp/gpu_lock_00
    > export TMPDIR=/tmp/gpu_lock_01
