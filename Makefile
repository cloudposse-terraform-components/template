# TODO DELETE AFTER MERGING PR https://github.com/cloudposse-terraform-components/.github/pull/2$
export README_TEMPLATE_REPO_REF ?= "DEV-1548/component-requirements"
export README_DEPS ?= docs/targets.md docs/terraform.md

-include $(shell curl -sSL -o .build-harness "https://cloudposse.tools/build-harness"; echo .build-harness)

all: init readme

readme: readme/build

test::
	@echo "🚀 Starting tests..."
	./test/run.sh
	@echo "✅ All tests passed."
