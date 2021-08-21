<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Illuminate\Support\Facades\DB;
use App\User;
use Illuminate\Database\QueryException;
use Nette\Schema\ValidationException;
use Illuminate\Support\Facades\Validator;
use App\Services\UserService;

class UsersController extends Controller
{
    public function insertdata(){
        return view('user_create');
    }
    public function createUser(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'first_name' => 'required',
            'last_name' => 'required',
            'email' => 'required|email',
            'mobile_number'=>'required|regex:/^[6-9][0-9]{9}$/'
        ]);
        if ($validator->fails()) {
            return response()->json($validator->errors(),Response::HTTP_BAD_REQUEST);
        }
        $firstname = $request->input('first_name');
        $lastname = $request->input('last_name');
        $mobilenumber = $request->input('mobile_number');
        $email = $request->input('email');

        $data=array('first_name'=>$firstname,'last_name'=>$lastname,'mobile_number'=>$mobilenumber,'email'=>$email);
        try{
            $id = DB::table('mobileusers')->insertGetId($data);
        }
        catch (QueryException $e){
            return response()->json($e->errorInfo, Response::HTTP_BAD_REQUEST);

        }
        return response()->json([
            'id' => $id ,
            'first_name' => $firstname,
            'last_name'  => $lastname,
            'mobile_number'  => $mobilenumber,
            'email'  => $email,
        ]);


        //echo '<a href = "/insert">Click Here</a> to go back to insert more data.';
    }

        public function getUserByName($firstname)
    {
        $users = DB::table('mobileusers')->where('first_name' , $firstname )->get();
        if (count($users) == 0 )
        {
            return response()->json([
                'message' => "No user found with this name-$firstname"
            ]);
        }
        return response()->json([
            "id" => $users[0]->id,
            "first_name" => $users[0]->first_name,
            "last_name" => $users[0]->last_name,
            "mobile_number" => $users[0]->mobile_number,
            "email" => $users[0]->email

        ]);

        //return view('user_view',['users' =>$users]);
    }
    public function getUserByEmail($email)
    {
        $users = DB::table('mobileusers')->where('email' , $email )->get();
        if (count($users) == 0 )
        {
            return response()->json([
                'message' => "No user found with this email-$email"
            ]);
        }
        return response()->json([
            "id" => $users[0]->id,
            "first_name" => $users[0]->first_name,
            "last_name" => $users[0]->last_name,
            "mobile_number" => $users[0]->mobile_number,
            "email" => $users[0]->email

        ]);

        //return view('user_view',['users' =>$users]);
    }
    public function getUserByMobileNumber($mobilenumber)
    {
        $users = DB::table('mobileusers')->where('mobile_number' , $mobilenumber )->get();
        if (count($users) == 0 )
        {
            return response()->json([
                'message' => "No user found with this mobile number-$mobilenumber"
            ]);
        }
        return response()->json([
            "id" => $users[0]->id,
            "first_name" => $users[0]->first_name,
            "last_name" => $users[0]->last_name,
            "mobile_number" => $users[0]->mobile_number,
            "email" => $users[0]->email

        ]);

        //return view('user_view',['users' =>$users]);
    }


    public function getUserById($id)
    {
        $users = DB::table('mobileusers')->where('id' , $id )->get();
        if (count($users) == 0 )
        {
            return response()->json([
                'message' => "No user found with this id-$id"
            ]);
        }
        return response()->json([
            "id" => $users[0]->id,
            "first_name" => $users[0]->first_name,
            "last_name" => $users[0]->last_name,
            "mobile_number" => $users[0]->mobile_number,
            "email" => $users[0]->email

        ]);

        //return view('user_view',['users' =>$users]);
    }

    public function getAllUsers()
    {
        $users = DB::table('mobileusers')
            ->get();

        return response()->json([
            'All users' => $users
        ]);

        //return view('user_view',['users'=>$users]);
    }

    public function deleteUserByMobileNumber($mobilenumber)
    {
        $count = DB::table('mobileusers')
            ->where('mobile_number' , $mobilenumber)
            ->count();
        if ($count == 0)
        {
            return response()->json([
                'message' => "User with mobile number-$mobilenumber does not exists"
            ]);
        }
        DB::table('mobileusers')
            ->where('mobile_number' , $mobilenumber)
            ->delete();

        return response()->json([
            'message' => "user with mobile number-$mobilenumber deleted"
        ]);
    }


public function deleteUserByFirstName($firstname)
{
    $count = DB::table('mobileusers')
        ->where('first_name' , $firstname)
        ->count();
    if ($count == 0)
    {
        return response()->json([
            'message' => "User with first name-$firstname does not exists"
        ]);
    }
    DB::table('mobileusers')
        ->where('first_name' , $firstname)
        ->delete();

    return response()->json([
        'message' => "user with first name-$firstname deleted"
    ]);
}
    public function deleteUserByEmail($email)
    {
        $count = DB::table('mobileusers')
            ->where('email' , $email)
            ->count();
        if ($count == 0)
        {
            return response()->json([
                'message' => "User with email-$email does not exists"
            ]);
        }
        DB::table('mobileusers')
            ->where('email' , $email)
            ->delete();

        return response()->json([
            'message' => "user with email-$email deleted"
        ]);
    }
}

