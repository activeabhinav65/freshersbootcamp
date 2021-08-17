<?php
$newArray = array();
function getFlatArray($array) {
    return flattenArray($array);
}

function flattenArray($array) {
    foreach ($array as $value) {
        if (is_array($value)) {
            flattenArray($value);
        }
        else {
            global $newArray;
            array_push($newArray, $value);
        }
    }
}
getFlatArray([2,3, [4,5], [6,7],8]);
print_r($newArray);
