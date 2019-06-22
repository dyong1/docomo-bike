const request = require("request-promise-native")

const EVENT_NO = 21401
const API_ENDPOINT = "https://tcc.docomo-cycle.jp/cycle/TYO/cs_web_main.php"

module.exports.login = async function login(userId, password) {
    const options = {
        form: {
            "EventNo": EVENT_NO,
            "MemberID": userId,
            "Password": password,
        }
    }
    const res = await request.post(API_ENDPOINT, options)
    const sessionLine = res.split("\n").filter(line => line.includes("SessionID"))[0]
    if (!sessionLine) {
        throw new Error("No session line is found in the response")
    }
    const matches = /value="(.+)"/.exec(sessionLine)
    return {
        userId: userId,
        sessionKey: matches[1]
    }
}
