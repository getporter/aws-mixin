# AWS CLI Mixin for Porter

This is a mixin for Porter that provides the AWS CLI.

The [Buckets Example](./examples/buckets) provides a full working bundle demonstrating how to use this mixin.

## Mixin Syntax

See the [AWS CLI Command Reference](https://docs.aws.amazon.com/cli/latest/reference/index.html#cli-aws) for the supported commands

```yaml
aws:
  description: "Description of the command"
  service: SERVICE
  operation: OP
  arguments:
  - arg1
  - arg2
  flags:
    FLAGNAME: FLAGVALUE
    REPEATED_FLAG:
    - FLAGVALUE1
    - FLAGVALUE2
```

## Examples

### Provision a VM

```yaml
aws:
  description: "Provision VM"
  service: ec2
  operation: run-instances
  flags:
    image-id: ami-xxxxxxxx
    instance-type: t2.micro
```

### Create a Bucket

```yaml
aws:
  description: "Create Bucket"
  service: s3api
  operation: create-bucket
  flags:
    bucket: my-buckkit
    region: us-east-1
```