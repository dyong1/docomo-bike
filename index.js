const express = require("express")
const bodyParser = require("body-parser")

const { getStation } = require("./handlers/station")

const app = express()
app.use(bodyParser.json())
app.get("/stations/:stationName", getStation)

app.listen(process.env.PORT)
