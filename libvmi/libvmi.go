package libvmi

// #cgo LDFLAGS: -lvmi
// #include <libvmi/libvmi.h>
// #include <sys/mman.h>
// #include <errno.h>
// #include <inttypes.h>
// #include <stdlib.h>
//
//addr_t
//get_addr_t(unsigned long long val)
//{
//  addr_t address = val;
//  return address;
//
//}
//
//vmi_pid_t
//get_vmi_pid_t(int val)
//{
//  vmi_pid_t pid = val;
//  return pid;
//}
//size_t
//get_size_t(unsigned int val)
//{
//  size_t size = val;
//  return size;
//}
//
import "C"

import (
  "fmt"
  "unsafe"
)

const (
  VMI_INIT_COMPLETE = C.VMI_INIT_COMPLETE
  VMI_INIT_PARTIAL = C.VMI_INIT_PARTIAL
  VMI_AUTO = C.VMI_AUTO
  VMI_FAILURE = C.VMI_FAILURE
  VMI_SUCCESS = C.VMI_SUCCESS
  VMI_FILE = C.VMI_FILE
  VMI_OS_UNKNOWN = 0
  VMI_OS_LINUX = 1
  VMI_OS_WINDOWS = 2

)


type LibVMI struct{
  vmi C.vmi_instance_t
  initialized bool
}


func (i *LibVMI) Get_vmid()uint64{
  id := C.vmi_get_vmid(i.vmi)
  return uint64(id)
}

func (i *LibVMI) Get_access_mode() uint32 {
  mode := C.vmi_get_access_mode(i.vmi)
  return uint32(mode)
}

func (i *LibVMI) Get_name()string{
  name := C.vmi_get_name(i.vmi)
  defer C.free(unsafe.Pointer(name))
  return C.GoString(name)
}

func (i *LibVMI) Pause_vm()int{
  var status int
  if C.vmi_pause_vm(i.vmi) == C.VMI_FAILURE{
    status = VMI_FAILURE
  }else{
    status = VMI_SUCCESS
  }

  return status
}

func (i *LibVMI) Get_offset(offset_name string)uint64{
  offset := C.vmi_get_offset(i.vmi,C.CString(offset_name))
  return uint64(offset)
}

func (i *LibVMI) Get_ostype()int{

 os := C.vmi_get_ostype(i.vmi)
 if os == C.VMI_OS_LINUX{
   return VMI_OS_LINUX
 }else if os == C.VMI_OS_WINDOWS{
   return VMI_OS_WINDOWS
 }else{
   return VMI_OS_UNKNOWN
 }

}

//TODO: fix uintptr
func (i *LibVMI) Read_va(addr uint64, pid int32, buf uintptr, count uint){

  C.vmi_read_va(i.vmi,C.get_addr_t(C.ulonglong(addr)),C.get_vmi_pid_t(C.int(pid)),buf,C.get_size_t(C.uint(count)))

}

func Init_complete(config string)(LibVMI, int){
    var vmi C.vmi_instance_t
    var status int

    if (C.vmi_init_complete(&vmi,C.CString(config)) == C.VMI_FAILURE) {
         fmt.Println("Failed to init LibVMI library")
         status = VMI_FAILURE
     }else{
       fmt.Println("Libvmi initialized successfully\n")
         status = VMI_SUCCESS
     }

    obj := LibVMI{vmi: vmi}

    return obj,status
}

//get a new libvmi instance for a given vm
func Init(vmName string, flags C.uint32_t)(LibVMI, int){
  fmt.Println("Creating new vmi instance for "+vmName)

    var vmi C.vmi_instance_t
    var status int

    if (C.vmi_init(&vmi, flags,C.CString(vmName)) == C.VMI_FAILURE) {
         fmt.Println("Failed to init LibVMI library")
         status = VMI_FAILURE
     }else{
       fmt.Println("Libvmi initialized successfully\n")
         status = VMI_SUCCESS
     }

    obj := LibVMI{vmi: vmi}

    return obj,status
}

//Destroy this libvmi instance
func(i *LibVMI) Destroy(){
  C.vmi_destroy(i.vmi)
}
