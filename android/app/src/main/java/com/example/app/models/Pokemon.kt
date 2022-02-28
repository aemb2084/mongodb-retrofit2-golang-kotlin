package com.example.app.models

import com.google.gson.annotations.SerializedName

data class Evolution(
    @SerializedName("num") var Num: String,
    @SerializedName("name") var Name: String
)

data class Pokemon (
    @SerializedName("id") var Id: String,
    @SerializedName("num") var Num: String,
    @SerializedName("name") var Name: String,
    @SerializedName("image") var Image: String,
    @SerializedName("type") var Type: Array<String>,
    @SerializedName("heigth") var Heigth: String,
    @SerializedName("weigth")var Weigth: String,
    @SerializedName("candy") var Candy: String,
    @SerializedName("candy_count") var CandyCount: Int,
    @SerializedName("egg") var Egg: String,
    @SerializedName("spawn_chance") var SpawnChance: Double,
    @SerializedName("avg_spawns") var AvgSpawns: Double,
    @SerializedName("spawn_time") var SpawnTime: String,
    @SerializedName("multipliers") var Multipliers: Array<Double>,
    @SerializedName("weaknesses") var Weaknesses: Array<String>,
    @SerializedName("next_evolution") var NextEvolution: Array<Evolution>,
    @SerializedName("prev_evolution")var PrevEvolution: Array<Evolution>
)