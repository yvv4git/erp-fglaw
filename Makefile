BUILD_DIR=./

build:
	GOOS=windows GOARCH=386 go build -o ${BUILD_DIR}/erpfglaw.exe main.go
	GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/erpfglaw64.exe main.go
	GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/erpfglaw64.bin main.go
	GOOS=linux GOARCH=arm go build -o ${BUILD_DIR}/erpfglawArm.bin main.go
	GOOS=freebsd GOARCH=amd64 go build -o ${BUILD_DIR}/erpfglawFreebsd64.bin main.go

tests_integration:
	go test -v -tags=integration ./...

tests_unit:
	go test -v -tags=unit ./...

run:
	go run main.go

migrations_init:
	sql-migrate new -env="development" create_clients_table
	sql-migrate new -env="development" create_client_types
	sql-migrate new -env="development" create_deposite
	sql-migrate new -env="development" file_types
	sql-migrate new -env="development" create_photo
	sql-migrate new -env="development" create_product_types
	sql-migrate new -env="development" create_products
	sql-migrate new -env="development" create_stock
	sql-migrate new -env="development" create_suppliers
	sql-migrate new -env="development" create_transactions
	sql-migrate new -env="development" create_transaction_status_types

migrations_dev_up:
	sql-migrate up -env="development"

migrations_dev_down:
	sql-migrate down -env="development" -dryrun

migrations_prod_up:
	sql-migrate up -env="production"

migrations_prod_down:
	sql-migrate down -env="production" -dryrun