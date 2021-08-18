<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('user', 'Api\UserController@getAllUsers');
Route::get('user/{id}', 'Api\UserController@getUser');
Route::post('users', 'Api\UserController@createUser');
Route::put('user/{id}', 'Api\UserController@updateUser');
Route::delete('user/{id}','Api\UserController@deleteUser');

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});
