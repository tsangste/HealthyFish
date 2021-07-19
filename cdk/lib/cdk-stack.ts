import * as cdk from '@aws-cdk/core'
import { Vpc } from '@aws-cdk/aws-ec2'
import { Cluster, ContainerImage } from '@aws-cdk/aws-ecs'
import { ApplicationLoadBalancedFargateService } from '@aws-cdk/aws-ecs-patterns'

export class CdkStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props)

    const vpc = new Vpc(this,'MyVpc',{ maxAzs: 2 })
    const cluster = new Cluster(this ,'Cluster',{ vpc })

    const myService = new ApplicationLoadBalancedFargateService(this, 'HealthyFish', {
      cluster,
      desiredCount:3,
      taskImageOptions: {
        image: ContainerImage.fromAsset('apps/myapp')
      }
    })

    myService.targetGroup.configureHealthCheck({
      path: "/",
    })
  }
}
