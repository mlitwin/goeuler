
.PHONY: docs

docs:
	mkdir -p docs
	gomarkdoc ./arith > ./docs/arith.md
	gomarkdoc ./algo > ./docs/algo.md
	gomarkdoc ./textutil > ./docs/textutil.md