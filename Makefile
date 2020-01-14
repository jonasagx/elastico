try: clean build

build:
	@echo "building elastico"
	pushd src && go build main.go && mv main ../elastico && popd

clean:
	rm elastico &