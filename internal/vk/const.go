package vk

const (
	STRUCTURE_TYPE_APPLICATION_INFO     = 0
	STRUCTURE_TYPE_INSTANCE_CREATE_INFO = 1
)

const (
	SUCCESS                     = 0
	NOT_READY                   = 1
	TIMEOUT                     = 2
	EVENT_SET                   = 3
	EVENT_RESET                 = 4
	INCOMPLETE                  = 5
	ERROR_OUT_OF_HOST_MEMORY    = -1
	ERROR_OUT_OF_DEVICE_MEMORY  = -2
	ERROR_INITIALIZATION_FAILED = -3
	ERROR_DEVICE_LOST           = -4
	ERROR_MEMORY_MAP_FAILED     = -5
	ERROR_LAYER_NOT_PRESENT     = -6
	ERROR_EXTENSION_NOT_PRESENT = -7
	ERROR_FEATURE_NOT_PRESENT   = -8
	ERROR_INCOMPATIBLE_DRIVER   = -9
	ERROR_TOO_MANY_OBJECTS      = -10
	ERROR_FORMAT_NOT_SUPPORTED  = -11
	ERROR_FRAGMENTED_POOL       = -12
	ERROR_UNKNOWN               = -13
)

const (
	PHYSICAL_DEVICE_TYPE_OTHER          = 0
	PHYSICAL_DEVICE_TYPE_INTEGRATED_GPU = 1
	PHYSICAL_DEVICE_TYPE_DISCRETE_GPU   = 2
	PHYSICAL_DEVICE_TYPE_VIRTUAL_GPU    = 3
	PHYSICAL_DEVICE_TYPE_CPU            = 4
)
