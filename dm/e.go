/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_852 struct{}

var Dm_build_853 = &dm_build_852{}

func (Dm_build_855 *dm_build_852) Dm_build_854(dm_build_856 []byte, dm_build_857 int, dm_build_858 byte) int {
	dm_build_856[dm_build_857] = dm_build_858
	return 1
}

func (Dm_build_860 *dm_build_852) Dm_build_859(dm_build_861 []byte, dm_build_862 int, dm_build_863 int8) int {
	dm_build_861[dm_build_862] = byte(dm_build_863)
	return 1
}

func (Dm_build_865 *dm_build_852) Dm_build_864(dm_build_866 []byte, dm_build_867 int, dm_build_868 int16) int {
	dm_build_866[dm_build_867] = byte(dm_build_868)
	dm_build_867++
	dm_build_866[dm_build_867] = byte(dm_build_868 >> 8)
	return 2
}

func (Dm_build_870 *dm_build_852) Dm_build_869(dm_build_871 []byte, dm_build_872 int, dm_build_873 int32) int {
	dm_build_871[dm_build_872] = byte(dm_build_873)
	dm_build_872++
	dm_build_871[dm_build_872] = byte(dm_build_873 >> 8)
	dm_build_872++
	dm_build_871[dm_build_872] = byte(dm_build_873 >> 16)
	dm_build_872++
	dm_build_871[dm_build_872] = byte(dm_build_873 >> 24)
	dm_build_872++
	return 4
}

func (Dm_build_875 *dm_build_852) Dm_build_874(dm_build_876 []byte, dm_build_877 int, dm_build_878 int64) int {
	dm_build_876[dm_build_877] = byte(dm_build_878)
	dm_build_877++
	dm_build_876[dm_build_877] = byte(dm_build_878 >> 8)
	dm_build_877++
	dm_build_876[dm_build_877] = byte(dm_build_878 >> 16)
	dm_build_877++
	dm_build_876[dm_build_877] = byte(dm_build_878 >> 24)
	dm_build_877++
	dm_build_876[dm_build_877] = byte(dm_build_878 >> 32)
	dm_build_877++
	dm_build_876[dm_build_877] = byte(dm_build_878 >> 40)
	dm_build_877++
	dm_build_876[dm_build_877] = byte(dm_build_878 >> 48)
	dm_build_877++
	dm_build_876[dm_build_877] = byte(dm_build_878 >> 56)
	return 8
}

func (Dm_build_880 *dm_build_852) Dm_build_879(dm_build_881 []byte, dm_build_882 int, dm_build_883 float32) int {
	return Dm_build_880.Dm_build_899(dm_build_881, dm_build_882, math.Float32bits(dm_build_883))
}

func (Dm_build_885 *dm_build_852) Dm_build_884(dm_build_886 []byte, dm_build_887 int, dm_build_888 float64) int {
	return Dm_build_885.Dm_build_904(dm_build_886, dm_build_887, math.Float64bits(dm_build_888))
}

func (Dm_build_890 *dm_build_852) Dm_build_889(dm_build_891 []byte, dm_build_892 int, dm_build_893 uint8) int {
	dm_build_891[dm_build_892] = byte(dm_build_893)
	return 1
}

func (Dm_build_895 *dm_build_852) Dm_build_894(dm_build_896 []byte, dm_build_897 int, dm_build_898 uint16) int {
	dm_build_896[dm_build_897] = byte(dm_build_898)
	dm_build_897++
	dm_build_896[dm_build_897] = byte(dm_build_898 >> 8)
	return 2
}

func (Dm_build_900 *dm_build_852) Dm_build_899(dm_build_901 []byte, dm_build_902 int, dm_build_903 uint32) int {
	dm_build_901[dm_build_902] = byte(dm_build_903)
	dm_build_902++
	dm_build_901[dm_build_902] = byte(dm_build_903 >> 8)
	dm_build_902++
	dm_build_901[dm_build_902] = byte(dm_build_903 >> 16)
	dm_build_902++
	dm_build_901[dm_build_902] = byte(dm_build_903 >> 24)
	return 3
}

func (Dm_build_905 *dm_build_852) Dm_build_904(dm_build_906 []byte, dm_build_907 int, dm_build_908 uint64) int {
	dm_build_906[dm_build_907] = byte(dm_build_908)
	dm_build_907++
	dm_build_906[dm_build_907] = byte(dm_build_908 >> 8)
	dm_build_907++
	dm_build_906[dm_build_907] = byte(dm_build_908 >> 16)
	dm_build_907++
	dm_build_906[dm_build_907] = byte(dm_build_908 >> 24)
	dm_build_907++
	dm_build_906[dm_build_907] = byte(dm_build_908 >> 32)
	dm_build_907++
	dm_build_906[dm_build_907] = byte(dm_build_908 >> 40)
	dm_build_907++
	dm_build_906[dm_build_907] = byte(dm_build_908 >> 48)
	dm_build_907++
	dm_build_906[dm_build_907] = byte(dm_build_908 >> 56)
	return 3
}

func (Dm_build_910 *dm_build_852) Dm_build_909(dm_build_911 []byte, dm_build_912 int, dm_build_913 []byte, dm_build_914 int, dm_build_915 int) int {
	copy(dm_build_911[dm_build_912:dm_build_912+dm_build_915], dm_build_913[dm_build_914:dm_build_914+dm_build_915])
	return dm_build_915
}

func (Dm_build_917 *dm_build_852) Dm_build_916(dm_build_918 []byte, dm_build_919 int, dm_build_920 []byte, dm_build_921 int, dm_build_922 int) int {
	dm_build_919 += Dm_build_917.Dm_build_899(dm_build_918, dm_build_919, uint32(dm_build_922))
	return 4 + Dm_build_917.Dm_build_909(dm_build_918, dm_build_919, dm_build_920, dm_build_921, dm_build_922)
}

func (Dm_build_924 *dm_build_852) Dm_build_923(dm_build_925 []byte, dm_build_926 int, dm_build_927 []byte, dm_build_928 int, dm_build_929 int) int {
	dm_build_926 += Dm_build_924.Dm_build_894(dm_build_925, dm_build_926, uint16(dm_build_929))
	return 2 + Dm_build_924.Dm_build_909(dm_build_925, dm_build_926, dm_build_927, dm_build_928, dm_build_929)
}

func (Dm_build_931 *dm_build_852) Dm_build_930(dm_build_932 []byte, dm_build_933 int, dm_build_934 string, dm_build_935 string, dm_build_936 *DmConnection) int {
	dm_build_937 := Dm_build_931.Dm_build_1063(dm_build_934, dm_build_935, dm_build_936)
	dm_build_933 += Dm_build_931.Dm_build_899(dm_build_932, dm_build_933, uint32(len(dm_build_937)))
	return 4 + Dm_build_931.Dm_build_909(dm_build_932, dm_build_933, dm_build_937, 0, len(dm_build_937))
}

func (Dm_build_939 *dm_build_852) Dm_build_938(dm_build_940 []byte, dm_build_941 int, dm_build_942 string, dm_build_943 string, dm_build_944 *DmConnection) int {
	dm_build_945 := Dm_build_939.Dm_build_1063(dm_build_942, dm_build_943, dm_build_944)

	dm_build_941 += Dm_build_939.Dm_build_894(dm_build_940, dm_build_941, uint16(len(dm_build_945)))
	return 2 + Dm_build_939.Dm_build_909(dm_build_940, dm_build_941, dm_build_945, 0, len(dm_build_945))
}

func (Dm_build_947 *dm_build_852) Dm_build_946(dm_build_948 []byte, dm_build_949 int) byte {
	return dm_build_948[dm_build_949]
}

func (Dm_build_951 *dm_build_852) Dm_build_950(dm_build_952 []byte, dm_build_953 int) int16 {
	var dm_build_954 int16
	dm_build_954 = int16(dm_build_952[dm_build_953] & 0xff)
	dm_build_953++
	dm_build_954 |= int16(dm_build_952[dm_build_953]&0xff) << 8
	return dm_build_954
}

func (Dm_build_956 *dm_build_852) Dm_build_955(dm_build_957 []byte, dm_build_958 int) int32 {
	var dm_build_959 int32
	dm_build_959 = int32(dm_build_957[dm_build_958] & 0xff)
	dm_build_958++
	dm_build_959 |= int32(dm_build_957[dm_build_958]&0xff) << 8
	dm_build_958++
	dm_build_959 |= int32(dm_build_957[dm_build_958]&0xff) << 16
	dm_build_958++
	dm_build_959 |= int32(dm_build_957[dm_build_958]&0xff) << 24
	return dm_build_959
}

func (Dm_build_961 *dm_build_852) Dm_build_960(dm_build_962 []byte, dm_build_963 int) int64 {
	var dm_build_964 int64
	dm_build_964 = int64(dm_build_962[dm_build_963] & 0xff)
	dm_build_963++
	dm_build_964 |= int64(dm_build_962[dm_build_963]&0xff) << 8
	dm_build_963++
	dm_build_964 |= int64(dm_build_962[dm_build_963]&0xff) << 16
	dm_build_963++
	dm_build_964 |= int64(dm_build_962[dm_build_963]&0xff) << 24
	dm_build_963++
	dm_build_964 |= int64(dm_build_962[dm_build_963]&0xff) << 32
	dm_build_963++
	dm_build_964 |= int64(dm_build_962[dm_build_963]&0xff) << 40
	dm_build_963++
	dm_build_964 |= int64(dm_build_962[dm_build_963]&0xff) << 48
	dm_build_963++
	dm_build_964 |= int64(dm_build_962[dm_build_963]&0xff) << 56
	return dm_build_964
}

func (Dm_build_966 *dm_build_852) Dm_build_965(dm_build_967 []byte, dm_build_968 int) float32 {
	return math.Float32frombits(Dm_build_966.Dm_build_982(dm_build_967, dm_build_968))
}

func (Dm_build_970 *dm_build_852) Dm_build_969(dm_build_971 []byte, dm_build_972 int) float64 {
	return math.Float64frombits(Dm_build_970.Dm_build_987(dm_build_971, dm_build_972))
}

func (Dm_build_974 *dm_build_852) Dm_build_973(dm_build_975 []byte, dm_build_976 int) uint8 {
	return uint8(dm_build_975[dm_build_976] & 0xff)
}

func (Dm_build_978 *dm_build_852) Dm_build_977(dm_build_979 []byte, dm_build_980 int) uint16 {
	var dm_build_981 uint16
	dm_build_981 = uint16(dm_build_979[dm_build_980] & 0xff)
	dm_build_980++
	dm_build_981 |= uint16(dm_build_979[dm_build_980]&0xff) << 8
	return dm_build_981
}

func (Dm_build_983 *dm_build_852) Dm_build_982(dm_build_984 []byte, dm_build_985 int) uint32 {
	var dm_build_986 uint32
	dm_build_986 = uint32(dm_build_984[dm_build_985] & 0xff)
	dm_build_985++
	dm_build_986 |= uint32(dm_build_984[dm_build_985]&0xff) << 8
	dm_build_985++
	dm_build_986 |= uint32(dm_build_984[dm_build_985]&0xff) << 16
	dm_build_985++
	dm_build_986 |= uint32(dm_build_984[dm_build_985]&0xff) << 24
	return dm_build_986
}

func (Dm_build_988 *dm_build_852) Dm_build_987(dm_build_989 []byte, dm_build_990 int) uint64 {
	var dm_build_991 uint64
	dm_build_991 = uint64(dm_build_989[dm_build_990] & 0xff)
	dm_build_990++
	dm_build_991 |= uint64(dm_build_989[dm_build_990]&0xff) << 8
	dm_build_990++
	dm_build_991 |= uint64(dm_build_989[dm_build_990]&0xff) << 16
	dm_build_990++
	dm_build_991 |= uint64(dm_build_989[dm_build_990]&0xff) << 24
	dm_build_990++
	dm_build_991 |= uint64(dm_build_989[dm_build_990]&0xff) << 32
	dm_build_990++
	dm_build_991 |= uint64(dm_build_989[dm_build_990]&0xff) << 40
	dm_build_990++
	dm_build_991 |= uint64(dm_build_989[dm_build_990]&0xff) << 48
	dm_build_990++
	dm_build_991 |= uint64(dm_build_989[dm_build_990]&0xff) << 56
	return dm_build_991
}

func (Dm_build_993 *dm_build_852) Dm_build_992(dm_build_994 []byte, dm_build_995 int) []byte {
	dm_build_996 := Dm_build_993.Dm_build_982(dm_build_994, dm_build_995)
	return dm_build_994[dm_build_995+4 : dm_build_995+4+int(dm_build_996)]

}

func (Dm_build_998 *dm_build_852) Dm_build_997(dm_build_999 []byte, dm_build_1000 int) []byte {
	dm_build_1001 := Dm_build_998.Dm_build_977(dm_build_999, dm_build_1000)
	return dm_build_999[dm_build_1000+2 : dm_build_1000+2+int(dm_build_1001)]

}

func (Dm_build_1003 *dm_build_852) Dm_build_1002(dm_build_1004 []byte, dm_build_1005 int, dm_build_1006 int) []byte {
	return dm_build_1004[dm_build_1005 : dm_build_1005+dm_build_1006]

}

func (Dm_build_1008 *dm_build_852) Dm_build_1007(dm_build_1009 []byte, dm_build_1010 int, dm_build_1011 int, dm_build_1012 string, dm_build_1013 *DmConnection) string {
	return Dm_build_1008.Dm_build_1100(dm_build_1009[dm_build_1010:dm_build_1010+dm_build_1011], dm_build_1012, dm_build_1013)
}

func (Dm_build_1015 *dm_build_852) Dm_build_1014(dm_build_1016 []byte, dm_build_1017 int, dm_build_1018 string, dm_build_1019 *DmConnection) string {
	dm_build_1020 := Dm_build_1015.Dm_build_982(dm_build_1016, dm_build_1017)
	dm_build_1017 += 4
	return Dm_build_1015.Dm_build_1007(dm_build_1016, dm_build_1017, int(dm_build_1020), dm_build_1018, dm_build_1019)
}

func (Dm_build_1022 *dm_build_852) Dm_build_1021(dm_build_1023 []byte, dm_build_1024 int, dm_build_1025 string, dm_build_1026 *DmConnection) string {
	dm_build_1027 := Dm_build_1022.Dm_build_977(dm_build_1023, dm_build_1024)
	dm_build_1024 += 2
	return Dm_build_1022.Dm_build_1007(dm_build_1023, dm_build_1024, int(dm_build_1027), dm_build_1025, dm_build_1026)
}

func (Dm_build_1029 *dm_build_852) Dm_build_1028(dm_build_1030 byte) []byte {
	return []byte{dm_build_1030}
}

func (Dm_build_1032 *dm_build_852) Dm_build_1031(dm_build_1033 int16) []byte {
	return []byte{byte(dm_build_1033), byte(dm_build_1033 >> 8)}
}

func (Dm_build_1035 *dm_build_852) Dm_build_1034(dm_build_1036 int32) []byte {
	return []byte{byte(dm_build_1036), byte(dm_build_1036 >> 8), byte(dm_build_1036 >> 16), byte(dm_build_1036 >> 24)}
}

func (Dm_build_1038 *dm_build_852) Dm_build_1037(dm_build_1039 int64) []byte {
	return []byte{byte(dm_build_1039), byte(dm_build_1039 >> 8), byte(dm_build_1039 >> 16), byte(dm_build_1039 >> 24), byte(dm_build_1039 >> 32),
		byte(dm_build_1039 >> 40), byte(dm_build_1039 >> 48), byte(dm_build_1039 >> 56)}
}

func (Dm_build_1041 *dm_build_852) Dm_build_1040(dm_build_1042 float32) []byte {
	return Dm_build_1041.Dm_build_1052(math.Float32bits(dm_build_1042))
}

func (Dm_build_1044 *dm_build_852) Dm_build_1043(dm_build_1045 float64) []byte {
	return Dm_build_1044.Dm_build_1055(math.Float64bits(dm_build_1045))
}

func (Dm_build_1047 *dm_build_852) Dm_build_1046(dm_build_1048 uint8) []byte {
	return []byte{byte(dm_build_1048)}
}

func (Dm_build_1050 *dm_build_852) Dm_build_1049(dm_build_1051 uint16) []byte {
	return []byte{byte(dm_build_1051), byte(dm_build_1051 >> 8)}
}

func (Dm_build_1053 *dm_build_852) Dm_build_1052(dm_build_1054 uint32) []byte {
	return []byte{byte(dm_build_1054), byte(dm_build_1054 >> 8), byte(dm_build_1054 >> 16), byte(dm_build_1054 >> 24)}
}

func (Dm_build_1056 *dm_build_852) Dm_build_1055(dm_build_1057 uint64) []byte {
	return []byte{byte(dm_build_1057), byte(dm_build_1057 >> 8), byte(dm_build_1057 >> 16), byte(dm_build_1057 >> 24), byte(dm_build_1057 >> 32), byte(dm_build_1057 >> 40), byte(dm_build_1057 >> 48), byte(dm_build_1057 >> 56)}
}

func (Dm_build_1059 *dm_build_852) Dm_build_1058(dm_build_1060 []byte, dm_build_1061 string, dm_build_1062 *DmConnection) []byte {
	if dm_build_1061 == "UTF-8" {
		return dm_build_1060
	}

	if dm_build_1062 == nil {
		if e := dm_build_1105(dm_build_1061); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1060), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1062.encodeBuffer == nil {
		dm_build_1062.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1062.encode = dm_build_1105(dm_build_1062.getServerEncoding())
		dm_build_1062.transformReaderDst = make([]byte, 4096)
		dm_build_1062.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1062.encode; e != nil {

		dm_build_1062.encodeBuffer.Reset()

		n, err := dm_build_1062.encodeBuffer.ReadFrom(
			Dm_build_1119(bytes.NewReader(dm_build_1060), e.NewEncoder(), dm_build_1062.transformReaderDst, dm_build_1062.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1062.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1064 *dm_build_852) Dm_build_1063(dm_build_1065 string, dm_build_1066 string, dm_build_1067 *DmConnection) []byte {
	return Dm_build_1064.Dm_build_1058([]byte(dm_build_1065), dm_build_1066, dm_build_1067)
}

func (Dm_build_1069 *dm_build_852) Dm_build_1068(dm_build_1070 []byte) byte {
	return Dm_build_1069.Dm_build_946(dm_build_1070, 0)
}

func (Dm_build_1072 *dm_build_852) Dm_build_1071(dm_build_1073 []byte) int16 {
	return Dm_build_1072.Dm_build_950(dm_build_1073, 0)
}

func (Dm_build_1075 *dm_build_852) Dm_build_1074(dm_build_1076 []byte) int32 {
	return Dm_build_1075.Dm_build_955(dm_build_1076, 0)
}

func (Dm_build_1078 *dm_build_852) Dm_build_1077(dm_build_1079 []byte) int64 {
	return Dm_build_1078.Dm_build_960(dm_build_1079, 0)
}

func (Dm_build_1081 *dm_build_852) Dm_build_1080(dm_build_1082 []byte) float32 {
	return Dm_build_1081.Dm_build_965(dm_build_1082, 0)
}

func (Dm_build_1084 *dm_build_852) Dm_build_1083(dm_build_1085 []byte) float64 {
	return Dm_build_1084.Dm_build_969(dm_build_1085, 0)
}

func (Dm_build_1087 *dm_build_852) Dm_build_1086(dm_build_1088 []byte) uint8 {
	return Dm_build_1087.Dm_build_973(dm_build_1088, 0)
}

func (Dm_build_1090 *dm_build_852) Dm_build_1089(dm_build_1091 []byte) uint16 {
	return Dm_build_1090.Dm_build_977(dm_build_1091, 0)
}

func (Dm_build_1093 *dm_build_852) Dm_build_1092(dm_build_1094 []byte) uint32 {
	return Dm_build_1093.Dm_build_982(dm_build_1094, 0)
}

func (Dm_build_1096 *dm_build_852) Dm_build_1095(dm_build_1097 []byte, dm_build_1098 string, dm_build_1099 *DmConnection) []byte {
	if dm_build_1098 == "UTF-8" {
		return dm_build_1097
	}

	if dm_build_1099 == nil {
		if e := dm_build_1105(dm_build_1098); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1097), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1099.encodeBuffer == nil {
		dm_build_1099.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1099.encode = dm_build_1105(dm_build_1099.getServerEncoding())
		dm_build_1099.transformReaderDst = make([]byte, 4096)
		dm_build_1099.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1099.encode; e != nil {

		dm_build_1099.encodeBuffer.Reset()

		n, err := dm_build_1099.encodeBuffer.ReadFrom(
			Dm_build_1119(bytes.NewReader(dm_build_1097), e.NewDecoder(), dm_build_1099.transformReaderDst, dm_build_1099.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1099.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1101 *dm_build_852) Dm_build_1100(dm_build_1102 []byte, dm_build_1103 string, dm_build_1104 *DmConnection) string {
	return string(Dm_build_1101.Dm_build_1095(dm_build_1102, dm_build_1103, dm_build_1104))
}

func dm_build_1105(dm_build_1106 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1106); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1107 struct {
	dm_build_1108 io.Reader
	dm_build_1109 transform.Transformer
	dm_build_1110 error

	dm_build_1111                []byte
	dm_build_1112, dm_build_1113 int

	dm_build_1114                []byte
	dm_build_1115, dm_build_1116 int

	dm_build_1117 bool
}

const dm_build_1118 = 4096

func Dm_build_1119(dm_build_1120 io.Reader, dm_build_1121 transform.Transformer, dm_build_1122 []byte, dm_build_1123 []byte) *Dm_build_1107 {
	dm_build_1121.Reset()
	return &Dm_build_1107{
		dm_build_1108: dm_build_1120,
		dm_build_1109: dm_build_1121,
		dm_build_1111: dm_build_1122,
		dm_build_1114: dm_build_1123,
	}
}

func (dm_build_1125 *Dm_build_1107) Read(dm_build_1126 []byte) (int, error) {
	dm_build_1127, dm_build_1128 := 0, error(nil)
	for {

		if dm_build_1125.dm_build_1112 != dm_build_1125.dm_build_1113 {
			dm_build_1127 = copy(dm_build_1126, dm_build_1125.dm_build_1111[dm_build_1125.dm_build_1112:dm_build_1125.dm_build_1113])
			dm_build_1125.dm_build_1112 += dm_build_1127
			if dm_build_1125.dm_build_1112 == dm_build_1125.dm_build_1113 && dm_build_1125.dm_build_1117 {
				return dm_build_1127, dm_build_1125.dm_build_1110
			}
			return dm_build_1127, nil
		} else if dm_build_1125.dm_build_1117 {
			return 0, dm_build_1125.dm_build_1110
		}

		if dm_build_1125.dm_build_1115 != dm_build_1125.dm_build_1116 || dm_build_1125.dm_build_1110 != nil {
			dm_build_1125.dm_build_1112 = 0
			dm_build_1125.dm_build_1113, dm_build_1127, dm_build_1128 = dm_build_1125.dm_build_1109.Transform(dm_build_1125.dm_build_1111, dm_build_1125.dm_build_1114[dm_build_1125.dm_build_1115:dm_build_1125.dm_build_1116], dm_build_1125.dm_build_1110 == io.EOF)
			dm_build_1125.dm_build_1115 += dm_build_1127

			switch {
			case dm_build_1128 == nil:
				if dm_build_1125.dm_build_1115 != dm_build_1125.dm_build_1116 {
					dm_build_1125.dm_build_1110 = nil
				}

				dm_build_1125.dm_build_1117 = dm_build_1125.dm_build_1110 != nil
				continue
			case dm_build_1128 == transform.ErrShortDst && (dm_build_1125.dm_build_1113 != 0 || dm_build_1127 != 0):

				continue
			case dm_build_1128 == transform.ErrShortSrc && dm_build_1125.dm_build_1116-dm_build_1125.dm_build_1115 != len(dm_build_1125.dm_build_1114) && dm_build_1125.dm_build_1110 == nil:

			default:
				dm_build_1125.dm_build_1117 = true

				if dm_build_1125.dm_build_1110 == nil || dm_build_1125.dm_build_1110 == io.EOF {
					dm_build_1125.dm_build_1110 = dm_build_1128
				}
				continue
			}
		}

		if dm_build_1125.dm_build_1115 != 0 {
			dm_build_1125.dm_build_1115, dm_build_1125.dm_build_1116 = 0, copy(dm_build_1125.dm_build_1114, dm_build_1125.dm_build_1114[dm_build_1125.dm_build_1115:dm_build_1125.dm_build_1116])
		}
		dm_build_1127, dm_build_1125.dm_build_1110 = dm_build_1125.dm_build_1108.Read(dm_build_1125.dm_build_1114[dm_build_1125.dm_build_1116:])
		dm_build_1125.dm_build_1116 += dm_build_1127
	}
}
