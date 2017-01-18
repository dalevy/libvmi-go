package libvmi

/*
#cgo LDFLAGS: -lvmi
#include <sys/mman.h>
#include <errno.h>
#include <inttypes.h>
#include <stdlib.h>
#include <libvmi/libvmi.h>
#include <libvmi/events.h>

//generic event handler that will be passed to C to call go
extern void go_libvmi_event_callback_proxy(vmi_instance_t vmi, vmi_event_t *event, uint64_t id);
extern  event_response_t generic_event_handler(vmi_instance_t vmi, vmi_event_t *event);

*/
import "C"

//export go_libvmi_event_callback_proxy
func go_libvmi_event_callback_proxy(vmi C.vmi_instance_t,event *C.vmi_event_t, id C.uint64_t){
  var event_wrapper Libvmi_Event
  event_wrapper = lookup(id)

  //set the event
  event_wrapper.setEvent(event)

  //another wrapper
  var instance Libvmi
  instance.vmi = vmi
  instance.initialized = true

  //retrieve the callback function
  callback_function := event_wrapper.Callback

  //finally - call it
  callback_function(instance,event_wrapper)

}
