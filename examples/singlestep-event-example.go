package main

import (
  "fmt"
  "libvmi-go/libvmi"
  "os"
  "os/signal"
  "syscall"

)

func single_step_callback(vmi libvmi.Libvmi, event libvmi.Libvmi_Event){
  fmt.Println("Doing stuff")
}

func main(){

  if len(os.Args) != 2 {
    fmt.Println("Usage: <vmname>")
    return
  }

  vmName := os.Args[1]

  vmi,status := libvmi.Init(libvmi.VMI_XEN | libvmi.VMI_INIT_PARTIAL | libvmi.VMI_INIT_EVENTS, vmName)

  /*for a clean exit*/
  signalchannel := make(chan os.Signal, 1)
  signal.Notify(signalchannel, syscall.SIGINT, syscall.SIGTERM)
  go func(){
    <-signalchannel
    fmt.Println("Forced exit. Releasing resource")
    if vmi.IsInitialized() == true {
      vmi.Resume_vm()
      vmi.Destroy()
      os.Exit(3)
    }else{
      os.Exit(3)
    }
  }()

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
  single_event.EnableSingleStepEvent = true

  libvmi.Vmi_register_event(vmi,single_event)

  for {
    fmt.Println("Waiting for events...")
    libvmi.Vmi_events_listen(vmi,500)
  }

}
