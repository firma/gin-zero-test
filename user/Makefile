########## migrate ##########
migrate-new:
	docker run --rm -v ${PWD}/migration:/migration -w /migration migrate/migrate create -ext sql -dir /migration ${ACTION}

migrate-up:
	docker run --rm -v ${PWD}/migration:/migration -w /migration migrate/migrate -database="mysql://root:root@tcp(host.docker.internal:3306)/target_db?x-migrations-table=schema_migrations" -path /migration up

migrate-down:
	docker run --rm -v ${PWD}/migration:/migration -w /migration migrate/migrate -database="mysql://root:root@tcp(host.docker.internal:3306)/target_db?x-migrations-table=schema_migrations" -path /migration down -all

########## model ##########
model-gen:
	sh ${PWD}/script/dao_gen.sh
