package libvmi

// #cgo LDFLAGS: -lvmi
// #include <sys/mman.h>
// #include <errno.h>
// #include <inttypes.h>
// #include <stdlib.h>
// #include <libvmi/libvmi.h>
// #include <libvmi/events.h>
// #include "go_libvmi.h"
//
// void SETUP_REG_EVENT_WRAPPER(vmi_event_t *event, unsigned long long register, unsigned char access, unsigned long long equal)
// {
//    SETUP_REG_EVENT(event,register,access,equal);
// }
import "C"

import (
  "fmt"
  "unsafe"
  "errors"
)

const (
  VMI_INIT_COMPLETE = C.VMI_INIT_COMPLETE
  VMI_INIT_PARTIAL = C.VMI_INIT_PARTIAL
  VMI_AUTO = C.VMI_AUTO
  VMI_EVENTS_VERSION = 0x00000002
  VMI_FAILURE = C.VMI_FAILURE
  VMI_SUCCESS = C.VMI_SUCCESS
  VMI_FILE = C.VMI_FILE
  VMI_XEN = C.VMI_XEN
  VMI_OS_UNKNOWN = 0
  VMI_OS_LINUX = 1
  VMI_OS_WINDOWS = 2

)

const (
  CR3 = C.CR3
)

type Reg_t struct{
  val uint64
}

type LibVMI_Event struct{
  version uint32
  slat_id uint16
  data uintptr
}

type LibVMI struct{
  vmi C.vmi_instance_t
  initialized bool
}

func SETUP_INTERRUPT_EVENT(event LibVMI_Event,reinject uint){
  var interrupt_event C.vmi_event_t
  interrupt_event.version = VMI_EVENTS_VERSION
}

//get a new libvmi instance for a given vm
func Init(flags C.uint32_t,vmName string)(LibVMI, int){
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

func Init_custom(flags C.uint32_t, buf uintptr)(LibVMI, int){
    var vmi C.vmi_instance_t
    var status int

    if (C.vmi_init_custom(&vmi,flags,buf) == C.VMI_FAILURE) {
         fmt.Println("Failed to init LibVMI library")
         status = VMI_FAILURE
     }else{
       fmt.Println("Libvmi initialized successfully\n")
         status = VMI_SUCCESS
     }

    obj := LibVMI{vmi: vmi}

    return obj,status
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

func Init_complete_custom(flags C.uint32_t, buf uintptr)(LibVMI, int){
    var vmi C.vmi_instance_t
    var status int

    if (C.vmi_init_complete_custom(&vmi,flags,buf) == C.VMI_FAILURE) {
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

func (i *LibVMI) Read_addr_ksym(symbol string)(uint64,int){
  var status int
  value := C.get_addr_t(0)
  if C.vmi_read_addr_ksym(i.vmi,C.CString(symbol),&value) == C.VMI_FAILURE{
    status = VMI_FAILURE
  }else{
    status = VMI_SUCCESS
  }

  return uint64(C.convert_addr_t(value)),status
}

func (i *LibVMI) Translate_ksym2v(symbol string)(uint64, error){
  address := C.vmi_translate_ksym2v(i.vmi,C.CString(symbol))
  if address == 0{
    return 0, errors.New("vmi kernel symbol to virtual address translation error")
  }
  return uint64(address),nil
}

/*
* Reads a null terminated string from memory, starting at the given virtual address. T
* The returned value is a Go string and does not need to be freed by the caller.
*/
func (i *LibVMI) Read_str_va(vaddr uint64,pid int32)(string,error){
  var value string
  cstring :=C.vmi_read_str_va(i.vmi,C.get_addr_t(C.ulonglong(vaddr)),C.get_vmi_pid_t(C.int(pid)))
  value = C.GoString(cstring)

  if cstring == nil{
    return "", errors.New("vmi kernel symbol to virtual address translation error")
  }else{
      defer C.free(unsafe.Pointer(cstring))
  }
  value = C.GoString(cstring)
  return value, nil

}

/*
* Reads 8 bits from memory, given a virtual address.
*/
func (i *LibVMI) Read_8_va(vaddr uint64, pid int32)(uint8, int){
  var value C.uint8_t
  var status int
  if C.vmi_read_8_va(i.vmi,C.get_addr_t(C.ulonglong(vaddr)),C.get_vmi_pid_t(C.int(pid)),&value) == C.VMI_FAILURE{
    status = VMI_FAILURE
  }else{
    status = VMI_SUCCESS
  }
  return uint8(value),status
}

/*
* Reads 16 bits from memory, given a virtual address.
*/
func (i *LibVMI) Read_16_va(vaddr uint64, pid int32)(uint16, int){
  var value C.uint16_t
  var status int
  if C.vmi_read_16_va(i.vmi,C.get_addr_t(C.ulonglong(vaddr)),C.get_vmi_pid_t(C.int(pid)),&value) == C.VMI_FAILURE{
    status = VMI_FAILURE
  }else{
    status = VMI_SUCCESS
  }
  return uint16(value),status
}

/*
* Reads 32 bits from memory, given a virtual address.
*/
func (i *LibVMI) Read_32_va(vaddr uint64, pid int32)(uint32, int){
  var value C.uint32_t
  var status int
  if C.vmi_read_32_va(i.vmi,C.get_addr_t(C.ulonglong(vaddr)),C.get_vmi_pid_t(C.int(pid)),&value) == C.VMI_FAILURE{
    status = VMI_FAILURE
  }else{
    status = VMI_SUCCESS
  }
  return uint32(value),status
}

/*
* Reads 64 bits from memory, given a virtual address.
*/
func (i *LibVMI) Read_64_va(vaddr uint64, pid int32)(uint64, int){
  var value C.uint64_t
  var status int
  if C.vmi_read_64_va(i.vmi,C.get_addr_t(C.ulonglong(vaddr)),C.get_vmi_pid_t(C.int(pid)),&value) == C.VMI_FAILURE{
    status = VMI_FAILURE
  }else{
    status = VMI_SUCCESS
  }
  return uint64(value),status
}


/*
* Reads an address from memory, given a virtual address. The number of bytes read is 8 for 64-bit systems and 4 for 32-bit systems.
*/
func (i *LibVMI) Read_addr_va(vaddr uint64, pid int32)(uint64,int){
  var status int
  value := C.get_addr_t(0)
  if C.vmi_read_addr_va(i.vmi,C.get_addr_t(C.ulonglong(vaddr)),C.get_vmi_pid_t(C.int(pid)),&value) == C.VMI_FAILURE{
    status = VMI_FAILURE
  }else{
    status = VMI_SUCCESS
  }

  return uint64(C.convert_addr_t(value)), status
}

//TODO: fix uintptr buffer, does it work?
func (i *LibVMI) Read_va(vaddr uint64, pid int32, buf uintptr, count uint){

  C.vmi_read_va(i.vmi,C.get_addr_t(C.ulonglong(vaddr)),C.get_vmi_pid_t(C.int(pid)),buf,C.get_size_t(C.uint(count)))

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

func (i *LibVMI) Resume_vm()int{
  var status int
  if C.vmi_resume_vm(i.vmi) == C.VMI_FAILURE{
    status = VMI_FAILURE
  }else{
    status = VMI_SUCCESS
  }
  return status
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
