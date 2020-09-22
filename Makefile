
default: build

build:
	go install
	cd example; goguette -type=Foo,Bar
