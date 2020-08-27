build:
	docker run -v `pwd`:/defs namely/gen-grpc-gateway -f proto/api/reportapi.proto -s ReportService -o api
	docker run -v `pwd`:/defs namely/protoc-all -f proto/storage/reportstorage.proto -l go -o storage
	docker build gen/reportapi -t lpegoraro/timereport
run:
	docker run lpegoraro/reportapi
