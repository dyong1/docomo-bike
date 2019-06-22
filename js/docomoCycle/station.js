const request = require("request-promise-native")
const iconv = require("iconv-lite")

const EVENT_NO = 25701
const API_ENDPOINT = "https://tcc.docomo-cycle.jp/cycle/TYO/cs_web_main.php"
const RESPONSE_ENCODING = "Shift_JIS"

module.exports.stations = {
    roppongiHills: 10082,
    nishiSimbashi1Chome: 10070,
}

module.exports.countBikesAt = async function countBikeAt(auth, stationId) {
    const station = await getStation(auth, stationId)
    return station.totalBikes
}

module.exports.getStation = async function getStation(auth, stationId) {
    const options = {
        form: {
            "EventNo": EVENT_NO,
            "MemberID": auth.userId,
            "SessionID": auth.sessionKey,
            "ParkingID": stationId,
            "GetInfoNum": 20,
            "GetInfoTopNum": 1,
            "UserID": "TYO", // Required, don't know why TYO is okay
            "ParkingEntID": "TYO", // Required, don't know why TYO is okay
        }
    }
    const res = await request.post(API_ENDPOINT, options)
    const decoded = iconv.decode(Buffer.from(res, "binary"), RESPONSE_ENCODING)
    const lines = decoded.split("\n")
    const bikeLines = lines.filter(line => /<a.*class=".*cycle_list_btn.*".*/.test(line))
    const stationNameHeaderLine = lines.find(line => line.includes("Port name"))
    return {
        name: lines[lines.indexOf(stationNameHeaderLine)+2],
        totalBikes: bikeLines.length
    }
}
