package main

import (
  "fmt"
  "libvmi-go/libvmi"
  "os"
)

var cr3 uint64

func main(){

  if len(os.Args) != 2 {
    fmt.Println("Usage: <vmname>")
    return
  }

  vmName := os.Args[1]

  var lstar, phys_lstar, cstar, phys_cstar, sysenter_ip, phys_sysenter_ip uint64
  var ia32_sysenter_target, phys_ia32_sysenter_target, vsyscall, phys_vsyscall uint64

  var pd uint32

  vmi,status := libvmi.Init(libvmi.VMI_XEN | libvmi.VMI_AUTO | libvmi.VMI_INIT_COMPLETE, vmName)

  if status == libvmi.VMI_SUCCESS{
    //destroy the libvmi instance when the main function returns
    defer vmi.Destroy()
  }else{
    fmt.Println("Libvmi was not initialized properly")
    return
  }

  //Get the cr3 for this process


}
