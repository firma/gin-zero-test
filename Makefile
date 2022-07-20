########## migrate ##########
migrate-new:
	docker run --rm -v ${PWD}/api/migration:/migration -w /migration migrate/migrate create -ext sql -dir /migration ${ACTION}

migrate-up:
	docker run --rm -v ${PWD}/api/migration:/migration -w /migration migrate/migrate -database="mysql://root:root@tcp(host.docker.internal:3306)/target_db?x-migrations-table=schema_migrations" -path /migration up

migrate-down:
	docker run --rm -v ${PWD}/api/migration:/migration -w /migration migrate/migrate -database="mysql://root:root@tcp(host.docker.internal:3306)/target_db?x-migrations-table=schema_migrations" -path /migration down -all

########## model ##########
model-api-gen:
	sh ${PWD}/api/script/dao_gen.sh

model-user-rpc-gen:
	sh ${PWD}/user/script/dao_gen.sh
