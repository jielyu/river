# river

**开发中......**

这些年一直都在吐槽C++缺乏一个好用的包管理工具，也一直在期待官方的作为，很显然目前还没等到。在Rust生态逐渐扩大的时候，了解了一些Cargo包管理工具，令人赞叹不已。现在包管理机制做得比较好的还有 golang 和 dart 等。这时，我决定试着开发一个好用的C++的包管理工具，类似Rust的Cargo。不知道能不能完成，也不知道需要多久完成，不抱太大期待，就当是打发时间。

为了方便开发，减少个人惰性带来的阻力，本项目在开发阶段只使用中文；后续如果能发布成品的话，再考虑国际化。

## 已有功能

无

## 未来功能

### 1. 创建工程

创建可执行工程

```shell
river new binary_proj_name
```

创建库工程

```shell
river new lib_proj_name --lib
```

### 2. 编译

```shell
river build
```

### 3. 运行


```shell
river run
```

### 4. 添加依赖包

```shell
river add package_name
```

### 5. 升级依赖包

只升级到最新的小版本，即 `a.b.x`

```shell
river upgrade
```

### 6. 发布自定义包

```shell
river publish
```





