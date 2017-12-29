部署Cloud Function时，出现如下错误:

```bash
ERROR: gcloud crashed (AttributeError): 'str' object has no attribute 'items'

If you would like to report this issue, please run the following command:
  gcloud feedback

To check gcloud for common problems, please run the following command:
  gcloud info --run-diagnostics
```

部署脚本如下：

```bash
#!/bin/bash

function=PostgresDemo
gcloud functions deploy ${function} --runtime go111 --trigger-http --memory=128 --env-vars-file .env.yaml
```

原因是环境变量文件`.env.yaml`的语法错误，`.env.yaml`如下：

```yaml
POSTGRES_INSTANCE_CONNECTION_NAME:xxx
POSTGRES_USER:xxx
POSTGRES_PASSWORD:xxx
POSTGRES_DB:xxx
```

环境变量`Key`与`Value`之间没有空格。修改后正确的`.env.yaml`文件如下：

```yaml
POSTGRES_INSTANCE_CONNECTION_NAME: xxx
POSTGRES_USER: xxx
POSTGRES_PASSWORD: xxx
POSTGRES_DB: xxx
```

