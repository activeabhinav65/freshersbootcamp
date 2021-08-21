<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\UsersController;
/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});
Route::post('/mobileusers',  [UsersController::class, 'createUser']); //....... add user
Route::get('/mobileusers1/{first_name}',[UsersController::class, 'getUserByName']); //....... fetch user by name
Route::get('/mobileusers2/{email}',[UsersController::class, 'getUserByEmail']); //....... fetch user by email
Route::get('/mobileusers3/{mobile_number}',[UsersController::class, 'getUserByMobileNumber']); //....... fetch user by email
Route::get('/mobileusers',[UsersController::class, 'getAllUsers']); //....... fetch all users
Route::delete('/mobileusers/{mobile_number}',[UsersController::class, 'deleteUserByMobileNumber']); //....... delete user by mobile number
Route::delete('/mobileusers1/{first_name}',[UsersController::class, 'deleteUserByFirstName']); //....... delete user by first name
Route::delete('/mobileusers2/{email}',[UsersController::class, 'deleteUserByEmail']); //....... delete user by eamil

Route::get('/mobileusers/{id}',[UsersController::class, 'getUserById']); //....... fetch user by id


