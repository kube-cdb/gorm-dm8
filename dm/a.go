/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"crypto/tls"
	"dm/security"
	"net"
	"strconv"
	"time"
)

const (
	Dm_build_0 = 8192
	Dm_build_1 = 2 * time.Second
)

type dm_build_2 struct {
	dm_build_3  *net.TCPConn
	dm_build_4  *tls.Conn
	dm_build_5  *Dm_build_1207
	dm_build_6  *DmConnection
	dm_build_7  security.Cipher
	dm_build_8  bool
	dm_build_9  bool
	dm_build_10 *security.DhKey
	dm_build_11 string
	dm_build_12 bool
}

func dm_build_13(dm_build_14 *DmConnection) (*dm_build_2, error) {
	dm_build_15, dm_build_16 := dm_build_18(dm_build_14.dmConnector.host+":"+strconv.Itoa(dm_build_14.dmConnector.port), time.Duration(dm_build_14.dmConnector.socketTimeout)*time.Second)
	if dm_build_16 != nil {
		return nil, dm_build_16
	}

	dm_build_17 := dm_build_2{}
	dm_build_17.dm_build_3 = dm_build_15
	dm_build_17.dm_build_5 = Dm_build_1210(Dm_build_260)
	dm_build_17.dm_build_6 = dm_build_14
	dm_build_17.dm_build_8 = false
	dm_build_17.dm_build_9 = false
	dm_build_17.dm_build_11 = ""
	dm_build_17.dm_build_12 = false
	dm_build_14.Access = &dm_build_17

	return &dm_build_17, nil
}

func dm_build_18(dm_build_19 string, dm_build_20 time.Duration) (*net.TCPConn, error) {
	dm_build_21, dm_build_22 := net.DialTimeout("tcp", dm_build_19, dm_build_20)
	if dm_build_22 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_19).throw()
	}

	if tcpConn, ok := dm_build_21.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_1)
		tcpConn.SetNoDelay(true)

		tcpConn.SetReadBuffer(Dm_build_0)
		tcpConn.SetWriteBuffer(Dm_build_0)
		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_24 *dm_build_2) dm_build_23(dm_build_25 dm_build_377) bool {
	var dm_build_26 = dm_build_24.dm_build_6.dmConnector.compress
	if dm_build_25.dm_build_391() == Dm_build_287 || dm_build_26 == Dm_build_335 {
		return false
	}

	if dm_build_26 == Dm_build_333 {
		return true
	} else if dm_build_26 == Dm_build_334 {
		return !dm_build_24.dm_build_6.Local && dm_build_25.dm_build_389() > Dm_build_332
	}

	return false
}

func (dm_build_28 *dm_build_2) dm_build_27(dm_build_29 dm_build_377) bool {
	var dm_build_30 = dm_build_28.dm_build_6.dmConnector.compress
	if dm_build_29.dm_build_391() == Dm_build_287 || dm_build_30 == Dm_build_335 {
		return false
	}

	if dm_build_30 == Dm_build_333 {
		return true
	} else if dm_build_30 == Dm_build_334 {
		return dm_build_28.dm_build_5.Dm_build_1470(Dm_build_295) == 1
	}

	return false
}

func (dm_build_32 *dm_build_2) dm_build_31(dm_build_33 dm_build_377) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_35 := dm_build_33.dm_build_389()

	if dm_build_35 > 0 {

		if dm_build_32.dm_build_23(dm_build_33) {
			var retBytes, err = Compress(dm_build_32.dm_build_5, Dm_build_288, int(dm_build_35), int(dm_build_32.dm_build_6.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_32.dm_build_5.Dm_build_1221(Dm_build_288)

			dm_build_32.dm_build_5.Dm_build_1258(dm_build_35)

			dm_build_32.dm_build_5.Dm_build_1286(retBytes)

			dm_build_33.dm_build_390(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_32.dm_build_5.Dm_build_1390(Dm_build_295, 1)
		}

		if dm_build_32.dm_build_9 {
			dm_build_35 = dm_build_33.dm_build_389()
			var retBytes = dm_build_32.dm_build_7.Encrypt(dm_build_32.dm_build_5.Dm_build_1497(Dm_build_288, int(dm_build_35)), true)

			dm_build_32.dm_build_5.Dm_build_1221(Dm_build_288)

			dm_build_32.dm_build_5.Dm_build_1286(retBytes)

			dm_build_33.dm_build_390(int32(len(retBytes)))
		}
	}

	dm_build_33.dm_build_386()
	if dm_build_32.dm_build_251(dm_build_33) {
		if dm_build_32.dm_build_4 != nil {
			dm_build_32.dm_build_5.Dm_build_1224(0)
			dm_build_32.dm_build_5.Dm_build_1243(dm_build_32.dm_build_4)
		}
	} else {
		dm_build_32.dm_build_5.Dm_build_1224(0)
		dm_build_32.dm_build_5.Dm_build_1243(dm_build_32.dm_build_3)
	}
	return nil
}

func (dm_build_37 *dm_build_2) dm_build_36(dm_build_38 dm_build_377) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_40 := int32(0)
	if dm_build_37.dm_build_251(dm_build_38) {
		if dm_build_37.dm_build_4 != nil {
			dm_build_37.dm_build_5.Dm_build_1221(0)
			dm_build_37.dm_build_5.Dm_build_1237(dm_build_37.dm_build_4, Dm_build_288)
			dm_build_40 = dm_build_38.dm_build_389()
			if dm_build_40 > 0 {
				dm_build_37.dm_build_5.Dm_build_1237(dm_build_37.dm_build_4, int(dm_build_40))
			}
		}
	} else {

		dm_build_37.dm_build_5.Dm_build_1221(0)
		dm_build_37.dm_build_5.Dm_build_1237(dm_build_37.dm_build_3, Dm_build_288)
		dm_build_40 = dm_build_38.dm_build_389()

		if dm_build_40 > 0 {
			dm_build_37.dm_build_5.Dm_build_1237(dm_build_37.dm_build_3, int(dm_build_40))
		}
	}

	dm_build_38.dm_build_387()

	if dm_build_40 <= 0 {
		return nil
	}

	if dm_build_37.dm_build_9 {
		dm_build_40 = dm_build_38.dm_build_389()
		ebytes := dm_build_37.dm_build_5.Dm_build_1497(Dm_build_288, int(dm_build_40))
		bytes, err := dm_build_37.dm_build_7.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_37.dm_build_5.Dm_build_1221(Dm_build_288)
		dm_build_37.dm_build_5.Dm_build_1286(bytes)
		dm_build_38.dm_build_390(int32(len(bytes)))
	}

	if dm_build_37.dm_build_27(dm_build_38) {

		dm_build_40 = dm_build_38.dm_build_389()
		cbytes := dm_build_37.dm_build_5.Dm_build_1497(Dm_build_288+ULINT_SIZE, int(dm_build_40-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_37.dm_build_6.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_37.dm_build_5.Dm_build_1221(Dm_build_288)
		dm_build_37.dm_build_5.Dm_build_1286(bytes)
		dm_build_38.dm_build_390(int32(len(bytes)))
	}
	return nil
}

func (dm_build_42 *dm_build_2) dm_build_41(dm_build_43 dm_build_377) (dm_build_44 interface{}, dm_build_45 error) {
	dm_build_45 = dm_build_43.dm_build_381(dm_build_43)
	if dm_build_45 != nil {
		return nil, dm_build_45
	}

	dm_build_45 = dm_build_42.dm_build_31(dm_build_43)
	if dm_build_45 != nil {
		return nil, dm_build_45
	}

	dm_build_45 = dm_build_42.dm_build_36(dm_build_43)
	if dm_build_45 != nil {
		return nil, dm_build_45
	}

	return dm_build_43.dm_build_385(dm_build_43)
}

func (dm_build_47 *dm_build_2) dm_build_46() (*dm_build_795, error) {

	Dm_build_48 := dm_build_801(dm_build_47)
	_, dm_build_49 := dm_build_47.dm_build_41(Dm_build_48)
	if dm_build_49 != nil {
		return nil, dm_build_49
	}

	return Dm_build_48, nil
}

func (dm_build_51 *dm_build_2) dm_build_50() error {

	dm_build_52 := dm_build_672(dm_build_51)
	_, dm_build_53 := dm_build_51.dm_build_41(dm_build_52)
	if dm_build_53 != nil {
		return dm_build_53
	}

	return nil
}

func (dm_build_55 *dm_build_2) dm_build_54() error {

	var dm_build_56 *dm_build_795
	var err error
	if dm_build_56, err = dm_build_55.dm_build_46(); err != nil {
		return err
	}

	if dm_build_55.dm_build_6.sslEncrypt == 2 {
		if err = dm_build_55.dm_build_247(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_55.dm_build_6.sslEncrypt == 1 {
		if err = dm_build_55.dm_build_247(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_55.dm_build_9 || dm_build_55.dm_build_8 {
		k, err := dm_build_55.dm_build_237()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_56.Dm_build_799)
		encryptType := dm_build_56.dm_build_797
		hashType := int(dm_build_56.Dm_build_798)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_55.dm_build_240(encryptType, sessionKey, dm_build_55.dm_build_6.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_55.dm_build_50(); err != nil {
		return err
	}
	return nil
}

func (dm_build_59 *dm_build_2) Dm_build_58(dm_build_60 *DmStatement) error {
	dm_build_61 := dm_build_824(dm_build_59, dm_build_60)
	_, dm_build_62 := dm_build_59.dm_build_41(dm_build_61)
	if dm_build_62 != nil {
		return dm_build_62
	}

	return nil
}

func (dm_build_64 *dm_build_2) Dm_build_63(dm_build_65 int32) error {
	dm_build_66 := dm_build_834(dm_build_64, dm_build_65)
	_, dm_build_67 := dm_build_64.dm_build_41(dm_build_66)
	if dm_build_67 != nil {
		return dm_build_67
	}

	return nil
}

func (dm_build_69 *dm_build_2) Dm_build_68(dm_build_70 *DmStatement, dm_build_71 bool, dm_build_72 int16) (*execInfo, error) {
	dm_build_73 := dm_build_708(dm_build_69, dm_build_70, dm_build_71, dm_build_72)
	dm_build_74, dm_build_75 := dm_build_69.dm_build_41(dm_build_73)
	if dm_build_75 != nil {
		return nil, dm_build_75
	}
	return dm_build_74.(*execInfo), nil
}

func (dm_build_77 *dm_build_2) Dm_build_76(dm_build_78 *DmStatement, dm_build_79 int16) (*execInfo, error) {
	return dm_build_77.Dm_build_68(dm_build_78, false, Dm_build_339)
}

func (dm_build_81 *dm_build_2) Dm_build_80(dm_build_82 *DmStatement, dm_build_83 []OptParameter) (*execInfo, error) {
	dm_build_84, dm_build_85 := dm_build_81.dm_build_41(dm_build_470(dm_build_81, dm_build_82, dm_build_83))
	if dm_build_85 != nil {
		return nil, dm_build_85
	}

	return dm_build_84.(*execInfo), nil
}

func (dm_build_87 *dm_build_2) Dm_build_86(dm_build_88 *DmStatement, dm_build_89 int16) (*execInfo, error) {
	return dm_build_87.Dm_build_68(dm_build_88, true, dm_build_89)
}

func (dm_build_91 *dm_build_2) Dm_build_90(dm_build_92 *DmStatement, dm_build_93 [][]interface{}) (*execInfo, error) {
	dm_build_94 := dm_build_493(dm_build_91, dm_build_92, dm_build_93)
	dm_build_95, dm_build_96 := dm_build_91.dm_build_41(dm_build_94)
	if dm_build_96 != nil {
		return nil, dm_build_96
	}
	return dm_build_95.(*execInfo), nil
}

func (dm_build_98 *dm_build_2) Dm_build_97(dm_build_99 *DmStatement, dm_build_100 [][]interface{}) (*execInfo, error) {
	var dm_build_101, dm_build_102 = 0, 0
	var dm_build_103 = len(dm_build_100)
	var dm_build_104 [][]interface{}
	var dm_build_105 = NewExceInfo()
	dm_build_105.updateCounts = make([]int64, dm_build_103)
	var dm_build_106 = false
	var dm_build_107 = false
	for dm_build_101 < dm_build_103 {
		for dm_build_102 = dm_build_101; dm_build_102 < dm_build_103; dm_build_102++ {
			paramData := dm_build_100[dm_build_102]
			bindData := make([]interface{}, dm_build_99.paramCount)
			dm_build_106 = false
			for icol := 0; icol < int(dm_build_99.paramCount); icol++ {
				if dm_build_99.params[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_98.dm_build_220(bindData, paramData, icol) {
					dm_build_106 = true
					break
				}
			}

			if dm_build_106 {
				break
			}
			dm_build_104 = append(dm_build_104, bindData)
		}

		if dm_build_102 != dm_build_101 {
			tmpExecInfo, err := dm_build_98.Dm_build_90(dm_build_99, dm_build_104)
			if err != nil {
				return nil, err
			}
			dm_build_104 = dm_build_104[0:0]
			if dm_build_102-dm_build_101 == 1 {
				dm_build_105.updateCounts[dm_build_101] = tmpExecInfo.updateCount
			} else if tmpExecInfo.updateCounts != nil {
				copy(dm_build_105.updateCounts[dm_build_101:dm_build_102], tmpExecInfo.updateCounts[0:dm_build_102-dm_build_101])
			}
		}

		if dm_build_102 < dm_build_103 {
			tmpExecInfo, err := dm_build_98.Dm_build_108(dm_build_99, dm_build_100[dm_build_102], dm_build_107)
			if err != nil {
				return nil, err
			}

			dm_build_107 = true
			dm_build_105.updateCounts[dm_build_102] = tmpExecInfo.updateCount
		}

		dm_build_101 = dm_build_102 + 1
	}
	for _, i := range dm_build_105.updateCounts {
		dm_build_105.updateCount += i
	}
	return dm_build_105, nil
}

func (dm_build_109 *dm_build_2) Dm_build_108(dm_build_110 *DmStatement, dm_build_111 []interface{}, dm_build_112 bool) (*execInfo, error) {

	var dm_build_113 = make([]interface{}, dm_build_110.paramCount)
	for icol := 0; icol < int(dm_build_110.paramCount); icol++ {
		if dm_build_110.params[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_109.dm_build_220(dm_build_113, dm_build_111, icol) {

			if !dm_build_112 {
				preExecute := dm_build_698(dm_build_109, dm_build_110, dm_build_110.params)
				dm_build_109.dm_build_41(preExecute)
				dm_build_112 = true
			}

			dm_build_109.dm_build_226(dm_build_110, dm_build_110.params[icol], icol, dm_build_111[icol].(iOffRowBinder))
			dm_build_113[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_114 = make([][]interface{}, 1, 1)
	dm_build_114[0] = dm_build_113

	dm_build_115 := dm_build_493(dm_build_109, dm_build_110, dm_build_114)
	dm_build_116, dm_build_117 := dm_build_109.dm_build_41(dm_build_115)
	if dm_build_117 != nil {
		return nil, dm_build_117
	}
	return dm_build_116.(*execInfo), nil
}

func (dm_build_119 *dm_build_2) Dm_build_118(dm_build_120 *DmStatement, dm_build_121 int16) (*execInfo, error) {
	dm_build_122 := dm_build_686(dm_build_119, dm_build_120, dm_build_121)

	dm_build_123, dm_build_124 := dm_build_119.dm_build_41(dm_build_122)
	if dm_build_124 != nil {
		return nil, dm_build_124
	}
	return dm_build_123.(*execInfo), nil
}

func (dm_build_126 *dm_build_2) Dm_build_125(dm_build_127 *innerRows, dm_build_128 int64) (*execInfo, error) {
	dm_build_129 := dm_build_593(dm_build_126, dm_build_127, dm_build_128, INT64_MAX)
	dm_build_130, dm_build_131 := dm_build_126.dm_build_41(dm_build_129)
	if dm_build_131 != nil {
		return nil, dm_build_131
	}
	return dm_build_130.(*execInfo), nil
}

func (dm_build_133 *dm_build_2) Commit() error {
	dm_build_134 := dm_build_456(dm_build_133)
	_, dm_build_135 := dm_build_133.dm_build_41(dm_build_134)
	if dm_build_135 != nil {
		return dm_build_135
	}

	return nil
}

func (dm_build_137 *dm_build_2) Rollback() error {
	dm_build_138 := dm_build_746(dm_build_137)
	_, dm_build_139 := dm_build_137.dm_build_41(dm_build_138)
	if dm_build_139 != nil {
		return dm_build_139
	}

	return nil
}

func (dm_build_141 *dm_build_2) Dm_build_140(dm_build_142 *DmConnection) error {
	dm_build_143 := dm_build_751(dm_build_141, dm_build_142.IsoLevel)
	_, dm_build_144 := dm_build_141.dm_build_41(dm_build_143)
	if dm_build_144 != nil {
		return dm_build_144
	}

	return nil
}

func (dm_build_146 *dm_build_2) Dm_build_145(dm_build_147 *DmStatement, dm_build_148 string) error {
	dm_build_149 := dm_build_461(dm_build_146, dm_build_147, dm_build_148)
	_, dm_build_150 := dm_build_146.dm_build_41(dm_build_149)
	if dm_build_150 != nil {
		return dm_build_150
	}

	return nil
}

func (dm_build_152 *dm_build_2) Dm_build_151(dm_build_153 []uint32) ([]int64, error) {
	dm_build_154 := dm_build_842(dm_build_152, dm_build_153)
	dm_build_155, dm_build_156 := dm_build_152.dm_build_41(dm_build_154)
	if dm_build_156 != nil {
		return nil, dm_build_156
	}
	return dm_build_155.([]int64), nil
}

func (dm_build_158 *dm_build_2) Close() error {
	if dm_build_158.dm_build_12 {
		return nil
	}

	dm_build_159 := dm_build_158.dm_build_3.Close()
	if dm_build_159 != nil {
		return dm_build_159
	}

	dm_build_158.dm_build_6 = nil
	dm_build_158.dm_build_12 = true
	return nil
}

func (dm_build_161 *dm_build_2) dm_build_160(dm_build_162 *lob) (int64, error) {
	dm_build_163 := dm_build_624(dm_build_161, dm_build_162)
	dm_build_164, dm_build_165 := dm_build_161.dm_build_41(dm_build_163)
	if dm_build_165 != nil {
		return 0, dm_build_165
	}
	return dm_build_164.(int64), nil
}

func (dm_build_167 *dm_build_2) dm_build_166(dm_build_168 *DmBlob, dm_build_169 int32, dm_build_170 int32) ([]byte, error) {

	dm_build_171 := dm_build_611(dm_build_167, &dm_build_168.lob, int(dm_build_169), int(dm_build_170))
	dm_build_172, dm_build_173 := dm_build_167.dm_build_41(dm_build_171)
	if dm_build_173 != nil {
		return nil, dm_build_173
	}
	return dm_build_172.([]byte), nil
}

func (dm_build_175 *dm_build_2) dm_build_174(dm_build_176 *DmClob, dm_build_177 int32, dm_build_178 int32) (string, error) {

	dm_build_179 := dm_build_611(dm_build_175, &dm_build_176.lob, int(dm_build_177), int(dm_build_178))
	dm_build_180, dm_build_181 := dm_build_175.dm_build_41(dm_build_179)
	if dm_build_181 != nil {
		return "", dm_build_181
	}
	dm_build_182 := dm_build_180.([]byte)
	return Dm_build_853.Dm_build_1007(dm_build_182, 0, len(dm_build_182), dm_build_176.serverEncoding, dm_build_175.dm_build_6), nil
}

func (dm_build_184 *dm_build_2) dm_build_183(dm_build_185 *DmClob, dm_build_186 int, dm_build_187 string, dm_build_188 string) (int, error) {
	var dm_build_189 = Dm_build_853.Dm_build_1063(dm_build_187, dm_build_188, dm_build_184.dm_build_6)
	var dm_build_190 = 0
	var dm_build_191 = len(dm_build_189)
	var dm_build_192 = 0
	var dm_build_193 = 0
	var dm_build_194 = 0
	var dm_build_195 = dm_build_191/Dm_build_371 + 1
	var dm_build_196 byte = 0
	var dm_build_197 byte = 0x01
	var dm_build_198 byte = 0x02
	for i := 0; i < dm_build_195; i++ {
		dm_build_196 = 0
		if i == 0 {
			dm_build_196 |= dm_build_197
		}
		if i == dm_build_195-1 {
			dm_build_196 |= dm_build_198
		}
		dm_build_194 = dm_build_191 - dm_build_193
		if dm_build_194 > Dm_build_371 {
			dm_build_194 = Dm_build_371
		}

		setLobData := dm_build_765(dm_build_184, &dm_build_185.lob, dm_build_196, dm_build_186, dm_build_189, dm_build_190, dm_build_194)
		ret, err := dm_build_184.dm_build_41(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_192, nil
		} else {
			dm_build_186 += int(tmp)
			dm_build_192 += int(tmp)
			dm_build_193 += dm_build_194
			dm_build_190 += dm_build_194
		}
	}
	return dm_build_192, nil
}

func (dm_build_200 *dm_build_2) dm_build_199(dm_build_201 *DmBlob, dm_build_202 int, dm_build_203 []byte) (int, error) {
	var dm_build_204 = 0
	var dm_build_205 = len(dm_build_203)
	var dm_build_206 = 0
	var dm_build_207 = 0
	var dm_build_208 = 0
	var dm_build_209 = dm_build_205/Dm_build_371 + 1
	var dm_build_210 byte = 0
	var dm_build_211 byte = 0x01
	var dm_build_212 byte = 0x02
	for i := 0; i < dm_build_209; i++ {
		dm_build_210 = 0
		if i == 0 {
			dm_build_210 |= dm_build_211
		}
		if i == dm_build_209-1 {
			dm_build_210 |= dm_build_212
		}
		dm_build_208 = dm_build_205 - dm_build_207
		if dm_build_208 > Dm_build_371 {
			dm_build_208 = Dm_build_371
		}

		setLobData := dm_build_765(dm_build_200, &dm_build_201.lob, dm_build_210, dm_build_202, dm_build_203, dm_build_204, dm_build_208)
		ret, err := dm_build_200.dm_build_41(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_206, nil
		} else {
			dm_build_202 += int(tmp)
			dm_build_206 += int(tmp)
			dm_build_207 += dm_build_208
			dm_build_204 += dm_build_208
		}
	}
	return dm_build_206, nil
}

func (dm_build_214 *dm_build_2) dm_build_213(dm_build_215 *lob, dm_build_216 int) (int64, error) {
	dm_build_217 := dm_build_635(dm_build_214, dm_build_215, dm_build_216)
	dm_build_218, dm_build_219 := dm_build_214.dm_build_41(dm_build_217)
	if dm_build_219 != nil {
		return dm_build_215.length, dm_build_219
	}
	return dm_build_218.(int64), nil
}

func (dm_build_221 *dm_build_2) dm_build_220(dm_build_222 []interface{}, dm_build_223 []interface{}, dm_build_224 int) bool {
	var dm_build_225 = false
	if dm_build_224 >= len(dm_build_223) || dm_build_223[dm_build_224] == nil {
		dm_build_222[dm_build_224] = ParamDataEnum_Null
	} else if binder, ok := dm_build_223[dm_build_224].(iOffRowBinder); ok {
		dm_build_225 = true
		dm_build_222[dm_build_224] = ParamDataEnum_OFF_ROW
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_221.dm_build_6) {
			dm_build_222[dm_build_224] = &lobCtl{lob.buildCtlData()}
			dm_build_225 = false
		}
	} else {
		dm_build_222[dm_build_224] = dm_build_223[dm_build_224]
	}
	return dm_build_225
}

func (dm_build_227 *dm_build_2) dm_build_226(dm_build_228 *DmStatement, dm_build_229 parameter, dm_build_230 int, dm_build_231 iOffRowBinder) error {
	var dm_build_232 = Dm_build_1133()
	dm_build_231.read(dm_build_232)
	var dm_build_233 = 0
	for !dm_build_231.isReadOver() || dm_build_232.Dm_build_1134() > 0 {
		if !dm_build_231.isReadOver() && dm_build_232.Dm_build_1134() < Dm_build_371 {
			dm_build_231.read(dm_build_232)
		}
		if dm_build_232.Dm_build_1134() > Dm_build_371 {
			dm_build_233 = Dm_build_371
		} else {
			dm_build_233 = dm_build_232.Dm_build_1134()
		}

		putData := dm_build_736(dm_build_227, dm_build_228, int16(dm_build_230), dm_build_232, int32(dm_build_233))
		_, err := dm_build_227.dm_build_41(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_235 *dm_build_2) dm_build_234() ([]byte, error) {
	var dm_build_236 error
	if dm_build_235.dm_build_10 == nil {
		if dm_build_235.dm_build_10, dm_build_236 = security.NewClientKeyPair(); dm_build_236 != nil {
			return nil, dm_build_236
		}
	}
	return security.Bn2Bytes(dm_build_235.dm_build_10.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_238 *dm_build_2) dm_build_237() (*security.DhKey, error) {
	var dm_build_239 error
	if dm_build_238.dm_build_10 == nil {
		if dm_build_238.dm_build_10, dm_build_239 = security.NewClientKeyPair(); dm_build_239 != nil {
			return nil, dm_build_239
		}
	}
	return dm_build_238.dm_build_10, nil
}

func (dm_build_241 *dm_build_2) dm_build_240(dm_build_242 int, dm_build_243 []byte, dm_build_244 string, dm_build_245 int) (dm_build_246 error) {
	if dm_build_242 > 0 && dm_build_242 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_243 != nil {
		dm_build_241.dm_build_7, dm_build_246 = security.NewSymmCipher(dm_build_242, dm_build_243)
	} else if dm_build_242 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_241.dm_build_7, dm_build_246 = security.NewThirdPartCipher(dm_build_242, dm_build_243, dm_build_244, dm_build_245); dm_build_246 != nil {
			dm_build_246 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_246.Error()).throw()
		}
	}
	return
}

func (dm_build_248 *dm_build_2) dm_build_247(dm_build_249 bool) (dm_build_250 error) {
	if dm_build_248.dm_build_4, dm_build_250 = security.NewTLSFromTCP(dm_build_248.dm_build_3, dm_build_248.dm_build_6.dmConnector.sslCertPath, dm_build_248.dm_build_6.dmConnector.sslKeyPath, dm_build_248.dm_build_6.dmConnector.user); dm_build_250 != nil {
		return
	}
	if !dm_build_249 {
		dm_build_248.dm_build_4 = nil
	}
	return
}

func (dm_build_252 *dm_build_2) dm_build_251(dm_build_253 dm_build_377) bool {
	return dm_build_253.dm_build_391() != Dm_build_287 && dm_build_252.dm_build_6.sslEncrypt == 1
}
