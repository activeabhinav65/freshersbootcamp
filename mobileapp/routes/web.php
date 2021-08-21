<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\UsersController;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/welcome', function () {
    $name = 'Welcome User';
    return view('welcome',compact('name'));
});

Route::get('/token/', function () {
    return csrf_token();
});
Route::get('/insert',[UsersController::class, 'insertdata']);
Route::get('/mobileusers/{first_name}',[UsersController::class, 'getuserbyname']);
Route::get('/mobileusers1/{email}',[UsersController::class, 'getuserbyemail']);
Route::get('/mobileusers2/{mobile_number}',[UsersController::class, 'getuserbymobilenumber']);
Route::get('/mobileusers',[UsersController::class, 'getallusers']);
Route::get('/mobileusers3/{id}',[UsersController::class, 'getuserbyid']);
Route::delete('/delete/{id}',[UsersController::class, 'deleteuser']);
