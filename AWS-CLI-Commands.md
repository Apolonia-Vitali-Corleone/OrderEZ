# 设置Region
REGION=ap-east-2

# 获取Region的全部ECR Repo（只有名称）

```bash
REGION=ap-south-1

aws ecr describe-repositories --region "$REGION" --query 'repositories[].repositoryName' --output table
```



# 设置对应的Repo
REPO=gcy-app

# 获取Repo的所有已打tag镜像的 “Tag + Digest”
```bash
REPO=gcy-app
aws ecr list-images --region "$REGION" --repository-name "$REPO" \
  --filter tagStatus=TAGGED \
  --query 'imageIds[].{Tag:imageTag, Digest:imageDigest}' \
  --output table
```



# 获取Repo的所有没有tag镜像的Digest
```
aws ecr list-images --region "$REGION" --repository-name "$REPO" --filter tagStatus=UNTAGGED \
  --query 'imageIds[].{Digest:imageDigest}' --output table
```



# 获取Repo的所有镜像的Digest和Tag
aws ecr list-images --region "$REGION" --repository-name "$REPO" \
  --query 'imageIds[].{Tag:imageTag, Digest:imageDigest}' --output table

# 安装 jq
sudo apt update && sudo apt install -y jq

# 删除全部的有tag的Image
```
aws ecr list-images --region "$REGION" --repository-name "$REPO" \
  --filter tagStatus=TAGGED --output json \
| jq -r '.imageIds[] | "imageTag=\(.imageTag)"' \
| xargs -n 50 aws ecr batch-delete-image --region "$REGION" --repository-name "$REPO" --image-ids
```



# 获取Repo的所有已打tag镜像的 “Tag + Digest”
aws ecr list-images --region "$REGION" --repository-name "$REPO" \
  --filter tagStatus=TAGGED \
  --query 'imageIds[].{Tag:imageTag, Digest:imageDigest}' \
  --output table
补充：没有任何输出，因为都删除完了

# 删除全部的无tag的Image
```
aws ecr list-images --region "$REGION" --repository-name "$REPO" \
  --filter tagStatus=UNTAGGED --output json \
| jq -r '.imageIds[]? | select(.imageDigest) | "imageDigest=\(.imageDigest)"' \
| xargs -r -n 50 aws ecr batch-delete-image --region "$REGION" --repository-name "$REPO" --image-ids
```



# 获取Repo的所有没有tag镜像的Digest
aws ecr list-images --region "$REGION" --repository-name "$REPO" --filter tagStatus=UNTAGGED \
  --query 'imageIds[].{Digest:imageDigest}' --output table
补充：没有任何输出，因为都删除完了

# 获取Repo的所有镜像的Digest和Tag
aws ecr list-images --region "$REGION" --repository-name "$REPO" \
  --query 'imageIds[].{Tag:imageTag, Digest:imageDigest}' --output table
补充：没有任何输出，因为都删除完了


# 查看全部VPC---当前Region的以及VPC-ID
```bash

aws ec2 describe-vpcs \
  --query 'Vpcs[*].{
    VpcId: VpcId,
    Name: Tags[?Key==`Name`].Value|[0], 
    CidrBlock: CidrBlock,
    State: State
  }' \
  --output table
```



# 查看全部Subnet---当前VPC中的（需要在Values=输入VPCID）
```bash
VPC_ID=vpc-06ac1e3cb683ced3e

aws ec2 describe-subnets \
  --filters "Name=vpc-id,Values=$VPC_ID" \
  --query 'Subnets[*].{
    SubnetId: SubnetId,
    Name: Tags[?Key==`Name`].Value|[0],
    CidrBlock: CidrBlock,
    AZ: AvailabilityZone,
    State: State
  }' \
  --output table
```



# 查看security group---当前region全部（只展示关键信息）
```bash
aws ec2 describe-security-groups \
  --query 'SecurityGroups[*].{
    GroupId: GroupId,          
    GroupName: GroupName,      
    VpcId: VpcId,              
    Description: Description,  
    Name: Tags[?Key==`Name`].Value|[0]
  }' \
  --output table
```



ALB_SG_ID=sg-073278699b315ce7c



# 创建ALB---根据NAME_PREFIX、SUBNET、ALB_SG_ID
```bash
NAME_PREFIX=order-ez
SUBNET1=subnet-0b01e08b0a1eb2896
SUBNET2=subnet-0f51eef8f1231e5c6
SUBNET3=subnet-0ee659357bb931f38
ALB_SG_ID=sg-05bf106de7ec670a1
REGION=ap-south-1

ALB_ARN=$(aws elbv2 create-load-balancer \
  --name ${NAME_PREFIX}-alb \
  --type application \
  --scheme internet-facing \
  --subnets $SUBNET1 $SUBNET2 $SUBNET3 \
  --security-groups $ALB_SG_ID \
  --ip-address-type ipv4 \
  --region $REGION \
  --query 'LoadBalancers[0].LoadBalancerArn' --output text) && echo $ALB_ARN
```



# 创建ALB_DNS---根据ALB_ARN
```bash
ALB_ARN=arn:aws:elasticloadbalancing:ap-south-1:818719120332:loadbalancer/app/order-ez-alb/d2b4ce6850319d5a
REGION=ap-south-1
ALB_DNS=$(aws elbv2 describe-load-balancers \
  --load-balancer-arns $ALB_ARN --region $REGION \
  --query 'LoadBalancers[0].DNSName' --output text) && echo "ALB DNS: $ALB_DNS"
```



# 查看当前Region的全部的Load-Balancer

```bash
aws elbv2 describe-load-balancers \
  --query 'LoadBalancers[*].{Name:LoadBalancerName, ARN:LoadBalancerArn}' \
  --output table
```



# 查看ALB_DNS---根据ALB_NAME查看ALB的状态DNS
```bash
REGION=ap-south-1
ALB_NAME="order-ez-alb"
aws elbv2 describe-load-balancers --region "$REGION" --names "$ALB_NAME" \
  --query 'LoadBalancers[0].{Name:LoadBalancerName,State:State.Code,DNS:DNSName,Scheme:Scheme,Type:Type}' --output table
```



# 解析ARN-DNS

```bash
ALB_DNS=order-ez-alb-1876390975.ap-east-2.elb.amazonaws.com

nslookup "$ALB_DNS" || dig +short "$ALB_DNS"
```





# 根据ALB的Name查看 ALB 自身的运行状态

```bash
aws elbv2 describe-load-balancers \
  --names order-ez-alb \
  --query 'LoadBalancers[0].State' \
  --output text
```

返回

```
active
```



# 根据ALB_NAME查看ALB的监听器

```
ALB_NAME=order-ez-alb

aws elbv2 describe-listeners \
  --load-balancer-arn $(aws elbv2 describe-load-balancers --names $ALB_NAME --query 'LoadBalancers[0].LoadBalancerArn' --output text) \
  --output table
```



# 根据TARGET_GROUP_NAME和VPC_ID创建一个可以接受 IP 地址作为目标的目标组

```bash
TARGET_GROUP_NAME=dummy-target-group
VPC_ID=vpc-0ffa225f05aa313ff

aws elbv2 create-target-group \
  --name $TARGET_GROUP_NAME \
  --protocol HTTP \
  --port 80 \
  --vpc-id $VPC_ID\
  --target-type ip
```





# 根据TARGET_GROUP_NAME和ALB_ARN创建80端口侦听器

```

TARGET_GROUP_NAME=dummy-target-group
ALB_ARN=arn:aws:elasticloadbalancing:ap-east-2:818719120332:loadbalancer/app/order-ez-alb/17986e768fc53e7f

aws elbv2 create-listener \
  --load-balancer-arn $ALB_ARN \
  --protocol HTTP \
  --port 80 \
  --default-actions Type=forward,TargetGroupArn=$(aws elbv2 describe-target-groups --names $TARGET_GROUP_NAME --query 'TargetGroups[0].TargetGroupArn' --output text)
```

# 根据TARGET_GROUP_NAME和ALB_ARN创建443端口侦听器

```bash

TARGET_GROUP_NAME=dummy-target-group
ALB_ARN=arn:aws:elasticloadbalancing:ap-east-2:818719120332:loadbalancer/app/order-ez-alb/17986e768fc53e7f

aws elbv2 create-listener \
  --load-balancer-arn $ALB_ARN \
  --protocol HTTPS \
  --port 443 \
  --default-actions Type=forward,TargetGroupArn=$(aws elbv2 describe-target-groups --names $TARGET_GROUP_NAME --query 'TargetGroups[0].TargetGroupArn' --output text)
```

输出

```
An error occurred (ValidationError) when calling the CreateListener operation: A certificate must be specified for HTTPS listeners
```



# 根据ALB_ARN查看ALB的Listeners

```bash
ALB_ARN=
aws elbv2 describe-listeners --load-balancer-arn $ALB_ARN --output table
```





```

TARGET_GROUP_NAME=dummy-target-group

aws elbv2 describe-target-health \
  --target-group-arn $(aws elbv2 describe-target-groups --names $TARGET_GROUP_NAME --query 'TargetGroups[0].TargetGroupArn' --output text)
```





```
export IMAGE_URI="<account>.dkr.ecr.${REGION}.amazonaws.com/<repo>:<tag>"
export EXEC_ROLE_ARN="arn:aws:iam::<account>:role/ecsTaskExecutionRole"
export TASK_ROLE_ARN="arn:aws:iam::<account>:role/ecsAppRole"
export LOG_GROUP="/ecs/dev-web"
```





创建ALB

创建SG



mysql

admin

12345678

order-ez-instance-1.cn6i6q8c6duc.ap-south-1.rds.amazonaws.com:3306

order-ez-redis-cnpbde.serverless.aps1.cache.amazonaws.com:6379

redis



rabbitmq

guetst

guestguestguest

amqps://guetst:guestguestguest@b-ebc9761e-87c4-4c32-a0c7-61e129379c36.mq.ap-south-1.on.aws:5671/%2F

```
{
  "containerDefinitions": [
    {
      "name": "your-go-service",
      "image": "your-image:tag",
      "environment": [
        { "name": "MYSQL_ADDR", "value": "my-aurora-cluster.cluster-xxx.rds.amazonaws.com:3306" },
        { "name": "REDIS_ADDR", "value": "my-redis-cluster.eaogs8.0001.ap-south-1.cache.amazonaws.com:6379" },
        { "name": "RABBITMQ_ADDR", "value": "amqps://mybroker.mq.ap-south-1.amazonaws.com:5671" }
      ]
    }
  ]
}
```



# 创建测试EC2

## 先设变量（改成你的值）

```
REGION=ap-south-1
VPC_ID=vpc-06ac1e3cb683ced3e
SUBNET_ID=subnet-0724d0b1e091dcefc      # 你的“目标子网”（和要测的 DB/Redis/RabbitMQ 在同 VPC/可达路由）
INSTANCE_NAME=vpc-test-shell
INSTANCE_TYPE=t3.micro     
```

## 安全组（入站空、仅出站放行）

> 作为“测试壳子”不需要任何入站规则；只要默认的全部出站即可。

```
SG_ID=$(aws ec2 create-security-group \
  --region $REGION \
  --group-name ${INSTANCE_NAME}-sg \
  --description "Outbound-only SG for VPC test shell" \
  --vpc-id $VPC_ID \
  --query 'GroupId' --output text)

# 出站默认 Allow All（若你们安全基线禁止了，至少放行到目标端口/网段）

```

##  给实例挂 SSM 托管角色（免公网、免 SSH）

```
aws iam create-role --role-name EC2SSMRole \
  --assume-role-policy-document '{
    "Version":"2012-10-17",
    "Statement":[{"Effect":"Allow","Principal":{"Service":"ec2.amazonaws.com"},"Action":"sts:AssumeRole"}]
  }' >/dev/null 2>&1 || true

aws iam attach-role-policy --role-name EC2SSMRole \
  --policy-arn arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore >/dev/null 2>&1 || true

aws iam create-instance-profile --instance-profile-name EC2SSMInstanceProfile >/dev/null 2>&1 || true
aws iam add-role-to-instance-profile --instance-profile-name EC2SSMInstanceProfile --role-name EC2SSMRole >/dev/null 2>&1 || true
sleep 10

```

## 找最新 Amazon Linux 2023 AMI

```
$REGION=

AMI_ID=$(aws ssm get-parameters \
  --names /aws/service/ami-amazon-linux-latest/al2023-ami-kernel-6.1-x86_64 \
  --query 'Parameters[0].Value' --output text --region ${REGION})
echo AMI_ID=$AMI_ID
```



## User Data：安装常用连通性/客户端工具

```
cat > /tmp/userdata.sh <<"EOF"
#!/bin/bash
set -euxo pipefail

dnf -y update
dnf -y install jq curl bind-utils iproute nc telnet traceroute tcpdump openssl
dnf -y install mysql postgresql redis
# rabbitmqadmin（通过 pip 装，纯客户端，走 15672 HTTP API）
dnf -y install python3-pip
pip3 install --break-system-packages rabbitmq-admin==0.0.2 || true

# 确保 SSM agent 运行
systemctl enable --now amazon-ssm-agent || true
EOF

USER_DATA_B64=$(base64 -w0 /tmp/userdata.sh)

```



## 启动“测试壳子”EC2（私网，无公网 IP）

```
AMI_ID=ami-06fa3f12191aa3337
INSTANCE_ID=$(aws ec2 run-instances \
  --region $REGION \
  --image-id $AMI_ID \
  --instance-type $INSTANCE_TYPE \
  --subnet-id $SUBNET_ID \
  --security-group-ids $SG_ID \
  --iam-instance-profile Name=EC2SSMInstanceProfile \
  --tag-specifications "ResourceType=instance,Tags=[{Key=Name,Value=vpc-test-shell}]" \
  --user-data fileb:///tmp/userdata.sh \
  --no-associate-public-ip-address \
  --query 'Instances[0].InstanceId' --output text)
echo INSTANCE_ID=$INSTANCE_ID
aws ec2 wait instance-status-ok --region $REGION --instance-ids $INSTANCE_ID

```

##  直接开壳（SSM）

```

```





```
RABBIT_HOST="b-ebc9761e-87c4-4c32-a0c7-61e129379c36.mq.ap-south-1.on.aws"   # 例：b-abc123.mq.ap-southeast-1.amazonaws.com
AMQP_PLAIN_PORT=5672                 # 明文 AMQP（很多托管环境禁用）
AMQP_TLS_PORT=5671                   # TLS AMQP（Amazon MQ 常用）
MGMT_HTTP=15672                      # 管理控制台 HTTP（自建常用）
MGMT_HTTPS=15671                     # 管理控制台 HTTPS（常用）

```



```
aws ec2 terminate-instances --instance-ids $(aws ec2 describe-instances --query "Reservations[*].Instances[*].InstanceId" --output text)
```

