# 卡机 WGPU

## 编译

> make all

### 注意

* 编译时启用 supra,需要 ubuntu 20.04 cuda 12 以上版本

> export FFI_USE_CUDA_SUPRASEAL=1

* 多卡机使用多个进程时需要绑定显卡，并且使用不同的 TMPDIR 目录（默认/tmp, 更改不同的 TMPDIR 环境变量 ），否则会出现显卡锁冲突。
    > export TMPDIR=/tmp/gpu_lock_00
    > export TMPDIR=/tmp/gpu_lock_01
