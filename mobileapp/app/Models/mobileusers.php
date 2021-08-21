<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class mobileusers extends Model
{

    protected $fillable = [
        'first_name',
        'last_name',
        'mobile_number',
        'email'

    ];
}
