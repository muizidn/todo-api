config:
	echo "# Environment Config" > config.env
	echo "TODO_DBNAME=${DBNAME}" >> config.env
	echo "TODO_DBUSER=${DBUSER}" >> config.env
	echo "TODO_DBPASS=${DBPASS}" >> config.env
	echo "TODO_DBHOST=${DBHOST}" >> config.env
	echo "TODO_DBPORT=${DBPORT}" >> config.env
	echo "TODO_JWTSECRET=${JWTSECRET}" >> config.env

proto_gen_dir := pb
proto_dir := submodules/proto

protogen:
	rm -rf ${proto_gen_dir}
	make -C ${proto_dir} go
	mkdir -p ${proto_gen_dir}
	mv ${proto_dir}/proto/*.go ${proto_gen_dir}

submodule:
	git submodule update --init --recursive

bootstrap: submodule protobuf config
	 
test:
	go test ./...

run:
	go run .

ci: bootstrap test
