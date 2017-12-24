package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// Copies contents of src AVCodecParameter to dst. See ffmpeg docs for semantics
//
// C-Function: avcodec_parameters_copy
func (c *CodecParameters) CopyParams(dst *CodecParameters) {
	C.avcodec_parameters_copy((*C.struct_AVCodecParameters)(dst),
		(*C.struct_AVCodecParameters)(c))
}
