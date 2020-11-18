/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_1207 struct {
	dm_build_1208 []byte
	dm_build_1209 int
}

func Dm_build_1210(dm_build_1211 int) *Dm_build_1207 {
	return &Dm_build_1207{make([]byte, 0, dm_build_1211), 0}
}

func Dm_build_1212(dm_build_1213 []byte) *Dm_build_1207 {
	return &Dm_build_1207{dm_build_1213, 0}
}

func (dm_build_1215 *Dm_build_1207) dm_build_1214(dm_build_1216 int) *Dm_build_1207 {

	dm_build_1217 := len(dm_build_1215.dm_build_1208)
	dm_build_1218 := cap(dm_build_1215.dm_build_1208)

	if dm_build_1217+dm_build_1216 <= dm_build_1218 {
		dm_build_1215.dm_build_1208 = dm_build_1215.dm_build_1208[:dm_build_1217+dm_build_1216]
	} else {
		remain := dm_build_1216 + dm_build_1217 - dm_build_1218
		nbuf := make([]byte, dm_build_1216+dm_build_1217, 2*dm_build_1218+remain)
		copy(nbuf, dm_build_1215.dm_build_1208)
		dm_build_1215.dm_build_1208 = nbuf
	}

	return dm_build_1215
}

func (dm_build_1220 *Dm_build_1207) Dm_build_1219() int {
	return len(dm_build_1220.dm_build_1208)
}

func (dm_build_1222 *Dm_build_1207) Dm_build_1221(dm_build_1223 int) *Dm_build_1207 {
	dm_build_1222.dm_build_1208 = dm_build_1222.dm_build_1208[:dm_build_1223]
	return dm_build_1222
}

func (dm_build_1225 *Dm_build_1207) Dm_build_1224(dm_build_1226 int) *Dm_build_1207 {
	dm_build_1225.dm_build_1209 = dm_build_1226
	return dm_build_1225
}

func (dm_build_1228 *Dm_build_1207) Dm_build_1227() int {
	return dm_build_1228.dm_build_1209
}

func (dm_build_1230 *Dm_build_1207) Dm_build_1229(dm_build_1231 bool) int {
	return len(dm_build_1230.dm_build_1208) - dm_build_1230.dm_build_1209
}

func (dm_build_1233 *Dm_build_1207) Dm_build_1232(dm_build_1234 int, dm_build_1235 bool, dm_build_1236 bool) *Dm_build_1207 {

	if dm_build_1235 {
		if dm_build_1236 {
			dm_build_1233.dm_build_1214(dm_build_1234)
		} else {
			dm_build_1233.dm_build_1208 = dm_build_1233.dm_build_1208[:len(dm_build_1233.dm_build_1208)-dm_build_1234]
		}
	} else {
		if dm_build_1236 {
			dm_build_1233.dm_build_1209 += dm_build_1234
		} else {
			dm_build_1233.dm_build_1209 -= dm_build_1234
		}
	}

	return dm_build_1233
}

func (dm_build_1238 *Dm_build_1207) Dm_build_1237(dm_build_1239 io.Reader, dm_build_1240 int) int {
	dm_build_1241 := len(dm_build_1238.dm_build_1208)
	dm_build_1238.dm_build_1214(dm_build_1240)
	dm_build_1242 := 0
	for dm_build_1240 > 0 {
		n, err := dm_build_1239.Read(dm_build_1238.dm_build_1208[dm_build_1241+dm_build_1242:])
		if n > 0 && err == io.EOF {
			dm_build_1242 += n
			dm_build_1238.dm_build_1208 = dm_build_1238.dm_build_1208[:dm_build_1241+dm_build_1242]
			return dm_build_1242
		} else if n > 0 && err == nil {
			dm_build_1240 -= n
			dm_build_1242 += n
		} else if n == 0 && err != nil {
			panic("load err")
		}
	}

	return dm_build_1242
}

func (dm_build_1244 *Dm_build_1207) Dm_build_1243(dm_build_1245 io.Writer) *Dm_build_1207 {
	dm_build_1245.Write(dm_build_1244.dm_build_1208)
	return dm_build_1244
}

func (dm_build_1247 *Dm_build_1207) Dm_build_1246(dm_build_1248 bool) int {
	dm_build_1249 := len(dm_build_1247.dm_build_1208)
	dm_build_1247.dm_build_1214(1)

	if dm_build_1248 {
		return copy(dm_build_1247.dm_build_1208[dm_build_1249:], []byte{1})
	} else {
		return copy(dm_build_1247.dm_build_1208[dm_build_1249:], []byte{0})
	}
}

func (dm_build_1251 *Dm_build_1207) Dm_build_1250(dm_build_1252 byte) int {
	dm_build_1253 := len(dm_build_1251.dm_build_1208)
	dm_build_1251.dm_build_1214(1)

	return copy(dm_build_1251.dm_build_1208[dm_build_1253:], Dm_build_853.Dm_build_1028(dm_build_1252))
}

func (dm_build_1255 *Dm_build_1207) Dm_build_1254(dm_build_1256 int16) int {
	dm_build_1257 := len(dm_build_1255.dm_build_1208)
	dm_build_1255.dm_build_1214(2)

	return copy(dm_build_1255.dm_build_1208[dm_build_1257:], Dm_build_853.Dm_build_1031(dm_build_1256))
}

func (dm_build_1259 *Dm_build_1207) Dm_build_1258(dm_build_1260 int32) int {
	dm_build_1261 := len(dm_build_1259.dm_build_1208)
	dm_build_1259.dm_build_1214(4)

	return copy(dm_build_1259.dm_build_1208[dm_build_1261:], Dm_build_853.Dm_build_1034(dm_build_1260))
}

func (dm_build_1263 *Dm_build_1207) Dm_build_1262(dm_build_1264 uint8) int {
	dm_build_1265 := len(dm_build_1263.dm_build_1208)
	dm_build_1263.dm_build_1214(1)

	return copy(dm_build_1263.dm_build_1208[dm_build_1265:], Dm_build_853.Dm_build_1046(dm_build_1264))
}

func (dm_build_1267 *Dm_build_1207) Dm_build_1266(dm_build_1268 uint16) int {
	dm_build_1269 := len(dm_build_1267.dm_build_1208)
	dm_build_1267.dm_build_1214(2)

	return copy(dm_build_1267.dm_build_1208[dm_build_1269:], Dm_build_853.Dm_build_1049(dm_build_1268))
}

func (dm_build_1271 *Dm_build_1207) Dm_build_1270(dm_build_1272 uint32) int {
	dm_build_1273 := len(dm_build_1271.dm_build_1208)
	dm_build_1271.dm_build_1214(4)

	return copy(dm_build_1271.dm_build_1208[dm_build_1273:], Dm_build_853.Dm_build_1052(dm_build_1272))
}

func (dm_build_1275 *Dm_build_1207) Dm_build_1274(dm_build_1276 uint64) int {
	dm_build_1277 := len(dm_build_1275.dm_build_1208)
	dm_build_1275.dm_build_1214(8)

	return copy(dm_build_1275.dm_build_1208[dm_build_1277:], Dm_build_853.Dm_build_1055(dm_build_1276))
}

func (dm_build_1279 *Dm_build_1207) Dm_build_1278(dm_build_1280 float32) int {
	dm_build_1281 := len(dm_build_1279.dm_build_1208)
	dm_build_1279.dm_build_1214(4)

	return copy(dm_build_1279.dm_build_1208[dm_build_1281:], Dm_build_853.Dm_build_1052(math.Float32bits(dm_build_1280)))
}

func (dm_build_1283 *Dm_build_1207) Dm_build_1282(dm_build_1284 float64) int {
	dm_build_1285 := len(dm_build_1283.dm_build_1208)
	dm_build_1283.dm_build_1214(8)

	return copy(dm_build_1283.dm_build_1208[dm_build_1285:], Dm_build_853.Dm_build_1055(math.Float64bits(dm_build_1284)))
}

func (dm_build_1287 *Dm_build_1207) Dm_build_1286(dm_build_1288 []byte) int {
	dm_build_1289 := len(dm_build_1287.dm_build_1208)
	dm_build_1287.dm_build_1214(len(dm_build_1288))
	return copy(dm_build_1287.dm_build_1208[dm_build_1289:], dm_build_1288)
}

func (dm_build_1291 *Dm_build_1207) Dm_build_1290(dm_build_1292 []byte) int {
	return dm_build_1291.Dm_build_1258(int32(len(dm_build_1292))) + dm_build_1291.Dm_build_1286(dm_build_1292)
}

func (dm_build_1294 *Dm_build_1207) Dm_build_1293(dm_build_1295 []byte) int {
	return dm_build_1294.Dm_build_1262(uint8(len(dm_build_1295))) + dm_build_1294.Dm_build_1286(dm_build_1295)
}

func (dm_build_1297 *Dm_build_1207) Dm_build_1296(dm_build_1298 []byte) int {
	return dm_build_1297.Dm_build_1266(uint16(len(dm_build_1298))) + dm_build_1297.Dm_build_1286(dm_build_1298)
}

func (dm_build_1300 *Dm_build_1207) Dm_build_1299(dm_build_1301 []byte) int {
	return dm_build_1300.Dm_build_1286(dm_build_1301) + dm_build_1300.Dm_build_1250(0)
}

func (dm_build_1303 *Dm_build_1207) Dm_build_1302(dm_build_1304 string, dm_build_1305 string, dm_build_1306 *DmConnection) int {
	dm_build_1307 := Dm_build_853.Dm_build_1063(dm_build_1304, dm_build_1305, dm_build_1306)
	return dm_build_1303.Dm_build_1290(dm_build_1307)
}

func (dm_build_1309 *Dm_build_1207) Dm_build_1308(dm_build_1310 string, dm_build_1311 string, dm_build_1312 *DmConnection) int {
	dm_build_1313 := Dm_build_853.Dm_build_1063(dm_build_1310, dm_build_1311, dm_build_1312)
	return dm_build_1309.Dm_build_1293(dm_build_1313)
}

func (dm_build_1315 *Dm_build_1207) Dm_build_1314(dm_build_1316 string, dm_build_1317 string, dm_build_1318 *DmConnection) int {
	dm_build_1319 := Dm_build_853.Dm_build_1063(dm_build_1316, dm_build_1317, dm_build_1318)
	return dm_build_1315.Dm_build_1296(dm_build_1319)
}

func (dm_build_1321 *Dm_build_1207) Dm_build_1320(dm_build_1322 string, dm_build_1323 string, dm_build_1324 *DmConnection) int {
	dm_build_1325 := Dm_build_853.Dm_build_1063(dm_build_1322, dm_build_1323, dm_build_1324)
	return dm_build_1321.Dm_build_1299(dm_build_1325)
}

func (dm_build_1327 *Dm_build_1207) Dm_build_1326() byte {
	dm_build_1328 := Dm_build_853.Dm_build_946(dm_build_1327.dm_build_1208, dm_build_1327.dm_build_1209)
	dm_build_1327.dm_build_1209++
	return dm_build_1328
}

func (dm_build_1330 *Dm_build_1207) Dm_build_1329() int16 {
	dm_build_1331 := Dm_build_853.Dm_build_950(dm_build_1330.dm_build_1208, dm_build_1330.dm_build_1209)
	dm_build_1330.dm_build_1209 += 2
	return dm_build_1331
}

func (dm_build_1333 *Dm_build_1207) Dm_build_1332() int32 {
	dm_build_1334 := Dm_build_853.Dm_build_955(dm_build_1333.dm_build_1208, dm_build_1333.dm_build_1209)
	dm_build_1333.dm_build_1209 += 4
	return dm_build_1334
}

func (dm_build_1336 *Dm_build_1207) Dm_build_1335() int64 {
	dm_build_1337 := Dm_build_853.Dm_build_960(dm_build_1336.dm_build_1208, dm_build_1336.dm_build_1209)
	dm_build_1336.dm_build_1209 += 8
	return dm_build_1337
}

func (dm_build_1339 *Dm_build_1207) Dm_build_1338() float32 {
	dm_build_1340 := Dm_build_853.Dm_build_965(dm_build_1339.dm_build_1208, dm_build_1339.dm_build_1209)
	dm_build_1339.dm_build_1209 += 4
	return dm_build_1340
}

func (dm_build_1342 *Dm_build_1207) Dm_build_1341() float64 {
	dm_build_1343 := Dm_build_853.Dm_build_969(dm_build_1342.dm_build_1208, dm_build_1342.dm_build_1209)
	dm_build_1342.dm_build_1209 += 8
	return dm_build_1343
}

func (dm_build_1345 *Dm_build_1207) Dm_build_1344() uint8 {
	dm_build_1346 := Dm_build_853.Dm_build_973(dm_build_1345.dm_build_1208, dm_build_1345.dm_build_1209)
	dm_build_1345.dm_build_1209 += 1
	return dm_build_1346
}

func (dm_build_1348 *Dm_build_1207) Dm_build_1347() uint16 {
	dm_build_1349 := Dm_build_853.Dm_build_977(dm_build_1348.dm_build_1208, dm_build_1348.dm_build_1209)
	dm_build_1348.dm_build_1209 += 2
	return dm_build_1349
}

func (dm_build_1351 *Dm_build_1207) Dm_build_1350() uint32 {
	dm_build_1352 := Dm_build_853.Dm_build_982(dm_build_1351.dm_build_1208, dm_build_1351.dm_build_1209)
	dm_build_1351.dm_build_1209 += 4
	return dm_build_1352
}

func (dm_build_1354 *Dm_build_1207) Dm_build_1353(dm_build_1355 int) []byte {
	dm_build_1356 := Dm_build_853.Dm_build_1002(dm_build_1354.dm_build_1208, dm_build_1354.dm_build_1209, dm_build_1355)
	dm_build_1354.dm_build_1209 += dm_build_1355
	return dm_build_1356
}

func (dm_build_1358 *Dm_build_1207) Dm_build_1357() []byte {
	return dm_build_1358.Dm_build_1353(int(dm_build_1358.Dm_build_1332()))
}

func (dm_build_1360 *Dm_build_1207) Dm_build_1359() []byte {
	return dm_build_1360.Dm_build_1353(int(dm_build_1360.Dm_build_1326()))
}

func (dm_build_1362 *Dm_build_1207) Dm_build_1361() []byte {
	return dm_build_1362.Dm_build_1353(int(dm_build_1362.Dm_build_1329()))
}

func (dm_build_1364 *Dm_build_1207) Dm_build_1363(dm_build_1365 int) []byte {
	return dm_build_1364.Dm_build_1353(dm_build_1365)
}

func (dm_build_1367 *Dm_build_1207) Dm_build_1366() []byte {
	dm_build_1368 := 0
	for dm_build_1367.Dm_build_1326() != 0 {
		dm_build_1368++
	}
	dm_build_1367.Dm_build_1232(dm_build_1368, false, false)
	return dm_build_1367.Dm_build_1353(dm_build_1368)
}

func (dm_build_1370 *Dm_build_1207) Dm_build_1369(dm_build_1371 int, dm_build_1372 string, dm_build_1373 *DmConnection) string {
	return Dm_build_853.Dm_build_1100(dm_build_1370.Dm_build_1353(dm_build_1371), dm_build_1372, dm_build_1373)
}

func (dm_build_1375 *Dm_build_1207) Dm_build_1374(dm_build_1376 string, dm_build_1377 *DmConnection) string {
	return Dm_build_853.Dm_build_1100(dm_build_1375.Dm_build_1357(), dm_build_1376, dm_build_1377)
}

func (dm_build_1379 *Dm_build_1207) Dm_build_1378(dm_build_1380 string, dm_build_1381 *DmConnection) string {
	return Dm_build_853.Dm_build_1100(dm_build_1379.Dm_build_1359(), dm_build_1380, dm_build_1381)
}

func (dm_build_1383 *Dm_build_1207) Dm_build_1382(dm_build_1384 string, dm_build_1385 *DmConnection) string {
	return Dm_build_853.Dm_build_1100(dm_build_1383.Dm_build_1361(), dm_build_1384, dm_build_1385)
}

func (dm_build_1387 *Dm_build_1207) Dm_build_1386(dm_build_1388 string, dm_build_1389 *DmConnection) string {
	return Dm_build_853.Dm_build_1100(dm_build_1387.Dm_build_1366(), dm_build_1388, dm_build_1389)
}

func (dm_build_1391 *Dm_build_1207) Dm_build_1390(dm_build_1392 int, dm_build_1393 byte) int {
	return dm_build_1391.Dm_build_1426(dm_build_1392, Dm_build_853.Dm_build_1028(dm_build_1393))
}

func (dm_build_1395 *Dm_build_1207) Dm_build_1394(dm_build_1396 int, dm_build_1397 int16) int {
	return dm_build_1395.Dm_build_1426(dm_build_1396, Dm_build_853.Dm_build_1031(dm_build_1397))
}

func (dm_build_1399 *Dm_build_1207) Dm_build_1398(dm_build_1400 int, dm_build_1401 int32) int {
	return dm_build_1399.Dm_build_1426(dm_build_1400, Dm_build_853.Dm_build_1034(dm_build_1401))
}

func (dm_build_1403 *Dm_build_1207) Dm_build_1402(dm_build_1404 int, dm_build_1405 int64) int {
	return dm_build_1403.Dm_build_1426(dm_build_1404, Dm_build_853.Dm_build_1037(dm_build_1405))
}

func (dm_build_1407 *Dm_build_1207) Dm_build_1406(dm_build_1408 int, dm_build_1409 float32) int {
	return dm_build_1407.Dm_build_1426(dm_build_1408, Dm_build_853.Dm_build_1040(dm_build_1409))
}

func (dm_build_1411 *Dm_build_1207) Dm_build_1410(dm_build_1412 int, dm_build_1413 float64) int {
	return dm_build_1411.Dm_build_1426(dm_build_1412, Dm_build_853.Dm_build_1043(dm_build_1413))
}

func (dm_build_1415 *Dm_build_1207) Dm_build_1414(dm_build_1416 int, dm_build_1417 uint8) int {
	return dm_build_1415.Dm_build_1426(dm_build_1416, Dm_build_853.Dm_build_1046(dm_build_1417))
}

func (dm_build_1419 *Dm_build_1207) Dm_build_1418(dm_build_1420 int, dm_build_1421 uint16) int {
	return dm_build_1419.Dm_build_1426(dm_build_1420, Dm_build_853.Dm_build_1049(dm_build_1421))
}

func (dm_build_1423 *Dm_build_1207) Dm_build_1422(dm_build_1424 int, dm_build_1425 uint32) int {
	return dm_build_1423.Dm_build_1426(dm_build_1424, Dm_build_853.Dm_build_1052(dm_build_1425))
}

func (dm_build_1427 *Dm_build_1207) Dm_build_1426(dm_build_1428 int, dm_build_1429 []byte) int {
	return copy(dm_build_1427.dm_build_1208[dm_build_1428:], dm_build_1429)
}

func (dm_build_1431 *Dm_build_1207) Dm_build_1430(dm_build_1432 int, dm_build_1433 []byte) int {
	return dm_build_1431.Dm_build_1398(dm_build_1432, int32(len(dm_build_1433))) + dm_build_1431.Dm_build_1426(dm_build_1432+4, dm_build_1433)
}

func (dm_build_1435 *Dm_build_1207) Dm_build_1434(dm_build_1436 int, dm_build_1437 []byte) int {
	return dm_build_1435.Dm_build_1390(dm_build_1436, byte(len(dm_build_1437))) + dm_build_1435.Dm_build_1426(dm_build_1436+1, dm_build_1437)
}

func (dm_build_1439 *Dm_build_1207) Dm_build_1438(dm_build_1440 int, dm_build_1441 []byte) int {
	return dm_build_1439.Dm_build_1394(dm_build_1440, int16(len(dm_build_1441))) + dm_build_1439.Dm_build_1426(dm_build_1440+2, dm_build_1441)
}

func (dm_build_1443 *Dm_build_1207) Dm_build_1442(dm_build_1444 int, dm_build_1445 []byte) int {
	return dm_build_1443.Dm_build_1426(dm_build_1444, dm_build_1445) + dm_build_1443.Dm_build_1390(dm_build_1444+len(dm_build_1445), 0)
}

func (dm_build_1447 *Dm_build_1207) Dm_build_1446(dm_build_1448 int, dm_build_1449 string, dm_build_1450 string, dm_build_1451 *DmConnection) int {
	return dm_build_1447.Dm_build_1430(dm_build_1448, Dm_build_853.Dm_build_1063(dm_build_1449, dm_build_1450, dm_build_1451))
}

func (dm_build_1453 *Dm_build_1207) Dm_build_1452(dm_build_1454 int, dm_build_1455 string, dm_build_1456 string, dm_build_1457 *DmConnection) int {
	return dm_build_1453.Dm_build_1434(dm_build_1454, Dm_build_853.Dm_build_1063(dm_build_1455, dm_build_1456, dm_build_1457))
}

func (dm_build_1459 *Dm_build_1207) Dm_build_1458(dm_build_1460 int, dm_build_1461 string, dm_build_1462 string, dm_build_1463 *DmConnection) int {
	return dm_build_1459.Dm_build_1438(dm_build_1460, Dm_build_853.Dm_build_1063(dm_build_1461, dm_build_1462, dm_build_1463))
}

func (dm_build_1465 *Dm_build_1207) Dm_build_1464(dm_build_1466 int, dm_build_1467 string, dm_build_1468 string, dm_build_1469 *DmConnection) int {
	return dm_build_1465.Dm_build_1442(dm_build_1466, Dm_build_853.Dm_build_1063(dm_build_1467, dm_build_1468, dm_build_1469))
}

func (dm_build_1471 *Dm_build_1207) Dm_build_1470(dm_build_1472 int) byte {
	return Dm_build_853.Dm_build_1068(dm_build_1471.Dm_build_1497(dm_build_1472, 1))
}

func (dm_build_1474 *Dm_build_1207) Dm_build_1473(dm_build_1475 int) int16 {
	return Dm_build_853.Dm_build_1071(dm_build_1474.Dm_build_1497(dm_build_1475, 2))
}

func (dm_build_1477 *Dm_build_1207) Dm_build_1476(dm_build_1478 int) int32 {
	return Dm_build_853.Dm_build_1074(dm_build_1477.Dm_build_1497(dm_build_1478, 4))
}

func (dm_build_1480 *Dm_build_1207) Dm_build_1479(dm_build_1481 int) int64 {
	return Dm_build_853.Dm_build_1077(dm_build_1480.Dm_build_1497(dm_build_1481, 8))
}

func (dm_build_1483 *Dm_build_1207) Dm_build_1482(dm_build_1484 int) float32 {
	return Dm_build_853.Dm_build_1080(dm_build_1483.Dm_build_1497(dm_build_1484, 4))
}

func (dm_build_1486 *Dm_build_1207) Dm_build_1485(dm_build_1487 int) float64 {
	return Dm_build_853.Dm_build_1083(dm_build_1486.Dm_build_1497(dm_build_1487, 8))
}

func (dm_build_1489 *Dm_build_1207) Dm_build_1488(dm_build_1490 int) uint8 {
	return Dm_build_853.Dm_build_1086(dm_build_1489.Dm_build_1497(dm_build_1490, 1))
}

func (dm_build_1492 *Dm_build_1207) Dm_build_1491(dm_build_1493 int) uint16 {
	return Dm_build_853.Dm_build_1089(dm_build_1492.Dm_build_1497(dm_build_1493, 2))
}

func (dm_build_1495 *Dm_build_1207) Dm_build_1494(dm_build_1496 int) uint32 {
	return Dm_build_853.Dm_build_1092(dm_build_1495.Dm_build_1497(dm_build_1496, 4))
}

func (dm_build_1498 *Dm_build_1207) Dm_build_1497(dm_build_1499 int, dm_build_1500 int) []byte {
	return dm_build_1498.dm_build_1208[dm_build_1499 : dm_build_1499+dm_build_1500]
}

func (dm_build_1502 *Dm_build_1207) Dm_build_1501(dm_build_1503 int) []byte {
	dm_build_1504 := dm_build_1502.Dm_build_1476(dm_build_1503)
	return dm_build_1502.Dm_build_1497(dm_build_1503+4, int(dm_build_1504))
}

func (dm_build_1506 *Dm_build_1207) Dm_build_1505(dm_build_1507 int) []byte {
	dm_build_1508 := dm_build_1506.Dm_build_1470(dm_build_1507)
	return dm_build_1506.Dm_build_1497(dm_build_1507+1, int(dm_build_1508))
}

func (dm_build_1510 *Dm_build_1207) Dm_build_1509(dm_build_1511 int) []byte {
	dm_build_1512 := dm_build_1510.Dm_build_1473(dm_build_1511)
	return dm_build_1510.Dm_build_1497(dm_build_1511+2, int(dm_build_1512))
}

func (dm_build_1514 *Dm_build_1207) Dm_build_1513(dm_build_1515 int) []byte {
	dm_build_1516 := 0
	for dm_build_1514.Dm_build_1470(dm_build_1515) != 0 {
		dm_build_1515++
		dm_build_1516++
	}

	return dm_build_1514.Dm_build_1497(dm_build_1515-dm_build_1516, int(dm_build_1516))
}

func (dm_build_1518 *Dm_build_1207) Dm_build_1517(dm_build_1519 int, dm_build_1520 string, dm_build_1521 *DmConnection) string {
	return Dm_build_853.Dm_build_1100(dm_build_1518.Dm_build_1501(dm_build_1519), dm_build_1520, dm_build_1521)
}

func (dm_build_1523 *Dm_build_1207) Dm_build_1522(dm_build_1524 int, dm_build_1525 string, dm_build_1526 *DmConnection) string {
	return Dm_build_853.Dm_build_1100(dm_build_1523.Dm_build_1505(dm_build_1524), dm_build_1525, dm_build_1526)
}

func (dm_build_1528 *Dm_build_1207) Dm_build_1527(dm_build_1529 int, dm_build_1530 string, dm_build_1531 *DmConnection) string {
	return Dm_build_853.Dm_build_1100(dm_build_1528.Dm_build_1509(dm_build_1529), dm_build_1530, dm_build_1531)
}

func (dm_build_1533 *Dm_build_1207) Dm_build_1532(dm_build_1534 int, dm_build_1535 string, dm_build_1536 *DmConnection) string {
	return Dm_build_853.Dm_build_1100(dm_build_1533.Dm_build_1513(dm_build_1534), dm_build_1535, dm_build_1536)
}
