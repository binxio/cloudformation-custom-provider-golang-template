Copier template for a CloudFormation Custom Resource Provider in Golang
======================================================================
This [copier](https://copier.readthedocs.io/) template  allows you to  create a complete custom resource provider in minutes!

Out-of-the-box features include:
- create the source for a custom cloudformation resource provider
- support for semantic versioning your provider using [git-release-tag](https://github.com/binxio/git-release-tag)
- distribute the provider to buckets in all AWS regions in the world
- deployable [AWS Codebuild](https://aws.amazon.com/codebuild/) pipeline

## getting started!
Let's say you want to create a custom resource for a Custom Domain of an AWS AppRunner service,
because it does not yet exist. To create the project, type:

```shell
pip install copier
copier https://github.com/binxio/cloudformation-custom-provider-golang-template /tmp/cfn-container-image-provider
🎤 the name of your custom resource type?
   ContainerImage
🎤 The name of your resource provider?
   cfn-container-image-provider
🎤 a short description for the custom provider?
   manages container images
🎤 golang version to use
   1.20
🎤 Your full name?
   Mark van Holsteijn
🎤 Your email address?
   mark@binx.io
🎤 the go module name
   github.com/binxio/cfn-container-image-provider
🎤 the URL to git source repository?
   https://github.com/binxio/cfn-container-image-provider.git
🎤 the golang executable name
   cfn-container-image-provider
🎤 the AWS region name
   eu-central-1
🎤 prefix for the S3 bucket name to store the lambda zipfiles?
   binxio-public
🎤 Access to lambda zip files?
   public

Copying from template version 0.0.0.post28.dev0+dfac895
 identical  .
    create  Makefile.mk
    create  .gitignore
    create  go.mod
    create  .buildspec.yaml
  conflict  .copier-answers.yml
 overwrite  .copier-answers.yml
    create  .dockerignore
    create  Dockerfile.lambda
    create  doc
    create  doc/ContainerImage.md
    create  .release
    create  Makefile
    create  cloudformation
    create  cloudformation/cicd-pipeline.yaml
    create  cloudformation/cfn-container-image-provider.yaml
    create  cloudformation/demo.yaml
    create  main.go

 > Running task 1 of 1: [ ! -f go.sum ] &&  (go mod download || echo "WARNING: failed to run go mod">&2); [ ! -d .git ] && ( git init && git add . && git commit -m 'initial import' && git tag 0.0.0) || exit 0
Initialized empty Git repository in ...
[main (root-commit) c97b9e2] initial import
 15 files changed, 529 insertions(+)
... 

````
This creates a project with a working custom provider for the resource `ContainerImage`. Change to 
the directory and type `make deploy-provider` and `make demo`. Your provider will be up-and-running
in less than 5 minutes!

## what is in the box
When you type `make help`, you will get a list of all of available actions.

```text
build                -  build the lambda zip file
fmt                  -  formats the source code

test                 -  run unit tests
test-templates       -  validate CloudFormation templates

deploy               -  AWS lambda zipfile to bucket
deploy-all-regions   -  AWS lambda zipfiles to all regional buckets
undeploy-all-regions -  deletes AWS lambda zipfile of this release from all buckets in all regions

deploy-provider      -  deploys the custom provider
delete-provider      -  deletes the custom provider

deploy-pipeline      -  deploys the CI/CD deployment pipeline
delete-pipeline      -  deletes the CI/CD deployment pipeline

deploy-demo          -  deploys the demo stack
delete-demo          -  deletes the demo stack

tag-patch-release    -  create a tag for a new patch release
tag-minor-release    -  create a tag for a new minor release
tag-major-release    -  create a tag for new major release

show-version         -  shows the current version of the workspace
help                 -  Show this help.
```

### run the unit tests
To run the unit tests, type:

```shell
$ make test
```

### Deploy the zip file to the bucket
To copy the zip file with the source code of the AWS Lambda of the custom resource provider, type:
```shell
BUCKET=<bucket-prefix>-<bucket-region>
aws s3 mb s3://$BUCKET
aws s3api put-bucket-ownership-controls \
    --bucket $BUCKET --ownership-controls \
    'Rules=[{ObjectOwnership=BucketOwnerPreferred}]'
```
As you can see, the zipfile will be copied to a bucket name which consists of the prefix
and the region name.  This allows the zipfile to be made available for use in
all regions.

If you want to allow public access to the bucket, type:
```shell
aws s3api put-public-access-block \
   --bucket $BUCKET  \
   --public-access-block-configuration \
   "BlockPublicAcls=false,IgnorePublicAcls=false,BlockPublicPolicy=false,RestrictPublicBuckets=false"
```

### Deploy the custom resource provider into the account
Now the zip file is available, you deploy the custom resource provider, by typing:
```shell
$ make deploy-provider
```
This deploys the provider as an AWS Lambda function. To configure
the run-time parameters and permissions of the Lambda change the CloudFormation
template in the directory `./cloudformation`.

### Deploy the custom resource demo
To deploy the demo CloudFormation stack using the custom resource provider, type:
```shell
$ make deploy-demo
```
This deploys an CloudFormation stack with an example custom resource as a CloudFormation stack.
the run-time parameters and permissions of the Lambda change the CloudFormation
template in the file `./cloudformation/demo.yaml`. Change the configuration of the custom
resource to match your implementation.

### Version your custom resource provider
To version your custom resource provider, you can use the following commands:

```text
make tag-patch-release    -  create a tag for a new patch release
make tag-minor-release    -  create a tag for a new minor release
make tag-major-release    -  create a tag for new major release
```

This will:
- run the pre-tag command in the file `./release`
- commit all outstanding changes in the workspace
- tag the commit with the new version.

To show the current version of the workspace, type:

```shell
make show-version
```

The utility [git-release-tag](https://github.com/binxio/git-release-tag)
implements this functionality.

### Deploy provider to all regions
To deploy the current version of your provider to all regions, type:

```shell
make deploy-all-regions
```
This assumes you have buckets in all regions with the defined prefix.

### Deploy CI/CD pipeline
To deploy the CI/CD pipeline based on AWS Codebuild, make sure that the AWS account can 
access the source repository. If that is the case, type:

```shell
make deploy-pipeline
```
Now every time you tag a new release, it will automatically be deployed to all regions.

## conclusion
This [copier](https://copier.readthedocs.io/) template provides everything you need to quickly
build, deploy and maintain a new custom AWS CloudFormation Provider!
