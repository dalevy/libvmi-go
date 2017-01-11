package main

import (
  "fmt"
  "libvmi-go/libvmi"
  "os"
)

var mm_enabled bool

func main(){

  if len(os.Args) != 2 {
    fmt.Println("Usage: <vmname>")
    return
  }

  vmName := os.Args[1]

  vmi,status := libvmi.Init(libvmi.VMI_XEN | libvmi.VMI_AUTO | libvmi.VMI_INIT_COMPLETE, vmName)

  if status == libvmi.VMI_SUCCESS{
    //destroy the libvmi instance when the main function returns
    defer vmi.Destroy()
  }else{
    fmt.Println("Libvmi was not initialized properly")
    return
  }

  


}
