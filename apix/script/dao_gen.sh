#!/bin/sh
echo ${PWD}
model_cmd=$1
name="dao_gen"
network_name="dao_gen_net"

network_name_check=`docker network ls|grep ${network_name}|awk '{print $2}'`
if [ -z "${network_name_check}" ]; then
    docker network create ${network_name}

    echo "create docker network ${network_name}\n"
fi

dao_gen_container_check=`docker ps -a|grep ${name}|awk '{print $1}'`
if [ -z "${dao_gen_container_check}" ]; then
    docker run --platform linux/x86_64 --network=${network_name} --name ${name} -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
    echo "create docker container ${name}\nwait for a while...\n"

    alive_check=`docker run --rm --network=dao_gen_net mysql:5.7 mysqladmin ping -u root --password=root -h dao_gen -P 3306 --silent|grep alive`
    while [ -z "${alive_check}" ]; do
        echo "check mysql alive...\n"
        alive_check=`docker run --rm --network=dao_gen_net mysql:5.7 mysqladmin ping -u root --password=root -h dao_gen -P 3306 --silent|grep alive`
        sleep 1
    done

    docker exec -it ${name} mysql --password=root -e "show databases;"
    docker exec -it ${name} mysql --password=root -e "drop database if exists dao_gen_db;"
    docker exec -it ${name} mysql --password=root -e "create database dao_gen_db;"
fi

echo "started ${name} container\n"

docker run --rm --network=${network_name} -v ${PWD}/api/migration:/migration -w /migration migrate/migrate -database="mysql://root:root@tcp(dao_gen:3306)/dao_gen_db?x-migrations-table=schema_migrations" -path /migration up
echo "migrated \n"

go run ${PWD}/api/cmd/modelgen/main.go

docker stop ${name}
docker rm ${name}
docker network rm ${network_name}
