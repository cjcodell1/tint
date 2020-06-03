* `tint-linux` is compiled for the Linux OS with amd64 architecture.
* `tint-windows` is compiled for the Windows OS with amd64 architecture.
* `tint-mac` is compiled for the Mac OS with amd64 architecture.

Once you have the compiled program, you can rename the file to `tint` or whatever else you want to call this program.

If you have a different OS or architecture, I suggest using `go build` to compile this program.
`$ GOOS=target_os GOARCH=target_arch go build`
