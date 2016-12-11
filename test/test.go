package main

import (
  "fmt"
  "libvmi-go/libvmi"
)

func processlist(vmName string){
  vmi,status := libvmi.Init(libvmi.VMI_AUTO | libvmi.VMI_INIT_COMPLETE, vmName)
  var tasks_offset, pid_offset, list_head, next_list_entry, current_process, name_offset uint64
  var pid uint32
  var err error

  if status == libvmi.VMI_SUCCESS{
    defer vmi.Destroy()
  }else{
    fmt.Println("Libvmi was not initialized properly")
    return
  }

  /*init the offset values */
  if libvmi.VMI_OS_LINUX == vmi.Get_ostype(){
    tasks_offset = vmi.Get_offset("linux_tasks")
    name_offset = vmi.Get_offset("linux_name")
    pid_offset = vmi.Get_offset("linux_pid")
  }else if libvmi.VMI_OS_WINDOWS == vmi.Get_ostype() {
    tasks_offset = vmi.Get_offset("win_tasks")
    name_offset = vmi.Get_offset("win_pname")
    pid_offset = vmi.Get_offset("win_pid")
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
  }else{
    defer vmi.Resume_vm()
  }

  /* demonstrate name and id accessors */
  name2 := vmi.Get_name()

  if libvmi.VMI_FILE != vmi.Get_access_mode() {
        id := vmi.Get_vmid()
        fmt.Println("Process listing for VM", name2, id)
  }else {
        fmt.Println("Process listing for file", name2)
  }

  /*get the head of the list */
  if libvmi.VMI_OS_LINUX == vmi.Get_ostype(){
    list_head,err = vmi.Translate_ksym2v("init_task")
    list_head = list_head + tasks_offset
    if err != nil {
      fmt.Println(err)
    }
  }else if libvmi.VMI_OS_WINDOWS == vmi.Get_ostype() {
    list_head, status = vmi.Read_addr_ksym("PsActiveProcessHead")
    if libvmi.VMI_FAILURE == status {
      fmt.Println("Failed to find PsActiveProcessHead")
      return
    }
  }

  next_list_entry = list_head

  for{

    current_process = next_list_entry - tasks_offset

    /* Note: the task_struct that we are looking at has a lot of
         * information.  However, the process name and id are burried
         * nice and deep.  Instead of doing something sane like mapping
         * this data to a task_struct, I'm just jumping to the location
         * with the info that I want.  This helps to make the example
         * code cleaner, if not more fragile.  In a real app, you'd
         * want to do this a little more robust :-)  See
         * include/linux/sched.h for mode details */

         /* NOTE: _EPROCESS.UniqueProcessId is a really VOID*, but is never > 32 bits,
        * so this is safe enough for x64 Windows for example purposes */
        pid, status = vmi.Read_32_va(current_process+pid_offset,0)

        if status == libvmi.VMI_FAILURE{
          fmt.Println("Failed to find pid")
          return
        }

        procname, err := vmi.Read_str_va(current_process+name_offset,0)

      if err != nil{
        fmt.Println("Failed to find procname")
        return
      }

      fmt.Print("[ ",pid)
      fmt.Print("]  ",procname)
      fmt.Printf(" -- struct addr:%x\n",current_process)

      next_list_entry, status = vmi.Read_addr_va(next_list_entry,0)

      if status == libvmi.VMI_FAILURE{
        fmt.Println("Failed to read the next pointer in the loop at ", next_list_entry)
        return
      }


      if next_list_entry == list_head {
        break
      }

  }

}

func main(){

  processlist("Buntu5-1")
}
