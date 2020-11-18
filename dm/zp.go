/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"dm/util"
	"os"
	"strconv"
	"strings"
)

const (
	Dm_build_254 = "7.6.0.0"

	Dm_build_255 = "7.0.0.9"

	Dm_build_256 = "8.0.0.73"

	Dm_build_257 = "7.1.2.128"

	Dm_build_258 = "7.1.5.144"

	Dm_build_259 = "7.1.6.123"

	Dm_build_260 = 32768 - 128

	Dm_build_261 int16 = 1

	Dm_build_262 int16 = 2

	Dm_build_263 int16 = 3

	Dm_build_264 int16 = 4

	Dm_build_265 int16 = 5

	Dm_build_266 int16 = 6

	Dm_build_267 int16 = 7

	Dm_build_268 int16 = 8

	Dm_build_269 int16 = 9

	Dm_build_270 int16 = 13

	Dm_build_271 int16 = 14

	Dm_build_272 int16 = 15

	Dm_build_273 int16 = 17

	Dm_build_274 int16 = 21

	Dm_build_275 int16 = 24

	Dm_build_276 int16 = 27

	Dm_build_277 int16 = 29

	Dm_build_278 int16 = 30

	Dm_build_279 int16 = 31

	Dm_build_280 int16 = 32

	Dm_build_281 int16 = 44

	Dm_build_282 int16 = 52

	Dm_build_283 int16 = 60

	Dm_build_284 int16 = 71

	Dm_build_285 int16 = 90

	Dm_build_286 int16 = 91

	Dm_build_287 int16 = 200

	Dm_build_288 = 64

	Dm_build_289 = 20

	Dm_build_290 = 0

	Dm_build_291 = 4

	Dm_build_292 = 6

	Dm_build_293 = 10

	Dm_build_294 = 14

	Dm_build_295 = 18

	Dm_build_296 = 19

	Dm_build_297 = 128

	Dm_build_298 = 256

	Dm_build_299 = 0xffff

	Dm_build_300 int32 = 2

	Dm_build_301 = -1

	Dm_build_302 uint16 = 0xFFFE

	Dm_build_303 uint16 = uint16(Dm_build_302 - 3)

	Dm_build_304 uint16 = Dm_build_302

	Dm_build_305 int32 = 0xFFFF

	Dm_build_306 int32 = 0x80

	Dm_build_307 byte = 0x60

	Dm_build_308 uint16 = uint16(Dm_build_304)

	Dm_build_309 uint16 = uint16(Dm_build_305)

	Dm_build_310 int16 = 0x00

	Dm_build_311 int16 = 0x03

	Dm_build_312 int32 = 0x80

	Dm_build_313 byte = 0

	Dm_build_314 byte = 1

	Dm_build_315 byte = 2

	Dm_build_316 byte = 3

	Dm_build_317 byte = 4

	Dm_build_318 byte = Dm_build_313

	Dm_build_319 int = 10

	Dm_build_320 int32 = 32

	Dm_build_321 int32 = 65536

	Dm_build_322 byte = 0

	Dm_build_323 byte = 1

	Dm_build_324 int32 = 0x00000000

	Dm_build_325 int32 = 0x00000020

	Dm_build_326 int32 = 0x00000040

	Dm_build_327 int32 = 0x00000FFF

	Dm_build_328 int32 = 0

	Dm_build_329 int32 = 1

	Dm_build_330 int32 = 2

	Dm_build_331 int32 = 3

	Dm_build_332 = 8192

	Dm_build_333 = 1

	Dm_build_334 = 2

	Dm_build_335 = 0

	Dm_build_336 = 0

	Dm_build_337 = 1

	Dm_build_338 = -1

	Dm_build_339 int16 = 0

	Dm_build_340 int16 = 1

	Dm_build_341 int16 = 2

	Dm_build_342 int16 = 3

	Dm_build_343 int16 = 4

	Dm_build_344 int16 = 127

	Dm_build_345 int16 = Dm_build_344 + 20

	Dm_build_346 int16 = Dm_build_344 + 21

	Dm_build_347 int16 = Dm_build_344 + 22

	Dm_build_348 int16 = Dm_build_344 + 24

	Dm_build_349 int16 = Dm_build_344 + 25

	Dm_build_350 int16 = Dm_build_344 + 26

	Dm_build_351 int16 = Dm_build_344 + 30

	Dm_build_352 int16 = Dm_build_344 + 31

	Dm_build_353 int16 = Dm_build_344 + 32

	Dm_build_354 int16 = Dm_build_344 + 33

	Dm_build_355 int16 = Dm_build_344 + 35

	Dm_build_356 int16 = Dm_build_344 + 38

	Dm_build_357 int16 = Dm_build_344 + 39

	Dm_build_358 int16 = Dm_build_344 + 51

	Dm_build_359 int16 = Dm_build_344 + 71

	Dm_build_360 int16 = Dm_build_344 + 124

	Dm_build_361 int16 = Dm_build_344 + 125

	Dm_build_362 int16 = Dm_build_344 + 126

	Dm_build_363 int16 = Dm_build_344 + 127

	Dm_build_364 int16 = Dm_build_344 + 128

	Dm_build_365 int16 = Dm_build_344 + 129

	Dm_build_366 byte = 0

	Dm_build_367 byte = 2

	Dm_build_368 = 2048

	Dm_build_369 = -1

	Dm_build_370 = 0

	Dm_build_371 = 16000

	Dm_build_372 = 32000

	Dm_build_373 = 0x00000000

	Dm_build_374 = 0x00000020

	Dm_build_375 = 0x00000040

	Dm_build_376 = 0x00000FFF
)

type dm_build_377 interface {
	dm_build_378()
	dm_build_379() error
	dm_build_380()
	dm_build_381(imsg dm_build_377) error
	dm_build_382() error
	dm_build_383() (interface{}, error)
	dm_build_384()
	dm_build_385(imsg dm_build_377) (interface{}, error)
	dm_build_386()
	dm_build_387() error
	dm_build_388() byte
	dm_build_389() int32
	dm_build_390(length int32)
	dm_build_391() int16
}

type dm_build_392 struct {
	dm_build_393 *dm_build_2

	dm_build_394 int16

	dm_build_395 int32

	dm_build_396 *DmStatement
}

func (dm_build_398 *dm_build_392) dm_build_397(dm_build_399 *dm_build_2, dm_build_400 int16) *dm_build_392 {
	dm_build_398.dm_build_393 = dm_build_399
	dm_build_398.dm_build_394 = dm_build_400
	return dm_build_398
}

func (dm_build_402 *dm_build_392) dm_build_401(dm_build_403 *dm_build_2, dm_build_404 int16, dm_build_405 *DmStatement) *dm_build_392 {
	dm_build_402.dm_build_397(dm_build_403, dm_build_404).dm_build_396 = dm_build_405
	return dm_build_402
}

func dm_build_406(dm_build_407 *dm_build_2, dm_build_408 int16) *dm_build_392 {
	return new(dm_build_392).dm_build_397(dm_build_407, dm_build_408)
}

func dm_build_409(dm_build_410 *dm_build_2, dm_build_411 int16, dm_build_412 *DmStatement) *dm_build_392 {
	return new(dm_build_392).dm_build_401(dm_build_410, dm_build_411, dm_build_412)
}

func (dm_build_414 *dm_build_392) dm_build_378() {
	dm_build_414.dm_build_393.dm_build_5.Dm_build_1221(0)
	dm_build_414.dm_build_393.dm_build_5.Dm_build_1232(Dm_build_288, true, true)
}

func (dm_build_416 *dm_build_392) dm_build_379() error {
	return nil
}

func (dm_build_418 *dm_build_392) dm_build_380() {
	if dm_build_418.dm_build_396 == nil {
		dm_build_418.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_290, 0)
	} else {
		dm_build_418.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_290, dm_build_418.dm_build_396.id)
	}

	dm_build_418.dm_build_393.dm_build_5.Dm_build_1394(Dm_build_291, dm_build_418.dm_build_394)
	dm_build_418.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_292, int32(dm_build_418.dm_build_393.dm_build_5.Dm_build_1219()-Dm_build_288))
}

func (dm_build_420 *dm_build_392) dm_build_382() error {
	dm_build_420.dm_build_393.dm_build_5.Dm_build_1224(0)
	dm_build_420.dm_build_393.dm_build_5.Dm_build_1232(Dm_build_288, false, true)
	return dm_build_420.dm_build_425()
}

func (dm_build_422 *dm_build_392) dm_build_383() (interface{}, error) {
	return nil, nil
}

func (dm_build_424 *dm_build_392) dm_build_384() {
}

func (dm_build_426 *dm_build_392) dm_build_425() error {
	dm_build_426.dm_build_395 = dm_build_426.dm_build_393.dm_build_5.Dm_build_1476(Dm_build_293)
	if dm_build_426.dm_build_395 < 0 && dm_build_426.dm_build_395 != EC_RN_EXCEED_ROWSET_SIZE.ErrCode {
		return (&DmError{dm_build_426.dm_build_395, dm_build_426.dm_build_427(), nil, ""}).throw()
	} else if dm_build_426.dm_build_395 > 0 {

	} else if dm_build_426.dm_build_394 == Dm_build_287 || dm_build_426.dm_build_394 == Dm_build_261 {
		dm_build_426.dm_build_427()
	}

	return nil
}

func (dm_build_428 *dm_build_392) dm_build_427() string {

	dm_build_429 := dm_build_428.dm_build_393.dm_build_6.getServerEncoding()

	if dm_build_429 != "" && dm_build_429 == ENCODING_EUCKR && Locale != LANGUAGE_EN {
		dm_build_429 = ENCODING_GB18030
	}

	dm_build_428.dm_build_393.dm_build_5.Dm_build_1232(int(dm_build_428.dm_build_393.dm_build_5.Dm_build_1332()), false, true)

	dm_build_428.dm_build_393.dm_build_5.Dm_build_1232(int(dm_build_428.dm_build_393.dm_build_5.Dm_build_1332()), false, true)

	dm_build_428.dm_build_393.dm_build_5.Dm_build_1232(int(dm_build_428.dm_build_393.dm_build_5.Dm_build_1332()), false, true)

	return dm_build_428.dm_build_393.dm_build_5.Dm_build_1374(dm_build_429, dm_build_428.dm_build_393.dm_build_6)
}

func (dm_build_431 *dm_build_392) dm_build_381(dm_build_432 dm_build_377) (dm_build_433 error) {
	dm_build_432.dm_build_378()
	if dm_build_433 = dm_build_432.dm_build_379(); dm_build_433 != nil {
		return dm_build_433
	}
	dm_build_432.dm_build_380()
	return nil
}

func (dm_build_435 *dm_build_392) dm_build_385(dm_build_436 dm_build_377) (dm_build_437 interface{}, dm_build_438 error) {
	dm_build_438 = dm_build_436.dm_build_382()
	if dm_build_438 != nil {
		return nil, dm_build_438
	}
	dm_build_437, dm_build_438 = dm_build_436.dm_build_383()
	if dm_build_438 != nil {
		return nil, dm_build_438
	}
	dm_build_436.dm_build_384()
	return dm_build_437, nil
}

func (dm_build_440 *dm_build_392) dm_build_386() {
	dm_build_440.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_296, dm_build_440.dm_build_388())
}

func (dm_build_442 *dm_build_392) dm_build_387() error {
	dm_build_443 := dm_build_442.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_296)
	dm_build_444 := dm_build_442.dm_build_388()
	if dm_build_443 != dm_build_444 {
		return ECGO_MSG_CHECK_ERROR.throw()
	}
	return nil
}

func (dm_build_446 *dm_build_392) dm_build_388() byte {
	dm_build_447 := dm_build_446.dm_build_393.dm_build_5.Dm_build_1470(0)

	for i := 1; i < Dm_build_296; i++ {
		dm_build_447 ^= dm_build_446.dm_build_393.dm_build_5.Dm_build_1470(i)
	}

	return dm_build_447
}

func (dm_build_449 *dm_build_392) dm_build_389() int32 {
	return dm_build_449.dm_build_393.dm_build_5.Dm_build_1476(Dm_build_292)
}

func (dm_build_451 *dm_build_392) dm_build_390(dm_build_452 int32) {
	dm_build_451.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_292, dm_build_452)
}

func (dm_build_454 *dm_build_392) dm_build_391() int16 {
	return dm_build_454.dm_build_394
}

type dm_build_455 struct {
	dm_build_392
}

func dm_build_456(dm_build_457 *dm_build_2) *dm_build_455 {
	dm_build_458 := new(dm_build_455)
	dm_build_458.dm_build_397(dm_build_457, Dm_build_268)
	return dm_build_458
}

type dm_build_459 struct {
	dm_build_392
	dm_build_460 string
}

func dm_build_461(dm_build_462 *dm_build_2, dm_build_463 *DmStatement, dm_build_464 string) *dm_build_459 {
	dm_build_465 := new(dm_build_459)
	dm_build_465.dm_build_401(dm_build_462, Dm_build_276, dm_build_463)
	dm_build_465.dm_build_460 = dm_build_464
	dm_build_465.dm_build_396.cursorName = dm_build_464
	return dm_build_465
}

func (dm_build_467 *dm_build_459) dm_build_379() error {
	dm_build_467.dm_build_393.dm_build_5.Dm_build_1320(dm_build_467.dm_build_460, dm_build_467.dm_build_393.dm_build_6.getServerEncoding(), dm_build_467.dm_build_393.dm_build_6)
	dm_build_467.dm_build_393.dm_build_5.Dm_build_1258(1)
	return nil
}

type Dm_build_468 struct {
	dm_build_484
	dm_build_469 []OptParameter
}

func dm_build_470(dm_build_471 *dm_build_2, dm_build_472 *DmStatement, dm_build_473 []OptParameter) *Dm_build_468 {
	dm_build_474 := new(Dm_build_468)
	dm_build_474.dm_build_401(dm_build_471, Dm_build_286, dm_build_472)
	dm_build_474.dm_build_469 = dm_build_473
	return dm_build_474
}

func (dm_build_476 *Dm_build_468) dm_build_379() error {
	dm_build_477 := len(dm_build_476.dm_build_469)

	dm_build_476.dm_build_498(int16(dm_build_477), 1)

	dm_build_476.dm_build_393.dm_build_5.Dm_build_1320(dm_build_476.dm_build_396.nativeSql, dm_build_476.dm_build_396.dmConn.getServerEncoding(), dm_build_476.dm_build_396.dmConn)

	for _, param := range dm_build_476.dm_build_469 {
		dm_build_476.dm_build_393.dm_build_5.Dm_build_1250(param.ioType)
		dm_build_476.dm_build_393.dm_build_5.Dm_build_1258(int32(param.tp))
		dm_build_476.dm_build_393.dm_build_5.Dm_build_1258(int32(param.prec))
		dm_build_476.dm_build_393.dm_build_5.Dm_build_1258(int32(param.scale))
	}

	for _, param := range dm_build_476.dm_build_469 {
		if param.bytes == nil {
			dm_build_476.dm_build_393.dm_build_5.Dm_build_1266(Dm_build_304)
		} else {
			dm_build_476.dm_build_393.dm_build_5.Dm_build_1296(param.bytes[:len(param.bytes)])
		}
	}
	return nil
}

func (dm_build_479 *Dm_build_468) dm_build_383() (interface{}, error) {
	return dm_build_479.dm_build_484.dm_build_383()
}

const (
	Dm_build_480 int = 0x01

	Dm_build_481 int = 0x02

	Dm_build_482 int = 0x04

	Dm_build_483 int = 0x08
)

type dm_build_484 struct {
	dm_build_392
	dm_build_485 [][]interface{}
	dm_build_486 []parameter
	dm_build_487 bool
}

func dm_build_488(dm_build_489 *dm_build_2, dm_build_490 int16, dm_build_491 *DmStatement) *dm_build_484 {
	dm_build_492 := new(dm_build_484)
	dm_build_492.dm_build_401(dm_build_489, dm_build_490, dm_build_491)
	dm_build_492.dm_build_487 = true
	return dm_build_492
}

func dm_build_493(dm_build_494 *dm_build_2, dm_build_495 *DmStatement, dm_build_496 [][]interface{}) *dm_build_484 {
	dm_build_497 := new(dm_build_484)

	if dm_build_494.dm_build_6.Execute2 {
		dm_build_497.dm_build_401(dm_build_494, Dm_build_270, dm_build_495)
	} else {
		dm_build_497.dm_build_401(dm_build_494, Dm_build_266, dm_build_495)
	}

	dm_build_497.dm_build_486 = dm_build_495.params
	dm_build_497.dm_build_485 = dm_build_496
	dm_build_497.dm_build_487 = true
	return dm_build_497
}

func (dm_build_499 *dm_build_484) dm_build_498(dm_build_500 int16, dm_build_501 int64) {

	dm_build_502 := Dm_build_289

	if dm_build_499.dm_build_393.dm_build_6.autoCommit {
		dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1390(dm_build_502, 1)
	} else {
		dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1390(dm_build_502, 0)
	}

	dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1394(dm_build_502, dm_build_500)

	dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1390(dm_build_502, 1)

	dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1402(dm_build_502, dm_build_501)

	dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1402(dm_build_502, dm_build_499.dm_build_396.cursorUpdateRow)

	if dm_build_499.dm_build_396.maxRows <= 0 || dm_build_499.dm_build_396.dmConn.dmConnector.enRsCache {
		dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1402(dm_build_502, INT64_MAX)
	} else {
		dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1402(dm_build_502, dm_build_499.dm_build_396.maxRows)
	}

	dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1390(dm_build_502, 1)

	if dm_build_499.dm_build_393.dm_build_6.dmConnector.continueBatchOnError {
		dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1390(dm_build_502, 1)
	} else {
		dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1390(dm_build_502, 0)
	}

	dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1390(dm_build_502, 0)

	dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1390(dm_build_502, 0)

	if dm_build_499.dm_build_396.queryTimeout == 0 {
		dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1398(dm_build_502, -1)
	} else {
		dm_build_502 += dm_build_499.dm_build_393.dm_build_5.Dm_build_1398(dm_build_502, dm_build_499.dm_build_396.queryTimeout)
	}
}

func (dm_build_504 *dm_build_484) dm_build_379() error {
	var dm_build_505 int16
	var dm_build_506 int64

	if dm_build_504.dm_build_486 != nil {
		dm_build_505 = int16(len(dm_build_504.dm_build_486))
	} else {
		dm_build_505 = 0
	}

	if dm_build_504.dm_build_485 != nil {
		dm_build_506 = int64(len(dm_build_504.dm_build_485))
	} else {
		dm_build_506 = 0
	}

	dm_build_504.dm_build_498(dm_build_505, dm_build_506)

	if dm_build_505 > 0 {
		err := dm_build_504.dm_build_507(dm_build_504.dm_build_486)
		if err != nil {
			return err
		}
		if dm_build_504.dm_build_485 != nil && len(dm_build_504.dm_build_485) > 0 {
			for _, paramObject := range dm_build_504.dm_build_485 {
				if err := dm_build_504.dm_build_510(paramObject); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (dm_build_508 *dm_build_484) dm_build_507(dm_build_509 []parameter) error {
	for _, param := range dm_build_509 {

		if param.colType == CURSOR && param.ioType == IO_TYPE_OUT {
			dm_build_508.dm_build_393.dm_build_5.Dm_build_1250(IO_TYPE_INOUT)
		} else {
			dm_build_508.dm_build_393.dm_build_5.Dm_build_1250(param.ioType)
		}

		dm_build_508.dm_build_393.dm_build_5.Dm_build_1258(param.colType)

		lprec := param.prec
		lscale := param.scale
		typeDesc := param.typeDescriptor
		switch param.colType {
		case ARRAY, SARRAY:
			tmp, err := getPackArraySize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case PLTYPE_RECORD:
			tmp, err := getPackRecordSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case CLASS:
			tmp, err := getPackClassSize(typeDesc)
			if err != nil {
				return err
			}
			lprec = int32(tmp)
		case BLOB:
			if isComplexType(int(param.colType), int(param.scale)) {
				lprec = int32(typeDesc.getObjId())
				if lprec == 4 {
					lprec = int32(typeDesc.getOuterId())
				}
			}
		}

		dm_build_508.dm_build_393.dm_build_5.Dm_build_1258(lprec)

		dm_build_508.dm_build_393.dm_build_5.Dm_build_1258(lscale)

		switch param.colType {
		case ARRAY, SARRAY:
			err := packArray(typeDesc, dm_build_508.dm_build_393.dm_build_5)
			if err != nil {
				return err
			}

		case PLTYPE_RECORD:
			err := packRecord(typeDesc, dm_build_508.dm_build_393.dm_build_5)
			if err != nil {
				return err
			}

		case CLASS:
			err := packClass(typeDesc, dm_build_508.dm_build_393.dm_build_5)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func (dm_build_511 *dm_build_484) dm_build_510(dm_build_512 []interface{}) error {
	for i := 0; i < len(dm_build_511.dm_build_486); i++ {

		if dm_build_511.dm_build_486[i].colType == CURSOR {
			dm_build_511.dm_build_393.dm_build_5.Dm_build_1254(ULINT_SIZE)
			dm_build_511.dm_build_393.dm_build_5.Dm_build_1258(dm_build_511.dm_build_486[i].cursorStmt.id)
			continue
		}

		if dm_build_511.dm_build_486[i].ioType == IO_TYPE_OUT {
			continue
		}

		switch dm_build_512[i].(type) {
		case []byte:
			if dataBytes, ok := dm_build_512[i].([]byte); ok {
				if len(dataBytes) > Dm_build_299 {
					return ECGO_DATA_TOO_LONG.throw()
				}
				dm_build_511.dm_build_393.dm_build_5.Dm_build_1296(dataBytes)
			}
		case int:
			if dm_build_512[i] == ParamDataEnum_Null {
				dm_build_511.dm_build_393.dm_build_5.Dm_build_1266(Dm_build_304)
			} else if dm_build_512[i] == ParamDataEnum_OFF_ROW {
				dm_build_511.dm_build_393.dm_build_5.Dm_build_1254(0)
			}
		case lobCtl:
			dm_build_511.dm_build_393.dm_build_5.Dm_build_1266(uint16(Dm_build_303))
			dm_build_511.dm_build_393.dm_build_5.Dm_build_1286(dm_build_512[i].(lobCtl).value)
		default:
			panic("Bind param data failed by invalid param data type: ")
		}
	}

	return nil
}

func (dm_build_514 *dm_build_484) dm_build_383() (interface{}, error) {
	dm_build_515 := execInfo{}
	dm_build_516 := dm_build_514.dm_build_396.dmConn

	dm_build_517 := Dm_build_289

	dm_build_515.retSqlType = dm_build_514.dm_build_393.dm_build_5.Dm_build_1473(dm_build_517)
	dm_build_517 += USINT_SIZE

	dm_build_518 := dm_build_514.dm_build_393.dm_build_5.Dm_build_1473(dm_build_517)
	dm_build_517 += USINT_SIZE

	dm_build_515.updateCount = dm_build_514.dm_build_393.dm_build_5.Dm_build_1479(dm_build_517)
	dm_build_517 += DDWORD_SIZE

	dm_build_519 := dm_build_514.dm_build_393.dm_build_5.Dm_build_1473(dm_build_517)
	dm_build_517 += USINT_SIZE

	dm_build_515.rsUpdatable = dm_build_514.dm_build_393.dm_build_5.Dm_build_1470(dm_build_517) != 0
	dm_build_517 += BYTE_SIZE

	dm_build_520 := dm_build_514.dm_build_393.dm_build_5.Dm_build_1473(dm_build_517)
	dm_build_517 += ULINT_SIZE

	dm_build_515.printLen = dm_build_514.dm_build_393.dm_build_5.Dm_build_1476(dm_build_517)
	dm_build_517 += ULINT_SIZE

	var dm_build_521 int16 = -1
	if dm_build_515.retSqlType == Dm_build_354 || dm_build_515.retSqlType == Dm_build_355 {
		dm_build_515.rowid = 0

		dm_build_515.rsBdta = dm_build_514.dm_build_393.dm_build_5.Dm_build_1470(dm_build_517) == Dm_build_367
		dm_build_517 += BYTE_SIZE

		dm_build_521 = dm_build_514.dm_build_393.dm_build_5.Dm_build_1473(dm_build_517)
		dm_build_517 += USINT_SIZE
		dm_build_517 += 5
	} else {
		dm_build_515.rowid = dm_build_514.dm_build_393.dm_build_5.Dm_build_1479(dm_build_517)
		dm_build_517 += DDWORD_SIZE
	}

	dm_build_515.execId = dm_build_514.dm_build_393.dm_build_5.Dm_build_1476(dm_build_517)
	dm_build_517 += ULINT_SIZE

	dm_build_515.rsCacheOffset = dm_build_514.dm_build_393.dm_build_5.Dm_build_1476(dm_build_517)
	dm_build_517 += ULINT_SIZE

	dm_build_522 := dm_build_514.dm_build_393.dm_build_5.Dm_build_1470(dm_build_517)
	dm_build_517 += BYTE_SIZE
	dm_build_523 := (dm_build_522 & 0x01) == 0x01
	dm_build_524 := (dm_build_522 & 0x02) == 0x02

	dm_build_516.TrxStatus = dm_build_514.dm_build_393.dm_build_5.Dm_build_1476(dm_build_517)
	dm_build_516.setTrxFinish(dm_build_516.TrxStatus)
	dm_build_517 += ULINT_SIZE

	if dm_build_515.printLen > 0 {
		bytes := dm_build_514.dm_build_393.dm_build_5.Dm_build_1353(int(dm_build_515.printLen))
		dm_build_515.printMsg = Dm_build_853.Dm_build_1007(bytes, 0, len(bytes), dm_build_516.getServerEncoding(), dm_build_516)
	}

	if dm_build_519 > 0 {
		dm_build_515.outParamDatas = dm_build_514.dm_build_525(int(dm_build_519))
	}

	switch dm_build_515.retSqlType {
	case Dm_build_356:
		dm_build_516.dmConnector.localTimezone = dm_build_514.dm_build_393.dm_build_5.Dm_build_1329()
	case Dm_build_354:
		dm_build_515.hasResultSet = true
		if dm_build_518 > 0 {
			dm_build_514.dm_build_396.columns = dm_build_514.dm_build_534(int(dm_build_518), dm_build_515.rsBdta)
		}
		dm_build_514.dm_build_544(&dm_build_515, len(dm_build_514.dm_build_396.columns), int(dm_build_520), int(dm_build_521))
	case Dm_build_355:
		if dm_build_518 > 0 || dm_build_520 > 0 {
			dm_build_515.hasResultSet = true
		}
		if dm_build_518 > 0 {
			dm_build_514.dm_build_396.columns = dm_build_514.dm_build_534(int(dm_build_518), dm_build_515.rsBdta)
		}
		dm_build_514.dm_build_544(&dm_build_515, len(dm_build_514.dm_build_396.columns), int(dm_build_520), int(dm_build_521))
	case Dm_build_357:
		dm_build_516.IsoLevel = int32(dm_build_514.dm_build_393.dm_build_5.Dm_build_1329())
		dm_build_516.ReadOnly = dm_build_514.dm_build_393.dm_build_5.Dm_build_1326() == 1
	case Dm_build_350:
		dm_build_516.Schema = dm_build_514.dm_build_393.dm_build_5.Dm_build_1374(dm_build_516.getServerEncoding(), dm_build_516)
	case Dm_build_347:
		dm_build_515.explain = dm_build_514.dm_build_393.dm_build_5.Dm_build_1374(dm_build_516.getServerEncoding(), dm_build_516)

	case Dm_build_351, Dm_build_353, Dm_build_352:
		if dm_build_523 {

			counts := dm_build_514.dm_build_393.dm_build_5.Dm_build_1332()
			rowCounts := make([]int64, counts)
			for i := 0; i < int(counts); i++ {
				rowCounts[i] = dm_build_514.dm_build_393.dm_build_5.Dm_build_1335()
			}
			dm_build_515.updateCounts = rowCounts
		}

		if dm_build_524 {
			rows := dm_build_514.dm_build_393.dm_build_5.Dm_build_1332()

			var lastInsertId int64
			for i := 0; i < int(rows); i++ {
				lastInsertId = dm_build_514.dm_build_393.dm_build_5.Dm_build_1335()
			}
			dm_build_515.lastInsertId = lastInsertId

		} else if dm_build_515.updateCount == 1 {
			dm_build_515.lastInsertId = dm_build_515.rowid
		}

		if dm_build_514.dm_build_395 == EC_BP_WITH_ERROR.ErrCode {
			dm_build_514.dm_build_550(dm_build_515.updateCounts)
		}
	case Dm_build_360:
		len := dm_build_514.dm_build_393.dm_build_5.Dm_build_1344()
		dm_build_516.OracleDateFormat = dm_build_514.dm_build_393.dm_build_5.Dm_build_1369(int(len), dm_build_516.getServerEncoding(), dm_build_516)
	case Dm_build_362:

		len := dm_build_514.dm_build_393.dm_build_5.Dm_build_1344()
		dm_build_516.OracleTimestampFormat = dm_build_514.dm_build_393.dm_build_5.Dm_build_1369(int(len), dm_build_516.getServerEncoding(), dm_build_516)
	case Dm_build_363:

		len := dm_build_514.dm_build_393.dm_build_5.Dm_build_1344()
		dm_build_516.OracleTimestampTZFormat = dm_build_514.dm_build_393.dm_build_5.Dm_build_1369(int(len), dm_build_516.getServerEncoding(), dm_build_516)
	case Dm_build_361:
		len := dm_build_514.dm_build_393.dm_build_5.Dm_build_1344()
		dm_build_516.OracleTimeFormat = dm_build_514.dm_build_393.dm_build_5.Dm_build_1369(int(len), dm_build_516.getServerEncoding(), dm_build_516)
	case Dm_build_364:
		len := dm_build_514.dm_build_393.dm_build_5.Dm_build_1344()
		dm_build_516.OracleTimeTZFormat = dm_build_514.dm_build_393.dm_build_5.Dm_build_1369(int(len), dm_build_516.getServerEncoding(), dm_build_516)
	case Dm_build_365:
		dm_build_516.OracleDateLanguage = dm_build_514.dm_build_393.dm_build_5.Dm_build_1344()
	}

	return &dm_build_515, nil
}

func (dm_build_526 *dm_build_484) dm_build_525(dm_build_527 int) [][]byte {
	dm_build_528 := make([]int, dm_build_527)

	dm_build_529 := 0
	for i := 0; i < len(dm_build_526.dm_build_486); i++ {
		if dm_build_526.dm_build_486[i].ioType == IO_TYPE_INOUT || dm_build_526.dm_build_486[i].ioType == IO_TYPE_OUT {
			dm_build_528[dm_build_529] = i
			dm_build_529++
		}
	}

	dm_build_530 := make([][]byte, len(dm_build_526.dm_build_486))
	var dm_build_531 int32
	var dm_build_532 bool
	var dm_build_533 []byte = nil
	for i := 0; i < dm_build_527; i++ {
		dm_build_532 = false
		dm_build_531 = int32(dm_build_526.dm_build_393.dm_build_5.Dm_build_1347())

		if dm_build_531 == int32(Dm_build_304) {
			dm_build_531 = 0
			dm_build_532 = true
		} else if dm_build_531 == int32(Dm_build_305) {
			dm_build_531 = dm_build_526.dm_build_393.dm_build_5.Dm_build_1332()
		}

		if dm_build_532 {
			dm_build_530[dm_build_528[i]] = nil
		} else {
			dm_build_533 = dm_build_526.dm_build_393.dm_build_5.Dm_build_1353(int(dm_build_531))
			dm_build_530[dm_build_528[i]] = dm_build_533
		}
	}

	return dm_build_530
}

func (dm_build_535 *dm_build_484) dm_build_534(dm_build_536 int, dm_build_537 bool) []column {
	dm_build_538 := dm_build_535.dm_build_393.dm_build_6.getServerEncoding()
	var dm_build_539, dm_build_540, dm_build_541, dm_build_542 int16
	dm_build_543 := make([]column, dm_build_536)
	for i := 0; i < dm_build_536; i++ {
		dm_build_543[i].InitColumn()

		dm_build_543[i].colType = dm_build_535.dm_build_393.dm_build_5.Dm_build_1332()

		dm_build_543[i].prec = dm_build_535.dm_build_393.dm_build_5.Dm_build_1332()

		dm_build_543[i].scale = dm_build_535.dm_build_393.dm_build_5.Dm_build_1332()

		dm_build_543[i].nullable = dm_build_535.dm_build_393.dm_build_5.Dm_build_1332() != 0

		itemFlag := dm_build_535.dm_build_393.dm_build_5.Dm_build_1329()
		dm_build_543[i].lob = int(itemFlag)&Dm_build_481 != 0
		dm_build_543[i].identity = int(itemFlag)&Dm_build_480 != 0
		dm_build_543[i].readonly = int(itemFlag)&Dm_build_482 != 0

		dm_build_535.dm_build_393.dm_build_5.Dm_build_1232(4, false, true)

		dm_build_535.dm_build_393.dm_build_5.Dm_build_1232(2, false, true)

		dm_build_539 = dm_build_535.dm_build_393.dm_build_5.Dm_build_1329()

		dm_build_540 = dm_build_535.dm_build_393.dm_build_5.Dm_build_1329()

		dm_build_541 = dm_build_535.dm_build_393.dm_build_5.Dm_build_1329()

		dm_build_542 = dm_build_535.dm_build_393.dm_build_5.Dm_build_1329()
		dm_build_543[i].name = dm_build_535.dm_build_393.dm_build_5.Dm_build_1369(int(dm_build_539), dm_build_538, dm_build_535.dm_build_393.dm_build_6)
		dm_build_543[i].typeName = dm_build_535.dm_build_393.dm_build_5.Dm_build_1369(int(dm_build_540), dm_build_538, dm_build_535.dm_build_393.dm_build_6)
		dm_build_543[i].tableName = dm_build_535.dm_build_393.dm_build_5.Dm_build_1369(int(dm_build_541), dm_build_538, dm_build_535.dm_build_393.dm_build_6)
		dm_build_543[i].schemaName = dm_build_535.dm_build_393.dm_build_5.Dm_build_1369(int(dm_build_542), dm_build_538, dm_build_535.dm_build_393.dm_build_6)

		if dm_build_535.dm_build_396.readBaseColName {
			dm_build_543[i].baseName = dm_build_535.dm_build_393.dm_build_5.Dm_build_1382(dm_build_538, dm_build_535.dm_build_393.dm_build_6)
		}

		if dm_build_543[i].lob {
			dm_build_543[i].lobTabId = dm_build_535.dm_build_393.dm_build_5.Dm_build_1332()
			dm_build_543[i].lobColId = dm_build_535.dm_build_393.dm_build_5.Dm_build_1329()
		}

	}

	for i := 0; i < dm_build_536; i++ {

		if isComplexType(int(dm_build_543[i].colType), int(dm_build_543[i].scale)) {
			strDesc := newTypeDescriptor(dm_build_535.dm_build_393.dm_build_6)
			strDesc.unpack(dm_build_535.dm_build_393.dm_build_5)
			dm_build_543[i].typeDescriptor = strDesc
		}
	}

	return dm_build_543
}

func (dm_build_545 *dm_build_484) dm_build_544(dm_build_546 *execInfo, dm_build_547 int, dm_build_548 int, dm_build_549 int) {
	if dm_build_548 > 0 {
		startOffset := dm_build_545.dm_build_393.dm_build_5.Dm_build_1227()
		if dm_build_546.rsBdta {
			dm_build_546.rsDatas = dm_build_545.dm_build_563(dm_build_545.dm_build_396.columns, dm_build_549)
		} else {
			datas := make([][][]byte, dm_build_548)

			for i := 0; i < dm_build_548; i++ {

				datas[i] = make([][]byte, dm_build_547+1)

				dm_build_545.dm_build_393.dm_build_5.Dm_build_1232(2, false, true)

				datas[i][0] = dm_build_545.dm_build_393.dm_build_5.Dm_build_1353(LINT64_SIZE)

				dm_build_545.dm_build_393.dm_build_5.Dm_build_1232(2*dm_build_547, false, true)

				for j := 1; j < dm_build_547+1; j++ {

					colLen := dm_build_545.dm_build_393.dm_build_5.Dm_build_1347()
					if colLen == Dm_build_308 {
						datas[i][j] = nil
					} else if colLen != Dm_build_309 {
						datas[i][j] = dm_build_545.dm_build_393.dm_build_5.Dm_build_1353(int(colLen))
					} else {
						datas[i][j] = dm_build_545.dm_build_393.dm_build_5.Dm_build_1357()
					}
				}
			}

			dm_build_546.rsDatas = datas
		}
		dm_build_546.rsSizeof = dm_build_545.dm_build_393.dm_build_5.Dm_build_1227() - startOffset
	}

	if dm_build_546.rsCacheOffset > 0 {
		tbCount := dm_build_545.dm_build_393.dm_build_5.Dm_build_1329()

		ids := make([]int32, tbCount)
		tss := make([]int64, tbCount)

		for i := 0; i < int(tbCount); i++ {
			ids[i] = dm_build_545.dm_build_393.dm_build_5.Dm_build_1332()
			tss[i] = dm_build_545.dm_build_393.dm_build_5.Dm_build_1335()
		}

		dm_build_546.tbIds = ids[:]
		dm_build_546.tbTss = tss[:]
	}
}

func (dm_build_551 *dm_build_484) dm_build_550(dm_build_552 []int64) error {

	dm_build_551.dm_build_393.dm_build_5.Dm_build_1232(4, false, true)

	dm_build_553 := dm_build_551.dm_build_393.dm_build_5.Dm_build_1332()

	dm_build_554 := make([]string, 0, 8)
	for i := 0; i < int(dm_build_553); i++ {
		irow := dm_build_551.dm_build_393.dm_build_5.Dm_build_1332()

		dm_build_552[irow] = -3

		code := dm_build_551.dm_build_393.dm_build_5.Dm_build_1332()

		errStr := dm_build_551.dm_build_393.dm_build_5.Dm_build_1382(dm_build_551.dm_build_393.dm_build_6.getServerEncoding(), dm_build_551.dm_build_393.dm_build_6)

		dm_build_554 = append(dm_build_554, "row["+strconv.Itoa(int(irow))+"]:"+strconv.Itoa(int(code))+", "+errStr)
	}

	if len(dm_build_554) > 0 {
		builder := &strings.Builder{}
		for _, str := range dm_build_554 {
			builder.WriteString(util.LINE_SEPARATOR)
			builder.WriteString(str)
		}
		EC_BP_WITH_ERROR.ErrText += builder.String()
		return EC_BP_WITH_ERROR.throw()
	}
	return nil
}

const (
	Dm_build_555 = 0

	Dm_build_556 = Dm_build_555 + ULINT_SIZE

	Dm_build_557 = Dm_build_556 + USINT_SIZE

	Dm_build_558 = Dm_build_557 + ULINT_SIZE

	Dm_build_559 = Dm_build_558 + ULINT_SIZE

	Dm_build_560 = Dm_build_559 + BYTE_SIZE

	Dm_build_561 = -2

	Dm_build_562 = -3
)

func (dm_build_564 *dm_build_484) dm_build_563(dm_build_565 []column, dm_build_566 int) [][][]byte {

	dm_build_567 := dm_build_564.dm_build_393.dm_build_5.Dm_build_1350()

	dm_build_568 := dm_build_564.dm_build_393.dm_build_5.Dm_build_1347()

	var dm_build_569 bool

	if dm_build_566 >= 0 && int(dm_build_568) == len(dm_build_565)+1 {
		dm_build_569 = true
	} else {
		dm_build_569 = false
	}

	dm_build_564.dm_build_393.dm_build_5.Dm_build_1232(ULINT_SIZE, false, true)

	dm_build_564.dm_build_393.dm_build_5.Dm_build_1232(ULINT_SIZE, false, true)

	dm_build_564.dm_build_393.dm_build_5.Dm_build_1232(BYTE_SIZE, false, true)

	dm_build_570 := make([]uint16, dm_build_568)
	for icol := 0; icol < int(dm_build_568); icol++ {
		dm_build_570[icol] = dm_build_564.dm_build_393.dm_build_5.Dm_build_1347()
	}

	dm_build_571 := make([]uint32, dm_build_568)
	dm_build_572 := make([][][]byte, dm_build_567)

	for i := uint32(0); i < dm_build_567; i++ {
		dm_build_572[i] = make([][]byte, len(dm_build_565)+1)
	}

	for icol := 0; icol < int(dm_build_568); icol++ {
		dm_build_571[icol] = dm_build_564.dm_build_393.dm_build_5.Dm_build_1350()
	}

	for icol := 0; icol < int(dm_build_568); icol++ {

		dataCol := icol + 1
		if dm_build_569 && icol == dm_build_566 {
			dataCol = 0
		} else if dm_build_569 && icol > dm_build_566 {
			dataCol = icol
		}

		allNotNull := dm_build_564.dm_build_393.dm_build_5.Dm_build_1332() == 1
		var isNull []bool = nil
		if !allNotNull {
			isNull = make([]bool, dm_build_567)
			for irow := uint32(0); irow < dm_build_567; irow++ {
				isNull[irow] = dm_build_564.dm_build_393.dm_build_5.Dm_build_1326() == 0
			}
		}

		for irow := uint32(0); irow < dm_build_567; irow++ {
			if allNotNull || !isNull[irow] {
				dm_build_572[irow][dataCol] = dm_build_564.dm_build_573(int(dm_build_570[icol]))
			}
		}
	}

	if !dm_build_569 && dm_build_566 >= 0 {
		for irow := uint32(0); irow < dm_build_567; irow++ {
			dm_build_572[irow][0] = dm_build_572[irow][dm_build_566+1]
		}
	}

	return dm_build_572
}

func (dm_build_574 *dm_build_484) dm_build_573(dm_build_575 int) []byte {

	dm_build_576 := dm_build_574.dm_build_579(dm_build_575)

	dm_build_577 := int32(0)
	if dm_build_576 == Dm_build_561 {
		dm_build_577 = dm_build_574.dm_build_393.dm_build_5.Dm_build_1332()
		dm_build_576 = int(dm_build_574.dm_build_393.dm_build_5.Dm_build_1332())
	} else if dm_build_576 == Dm_build_562 {
		dm_build_576 = int(dm_build_574.dm_build_393.dm_build_5.Dm_build_1332())
	}

	dm_build_578 := dm_build_574.dm_build_393.dm_build_5.Dm_build_1353(dm_build_576 + int(dm_build_577))
	if dm_build_577 == 0 {
		return dm_build_578
	}

	for i := dm_build_576; i < len(dm_build_578); i++ {
		dm_build_578[i] = ' '
	}
	return dm_build_578
}

func (dm_build_580 *dm_build_484) dm_build_579(dm_build_581 int) int {

	dm_build_582 := 0
	switch dm_build_581 {
	case INT:
	case BIT:
	case TINYINT:
	case SMALLINT:
	case BOOLEAN:
	case NULL:
		dm_build_582 = 4

	case BIGINT:

		dm_build_582 = 8

	case CHAR:
	case VARCHAR2:
	case VARCHAR:
	case BINARY:
	case VARBINARY:
	case BLOB:
	case CLOB:
		dm_build_582 = Dm_build_561

	case DECIMAL:
		dm_build_582 = Dm_build_562

	case REAL:
		dm_build_582 = 4

	case DOUBLE:
		dm_build_582 = 8

	case DATE:
	case TIME:
	case DATETIME:
	case TIME_TZ:
	case DATETIME_TZ:
		dm_build_582 = 12

	case INTERVAL_YM:
		dm_build_582 = 12

	case INTERVAL_DT:
		dm_build_582 = 24

	default:
		panic(ECGO_INVALID_COLUMN_TYPE)
	}
	return dm_build_582
}

const (
	Dm_build_583 = Dm_build_289

	Dm_build_584 = Dm_build_583 + DDWORD_SIZE

	Dm_build_585 = Dm_build_584 + LINT64_SIZE

	Dm_build_586 = Dm_build_585 + USINT_SIZE

	Dm_build_587 = Dm_build_289

	Dm_build_588 = Dm_build_587 + DDWORD_SIZE
)

type dm_build_589 struct {
	dm_build_484
	dm_build_590 *innerRows
	dm_build_591 int64
	dm_build_592 int64
}

func dm_build_593(dm_build_594 *dm_build_2, dm_build_595 *innerRows, dm_build_596 int64, dm_build_597 int64) *dm_build_589 {
	dm_build_598 := new(dm_build_589)
	dm_build_598.dm_build_401(dm_build_594, Dm_build_267, dm_build_595.dmStmt)
	dm_build_598.dm_build_590 = dm_build_595
	dm_build_598.dm_build_591 = dm_build_596
	dm_build_598.dm_build_592 = dm_build_597
	return dm_build_598
}

func (dm_build_600 *dm_build_589) dm_build_379() error {

	dm_build_600.dm_build_393.dm_build_5.Dm_build_1402(Dm_build_583, dm_build_600.dm_build_591)

	dm_build_600.dm_build_393.dm_build_5.Dm_build_1402(Dm_build_584, dm_build_600.dm_build_592)

	dm_build_600.dm_build_393.dm_build_5.Dm_build_1394(Dm_build_585, dm_build_600.dm_build_590.id)

	dm_build_601 := dm_build_600.dm_build_590.dmStmt.dmConn.dmConnector.bufPrefetch
	var dm_build_602 int32
	if dm_build_600.dm_build_590.sizeOfRow != 0 && dm_build_600.dm_build_590.fetchSize != 0 {
		if dm_build_600.dm_build_590.sizeOfRow*dm_build_600.dm_build_590.fetchSize > int(INT32_MAX) {
			dm_build_602 = INT32_MAX
		} else {
			dm_build_602 = int32(dm_build_600.dm_build_590.sizeOfRow * dm_build_600.dm_build_590.fetchSize)
		}

		if dm_build_602 < Dm_build_320 {
			dm_build_601 = int(Dm_build_320)
		} else if dm_build_602 > Dm_build_321 {
			dm_build_601 = int(Dm_build_321)
		} else {
			dm_build_601 = int(dm_build_602)
		}

		dm_build_600.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_586, int32(dm_build_601))
	}

	return nil
}

func (dm_build_604 *dm_build_589) dm_build_383() (interface{}, error) {
	dm_build_605 := execInfo{}
	dm_build_605.rsBdta = dm_build_604.dm_build_590.isBdta

	dm_build_605.updateCount = dm_build_604.dm_build_393.dm_build_5.Dm_build_1479(Dm_build_587)
	dm_build_606 := dm_build_604.dm_build_393.dm_build_5.Dm_build_1476(Dm_build_588)

	dm_build_604.dm_build_544(&dm_build_605, len(dm_build_604.dm_build_590.columns), int(dm_build_606), -1)

	return &dm_build_605, nil
}

type dm_build_607 struct {
	dm_build_392
	dm_build_608 *lob
	dm_build_609 int
	dm_build_610 int
}

func dm_build_611(dm_build_612 *dm_build_2, dm_build_613 *lob, dm_build_614 int, dm_build_615 int) *dm_build_607 {
	dm_build_616 := new(dm_build_607)
	dm_build_616.dm_build_397(dm_build_612, Dm_build_280)
	dm_build_616.dm_build_608 = dm_build_613
	dm_build_616.dm_build_609 = dm_build_614
	dm_build_616.dm_build_610 = dm_build_615
	return dm_build_616
}

func (dm_build_618 *dm_build_607) dm_build_379() error {

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1250(byte(dm_build_618.dm_build_608.lobFlag))

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1258(dm_build_618.dm_build_608.tabId)

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1254(dm_build_618.dm_build_608.colId)

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1274(uint64(dm_build_618.dm_build_608.blobId))

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1254(dm_build_618.dm_build_608.groupId)

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1254(dm_build_618.dm_build_608.fileId)

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1258(dm_build_618.dm_build_608.pageNo)

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1254(dm_build_618.dm_build_608.curFileId)

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1258(dm_build_618.dm_build_608.curPageNo)

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1258(dm_build_618.dm_build_608.totalOffset)

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1258(int32(dm_build_618.dm_build_609))

	dm_build_618.dm_build_393.dm_build_5.Dm_build_1258(int32(dm_build_618.dm_build_610))

	if dm_build_618.dm_build_393.dm_build_6.NewLobFlag {
		dm_build_618.dm_build_393.dm_build_5.Dm_build_1274(uint64(dm_build_618.dm_build_608.rowId))
		dm_build_618.dm_build_393.dm_build_5.Dm_build_1254(dm_build_618.dm_build_608.exGroupId)
		dm_build_618.dm_build_393.dm_build_5.Dm_build_1254(dm_build_618.dm_build_608.exFileId)
		dm_build_618.dm_build_393.dm_build_5.Dm_build_1258(dm_build_618.dm_build_608.exPageNo)
	}

	return nil
}

func (dm_build_620 *dm_build_607) dm_build_383() (interface{}, error) {

	dm_build_620.dm_build_608.readOver = dm_build_620.dm_build_393.dm_build_5.Dm_build_1326() == 1
	var dm_build_621 = dm_build_620.dm_build_393.dm_build_5.Dm_build_1350()
	if dm_build_621 <= 0 {
		return []byte(nil), nil
	}
	dm_build_620.dm_build_608.curFileId = dm_build_620.dm_build_393.dm_build_5.Dm_build_1329()
	dm_build_620.dm_build_608.curPageNo = dm_build_620.dm_build_393.dm_build_5.Dm_build_1332()
	dm_build_620.dm_build_608.totalOffset = dm_build_620.dm_build_393.dm_build_5.Dm_build_1332()

	return dm_build_620.dm_build_393.dm_build_5.Dm_build_1363(int(dm_build_621)), nil
}

type dm_build_622 struct {
	dm_build_392
	dm_build_623 *lob
}

func dm_build_624(dm_build_625 *dm_build_2, dm_build_626 *lob) *dm_build_622 {
	dm_build_627 := new(dm_build_622)
	dm_build_627.dm_build_397(dm_build_625, Dm_build_277)
	dm_build_627.dm_build_623 = dm_build_626
	return dm_build_627
}

func (dm_build_629 *dm_build_622) dm_build_379() error {

	dm_build_629.dm_build_393.dm_build_5.Dm_build_1250(byte(dm_build_629.dm_build_623.lobFlag))

	dm_build_629.dm_build_393.dm_build_5.Dm_build_1274(uint64(dm_build_629.dm_build_623.blobId))

	dm_build_629.dm_build_393.dm_build_5.Dm_build_1254(dm_build_629.dm_build_623.groupId)

	dm_build_629.dm_build_393.dm_build_5.Dm_build_1254(dm_build_629.dm_build_623.fileId)

	dm_build_629.dm_build_393.dm_build_5.Dm_build_1258(dm_build_629.dm_build_623.pageNo)

	if dm_build_629.dm_build_393.dm_build_6.NewLobFlag {
		dm_build_629.dm_build_393.dm_build_5.Dm_build_1258(dm_build_629.dm_build_623.tabId)
		dm_build_629.dm_build_393.dm_build_5.Dm_build_1254(dm_build_629.dm_build_623.colId)
		dm_build_629.dm_build_393.dm_build_5.Dm_build_1274(uint64(dm_build_629.dm_build_623.rowId))

		dm_build_629.dm_build_393.dm_build_5.Dm_build_1254(dm_build_629.dm_build_623.exGroupId)
		dm_build_629.dm_build_393.dm_build_5.Dm_build_1254(dm_build_629.dm_build_623.exFileId)
		dm_build_629.dm_build_393.dm_build_5.Dm_build_1258(dm_build_629.dm_build_623.exPageNo)
	}

	return nil
}

func (dm_build_631 *dm_build_622) dm_build_383() (interface{}, error) {

	return int64(dm_build_631.dm_build_393.dm_build_5.Dm_build_1332()), nil
}

type dm_build_632 struct {
	dm_build_392
	dm_build_633 *lob
	dm_build_634 int
}

func dm_build_635(dm_build_636 *dm_build_2, dm_build_637 *lob, dm_build_638 int) *dm_build_632 {
	dm_build_639 := new(dm_build_632)
	dm_build_639.dm_build_397(dm_build_636, Dm_build_279)
	dm_build_639.dm_build_633 = dm_build_637
	dm_build_639.dm_build_634 = dm_build_638
	return dm_build_639
}

func (dm_build_641 *dm_build_632) dm_build_379() error {

	dm_build_641.dm_build_393.dm_build_5.Dm_build_1250(byte(dm_build_641.dm_build_633.lobFlag))

	dm_build_641.dm_build_393.dm_build_5.Dm_build_1274(uint64(dm_build_641.dm_build_633.blobId))

	dm_build_641.dm_build_393.dm_build_5.Dm_build_1254(dm_build_641.dm_build_633.groupId)

	dm_build_641.dm_build_393.dm_build_5.Dm_build_1254(dm_build_641.dm_build_633.fileId)

	dm_build_641.dm_build_393.dm_build_5.Dm_build_1258(dm_build_641.dm_build_633.pageNo)

	dm_build_641.dm_build_393.dm_build_5.Dm_build_1258(dm_build_641.dm_build_633.tabId)
	dm_build_641.dm_build_393.dm_build_5.Dm_build_1254(dm_build_641.dm_build_633.colId)
	dm_build_641.dm_build_393.dm_build_5.Dm_build_1274(uint64(dm_build_641.dm_build_633.rowId))
	dm_build_641.dm_build_393.dm_build_5.Dm_build_1286(Dm_build_853.Dm_build_1052(uint32(dm_build_641.dm_build_634)))

	if dm_build_641.dm_build_393.dm_build_6.NewLobFlag {
		dm_build_641.dm_build_393.dm_build_5.Dm_build_1254(dm_build_641.dm_build_633.exGroupId)
		dm_build_641.dm_build_393.dm_build_5.Dm_build_1254(dm_build_641.dm_build_633.exFileId)
		dm_build_641.dm_build_393.dm_build_5.Dm_build_1258(dm_build_641.dm_build_633.exPageNo)
	}
	return nil
}

func (dm_build_643 *dm_build_632) dm_build_383() (interface{}, error) {

	dm_build_644 := dm_build_643.dm_build_393.dm_build_5.Dm_build_1350()
	dm_build_643.dm_build_633.blobId = dm_build_643.dm_build_393.dm_build_5.Dm_build_1335()
	dm_build_643.dm_build_633.resetCurrentInfo()
	return int64(dm_build_644), nil
}

const (
	Dm_build_645 = Dm_build_289

	Dm_build_646 = Dm_build_645 + ULINT_SIZE

	Dm_build_647 = Dm_build_646 + ULINT_SIZE

	Dm_build_648 = Dm_build_647 + ULINT_SIZE

	Dm_build_649 = Dm_build_648 + BYTE_SIZE

	Dm_build_650 = Dm_build_649 + USINT_SIZE

	Dm_build_651 = Dm_build_650 + ULINT_SIZE

	Dm_build_652 = Dm_build_651 + BYTE_SIZE

	Dm_build_653 = Dm_build_652 + BYTE_SIZE

	Dm_build_654 = Dm_build_653 + BYTE_SIZE

	Dm_build_655 = Dm_build_289

	Dm_build_656 = Dm_build_655 + ULINT_SIZE

	Dm_build_657 = Dm_build_656 + ULINT_SIZE

	Dm_build_658 = Dm_build_657 + BYTE_SIZE

	Dm_build_659 = Dm_build_658 + ULINT_SIZE

	Dm_build_660 = Dm_build_659 + BYTE_SIZE

	Dm_build_661 = Dm_build_660 + BYTE_SIZE

	Dm_build_662 = Dm_build_661 + USINT_SIZE

	Dm_build_663 = Dm_build_662 + USINT_SIZE

	Dm_build_664 = Dm_build_663 + BYTE_SIZE

	Dm_build_665 = Dm_build_664 + USINT_SIZE

	Dm_build_666 = Dm_build_665 + BYTE_SIZE

	Dm_build_667 = Dm_build_666 + BYTE_SIZE

	Dm_build_668 = Dm_build_667 + ULINT_SIZE
)

type dm_build_669 struct {
	dm_build_392

	dm_build_670 *DmConnection

	dm_build_671 bool
}

func dm_build_672(dm_build_673 *dm_build_2) *dm_build_669 {
	dm_build_674 := new(dm_build_669)
	dm_build_674.dm_build_397(dm_build_673, Dm_build_261)
	dm_build_674.dm_build_670 = dm_build_673.dm_build_6
	return dm_build_674
}

func (dm_build_676 *dm_build_669) dm_build_379() error {

	dm_build_676.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_645, Dm_build_300)

	dm_build_676.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_646, g2dbIsoLevel(dm_build_676.dm_build_670.IsoLevel))
	dm_build_676.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_647, int32(Locale))
	dm_build_676.dm_build_393.dm_build_5.Dm_build_1394(Dm_build_649, dm_build_676.dm_build_670.dmConnector.localTimezone)

	if dm_build_676.dm_build_670.ReadOnly {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_648, Dm_build_323)
	} else {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_648, Dm_build_322)
	}

	dm_build_676.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_650, int32(dm_build_676.dm_build_670.dmConnector.sessionTimeout))

	if dm_build_676.dm_build_670.dmConnector.mppLocal {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_651, 1)
	} else {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_651, 0)
	}

	if dm_build_676.dm_build_670.dmConnector.rwSeparate {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_652, 1)
	} else {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_652, 0)
	}

	if dm_build_676.dm_build_670.NewLobFlag {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_653, 1)
	} else {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_653, 0)
	}

	dm_build_676.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_654, dm_build_676.dm_build_670.dmConnector.osAuthType)

	dm_build_677 := dm_build_676.dm_build_670.getServerEncoding()

	if dm_build_676.dm_build_393.dm_build_11 != "" {

	}

	dm_build_678 := Dm_build_853.Dm_build_1063(dm_build_676.dm_build_670.dmConnector.user, dm_build_677, dm_build_676.dm_build_393.dm_build_6)
	dm_build_679 := Dm_build_853.Dm_build_1063(dm_build_676.dm_build_670.dmConnector.password, dm_build_677, dm_build_676.dm_build_393.dm_build_6)
	if len(dm_build_678) > Dm_build_297 {
		return ECGO_USERNAME_TOO_LONG.throw()
	}
	if len(dm_build_679) > Dm_build_297 {
		return ECGO_PASSWORD_TOO_LONG.throw()
	}

	if dm_build_676.dm_build_393.dm_build_8 && dm_build_676.dm_build_670.dmConnector.loginCertificate != "" {

	} else if dm_build_676.dm_build_393.dm_build_8 {
		dm_build_678 = dm_build_676.dm_build_393.dm_build_7.Encrypt(dm_build_678, false)
		dm_build_679 = dm_build_676.dm_build_393.dm_build_7.Encrypt(dm_build_679, false)
	}

	dm_build_676.dm_build_393.dm_build_5.Dm_build_1290(dm_build_678)
	dm_build_676.dm_build_393.dm_build_5.Dm_build_1290(dm_build_679)

	dm_build_676.dm_build_393.dm_build_5.Dm_build_1302(dm_build_676.dm_build_670.dmConnector.appName, dm_build_677, dm_build_676.dm_build_393.dm_build_6)
	dm_build_676.dm_build_393.dm_build_5.Dm_build_1302(dm_build_676.dm_build_670.dmConnector.osName, dm_build_677, dm_build_676.dm_build_393.dm_build_6)

	if hostName, err := os.Hostname(); err != nil {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1302(hostName, dm_build_677, dm_build_676.dm_build_393.dm_build_6)
	} else {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1302("", dm_build_677, dm_build_676.dm_build_393.dm_build_6)
	}

	if dm_build_676.dm_build_670.dmConnector.rwStandby {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1250(1)
	} else {
		dm_build_676.dm_build_393.dm_build_5.Dm_build_1250(0)
	}

	return nil
}

func (dm_build_681 *dm_build_669) dm_build_383() (interface{}, error) {

	dm_build_681.dm_build_670.MaxRowSize = dm_build_681.dm_build_393.dm_build_5.Dm_build_1476(Dm_build_655)
	dm_build_681.dm_build_670.DDLAutoCommit = dm_build_681.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_657) == 1
	dm_build_681.dm_build_670.IsoLevel = dm_build_681.dm_build_393.dm_build_5.Dm_build_1476(Dm_build_658)
	dm_build_681.dm_build_670.dmConnector.caseSensitive = dm_build_681.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_659) == 1
	dm_build_681.dm_build_670.BackslashEscape = dm_build_681.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_660) == 1
	dm_build_681.dm_build_670.SvrStat = dm_build_681.dm_build_393.dm_build_5.Dm_build_1473(Dm_build_662)
	dm_build_681.dm_build_670.SvrMode = dm_build_681.dm_build_393.dm_build_5.Dm_build_1473(Dm_build_661)
	dm_build_681.dm_build_670.ConstParaOpt = dm_build_681.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_663) == 1
	dm_build_681.dm_build_670.DbTimezone = dm_build_681.dm_build_393.dm_build_5.Dm_build_1473(Dm_build_664)
	dm_build_681.dm_build_670.NewLobFlag = dm_build_681.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_666) == 1

	if dm_build_681.dm_build_670.dmConnector.bufPrefetch == 0 {
		dm_build_681.dm_build_670.dmConnector.bufPrefetch = int(dm_build_681.dm_build_393.dm_build_5.Dm_build_1476(Dm_build_667))
	}

	dm_build_681.dm_build_670.LifeTimeRemainder = dm_build_681.dm_build_393.dm_build_5.Dm_build_1473(Dm_build_668)

	dm_build_682 := dm_build_681.dm_build_670.getServerEncoding()

	dm_build_681.dm_build_670.InstanceName = dm_build_681.dm_build_393.dm_build_5.Dm_build_1374(dm_build_682, dm_build_681.dm_build_393.dm_build_6)
	dm_build_681.dm_build_670.Schema = dm_build_681.dm_build_393.dm_build_5.Dm_build_1374(dm_build_682, dm_build_681.dm_build_393.dm_build_6)
	dm_build_681.dm_build_670.LastLoginIP = dm_build_681.dm_build_393.dm_build_5.Dm_build_1374(dm_build_682, dm_build_681.dm_build_393.dm_build_6)
	dm_build_681.dm_build_670.LastLoginTime = dm_build_681.dm_build_393.dm_build_5.Dm_build_1374(dm_build_682, dm_build_681.dm_build_393.dm_build_6)
	dm_build_681.dm_build_670.FailedAttempts = dm_build_681.dm_build_393.dm_build_5.Dm_build_1332()
	dm_build_681.dm_build_670.LoginWarningID = dm_build_681.dm_build_393.dm_build_5.Dm_build_1332()
	dm_build_681.dm_build_670.GraceTimeRemainder = dm_build_681.dm_build_393.dm_build_5.Dm_build_1332()
	dm_build_681.dm_build_670.Guid = dm_build_681.dm_build_393.dm_build_5.Dm_build_1374(dm_build_682, dm_build_681.dm_build_393.dm_build_6)
	dm_build_681.dm_build_670.DbName = dm_build_681.dm_build_393.dm_build_5.Dm_build_1374(dm_build_682, dm_build_681.dm_build_393.dm_build_6)

	if dm_build_681.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_665) == 1 {
		dm_build_681.dm_build_670.StandbyHost = dm_build_681.dm_build_393.dm_build_5.Dm_build_1374(dm_build_682, dm_build_681.dm_build_393.dm_build_6)
		dm_build_681.dm_build_670.StandbyPort = dm_build_681.dm_build_393.dm_build_5.Dm_build_1332()
		dm_build_681.dm_build_670.StandbyCount = dm_build_681.dm_build_393.dm_build_5.Dm_build_1347()
	}

	if dm_build_681.dm_build_393.dm_build_5.Dm_build_1229(false) > 0 {
		dm_build_681.dm_build_670.SessionID = dm_build_681.dm_build_393.dm_build_5.Dm_build_1335()
	}

	if dm_build_681.dm_build_393.dm_build_5.Dm_build_1229(false) > 0 {
		if dm_build_681.dm_build_393.dm_build_5.Dm_build_1326() == 1 {

			dm_build_681.dm_build_670.OracleDateFormat = "DD-MON-YY"

			dm_build_681.dm_build_670.OracleTimestampFormat = "DD-MON-YY HH12.MI.SS.FF6 AM"

			dm_build_681.dm_build_670.OracleTimestampTZFormat = "DD-MON-YY HH12.MI.SS.FF6 AM +TZH:TZM"

			dm_build_681.dm_build_670.OracleTimeFormat = "HH12.MI.SS.FF6 AM"

			dm_build_681.dm_build_670.OracleTimeTZFormat = "HH12.MI.SS.FF6 AM +TZH:TZM"
		}
	}

	return nil, nil
}

const (
	Dm_build_683 = Dm_build_289
)

type dm_build_684 struct {
	dm_build_484
	dm_build_685 int16
}

func dm_build_686(dm_build_687 *dm_build_2, dm_build_688 *DmStatement, dm_build_689 int16) *dm_build_684 {
	dm_build_690 := new(dm_build_684)
	dm_build_690.dm_build_401(dm_build_687, Dm_build_281, dm_build_688)
	dm_build_690.dm_build_685 = dm_build_689
	return dm_build_690
}

func (dm_build_692 *dm_build_684) dm_build_379() error {
	dm_build_692.dm_build_393.dm_build_5.Dm_build_1394(Dm_build_683, dm_build_692.dm_build_685)
	return nil
}

func (dm_build_694 *dm_build_684) dm_build_383() (interface{}, error) {
	return dm_build_694.dm_build_484.dm_build_383()
}

const (
	Dm_build_695 = Dm_build_289
)

type dm_build_696 struct {
	dm_build_484
	dm_build_697 []parameter
}

func dm_build_698(dm_build_699 *dm_build_2, dm_build_700 *DmStatement, dm_build_701 []parameter) *dm_build_696 {
	dm_build_702 := new(dm_build_696)
	dm_build_702.dm_build_401(dm_build_699, Dm_build_285, dm_build_700)
	dm_build_702.dm_build_697 = dm_build_701
	return dm_build_702
}

func (dm_build_704 *dm_build_696) dm_build_379() error {

	if dm_build_704.dm_build_697 == nil {
		dm_build_704.dm_build_393.dm_build_5.Dm_build_1394(Dm_build_695, 0)
	} else {
		dm_build_704.dm_build_393.dm_build_5.Dm_build_1394(Dm_build_695, int16(len(dm_build_704.dm_build_697)))
	}

	return dm_build_704.dm_build_507(dm_build_704.dm_build_697)
}

type dm_build_705 struct {
	dm_build_484
	dm_build_706 bool
	dm_build_707 int16
}

func dm_build_708(dm_build_709 *dm_build_2, dm_build_710 *DmStatement, dm_build_711 bool, dm_build_712 int16) *dm_build_705 {
	dm_build_713 := new(dm_build_705)
	dm_build_713.dm_build_401(dm_build_709, Dm_build_265, dm_build_710)
	dm_build_713.dm_build_706 = dm_build_711
	dm_build_713.dm_build_707 = dm_build_712
	return dm_build_713
}

func (dm_build_715 *dm_build_705) dm_build_379() error {

	dm_build_716 := Dm_build_289

	if dm_build_715.dm_build_393.dm_build_6.autoCommit {
		dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, 1)
	} else {
		dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, 0)
	}

	if dm_build_715.dm_build_706 {
		dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, 1)
	} else {
		dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, 0)
	}

	dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, 0)

	dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, 1)

	if dm_build_715.dm_build_393.dm_build_6.CompatibleOracle() {
		dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, 0)
	} else {
		dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, 1)
	}

	dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1394(dm_build_716, dm_build_715.dm_build_707)

	if dm_build_715.dm_build_396.maxRows <= 0 || dm_build_715.dm_build_393.dm_build_6.dmConnector.enRsCache {
		dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1402(dm_build_716, INT64_MAX)
	} else {
		dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1402(dm_build_716, dm_build_715.dm_build_396.maxRows)
	}

	if dm_build_715.dm_build_393.dm_build_6.dmConnector.isBdtaRS {
		dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, Dm_build_367)
	} else {
		dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, Dm_build_366)
	}

	dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1394(dm_build_716, 0)

	dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, 1)

	dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, 0)

	dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1390(dm_build_716, 0)

	dm_build_716 += dm_build_715.dm_build_393.dm_build_5.Dm_build_1398(dm_build_716, dm_build_715.dm_build_396.queryTimeout)

	dm_build_715.dm_build_393.dm_build_5.Dm_build_1320(dm_build_715.dm_build_396.nativeSql, dm_build_715.dm_build_393.dm_build_6.getServerEncoding(), dm_build_715.dm_build_393.dm_build_6)

	return nil
}

func (dm_build_718 *dm_build_705) dm_build_383() (interface{}, error) {

	if dm_build_718.dm_build_706 {
		return dm_build_718.dm_build_484.dm_build_383()
	}

	dm_build_719 := NewExceInfo()
	dm_build_720 := Dm_build_289

	dm_build_719.retSqlType = dm_build_718.dm_build_393.dm_build_5.Dm_build_1473(dm_build_720)
	dm_build_720 += USINT_SIZE

	dm_build_721 := dm_build_718.dm_build_393.dm_build_5.Dm_build_1473(dm_build_720)
	dm_build_720 += USINT_SIZE

	dm_build_722 := dm_build_718.dm_build_393.dm_build_5.Dm_build_1473(dm_build_720)
	dm_build_720 += USINT_SIZE

	dm_build_718.dm_build_393.dm_build_5.Dm_build_1479(dm_build_720)
	dm_build_720 += DDWORD_SIZE

	dm_build_718.dm_build_393.dm_build_6.TrxStatus = dm_build_718.dm_build_393.dm_build_5.Dm_build_1476(dm_build_720)
	dm_build_720 += ULINT_SIZE

	if dm_build_721 > 0 {
		dm_build_718.dm_build_396.params = dm_build_718.dm_build_723(int(dm_build_721))
		dm_build_718.dm_build_396.paramCount = dm_build_721
	} else {
		dm_build_718.dm_build_396.params = make([]parameter, 0)
		dm_build_718.dm_build_396.paramCount = 0
	}

	if dm_build_722 > 0 {
		dm_build_718.dm_build_396.columns = dm_build_718.dm_build_534(int(dm_build_722), dm_build_719.rsBdta)
	} else {
		dm_build_718.dm_build_396.columns = make([]column, 0)
	}

	return dm_build_719, nil
}

func (dm_build_724 *dm_build_705) dm_build_723(dm_build_725 int) []parameter {

	var dm_build_726, dm_build_727, dm_build_728, dm_build_729 int16

	dm_build_730 := make([]parameter, dm_build_725)
	for i := 0; i < dm_build_725; i++ {

		dm_build_730[i].InitParameter()

		dm_build_730[i].colType = dm_build_724.dm_build_393.dm_build_5.Dm_build_1332()

		dm_build_730[i].prec = dm_build_724.dm_build_393.dm_build_5.Dm_build_1332()

		dm_build_730[i].scale = dm_build_724.dm_build_393.dm_build_5.Dm_build_1332()

		dm_build_730[i].nullable = dm_build_724.dm_build_393.dm_build_5.Dm_build_1332() != 0

		itemFlag := dm_build_724.dm_build_393.dm_build_5.Dm_build_1329()

		if int(itemFlag)&Dm_build_483 != 0 {
			dm_build_730[i].typeFlag = TYPE_FLAG_RECOMMEND
		} else {
			dm_build_730[i].typeFlag = TYPE_FLAG_EXACT
		}

		dm_build_730[i].lob = int(itemFlag)&Dm_build_481 != 0

		dm_build_724.dm_build_393.dm_build_5.Dm_build_1332()

		dm_build_730[i].ioType = byte(dm_build_724.dm_build_393.dm_build_5.Dm_build_1329())

		dm_build_726 = dm_build_724.dm_build_393.dm_build_5.Dm_build_1329()

		dm_build_727 = dm_build_724.dm_build_393.dm_build_5.Dm_build_1329()

		dm_build_728 = dm_build_724.dm_build_393.dm_build_5.Dm_build_1329()

		dm_build_729 = dm_build_724.dm_build_393.dm_build_5.Dm_build_1329()
		dm_build_730[i].name = dm_build_724.dm_build_393.dm_build_5.Dm_build_1369(int(dm_build_726), dm_build_724.dm_build_393.dm_build_6.getServerEncoding(), dm_build_724.dm_build_393.dm_build_6)
		dm_build_730[i].typeName = dm_build_724.dm_build_393.dm_build_5.Dm_build_1369(int(dm_build_727), dm_build_724.dm_build_393.dm_build_6.getServerEncoding(), dm_build_724.dm_build_393.dm_build_6)
		dm_build_730[i].tableName = dm_build_724.dm_build_393.dm_build_5.Dm_build_1369(int(dm_build_728), dm_build_724.dm_build_393.dm_build_6.getServerEncoding(), dm_build_724.dm_build_393.dm_build_6)
		dm_build_730[i].schemaName = dm_build_724.dm_build_393.dm_build_5.Dm_build_1369(int(dm_build_729), dm_build_724.dm_build_393.dm_build_6.getServerEncoding(), dm_build_724.dm_build_393.dm_build_6)

		if dm_build_730[i].lob {
			dm_build_730[i].lobTabId = dm_build_724.dm_build_393.dm_build_5.Dm_build_1332()
			dm_build_730[i].lobColId = dm_build_724.dm_build_393.dm_build_5.Dm_build_1329()
		}
	}

	for i := 0; i < dm_build_725; i++ {

		if isComplexType(int(dm_build_730[i].colType), int(dm_build_730[i].scale)) {

			strDesc := newTypeDescriptor(dm_build_724.dm_build_393.dm_build_6)
			strDesc.unpack(dm_build_724.dm_build_393.dm_build_5)
			dm_build_730[i].typeDescriptor = strDesc
		}
	}

	return dm_build_730
}

const (
	Dm_build_731 = Dm_build_289
)

type dm_build_732 struct {
	dm_build_392
	dm_build_733 int16
	dm_build_734 *Dm_build_1129
	dm_build_735 int32
}

func dm_build_736(dm_build_737 *dm_build_2, dm_build_738 *DmStatement, dm_build_739 int16, dm_build_740 *Dm_build_1129, dm_build_741 int32) *dm_build_732 {
	dm_build_742 := new(dm_build_732)
	dm_build_742.dm_build_401(dm_build_737, Dm_build_271, dm_build_738)
	dm_build_742.dm_build_733 = dm_build_739
	dm_build_742.dm_build_734 = dm_build_740
	dm_build_742.dm_build_735 = dm_build_741
	return dm_build_742
}

func (dm_build_744 *dm_build_732) dm_build_379() error {
	dm_build_744.dm_build_393.dm_build_5.Dm_build_1394(Dm_build_731, dm_build_744.dm_build_733)

	dm_build_744.dm_build_393.dm_build_5.Dm_build_1258(dm_build_744.dm_build_735)

	if dm_build_744.dm_build_393.dm_build_6.NewLobFlag {
		dm_build_744.dm_build_393.dm_build_5.Dm_build_1258(-1)
	}
	dm_build_744.dm_build_734.Dm_build_1136(dm_build_744.dm_build_393.dm_build_5, int(dm_build_744.dm_build_735))
	return nil
}

type dm_build_745 struct {
	dm_build_392
}

func dm_build_746(dm_build_747 *dm_build_2) *dm_build_745 {
	dm_build_748 := new(dm_build_745)
	dm_build_748.dm_build_397(dm_build_747, Dm_build_269)
	return dm_build_748
}

type dm_build_749 struct {
	dm_build_392
	dm_build_750 int32
}

func dm_build_751(dm_build_752 *dm_build_2, dm_build_753 int32) *dm_build_749 {
	dm_build_754 := new(dm_build_749)
	dm_build_754.dm_build_397(dm_build_752, Dm_build_282)
	dm_build_754.dm_build_750 = dm_build_753
	return dm_build_754
}

func (dm_build_756 *dm_build_749) dm_build_379() error {

	dm_build_757 := Dm_build_289
	dm_build_757 += dm_build_756.dm_build_393.dm_build_5.Dm_build_1398(dm_build_757, g2dbIsoLevel(dm_build_756.dm_build_750))
	return nil
}

type dm_build_758 struct {
	dm_build_392
	dm_build_759 *lob
	dm_build_760 byte
	dm_build_761 int
	dm_build_762 []byte
	dm_build_763 int
	dm_build_764 int
}

func dm_build_765(dm_build_766 *dm_build_2, dm_build_767 *lob, dm_build_768 byte, dm_build_769 int, dm_build_770 []byte,
	dm_build_771 int, dm_build_772 int) *dm_build_758 {
	dm_build_773 := new(dm_build_758)
	dm_build_773.dm_build_397(dm_build_766, Dm_build_278)
	dm_build_773.dm_build_759 = dm_build_767
	dm_build_773.dm_build_760 = dm_build_768
	dm_build_773.dm_build_761 = dm_build_769
	dm_build_773.dm_build_762 = dm_build_770
	dm_build_773.dm_build_763 = dm_build_771
	dm_build_773.dm_build_764 = dm_build_772
	return dm_build_773
}

func (dm_build_775 *dm_build_758) dm_build_379() error {

	dm_build_775.dm_build_393.dm_build_5.Dm_build_1250(byte(dm_build_775.dm_build_759.lobFlag))
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1250(dm_build_775.dm_build_760)
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1274(uint64(dm_build_775.dm_build_759.blobId))
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1254(dm_build_775.dm_build_759.groupId)
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1254(dm_build_775.dm_build_759.fileId)
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1258(dm_build_775.dm_build_759.pageNo)
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1254(dm_build_775.dm_build_759.curFileId)
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1258(dm_build_775.dm_build_759.curPageNo)
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1258(dm_build_775.dm_build_759.totalOffset)
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1258(dm_build_775.dm_build_759.tabId)
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1254(dm_build_775.dm_build_759.colId)
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1274(uint64(dm_build_775.dm_build_759.rowId))

	dm_build_775.dm_build_393.dm_build_5.Dm_build_1258(int32(dm_build_775.dm_build_761))
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1258(int32(dm_build_775.dm_build_764))
	dm_build_775.dm_build_393.dm_build_5.Dm_build_1286(dm_build_775.dm_build_762)

	if dm_build_775.dm_build_393.dm_build_6.NewLobFlag {
		dm_build_775.dm_build_393.dm_build_5.Dm_build_1254(dm_build_775.dm_build_759.exGroupId)
		dm_build_775.dm_build_393.dm_build_5.Dm_build_1254(dm_build_775.dm_build_759.exFileId)
		dm_build_775.dm_build_393.dm_build_5.Dm_build_1258(dm_build_775.dm_build_759.exPageNo)
	}
	return nil
}

func (dm_build_777 *dm_build_758) dm_build_383() (interface{}, error) {

	var dm_build_778 = dm_build_777.dm_build_393.dm_build_5.Dm_build_1332()
	dm_build_777.dm_build_759.blobId = dm_build_777.dm_build_393.dm_build_5.Dm_build_1335()
	dm_build_777.dm_build_759.fileId = dm_build_777.dm_build_393.dm_build_5.Dm_build_1329()
	dm_build_777.dm_build_759.pageNo = dm_build_777.dm_build_393.dm_build_5.Dm_build_1332()
	dm_build_777.dm_build_759.curFileId = dm_build_777.dm_build_393.dm_build_5.Dm_build_1329()
	dm_build_777.dm_build_759.curPageNo = dm_build_777.dm_build_393.dm_build_5.Dm_build_1332()
	dm_build_777.dm_build_759.totalOffset = dm_build_777.dm_build_393.dm_build_5.Dm_build_1332()
	return dm_build_778, nil
}

const (
	Dm_build_779 = Dm_build_289

	Dm_build_780 = Dm_build_779 + ULINT_SIZE

	Dm_build_781 = Dm_build_780 + ULINT_SIZE

	Dm_build_782 = Dm_build_781 + BYTE_SIZE

	Dm_build_783 = Dm_build_782 + BYTE_SIZE

	Dm_build_784 = Dm_build_783 + BYTE_SIZE

	Dm_build_785 = Dm_build_784 + BYTE_SIZE

	Dm_build_786 = Dm_build_289

	Dm_build_787 = Dm_build_786 + ULINT_SIZE

	Dm_build_788 = Dm_build_787 + ULINT_SIZE

	Dm_build_789 = Dm_build_788 + ULINT_SIZE

	Dm_build_790 = Dm_build_789 + ULINT_SIZE

	Dm_build_791 = Dm_build_790 + ULINT_SIZE

	Dm_build_792 = Dm_build_791 + BYTE_SIZE

	Dm_build_793 = Dm_build_792 + BYTE_SIZE

	Dm_build_794 = Dm_build_793 + BYTE_SIZE
)

type dm_build_795 struct {
	dm_build_392
	dm_build_796 *DmConnection
	dm_build_797 int
	Dm_build_798 int32
	Dm_build_799 []byte
	dm_build_800 byte
}

func dm_build_801(dm_build_802 *dm_build_2) *dm_build_795 {
	dm_build_803 := new(dm_build_795)
	dm_build_803.dm_build_397(dm_build_802, Dm_build_287)
	dm_build_803.dm_build_796 = dm_build_802.dm_build_6
	return dm_build_803
}

func dm_build_804(dm_build_805 string, dm_build_806 string) int {
	dm_build_807 := strings.Split(dm_build_805, ".")
	dm_build_808 := strings.Split(dm_build_806, ".")

	for i, serStr := range dm_build_807 {
		ser, _ := strconv.ParseInt(serStr, 10, 32)
		global, _ := strconv.ParseInt(dm_build_808[i], 10, 32)
		if ser < global {
			return -1
		} else if ser == global {
			continue
		} else {
			return 1
		}
	}

	return 0
}

func (dm_build_810 *dm_build_795) dm_build_379() error {

	dm_build_810.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_779, int32(0))
	dm_build_810.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_780, int32(dm_build_810.dm_build_796.dmConnector.compress))

	if dm_build_810.dm_build_796.dmConnector.loginEncrypt {
		dm_build_810.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_782, 2)
		dm_build_810.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_781, 1)
	} else {
		dm_build_810.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_782, 0)
		dm_build_810.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_781, 0)
	}

	if dm_build_810.dm_build_796.dmConnector.isBdtaRS {
		dm_build_810.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_783, Dm_build_367)
	} else {
		dm_build_810.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_783, Dm_build_366)
	}

	dm_build_810.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_784, byte(dm_build_810.dm_build_796.dmConnector.compressID))

	if dm_build_810.dm_build_796.dmConnector.loginCertificate != "" {
		dm_build_810.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_785, 1)
	} else {
		dm_build_810.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_785, 0)
	}

	dm_build_811 := dm_build_810.dm_build_796.getServerEncoding()
	dm_build_810.dm_build_393.dm_build_5.Dm_build_1302(Dm_build_254, dm_build_811, dm_build_810.dm_build_393.dm_build_6)

	var dm_build_812 byte
	if dm_build_810.dm_build_796.dmConnector.uKeyName != "" {
		dm_build_812 = 1
	} else {
		dm_build_812 = 0
	}

	dm_build_810.dm_build_393.dm_build_5.Dm_build_1250(0)

	if dm_build_812 == 1 {

	}

	if dm_build_810.dm_build_796.dmConnector.loginEncrypt {
		clientPubKey, err := dm_build_810.dm_build_393.dm_build_234()
		if err != nil {
			return err
		}
		dm_build_810.dm_build_393.dm_build_5.Dm_build_1290(clientPubKey)
	}

	return nil
}

func (dm_build_814 *dm_build_795) dm_build_382() error {
	dm_build_814.dm_build_393.dm_build_5.Dm_build_1224(0)
	dm_build_814.dm_build_393.dm_build_5.Dm_build_1232(Dm_build_288, false, true)
	return nil
}

func (dm_build_816 *dm_build_795) dm_build_383() (interface{}, error) {

	dm_build_816.dm_build_796.sslEncrypt = int(dm_build_816.dm_build_393.dm_build_5.Dm_build_1476(Dm_build_786))
	dm_build_816.dm_build_796.GlobalServerSeries = int(dm_build_816.dm_build_393.dm_build_5.Dm_build_1476(Dm_build_787))

	switch dm_build_816.dm_build_393.dm_build_5.Dm_build_1476(Dm_build_788) {
	case 1:
		dm_build_816.dm_build_796.serverEncoding = ENCODING_UTF8
	case 2:
		dm_build_816.dm_build_796.serverEncoding = ENCODING_EUCKR
	default:
		dm_build_816.dm_build_796.serverEncoding = ENCODING_GB18030
	}

	dm_build_816.dm_build_796.dmConnector.compress = int(dm_build_816.dm_build_393.dm_build_5.Dm_build_1476(Dm_build_789))
	dm_build_817 := dm_build_816.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_791)
	dm_build_818 := dm_build_816.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_792)
	dm_build_816.dm_build_796.dmConnector.isBdtaRS = dm_build_816.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_793) > 0
	dm_build_816.dm_build_796.dmConnector.compressID = int8(dm_build_816.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_794))

	dm_build_819 := dm_build_816.dm_build_425()
	if dm_build_819 != nil {
		return nil, dm_build_819
	}

	dm_build_820 := dm_build_816.dm_build_393.dm_build_5.Dm_build_1374(dm_build_816.dm_build_796.getServerEncoding(), dm_build_816.dm_build_393.dm_build_6)
	if dm_build_804(dm_build_820, Dm_build_255) < 0 {
		return nil, ECGO_ERROR_SERVER_VERSION.throw()
	}

	dm_build_816.dm_build_796.ServerVersion = dm_build_820
	dm_build_816.dm_build_796.Malini2 = dm_build_804(dm_build_820, Dm_build_256) > 0
	dm_build_816.dm_build_796.Execute2 = dm_build_804(dm_build_820, Dm_build_257) > 0
	dm_build_816.dm_build_796.LobEmptyCompOrcl = dm_build_804(dm_build_820, Dm_build_258) > 0

	if dm_build_816.dm_build_393.dm_build_6.dmConnector.uKeyName != "" {
		dm_build_816.dm_build_800 = 1
	} else {
		dm_build_816.dm_build_800 = 0
	}

	if dm_build_816.dm_build_800 == 1 {
		dm_build_816.dm_build_393.dm_build_11 = dm_build_816.dm_build_393.dm_build_5.Dm_build_1369(16, dm_build_816.dm_build_796.getServerEncoding(), dm_build_816.dm_build_393.dm_build_6)
	}

	dm_build_816.dm_build_797 = -1
	dm_build_821 := false
	dm_build_822 := false
	dm_build_816.Dm_build_798 = -1
	if dm_build_818 > 0 {
		dm_build_816.dm_build_797 = int(dm_build_816.dm_build_393.dm_build_5.Dm_build_1332())
	}

	if dm_build_817 > 0 {

		if dm_build_816.dm_build_797 == -1 {
			dm_build_821 = true
		} else {
			dm_build_822 = true
		}

		dm_build_816.Dm_build_799 = dm_build_816.dm_build_393.dm_build_5.Dm_build_1357()
	}

	if dm_build_818 == 2 {
		dm_build_816.Dm_build_798 = dm_build_816.dm_build_393.dm_build_5.Dm_build_1332()
	}
	dm_build_816.dm_build_393.dm_build_8 = dm_build_821
	dm_build_816.dm_build_393.dm_build_9 = dm_build_822

	return nil, nil
}

type dm_build_823 struct {
	dm_build_392
}

func dm_build_824(dm_build_825 *dm_build_2, dm_build_826 *DmStatement) *dm_build_823 {
	dm_build_827 := new(dm_build_823)
	dm_build_827.dm_build_401(dm_build_825, Dm_build_263, dm_build_826)
	return dm_build_827
}

func (dm_build_829 *dm_build_823) dm_build_379() error {

	dm_build_829.dm_build_393.dm_build_5.Dm_build_1390(Dm_build_289, 1)
	return nil
}

func (dm_build_831 *dm_build_823) dm_build_383() (interface{}, error) {

	dm_build_831.dm_build_396.id = dm_build_831.dm_build_393.dm_build_5.Dm_build_1476(Dm_build_290)

	dm_build_831.dm_build_396.readBaseColName = dm_build_831.dm_build_393.dm_build_5.Dm_build_1470(Dm_build_289) == 1
	return nil, nil
}

type dm_build_832 struct {
	dm_build_392
	dm_build_833 int32
}

func dm_build_834(dm_build_835 *dm_build_2, dm_build_836 int32) *dm_build_832 {
	dm_build_837 := new(dm_build_832)
	dm_build_837.dm_build_397(dm_build_835, Dm_build_264)
	dm_build_837.dm_build_833 = dm_build_836
	return dm_build_837
}

func (dm_build_839 *dm_build_832) dm_build_380() {
	dm_build_839.dm_build_392.dm_build_380()
	dm_build_839.dm_build_393.dm_build_5.Dm_build_1398(Dm_build_290, dm_build_839.dm_build_833)
}

type dm_build_840 struct {
	dm_build_392
	dm_build_841 []uint32
}

func dm_build_842(dm_build_843 *dm_build_2, dm_build_844 []uint32) *dm_build_840 {
	dm_build_845 := new(dm_build_840)
	dm_build_845.dm_build_397(dm_build_843, Dm_build_284)
	dm_build_845.dm_build_841 = dm_build_844
	return dm_build_845
}

func (dm_build_847 *dm_build_840) dm_build_379() error {

	dm_build_847.dm_build_393.dm_build_5.Dm_build_1418(Dm_build_289, uint16(len(dm_build_847.dm_build_841)))

	for _, tableID := range dm_build_847.dm_build_841 {
		dm_build_847.dm_build_393.dm_build_5.Dm_build_1270(uint32(tableID))
	}

	return nil
}

func (dm_build_849 *dm_build_840) dm_build_383() (interface{}, error) {
	dm_build_850 := dm_build_849.dm_build_393.dm_build_5.Dm_build_1491(Dm_build_289)
	if dm_build_850 <= 0 {
		return nil, nil
	}

	dm_build_851 := make([]int64, dm_build_850)
	for i := 0; i < int(dm_build_850); i++ {
		dm_build_851[i] = dm_build_849.dm_build_393.dm_build_5.Dm_build_1335()
	}

	return dm_build_851, nil
}
