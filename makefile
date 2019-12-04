# this makefile used in console environment.
# copy this file to project base directory.

# SET BIN NAME BY USER
binName:=pgo2-demo

goBin:=go

######## DO NOT CHANGE THE FLOWING CONTENT ########

# absolute path of makefile
mkPath:=$(abspath $(firstword $(MAKEFILE_LIST)))

# absolute base directory of project
baseDir:=$(strip $(patsubst %/, %, $(dir $(mkPath))))

binDir:=$(baseDir)/bin
pkgDir:=$(baseDir)/pkg

.PHONY: start stop build update pgo init

start: build
	$(binDir)/$(binName)

stop:
	-killall $(binName)

build:
	export GOPATH=$(baseDir) && $(goBin) build -o $(binDir)/$(binName) $(pkgDir)/Main/main.go

update:
	export GOPATH=$(baseDir) && cd $(pkgDir) && $(glideBin) update

install:
	cd $(baseDir) && $(glideBin) install

pgo2:
	cd $(baseDir) && $(goBin) get -u github.com/pinguo/pgo2

init:
    [ -d $(baseDir)/cmd ] || mkdir $(baseDir)/cmd
    [ -d $(baseDir)/cmd/$(binName) ] || mkdir $(baseDir)/cmd/$(binName)
    [ -d $(baseDir)/web ] || mkdir $(baseDir)/web
    [ -d $(baseDir)/web/app ] || mkdir $(baseDir)/web/app
    [ -d $(baseDir)/web/template ] || mkdir $(baseDir)/web/template
    [ -d $(baseDir)/web/static ] || mkdir $(baseDir)/web/static
	[ -d $(baseDir)/configs ] || mkdir $(baseDir)/configs
	[ -d $(baseDir)/assets ] || mkdir $(baseDir)/assets
	[ -d $(baseDir)/build ] || mkdir $(baseDir)/build
	[ -d $(pkgDir) ] || mkdir $(pkgDir)
	[ -d $(pkgDir)/command ] || mkdir $(pkgDir)/command
	[ -d $(pkgDir)/controller ] || mkdir $(pkgDir)/controller
	[ -d $(pkgDir)/lib ] || mkdir $(pkgDir)/lib
	[ -d $(pkgDir)/model ] || mkdir $(pkgDir)/model
	[ -d $(pkgDir)/service ] || mkdir $(pkgDir)/service
	[ -d $(pkgDir)/struct ] || mkdir $(pkgDir)/struct
	[ -d $(pkgDir)/test ] || mkdir $(pkgDir)/test
	[ -f $(pkgDir)/go.mod ] || (cd $(baseDir) && $(goBin) init $(binName))

help:
	@echo "make start       build and start $(binName)"
	@echo "make stop        stop process $(binName)"
	@echo "make build       build $(binName)"
	@echo "make update      glide update packages recursively"
	@echo "make install     glide install packages in glide.lock"
	@echo "make pgo         glide get pgo"
	@echo "make init        init project"
