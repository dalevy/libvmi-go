package main

import (
  "fmt"
  "libvmi-go/libvmi"
  "os"
)

func single_step_callback(vmi libvmi.Libvmi, event libvmi.Libvmi_Event){

}

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

  var single_event libvmi.Libvmi_Event
  single_event.Callback = single_step_callback
  single_event.Version = libvmi.VMI_EVENTS_VERSION
  single_event.Type = libvmi.VMI_EVENT_SINGLESTEP

  libvmi.Vmi_register_event(vmi,single_event)

  for {
    fmt.Println("Waiting for events...")
    libvmi.Vmi_events_listen(vmi,500)
  }

}
