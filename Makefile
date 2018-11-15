##########################################################################################
## Make config variables

MAKEFLAGS += --warn-undefined-variables
SHELL := /usr/bin/env bash
.DELETE_ON_ERROR:

##########################################################################################
## App level vars

include Makefile.app.vars

##########################################################################################
## Functions

rfind = $(shell find '$(1)' -path '$(2)' -type f)
rfind_exclude = $(shell find '$(1)' -path '$(2)' ! -name '$(3)')
uname_s = $(shell uname -s)
get_os = $(if $(findstring Darwin,$(call uname_s)),MAC,LINUX)
sha1 = $(shell echo $(1) | sha1sum | cut -d ' ' -f1)

##########################################################################################
## Variables to use in pattern subs

empty :=
space := $(empty) $(empty)
dash := $(empty)-$(empty)
comma := $(empty),$(empty)
at_the_rate := $(empty)@$(empty)
escaped_quotes := $(empty)\'$(empty)
escaped_semicolon := $(empty)\;$(empty)
pipe := $(empty)|$(empty)

##########################################################################################
## Variables

DEBUG := off # Set this to on only for debug purposes during development
AT_off := $(at_the_rate)
AT_on := $(empty)
AT := $(AT_$(DEBUG))

OS := $(call get_os)
ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
REPO_NAME := $(shell basename $(ROOT_DIR))
THIS_FILE := $(lastword $(MAKEFILE_LIST))
WIP := 0
HOTFIX := 0
BINTRAY_API_URL="https://api.bintray.com"

ifdef DEPLOY_GITHUB_TOKEN
	GIT_REPO_URL := https://$(DEPLOY_GITHUB_TOKEN)@github.com/$(GITHUB_USERNAME)/$(REPO_NAME).git
else
	$(error "ERROR : Please export DEPLOY_GITHUB_TOKEN to your shell")
endif

USER := $(shell id -un)
PREVIOUS_RELEASE_TAG := $(shell git describe --tags --match '$(SEMVER_REGEX)' --abbrev=0 HEAD)
GIT_SHA := $(shell git rev-parse --short --verify HEAD)
BUILD_TIME := $(shell date --utc +"%Y-%m-%d-%H-%M-%S-UTC")
TMPDIR_FOR_BUILD :=$(shell mktemp -d "/tmp/$(REPO_NAME).XXXXXXX")
BUILDER_IMAGE_NAME := $(REPO_NAME)-builder
BUILDER_CONTAINER_NAME := $(BUILDER_IMAGE_NAME)-container
BUILDER_DATA_DIR ?= $(ROOT_DIR)/dist

# This is programatically set if a TAG is passed from the cli
IS_TAG_FROM_CLI := 0
ifdef TAG
	IS_TAG_FROM_CLI := 1
	WIP := 0
endif

APPLICATION_FILES := $(call rfind,cmd/,**/*) \
	$(ROOT_DIR)/Gopkg.toml \
	$(ROOT_DIR)/Gopkg.toml

DEPS_STATEFILE = .make/done_deps

# If a tag is auto generated then it is a new build and perform the required checks
# Otherwise a known tag is being passed, so look for existing build
ifeq ($(IS_TAG_FROM_CLI), 1)
	BUILD_DEPENDENT_TARGETS := checks_for_existing_build
else
	BUILD_DEPENDENT_TARGETS := checks_for_new_build
endif

##########################################################################################
## Vars for cli options

ifeq ($(WIP),1)
	HOTFIX := 0
	VERSION := $(empty)
	TAG := $(WIP_REGEX)-$(BUILD_TIME)
else
ifeq ($(HOTFIX),1)
	VERSION := $(empty)
	TAG := $(HOTFIX_REGEX)-$(BUILD_TIME)
else
	# If not a hotfix then valid values can be <patch|minor|major>
	VERSION := $(PATCH)
	TAG := $(shell /usr/bin/env bash $(ROOT_DIR)/bin/semver.sh $(VERSION) $(PREVIOUS_RELEASE_TAG))
endif
endif

##########################################################################################
## Public targets

.DEFAULT_GOAL := release
.PHONY : deps \
	help \
	build \
	release \
	clean \
	fullclean

# Installs all the dependencies for running the build process
deps : $(DEPS_STATEFILE)
	$(info [INFO] --- Install dependencies for running make targets)
ifeq ($(MAKECMDGOALS),deps)
	rm -rf $(TMPDIR_FOR_BUILD)
endif

# Builds the artifacts inside a docker container
build : $(BUILD_DEPENDENT_TARGETS) Dockerfile $(APPLICATION_FILES)
	$(AT)echo "Was $(TAG) passed from the cli? - $(IS_TAG_FROM_CLI)"
ifeq ($(IS_TAG_FROM_CLI), 0)
	$(info [INFO] --- Build and tag container from commit sha)
	$(AT)cd $(TMPDIR_FOR_BUILD) \
		&& git clone $(GIT_REPO_URL) \
		&& cd $(REPO_NAME) \
		&& git checkout $(GIT_SHA) \
		&& docker build \
			--no-cache \
			--build-arg GITHUB_USERNAME=$(GITHUB_USERNAME) \
			--build-arg GIT_TAG=$(TAG) \
			--build-arg BUILD_TIME=$(BUILD_TIME) \
			--build-arg GIT_REF="$(GIT_REPO_URL)#$(GIT_SHA)" \
			--build-arg BUILD_USER="$(USER)" \
			--build-arg REPO_NAME="$(REPO_NAME)" \
			--build-arg NON_ROOT_UID=$(NON_ROOT_UID) \
			--build-arg NON_ROOT_GID=$(NON_ROOT_GID) \
			--build-arg NON_ROOT_USER=default \
			-f Dockerfile \
			-t $(BUILDER_IMAGE_NAME):$(TAG) .
	$(info [INFO] --- Create annotated semver tag marking commit sha as a release candidate)
	$(AT)docker tag $(BUILDER_IMAGE_NAME):$(TAG) $(DOCKERHUB_USERNAME)/$(BUILDER_IMAGE_NAME):$(TAG)
	$(AT)git tag $(TAG) -am "Version:$(TAG),User:$(USER),Time:$(BUILD_TIME)"
else
	$(info [INFO] --- Skipping building and tagging container from commit sha)
	$(info [INFO] --- Skipping tagging git commit sha)
endif
	$(AT)rm -rf $(TMPDIR_FOR_BUILD)
ifeq ($(MAKECMDGOALS),build)
	rm -rf $(TMPDIR_FOR_BUILD)
endif

# Uploads the artifact built from the code to bintray
# Pushes the docker image to dockerhub
# Pushes the git tag
release: build
	$(info [INFO] --- Create release candidate)
	$(AT)mkdir -p $(BUILDER_DATA_DIR)
	$(AT)docker run -u $(NON_ROOT_UID):$(NON_ROOT_GID) --name $(BUILDER_CONTAINER_NAME) $(DOCKERHUB_USERNAME)/$(BUILDER_IMAGE_NAME):$(TAG) \
		&& docker cp $(BUILDER_CONTAINER_NAME):/var/data/build/$(REPO_NAME).tar.gz $(BUILDER_DATA_DIR)/$(REPO_NAME)-$(TAG).tar.gz \
		&& curl -X POST -u$(BINTRAY_USERNAME):$(BINTRAY_API_KEY) -d '{"name": "$(TAG)", "vcs_tag": "$(TAG)", "released": "$(BUILD_TIME)"}' $(BINTRAY_API_URL)/packages/$(BINTRAY_USERNAME)/$(BINTRAY_REPO_NAME)/$(REPO_NAME)/versions \
		&& curl -T $(BUILDER_DATA_DIR)/$(REPO_NAME)-$(TAG).tar.gz -u$(BINTRAY_USERNAME):$(BINTRAY_API_KEY) -d '{"discard": "false"}' $(BINTRAY_API_URL)/content/$(BINTRAY_USERNAME)/$(BINTRAY_REPO_NAME)/$(REPO_NAME)/$(TAG)/$(REPO_NAME)-$(TAG).tar.gz?publish=1
	$(AT)docker stop $(BUILDER_CONTAINER_NAME)
	$(AT)docker rm $(BUILDER_CONTAINER_NAME)
	$(AT)docker push $(DOCKERHUB_USERNAME)/$(BUILDER_IMAGE_NAME):$(TAG)
	$(AT)git push origin $(TAG)
ifeq ($(MAKECMDGOALS),release)
	rm -rf $(TMPDIR_FOR_BUILD)
endif

# Cleans unused dirs/images/containers
clean :
	$(info [INFO] --- Clean stopped containers, intermediate images and build artifacts)
ifneq ($(shell docker ps -a -q --filter "name=$(BUILDER_CONTAINER_NAME)"),)
	$(AT)docker stop $(shell docker ps -a -q --filter "name=$(BUILDER_CONTAINER_NAME)")
	$(AT)docker rm $(shell docker ps -a -q --filter "name=$(BUILDER_CONTAINER_NAME)")
endif
ifneq ($(shell docker images | egrep -i "($(BUILDER_IMAGE_NAME)|none)" | awk '{print $$3}'),)
	$(AT)docker images | egrep -i "($(BUILDER_IMAGE_NAME)|none)" | awk '{print $$3}' | xargs docker rmi -f || echo
endif
	$(AT)rm -rf $(ROOT_DIR)/dist
ifeq ($(MAKECMDGOALS),clean)
	rm -rf $(TMPDIR_FOR_BUILD)
endif

fullclean : clean
	$(AT)rm -rf .make
ifeq ($(MAKECMDGOALS),fullclean)
	rm -rf $(TMPDIR_FOR_BUILD)
endif

help :
	$(AT)echo make deps - This install dependancies to run the application
	$(AT)echo make build - This builds the golang binaries within a docker container
	$(AT)echo make release - This pushes the golang binary artifacts to bintray/s3/github/artifactory etc
	$(AT)echo make clean - This remove image caches, stale containers etc
	$(AT)echo make help - This is the help menu
ifeq ($(MAKECMDGOALS),help)
	rm -rf $(TMPDIR_FOR_BUILD)
endif

##########################################################################################
## Plumbing
.PHONY : check_deps \
	checks_for_env_vars \
	checks_bintray_repo \
	checks_or_makes_bintray_package \
	check_working_dir_status \
	check_branch \
	check_no_existing_tag_on_commit \
	check_no_existing_tag_on_remote \
	check_no_existing_tag_locally \
	checks_for_new_build \
	check_existing_docker_tag \
	checks_for_existing_build \
	checks_logged_into_dockerhub

# Checks for dependencies
# Check for docker
# Check for git
# Check for xargs
# Check for jq
check_deps:
	$(info [INFO] --- Checks that you have the required dependencies for building and releasing)
	$(AT)command -v docker \
	&& command -v git \
	&& command -v jq \
	&& command -v xargs

# Checks that required env vars are present
checks_for_env_vars :
	$(info [INFO] --- Checks that required env vars are present)
	$(AT)test ! -z "$$GITHUB_USERNAME" \
	&& test ! -z "$$DOCKERHUB_USERNAME" \
	&& test ! -z "$$BINTRAY_USERNAME" \
	&& test ! -z "$$BINTRAY_REPO_NAME" \
	&& test ! -z "$$BINTRAY_API_KEY"

# Checks that bintray repo is present
checks_bintray_repo :
	$(info [INFO] --- Checks that bintray repo is present)
	$(AT)! curl -u$(BINTRAY_USERNAME):$(BINTRAY_API_KEY) $(BINTRAY_API_URL)/repos/$(BINTRAY_USERNAME)/$(BINTRAY_REPO_NAME) | grep -i 'was not found'

# Checks that bintray package is present or create one
checks_or_makes_bintray_package :
	$(info [INFO] --- Checks that bintray repo is present)
	$(AT)! curl -u$(BINTRAY_USERNAME):$(BINTRAY_API_KEY) $(BINTRAY_API_URL)/packages/$(BINTRAY_USERNAME)/$(BINTRAY_REPO_NAME)/$(REPO_NAME) | grep -i 'was not found' \
	|| curl -X POST -u$(BINTRAY_USERNAME):$(BINTRAY_API_KEY) -H "Content-Type: application/json" -H "Accept: application/json" -d '{"name": "$(REPO_NAME)", "licenses": ["MIT"], "vcs_url": "https://github.com/$(GITHUB_USERNAME)/$(REPO_NAME).git"}' $(BINTRAY_API_URL)/packages/$(BINTRAY_USERNAME)/$(BINTRAY_REPO_NAME)

# Checks is HEAD is detached
# Checks if there are uncommited changes
# Checks if local is upto date with remote
# Checks if you have unpushed tags or not
check_working_dir_status :
	$(info [INFO] --- Checks the status of the working directory)
	$(AT)git symbolic-ref --short --quiet HEAD \
	&& git diff-index --quiet HEAD -- \
	&& git status -uno | grep -E '(up-to-date|up\ to\ date)' \
	&& ! git show-ref --tags | grep -v -F "$$(git ls-remote --tags $(GIT_REPO_URL) | grep -v '\^{}' | cut -f2)"

# Checks if you have the WIP flag enabled then you are not on stable branches like master or develop
# Checks that if you have a HOTFIX flag enabled then you are on a hotfix branch
# Checks that you are on master or develop for all other cases
check_branch :
	$(info [INFO] --- Checks that you are on the correct branch)
ifeq ($(WIP),1)
	$(AT)! git rev-parse --symbolic-full-name --abbrev-ref HEAD | egrep -e '($(subst $(space),$(pipe),$(STABLE_BRANCHES)))'
else
ifeq ($(HOTFIX),1)
	$(AT)git rev-parse --symbolic-full-name --abbrev-ref HEAD | grep '$(HOTFIX_REGEX)'
else
	$(AT)git rev-parse --symbolic-full-name --abbrev-ref HEAD | egrep -e '($(subst $(space),$(pipe),$(STABLE_BRANCHES)))'
endif
endif

# Checks that the commit you are trying to build from locally is not previously tagged already from a previous build
check_no_existing_tag_on_commit :
	$(info [INFO] --- Checks that the commit is not previously tagged)
ifeq ($(WIP),1)
	$(AT)! git describe --exact-match --tags $$(git log -n1 --pretty='%h')
else
ifeq ($(HOTFIX),1)
	$(AT)! git describe --exact-match --tags $$(git log -n1 --pretty='%h') | grep '$(HOTFIX_REGEX)'
else
	$(AT)! git describe --exact-match --tags $$(git log -n1 --pretty='%h') | grep 'v[0-9]\{1,5\}\.[0-9]\{1,5\}\.[0-9]\{1,5\}'
endif
endif

# Check for git tag on remote
# Check for docker tag on docker repo
check_no_existing_tag_on_remote :
	echo '$(VERSION)'
	echo '$(TAG)'
	$(info [INFO] --- Checks that the new tag does not exist on git remote or on remote docker repo)
	$(AT)! git ls-remote --tags $(GIT_REPO_URL) | grep '$(TAG)'

# Check for git tag locally
# Check for docker tag locally
check_no_existing_tag_locally :
	echo '$(VERSION)'
	echo '$(TAG)'
	$(info [INFO] --- Checks that the new tag does not exist on git locally or on docker client locally)
	$(AT)! git tag | grep '$(TAG)' \
	&& ! docker images | grep -i "$(BUILDER_IMAGE_NAME)" | grep "$(TAG)"

# Checks to pass for the new build
# Checks that all dependencies required to build the release exist
# Check if working dir is in a clean state
# Check if branch is correct based on the WIP and HOTFIX flag
# Check if there are any existing tags already for the commit you are trying to build from
# Check new tag is not already on remote on git
# Check new tag is not already on docker repo
checks_for_new_build : check_deps check_working_dir_status check_branch check_no_existing_tag_on_commit check_no_existing_tag_on_remote check_no_existing_tag_locally checks_for_env_vars checks_bintray_repo checks_or_makes_bintray_package checks_logged_into_dockerhub

# Check for docker tag locally
# Check for docker tag on remote
check_existing_docker_tag :
	echo '$(VERSION)'
	echo '$(TAG)'
	$(info [INFO] --- Checks that the tag passed from the cli exists on docker client locally or on docker repo on remote)
	$(AT)docker images | grep -i "$(BUILDER_IMAGE_NAME)" | grep "$(TAG)"

# Check for git tag locally
# Check for git tag on remote
check_existing_git_tag :
	echo '$(VERSION)'
	echo '$(TAG)'
	$(info [INFO] --- Checks that the tag passed from the cli exists on git locally or on git repo on remote)
	$(AT)git ls-remote --tags $(GIT_REPO_URL) | grep '$(TAG)' \
	|| git tag | grep '$(TAG)'

checks_for_existing_build : check_deps check_existing_docker_tag check_existing_git_tag checks_for_env_vars checks_logged_into_dockerhub

# Checks that user is logged into dockerhub
checks_logged_into_dockerhub :
	$(info [INFO] --- Checks that user is logged into dockerhub)
	$(AT){ test -d "$$HOME/.docker" && test ! -z "$$(cat "$$HOME/.docker/config.json" | jq -r '.auths | ."https://index.docker.io/v1/" | .auth')"; } \
	|| { test -f "$$HOME/.dockercfg" && test ! -z "$$(cat "$$HOME/.dockercfg" | jq -r '.auths | ."https://index.docker.io/v1/" | .auth')"; }

$(DEPS_STATEFILE) :
	$(info [INFO] --- Installs the dependencies to run the make targets)
	$(AT)mkdir -p .make
	$(AT)touch $(DEPS_STATEFILE)

