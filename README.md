# 🔋Go_Web 开发脚手架
- mysql
- gin 框架
- sqlx 操作数据库
- 采用 viper 来配置项目
- docker 部署项目
- air 实现项目热加载

项目目录结构：

```sh
├─api            # 接口模型
├─config         # 配置文件
├─controller     # 控制层
├─dao            # 操作数据库
│  ├─mysql       
│  └─redis
├─docs           # swagger文档
├─log            # 日志操作
├─model          # 数据模型
├─pkg            # 第三方库
├─repository     # 数据库建表
├─router         # 路由
└─service        # 业务
```