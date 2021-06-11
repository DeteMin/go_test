.PHONY: repair

repair:
	protoc \
	-I . \
	--micro_out=. \
	--go_out=. \
	--validate_out="lang=go:." \
	./repair/repair.proto
