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

