# go-aws-s3-bucket-list

export AWS_VAULT_FILE_PASSPHRASE="$(cat /root/.awsvaultk)"

aws-vault exec dev -- terraform -chdir=./terraform init
aws-vault exec dev -- terraform -chdir=./terraform apply --auto-approve

sh ./terraform/terraform.tmp

go mod init awsS3BucketList
go mod tidy
go run cmd/awsS3BucketList/main.go
go build -o awsS3BucketList cmd/awsS3BucketList/main.go
