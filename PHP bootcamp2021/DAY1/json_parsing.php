<?php
function getallnames($string)
{
    $input = json_decode($string, true);
    $names = array();
    $ages = array();
    $cities = array();
    foreach ($input["players"] as $value) {
        array_push($names, $value['name']);
        array_push($ages, $value['age']);
        array_push($cities, $value['address']['city']);

    }
    print_r($names);
    print_r($ages);
    print_r($cities);
    $uniqueNames=array_unique($names);
    print_r($uniqueNames);
    $max = max($ages);
    $oldest = array();
    foreach($ages as $idx=>$age){
        if ($age== $max){
            array_push($oldest,$names[$idx]);
        }
    }
    print_r($oldest);

}




$string= "{\"players\":[{\"name\":\"Ganguly\",\"age\":45,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Dravid\",\"age\":45,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Dhoni\",\"Age\":37,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Virat\",\"age\":35,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Jadeja\",\"age\":35,\"address\":{\"city\":\"Hyderabad\"}},{\"name\":\"Jadeja\",\"age\":35,\"address\":{\"city\":\"Hyderabad\"}}]}";
getallnames($string);