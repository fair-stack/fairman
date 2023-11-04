# FairMan

<img src="https://qn.0x3.cn/FairMan-1000x1000-%E6%A0%87%E5%87%86%E6%A0%B7%E5%BC%8F.png" align="right" width="240" />

This software provides capabilities such as cloud platform management, fundamental software oversight, image packaging, cluster orchestration, and software distribution. Image packaging involves encapsulating modules and operators (plugins such as tools, algorithms, models, etc., namely) through virtualization and containerization technologies, creating reproducible and deployable software images. On cluster orchestration, it allows concurrent deployment of interdependent modules on virtual machines and containers, enabling selective deployment and initial cluster size configuration. It dynamically scales the cluster by evaluating storage and computing resource costs. Software distribution involves disseminating modules and plugins through the software market, managing versions, monitoring operations, and ensuring consistent software stack performance across diverse hardware environments.

## 📦 Local development

### Environmental requirements

go 1.18

### Get the code

```bash
git clone https://github.com/LLiuHuan/fairman
```

### Startup instructions

#### Server startup instructions

```bash
# Enter the fairman backend project
cd ./fairman

# Update dependencies
go mod tidy

# Compile the project
go build

# Change setting 
# File path fairman/config_dev.yaml
vi ./config_dev.yaml
```
