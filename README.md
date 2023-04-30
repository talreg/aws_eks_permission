AWS sample tester
==========================================
This codebase is a sample of how to authenticate against service account iam role.

*Terraform files not included*

Preinstall:
---------------------
1. `make`
2. `docker`
3. `go` version 19 or above

Building
-------------------
You will need `make`. Just run `make`. The product will be under the local bin folder.<br>
Docker container: `make dockerize`


sample pod file
--------------------------
A sample pod file can be found [here](samples%2Fpod-sample.yaml)<br>
1. Please note that all the envars are mandatory!<br>
2. The last argument, `SELECTED_SECRET_NAME` has to be your value, you can't use the defaults here, unless you happen to have the same name.


Pushing to ECR
----------------------
NOTE! the order is important!<br>
1. First export the profile and the region: 
    `export AWS_PROFILE=<ACCOUNT/PROFILE ID>` and `export AWS_DEFAULT_REGION=<AWS REGION>`<br>
2. Make sure that the image is build (`make`)<br>
3. Login to the repo `aws ecr get-login-password --region <AWS REGION> | docker login --username AWS --password-stdin <ECR REPO ADDRESS>`<br>
4. Push