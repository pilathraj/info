# Terraform

```hcl
provider "aws" {
  region = "ap-south-1"  
}
resource "provider_<resource_type>" "name"{
// Config options
 key1 = "value1"
 key2 = "value2"
}

```
```
terraform -v  ===> return version number
terraform init  ==> initialize terraform project
terraform plan
terraform apply
terraform destroy
```
