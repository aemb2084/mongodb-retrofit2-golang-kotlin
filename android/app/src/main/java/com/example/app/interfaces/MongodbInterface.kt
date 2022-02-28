package com.example.app.interfaces

import com.example.app.models.Pokemon
import retrofit2.Response
import retrofit2.http.GET
import retrofit2.http.Url

interface MongodbInterface {
    @GET
    suspend fun listDb(@Url url:String): Response<List<String>>

    @GET
    suspend fun getDocument(@Url url:String): Response<List<Pokemon>>
}