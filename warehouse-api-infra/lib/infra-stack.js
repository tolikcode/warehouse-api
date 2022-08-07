const { Stack } = require("aws-cdk-lib");
const ec2 = require("aws-cdk-lib/aws-ec2");
const ecs = require("aws-cdk-lib/aws-ecs")
const ecr = require("aws-cdk-lib/aws-ecr")
const ecsPatterns = require("aws-cdk-lib/aws-ecs-patterns");
const logs = require("aws-cdk-lib/aws-logs")

class WarehouseStack extends Stack {
  constructor(scope, id, props) {
    super(scope, id, props);

    var vpc = ec2.Vpc.fromLookup(this, "my-vpc", {
      isDefault: true,
    });

    const cluster = new ecs.Cluster(this, 'warehouse-cluster', {
      clusterName: 'warehouse-cluster',
      containerInsights: true,
      vpc: vpc,
    });

    const repository = new ecr.Repository(this, 'warehouse-ecr-repo', {
      repositoryName: 'warehouse-ecr-repo',
    })

    const image = ecs.ContainerImage.fromEcrRepository(repository);

    const loadBalancedFargateService = new ecsPatterns.ApplicationLoadBalancedFargateService(
      this,
      'warehouse-service',
      {
        cluster,
        circuitBreaker: {
          rollback: true,
        },
        cpu: 512,
        desiredCount: 1,
        memoryLimitMiB: 1024,
        taskImageOptions: {
          image: image,
          containerPort: 8080,
          logDriver: ecs.LogDrivers.awsLogs({
            streamPrefix: id,
            logRetention: logs.RetentionDays.ONE_DAY,
          }),
        },
      }
    );
  }
}

module.exports = { WarehouseStack };
