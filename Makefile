# TODO, figure out remote template
# export CLOUDPOSSE_TERRAFORM_COMPONENTS_PATH ?= "https://raw.githubusercontent.com/cloudposse-terraform-components/.github/main"
# export README_TEMPLATE_FILE ?= $(CLOUDPOSSE_TERRAFORM_COMPONENTS_PATH)/templates/README.md.gotmpl
export README_DEPS ?= src/README.md

-include $(shell curl -sSL -o .build-harness "https://cloudposse.tools/build-harness"; echo .build-harness)

all: init readme

test::
	@echo "🚀 Starting tests..."
	./test/run.sh
	@echo "✅ All tests passed."
