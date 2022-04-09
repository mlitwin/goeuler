
.PHONY: docs

docs:
	mkdir -p docs
	(cd arith; gomarkdoc > ../docs/arith.md)
	(cd algo; gomarkdoc > ../docs/algo.md)