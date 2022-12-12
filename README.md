# Terraform Provider Example Project

This Github Repository provides an example Terraform Provider. This can be used
for educational purposes or as starting point to create and extend your own
Provider.

On my Blog I documented the process in a bit more detail, check it our `HERE: To be added`

## Usage
You need Go & Terraform installed

### Build the binary of the Provider
This generates the go binary

```
cd terraform-provider-filr
make build
cd ..
```

### Run the Provider
After you have built the binary, you are good to run. I set the
TF_CLI_CONFIG_FILE to make sure that the local copy gets picked up.

```
export TF_CLI_CONFIG_FILE=dev.tfrc
terraform apply
```
