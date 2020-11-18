/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_1129 struct {
	dm_build_1130 *list.List
	dm_build_1131 *dm_build_1183
	dm_build_1132 int
}

func Dm_build_1133() *Dm_build_1129 {
	return &Dm_build_1129{
		dm_build_1130: list.New(),
		dm_build_1132: 0,
	}
}

func (dm_build_1135 *Dm_build_1129) Dm_build_1134() int {
	return dm_build_1135.dm_build_1132
}

func (dm_build_1137 *Dm_build_1129) Dm_build_1136(dm_build_1138 *Dm_build_1207, dm_build_1139 int) int {
	var dm_build_1140 = 0
	var dm_build_1141 = 0
	for dm_build_1140 < dm_build_1139 && dm_build_1137.dm_build_1131 != nil {
		dm_build_1141 = dm_build_1137.dm_build_1131.dm_build_1191(dm_build_1138, dm_build_1139-dm_build_1140)
		if dm_build_1137.dm_build_1131.dm_build_1186 == 0 {
			dm_build_1137.dm_build_1173()
		}
		dm_build_1140 += dm_build_1141
		dm_build_1137.dm_build_1132 -= dm_build_1141
	}
	return dm_build_1140
}

func (dm_build_1143 *Dm_build_1129) Dm_build_1142(dm_build_1144 []byte, dm_build_1145 int, dm_build_1146 int) int {
	var dm_build_1147 = 0
	var dm_build_1148 = 0
	for dm_build_1147 < dm_build_1146 && dm_build_1143.dm_build_1131 != nil {
		dm_build_1148 = dm_build_1143.dm_build_1131.dm_build_1195(dm_build_1144, dm_build_1145, dm_build_1146-dm_build_1147)
		if dm_build_1143.dm_build_1131.dm_build_1186 == 0 {
			dm_build_1143.dm_build_1173()
		}
		dm_build_1147 += dm_build_1148
		dm_build_1143.dm_build_1132 -= dm_build_1148
		dm_build_1145 += dm_build_1148
	}
	return dm_build_1147
}

func (dm_build_1150 *Dm_build_1129) Dm_build_1149(dm_build_1151 io.Writer, dm_build_1152 int) int {
	var dm_build_1153 = 0
	var dm_build_1154 = 0
	for dm_build_1153 < dm_build_1152 && dm_build_1150.dm_build_1131 != nil {
		dm_build_1154 = dm_build_1150.dm_build_1131.dm_build_1200(dm_build_1151, dm_build_1152-dm_build_1153)
		if dm_build_1150.dm_build_1131.dm_build_1186 == 0 {
			dm_build_1150.dm_build_1173()
		}
		dm_build_1153 += dm_build_1154
		dm_build_1150.dm_build_1132 -= dm_build_1154
	}
	return dm_build_1153
}

func (dm_build_1156 *Dm_build_1129) Dm_build_1155(dm_build_1157 []byte, dm_build_1158 int, dm_build_1159 int) {
	if dm_build_1159 == 0 {
		return
	}
	var dm_build_1160 = dm_build_1187(dm_build_1157, dm_build_1158, dm_build_1159)
	if dm_build_1156.dm_build_1131 == nil {
		dm_build_1156.dm_build_1131 = dm_build_1160
	} else {
		dm_build_1156.dm_build_1130.PushBack(dm_build_1160)
	}
	dm_build_1156.dm_build_1132 += dm_build_1159
}

func (dm_build_1162 *Dm_build_1129) dm_build_1161(dm_build_1163 int) byte {
	var dm_build_1164 = dm_build_1163
	var dm_build_1165 = dm_build_1162.dm_build_1131
	for dm_build_1164 > 0 && dm_build_1165 != nil {
		if dm_build_1165.dm_build_1186 == 0 {
			continue
		}
		if dm_build_1164 > dm_build_1165.dm_build_1186-1 {
			dm_build_1164 -= dm_build_1165.dm_build_1186
			dm_build_1165 = dm_build_1162.dm_build_1130.Front().Value.(*dm_build_1183)
		} else {
			break
		}
	}
	return dm_build_1165.dm_build_1204(dm_build_1164)
}
func (dm_build_1167 *Dm_build_1129) Dm_build_1166(dm_build_1168 *Dm_build_1129) {
	if dm_build_1168.dm_build_1132 == 0 {
		return
	}
	var dm_build_1169 = dm_build_1168.dm_build_1131
	for dm_build_1169 != nil {
		dm_build_1167.dm_build_1170(dm_build_1169)
		dm_build_1168.dm_build_1173()
		dm_build_1169 = dm_build_1168.dm_build_1131
	}
	dm_build_1168.dm_build_1132 = 0
}
func (dm_build_1171 *Dm_build_1129) dm_build_1170(dm_build_1172 *dm_build_1183) {
	if dm_build_1172.dm_build_1186 == 0 {
		return
	}
	if dm_build_1171.dm_build_1131 == nil {
		dm_build_1171.dm_build_1131 = dm_build_1172
	} else {
		dm_build_1171.dm_build_1130.PushBack(dm_build_1172)
	}
	dm_build_1171.dm_build_1132 += dm_build_1172.dm_build_1186
}

func (dm_build_1174 *Dm_build_1129) dm_build_1173() {
	var dm_build_1175 = dm_build_1174.dm_build_1130.Front()
	if dm_build_1175 == nil {
		dm_build_1174.dm_build_1131 = nil
	} else {
		dm_build_1174.dm_build_1131 = dm_build_1175.Value.(*dm_build_1183)
		dm_build_1174.dm_build_1130.Remove(dm_build_1175)
	}
}

func (dm_build_1177 *Dm_build_1129) Dm_build_1176() []byte {
	var dm_build_1178 = make([]byte, dm_build_1177.dm_build_1132)
	var dm_build_1179 = dm_build_1177.dm_build_1131
	var dm_build_1180 = 0
	var dm_build_1181 = len(dm_build_1178)
	var dm_build_1182 = 0
	for dm_build_1179 != nil {
		if dm_build_1179.dm_build_1186 > 0 {
			if dm_build_1181 > dm_build_1179.dm_build_1186 {
				dm_build_1182 = dm_build_1179.dm_build_1186
			} else {
				dm_build_1182 = dm_build_1181
			}
			copy(dm_build_1178[dm_build_1180:dm_build_1180+dm_build_1182], dm_build_1179.dm_build_1184[dm_build_1179.dm_build_1185:dm_build_1179.dm_build_1185+dm_build_1182])
			dm_build_1180 += dm_build_1182
			dm_build_1181 -= dm_build_1182
		}
		dm_build_1179 = dm_build_1177.dm_build_1130.Front().Value.(*dm_build_1183)
	}
	return dm_build_1178
}

type dm_build_1183 struct {
	dm_build_1184 []byte
	dm_build_1185 int
	dm_build_1186 int
}

func dm_build_1187(dm_build_1188 []byte, dm_build_1189 int, dm_build_1190 int) *dm_build_1183 {
	return &dm_build_1183{
		dm_build_1188,
		dm_build_1189,
		dm_build_1190,
	}
}

func (dm_build_1192 *dm_build_1183) dm_build_1191(dm_build_1193 *Dm_build_1207, dm_build_1194 int) int {
	if dm_build_1192.dm_build_1186 <= dm_build_1194 {
		dm_build_1194 = dm_build_1192.dm_build_1186
	}
	dm_build_1193.Dm_build_1286(dm_build_1192.dm_build_1184[dm_build_1192.dm_build_1185 : dm_build_1192.dm_build_1185+dm_build_1194])
	dm_build_1192.dm_build_1185 += dm_build_1194
	dm_build_1192.dm_build_1186 -= dm_build_1194
	return dm_build_1194
}

func (dm_build_1196 *dm_build_1183) dm_build_1195(dm_build_1197 []byte, dm_build_1198 int, dm_build_1199 int) int {
	if dm_build_1196.dm_build_1186 <= dm_build_1199 {
		dm_build_1199 = dm_build_1196.dm_build_1186
	}
	copy(dm_build_1197[dm_build_1198:dm_build_1198+dm_build_1199], dm_build_1196.dm_build_1184[dm_build_1196.dm_build_1185:dm_build_1196.dm_build_1185+dm_build_1199])
	dm_build_1196.dm_build_1185 += dm_build_1199
	dm_build_1196.dm_build_1186 -= dm_build_1199
	return dm_build_1199
}

func (dm_build_1201 *dm_build_1183) dm_build_1200(dm_build_1202 io.Writer, dm_build_1203 int) int {
	if dm_build_1201.dm_build_1186 <= dm_build_1203 {
		dm_build_1203 = dm_build_1201.dm_build_1186
	}
	dm_build_1202.Write(dm_build_1201.dm_build_1184[dm_build_1201.dm_build_1185 : dm_build_1201.dm_build_1185+dm_build_1203])
	dm_build_1201.dm_build_1185 += dm_build_1203
	dm_build_1201.dm_build_1186 -= dm_build_1203
	return dm_build_1203
}
func (dm_build_1205 *dm_build_1183) dm_build_1204(dm_build_1206 int) byte {
	return dm_build_1205.dm_build_1184[dm_build_1205.dm_build_1185+dm_build_1206]
}
