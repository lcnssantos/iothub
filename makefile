dev:
	docker-compose -f infra/dev/docker-compose.yml up -d

terraform:
	docker-compose -f infra/terraform/docker-compose.yml up -d 

terraform-init:
	docker-compose -f infra/terraform/docker-compose.yml run --rm terraform init

terraform-plan:
	docker-compose -f infra/terraform/docker-compose.yml run --rm terraform plan

terraform-validate:
	docker-compose -f infra/terraform/docker-compose.yml run --rm terraform validate

terraform-apply:
	docker-compose -f infra/terraform/docker-compose.yml run --rm terraform apply -auto-approve

terraform-destroy:
	docker-compose -f infra/terraform/docker-compose.yml run --rm terraform destroy -auto-approve

terraform-deploy:
	make terraform-validate
	make terraform-plan
	make terraform-apply