SHELL=/bin/bash
PWD=$(shell pwd)

define run
	docker run -it --rm -v $(PWD):/usr/src/app --workdir /usr/src/app golang:1.18-alpine $(1)
endef

.PHONY: shell
shell:
	$(call run,sh)