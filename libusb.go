//-----------------------------------------------------------------------------
/*

Golang wrapper for libusb-1.0

*/
//-----------------------------------------------------------------------------

package libusb

/*
#cgo LDFLAGS: -lusb-1.0
#include <libusb-1.0/libusb.h>
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

//-----------------------------------------------------------------------------

const (
	LIBUSB_LOG_LEVEL_NONE    = C.LIBUSB_LOG_LEVEL_NONE
	LIBUSB_LOG_LEVEL_ERROR   = C.LIBUSB_LOG_LEVEL_ERROR
	LIBUSB_LOG_LEVEL_WARNING = C.LIBUSB_LOG_LEVEL_WARNING
	LIBUSB_LOG_LEVEL_INFO    = C.LIBUSB_LOG_LEVEL_INFO
	LIBUSB_LOG_LEVEL_DEBUG   = C.LIBUSB_LOG_LEVEL_DEBUG
)

const (
	LIBUSB_SUCCESS             = C.LIBUSB_SUCCESS
	LIBUSB_ERROR_IO            = C.LIBUSB_ERROR_IO
	LIBUSB_ERROR_INVALID_PARAM = C.LIBUSB_ERROR_INVALID_PARAM
	LIBUSB_ERROR_ACCESS        = C.LIBUSB_ERROR_ACCESS
	LIBUSB_ERROR_NO_DEVICE     = C.LIBUSB_ERROR_NO_DEVICE
	LIBUSB_ERROR_NOT_FOUND     = C.LIBUSB_ERROR_NOT_FOUND
	LIBUSB_ERROR_BUSY          = C.LIBUSB_ERROR_BUSY
	LIBUSB_ERROR_TIMEOUT       = C.LIBUSB_ERROR_TIMEOUT
	LIBUSB_ERROR_OVERFLOW      = C.LIBUSB_ERROR_OVERFLOW
	LIBUSB_ERROR_PIPE          = C.LIBUSB_ERROR_PIPE
	LIBUSB_ERROR_INTERRUPTED   = C.LIBUSB_ERROR_INTERRUPTED
	LIBUSB_ERROR_NO_MEM        = C.LIBUSB_ERROR_NO_MEM
	LIBUSB_ERROR_NOT_SUPPORTED = C.LIBUSB_ERROR_NOT_SUPPORTED
	LIBUSB_ERROR_OTHER         = C.LIBUSB_ERROR_OTHER
)

const (
	LIBUSB_ENDPOINT_IN  = C.LIBUSB_ENDPOINT_IN  // In: device-to-host.
	LIBUSB_ENDPOINT_OUT = C.LIBUSB_ENDPOINT_OUT // Out: host-to-device.
)

const LIBUSB_API_VERSION = C.LIBUSB_API_VERSION

//-----------------------------------------------------------------------------
// structures

/*

struct libusb_endpoint_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bEndpointAddress;
	uint8_t  bmAttributes;
	uint16_t wMaxPacketSize;
	uint8_t  bInterval;
	uint8_t  bRefresh;
	uint8_t  bSynchAddress;
	const unsigned char *extra;
	int extra_length;
};

struct libusb_interface_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bInterfaceNumber;
	uint8_t  bAlternateSetting;
	uint8_t  bNumEndpoints;
	uint8_t  bInterfaceClass;
	uint8_t  bInterfaceSubClass;
	uint8_t  bInterfaceProtocol;
	uint8_t  iInterface;
	const struct libusb_endpoint_descriptor *endpoint;
	const unsigned char *extra;
	int extra_length;
};

struct libusb_interface {
	const struct libusb_interface_descriptor *altsetting;
	int num_altsetting;
};

struct libusb_config_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint16_t wTotalLength;
	uint8_t  bNumInterfaces;
	uint8_t  bConfigurationValue;
	uint8_t  iConfiguration;
	uint8_t  bmAttributes;
	uint8_t  MaxPower;
	const struct libusb_interface *interface;
	const unsigned char *extra;
	int extra_length;
};

struct libusb_ss_endpoint_companion_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bMaxBurst;
	uint8_t  bmAttributes;
	uint16_t wBytesPerInterval;
};

struct libusb_bos_dev_capability_descriptor {
	uint8_t bLength;
	uint8_t bDescriptorType;
	uint8_t bDevCapabilityType;
	uint8_t dev_capability_data
};

struct libusb_bos_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint16_t wTotalLength;
	uint8_t  bNumDeviceCaps;
	struct libusb_bos_dev_capability_descriptor *dev_capability
};

struct libusb_usb_2_0_extension_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bDevCapabilityType;
	uint32_t  bmAttributes;
};

struct libusb_ss_usb_device_capability_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bDevCapabilityType;
	uint8_t  bmAttributes;
	uint16_t wSpeedSupported;
	uint8_t  bFunctionalitySupport;
	uint8_t  bU1DevExitLat;
	uint16_t bU2DevExitLat;
};

struct libusb_container_id_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bDevCapabilityType;
	uint8_t bReserved;
	uint8_t  ContainerID[16];
};

struct libusb_control_setup {
	uint8_t  bmRequestType;
	uint8_t  bRequest;
	uint16_t wValue;
	uint16_t wIndex;
	uint16_t wLength;
};

*/

type Device_Descriptor struct {
	BLength            uint8
	BDescriptorType    uint8
	BcdUSB             uint16
	BDeviceClass       uint8
	BDeviceSubClass    uint8
	BDeviceProtocol    uint8
	BMaxPacketSize0    uint8
	IdVendor           uint16
	IdProduct          uint16
	BcdDevice          uint16
	IManufacturer      uint8
	IProduct           uint8
	ISerialNumber      uint8
	BNumConfigurations uint8
}

type Version struct {
	Major    uint16
	Minor    uint16
	Micro    uint16
	Nano     uint16
	Rc       string
	Describe string
}

type Context *C.struct_libusb_context
type Device *C.struct_libusb_device
type Device_Handle *C.struct_libusb_device_handle
type Hotplug_Callback *C.struct_libusb_hotplug_callback

//-----------------------------------------------------------------------------
// errors

type libusb_error_t struct {
	name string
	code int
}

func (e *libusb_error_t) Error() string {
	return fmt.Sprintf("libusb_error: %s returned %d(%s)", e.name, e.code, Error_Name(e.code))
}

func libusb_error(name string, code int) error {
	return &libusb_error_t{
		name: name,
		code: code,
	}
}

//-----------------------------------------------------------------------------
// Library initialization/deinitialization

func Set_Debug(ctx Context, level int) {
	C.libusb_set_debug(ctx, C.int(level))
}

func Init(ctx *Context) error {
	rc := int(C.libusb_init((**C.struct_libusb_context)(ctx)))
	if rc != LIBUSB_SUCCESS {
		return libusb_error("libusb_init", rc)
	}
	return nil
}

func Exit(ctx Context) {
	C.libusb_exit(ctx)
}

//-----------------------------------------------------------------------------
// Device handling and enumeration

func Get_Device_List(ctx Context) ([]Device, error) {
	var hdl **C.struct_libusb_device
	rc := int(C.libusb_get_device_list(ctx, (***C.struct_libusb_device)(&hdl)))
	if rc < 0 {
		return nil, libusb_error("libusb_get_device_list", rc)
	}
	// turn the c array into a slice of device pointers
	var list []Device
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&list))
	hdr.Cap = rc
	hdr.Len = rc
	hdr.Data = uintptr(unsafe.Pointer(hdl))
	return list, nil
}

func Free_Device_List(list []Device, unref_devices int) {
	if list == nil {
		return
	}
	C.libusb_free_device_list((**C.struct_libusb_device)(&list[0]), C.int(unref_devices))
}

func Get_Bus_Number(dev Device) uint8 {
	return uint8(C.libusb_get_bus_number(dev))
}

func Get_Port_Number(dev Device) uint8 {
	return uint8(C.libusb_get_port_number(dev))
}

func Get_Port_Numbers(dev Device) ([]uint8, error) {
	ports := make([]uint8, 16)
	rc := int(C.libusb_get_port_numbers(dev, (*C.uint8_t)(&ports[0]), (C.int)(len(ports))))
	if rc < 0 {
		return nil, libusb_error("libusb_get_port_numbers", rc)
	}
	return ports[:rc], nil
}

func Get_Parent(dev Device) Device {
	return C.libusb_get_parent(dev)
}

func Get_Device_Address(dev Device) uint8 {
	return uint8(C.libusb_get_device_address(dev))
}

func Get_Device_Speed(dev Device) int {
	return int(C.libusb_get_device_speed(dev))
}

func Get_Max_Packet_Size(dev Device, endpoint uint8) int {
	return int(C.libusb_get_max_packet_size(dev, (C.uchar)(endpoint)))
}

func Get_Max_ISO_Packet_Size(dev Device, endpoint uint8) int {
	return int(C.libusb_get_max_iso_packet_size(dev, (C.uchar)(endpoint)))
}

func Ref_Device(dev Device) Device {
	return C.libusb_ref_device(dev)
}

func Unref_Device(dev Device) {
	C.libusb_unref_device(dev)
}

func Open(dev Device) (Device_Handle, error) {
	var hdl Device_Handle
	rc := int(C.libusb_open(dev, (**C.struct_libusb_device_handle)(&hdl)))
	if rc < 0 {
		return nil, libusb_error("libusb_open", rc)
	}
	return hdl, nil
}

func Open_Device_With_VID_PID(ctx Context, vendor_id uint16, product_id uint16) Device_Handle {
	return C.libusb_open_device_with_vid_pid(ctx, (C.uint16_t)(vendor_id), (C.uint16_t)(product_id))
}

func Close(hdl Device_Handle) {
	C.libusb_close(hdl)
}

func Get_Device(hdl Device_Handle) Device {
	return C.libusb_get_device(hdl)
}

func Get_Configuration(hdl Device_Handle) (int, error) {
	var config C.int
	rc := int(C.libusb_get_configuration(hdl, &config))
	if rc < 0 {
		return 0, libusb_error("libusb_get_configuration", rc)
	}
	return int(config), nil
}

func Set_Configuration(hdl Device_Handle, configuration int) error {
	rc := int(C.libusb_set_configuration(hdl, (C.int)(configuration)))
	if rc < 0 {
		return libusb_error("libusb_set_configuration", rc)
	}
	return nil
}

func Claim_Interface(hdl Device_Handle, interface_number int) error {
	rc := int(C.libusb_claim_interface(hdl, (C.int)(interface_number)))
	if rc < 0 {
		return libusb_error("libusb_claim_interface", rc)
	}
	return nil
}

func Release_Interface(hdl Device_Handle, interface_number int) error {
	rc := int(C.libusb_release_interface(hdl, (C.int)(interface_number)))
	if rc < 0 {
		return libusb_error("libusb_release_interface", rc)
	}
	return nil
}

func Set_Interface_Alt_Setting(hdl Device_Handle, interface_number int, alternate_setting int) error {
	rc := int(C.libusb_set_interface_alt_setting(hdl, (C.int)(interface_number), (C.int)(alternate_setting)))
	if rc < 0 {
		return libusb_error("libusb_set_interface_alt_setting", rc)
	}
	return nil
}

func Clear_Halt(hdl Device_Handle, endpoint uint8) error {
	rc := int(C.libusb_clear_halt(hdl, (C.uchar)(endpoint)))
	if rc < 0 {
		return libusb_error("libusb_clear_halt", rc)
	}
	return nil
}

func Reset_Device(hdl Device_Handle) error {
	rc := int(C.libusb_reset_device(hdl))
	if rc < 0 {
		return libusb_error("libusb_reset_device", rc)
	}
	return nil
}

func Kernel_Driver_Active(hdl Device_Handle, interface_number int) (bool, error) {
	rc := int(C.libusb_kernel_driver_active(hdl, (C.int)(interface_number)))
	if rc < 0 {
		return false, libusb_error("libusb_kernel_driver_active", rc)
	}
	return rc != 0, nil
}

func Detach_Kernel_Driver(hdl Device_Handle, interface_number int) error {
	rc := int(C.libusb_detach_kernel_driver(hdl, (C.int)(interface_number)))
	if rc < 0 {
		return libusb_error("libusb_detach_kernel_driver", rc)
	}
	return nil
}

func Attach_Kernel_Driver(hdl Device_Handle, interface_number int) error {
	rc := int(C.libusb_attach_kernel_driver(hdl, (C.int)(interface_number)))
	if rc < 0 {
		return libusb_error("libusb_attach_kernel_driver", rc)
	}
	return nil
}

func Set_Auto_Detach_Kernel_Driver(hdl Device_Handle, enable int) error {
	rc := int(C.libusb_set_auto_detach_kernel_driver(hdl, (C.int)(enable)))
	if rc < 0 {
		return libusb_error("libusb_set_auto_detach_kernel_driver", rc)
	}
	return nil
}

//-----------------------------------------------------------------------------
// Miscellaneous

func Has_Capability(capability uint32) bool {
	rc := int(C.libusb_has_capability((C.uint32_t)(capability)))
	return rc != 0
}

func Error_Name(code int) string {
	return C.GoString(C.libusb_error_name(C.int(code)))
}

func Get_Version() *Version {
	ver := (*C.struct_libusb_version)(unsafe.Pointer(C.libusb_get_version()))
	return &Version{
		Major:    uint16(ver.major),
		Minor:    uint16(ver.minor),
		Micro:    uint16(ver.micro),
		Nano:     uint16(ver.nano),
		Rc:       C.GoString(ver.rc),
		Describe: C.GoString(ver.describe),
	}
}

func CPU_To_LE16(x uint16) uint16 {
	return uint16(C.libusb_cpu_to_le16((C.uint16_t)(x)))
}

func Setlocale(locale string) error {
	cstr := C.CString(locale)
	rc := int(C.libusb_setlocale(cstr))
	if rc < 0 {
		return libusb_error("libusb_setlocale", rc)
	}
	return nil
}

func Strerror(errcode int) string {
	return C.GoString(C.libusb_strerror(int32(errcode)))
}

//-----------------------------------------------------------------------------
// USB descriptors

func Get_Device_Descriptor(dev Device) (*Device_Descriptor, error) {
	var dd C.struct_libusb_device_descriptor
	rc := int(C.libusb_get_device_descriptor(dev, &dd))
	if rc != LIBUSB_SUCCESS {
		return nil, libusb_error("libusb_get_device_descriptor", rc)
	}
	return &Device_Descriptor{
		BLength:            uint8(dd.bLength),
		BDescriptorType:    uint8(dd.bDescriptorType),
		BcdUSB:             uint16(dd.bcdUSB),
		BDeviceClass:       uint8(dd.bDeviceClass),
		BDeviceSubClass:    uint8(dd.bDeviceSubClass),
		BDeviceProtocol:    uint8(dd.bDeviceProtocol),
		BMaxPacketSize0:    uint8(dd.bMaxPacketSize0),
		IdVendor:           uint16(dd.idVendor),
		IdProduct:          uint16(dd.idProduct),
		BcdDevice:          uint16(dd.bcdDevice),
		IManufacturer:      uint8(dd.iManufacturer),
		IProduct:           uint8(dd.iProduct),
		ISerialNumber:      uint8(dd.iSerialNumber),
		BNumConfigurations: uint8(dd.bNumConfigurations),
	}, nil
}

// int 	libusb_get_active_config_descriptor (libusb_device *dev, struct libusb_config_descriptor **config)
// int 	libusb_get_config_descriptor (libusb_device *dev, uint8_t config_index, struct libusb_config_descriptor **config)
// int 	libusb_get_config_descriptor_by_value (libusb_device *dev, uint8_t bConfigurationValue, struct libusb_config_descriptor **config)
// void 	libusb_free_config_descriptor (struct libusb_config_descriptor *config)
// int 	libusb_get_ss_endpoint_companion_descriptor (struct libusb_context *ctx, const struct libusb_endpoint_descriptor *endpoint, struct libusb_ss_endpoint_companion_descriptor **ep_comp)
// void 	libusb_free_ss_endpoint_companion_descriptor (struct libusb_ss_endpoint_companion_descriptor *ep_comp)
// int 	libusb_get_bos_descriptor (libusb_device_handle *handle, struct libusb_bos_descriptor **bos)
// void 	libusb_free_bos_descriptor (struct libusb_bos_descriptor *bos)
// int 	libusb_get_usb_2_0_extension_descriptor (struct libusb_context *ctx, struct libusb_bos_dev_capability_descriptor *dev_cap, struct libusb_usb_2_0_extension_descriptor **usb_2_0_extension)
// void 	libusb_free_usb_2_0_extension_descriptor (struct libusb_usb_2_0_extension_descriptor *usb_2_0_extension)
// int 	libusb_get_ss_usb_device_capability_descriptor (struct libusb_context *ctx, struct libusb_bos_dev_capability_descriptor *dev_cap, struct libusb_ss_usb_device_capability_descriptor **ss_usb_device_cap)
// void 	libusb_free_ss_usb_device_capability_descriptor (struct libusb_ss_usb_device_capability_descriptor *ss_usb_device_cap)
// int 	libusb_get_container_id_descriptor (struct libusb_context *ctx, struct libusb_bos_dev_capability_descriptor *dev_cap, struct libusb_container_id_descriptor **container_id)
// void 	libusb_free_container_id_descriptor (struct libusb_container_id_descriptor *container_id)
// int 	libusb_get_string_descriptor_ascii (libusb_device_handle *dev, uint8_t desc_index, unsigned char *data, int length)
// static int 	libusb_get_descriptor (libusb_device_handle *dev, uint8_t desc_type, uint8_t desc_index, unsigned char *data, int length)
// static int 	libusb_get_string_descriptor (libusb_device_handle *dev, uint8_t desc_index, uint16_t langid, unsigned char *data, int length)

//-----------------------------------------------------------------------------
// Device hotplug event notification

//int 	libusb_hotplug_register_callback (libusb_context *ctx, libusb_hotplug_event events, libusb_hotplug_flag flags, int vendor_id, int product_id, int dev_class, libusb_hotplug_callback_fn cb_fn, void *user_data, libusb_hotplug_callback_handle *handle)
//void 	libusb_hotplug_deregister_callback (libusb_context *ctx, libusb_hotplug_callback_handle handle)

//-----------------------------------------------------------------------------
//Asynchronous device I/O

// int 	libusb_alloc_streams (libusb_device_handle *dev, uint32_t num_streams, unsigned char *endpoints, int num_endpoints)
// int 	libusb_free_streams (libusb_device_handle *dev, unsigned char *endpoints, int num_endpoints)
// struct libusb_transfer * 	libusb_alloc_transfer (int iso_packets)
// void 	libusb_free_transfer (struct libusb_transfer *transfer)
// int 	libusb_submit_transfer (struct libusb_transfer *transfer)
// int 	libusb_cancel_transfer (struct libusb_transfer *transfer)
// void 	libusb_transfer_set_stream_id (struct libusb_transfer *transfer, uint32_t stream_id)
// uint32_t 	libusb_transfer_get_stream_id (struct libusb_transfer *transfer)
// static unsigned char * 	libusb_control_transfer_get_data (struct libusb_transfer *transfer)
// static struct libusb_control_setup * 	libusb_control_transfer_get_setup (struct libusb_transfer *transfer)
// static void 	libusb_fill_control_setup (unsigned char *buffer, uint8_t bmRequestType, uint8_t bRequest, uint16_t wValue, uint16_t wIndex, uint16_t wLength)
// static void 	libusb_fill_control_transfer (struct libusb_transfer *transfer, libusb_device_handle *dev_handle, unsigned char *buffer, libusb_transfer_cb_fn callback, void *user_data, unsigned int timeout)
// static void 	libusb_fill_bulk_transfer (struct libusb_transfer *transfer, libusb_device_handle *dev_handle, unsigned char endpoint, unsigned char *buffer, int length, libusb_transfer_cb_fn callback, void *user_data, unsigned int timeout)
// static void 	libusb_fill_bulk_stream_transfer (struct libusb_transfer *transfer, libusb_device_handle *dev_handle, unsigned char endpoint, uint32_t stream_id, unsigned char *buffer, int length, libusb_transfer_cb_fn callback, void *user_data, unsigned int timeout)
// static void 	libusb_fill_interrupt_transfer (struct libusb_transfer *transfer, libusb_device_handle *dev_handle, unsigned char endpoint, unsigned char *buffer, int length, libusb_transfer_cb_fn callback, void *user_data, unsigned int timeout)
// static void 	libusb_fill_iso_transfer (struct libusb_transfer *transfer, libusb_device_handle *dev_handle, unsigned char endpoint, unsigned char *buffer, int length, int num_iso_packets, libusb_transfer_cb_fn callback, void *user_data, unsigned int timeout)
// static void 	libusb_set_iso_packet_lengths (struct libusb_transfer *transfer, unsigned int length)
// static unsigned char * 	libusb_get_iso_packet_buffer (struct libusb_transfer *transfer, unsigned int packet)
// static unsigned char * 	libusb_get_iso_packet_buffer_simple (struct libusb_transfer *transfer, unsigned int packet)

//-----------------------------------------------------------------------------
// Polling and timing

// int 	libusb_try_lock_events (libusb_context *ctx)
// void 	libusb_lock_events (libusb_context *ctx)
// void 	libusb_unlock_events (libusb_context *ctx)
// int 	libusb_event_handling_ok (libusb_context *ctx)
// int 	libusb_event_handler_active (libusb_context *ctx)
// void 	libusb_lock_event_waiters (libusb_context *ctx)
// void 	libusb_unlock_event_waiters (libusb_context *ctx)
// int 	libusb_wait_for_event (libusb_context *ctx, struct timeval *tv)
// int 	libusb_handle_events_timeout_completed (libusb_context *ctx, struct timeval *tv, int *completed)
// int 	libusb_handle_events_timeout (libusb_context *ctx, struct timeval *tv)
// int 	libusb_handle_events (libusb_context *ctx)
// int 	libusb_handle_events_completed (libusb_context *ctx, int *completed)
// int 	libusb_handle_events_locked (libusb_context *ctx, struct timeval *tv)
// int 	libusb_pollfds_handle_timeouts (libusb_context *ctx)
// int 	libusb_get_next_timeout (libusb_context *ctx, struct timeval *tv)
// void 	libusb_set_pollfd_notifiers (libusb_context *ctx, libusb_pollfd_added_cb added_cb, libusb_pollfd_removed_cb removed_cb, void *user_data)
// const struct libusb_pollfd ** 	libusb_get_pollfds (libusb_context *ctx)
// void 	libusb_free_pollfds (const struct libusb_pollfd **pollfds)

//-----------------------------------------------------------------------------
// Synchronous device I/O

func Control_Transfer(hdl Device_Handle, bmRequestType uint8, bRequest uint8, wValue uint16, wIndex uint16, data []byte, timeout uint) ([]byte, error) {
	//int 	libusb_control_transfer (libusb_device_handle *dev_handle, uint8_t bmRequestType, uint8_t bRequest, uint16_t wValue, uint16_t wIndex, unsigned char *data, uint16_t wLength, unsigned int timeout)
	return nil, nil
}

func Bulk_Transfer(hdl Device_Handle, endpoint uint8, data []byte, timeout uint) ([]byte, error) {
	var length int
	var transferred C.int
	if endpoint&LIBUSB_ENDPOINT_IN != 0 {
		// read device
		length = cap(data)
	} else {
		// write device
		length = len(data)
	}
	rc := int(C.libusb_bulk_transfer(hdl, (C.uchar)(endpoint), (*C.uchar)(&data[0]), (C.int)(length), &transferred, (C.uint)(timeout)))
	if rc != LIBUSB_SUCCESS {
		return nil, libusb_error("libusb_bulk_transfer", rc)
	}
	return data[:int(transferred)], nil
}

func Interrupt_Transfer(hdl Device_Handle, endpoint uint8, data []byte, timeout uint) ([]byte, error) {
	var length int
	var transferred C.int
	if endpoint&LIBUSB_ENDPOINT_IN != 0 {
		// read device
		length = cap(data)
	} else {
		// write device
		length = len(data)
	}
	rc := int(C.libusb_interrupt_transfer(hdl, (C.uchar)(endpoint), (*C.uchar)(&data[0]), (C.int)(length), &transferred, (C.uint)(timeout)))
	if rc != LIBUSB_SUCCESS {
		return nil, libusb_error("libusb_interrupt_transfer", rc)
	}
	return data[:int(transferred)], nil
}

//-----------------------------------------------------------------------------
