package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/log.h>
import "C"

// Sets the log level.
// C-Function: av_log_set_level
func LogSetLevel(level int) {
  C.av_log_set_level(C.int(level))
}
