import * as cdk from '@aws-cdk/core'
import { Vpc } from '@aws-cdk/aws-ec2'
import { Cluster, ContainerImage } from '@aws-cdk/aws-ecs'
import { Repository } from '@aws-cdk/aws-ecr'
import { ApplicationLoadBalancedFargateService } from '@aws-cdk/aws-ecs-patterns'

export class CdkStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props)

    const vpc = new Vpc(this,'MyVpc',{ maxAzs: 2 })
    const cluster = new Cluster(this ,'Cluster',{ vpc })

    const repo = new Repository(this, 'healthyfish')

    const myService = new ApplicationLoadBalancedFargateService(this, 'HealthyFish', {
      cluster,
      desiredCount:1,
      taskImageOptions: {
        image: ContainerImage.fromEcrRepository(repo, '1.0.1'),
        containerPort: 8080
      }
    })

    myService.targetGroup.configureHealthCheck({
      path: "/",
    })
  }
}
