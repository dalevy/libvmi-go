package main

import (
  "fmt"
  "libvmi-go/libvmi"
)

func processlist(vmName string){
  vmi,status := libvmi.Init(vmName, libvmi.VMI_AUTO | libvmi.VMI_INIT_COMPLETE)
  var tasks_offset, pid_offset, name_offset uint64

  if status == libvmi.VMI_SUCCESS{
    defer vmi.Destroy()
  }else{
    fmt.Println("Libvmi was not initialized properly")
    return
  }

  /*init the offset values */
  if libvmi.VMI_OS_LINUX == vmi.Get_ostype(){
    tasks_offset = vmi.Get_offset("linux_tasks")
    pid_offset = vmi.Get_offset("linux_name")
    name_offset = vmi.Get_offset("linux_pid")
  }else if libvmi.VMI_OS_LINUX == vmi.Get_ostype() {
    tasks_offset = vmi.Get_offset("win_tasks")
    pid_offset = vmi.Get_offset("win_pname")
    name_offset = vmi.Get_offset("win_pid")
  }

  if 0 == tasks_offset {
       fmt.Println("Failed to find win_tasks")
       return
   }
   if 0 == pid_offset {
       fmt.Println("Failed to find win_pid")
       return
   }
   if 0 == name_offset{
       fmt.Println("Failed to find win_pname")
       return
   }

   /* pause the vm for consistent memory access */
   if vmi.Pause_vm() != libvmi.VMI_SUCCESS {
       fmt.Println("Failed to pause VM")
       return
  }

  /* demonstrate name and id accessors */
  name2 := vmi.Get_name()

  if libvmi.VMI_FILE != vmi.Get_access_mode() {
        id := vmi.Get_vmid()
        fmt.Println("Process listing for VM", name2, id)
    }else {
        fmt.Println("Process listing for file", name2)
  }

  /*init the offset values */
  if libvmi.VMI_OS_LINUX == vmi.Get_ostype(){
    tasks_offset = vmi.Get_offset("linux_tasks")
    pid_offset = vmi.Get_offset("linux_name")
    name_offset = vmi.Get_offset("linux_pid")
  }else if libvmi.VMI_OS_LINUX == vmi.Get_ostype() {
    tasks_offset = vmi.Get_offset("win_tasks")
    pid_offset = vmi.Get_offset("win_pname")
    name_offset = vmi.Get_offset("win_pid")
  }

}

func main(){

  processlist("Buntu5-1")
}
