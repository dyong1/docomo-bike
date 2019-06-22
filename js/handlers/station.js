const { login } = require("../docomoCycle/login")
const { stations, getStation } = require("../docomoCycle/station")

module.exports.getStation = async function (req, res) {
    try {
        const userId = req.query.userId
        const password = req.query.password
        const stationId = stations[req.params.stationName]
        validateInput({
            userId,
            password,
            stationId,
        })

        const auth = await login(userId, password)
        const station = await getStation(auth, stationId)
        res.status(200).json(station)
    } catch (e) {
        res.status(500).json({
            reason: e.toString(),
        })
    }
}

function validateInput() {
    //TODO:
}
