# libvmi-go
Go bindings for libvmi. Please note, this is a work in progress and the bindings are not complete at this time but pull requests are still welcome.

# Dependencies
The LibVMI virtual machine introspection library is required to use these bindings. Please follow the instructions at https://github.com/libvmi/libvmi to install the LibVMI library.

# Examples
The examples can be run by either typing go run [params] or go build and then ./[executable]
For instance...
 #go run processlist.go UbuntuTestVM
Or ...
 #go build processlist.go
 #./processlist UbuntuTestVM
