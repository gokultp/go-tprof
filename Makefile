build: build-ui gen-statik compile

compile: 
	cd cmd/tprof && go build -o ${GOPATH}/bin/tprof

build-nuxt:
	cd web && yarn build 

copy-assets:
	cd web && cp -r dist ../assets

clean:
	rm -rf assets ?


gen-template:
	sed -i 's/window.data=/window.foo=/g' assets/_nuxt/*.js
	sed -i 's/<\/body>/<script>window.data={{data}}<\/script>\n<\/body>/g' assets/index.html

build-ui: clean build-nuxt copy-assets	gen-template

gen-statik:
	cd internal && statik -src=../assets

config:
	cd web && yarn install
.PHONY:
	build
