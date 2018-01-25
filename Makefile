svgtpl: 
	# go-bindata -pkg svg -ignore svg/*.go -o svg/gobindata.go svg
	go-bindata -pkg svg -o svg/gobindata.go svg