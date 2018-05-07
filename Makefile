.PHONY: all

VERSION = 100

deploy: 
	gcloud app deploy --project moshimo-works --version ${VERSION}
