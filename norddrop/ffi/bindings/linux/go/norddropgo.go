/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ffi/bindings/norddrop.i

package norddropgo

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>
#include "wtf.h"


typedef int intgo;
typedef unsigned int uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef void* swig_type_1;
typedef void* swig_type_2;
typedef void* swig_type_3;
typedef void* swig_type_4;
typedef void* swig_type_5;
typedef void* swig_type_6;
typedef _gostring_ swig_type_7;
typedef _gostring_ swig_type_8;
typedef _gostring_ swig_type_9;
typedef _gostring_ swig_type_10;
typedef _gostring_ swig_type_11;
typedef _gostring_ swig_type_12;
typedef _gostring_ swig_type_13;
typedef _gostring_ swig_type_14;
typedef _gostring_ swig_type_15;
typedef _gostring_ swig_type_16;
typedef _gostring_ swig_type_17;
typedef _gostring_ swig_type_18;
typedef _gostring_ swig_type_19;
typedef _gostring_ swig_type_20;
typedef _gostring_ swig_type_21;
extern void _wrap_Swig_free_norddropgo_66601ed8ddecf11f(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_norddropgo_66601ed8ddecf11f(swig_intgo arg1);
extern swig_intgo _wrap_NORDDROPLOGCRITICAL_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPLOGERROR_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPLOGWARNING_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPLOGINFO_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPLOGDEBUG_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPLOGTRACE_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESOK_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESERROR_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESINVALIDSTRING_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESBADINPUT_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESJSONPARSE_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESTRANSFERCREATE_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESNOTSTARTED_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESADDRINUSE_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESINSTANCESTART_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESINSTANCESTOP_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESINVALIDPRIVKEY_norddropgo_66601ed8ddecf11f(void);
extern swig_intgo _wrap_NORDDROPRESDBERROR_norddropgo_66601ed8ddecf11f(void);
extern void _wrap_NorddropEventCb_Ctx_set_norddropgo_66601ed8ddecf11f(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_NorddropEventCb_Ctx_get_norddropgo_66601ed8ddecf11f(uintptr_t arg1);
extern void _wrap_NorddropEventCb_Cb_set_norddropgo_66601ed8ddecf11f(uintptr_t arg1, swig_type_1 arg2);
extern swig_type_2 _wrap_NorddropEventCb_Cb_get_norddropgo_66601ed8ddecf11f(uintptr_t arg1);
extern uintptr_t _wrap_new_NorddropEventCb_norddropgo_66601ed8ddecf11f(void);
extern void _wrap_delete_NorddropEventCb_norddropgo_66601ed8ddecf11f(uintptr_t arg1);
extern void _wrap_NorddropLoggerCb_Ctx_set_norddropgo_66601ed8ddecf11f(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_NorddropLoggerCb_Ctx_get_norddropgo_66601ed8ddecf11f(uintptr_t arg1);
extern void _wrap_NorddropLoggerCb_Cb_set_norddropgo_66601ed8ddecf11f(uintptr_t arg1, swig_type_3 arg2);
extern swig_type_4 _wrap_NorddropLoggerCb_Cb_get_norddropgo_66601ed8ddecf11f(uintptr_t arg1);
extern uintptr_t _wrap_new_NorddropLoggerCb_norddropgo_66601ed8ddecf11f(void);
extern void _wrap_delete_NorddropLoggerCb_norddropgo_66601ed8ddecf11f(uintptr_t arg1);
extern void _wrap_NorddropPubkeyCb_Ctx_set_norddropgo_66601ed8ddecf11f(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_NorddropPubkeyCb_Ctx_get_norddropgo_66601ed8ddecf11f(uintptr_t arg1);
extern void _wrap_NorddropPubkeyCb_Cb_set_norddropgo_66601ed8ddecf11f(uintptr_t arg1, swig_type_5 arg2);
extern swig_type_6 _wrap_NorddropPubkeyCb_Cb_get_norddropgo_66601ed8ddecf11f(uintptr_t arg1);
extern uintptr_t _wrap_new_NorddropPubkeyCb_norddropgo_66601ed8ddecf11f(void);
extern void _wrap_delete_NorddropPubkeyCb_norddropgo_66601ed8ddecf11f(uintptr_t arg1);

#include <string.h>

extern uintptr_t _wrap_new_Norddrop_norddropgo_66601ed8ddecf11f(norddrop_event_cb arg1, swig_intgo arg2, norddrop_logger_cb arg3, norddrop_pubkey_cb arg4, swig_type_7 arg5);
extern void _wrap_delete_Norddrop_norddropgo_66601ed8ddecf11f(uintptr_t arg1);
extern swig_intgo _wrap_Norddrop_Start_norddropgo_66601ed8ddecf11f(uintptr_t arg1, swig_type_8 arg2, swig_type_9 arg3);
extern swig_intgo _wrap_Norddrop_Stop_norddropgo_66601ed8ddecf11f(uintptr_t arg1);
extern swig_intgo _wrap_Norddrop_CancelTransfer_norddropgo_66601ed8ddecf11f(uintptr_t arg1, swig_type_10 arg2);
extern swig_intgo _wrap_Norddrop_CancelFile_norddropgo_66601ed8ddecf11f(uintptr_t arg1, swig_type_11 arg2, swig_type_12 arg3);
extern swig_intgo _wrap_Norddrop_Download_norddropgo_66601ed8ddecf11f(uintptr_t arg1, swig_type_13 arg2, swig_type_14 arg3, swig_type_15 arg4);
extern swig_type_16 _wrap_Norddrop_NewTransfer_norddropgo_66601ed8ddecf11f(uintptr_t arg1, swig_type_17 arg2, swig_type_18 arg3);
extern swig_intgo _wrap_Norddrop_PurgeTransfers_norddropgo_66601ed8ddecf11f(uintptr_t arg1, swig_type_19 arg2);
extern swig_intgo _wrap_Norddrop_PurgeTransfersUntil_norddropgo_66601ed8ddecf11f(uintptr_t arg1, uintptr_t arg2);
extern swig_type_20 _wrap_Norddrop_GetTransfersSince_norddropgo_66601ed8ddecf11f(uintptr_t arg1, uintptr_t arg2);
extern swig_type_21 _wrap_Norddrop_Version_norddropgo_66601ed8ddecf11f(void);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex


type swig_gostring struct { p uintptr; n int32 }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_norddropgo_66601ed8ddecf11f(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type Enum_SS_norddrop_log_level int
func _swig_getNORDDROPLOGCRITICAL() (_swig_ret Enum_SS_norddrop_log_level) {
	var swig_r Enum_SS_norddrop_log_level
	swig_r = (Enum_SS_norddrop_log_level)(C._wrap_NORDDROPLOGCRITICAL_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPLOGCRITICAL Enum_SS_norddrop_log_level = _swig_getNORDDROPLOGCRITICAL()
func _swig_getNORDDROPLOGERROR() (_swig_ret Enum_SS_norddrop_log_level) {
	var swig_r Enum_SS_norddrop_log_level
	swig_r = (Enum_SS_norddrop_log_level)(C._wrap_NORDDROPLOGERROR_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPLOGERROR Enum_SS_norddrop_log_level = _swig_getNORDDROPLOGERROR()
func _swig_getNORDDROPLOGWARNING() (_swig_ret Enum_SS_norddrop_log_level) {
	var swig_r Enum_SS_norddrop_log_level
	swig_r = (Enum_SS_norddrop_log_level)(C._wrap_NORDDROPLOGWARNING_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPLOGWARNING Enum_SS_norddrop_log_level = _swig_getNORDDROPLOGWARNING()
func _swig_getNORDDROPLOGINFO() (_swig_ret Enum_SS_norddrop_log_level) {
	var swig_r Enum_SS_norddrop_log_level
	swig_r = (Enum_SS_norddrop_log_level)(C._wrap_NORDDROPLOGINFO_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPLOGINFO Enum_SS_norddrop_log_level = _swig_getNORDDROPLOGINFO()
func _swig_getNORDDROPLOGDEBUG() (_swig_ret Enum_SS_norddrop_log_level) {
	var swig_r Enum_SS_norddrop_log_level
	swig_r = (Enum_SS_norddrop_log_level)(C._wrap_NORDDROPLOGDEBUG_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPLOGDEBUG Enum_SS_norddrop_log_level = _swig_getNORDDROPLOGDEBUG()
func _swig_getNORDDROPLOGTRACE() (_swig_ret Enum_SS_norddrop_log_level) {
	var swig_r Enum_SS_norddrop_log_level
	swig_r = (Enum_SS_norddrop_log_level)(C._wrap_NORDDROPLOGTRACE_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPLOGTRACE Enum_SS_norddrop_log_level = _swig_getNORDDROPLOGTRACE()
type Enum_SS_norddrop_result int
func _swig_getNORDDROPRESOK() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESOK_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESOK Enum_SS_norddrop_result = _swig_getNORDDROPRESOK()
func _swig_getNORDDROPRESERROR() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESERROR_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESERROR Enum_SS_norddrop_result = _swig_getNORDDROPRESERROR()
func _swig_getNORDDROPRESINVALIDSTRING() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESINVALIDSTRING_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESINVALIDSTRING Enum_SS_norddrop_result = _swig_getNORDDROPRESINVALIDSTRING()
func _swig_getNORDDROPRESBADINPUT() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESBADINPUT_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESBADINPUT Enum_SS_norddrop_result = _swig_getNORDDROPRESBADINPUT()
func _swig_getNORDDROPRESJSONPARSE() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESJSONPARSE_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESJSONPARSE Enum_SS_norddrop_result = _swig_getNORDDROPRESJSONPARSE()
func _swig_getNORDDROPRESTRANSFERCREATE() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESTRANSFERCREATE_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESTRANSFERCREATE Enum_SS_norddrop_result = _swig_getNORDDROPRESTRANSFERCREATE()
func _swig_getNORDDROPRESNOTSTARTED() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESNOTSTARTED_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESNOTSTARTED Enum_SS_norddrop_result = _swig_getNORDDROPRESNOTSTARTED()
func _swig_getNORDDROPRESADDRINUSE() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESADDRINUSE_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESADDRINUSE Enum_SS_norddrop_result = _swig_getNORDDROPRESADDRINUSE()
func _swig_getNORDDROPRESINSTANCESTART() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESINSTANCESTART_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESINSTANCESTART Enum_SS_norddrop_result = _swig_getNORDDROPRESINSTANCESTART()
func _swig_getNORDDROPRESINSTANCESTOP() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESINSTANCESTOP_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESINSTANCESTOP Enum_SS_norddrop_result = _swig_getNORDDROPRESINSTANCESTOP()
func _swig_getNORDDROPRESINVALIDPRIVKEY() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESINVALIDPRIVKEY_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESINVALIDPRIVKEY Enum_SS_norddrop_result = _swig_getNORDDROPRESINVALIDPRIVKEY()
func _swig_getNORDDROPRESDBERROR() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	swig_r = (Enum_SS_norddrop_result)(C._wrap_NORDDROPRESDBERROR_norddropgo_66601ed8ddecf11f())
	return swig_r
}

var NORDDROPRESDBERROR Enum_SS_norddrop_result = _swig_getNORDDROPRESDBERROR()
type SwigcptrNorddropEventCb uintptr

func (p SwigcptrNorddropEventCb) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNorddropEventCb) SwigIsNorddropEventCb() {
}

func (arg1 SwigcptrNorddropEventCb) SetCtx(arg2 uintptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_NorddropEventCb_Ctx_set_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrNorddropEventCb) GetCtx() (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_NorddropEventCb_Ctx_get_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrNorddropEventCb) SetCb(arg2 _swig_fnptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_NorddropEventCb_Cb_set_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), C.swig_type_1(_swig_i_1))
}

func (arg1 SwigcptrNorddropEventCb) GetCb() (_swig_ret _swig_fnptr) {
	var swig_r _swig_fnptr
	_swig_i_0 := arg1
	swig_r = (_swig_fnptr)(C._wrap_NorddropEventCb_Cb_get_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewNorddropEventCb() (_swig_ret NorddropEventCb) {
	var swig_r NorddropEventCb
	swig_r = (NorddropEventCb)(SwigcptrNorddropEventCb(C._wrap_new_NorddropEventCb_norddropgo_66601ed8ddecf11f()))
	return swig_r
}

func DeleteNorddropEventCb(arg1 NorddropEventCb) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_NorddropEventCb_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0))
}

type NorddropEventCb interface {
	Swigcptr() uintptr
	SwigIsNorddropEventCb()
	SetCtx(arg2 uintptr)
	GetCtx() (_swig_ret uintptr)
	SetCb(arg2 _swig_fnptr)
	GetCb() (_swig_ret _swig_fnptr)
}

type SwigcptrNorddropLoggerCb uintptr

func (p SwigcptrNorddropLoggerCb) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNorddropLoggerCb) SwigIsNorddropLoggerCb() {
}

func (arg1 SwigcptrNorddropLoggerCb) SetCtx(arg2 uintptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_NorddropLoggerCb_Ctx_set_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrNorddropLoggerCb) GetCtx() (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_NorddropLoggerCb_Ctx_get_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrNorddropLoggerCb) SetCb(arg2 _swig_fnptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_NorddropLoggerCb_Cb_set_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), C.swig_type_3(_swig_i_1))
}

func (arg1 SwigcptrNorddropLoggerCb) GetCb() (_swig_ret _swig_fnptr) {
	var swig_r _swig_fnptr
	_swig_i_0 := arg1
	swig_r = (_swig_fnptr)(C._wrap_NorddropLoggerCb_Cb_get_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewNorddropLoggerCb() (_swig_ret NorddropLoggerCb) {
	var swig_r NorddropLoggerCb
	swig_r = (NorddropLoggerCb)(SwigcptrNorddropLoggerCb(C._wrap_new_NorddropLoggerCb_norddropgo_66601ed8ddecf11f()))
	return swig_r
}

func DeleteNorddropLoggerCb(arg1 NorddropLoggerCb) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_NorddropLoggerCb_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0))
}

type NorddropLoggerCb interface {
	Swigcptr() uintptr
	SwigIsNorddropLoggerCb()
	SetCtx(arg2 uintptr)
	GetCtx() (_swig_ret uintptr)
	SetCb(arg2 _swig_fnptr)
	GetCb() (_swig_ret _swig_fnptr)
}

type SwigcptrNorddropPubkeyCb uintptr

func (p SwigcptrNorddropPubkeyCb) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNorddropPubkeyCb) SwigIsNorddropPubkeyCb() {
}

func (arg1 SwigcptrNorddropPubkeyCb) SetCtx(arg2 uintptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_NorddropPubkeyCb_Ctx_set_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrNorddropPubkeyCb) GetCtx() (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_NorddropPubkeyCb_Ctx_get_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrNorddropPubkeyCb) SetCb(arg2 _swig_fnptr) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_NorddropPubkeyCb_Cb_set_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), C.swig_type_5(_swig_i_1))
}

func (arg1 SwigcptrNorddropPubkeyCb) GetCb() (_swig_ret _swig_fnptr) {
	var swig_r _swig_fnptr
	_swig_i_0 := arg1
	swig_r = (_swig_fnptr)(C._wrap_NorddropPubkeyCb_Cb_get_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewNorddropPubkeyCb() (_swig_ret NorddropPubkeyCb) {
	var swig_r NorddropPubkeyCb
	swig_r = (NorddropPubkeyCb)(SwigcptrNorddropPubkeyCb(C._wrap_new_NorddropPubkeyCb_norddropgo_66601ed8ddecf11f()))
	return swig_r
}

func DeleteNorddropPubkeyCb(arg1 NorddropPubkeyCb) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_NorddropPubkeyCb_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0))
}

type NorddropPubkeyCb interface {
	Swigcptr() uintptr
	SwigIsNorddropPubkeyCb()
	SetCtx(arg2 uintptr)
	GetCtx() (_swig_ret uintptr)
	SetCb(arg2 _swig_fnptr)
	GetCb() (_swig_ret _swig_fnptr)
}


var eventCallbacks = map[uintptr]func(string){}
var loggerCallbacks = map[uintptr]func(int, string){}
var pubkeyCallbacks = map[uintptr]func(string) []byte{}
// Note: This can only ensure enough place for 8 callbacks
// Application can crash when creating more if these the last
// items on stack
// The real fix for this would be to avoid using pointers where not necessary. In this case - key to hashmap
var arbitraryValue = uint64(0)
var arbitraryAddress = uintptr(unsafe.Pointer(&arbitraryValue))

func maxEventCbIndex() uintptr {
        maxI := arbitraryAddress
        for i := range eventCallbacks {
                if i > maxI {
                        maxI = i
                }
        }
        return maxI
}

func maxLoggerCbIndex() uintptr {
        maxI := arbitraryAddress
        for i := range loggerCallbacks {
                if i > maxI {
                        maxI = i
                }
        }
        return maxI
}

func maxPubkeyCbIndex() uintptr {
        maxI := arbitraryAddress
        for i := range pubkeyCallbacks {
                if i > maxI {
                        maxI = i
                }
        }
        return maxI
}


//export call_norddrop_event_cb
func call_norddrop_event_cb(ctx uintptr, str *C.char) {
        if callback, ok := eventCallbacks[ctx]; ok {
                callback(C.GoString(str))
        }
}


//export call_norddrop_logger_cb
func call_norddrop_logger_cb(ctx uintptr, level C.int, str *C.char) {
        if callback, ok := loggerCallbacks[ctx]; ok {
                callback(int(level), C.GoString(str))
        }
}


//export call_norddrop_pubkey_cb
func call_norddrop_pubkey_cb(ctx uintptr, ip *C.char, pubkey *C.char) C.int {
        if callback, ok := pubkeyCallbacks[ctx]; ok {
                gokey := callback(C.GoString(ip))

                if gokey != nil {
                        ckey := C.CBytes(gokey)
                        C.memcpy(unsafe.Pointer(pubkey), ckey, 32)
                        C.free(ckey)

                        return 0
                }
        }

        return 1
}

type SwigcptrNorddrop uintptr

func (p SwigcptrNorddrop) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNorddrop) SwigIsNorddrop() {
}

func NewNorddrop(arg1 func(string), arg2 Enum_SS_norddrop_log_level, arg3 func(int, string), arg4 func(string) []byte, arg5 string) (_swig_ret Norddrop) {
	var swig_r Norddrop
	var _swig_i_0 C.norddrop_event_cb
{
        index := maxEventCbIndex() + 1
        cb := C.norddrop_event_cb{
                ctx: unsafe.Pointer(index),
                cb: (C.norddrop_event_fn)(C.call_norddrop_event_cb),
        }
        eventCallbacks[index] = arg1
        _swig_i_0 = cb
}
	_swig_i_1 := arg2
	var _swig_i_2 C.norddrop_logger_cb
{
        index := maxLoggerCbIndex() + 1
        cb := C.norddrop_logger_cb{
                ctx: unsafe.Pointer(index),
                cb: (C.norddrop_logger_fn)(C.call_norddrop_logger_cb),
        }
        loggerCallbacks[index] = arg3
        _swig_i_2 = cb
}
	var _swig_i_3 C.norddrop_pubkey_cb
{
        index := maxPubkeyCbIndex() + 1
        cb := C.norddrop_pubkey_cb{
                ctx: unsafe.Pointer(index),
                cb: (C.norddrop_pubkey_fn)(C.call_norddrop_pubkey_cb),
        }
        pubkeyCallbacks[index] = arg4
        _swig_i_3 = cb
}
	_swig_i_4 := arg5
	swig_r = (Norddrop)(SwigcptrNorddrop(C._wrap_new_Norddrop_norddropgo_66601ed8ddecf11f(C.norddrop_event_cb(_swig_i_0), C.swig_intgo(_swig_i_1), C.norddrop_logger_cb(_swig_i_2), C.norddrop_pubkey_cb(_swig_i_3), *(*C.swig_type_7)(unsafe.Pointer(&_swig_i_4)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg5
	}
	var swig_r_1 Norddrop
{
    if swig_r == SwigcptrNorddrop(0) {
        swig_r_1 = nil
    }

    swig_r_1 = swig_r
}
	return swig_r_1
}

func DeleteNorddrop(arg1 Norddrop) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Norddrop_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrNorddrop) Start(arg2 string, arg3 string) (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Enum_SS_norddrop_result)(C._wrap_Norddrop_Start_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), *(*C.swig_type_8)(unsafe.Pointer(&_swig_i_1)), *(*C.swig_type_9)(unsafe.Pointer(&_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	return swig_r
}

func (arg1 SwigcptrNorddrop) Stop() (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	_swig_i_0 := arg1
	swig_r = (Enum_SS_norddrop_result)(C._wrap_Norddrop_Stop_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrNorddrop) CancelTransfer(arg2 string) (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Enum_SS_norddrop_result)(C._wrap_Norddrop_CancelTransfer_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), *(*C.swig_type_10)(unsafe.Pointer(&_swig_i_1))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrNorddrop) CancelFile(arg2 string, arg3 string) (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Enum_SS_norddrop_result)(C._wrap_Norddrop_CancelFile_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), *(*C.swig_type_11)(unsafe.Pointer(&_swig_i_1)), *(*C.swig_type_12)(unsafe.Pointer(&_swig_i_2))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	return swig_r
}

func (arg1 SwigcptrNorddrop) Download(arg2 string, arg3 string, arg4 string) (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (Enum_SS_norddrop_result)(C._wrap_Norddrop_Download_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), *(*C.swig_type_13)(unsafe.Pointer(&_swig_i_1)), *(*C.swig_type_14)(unsafe.Pointer(&_swig_i_2)), *(*C.swig_type_15)(unsafe.Pointer(&_swig_i_3))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg4
	}
	return swig_r
}

func (arg1 SwigcptrNorddrop) NewTransfer(arg2 string, arg3 string) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r_p := C._wrap_Norddrop_NewTransfer_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), *(*C.swig_type_17)(unsafe.Pointer(&_swig_i_1)), *(*C.swig_type_18)(unsafe.Pointer(&_swig_i_2)))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrNorddrop) PurgeTransfers(arg2 string) (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Enum_SS_norddrop_result)(C._wrap_Norddrop_PurgeTransfers_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), *(*C.swig_type_19)(unsafe.Pointer(&_swig_i_1))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrNorddrop) PurgeTransfersUntil(arg2 Int64_t) (_swig_ret Enum_SS_norddrop_result) {
	var swig_r Enum_SS_norddrop_result
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Enum_SS_norddrop_result)(C._wrap_Norddrop_PurgeTransfersUntil_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrNorddrop) GetTransfersSince(arg2 Int64_t) (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	swig_r_p := C._wrap_Norddrop_GetTransfersSince_norddropgo_66601ed8ddecf11f(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func NorddropVersion() (_swig_ret string) {
	var swig_r string
	swig_r_p := C._wrap_Norddrop_Version_norddropgo_66601ed8ddecf11f()
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

type Norddrop interface {
	Swigcptr() uintptr
	SwigIsNorddrop()
	Start(arg2 string, arg3 string) (_swig_ret Enum_SS_norddrop_result)
	Stop() (_swig_ret Enum_SS_norddrop_result)
	CancelTransfer(arg2 string) (_swig_ret Enum_SS_norddrop_result)
	CancelFile(arg2 string, arg3 string) (_swig_ret Enum_SS_norddrop_result)
	Download(arg2 string, arg3 string, arg4 string) (_swig_ret Enum_SS_norddrop_result)
	NewTransfer(arg2 string, arg3 string) (_swig_ret string)
	PurgeTransfers(arg2 string) (_swig_ret Enum_SS_norddrop_result)
	PurgeTransfersUntil(arg2 Int64_t) (_swig_ret Enum_SS_norddrop_result)
	GetTransfersSince(arg2 Int64_t) (_swig_ret string)
}


type SwigcptrInt64_t uintptr
type Int64_t interface {
	Swigcptr() uintptr;
}
func (p SwigcptrInt64_t) Swigcptr() uintptr {
	return uintptr(p)
}

