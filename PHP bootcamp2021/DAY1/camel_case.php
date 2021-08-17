<?php
function camelize($input, $separator = '_') {
    return lcfirst(str_replace($separator, '', ucwords($input, $separator)));
}
$array=["snake_case","camel_case"];
$newArray=array();
foreach ($array as $value)
    array_push($newArray,camelize($value));
print_r($newArray);