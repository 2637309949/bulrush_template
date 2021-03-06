all: target

target: dependence doc bin migration clean

dependence:
	@echo "\033[32mInstalling dependence\033[0m"
	@go get  github.com/json-iterator/go
	@go mod tidy
bin:
	@echo "\033[32mBuilding bin file\033[0m"
	@go build -tags=jsoniter -o app .

doc:
	@echo "\033[32mBuilding doc file\033[0m"
	@apidoc

migration:
	@echo "\033[32mMigrating files\033[0m"
	@rm -rf ./build
	@mkdir -p ./build/conf
	@cp -rf Dockerfile ./build
	@cp -rf assets ./build
	@cp -rf doc/api_data.js ./build
	@cp -rf doc/api_project.js ./build
	@cd conf && cp -rf `ls | grep -v index.go | xargs` ../build/conf && cd ../
	@cp -rf logs ./build
	@cp app ./build
clean:
	@echo "\033[32mClean files\033[0m"
	@rm -rf app

.PHONY: all target dependence doc bin migration clean 