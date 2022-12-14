# Terraform Provider Example Project

This Github Repository provides an example Terraform Provider. This can be used
for educational purposes or as starting point to create and extend your own
Provider.

On my Blog I documented the process in a bit more detail, check it our `HERE: To be added`

## Usage
You need Go & Terraform installed

### Build the binary of the Provider
This generates the binary of the Provider which is then used by Terraform.
There is no way (as far as I'm aware of) to directly use Go code as Provider
without compiling it.

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

### Output
The Provider creates two (or how many you specify in the main.tf file) files in
the data directory. The name is a random UUID and gets tracked by Terraform and
its state.

Overview of the output
```
.
├── data
│   ├── bd4064f7-bb98-cde4-73e6-dcc3abb8fbd4
│   └── ef092c14-3145-71f3-87b2-3f7942be9fdb
├── main.tf
└── terraform.tfstate
```
