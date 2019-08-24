build:
	cd cmd/tprof && go build -o ~/bin/tprof

build-nuxt:
	cd web && yarn build 
copy-assets:
	cd web && cp -r dist ../assets
clean:
	rm -r assets
generate-template:
	sed -i 's/window.data=/window.foo=/g' assets/_nuxt/*.js
	sed -i 's/<\/body>/<script>window.data={{data}}<\/script>\n<\/body>/g' assets/index.html

build-ui:
	clean
	build-nuxt 
	generate-template

.PHONY:
	build
