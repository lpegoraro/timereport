build:
	docker run -v `pwd`:/defs namely/gen-grpc-gateway -f reportapi.proto -s ReportService -o gen/reportapi
	docker run -v `pwd`:/defs namely/protoc-all -f reportstorage.proto -l go -o cmd/timereport
	docker build gen/reportapi -t lpegoraro/timereport
run:
	docker run lpegoraro/reportapi
